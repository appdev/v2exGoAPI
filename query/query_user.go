package query

import (
	"../module"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"log"
)

func UserTopics(name string) (data interface{}, error string) {
	var requestUrl = BaseUrl + "member/" + name + "/topics"
	doc, err := RequestClient(requestUrl)

	if err != "" {
		return "", err
	}
	// Find the review itemsclass="cell"
	var contentDoc = doc.Find("#Main").Find(".box").Find(".cell")
	//切片方式
	var list [] module.Items
	contentDoc.Each(func(i int, selection *goquery.Selection) {
		titleDoc := selection.Find("span[class=item_title]")
		link, _ := titleDoc.Find("a").Attr("href")
		image, _ := selection.Find("img").Attr("src")
		user := selection.Find("span strong").First().Find("a")
		nodeDoc := selection.Find("a[class=node]")
		nodeLink, _ := nodeDoc.Attr("href")
		replyDoc := selection.Find(`span[class="topic_info"]`)
		replayTime := ""
		replayName := ""
		if strings.Contains(replyDoc.Text(), "最后回复") {
			//有回复人 时间
			replayName = replyDoc.Find("strong").Last().Text()
			replyDoc.Find("strong").Remove()
			replyDoc.Find("div").Remove()
			replyDoc.Find("a").Remove()
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

		reply, _ := strconv.Atoi(selection.Find("a[class=count_livid]").Text())
		items.ReplyCount = reply

		items.UserName = user.Text()
		list = append(list, *items)
		//list.PushBack(*items)

	})

	return list, ""
}

func UserReplies(name string, page string) (json interface{}, error string) {
	var requestUrl = BaseUrl + "member/" + name + "/replies?p=" + page
	doc, err := RequestClient(requestUrl)
	if err != "" {
		return "", err
	}
	// Find the review itemsclass="cell"
	//切片方式
	var list [] module.UserReplies
	var contentDoc = doc.Find("#Main").Find(".dock_area")
	//list:=list.New()
	contentDoc.Each(func(i int, selection *goquery.Selection) {
		var reply = new(module.UserReplies)
		//帖子信息
		reply.PostTime = selection.Find(`span[class="fade"]`).Text()
		selection.Find(`span[class="gray"]`).Find("a").Each(func(i int, info *goquery.Selection) {
			switch i {
			case 0:
				reply.TopicsUser = info.Text()
				break
			case 1:
				node, nodeErr := info.Attr("href")
				if nodeErr {
					reply.Node = replace(node, "/go/", "")
				}
				break
			case 2:
				reply.TopicsTitle = info.Text()
				topics, topicsErr := info.Attr("href")
				if topicsErr {
					reply.TopicsId = strings.Replace(strings.Split(topics, "#")[0], "/t/", "", -1)
				}
				break
			}

		})

		//回复内容
		content, contentErr := selection.Next().Find(".reply_content").Html()
		if contentErr == nil {
			reply.ReplyContent = content
		}

		list = append(list, *reply)
	})
	return list, ""
}

func UserInfo(name string) (json interface{}, error string) {
	var requestUrl = BaseUrl + "member/" + name
	doc, err := RequestClient(requestUrl)
	if err != "" {
		return "", err
	}
	var contentDoc = doc.Find("#Main").Find(".box").First()
	log.Println(contentDoc.Html())
	var info = new(module.UserInfo)
	info.Name = contentDoc.Find("h1").Text()
	src, errImg := contentDoc.Find("img").Attr("src")
	if errImg {
		info.Avatar = src
	} else {
		info.Avatar = ""
	}
	span := contentDoc.Find("span[class=gray]")
	//活跃度
	info.Activity = span.Find("a").Text()

	//加入时间
	span.Find("div").Remove()
	span.Find("a").Remove()
	info.JoinTime = strings.TrimSpace(replace(span.Text(), "今日活跃度排名", ""))
	return info, ""
}
