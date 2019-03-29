# v2exGoAPI
使用Go语言通过爬虫方式实现的V2ex的API

V2ex作为国内（中国）程序员的聚集地之一广为人知，刚好在学习Flutter，所以打算写个Flutter的V2客户端练手，然后发现V2自身提供的客户端很少，功能不全，所以就有了这个项目！

> 其实 go 也是刚学不到半个月，这个项目也算是练手。代码上如果有问题，还希望高手指点！🙏

目前实现的功能:
- 首页几个tab的数据(姑且这么叫吧🤣)
- 帖子内容、帖子回复
- 所有节点
- 节点详情
- 用户详情
- 用户发布的主题
- 用户的回复
> 如非特殊说明，所有请求方式均为GET


### 首页Tab
请求：
> /  

参数：
> tab

例子：
> /?tab=hot 请求最热  
> /?tab=qna 请求问与答

结果:
```json
{
    "code": 200,
    "data": [
        {
            "avatar": "//cdn.v2ex.com/avatar/fb2e/6365/24324_normal.png?m=1553556157",
            "title": "在 V 站众筹一款正在开发的 SSH 客户端： Termix",
            "id": "549770",
            "node": "分享创造",
            "nodeId": "create",
            "userName": "plqws",
            "replyCount": 395,
            "replyName": "liyangyijie",
            "replyTime": "几秒前"
        }
            ],
    "msg": ""
}
```

这里说明一下，V站首页有个标签 **全部** 这里的全部并不是真的全部，可以自己操作一下就能明白区别。
想要获取全部帖子 `/?tab=recent&page=1`即可(可能需要登陆后的cookie)

### 帖子内容

请求：
> /detail  

参数：
> id

例子：
> /detail?id=549885   


结果:
```json
{
    "code": 200,
    "data": {
        "avatar": "//cdn.v2ex.com/avatar/0c7b/7138/299623_large.png?m=1552894684",
        "title": "[深圳] 7 天求职感受",
        "userName": "duzhihao",
        "SendTime": "·22小时47分钟前·4669次点击",
        "node": "深圳",
        "nodeId": "shenzhen",
        "Content": "<p>上周四来到深圳...现在找工作确实有点困难，特别是技术有限的人来说，大家平时还是努力学习啊。接下来我会去杭州碰碰运气，先祝自己好运。</p>\n",
        "ContentSub": [ //这里是附言
            {
                "Content": "谢谢大家的回帖，我在给自己一周时间吧。\n我的理想公司：\n\n有一个稳定的产品驱动。\n研发团队最好20人以上，前端最好5个以上。\n研发团队分配合理，不能没有产品、测试。\n非单休。\n\n希望我能找到这样的公司，好运吧！\n",
                "SendTime": "18 小时 0 分钟前"
            }
        ]
    },
    "msg": ""
}
```
帖子内容返回的是 带有HTML标签的文本(因为帖子有格式)，方便在各个平台展示。ContentSub是帖子的附言，没有则为null

### 帖子回复


请求：
> /reply  

参数：
> id
> page   分页 默认为1

例子：
> reply?id=549885&page=1


结果:
```json
{
    "code": 200,
    "data": [
        {
            "avatar": "//cdn.v2ex.com/avatar/0c7b/7138/299623_normal.png?m=1552894684",
            "userName": "duzhihao",
            "SendTime": "22小时40分钟前",
            "Content": "@hahasong 主要住在别人宿舍，不太好意思，7 天也基本摸清这里的环境了。很难有运气砸到头上了...",
            "Thank": "",
            "ReplyAt": [
                {
                    "userName": "hahasong"
                }
            ]
        }
    ],
    "msg": ""
}
```
这里单独吧 被@ 的用户提取出来放在ReplyAt数组中！


### 所有节点


请求：
> /node  


结果:
```json
{
    "code": 200,
    "data": [
        {
            "Title": "混沌海",
            "ChildNode": [
                {
                    "Title": "地球",
                    "AliasName": "earth"
                },
                {
                    "Title": "招商银行",
                    "AliasName": "cmb"
                }
            ]
        }
    ],
    "msg": ""
}
```  

### 节点详情

请求：
> /go  

参数：
> node
> page   分页 默认为1

例子：
> /go?node=shanghai&page=2


结果:
```json
{
    "code": 200,
    "data": {
        "NodeImage": "//cdn.v2ex.com/navatar/6f49/22f4/18_xxlarge.png?m=1542418426",
        "NodeName": "上海",
        "TopicsCount": "1634",
        "Topics": [
            {
                "avatar": "//cdn.v2ex.com/gravatar/b279806db1b06506f55355c45a4301f9?s=48&d=retro",
                "title": "大家平时会去上海周边什么地方自驾游吗？",
                "id": "549718",
                "node": "shanghai",
                "nodeId": "上海",
                "userName": "skunktalks",
                "replyCount": 3,
                "replyName": "Lawlieti",
                "replyTime": "4小时30分钟前"
            }
        ]
    },
    "msg": ""
}
```

### 用户详情

请求：
> /user  

参数：
> name

例子：
> /user?name=yfixx


结果:
```json
{
    "code": 200,
    "data": {
        "Name": "yfixx",
        "Avatar": "//cdn.v2ex.com/gravatar/6138b0dbd40e3fc92a5321892282dd7a?s=73&d=retro",
        "JoinTime": "V2EX 第 357200 号会员，加入于 2018-10-19 15:39:35 +08:00",
        "Activity": "2372"
    },
    "msg": ""
}
```

### 用户发布的主题  

请求：
> /user/topics  

参数：
> name

例子：
> /user/topics?name=yfixx


结果:
```json
{
    "code": 200,
    "data": [
        {
            "avatar": "",
            "title": "此时的我该如何选择？",
            "id": "499872",
            "node": "职场话题",
            "nodeId": "career",
            "userName": "yfixx",
            "replyCount": 6,
            "replyName": "loryyang",
            "replyTime": "157天前"
        }
    ],
    "msg": ""
}
```
### 用户的回复


请求：
> /user/replies  

参数：
> name

例子：
> /user/replies?name=yfixx


结果:
```json
{
    "code": 200,
    "data": [
        {
            "TopicsTitle": "问题来了，程序员们你们用的是哪种 PDF 阅读器？",
            "Node": "programmer",
            "TopicsUser": "yfixx",
            "TopicsId": "549362",
            "PostTime": "7 小时 30 分钟前",
            "ReplyContent": "@<a target=\"_blank\" href=\"/member/momo5269\" rel=\"nofollow\">momo5269</a> 用过的真多呀"
        }
    ],
    "msg": ""
}
```
同样，回复内容包含了HTML标签，方便展示。
