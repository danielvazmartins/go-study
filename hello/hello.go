package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const qtdeMonitoramentos int = 3
const intervaloMonitoramento time.Duration = 5 * time.Second
const logFile = "log.txt"

func main() {
	bemVindo()

	for {
		exibeMenu()
		menuSelected := leMenu()

		switch menuSelected {
		case 1:
			iniciaMonitoramento()
		case 2:
			exibirLogs()
		case 3:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida!")
			os.Exit(-1)
		}
	}
}

func bemVindo() {
	var nome string = "Daniel"
	var nome2 = "Daniel"
	nome3 := "Daniel"
	var idade int = 40
	var versao float32 = 1.1
	fmt.Println("Hello", nome)
	fmt.Println("Sua idade é", idade, "anos")
	fmt.Println("Versão desse programa:", versao)

	fmt.Println("A variável nome2 é do tipo", reflect.TypeOf(nome2))
	fmt.Println("A variável nome3 é do tipo", reflect.TypeOf(nome3), "e tem o valor de", nome3)
	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("Escolha uma opção do Menu")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func leMenu() int {
	var menuSelected int
	fmt.Scan(&menuSelected)

	fmt.Println("Você selecionou a opção", menuSelected)
	fmt.Println("")
	return menuSelected
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")

	sites := carregaSites()

	for i := 0; i < len(sites); i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("")
		time.Sleep(intervaloMonitoramento)
	}
	fmt.Println("")
}

func testaSite(site string) {
	response, error := http.Get(site)

	if error != nil {
		fmt.Println(error)
		return
	}

	if response.StatusCode == 200 {
		fmt.Println("O site:", site, "está funcionando corretamente!")
		registraLog(site, true)
	} else {
		fmt.Println("O site:", site, "está com algum problema. StatusCode:", response.StatusCode)
		registraLog(site, false)
	}
}

func carregaSites() []string {
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Erro ao ler arquivo", err)
		os.Exit(-1)
	}
	reader := bufio.NewReader(file)

	var sites []string
	for {
		linha, err := reader.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}

func registraLog(site string, online bool) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao ler arquivo", err)
		os.Exit(-1)
	}
	msg := time.Now().Format("02/01/2006 15:04:05 - ") + site + " - Online: " + strconv.FormatBool(online) + "\n"
	file.WriteString(msg)
	file.Close()
}

func exibirLogs() {
	fmt.Println("Exibindo logs...")
	file, err := os.ReadFile(logFile)
	if err != nil {
		fmt.Println("Erro ao ler arquivo", err)
		os.Exit(-1)
	}

	fmt.Println(string(file))
}
