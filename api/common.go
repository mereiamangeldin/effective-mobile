package api

type Error struct {
	Message string `json:"message"`
}

type Ok struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
