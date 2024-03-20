package main

import (
	"bank/cliente"
	"bank/conta"
	"fmt"
)

func main() {
	conta.ModContaCorrente()
	conta.ModContaPoupanca()

	titularA := cliente.Titular{
		Nome:      "Daniel",
		CPF:       12312312312,
		Profissao: "Analista de TI",
	}
	contaA := conta.ContaCorrente{
		Titular:    titularA,
		NumAgencia: 123,
		NumConta:   123456,
	}
	contaA.Depositar(50.25)

	titularB := cliente.Titular{"João", 11122233344, "Manobrista"}
	contaB := conta.ContaCorrente{
		Titular:    titularB,
		NumAgencia: 256,
		NumConta:   123123}
	contaB.Depositar(500)

	var contaC *conta.ContaCorrente
	contaC = new(conta.ContaCorrente)
	contaC.Titular = cliente.Titular{Nome: "Maria"}

	contaD := new(conta.ContaCorrente)
	contaD.Titular = cliente.Titular{Nome: "Marcos"}

	fmt.Println(contaA)
	fmt.Println(contaB)
	fmt.Println(contaC)
	fmt.Println(*contaC)
	fmt.Println(contaD)

	fmt.Println("Saldo da contaA:", contaA.ObterSaldo())
	fmt.Println(contaA.Sacar(10))
	fmt.Println("Saldo da contaA:", contaA.ObterSaldo())
	fmt.Println(contaA.Depositar(30))
	fmt.Println("Saldo da contaA:", contaA.ObterSaldo())

	fmt.Println("Realizando transferência de contaA para contaB")
	msg, status := contaA.Transferir(1500, &contaB)
	fmt.Println(msg)
	if status {
		fmt.Println("Saldo da contaA:", contaA.ObterSaldo())
		fmt.Println("Saldo da contaB:", contaB.ObterSaldo())
	}

	titularAPoup := cliente.Titular{"Joana", 11122211134, "Faxineira"}
	contaAPoup := conta.ContaPoupanca{
		Titular:    titularAPoup,
		NumAgencia: 1234,
		NumConta:   4321,
		Operacao:   1}
	contaAPoup.Depositar(200)
	fmt.Println(contaAPoup)
	fmt.Println("Saldo da contaAPoup:", conta.ObterSaldo(&contaAPoup))
}
