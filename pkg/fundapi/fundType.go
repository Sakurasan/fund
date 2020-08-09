package fundapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"fund/pkg/utils"
	"io/ioutil"
	"net/http"
	"os"
)

type Fund struct {
	FundNum    string `json:"fundNum,omitempty"`
	FundName   string `json:"fundName,omitempty"`
	FundNameJP string
	FundType   string
	D0         float64 `json:"d0,omitempty"`
	D1         float64 `json:"d1,omitempty"`
	D2         float64 `json:"d2,omitempty"`
	D3         float64 `json:"d3,omitempty"`
	D4         float64 `json:"d4,omitempty"`
	D5         float64 `json:"d5,omitempty"`
	D6         float64 `json:"d6,omitempty"`
	D7         float64 `json:"d7,omitempty"`

	Low  float64
	High float64

	Delta    int
	DeltaSum float64

	FArray []*FundItem
}

type FundItem struct {
	FSRQ  string `json:"FSRQ"`  //净值日期
	DWJZ  string `json:"DWJZ"`  //单位净值 (当前净值)
	LJJZ  string `json:"LJJZ"`  //累计净值
	JZZZL string `json:"JZZZL"` //净值增长率
}

func GetAllFundData() ([]*Fund, error) {
	fundMap, err := GetAllFund()
	if err != nil {
		return nil, err
	}
	var allFund []*Fund

	for k, v := range fundMap {
		f := new(Fund)
		f.FundNum = k
		f.FundName = v[2]
		f.FundNameJP = v[1]
		f.FundType = v[3]
		allFund = append(allFund, f)
	}

	return allFund, nil
}

func GetAllFund() (map[string][]string, error) {
	var data []byte
	var err error

	if !utils.IsFileExist("fund.json") {
		rsp, err := http.DefaultClient.Get("http://fund.eastmoney.com/js/fundcode_search.js")
		if err != nil {
			return nil, err
		}
		if rsp.StatusCode != http.StatusOK {
			return nil, errors.New(fmt.Sprintf("Err StatusCode:%d", rsp.StatusCode))
		}
		data, _ = ioutil.ReadAll(rsp.Body)
		// var r = [["0
		ioutil.WriteFile("fund.json", data[bytes.Index(data, []byte{'['}):len(data)-1], os.ModePerm)
	} else {
		data, err = ioutil.ReadFile("fund.json")
		if err != nil {
			return nil, err
		}
	}

	var fss [][]string
	err = json.Unmarshal(data, &fss)
	if err != nil {
		return nil, err
	}

	fundMap := map[string][]string{}
	for _, v := range fss {
		fundMap[v[0]] = v
	}
	return fundMap, nil
}
