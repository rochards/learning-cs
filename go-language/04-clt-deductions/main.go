package main

import (
	"bufio"
	"fmt"
	"main/inss"
	"main/irpf"
	"main/utils"
	"os"
	"strconv"
	"strings"
)

func leSalario() (float64, error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Informe seu salário. (Exemplo: 2.500,00)")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("entrada inválida")
	}

	cleanedInput := strings.ReplaceAll(input, ".", "")
	cleanedInput = strings.ReplaceAll(cleanedInput, ",", ".")

	salario, err := strconv.ParseFloat(strings.TrimSpace(cleanedInput), 64)
	if err != nil {
		return 0, fmt.Errorf("entrada inválida")
	}

	return salario, nil
}

func main() {

	for {
		fmt.Printf("\nDigite um número para escolher: \n" +
			"0 - Sair\n" +
			"1 - Calcular taxas\n")
		var option int
		fmt.Scan(&option)

		switch option {
		case 0:
			return

		case 1:
			salario, err := leSalario()
			if err != nil {
				fmt.Println("Erro: " + err.Error())
				fmt.Println("Tente novamente!")
			}

			descontoINSS, err := inss.CalculaContribuicaoINSS(salario)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println()
			descontoIRPF := irpf.CalculaContribuicaoIRPF(salario, descontoINSS, 2)

			salarioLiquido := salario - descontoINSS - descontoIRPF

			fmt.Printf("\n==> Salário líquido: R$ %s - R$ %s - R$ %s = R$ %s\n",
				utils.FormatDecimalInBRL(salario),
				utils.FormatDecimalInBRL(descontoINSS),
				utils.FormatDecimalInBRL(descontoIRPF),
				utils.FormatDecimalInBRL(salarioLiquido),
			)

		default:
			fmt.Println("Opção desconhecida. Tente novamente!")
		}
	}
}
