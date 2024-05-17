package pagamentos

import (
	"banco/interfaces"
)

func EfetuarPagamento(conta interfaces.Conta, valor float64) bool {

	return conta.Sacar(valor)
	
}