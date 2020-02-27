package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	url_parse = `http://fundgz.1234567.com.cn/js/%s.js`
	url_list  []string
	wg        sync.WaitGroup
)

type Config struct {
	Id []string `yaml:"id"`
}

func main() {
	conf := Config{}
	file, _ := ioutil.ReadFile("config.yaml")
	err := yaml.Unmarshal(file, &conf)
	if err != nil {
		panic(err)
	}
	for _, data := range conf.Id {
		url_list = append(url_list, fmt.Sprintf(url_parse, data[strings.Index(data, "[")+1:len(data)-1]))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := justdoit()
		w.Write([]byte(data))
	})

	http.ListenAndServe(":12345", mux)
}

func justdoit() []byte {
	var msgList []string
	for _, url := range url_list {
		wg.Add(1)
		go func(url string) {
			resp, err := browser(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			msgList = append(msgList, resp)
			wg.Done()
		}(url)
		wg.Wait()
	}
	data, _ := json.Marshal(msgList)
	return data
}

func browser(url string) (string, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Accept-Charset", "UTF-8,*;q=0.5")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:60.0) Gecko/20100101 Firefox/60.0")
	req.Header.Add("referer", "http://fund.eastmoney.com")

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return string(data), nil

}

// Postman GET&POST tool
func Postman(method, url string, body io.Reader) (string, error) {
	req := new(http.Request)
	var err error
	switch method {
	case http.MethodGet:
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return "", nil
		}
	case http.MethodPost:
		req, err = http.NewRequest(method, url, body)
		if err != nil {
			return "", nil
		}
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
