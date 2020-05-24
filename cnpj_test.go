//
// Autor: Marcelo Gomes Jr <marcelo.gomes.junior@gmail.com>
// Criado: abr/2020
//
// cpfcnpj
//

package cpfcnpj_test

import (
	"testing"

	"github.com/cll0ud/cpfcnpj"
)

func TestNewValidCNPJ(t *testing.T) {
	expected := "00000000000191"
	cnpj, err := cpfcnpj.NewValidCNPJ("00000000000191")
	expectNil(t, err)
	expect(t, string(cnpj), expected)

	expected = "00.000.000/0001-91"
	cnpj, err = cpfcnpj.NewValidCNPJ("00.000.000/0001-91")
	expectNil(t, err)
	expect(t, string(cnpj), expected)

	expected = ""
	cnpj, err = cpfcnpj.NewValidCNPJ("00.000.000/0001-92")
	expect(t, err, cpfcnpj.ErrCNPJInvalid)
	expect(t, string(cnpj), expected)
}

func TestCNPJ_IsValid(t *testing.T) {
	cnpjList := []cpfcnpj.CNPJ{
		"00.000.000/0001-91",
		"11.444.777/0001-61",
		"00.365.771/0001-82",
		"32.797.739/0001-62",
		"48.316.731/0001-77",
		"61.774.326/0001-60",
		"48.704.729/0001-75",
		"39.862.504/0001-56",
		"56.468.349/0001-07",
		"62.120.481/0001-26",
		"77.578.781/0001-20",
		"20.339.302/0001-04",
		"98.345.740/0001-64",
		"36.572.030/0001-10",
		"61.081.792/0001-60",
		"60.042.957/0001-22",
		"54.359.649/0001-22",
		"87.046.186/0001-06",
		"94.201.645/0001-36",
		"42.633.594/0001-18",
		"40.635.841/0001-90",
		"86.761.822/0001-00",
		"78.707.461/0001-96",
		"19.613.446/0001-10",
		"76.510.184/0001-00",
	}
	for _, c := range cnpjList {
		valid := c.IsValid()
		if valid {
			t.Logf("'%s'... ok", c)
		} else {
			t.Errorf("'%s'... fail", c)
		}
	}
}

func TestCNPJ_Format(t *testing.T) {
	expected := "00.000.000/0001-91"
	cnpj, err := cpfcnpj.NewValidCNPJ("00000000000191")
	expectNil(t, err)
	expectNil(t, cnpj.Format())
	expect(t, string(cnpj), expected)
}

func TestCNPJ_Unformat(t *testing.T) {
	expected := "00000000000191"
	cnpj, err := cpfcnpj.NewValidCNPJ("00.000.000/0001-91")
	expectNil(t, err)
	expectNil(t, cnpj.Unformat())
	expect(t, string(cnpj), expected)
}
