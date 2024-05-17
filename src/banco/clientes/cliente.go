package clientes

type Cliente struct {
	Nome string
	Cpf string
	Profissao string
}

func (t * Cliente) Imprimir() {
	println("Cliente:", t.Nome, "| Cpf:", t.Cpf, "| Profissão:", t.Profissao)
}

func (t * Cliente) Descricao() string {

	cliente := "Cliente: " + t.Nome + " | Cpf: " + t.Cpf + " | Profissão: " + t.Profissao 
	return cliente
}