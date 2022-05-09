package metaphoneptbr

import "testing"

func BenchmarkCache(b *testing.B) {
	srts := []string{
		"oi",
		"Oi",
		"beleza",
		"Francisco",
		"de",
		"Assis",
		"francisco",
		"BéLEzÁ",
	}

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		cache := NovoCache(20)
		b.StartTimer()
		for _, v := range srts {
			cache.Metaphone_PTBR(v)
		}
	}
}
