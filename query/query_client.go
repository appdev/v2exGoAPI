package query

import (
	"net/http"
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	"log"
)

var BaseUrl = "https://www.v2ex.com/"

func RequestClient(url string) (*goquery.Document, string) {
	log.Println(url)

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}
	client := &http.Client{Transport: transCfg}
	req, _ := http.NewRequest("GET", url, nil)
	//手机页面
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Mobile Safari/537.36")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	req.Header.Add("Referer", url)
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")

	res, err := client.Do(req)

	if err != nil {
		return nil, "网络请求出错"
	} else {
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return nil, "网络请求出错"
		} else {
			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				return nil, "网络请求出错"
			} else {
				return doc, ""
			}
		}
	}
}
