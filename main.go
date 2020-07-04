package main

import (
	"fmt"

	"github.com/user/geren-conta-banco_go/contas"
)

// PagarBoleto é uma função para o método de pagamento de boleto
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoBruno := contas.ContaPoupanca{}
	contaDoBruno.Depositar(100)
	PagarBoleto(&contaDoBruno, 60)

	fmt.Println(contaDoBruno.ObterSaldo())

	contaDaNik := contas.ContaCorrente{}
	contaDaNik.Depositar(500)
	PagarBoleto(&contaDaNik, 1000)

	fmt.Println(contaDaNik.ObterSaldo())

}
