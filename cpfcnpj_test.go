//
// Autor: Marcelo Gomes Jr <marcelo.gomes.junior@gmail.com>
// Criado: abr/2020
//
// cpfcnpj
//

package cpfcnpj_test

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	t.Helper()
	if a != b {
		t.Errorf("Expected: '%#v' (%v) - Got: '%#v' (%v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func expectNil(t *testing.T, a interface{}) {
	t.Helper()
	if a != nil {
		t.Errorf("Expected: '%#v' - Got: '%#v' (%v)", nil, a, reflect.TypeOf(a))
	}
}
