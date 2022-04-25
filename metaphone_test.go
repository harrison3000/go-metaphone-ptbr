package metaphoneptbr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pares = map[string]string{
	"casa":           "KZ",
	"paralelepipedo": "PRLLPPD",
	"phone":          "FN",
	"teophilo":       "TFL",
	"alado":          "ALD",
	"caldo":          "KD",
	"fala":           "FL",
	"andar":          "AND2",
	"carro":          "K2",
	"rato":           "2T",
	"arara":          "ARR",
	"arsenico":       "ARSNK",
	"algoz":          "AGS",
	"zebra":          "ZBR",
	"azazel":         "AZZ",
	"avon":           "AVM",
	"manha":          "M3",
	"anna":           "AN",
	"assar":          "AS2",
	"sheila":         "XL",
	"asa":            "AZ",
	"ascender":       "ASND2",
	"mascavo":        "MSKV",
	"lascivia":       "LSV",
	"scheila":        "XL",
	"mesclado":       "MSLD",
	"fax":            "FX",
	"exonerar":       "EZNR2",
	"exército":       "EZRST",
	"méxico":         "MXK",
	"nexo":           "NKS",
	"alex":           "ALX",
	"texugo":         "TKSG",
	"exceção":        "ESS",
	"exceto":         "EST",
	"expatriar":      "ESPTR2",
	"experimento":    "ESPRMNT",
	"faxina":         "FKSN",
	"caixa":          "KX",
	"trouxe":         "TRX",
	"coxa":           "KX",
	"gaxeta":         "GXT",
	"laxante":        "LXNT",
	"roxo":           "2X",
	"xaxim":          "XXM",
	"taxi":           "TKS",
	"fixo":           "FKS",
	"enxame":         "ENXM",
}

func TestMeta(t *testing.T) {
	for k, v := range pares {
		res := Metaphone_PTBR(k, 99)
		assert.Equal(t, v, res, "Teste falhou em: %s", k)
	}

	resl := Metaphone_PTBR("paralelepipedo", 4)
	assert.Equal(t, "PRLL", resl, "Teste com limite falhou")

	ress := Metaphone_PTBR_s("odio.amor.cágado.jabutí.ryzen.coreissete", 99, '.')
	assert.Equal(t, "OD.AM2.KGD.JBT.2ZM.KRST", ress, "Teste com separador falhou")
}
