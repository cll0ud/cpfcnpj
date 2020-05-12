//
// Autor: Marcelo Gomes Jr <marcelo.gomes.junior@gmail.com>
// Criado: abr/2020
//
// golang
//

package cpfcnpj

import (
	"errors"
)

const (
	cnpjPattern       = `^\d{2}\.\d{3}\.\d{3}\/\d{4}\-\d{2}$`
	cnpjSize          = 14
	cnpjFormattedSize = 18
)

var (
	ErrCNPJInvalidFormat = errors.New("o cnpj deve ter 14 digitos (ou 18 com pontuação)")
	ErrCNPJInvalid       = errors.New("cnpj inválido")
)

// CNPJ representa o documento Cadastro Nacional da Pessoa Jurídica
//
// Não é necessária a utilização das funções 'NewCNPJ' ou 'NewValidCNPJ',
// elas são apenas helpers e é possível criar um objeto CNPJ duretamente com
// uma simples conversão (cpfcnpj.CNPJ("00000000000191") ou cpfcnpj.CNPJ("00.000.000/0001-91"))
type CNPJ string

// NewValidCNPJ retorna um CNPJ baseado na string informada, caso ele
// obedeça o formato correto e seja válido
func NewValidCNPJ(v string) (CNPJ, error) {
	if !patternOrSize(v, cnpjPattern, cnpjSize) {
		return "", ErrCNPJInvalidFormat
	}
	cnpj := CNPJ(v)
	if !cnpj.IsValid() {
		return "", ErrCNPJInvalid
	}
	return cnpj, nil
}

// IsValid checa se o CNPJ é válido checando através de um algoritmo
// se os dois últimos digitos são válidos de acordo com o resto da
// string
func (c *CNPJ) IsValid() bool {
	return isValid(string(*c), cnpjSize, cnpjFormattedSize)
}

// Format formata o CNPJ no padrão "00.000.000/0001-91" retornando erro
// caso ocorra algum problema durante o processo
func (c *CNPJ) Format() error {
	v, err := format(string(*c), cnpjPattern, cnpjSize, ErrCNPJInvalidFormat)
	if err != nil {
		return err
	}
	*c = CNPJ(v)
	return nil
}

// Unformat remove a pontuação do CNPJ ("00000000000191") retornando
// erro caso ocorra algum problema durante o processo
func (c *CNPJ) Unformat() error {
	v, err := unformat(string(*c), cnpjSize, ErrCNPJInvalidFormat)
	if err != nil {
		return err
	}
	*c = CNPJ(v)
	return nil
}
