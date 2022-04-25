package metaphoneptbr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimpeza(t *testing.T) {
	s1 := makeUpperAndClean("eita é teste")
	s2 := "EITA E TESTE"

	assert.Equal(t, s1, s2)

	d1 := makeUpperAndClean("carro passa lojja")
	d2 := "CARRO PASSA LOJA"

	assert.Equal(t, d1, d2)
}
