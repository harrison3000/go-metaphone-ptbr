package metaphoneptbr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Valores calculados usando o programa em C
// foi só mudar a linha de printar o resultado para: printf("\"%s\":\"%s\",\n", argv[count-1], code);
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
	"Jacques":        "JKS",
	"aranha":         "AR3",
	"chuva":          "XV",
	"caçada":         "KSD",
	"quero":          "KR",
	"quase":          "KZ",
	"alho":           "A1",
	"theos":          "TS",
	"gente":          "JNT",
	"girar":          "JR2",
	"gosto":          "GST",
	"gheto":          "JT",
	"ghi":            "J",
	"gho":            "J",
	"christiano":     "KRSTN",
	"hoje":           "OJ",
	"homem":          "OMM",
	"loha":           "L",
	"hosana":         "OSN",
	"Wladimir":       "VLDM2",
	"vladmir":        "VLDM2",
	"welington":      "VLNGTM",
	"ótimo":          "OTM",
	"último":         "UTM",
	"ghst":           "GJST",
}

func TestMeta(t *testing.T) {
	for k, v := range pares {
		res := Metaphone(k)
		assert.Equal(t, v, res, "Teste falhou em: %s", k)
	}

	ress := Metaphone_s("odio.amor.cágado.jabutí.ryzen.coreissete", '.')
	assert.Equal(t, "OD.AM2.KGD.JBT.2ZM.KRST", ress, "Teste com separador falhou")
}

func BenchmarkMeta(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Metaphone("Anticonstitucionalissimamente")
	}
}
