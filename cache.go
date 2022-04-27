package metaphoneptbr

import "sync"

//TODO documentação

type Cache struct {
	mu sync.RWMutex

	cache map[string]string

	ml  int
	sep rune
}

func NovoCache(max_length int, separator ...rune) *Cache {
	var sep rune
	if len(separator) > 0 {
		sep = separator[0]
	}

	return &Cache{
		cache: make(map[string]string),

		ml:  max_length,
		sep: sep,
	}
}

func (c *Cache) Metaphone_PTBR(s string) string {
	if res, ok := c.tentaRecup(s); ok {
		//cache hit
		return res
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	mp := Metaphone_PTBR_s(s, c.ml, c.sep)

	c.cache[s] = mp
	return mp
}

func (c *Cache) tentaRecup(s string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	r, o := c.cache[s]
	return r, o
}

func (c *Cache) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[string]string)
}
