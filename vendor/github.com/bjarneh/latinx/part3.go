// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

// ISO 8859-3

// Generated from http://unicode.org/Public/MAPPINGS/ISO8859/8859-3.TXT

func newISO_8859_3Converter() *Converter {

    converter := new(Converter)
    converter.id = "Latin-3/ISO-8859-3"
    converter.latinToUtf8 = make(map[byte][]byte)

    converter.latinToUtf8[160] = []byte(" ") // NO-BREAK SPACE
    converter.latinToUtf8[161] = []byte("Ħ") // LATIN CAPITAL LETTER H WITH STROKE
    converter.latinToUtf8[162] = []byte("˘") // BREVE
    converter.latinToUtf8[163] = []byte("£") // POUND SIGN
    converter.latinToUtf8[164] = []byte("¤") // CURRENCY SIGN
    converter.latinToUtf8[166] = []byte("Ĥ") // LATIN CAPITAL LETTER H WITH CIRCUMFLEX
    converter.latinToUtf8[167] = []byte("§") // SECTION SIGN
    converter.latinToUtf8[168] = []byte("¨") // DIAERESIS
    converter.latinToUtf8[169] = []byte("İ") // LATIN CAPITAL LETTER I WITH DOT ABOVE
    converter.latinToUtf8[170] = []byte("Ş") // LATIN CAPITAL LETTER S WITH CEDILLA
    converter.latinToUtf8[171] = []byte("Ğ") // LATIN CAPITAL LETTER G WITH BREVE
    converter.latinToUtf8[172] = []byte("Ĵ") // LATIN CAPITAL LETTER J WITH CIRCUMFLEX
    converter.latinToUtf8[173] = []byte("­") // SOFT HYPHEN
    converter.latinToUtf8[175] = []byte("Ż") // LATIN CAPITAL LETTER Z WITH DOT ABOVE
    converter.latinToUtf8[176] = []byte("°") // DEGREE SIGN
    converter.latinToUtf8[177] = []byte("ħ") // LATIN SMALL LETTER H WITH STROKE
    converter.latinToUtf8[178] = []byte("²") // SUPERSCRIPT TWO
    converter.latinToUtf8[179] = []byte("³") // SUPERSCRIPT THREE
    converter.latinToUtf8[180] = []byte("´") // ACUTE ACCENT
    converter.latinToUtf8[181] = []byte("µ") // MICRO SIGN
    converter.latinToUtf8[182] = []byte("ĥ") // LATIN SMALL LETTER H WITH CIRCUMFLEX
    converter.latinToUtf8[183] = []byte("·") // MIDDLE DOT
    converter.latinToUtf8[184] = []byte("¸") // CEDILLA
    converter.latinToUtf8[185] = []byte("ı") // LATIN SMALL LETTER DOTLESS I
    converter.latinToUtf8[186] = []byte("ş") // LATIN SMALL LETTER S WITH CEDILLA
    converter.latinToUtf8[187] = []byte("ğ") // LATIN SMALL LETTER G WITH BREVE
    converter.latinToUtf8[188] = []byte("ĵ") // LATIN SMALL LETTER J WITH CIRCUMFLEX
    converter.latinToUtf8[189] = []byte("½") // VULGAR FRACTION ONE HALF
    converter.latinToUtf8[191] = []byte("ż") // LATIN SMALL LETTER Z WITH DOT ABOVE
    converter.latinToUtf8[192] = []byte("À") // LATIN CAPITAL LETTER A WITH GRAVE
    converter.latinToUtf8[193] = []byte("Á") // LATIN CAPITAL LETTER A WITH ACUTE
    converter.latinToUtf8[194] = []byte("Â") // LATIN CAPITAL LETTER A WITH CIRCUMFLEX
    converter.latinToUtf8[196] = []byte("Ä") // LATIN CAPITAL LETTER A WITH DIAERESIS
    converter.latinToUtf8[197] = []byte("Ċ") // LATIN CAPITAL LETTER C WITH DOT ABOVE
    converter.latinToUtf8[198] = []byte("Ĉ") // LATIN CAPITAL LETTER C WITH CIRCUMFLEX
    converter.latinToUtf8[199] = []byte("Ç") // LATIN CAPITAL LETTER C WITH CEDILLA
    converter.latinToUtf8[200] = []byte("È") // LATIN CAPITAL LETTER E WITH GRAVE
    converter.latinToUtf8[201] = []byte("É") // LATIN CAPITAL LETTER E WITH ACUTE
    converter.latinToUtf8[202] = []byte("Ê") // LATIN CAPITAL LETTER E WITH CIRCUMFLEX
    converter.latinToUtf8[203] = []byte("Ë") // LATIN CAPITAL LETTER E WITH DIAERESIS
    converter.latinToUtf8[204] = []byte("Ì") // LATIN CAPITAL LETTER I WITH GRAVE
    converter.latinToUtf8[205] = []byte("Í") // LATIN CAPITAL LETTER I WITH ACUTE
    converter.latinToUtf8[206] = []byte("Î") // LATIN CAPITAL LETTER I WITH CIRCUMFLEX
    converter.latinToUtf8[207] = []byte("Ï") // LATIN CAPITAL LETTER I WITH DIAERESIS
    converter.latinToUtf8[209] = []byte("Ñ") // LATIN CAPITAL LETTER N WITH TILDE
    converter.latinToUtf8[210] = []byte("Ò") // LATIN CAPITAL LETTER O WITH GRAVE
    converter.latinToUtf8[211] = []byte("Ó") // LATIN CAPITAL LETTER O WITH ACUTE
    converter.latinToUtf8[212] = []byte("Ô") // LATIN CAPITAL LETTER O WITH CIRCUMFLEX
    converter.latinToUtf8[213] = []byte("Ġ") // LATIN CAPITAL LETTER G WITH DOT ABOVE
    converter.latinToUtf8[214] = []byte("Ö") // LATIN CAPITAL LETTER O WITH DIAERESIS
    converter.latinToUtf8[215] = []byte("×") // MULTIPLICATION SIGN
    converter.latinToUtf8[216] = []byte("Ĝ") // LATIN CAPITAL LETTER G WITH CIRCUMFLEX
    converter.latinToUtf8[217] = []byte("Ù") // LATIN CAPITAL LETTER U WITH GRAVE
    converter.latinToUtf8[218] = []byte("Ú") // LATIN CAPITAL LETTER U WITH ACUTE
    converter.latinToUtf8[219] = []byte("Û") // LATIN CAPITAL LETTER U WITH CIRCUMFLEX
    converter.latinToUtf8[220] = []byte("Ü") // LATIN CAPITAL LETTER U WITH DIAERESIS
    converter.latinToUtf8[221] = []byte("Ŭ") // LATIN CAPITAL LETTER U WITH BREVE
    converter.latinToUtf8[222] = []byte("Ŝ") // LATIN CAPITAL LETTER S WITH CIRCUMFLEX
    converter.latinToUtf8[223] = []byte("ß") // LATIN SMALL LETTER SHARP S
    converter.latinToUtf8[224] = []byte("à") // LATIN SMALL LETTER A WITH GRAVE
    converter.latinToUtf8[225] = []byte("á") // LATIN SMALL LETTER A WITH ACUTE
    converter.latinToUtf8[226] = []byte("â") // LATIN SMALL LETTER A WITH CIRCUMFLEX
    converter.latinToUtf8[228] = []byte("ä") // LATIN SMALL LETTER A WITH DIAERESIS
    converter.latinToUtf8[229] = []byte("ċ") // LATIN SMALL LETTER C WITH DOT ABOVE
    converter.latinToUtf8[230] = []byte("ĉ") // LATIN SMALL LETTER C WITH CIRCUMFLEX
    converter.latinToUtf8[231] = []byte("ç") // LATIN SMALL LETTER C WITH CEDILLA
    converter.latinToUtf8[232] = []byte("è") // LATIN SMALL LETTER E WITH GRAVE
    converter.latinToUtf8[233] = []byte("é") // LATIN SMALL LETTER E WITH ACUTE
    converter.latinToUtf8[234] = []byte("ê") // LATIN SMALL LETTER E WITH CIRCUMFLEX
    converter.latinToUtf8[235] = []byte("ë") // LATIN SMALL LETTER E WITH DIAERESIS
    converter.latinToUtf8[236] = []byte("ì") // LATIN SMALL LETTER I WITH GRAVE
    converter.latinToUtf8[237] = []byte("í") // LATIN SMALL LETTER I WITH ACUTE
    converter.latinToUtf8[238] = []byte("î") // LATIN SMALL LETTER I WITH CIRCUMFLEX
    converter.latinToUtf8[239] = []byte("ï") // LATIN SMALL LETTER I WITH DIAERESIS
    converter.latinToUtf8[241] = []byte("ñ") // LATIN SMALL LETTER N WITH TILDE
    converter.latinToUtf8[242] = []byte("ò") // LATIN SMALL LETTER O WITH GRAVE
    converter.latinToUtf8[243] = []byte("ó") // LATIN SMALL LETTER O WITH ACUTE
    converter.latinToUtf8[244] = []byte("ô") // LATIN SMALL LETTER O WITH CIRCUMFLEX
    converter.latinToUtf8[245] = []byte("ġ") // LATIN SMALL LETTER G WITH DOT ABOVE
    converter.latinToUtf8[246] = []byte("ö") // LATIN SMALL LETTER O WITH DIAERESIS
    converter.latinToUtf8[247] = []byte("÷") // DIVISION SIGN
    converter.latinToUtf8[248] = []byte("ĝ") // LATIN SMALL LETTER G WITH CIRCUMFLEX
    converter.latinToUtf8[249] = []byte("ù") // LATIN SMALL LETTER U WITH GRAVE
    converter.latinToUtf8[250] = []byte("ú") // LATIN SMALL LETTER U WITH ACUTE
    converter.latinToUtf8[251] = []byte("û") // LATIN SMALL LETTER U WITH CIRCUMFLEX
    converter.latinToUtf8[252] = []byte("ü") // LATIN SMALL LETTER U WITH DIAERESIS
    converter.latinToUtf8[253] = []byte("ŭ") // LATIN SMALL LETTER U WITH BREVE
    converter.latinToUtf8[254] = []byte("ŝ") // LATIN SMALL LETTER S WITH CIRCUMFLEX
    converter.latinToUtf8[255] = []byte("˙") // DOT ABOVE


    converter.utf8ToLatin = make(map[int]byte)

    converter.utf8ToLatin[160] = 160 // [ ]  U+00A0 => A0
    converter.utf8ToLatin[294] = 161 // [Ħ]  U+0126 => A1
    converter.utf8ToLatin[728] = 162 // [˘]  U+02D8 => A2
    converter.utf8ToLatin[163] = 163 // [£]  U+00A3 => A3
    converter.utf8ToLatin[164] = 164 // [¤]  U+00A4 => A4
    converter.utf8ToLatin[292] = 166 // [Ĥ]  U+0124 => A6
    converter.utf8ToLatin[167] = 167 // [§]  U+00A7 => A7
    converter.utf8ToLatin[168] = 168 // [¨]  U+00A8 => A8
    converter.utf8ToLatin[304] = 169 // [İ]  U+0130 => A9
    converter.utf8ToLatin[350] = 170 // [Ş]  U+015E => AA
    converter.utf8ToLatin[286] = 171 // [Ğ]  U+011E => AB
    converter.utf8ToLatin[308] = 172 // [Ĵ]  U+0134 => AC
    converter.utf8ToLatin[173] = 173 // [­]  U+00AD => AD
    converter.utf8ToLatin[379] = 175 // [Ż]  U+017B => AF
    converter.utf8ToLatin[176] = 176 // [°]  U+00B0 => B0
    converter.utf8ToLatin[295] = 177 // [ħ]  U+0127 => B1
    converter.utf8ToLatin[178] = 178 // [²]  U+00B2 => B2
    converter.utf8ToLatin[179] = 179 // [³]  U+00B3 => B3
    converter.utf8ToLatin[180] = 180 // [´]  U+00B4 => B4
    converter.utf8ToLatin[181] = 181 // [µ]  U+00B5 => B5
    converter.utf8ToLatin[293] = 182 // [ĥ]  U+0125 => B6
    converter.utf8ToLatin[183] = 183 // [·]  U+00B7 => B7
    converter.utf8ToLatin[184] = 184 // [¸]  U+00B8 => B8
    converter.utf8ToLatin[305] = 185 // [ı]  U+0131 => B9
    converter.utf8ToLatin[351] = 186 // [ş]  U+015F => BA
    converter.utf8ToLatin[287] = 187 // [ğ]  U+011F => BB
    converter.utf8ToLatin[309] = 188 // [ĵ]  U+0135 => BC
    converter.utf8ToLatin[189] = 189 // [½]  U+00BD => BD
    converter.utf8ToLatin[380] = 191 // [ż]  U+017C => BF
    converter.utf8ToLatin[192] = 192 // [À]  U+00C0 => C0
    converter.utf8ToLatin[193] = 193 // [Á]  U+00C1 => C1
    converter.utf8ToLatin[194] = 194 // [Â]  U+00C2 => C2
    converter.utf8ToLatin[196] = 196 // [Ä]  U+00C4 => C4
    converter.utf8ToLatin[266] = 197 // [Ċ]  U+010A => C5
    converter.utf8ToLatin[264] = 198 // [Ĉ]  U+0108 => C6
    converter.utf8ToLatin[199] = 199 // [Ç]  U+00C7 => C7
    converter.utf8ToLatin[200] = 200 // [È]  U+00C8 => C8
    converter.utf8ToLatin[201] = 201 // [É]  U+00C9 => C9
    converter.utf8ToLatin[202] = 202 // [Ê]  U+00CA => CA
    converter.utf8ToLatin[203] = 203 // [Ë]  U+00CB => CB
    converter.utf8ToLatin[204] = 204 // [Ì]  U+00CC => CC
    converter.utf8ToLatin[205] = 205 // [Í]  U+00CD => CD
    converter.utf8ToLatin[206] = 206 // [Î]  U+00CE => CE
    converter.utf8ToLatin[207] = 207 // [Ï]  U+00CF => CF
    converter.utf8ToLatin[209] = 209 // [Ñ]  U+00D1 => D1
    converter.utf8ToLatin[210] = 210 // [Ò]  U+00D2 => D2
    converter.utf8ToLatin[211] = 211 // [Ó]  U+00D3 => D3
    converter.utf8ToLatin[212] = 212 // [Ô]  U+00D4 => D4
    converter.utf8ToLatin[288] = 213 // [Ġ]  U+0120 => D5
    converter.utf8ToLatin[214] = 214 // [Ö]  U+00D6 => D6
    converter.utf8ToLatin[215] = 215 // [×]  U+00D7 => D7
    converter.utf8ToLatin[284] = 216 // [Ĝ]  U+011C => D8
    converter.utf8ToLatin[217] = 217 // [Ù]  U+00D9 => D9
    converter.utf8ToLatin[218] = 218 // [Ú]  U+00DA => DA
    converter.utf8ToLatin[219] = 219 // [Û]  U+00DB => DB
    converter.utf8ToLatin[220] = 220 // [Ü]  U+00DC => DC
    converter.utf8ToLatin[364] = 221 // [Ŭ]  U+016C => DD
    converter.utf8ToLatin[348] = 222 // [Ŝ]  U+015C => DE
    converter.utf8ToLatin[223] = 223 // [ß]  U+00DF => DF
    converter.utf8ToLatin[224] = 224 // [à]  U+00E0 => E0
    converter.utf8ToLatin[225] = 225 // [á]  U+00E1 => E1
    converter.utf8ToLatin[226] = 226 // [â]  U+00E2 => E2
    converter.utf8ToLatin[228] = 228 // [ä]  U+00E4 => E4
    converter.utf8ToLatin[267] = 229 // [ċ]  U+010B => E5
    converter.utf8ToLatin[265] = 230 // [ĉ]  U+0109 => E6
    converter.utf8ToLatin[231] = 231 // [ç]  U+00E7 => E7
    converter.utf8ToLatin[232] = 232 // [è]  U+00E8 => E8
    converter.utf8ToLatin[233] = 233 // [é]  U+00E9 => E9
    converter.utf8ToLatin[234] = 234 // [ê]  U+00EA => EA
    converter.utf8ToLatin[235] = 235 // [ë]  U+00EB => EB
    converter.utf8ToLatin[236] = 236 // [ì]  U+00EC => EC
    converter.utf8ToLatin[237] = 237 // [í]  U+00ED => ED
    converter.utf8ToLatin[238] = 238 // [î]  U+00EE => EE
    converter.utf8ToLatin[239] = 239 // [ï]  U+00EF => EF
    converter.utf8ToLatin[241] = 241 // [ñ]  U+00F1 => F1
    converter.utf8ToLatin[242] = 242 // [ò]  U+00F2 => F2
    converter.utf8ToLatin[243] = 243 // [ó]  U+00F3 => F3
    converter.utf8ToLatin[244] = 244 // [ô]  U+00F4 => F4
    converter.utf8ToLatin[289] = 245 // [ġ]  U+0121 => F5
    converter.utf8ToLatin[246] = 246 // [ö]  U+00F6 => F6
    converter.utf8ToLatin[247] = 247 // [÷]  U+00F7 => F7
    converter.utf8ToLatin[285] = 248 // [ĝ]  U+011D => F8
    converter.utf8ToLatin[249] = 249 // [ù]  U+00F9 => F9
    converter.utf8ToLatin[250] = 250 // [ú]  U+00FA => FA
    converter.utf8ToLatin[251] = 251 // [û]  U+00FB => FB
    converter.utf8ToLatin[252] = 252 // [ü]  U+00FC => FC
    converter.utf8ToLatin[365] = 253 // [ŭ]  U+016D => FD
    converter.utf8ToLatin[349] = 254 // [ŝ]  U+015D => FE
    converter.utf8ToLatin[729] = 255 // [˙]  U+02D9 => FF

    return converter
}
