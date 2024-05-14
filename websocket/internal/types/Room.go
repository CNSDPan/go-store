package types

type Room struct {
	IId    int64  `json:"iId,string"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	People int    `json:"people,string"`
}
