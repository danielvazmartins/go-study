package conta

type contaGenerica interface {
	ObterSaldo() float64
}

func ObterSaldo(conta contaGenerica) float64 {
	return conta.ObterSaldo()
}
