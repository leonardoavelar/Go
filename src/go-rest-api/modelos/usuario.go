package modelos

type Usuario struct {
	Id        int    `json:"id"`
	Nome      string `json:"nome"`
	Profissao string `json:"profissao"`
}
