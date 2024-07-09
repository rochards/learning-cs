package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ContribuicaoINSS struct {
	salarioContribuicao float64
	aliquota            float64
}

// dados da tabela de contribuição com valores de 2024:
// https://www.gov.br/inss/pt-br/assuntos/confira-as-aliquotas-de-contribuicao-ao-inss-com-o-aumento-do-salario-minimo
var TETO_FAIXA_1 = ContribuicaoINSS{salarioContribuicao: 1412, aliquota: 7.5 / 100}
var TETO_FAIXA_2 = ContribuicaoINSS{salarioContribuicao: 2666.68, aliquota: 9.0 / 100}
var TETO_FAIXA_3 = ContribuicaoINSS{salarioContribuicao: 4000.03, aliquota: 12.0 / 100}
var TETO_FAIXA_4 = ContribuicaoINSS{salarioContribuicao: 7786.02, aliquota: 14.0 / 100}

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

func calculaContribuicaoINSS(salario float64) float64 {

	if salario < TETO_FAIXA_1.salarioContribuicao {
		fmt.Println("Valor abaixo de um salário mínimo. Tente novamente!")
		return -1
	}

	if salario > TETO_FAIXA_4.salarioContribuicao {
		// valor acima da quarta faixa:
		// - acima de R$ 7.786,02
		deducaoParcial := (TETO_FAIXA_4.salarioContribuicao - TETO_FAIXA_3.salarioContribuicao) * TETO_FAIXA_4.aliquota
		exibeCalculoINSS(4, TETO_FAIXA_4.salarioContribuicao, TETO_FAIXA_3.salarioContribuicao, TETO_FAIXA_4.aliquota, deducaoParcial)

		return deducaoParcial + calculaContribuicaoINSS(TETO_FAIXA_3.salarioContribuicao)
	}

	if salario > TETO_FAIXA_3.salarioContribuicao {
		// valor acima do teto da terceira faixa, por tanto dentro da quarta faixa:
		// - acima de R$ 4.000,03, até R$ 7786,02
		deducaoParcial := (salario - TETO_FAIXA_3.salarioContribuicao) * TETO_FAIXA_4.aliquota
		exibeCalculoINSS(4, salario, TETO_FAIXA_3.salarioContribuicao, TETO_FAIXA_4.aliquota, deducaoParcial)

		return deducaoParcial + calculaContribuicaoINSS(TETO_FAIXA_3.salarioContribuicao)
	}

	if salario > TETO_FAIXA_2.salarioContribuicao {
		// valor acima do teto da segunda faixa, por tanto dentro da terceira faixa:
		// - acima de R$ 2.666,68, até R$ 4.000,03
		deducaoParcial := (salario - TETO_FAIXA_2.salarioContribuicao) * TETO_FAIXA_3.aliquota
		exibeCalculoINSS(3, salario, TETO_FAIXA_2.salarioContribuicao, TETO_FAIXA_3.aliquota, deducaoParcial)

		return deducaoParcial + calculaContribuicaoINSS(TETO_FAIXA_2.salarioContribuicao)
	}

	if salario > TETO_FAIXA_1.salarioContribuicao {
		// valor acima do teto da primeira faixa, por tanto dentro da segunda faixa:
		// - acima de R$ 1.412,00, até R$ 2.666,68
		deducalParcial := (salario - TETO_FAIXA_1.salarioContribuicao) * TETO_FAIXA_2.aliquota
		exibeCalculoINSS(2, salario, TETO_FAIXA_1.salarioContribuicao, TETO_FAIXA_2.aliquota, deducalParcial)

		return deducalParcial + calculaContribuicaoINSS(TETO_FAIXA_1.salarioContribuicao)

	} else {
		// caso básico da função

		// salário dentro da primeira faixa, usa-se o valor do teto para o cálculo:
		// - até R$ 1.412,00
		deducaoParcial := TETO_FAIXA_1.salarioContribuicao * TETO_FAIXA_1.aliquota
		exibeCalculoINSS(1, TETO_FAIXA_1.salarioContribuicao, 0, TETO_FAIXA_1.aliquota, deducaoParcial)

		return deducaoParcial
	}
}

func exibeCalculoINSS(row int, salary, ceilingIncome, contribution, contributionTax float64) {
	fmt.Printf("%dª faixa => R$ %s - R$ %s = R$ %s x %s = R$ %s\n",
		row,
		formataDecimalBRL(salary),
		formataDecimalBRL(ceilingIncome),
		formataDecimalBRL(salary-ceilingIncome),
		formataDecimalBRL(contribution),
		formataDecimalBRL(contributionTax),
	)
}

func formataDecimalBRL(value float64) string {

	parts := strings.Split(fmt.Sprintf("%.2f", value), ".")
	integerPart := parts[0]
	decimalPart := parts[1]

	var result strings.Builder
	integerPartLength := len(integerPart)
	for i, digit := range integerPart {
		if i > 0 && (integerPartLength-i)%3 == 0 {
			result.WriteString(".")
		}
		result.WriteRune(digit)
	}

	result.WriteString(",")
	result.WriteString(decimalPart)

	return result.String()
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

			fmt.Printf("\nSalário informado: R$ %s\n", formataDecimalBRL(salario))
			descontoINSS := calculaContribuicaoINSS(salario)
			if descontoINSS > 0 {
				// evitei retornar error na função calculaContribuicaoINSS para não ter que ficar lidando
				// com erro nas chamadas recursivas
				fmt.Printf("==> Desconto do INSS: R$ %s\n", formataDecimalBRL(descontoINSS))
			}

		default:
			fmt.Println("Opção desconhecida. Tente novamente!")
		}
	}
}
