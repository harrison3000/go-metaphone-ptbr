/*
  Copyright 2022, Hárrison Leão Ferreira <harrisonf100@gmail.com>.
  All rights reserved.

  Redistribution and use in source and binary forms, with or without modification,
  are permitted provided that the following conditions are met,
  1. Redistributions of source code must retain the above copyright notice, this
     list of conditions and the following disclaimer.
  2. Redistributions in binary form must reproduce the above copyright notice, this
     list of conditions and the following disclaimer in the documentation and/or
     other materials provided with the distribution.


  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
  ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
  ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


***********************************************************************/

package metaphoneptbr

import (
	"strings"
	"unicode"
)

func Metaphone_PTBR_s(s string, separator rune) string {
	s = makeUpperAndClean(s)

	primary := &strings.Builder{}

	MetaphAddChr := (*strings.Builder).WriteRune
	WORD_EDGE := func(c rune) bool {
		return c == 0 || unicode.IsSpace(c) || c == separator
	}

	var last_char rune

	r := []rune(s)
	r = append(r, 0) //NULL terminator falso, pra WORD_EDGE funcionar, entre outras coisas

	//Neste loop eu tentei ao maximo preservar os nome de variaveis e
	//comentários da versão original em C
	for current, current_char := range r {

		if separator == current_char {
			MetaphAddChr(primary, separator)
			last_char = current_char
			continue
		}

		switch current_char {
		case 'A', 'E', 'I', 'O', 'U':
			/* initials vowels after any space must stay too */
			if WORD_EDGE(last_char) {
				MetaphAddChr(primary, current_char)
			}

			//TODO continuar tradução
			_ = current
		}

		last_char = current_char
	}

	return primary.String()
}
