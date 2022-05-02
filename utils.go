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

import "strings"

func tiraAcentos(c rune) rune {
	switch c {
	case 'Á', 'À', 'Ã', 'Â', 'Ä':
		return 'A'

	case 'É', 'È', 'Ẽ', 'Ê', 'Ë':
		return 'E'

	case 'Y', 'Í', 'Ì', 'Ĩ', 'Î', 'Ï':
		return 'I'

	case 'Ó', 'Ò', 'Õ', 'Ô', 'Ö':
		return 'O'

	case 'Ú', 'Ù', 'Ũ', 'Û', 'Ü':
		return 'U'
	}

	return c
}

func makeUpperAndClean(s string) []rune {
	//Não tinha isso no original mas coloquei mesmo assim
	//TODO validar
	s = strings.TrimSpace(s)

	//maiuscula e sem acentos
	s = strings.ToUpper(s)
	s = strings.Map(tiraAcentos, s)

	res := make([]rune, 0, 24)

	var ultimo rune
	for _, v := range s {
		if ultimo == v && v != 'R' && v != 'S' {
			continue
		}

		ultimo = v
		res = append(res, v)
	}

	return res
}

func isVowel(c rune) bool {
	switch c {
	case 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func getAt(r []rune, i int) rune {
	if i < 0 || i >= len(r) {
		return 0
	}
	return r[i]
}
