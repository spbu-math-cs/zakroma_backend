package schemas

type Group struct {
	Id   int    `json:"id" example:"3"`
	Hash string `json:"hash" example:"bf633968b5fcc31f1eaa90764d061083bbed8ca29a59662df9e5fecf41d221f0"`
	Name string `json:"name" example:"семья"`
}
