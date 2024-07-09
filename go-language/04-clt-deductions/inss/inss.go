package inss

import (
	"fmt"
	"main/utils"
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

func CalculaContribuicaoINSS(salario float64) float64 {

	if salario < TETO_FAIXA_1.salarioContribuicao {
		fmt.Println("Valor abaixo de um salário mínimo. Tente novamente!")
		return -1
	}

	if salario > TETO_FAIXA_4.salarioContribuicao {
		// valor acima da quarta faixa:
		// - acima de R$ 7.786,02
		deducaoParcial := (TETO_FAIXA_4.salarioContribuicao - TETO_FAIXA_3.salarioContribuicao) * TETO_FAIXA_4.aliquota
		exibeCalculoINSS(4, TETO_FAIXA_4.salarioContribuicao, TETO_FAIXA_3.salarioContribuicao, TETO_FAIXA_4.aliquota, deducaoParcial)

		return deducaoParcial + CalculaContribuicaoINSS(TETO_FAIXA_3.salarioContribuicao)
	}

	if salario > TETO_FAIXA_3.salarioContribuicao {
		// valor acima do teto da terceira faixa, por tanto dentro da quarta faixa:
		// - acima de R$ 4.000,03, até R$ 7786,02
		deducaoParcial := (salario - TETO_FAIXA_3.salarioContribuicao) * TETO_FAIXA_4.aliquota
		exibeCalculoINSS(4, salario, TETO_FAIXA_3.salarioContribuicao, TETO_FAIXA_4.aliquota, deducaoParcial)

		return deducaoParcial + CalculaContribuicaoINSS(TETO_FAIXA_3.salarioContribuicao)
	}

	if salario > TETO_FAIXA_2.salarioContribuicao {
		// valor acima do teto da segunda faixa, por tanto dentro da terceira faixa:
		// - acima de R$ 2.666,68, até R$ 4.000,03
		deducaoParcial := (salario - TETO_FAIXA_2.salarioContribuicao) * TETO_FAIXA_3.aliquota
		exibeCalculoINSS(3, salario, TETO_FAIXA_2.salarioContribuicao, TETO_FAIXA_3.aliquota, deducaoParcial)

		return deducaoParcial + CalculaContribuicaoINSS(TETO_FAIXA_2.salarioContribuicao)
	}

	if salario > TETO_FAIXA_1.salarioContribuicao {
		// valor acima do teto da primeira faixa, por tanto dentro da segunda faixa:
		// - acima de R$ 1.412,00, até R$ 2.666,68
		deducalParcial := (salario - TETO_FAIXA_1.salarioContribuicao) * TETO_FAIXA_2.aliquota
		exibeCalculoINSS(2, salario, TETO_FAIXA_1.salarioContribuicao, TETO_FAIXA_2.aliquota, deducalParcial)

		return deducalParcial + CalculaContribuicaoINSS(TETO_FAIXA_1.salarioContribuicao)

	} else {
		// caso básico da função

		// salário dentro da primeira faixa, usa-se o valor do teto para o cálculo:
		// - até R$ 1.412,00
		deducaoParcial := TETO_FAIXA_1.salarioContribuicao * TETO_FAIXA_1.aliquota
		exibeCalculoINSS(1, TETO_FAIXA_1.salarioContribuicao, 0, TETO_FAIXA_1.aliquota, deducaoParcial)

		return deducaoParcial
	}
}

func exibeCalculoINSS(numeroFaixa int, salario, tetoFaixa, aliquota, deducao float64) {
	fmt.Printf("%dª faixa => R$ %s - R$ %s = R$ %s x %s = R$ %s\n",
		numeroFaixa,
		utils.FormatDecimalInBRL(salario),
		utils.FormatDecimalInBRL(tetoFaixa),
		utils.FormatDecimalInBRL(salario-tetoFaixa),
		utils.FormatDecimalInBRLWithPrecision(aliquota, 3),
		utils.FormatDecimalInBRL(deducao),
	)
}
