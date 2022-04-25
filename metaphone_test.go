package metaphoneptbr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pares = map[string]string{
	"casa": "KZ",
}

func TestMeta(t *testing.T) {
	for k, v := range pares {
		res := Metaphone_PTBR(k, 9999)
		assert.Equal(t, v, res)
	}

	ress := Metaphone_PTBR_s("odio.amor.cágado.jabutí.ryzen.coreissete", 99, '.')
	assert.Equal(t, "OD.AM2.KGD.JBT.2ZM.KRST", ress)
}
