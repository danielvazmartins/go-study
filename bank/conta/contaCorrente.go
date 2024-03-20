package conta

import (
	"bank/cliente"
	"fmt"
)

func ModContaCorrente() {
	fmt.Println("Module de Conta Corrente")
}

type ContaCorrente struct {
	Titular    cliente.Titular
	NumAgencia int
	NumConta   int
	saldo      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, bool) {
	if valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!", true
	} else {
		return "Saldo insuficiente!", false
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito >= 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso!"
	} else {
		return "Valor do depósito inválido!"
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (string, bool) {
	msg, status := c.Sacar(valorDaTransferencia)
	if status {
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferência realizada com sucesso!", true
	} else {
		return msg, false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
