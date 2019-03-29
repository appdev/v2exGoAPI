package query

import (
	"strings"
	"../module"
	"github.com/PuerkitoBio/goquery"
)

func DetailsTopics(ids string) (json interface{}, error string) {
	var requestUrl = BaseUrl + "t/" + ids
	doc, err := RequestClient(requestUrl)

	if err != "" {
		return "", err
	}
	// Find the review itemsclass="cell"
	var contentDoc = doc.Find("#Main").Find(".box").First()
	//切片方式
	var details = new(module.PageDetails)
	//list:=list.New()
	headerDoc := contentDoc.Find(".header").First()
	image, _ := headerDoc.Find(".fr").Find("img").Attr("src")
	link, _ := headerDoc.Find(".fr").Find("a").Attr("href")
	name := strings.Replace(link, "/member/", "", -1)

	//时间
	//移除a标签
	headerDoc.Find("small[class=gray]").Find("a").Remove()
	sendDoc := replace(headerDoc.Find("small[class=gray]").Text(), "By", "")
	timeDoc := replace(strings.Split(replace(sendDoc, "at", ""), ",")[0], " ", "")

	html, htmlErr := contentDoc.Find(".cell").Find(".topic_content").Find(".markdown_body").Html()

	//节点
	nodeDoc := headerDoc.Find(".chevron").Next()

	//附言
	subtle := contentDoc.Find(".subtle")
	if subtle == nil {
		details.ContentSub = nil
	} else {
		var subItem [] module.ContentSub
		subtle.Each(func(i int, selection *goquery.Selection) {
			var sub = new(module.ContentSub)
			sub.SendTime = strings.TrimSpace(strings.Split(selection.Find(".fade").Text(), "·")[1])
			sub.Content = selection.Find(".topic_content").Text()
			subItem = append(subItem, *sub)
		})
		details.ContentSub = subItem
	}
	//node
	details.Node = nodeDoc.Text()
	nodeID, _ := nodeDoc.Attr("href")
	details.NodeId = strings.Replace(nodeID, "/go/", "", -1)

	details.Images = image
	details.UserName = name
	details.Title = headerDoc.Find("h1").Text()
	details.SendTime = strings.TrimSpace(timeDoc)
	if htmlErr == nil {
		details.Content = html
	} else {
		details.Content = ""
	}

	//lang, err := util.DataToJson(list)
	//return string(lang), err
	return details, ""
}

func TopicsReply(ids string, page string) (json interface{}, error string) {
	var requestUrl = BaseUrl + "t/" + ids + "?p=" + page
	doc, err := RequestClient(requestUrl)

	if err != "" {
		return "", err
	}
	// Find the review itemsclass="cell"
	var contentDoc = doc.Find("#Main").Find(".box").Eq(1)

	//切片方式
	var list [] module.TopicsReply

	//list:=list.New()
	contentDoc.Find(`div[class=cell]`).Each(func(i int, selection *goquery.Selection) {
		var reply = new(module.TopicsReply)
		//头像
		src, _ := selection.Find("td").Find(".avatar").Attr("src")
		reply.Images = src
		//名字
		reply.UserName = selection.Find("a[class=dark]").Text()
		//时间 ♥ 1
		reply.SendTime = replace(selection.Find(`span[class="ago"]`).Text(), " ", "")
		//感谢
		reply.Thank = strings.TrimSpace(replace(selection.Find(`span[class="small fade"]`).Text(), "♥", ""))
		//内容 `span[class="small fade"]`
		replyDoc := selection.Find(`div[class="reply_content"]`)
		replyData := replyDoc.Find(`a[href^="/member"]`)
		if replyData == nil {
			reply.ReplyAt = nil
		} else {
			var replyAtList [] module.ReplyAt
			replyDoc.Find(`a[href^="/member/"]`).Each(func(i int, selection *goquery.Selection) {
				var replyAt = new(module.ReplyAt)
				replyAt.UserName = selection.Text()
				replyAtList = append(replyAtList, *replyAt)
			})
			reply.ReplyAt = replyAtList
		}

		_, htmlErr := selection.Find(`div[class="reply_content"]`).Html()
		if htmlErr == nil {
			reply.Content = selection.Find(`div[class="reply_content"]`).Text()
		} else {
			reply.Content = ""
		}
		list = append(list, *reply)
	})
	list = list[2:]
	return list, ""
}

func replace(s string, old string, new string) string {
	return strings.Replace(s, old, new, -1)
}
