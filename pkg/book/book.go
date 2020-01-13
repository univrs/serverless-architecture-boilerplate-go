package book

type Book struct {
	Hashkey      string      `json:"hashkey"`
	Title        string      `json:"title"`
	Author       string      `json:"author"`
	Price        float64     `json:"price"`
	Updated      string      `json:"updated"`
	Created      string      `json:"created"`
	CustomStruct interface{} `json:",omitempty"`
}
