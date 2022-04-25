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
}
