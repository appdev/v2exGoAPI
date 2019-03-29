package query

import (
	"../module"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
)

func ItemByTab(tab string, page string) (json interface{}, error string) {

	// Create Http Client
	var requestUrl = ""
	if page == "" {
		requestUrl = BaseUrl + "?tab=" + tab
	} else {
		requestUrl = BaseUrl + tab + "?p=" + page
	}
	doc, err := RequestClient(requestUrl)
	if err != "" {
		return "", err
	}

	// Find the review itemsclass="cell"
	var domList = doc.Find("div[class*=item]")
	//切片方式
	var list [] module.Items
	//list:=list.New()
	domList.Each(func(i int, selection *goquery.Selection) {
		selection.Find("tr").Each(func(i int, itemSelect *goquery.Selection) {
			titleDoc := itemSelect.Find("span[class=item_title]")
			link, _ := titleDoc.Find("a").Attr("href")
			image, _ := itemSelect.Find("img").Attr("src")
			nodeDoc := itemSelect.Find("a[class=node]")
			nodeLink, _ := nodeDoc.Attr("href")
			user := itemSelect.Find("span strong").First().Find("a")

			replyDoc := itemSelect.Find("span[class=topic_info]")
			replayTime := ""
			replayName := ""
			if strings.Contains(replyDoc.Text(), "最后回复") {
				//有回复人 时间
				replayName = replyDoc.Find("strong").Last().Text()
				replyDoc.Find("a").Remove()
				replyDoc.Find("strong").Remove()
				replyDoc.Find(".votes").Remove()
				replayTime = strings.TrimSpace(
					replace(
						replace(
							replace(replyDoc.Text(), "•", ""), "最后回复来自", ""), " ", ""))
			}

			var items = new(module.Items)

			items.Title = titleDoc.Text()
			items.TopicsId = strings.Replace(strings.Split(link, "#")[0], "/t/", "", -1)
			items.Node = nodeDoc.Text()
			items.NodeId = strings.Replace(nodeLink, "/go/", "", -1)
			items.Images = image
			//回复时间
			items.ReplyName = replayName
			items.ReplyTime = replayTime

			reply, _ := strconv.Atoi(itemSelect.Find("a[class=count_livid]").Text())
			items.ReplyCount = reply

			items.UserName = user.Text()
			list = append(list, *items)
			//list.PushBack(*items)
		})

	})
	return list, ""
}
