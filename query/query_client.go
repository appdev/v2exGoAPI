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
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Mobile Safari/537.36")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	req.Header.Add("Referer", url)
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("cookie", `_ga=GA1.2.1172485124.1526525960; V2EX_REFERRER="2|1:0|10:1553751586|13:V2EX_REFERRER|12:VmljdG9yMjE1|62b2b38fd34b3069747e9466082890913f673897ce3ef39ffd47a1ef714e5005"; A2="2|1:0|10:1553772217|2:A2|56:Y2ZjMmJkMTM5MTg4MGU4MTNlMGUxYmM2MjM2NTUzYzkxMDFhZTEyMA==|974b81fef5113823b03d54aeff801473ed9b76b56979082de19b7cf588eaaaa6"; V2EX_LANG=zhcn; PB3_SESSION="2|1:0|10:1553839299|11:PB3_SESSION|40:djJleDoyMjIuMjEyLjE4NC4xNDY6MTU4NDIwNjI=|f742180f79aaeb97803bf2cf92e99af1606e5a783ae58a15c2ca17526c4f9d9f"; V2EX_TAB="2|1:0|10:1553843767|8:V2EX_TAB|4:YWxs|bdeb62def53bcf58d9a3361d7114e5a6d0b9d393badec243a2989e18a39b3475"`)


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
