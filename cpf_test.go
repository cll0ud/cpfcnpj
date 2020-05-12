//
// Autor: Marcelo Gomes Jr <marcelo.gomes.junior@gmail.com>
// Criado: abr/2020
//
// cpfcnpj
//

package cpfcnpj_test

import (
	"testing"

	"cpfcnpj"
)

func TestNewValidCPF(t *testing.T) {
	expected := "11144477735"

	cpf, err := cpfcnpj.NewValidCPF("11144477735")
	expectNil(t, err)
	expect(t, string(cpf), expected)

	expected = "111.444.777-35"
	cpf, err = cpfcnpj.NewValidCPF("111.444.777-35")
	expectNil(t, err)
	expect(t, string(cpf), expected)

	expected = ""
	cpf, err = cpfcnpj.NewValidCPF("111.444.777-36")
	expect(t, err, cpfcnpj.ErrCPFInvalid)
	expect(t, string(cpf), expected)
}

func TestCPF_IsValid(t *testing.T) {
	cpfList := []cpfcnpj.CPF{
		"111.444.777-35",
		"621.283.446-62",
		"072.512.964-62",
		"477.811.768-98",
		"145.442.945-33",
		"627.348.187-36",
		"534.537.926-29",
		"383.742.685-81",
		"567.888.244-95",
		"938.606.954-79",
		"611.077.866-49",
		"667.477.577-00",
		"948.846.886-60",
		"067.932.419-40",
		"386.916.779-37",
		"957.252.429-19",
		"816.476.876-67",
	}
	for _, c := range cpfList {
		valid := c.IsValid()
		if valid {
			t.Logf("'%s'... ok", c)
		} else {
			t.Errorf("'%s'... fail", c)
		}
	}
}

func TestCPF_Format(t *testing.T) {
	expected := "111.444.777-35"
	cpf, err := cpfcnpj.NewValidCPF("11144477735")
	expectNil(t, err)
	expectNil(t, cpf.Format())
	expect(t, string(cpf), expected)
}

func TestCPF_Unformat(t *testing.T) {
	expected := "11144477735"
	cpf, err := cpfcnpj.NewValidCPF("11144477735")
	expectNil(t, err)
	expectNil(t, cpf.Unformat())
	expect(t, string(cpf), expected)
}
