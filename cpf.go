//
// Autor: Marcelo Gomes Jr <marcelo.gomes.junior@gmail.com>
// Criado: abr/2020
//
// cpfcnpj
//

package cpfcnpj

import (
	"errors"
)

const (
	cpfPattern       = `^\d{3}\.\d{3}\.\d{3}-\d{2}$`
	cpfSize          = 11
	cpfFormattedSize = 14
)

var (
	ErrCPFInvalidFormat = errors.New("o cpf deve ter 11 digitos (ou 14 com pontuação)")
	ErrCPFInvalid       = errors.New("cpf inválido")
)

// CPF representa o documento Cadastro de Pessoas Físicas, registro
// mantido pela Receita Federal do Brail
//
// Não é necessário utilizar as funções 'NewCPF' ou 'NewValidCPF', elas
// são apenas helpers. É possível criar um objeto CPF diretamente com uma
// simples conversão (cpfcnpj.CPF("11144477735") ou cpfcnpj.CPF("111.444.777-35"))
type CPF string

// NewValidCPF retorna um CPF baseado na string informada caso ele seja válido
func NewValidCPF(v string) (CPF, error) {
	if !patternOrSize(v, cpfPattern, cpfSize) {
		return "", ErrCPFInvalidFormat
	}
	cpf := CPF(v)
	if !cpf.IsValid() {
		return "", ErrCPFInvalid
	}
	return cpf, nil
}

// IsValid checa se o CPF é válido checando através de um algoritmo
// se os dois últimos digitos são válidos de acordo com o resto
// da string
func (c *CPF) IsValid() bool {
	return isValid(string(*c), cpfSize, cpfFormattedSize)
}

// Format formata o CPF no padrão "999.999.999-99" e retorna
// erro caso ocorra algum problema durante o processo
func (c *CPF) Format() error {
	v, err := format(string(*c), cpfPattern, cpfSize, ErrCPFInvalidFormat)
	if err != nil {
		return err
	}
	*c = CPF(v)
	return nil
}

// Unformat remove a pontuação ("99999999999") do CPF e
// retorna erro caso ocorra algum problema durante o processo
func (c *CPF) Unformat() error {
	v, err := unformat(string(*c), cpfSize, ErrCPFInvalidFormat)
	if err != nil {
		return err
	}
	*c = CPF(v)
	return nil
}
