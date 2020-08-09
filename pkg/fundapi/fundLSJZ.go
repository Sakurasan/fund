package fundapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

type FundLSJZ struct {
	Data struct {
		LSJZList []struct {
			FSRQ      string      `json:"FSRQ"` //净值日期
			DWJZ      string      `json:"DWJZ"` //单位净值 (当前净值)
			LJJZ      string      `json:"LJJZ"` //累计净值
			SDATE     interface{} `json:"SDATE"`
			ACTUALSYI string      `json:"ACTUALSYI"`
			NAVTYPE   string      `json:"NAVTYPE"`
			JZZZL     string      `json:"JZZZL"` //净值增长率
			SGZT      string      `json:"SGZT"`  //申购状态
			SHZT      string      `json:"SHZT"`  //赎回状态
			FHFCZ     string      `json:"FHFCZ"`
			FHFCBZ    string      `json:"FHFCBZ"`
			DTYPE     interface{} `json:"DTYPE"`
			FHSP      string      `json:"FHSP"`
		} `json:"LSJZList"`
		FundType  string      `json:"FundType"`
		SYType    interface{} `json:"SYType"`
		IsNewType bool        `json:"isNewType"`
		Feature   string      `json:"Feature"`
	} `json:"Data"`
	ErrCode    int         `json:"ErrCode"`
	ErrMsg     interface{} `json:"ErrMsg"`
	TotalCount int         `json:"TotalCount"`
	Expansion  interface{} `json:"Expansion"`
	PageSize   int         `json:"PageSize"`
	PageIndex  int         `json:"PageIndex"`
}

func (t *FundLSJZ) Get() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	fundurl := fmt.Sprintf("http://api.fund.eastmoney.com/f10/lsjz?callback=jQuery183019912269895979295_1596967948165&fundCode=100038&pageIndex=1&pageSize=20&startDate=&endDate=&_=1596968113385")
	req, err := http.NewRequest("GET", fundurl, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Referer", "http://fundf10.eastmoney.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,ja;q=0.5")
	// req.Header.Set("Cookie", os.ExpandEnv("st_si=40821216490856; st_asi=delete; EMFUND1=null; EMFUND2=null; EMFUND3=null; EMFUND4=null; EMFUND5=null; EMFUND0=null; EMFUND6=08-09%2000%3A32%3A29@%23%24%u8BFA%u5B89%u6210%u957F%u6DF7%u5408@%23%24320007; EMFUND7=08-09%2018%3A11%3A04@%23%24%u5BCC%u56FD%u6CAA%u6DF1300%u6307%u6570%u589E%u5F3A@%23%24100038; EMFUND8=08-09%2001%3A22%3A16@%23%24%u4E07%u5BB6%u4EBA%u5DE5%u667A%u80FD%u6DF7%u5408@%23%24006281; EMFUND9=08-09 18:11:50@#$%u6613%u65B9%u8FBE%u56FD%u9632%u519B%u5DE5%u6DF7%u5408@%23%24001475; st_pvi=15759433755952; st_sp=2020-02-14%2013%3A24%3A29; st_inirUrl=https%3A%2F%2Fwww.baidu.com%2Flink; st_sn=65; st_psi=20200809181228116-0-4731647508"))

	resp, err := client.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}
