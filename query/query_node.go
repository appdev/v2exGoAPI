package query

import (
	"../module"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"log"
)

func NodeDetail(node string, page string) (json interface{}, error string) {
	var requestUrl = BaseUrl + "go/" + node + "?p=" + page
	doc, err := RequestClient(requestUrl)

	if err != "" {
		return "", err
	}
	// Find the review itemsclass="cell"
	var contentDoc = doc.Find("#Main")
	// node 的 信息
	nodeDoc := contentDoc.Find(".node_header")
	var info = new(module.NodeInfo)
	image, imageErr := nodeDoc.Find(".node_avatar").Find("img").Attr("src")
	if imageErr {
		info.NodeImage = image
	} else {
		info.NodeImage = ""
	}
	info.TopicsCount = nodeDoc.Find(".node_info").Find("strong").Text()

	//名字
	nodeName := strings.Split(nodeDoc.Find(".node_info").Text(), "›")[1]
	log.Println(replace(nodeName, "全部主题", ""))
	info.NodeName = strings.TrimSpace(replace(replace(nodeName, "全部主题", ""), " ", ""))

	//切片方式
	var list [] module.Items
	//list:=list.New()
	contentDoc.Find("#TopicsNode").Find(".cell").Each(func(i int, selection *goquery.Selection) {
		titleDoc := selection.Find("span[class=item_title]")
		link, _ := titleDoc.Find("a").Attr("href")
		image, _ := selection.Find("img").Attr("src")
		user := selection.Find("span strong").First().Find("a")

		replyDoc := selection.Find(`span[class="small fade"]`)
		replayTime := ""
		replayName := ""
		if strings.Contains(replyDoc.Text(), "最后回复") {
			//有回复人 时间
			replayName = replyDoc.Find("strong").Last().Text()
			replyDoc.Find("strong").Remove()
			replayTime = strings.TrimSpace(
				replace(
					replace(
						replace(replyDoc.Text(), "•", ""), "最后回复来自", ""), " ", ""))
		}

		var items = new(module.Items)

		items.Title = titleDoc.Text()
		items.TopicsId = strings.Replace(strings.Split(link, "#")[0], "/t/", "", -1)
		items.Node = node
		items.NodeId = info.NodeName
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
	info.Topics = list
	//lang, err := util.DataToJson(list)
	//return string(lang), err
	return info, ""
}

func AllNode() (json interface{}, error string) {
	var requestUrl = BaseUrl + "planes"
	doc, err := RequestClient(requestUrl)

	if err != "" {
		return "", err
	}
	// Find the review itemsclass="cell"
	//切片方式
	var list [] module.RootNode

	//list:=list.New()
	doc.Find("#Wrapper").Find(".content").Find("div[class=box]").Each(
		func(i int, selection *goquery.Selection) {
			if i > 0 {
				var reply = new(module.RootNode)
				//根节点 name
				rootNode := selection.Find("div[class=header]")
				rootNode.Find(`span[class="fr fade"]`).Remove()
				reply.Title = rootNode.Text()
				//子节点
				var nodes [] module.ChildNode
				selection.Find("div[class=inner]").Find("a").Each(
					func(i int, selection *goquery.Selection) {
						var node = new(module.ChildNode)
						node.Title = selection.Text()
						href, err := selection.Attr("href")
						if err {
							node.AliasName = replace(href, "/go/", "")
						} else {
							node.AliasName = ""
						}
						nodes = append(nodes, *node)
					})

				reply.ChildNode = nodes

				list = append(list, *reply)
			}
		})
	return list, ""
}
