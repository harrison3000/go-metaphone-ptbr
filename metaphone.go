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

func Metaphone_PTBR_s(s string, max_length int, separator rune) string {
	s = makeUpperAndClean(s)

	primary := &strings.Builder{}

	MetaphAddChr := (*strings.Builder).WriteRune
	WORD_EDGE := func(c rune) bool {
		return c == 0 || unicode.IsSpace(c) || c == separator
	}
	GetAt := getAt

	var last_char rune

	original := []rune(s)
	original = append(original, 0) //NULL terminator falso, pra WORD_EDGE funcionar, entre outras coisas

	//Neste loop eu tentei ao maximo preservar os nome de variaveis e
	//comentários da versão original em C
	for current := 0; primary.Len() < max_length && current < len(original); {
		current_char := GetAt(original, current)

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
		case 'L':
			ahead_char := GetAt(original, current+1)
			/* lha, lho. Adicionado 2009-11-09. Thx Peter Krauss. Ele estava mal-colocado */
			if ahead_char == 'H' {
				MetaphAddChr(primary, '1')
			} else if isVowel(ahead_char) || WORD_EDGE(last_char) {
				/* como em Louco, aloprado, alado, lampada, etc */
				MetaphAddChr(primary, 'L')
			}
			/* atualmente ignora L antes de consoantes */

		case 'T', 'P':
			/* those are special cases, from foreign names or
			 * old portuguese names sintax.
			 * Besides, should behavior as the others.
			 */
			ahead_char := GetAt(original, current+1)
			if ahead_char == 'H' {
				/* phone, pharmacia, teophilo */
				if current_char == 'P' {
					MetaphAddChr(primary, 'F')
				} else {
					MetaphAddChr(primary, 'T')
					current++
				}
			}
			//TODO continuar tradução
			_ = current
		}

		/* next char */
		current++
		last_char = current_char
	}

	return primary.String()
}
