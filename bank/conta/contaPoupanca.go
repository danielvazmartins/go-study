package conta

import (
	"bank/cliente"
	"fmt"
)

type ContaPoupanca struct {
	Titular    cliente.Titular
	NumAgencia int
	NumConta   int
	Operacao   int
	saldo      float64
}

func ModContaPoupanca() {
	fmt.Println("Module de Conta Poupanca")
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito >= 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso!"
	} else {
		return "Valor do depósito inválido!"
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
