package view

type ReqCreateSource struct {
	Iid int `json:"iid" form:"iid" binding:"required"`
	ReqUpdateSource
}

type ReqUpdateSource struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Desc     string `json:"desc" form:"desc"`
	URL      string `json:"url" form:"url"`
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Typ      int    `json:"typ" form:"typ"`
}

type ReqListSource struct {
	Typ int `json:"typ" form:"typ"`
}