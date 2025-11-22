package getthemes

type output struct {
	Themes []theme `json:"tags"`
}

type theme struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
