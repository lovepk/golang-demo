package serializer

type Response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Error string `json:"error"`
}

//DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"items"`
	Total uint        `json:"total"`
}

type TokenData struct {
	User interface{} `json:"user"`
	Token string `json:"token"`
}

// BulidListResponse 带有总数的列表构建器
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList {
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}