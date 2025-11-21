package gettags

type output struct {
	Tags []tag `json:"tags"`
}

type tag struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	ColorId int    `json:"color_id"`
}
