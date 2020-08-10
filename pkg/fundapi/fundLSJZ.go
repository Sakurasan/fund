package fundapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type FundLSJZ struct {
	Data struct {
		LSJZList []struct {
			FSRQ      string      `json:"FSRQ,omitempty"` //净值日期*
			DWJZ      string      `json:"DWJZ,omitempty"` //单位净值 (当前净值)*
			LJJZ      string      `json:"LJJZ,omitempty"` //累计净值*
			SDATE     interface{} `json:"SDATE,omitempty"`
			ACTUALSYI string      `json:"ACTUALSYI,omitempty"`
			NAVTYPE   string      `json:"NAVTYPE,omitempty"`
			JZZZL     string      `json:"JZZZL,omitempty"` //净值增长率*
			SGZT      string      `json:"SGZT,omitempty"`  //申购状态
			SHZT      string      `json:"SHZT,omitempty"`  //赎回状态
			FHFCZ     string      `json:"FHFCZ,omitempty"`
			FHFCBZ    string      `json:"FHFCBZ,omitempty"`
			DTYPE     interface{} `json:"DTYPE,omitempty"`
			FHSP      string      `json:"FHSP,omitempty"`
		} `json:"LSJZList,omitempty"`
		FundType  string      `json:"FundType,omitempty"`
		SYType    interface{} `json:"SYType,omitempty"`
		IsNewType bool        `json:"isNewType,omitempty"`
		Feature   string      `json:"Feature,omitempty"`
	} `json:"Data,omitempty"`
	ErrCode    int         `json:"ErrCode,omitempty"`
	ErrMsg     interface{} `json:"ErrMsg,omitempty"`
	TotalCount int         `json:"TotalCount,omitempty"`
	Expansion  interface{} `json:"Expansion,omitempty"`
	PageSize   int         `json:"PageSize,omitempty"`
	PageIndex  int         `json:"PageIndex,omitempty"`
}

/*
历史净值数据

pageSize 条数 default 30
*/
func (t *FundLSJZ) GetAllHistoryList(pageSize int) error {
	if pageSize == 0 {
		pageSize = 30
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	fundurl := fmt.Sprintf("http://api.fund.eastmoney.com/f10/lsjz?callback=&fundCode=100038&pageIndex=1&pageSize=%d&startDate=&endDate=", pageSize)
	req, err := http.NewRequest("GET", fundurl, nil)
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
	req.Header.Set("Referer", "http://fundf10.eastmoney.com/")
	// req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,ja;q=0.5")
	// req.Header.Set("Cookie", os.ExpandEnv("st_si=40821216490856; st_asi=delete; EMFUND1=null; EMFUND2=null; EMFUND3=null; EMFUND4=null; EMFUND5=null; EMFUND0=null; EMFUND6=08-09%2000%3A32%3A29@%23%24%u8BFA%u5B89%u6210%u957F%u6DF7%u5408@%23%24320007; EMFUND7=08-09%2018%3A11%3A04@%23%24%u5BCC%u56FD%u6CAA%u6DF1300%u6307%u6570%u589E%u5F3A@%23%24100038; EMFUND8=08-09%2001%3A22%3A16@%23%24%u4E07%u5BB6%u4EBA%u5DE5%u667A%u80FD%u6DF7%u5408@%23%24006281; EMFUND9=08-09 18:11:50@#$%u6613%u65B9%u8FBE%u56FD%u9632%u519B%u5DE5%u6DF7%u5408@%23%24001475; st_pvi=15759433755952; st_sp=2020-02-14%2013%3A24%3A29; st_inirUrl=https%3A%2F%2Fwww.baidu.com%2Flink; st_sn=65; st_psi=20200809181228116-0-4731647508"))

	resp, err := client.Do(req)
	if err != nil {
		// handle err
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(t)

}
