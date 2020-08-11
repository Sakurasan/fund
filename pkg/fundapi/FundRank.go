package fundapi

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
获取基金排行

sc
	rzdf 日增长率
	zzf 近一周
	3yzf 近三月
	6yzf
	1nzf 近一年
	2nzf
	3nzf
	jnzf 今年来
	lnzf 成立来
	qjzf 自定义
		sd: 2019-08-11
		ed: 2020-08-11

*/
func GetfundRank() error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", "http://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=all&rs=&gs=0&sc=rzdf&st=desc&sd=&ed=&qdii=&tabSubtype=,,,,,&pi=1&pn=50&dx=1&v=0.25340474789879774", nil)
	if err != nil {
		// handle err
		return err
	}
	// req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Pragma", "no-cache")
	// req.Header.Set("Cache-Control", "no-cache")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	// req.Header.Set("Dnt", "1")
	// req.Header.Set("Accept", "*/*")
	req.Header.Set("Referer", "http://fund.eastmoney.com/data/fundranking.html")
	// req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,ja;q=0.5")

	resp, err := client.Do(req)
	if err != nil {
		// handle err
		return err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data[bytes.Index(data, []byte("[")) : bytes.Index(data, []byte("]"))+1]))

	return nil
}
