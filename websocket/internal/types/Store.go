package types

type Store struct {
	IId   int64  `json:"iId,string"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
