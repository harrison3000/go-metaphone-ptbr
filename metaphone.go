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
	MetaphAdd := (*strings.Builder).WriteString
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
				}
				current++
				break
			}
			fallthrough
		case 'B', 'D', 'F', 'J', 'K', 'M', 'V':
			MetaphAddChr(primary, current_char)
		/* checar consoantes com som confuso e similares */
		case 'G':
			ahead_char := GetAt(original, current+1)
			switch ahead_char {
			case 'H':
				/* H sempre complica a vida. Se não for vogal, tratar como 'G',
				   caso contrário segue o fluxo abaixo. */
				if !isVowel(GetAt(original, current+2)) {
					MetaphAddChr(primary, 'G')
				}
				fallthrough //ta certo? TODO testar
			case 'E', 'I':
				MetaphAddChr(primary, 'J')
			default:
				MetaphAddChr(primary, 'G')
			}

		case 'R':
			ahead_char := GetAt(original, current+1)

			/* como em andar, carro, rato */
			if WORD_EDGE(last_char) || WORD_EDGE(ahead_char) {
				MetaphAddChr(primary, '2')
			} else if ahead_char == 'R' {
				MetaphAddChr(primary, '2')
				current++
			} else if isVowel(last_char) && isVowel(ahead_char) {
				/* como em arara */
				MetaphAddChr(primary, 'R')
				current++
			} else {
				/* todo o resto, como em arsenico */
				MetaphAddChr(primary, 'R')
			}

		case 'Z':
			ahead_char := GetAt(original, current+1)

			if WORD_EDGE(ahead_char) {
				/* termina com, como em algoz */
				MetaphAddChr(primary, 'S')
			} else {
				MetaphAddChr(primary, 'Z')
			}

		case 'N':
			ahead_char := GetAt(original, current+1)

			/* no português, todas as palavras terminam com 'M', exceto
			 * no caso de nomes próprios, ou estrangeiros. Para todo caso,
			 * tem som de 'M'
			 */
			if WORD_EDGE(ahead_char) {
				MetaphAddChr(primary, 'M')
			} else if ahead_char == 'H' {
				/* aranha, nhoque, manha */
				MetaphAddChr(primary, '3')
				current++
			} else if last_char != 'N' {
				/* duplicado... */
				MetaphAddChr(primary, 'N')
			}

		case 'S':
			ahead_char := GetAt(original, current+1)

			if ahead_char == 'S' {
				/* aSSar */
				MetaphAddChr(primary, 'S')
				last_char = ahead_char
				current++
			} else if ahead_char == 'H' {
				/* mais estrangeirismo: sheila, mishel, e compatibilidade sonora com sobrenomes estrangeiros (japoneses) */
				MetaphAddChr(primary, 'X')
				current++
			} else if isVowel(last_char) && isVowel(ahead_char) {
				/* como em asa */
				MetaphAddChr(primary, 'Z')
			} else if ahead_char == 'C' {
				/* special cases = 'SC' */
				ahead2_char := GetAt(original, current+2)
				switch ahead2_char {
				case 'E', 'I':
					/* aSCEnder, laSCIvia */
					MetaphAddChr(primary, 'S')
					current += 2

				case 'A', 'O', 'U':
					/* maSCAvo, aSCO, auSCUltar */
					MetaphAdd(primary, "SK")
					current += 2

				case 'H':
					/* estrangeirismo tal como scheila. */
					MetaphAddChr(primary, 'X')
					current += 2

				default:
					/* mesclado */
					MetaphAddChr(primary, 'S')
					current++
				}
			} else {
				/* catch all - deve pegar atrás e sapato */
				MetaphAddChr(primary, 'S')
			}

		}

		/* next char */
		current++
		last_char = current_char
	}

	return primary.String()
}
