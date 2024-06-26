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

	if salary > FOURTH_CEILING_RATE.income {
		// valor acima da quarta faixa:
		// - acima de R$ 7.786,02
		diff := FOURTH_CEILING_RATE.income - THIRD_CEILING_RATE.income
		partialContribution := diff * FOURTH_CEILING_RATE.contribution

		fmt.Printf("4ª faixa => R$ %s - R$ %s = R$ %s x %s = R$ %s\n",
			formatSalaryInBRL(FOURTH_CEILING_RATE.income),
			formatSalaryInBRL(THIRD_CEILING_RATE.income),
			formatSalaryInBRL(diff),
			formatSalaryInBRL(FOURTH_CEILING_RATE.contribution),
			formatSalaryInBRL(partialContribution),
		)
		return partialContribution + calculateINSSDeductions(THIRD_CEILING_RATE.income)
	}

	if salary > THIRD_CEILING_RATE.income {
		// valor acima do teto da terceira faixa, por tanto dentro da quarta faixa:
		// - acima de R$ 4.000,03, até R$ 7786,02
		diff := salary - THIRD_CEILING_RATE.income
		partialContribution := diff * FOURTH_CEILING_RATE.contribution

		fmt.Printf("4ª faixa => R$ %s - R$ %s = R$ %s x %s = R$ %s\n",
			formatSalaryInBRL(salary),
			formatSalaryInBRL(THIRD_CEILING_RATE.income),
			formatSalaryInBRL(diff),
			formatSalaryInBRL(FOURTH_CEILING_RATE.contribution),
			formatSalaryInBRL(partialContribution),
		)

		return partialContribution + calculateINSSDeductions(THIRD_CEILING_RATE.income)
	}

	if salary > SECOND_CEILING_RATE.income {
		// valor acima do teto da segunda faixa, por tanto dentro da terceira faixa:
		// - acima de R$ 2.666,68, até R$ 4.000,03
		diff := salary - SECOND_CEILING_RATE.income
		partialContribution := diff * THIRD_CEILING_RATE.contribution

		fmt.Printf("3ª faixa => R$ %s - R$ %s = R$ %s x %s = R$ %s\n",
			formatSalaryInBRL(salary),
			formatSalaryInBRL(SECOND_CEILING_RATE.income),
			formatSalaryInBRL(diff),
			formatSalaryInBRL(THIRD_CEILING_RATE.contribution),
			formatSalaryInBRL(partialContribution),
		)

		return partialContribution + calculateINSSDeductions(SECOND_CEILING_RATE.income)
	}

	if salary > FIRST_CEILING_RATE.income {
		// valor acima do teto da primeira faixa, por tanto dentro da segunda faixa:
		// - acima de R$ 1.412,00, até R$ 2.666,68
		diff := salary - FIRST_CEILING_RATE.income
		partialContribution := diff * SECOND_CEILING_RATE.contribution

		fmt.Printf("2ª faixa => R$ %s - R$ %s = R$ %s x %s = R$ %s\n",
			formatSalaryInBRL(salary),
			formatSalaryInBRL(FIRST_CEILING_RATE.income),
			formatSalaryInBRL(diff),
			formatSalaryInBRL(SECOND_CEILING_RATE.contribution),
			formatSalaryInBRL(partialContribution),
		)

		return partialContribution + calculateINSSDeductions(FIRST_CEILING_RATE.income)

	} else {
		// caso básico da função

		// salário dentro da primeira faixa, usa-se o valor do teto para o cálculo:
		// - até R$ 1.412,00

		partialContribution := FIRST_CEILING_RATE.income * FIRST_CEILING_RATE.contribution

		fmt.Printf("1ª faixa => R$ %s - R$ %s = R$ %s x %s = R$ %s\n",
			formatSalaryInBRL(FIRST_CEILING_RATE.income),
			formatSalaryInBRL(0),
			formatSalaryInBRL(FIRST_CEILING_RATE.income),
			formatSalaryInBRL(FIRST_CEILING_RATE.contribution),
			formatSalaryInBRL(partialContribution),
		)

		return partialContribution
	}

}

func main() {

	for {
		fmt.Printf("\nDigite o número para escolher: \n" +
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
			fmt.Println("Desconto do INSS: R$ ", formatSalaryInBRL(deductionTax))

		default:
			fmt.Println("Unknown option")
		}
	}
}
