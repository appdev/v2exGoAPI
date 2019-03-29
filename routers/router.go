package routers

import (
	"net/http"
	"../module"
	"github.com/gin-gonic/gin"
	"../query"
)

// LoadRouters 初始化router
func LoadRouters(router *gin.Engine) {
	loadRouters(router)
}
func loadRouters(router *gin.Engine) {

	router.GET("/tab", func(context *gin.Context) {
		url := context.DefaultQuery("url", "hot")
		data, err := query.ItemByTab(url, "")
		resultData(context, data, err)
	})
	router.GET("/detail", func(context *gin.Context) {
		id := context.DefaultQuery("id", "")
		if id == "" {
			errorRequest(context)
		} else {
			data, err := query.DetailsTopics(id)
			resultData(context, data, err)
		}
	})
	router.GET("/reply", func(context *gin.Context) {
		id := context.DefaultQuery("id", "")
		page := context.DefaultQuery("page", "1")
		if id == "" {
			errorRequest(context)
		} else {
			data, err := query.TopicsReply(id, page)
			resultData(context, data, err)
		}
	})

	router.GET("/node", func(context *gin.Context) {
		data, err := query.AllNode()
		resultData(context, data, err)
	})
	router.GET("/go", func(context *gin.Context) {
		node := context.DefaultQuery("node", "")
		page := context.DefaultQuery("page", "1")
		if node == "" {
			errorRequest(context)
		} else {
			data, err := query.NodeDetail(node, page)
			resultData(context, data, err)
		}
	})
	router.GET("/user/replies", func(context *gin.Context) {
		name := context.DefaultQuery("name", "")
		page := context.DefaultQuery("page", "1")
		if name == "" {
			errorRequest(context)
		} else {
			data, err := query.UserReplies(name,page)
			resultData(context, data, err)
		}
	})
	router.GET("/user/topics", func(context *gin.Context) {
		name := context.DefaultQuery("name", "")
		if name == "" {
			errorRequest(context)
		} else {
			data, err := query.UserTopics(name)
			resultData(context, data, err)
		}
	})
	router.GET("/user", func(context *gin.Context) {
		name := context.DefaultQuery("name", "")
		if name == "" {
			errorRequest(context)
		} else {
			data, err := query.UserInfo(name)
			resultData(context, data, err)
		}
	})
}

func errorRequest(context *gin.Context) {
	resultData(context, nil, "请求地址出错")
}

func resultData(context *gin.Context, data interface{}, err string) {

	var status = http.StatusOK
	if err != "" {
		status = http.StatusInternalServerError
	}
	context.JSON(http.StatusOK, module.Response{status, data, err})
}
