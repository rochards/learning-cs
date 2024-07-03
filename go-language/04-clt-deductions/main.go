package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TaxRate struct {
	income       float64
	contribution float64
}

// dados da tabela de contribuição com valores de 2024:
// https://www.gov.br/inss/pt-br/assuntos/confira-as-aliquotas-de-contribuicao-ao-inss-com-o-aumento-do-salario-minimo
var FIRST_CEILING_RATE = TaxRate{income: 1412, contribution: 7.5 / 100}
var SECOND_CEILING_RATE = TaxRate{income: 2666.68, contribution: 9.0 / 100}
var THIRD_CEILING_RATE = TaxRate{income: 4000.03, contribution: 12.0 / 100}
var FOURTH_CEILING_RATE = TaxRate{income: 7786.02, contribution: 14.0 / 100}

func readSalary() (float64, error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Informe seu salário. (Exemplo: 2.500,00)")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("entrada inválida")
	}

	cleanedInput := strings.ReplaceAll(input, ".", "")
	cleanedInput = strings.ReplaceAll(cleanedInput, ",", ".")

	amount, err := strconv.ParseFloat(strings.TrimSpace(cleanedInput), 64)
	if err != nil {
		return 0, fmt.Errorf("entrada inválida")
	}

	return amount, nil
}

func formatSalaryInBRL(salary float64) string {
	salaryAsString := fmt.Sprintf("%.2f", salary)
	salaryAsString = strings.ReplaceAll(salaryAsString, ".", ",")
	return salaryAsString
}

func calculateINSSDeductions(salary float64) float64 {

	if salary < FIRST_CEILING_RATE.income {
		fmt.Println("Valor abaixo de um salário mínimo. Tente novamente!")
		return -1
	}

	if salary > FOURTH_CEILING_RATE.income {
		// valor acima da quarta faixa:
		// - acima de R$ 7.786,02
		partialTax := (FOURTH_CEILING_RATE.income - THIRD_CEILING_RATE.income) * FOURTH_CEILING_RATE.contribution
		printFormattedINNSInfo(4, FOURTH_CEILING_RATE.income, THIRD_CEILING_RATE.income, FOURTH_CEILING_RATE.contribution, partialTax)

		return partialTax + calculateINSSDeductions(THIRD_CEILING_RATE.income)
	}

	if salary > THIRD_CEILING_RATE.income {
		// valor acima do teto da terceira faixa, por tanto dentro da quarta faixa:
		// - acima de R$ 4.000,03, até R$ 7786,02
		partialTax := (salary - THIRD_CEILING_RATE.income) * FOURTH_CEILING_RATE.contribution
		printFormattedINNSInfo(4, salary, THIRD_CEILING_RATE.income, FOURTH_CEILING_RATE.contribution, partialTax)

		return partialTax + calculateINSSDeductions(THIRD_CEILING_RATE.income)
	}

	if salary > SECOND_CEILING_RATE.income {
		// valor acima do teto da segunda faixa, por tanto dentro da terceira faixa:
		// - acima de R$ 2.666,68, até R$ 4.000,03
		partialTax := (salary - SECOND_CEILING_RATE.income) * THIRD_CEILING_RATE.contribution
		printFormattedINNSInfo(3, salary, SECOND_CEILING_RATE.income, THIRD_CEILING_RATE.contribution, partialTax)

		return partialTax + calculateINSSDeductions(SECOND_CEILING_RATE.income)
	}

	if salary > FIRST_CEILING_RATE.income {
		// valor acima do teto da primeira faixa, por tanto dentro da segunda faixa:
		// - acima de R$ 1.412,00, até R$ 2.666,68
		partialTax := (salary - FIRST_CEILING_RATE.income) * SECOND_CEILING_RATE.contribution
		printFormattedINNSInfo(2, salary, FIRST_CEILING_RATE.income, SECOND_CEILING_RATE.contribution, partialTax)

		return partialTax + calculateINSSDeductions(FIRST_CEILING_RATE.income)

	} else {
		// caso básico da função

		// salário dentro da primeira faixa, usa-se o valor do teto para o cálculo:
		// - até R$ 1.412,00
		partialTax := FIRST_CEILING_RATE.income * FIRST_CEILING_RATE.contribution
		printFormattedINNSInfo(1, FIRST_CEILING_RATE.income, 0, FIRST_CEILING_RATE.contribution, partialTax)

		return partialTax
	}
}

func printFormattedINNSInfo(row int, salary, ceilingIncome, contribution, contributionTax float64) {
	fmt.Printf("%dª faixa => R$ %s - R$ %s = R$ %s x %s = R$ %s\n",
		row,
		formatSalaryInBRL(salary),
		formatSalaryInBRL(ceilingIncome),
		formatSalaryInBRL(salary-ceilingIncome),
		formatSalaryInBRL(contribution),
		formatSalaryInBRL(contributionTax),
	)
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
			salary, err := readSalary()
			if err != nil {
				fmt.Println("Erro: " + err.Error())
				fmt.Println("Tente novamente!")
			}

			deductionTax := calculateINSSDeductions(salary)
			if deductionTax > 0 {
				// evitei retornar error na função calculateINSSDeductions para não ter que ficar lidando
				// com erro nas chamadas recursivas
				fmt.Println("Desconto do INSS: R$ ", formatSalaryInBRL(deductionTax))
			}

		default:
			fmt.Println("Unknown option")
		}
	}
}
