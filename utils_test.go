package metaphoneptbr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimpeza(t *testing.T) {
	s1 := string(makeUpperAndClean("eita é teste"))
	s2 := "EITA E TESTE"

	assert.Equal(t, s1, s2)

	d1 := string(makeUpperAndClean("carro passa lojja LlLl"))
	d2 := "CARRO PASSA LOJA L"

	assert.Equal(t, d1, d2)
}
