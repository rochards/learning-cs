package irpf

import (
	"fmt"
	"main/utils"
)

type ContribuicaoIRPF struct {
	valorBase float64
	aliquota  float64
	deducao   float64
}

// dados da tabela de contribuição com valores de 2024
// https://www.gov.br/receitafederal/pt-br/assuntos/meu-imposto-de-renda/tabelas/2024
// var BASE_FAIXA_1 = ContribuicaoIRPF{valorBase: 0, aliquota: 0, deducao: 0}
var BASE_FAIXA_2 = ContribuicaoIRPF{valorBase: 2259.21, aliquota: 7.5 / 100, deducao: 169.44}
var BASE_FAIXA_3 = ContribuicaoIRPF{valorBase: 2826.66, aliquota: 15.0 / 100, deducao: 381.44}
var BASE_FAIXA_4 = ContribuicaoIRPF{valorBase: 3751.06, aliquota: 22.5 / 100, deducao: 662.77}
var BASE_FAIXA_5 = ContribuicaoIRPF{valorBase: 4664.69, aliquota: 27.5 / 100, deducao: 896.00}

const DEDUCAO_POR_DEPENDENTE = 189.59

func CalculaContribuicaoIRPF(salario, deducaoINSS float64, numeroFilhos int) float64 {

	deducaoDependentes := DEDUCAO_POR_DEPENDENTE * float64(numeroFilhos)
	baseCalculo := salario - deducaoINSS - deducaoDependentes

	fmt.Println("== Cálculos referentes ao IRPF ==")
	fmt.Printf("-> Salário bruto: R$ %s\n"+
		"-> Dedução INSS: R$ %s\n"+
		"-> Dedução pelos dependentes: %d x R$ %s = %s\n"+
		"-> Base de cálculo: R$ %s - R$ %s - R$ %s = R$ %s\n",
		utils.FormatDecimalInBRL(salario),
		utils.FormatDecimalInBRL(deducaoINSS),
		numeroFilhos, utils.FormatDecimalInBRL(DEDUCAO_POR_DEPENDENTE), utils.FormatDecimalInBRL(deducaoDependentes),
		utils.FormatDecimalInBRL(salario), utils.FormatDecimalInBRL(deducaoINSS), utils.FormatDecimalInBRL(deducaoDependentes), utils.FormatDecimalInBRL(baseCalculo),
	)

	if baseCalculo >= BASE_FAIXA_5.valorBase {
		return calculaIRPF(baseCalculo, BASE_FAIXA_5)
	}

	if baseCalculo >= BASE_FAIXA_4.valorBase {
		return calculaIRPF(baseCalculo, BASE_FAIXA_4)
	}

	if baseCalculo >= BASE_FAIXA_3.valorBase {
		return calculaIRPF(baseCalculo, BASE_FAIXA_3)
	}

	if baseCalculo >= BASE_FAIXA_2.valorBase {
		return calculaIRPF(baseCalculo, BASE_FAIXA_2)
	}

	return 0 // faixa 1 é insento
}

func calculaIRPF(salarioBase float64, contribuicaoIRPF ContribuicaoIRPF) float64 {

	impostoTotal := salarioBase*contribuicaoIRPF.aliquota - contribuicaoIRPF.deducao

	fmt.Printf("==> IRPF: R$ %s x %s - R$ %s = R$ %s\n",
		utils.FormatDecimalInBRL(salarioBase),
		utils.FormatDecimalInBRLWithPrecision(contribuicaoIRPF.aliquota, 3),
		utils.FormatDecimalInBRL(contribuicaoIRPF.deducao),
		utils.FormatDecimalInBRL(impostoTotal),
	)

	return impostoTotal
}
