package getcardcolor

type output struct {
	Colors []color `json:"colors"`
}

type color struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Variable string `json:"variable"`
}
