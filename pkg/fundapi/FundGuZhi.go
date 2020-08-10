package fundapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FundGuZhi struct {
	Data struct {
		TypeStr  string `json:"typeStr,omitempty"`
		Sort     string `json:"sort,omitempty"`
		SortType string `json:"sortType,omitempty"`
		Canbuy   string `json:"canbuy,omitempty"` //开房购买
		Gzrq     string `json:"gzrq,omitempty"`   //估值日期
		Gxrq     string `json:"gxrq,omitempty"`   //更新日期*
		List     []struct {
			Bzdm        string      `json:"bzdm,omitempty"` //基金代码*
			ListTexch   string      `json:"ListTexch,omitempty"`
			FScaleType  string      `json:"FScaleType,omitempty"`
			PLevel      float64     `json:"PLevel,omitempty"`
			JJGSID      string      `json:"JJGSID,omitempty"`
			IsExchg     string      `json:"IsExchg,omitempty"`
			FType       string      `json:"FType,omitempty"` //基金类型
			Discount    float64     `json:"Discount,omitempty"`
			Rate        string      `json:"Rate,omitempty"`
			Feature     string      `json:"feature,omitempty"`
			Fundtype    string      `json:"fundtype,omitempty"`
			Gxrq        string      `json:"gxrq,omitempty"`
			Jjlx3       interface{} `json:"jjlx3,omitempty"`
			IsListTrade string      `json:"IsListTrade,omitempty"`
			Jjlx2       interface{} `json:"jjlx2,omitempty"`
			Shzt        interface{} `json:"shzt,omitempty"`
			Sgzt        string      `json:"sgzt,omitempty"` //申购状态
			Isbuy       string      `json:"isbuy,omitempty"`
			Gzrq        string      `json:"gzrq,omitempty"`
			Gspc        string      `json:"gspc,omitempty"`
			Gsz         string      `json:"gsz,omitempty"`   //估算值*
			Gszzl       string      `json:"gszzl,omitempty"` //估算增长率*
			Jzzzl       string      `json:"jzzzl,omitempty"`
			Dwjz        string      `json:"dwjz,omitempty"` //单位净值*
			Gbdwjz      string      `json:"gbdwjz,omitempty"`
			Jjjcpy      string      `json:"jjjcpy,omitempty"` //基金简称拼音
			Jjjc        string      `json:"jjjc,omitempty"`   //基金简称
			Jjlx        interface{} `json:"jjlx,omitempty"`
			Gszzlcolor  string      `json:"gszzlcolor,omitempty"`
			Jzzzlcolor  string      `json:"jzzzlcolor,omitempty"`
		} `json:"list,omitempty"`
	} `json:"Data,omitempty"`
	ErrCode    int         `json:"ErrCode,omitempty"`
	ErrMsg     interface{} `json:"ErrMsg,omitempty"`
	TotalCount int         `json:"TotalCount,omitempty"`
	Expansion  interface{} `json:"Expansion,omitempty"`
	PageSize   int         `json:"PageSize,omitempty"`
	PageIndex  int         `json:"PageIndex,omitempty"`
}

/*
历史估值数据

sortType default ”3“
	1 fundCode 基金代码
	2 fundGSZ 估算值
	3 gszzl 估算增长率

orderby  default desc
	"desc" or "asc"
*/
func GetFundGZList(sortType, orderby string) error {
	if sortType == "" {
		sortType = "3"
	}
	if orderby == "" {
		orderby = "desc"
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	// "http://api.fund.eastmoney.com/FundGuZhi/GetFundGZList?type=1&sort=3&orderType=asc&canbuy=0&pageIndex=1&pageSize=200&callback=&_=1597046961112"

	gzURL := fmt.Sprintf("http://api.fund.eastmoney.com/FundGuZhi/GetFundGZList?type=1&sort=%s&orderType=%s&canbuy=0&pageIndex=1&pageSize=200&callback=&_=1597045344249", sortType, orderby)
	req, err := http.NewRequest("GET", gzURL, nil)
	if err != nil {
		// handle err
		return err
	}
	// req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	// req.Header.Set("Accept", "*/*")
	req.Header.Set("Referer", "http://fund.eastmoney.com/fundguzhi.html")
	// req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,ja;q=0.5")

	resp, err := client.Do(req)
	if err != nil {
		// handle err
		return err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	var fundList FundGuZhi
	if err := json.Unmarshal(data, &fundList); err != nil {
		return err
	}

	fmt.Println(len(fundList.Data.List))
	return nil
}
