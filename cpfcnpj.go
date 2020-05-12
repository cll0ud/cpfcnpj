//
// Autor: Marcelo Gomes Jr <marcelo.gomes.junior@gmail.com>
// Criado: abr/2020
//
// cpfcnpj
//

package cpfcnpj

import (
	"errors"
	"fmt"
	"regexp"
)

// isValid checa se o CPF ou CNPJ é válido. A maioria dos processos
// são similares, somente o tamanho da string e o peso inicial da fórmula
// são diferentes
func isValid(v string, size, formattedSize int) bool {
	// Se o tamanho da string for inválido, retorna false imediatamente
	if v == "" || (len(v) != size && len(v) != formattedSize) {
		return false
	}

	// Remove todos os caracteres que não são digitos e checa se o tamanho
	// do resultado é válido
	plain, err := unformat(v, size, errors.New("formato inválido"))
	if err != nil {
		return false
	}

	// CPFs e CNPJs comspostos somente pelo mesmo número são considerados inválidos
	// (eg.: '111.111.111-11', '11.111.111/1111-11')
	// www.receita.fazenda.gov.br/publico/Legislacao/atos/AtosConjuntos/AnexoIADEConjuntoCoratCotec0012002.doc
	if isSameChar(plain) {
		return false
	}

	// Calcula os dois dígitos verificadores de acordo com o tamanho esperado
	// e depois guarda o valor em 'result'
	// 11 = CPF
	// 14 = CNPJ
	var result string
	switch size {
	case 11:
		d1 := calculateDigit(plain[0:9], 10)
		d2 := calculateDigit(plain[0:10], 11)
		result = fmt.Sprintf("%s%s%s", plain[0:9], d1, d2)
	case 14:
		d1 := calculateDigit(plain[0:12], 5)
		d2 := calculateDigit(plain[0:13], 6)
		result = fmt.Sprintf("%s%s%s", plain[0:12], d1, d2)
	}

	// Se o cálculo está correto, 'result' é igual a 'plain'
	// o que significa que o CPF/CNPJ informado é válido
	return plain == result
}

// isSameChar retorna true se todos os caracteres em uma dada string
// forem iguais
func isSameChar(v string) bool {
	first := v[0]
	for i := 1; i < len(v); i++ {
		if first != v[i] {
			return false
		}
	}
	return true
}

// calculateDigit calcula o dígito tendo como base a string informada
// e começando no peso informado.
//
// O cálculo é similar tanto para CPF e CNPJ e utiliza a seguinte fórmula.
// Para cada dígito verificador, é feita uma série de multiplicações e os
// resultados são somados. Depois eles são divididos por 11 e o resto da
// divisão é usado para definir o dígito verificador.
//
// CPF:  para o primeiro dígito verificador, multiplica cada um
//       dos primeiros 9 dígitos do CPF por uma tabela de pesos pré-definida
//       [10, 9, 8, 7, 6, 5, 4, 3, 2].
//       Para o segundo dígito, inclui o primeiro resultado e multiplica os
//       10 dígitos do CPF pela mesma tabela, mas dessa vez começa com 11
//       [11, 10, 9, 8, 7, 6, 5, 4, 3, 2].
//
// CNPJ: para o primeiro dígito verificador, multiplica cada um dos
//       primeiros 12 dígitos pela tabela de pesos:
//       [5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2].
//       Para o segundo dígito, inclui o primeiro resultado e multiplica
//       os 13 dígitos por uma tabela similar mas que começa com 6:
//       [6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2].
//
// Soma-se todos os resultados de cada multiplicação e este valor é dividido
// por 11. O restante da divisão define qual o dígito verificador:
// - Se o resto for menor que 2, o dígito é 0
// - Se o resto for igual ou superior à 2, o dígito é o resultado de (11 - resto)
//
func calculateDigit(v string, weight int) string {
	sum := 0
	for _, char := range v {
		// runes em go guardam um valor numérico que representa
		// o caractere em questão na tabela ASCII então se você tentar
		// converter diretamente para int utilizando `int(rune)` o valor
		// retornado para um número será algo entre 48 (0) e 57 (9).
		// É possível recuperar o caractere utilizando conversões de texto
		// mas como queremos apenas um número existe um jeito simples de
		// resolver: recuperar o valor int (entre 48 e 57 como dito acima)
		// e subtrair o valor de '0' (48), isso vai retornar o número que
		// queremos
		sum += int(char-'0') * weight
		weight--
		if weight < 2 {
			weight = 9
		}
	}
	sum %= 11
	if sum < 2 {
		return "0"
	}
	return fmt.Sprintf("%d", 11-sum)
}

// format devolve uma string formatada de acordo com o padrão recebido
// retornando erro caso aconteça algum problema durante o processo
func format(v, pattern string, size int, rerr error) (string, error) {
	// se o resultado bater com o padrão, significa que já está formatado
	// então só devolve o valor recebido
	if regexp.MustCompile(pattern).MatchString(v) {
		return v, nil
	}
	// remove caracteres não-numéricos e checa se o tamanho do resultado
	// bate com o esperado
	plain, err := unformat(v, size, rerr)
	if err != nil {
		return v, err
	}
	// formata a string de acordo com o tamanho esperado
	// 11 = CPF
	// 14 = CNPJ
	switch size {
	case 11:
		return fmt.Sprintf("%s.%s.%s-%s", plain[0:3], plain[3:6], plain[6:9], plain[9:11]), nil
	case 14:
		return fmt.Sprintf("%s.%s.%s/%s-%s", plain[0:2], plain[2:5], plain[5:8], plain[8:12], plain[12:14]), nil
	}
	return v, nil
}

// unformat devolve uma string sem caractéres não-numéricos
// retornando erro caso o resultado não bata com o tamanho
// informado
func unformat(v string, size int, err error) (string, error) {
	plain := regexp.MustCompile(`[^\d]+`).ReplaceAllString(v, "")
	if len(plain) != size {
		return v, err
	}
	return plain, nil
}

// patternOrSize retorna false caso a string não bata com o padrão
// nem com o tamanho fornecidos
func patternOrSize(v, pattern string, size int) bool {
	// Retorna caso a string bata com o padrão fornecido
	if regexp.MustCompile(pattern).MatchString(v) {
		return true
	}
	// Retorna uma string vazia caso o valor informado não bata com
	// o tamanho esperado
	if len(v) != size {
		return false
	}
	return true
}
