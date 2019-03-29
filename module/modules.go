package module

type Items struct {
	Images     string `json:"avatar"`
	Title      string `json:"title"`
	TopicsId   string `json:"id"`
	Node       string `json:"node"`
	NodeId     string `json:"nodeId"`
	UserName   string `json:"userName"`
	ReplyCount int    `json:"replyCount"`
	ReplyName  string `json:"replyName"`
	ReplyTime  string `json:"replyTime"`
}

type PageDetails struct {
	Images     string `json:"avatar"`
	Title      string `json:"title"`
	UserName   string `json:"userName"`
	SendTime   string
	Node       string `json:"node"`
	NodeId     string `json:"nodeId"`
	Content    string
	ContentSub []ContentSub
}
type RootNode struct {
	Title     string
	ChildNode []ChildNode
}

type ChildNode struct {
	Title     string
	AliasName string
}

type ContentSub struct {
	Content  string
	SendTime string
}
type TopicsReply struct {
	Images   string `json:"avatar"`
	UserName string `json:"userName"`
	SendTime string
	Content  string
	Thank    string
	ReplyAt  []ReplyAt
}

type ReplyAt struct {
	UserName string `json:"userName"`
}

type NodeInfo struct {
	NodeImage   string
	NodeName    string
	TopicsCount string
	Topics      []Items
}

type UserInfo struct {
	Name     string
	Avatar   string
	JoinTime string
	Activity string
}

type UserReplies struct {
	TopicsTitle  string
	Node         string
	TopicsUser   string
	TopicsId     string
	PostTime     string
	ReplyContent string
}
