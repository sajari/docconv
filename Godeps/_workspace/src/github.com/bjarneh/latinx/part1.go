// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

// ISO 8859-1

// Generated from http://unicode.org/Public/MAPPINGS/ISO8859/8859-1.TXT

func newISO_8859_1Converter() *Converter {

    converter := new(Converter)
    converter.id = "Latin-1/ISO-8859-1"
    converter.latinToUtf8 = make(map[byte][]byte)

    converter.latinToUtf8[160] = []byte(" ") // NO-BREAK SPACE
    converter.latinToUtf8[161] = []byte("¡") // INVERTED EXCLAMATION MARK
    converter.latinToUtf8[162] = []byte("¢") // CENT SIGN
    converter.latinToUtf8[163] = []byte("£") // POUND SIGN
    converter.latinToUtf8[164] = []byte("¤") // CURRENCY SIGN
    converter.latinToUtf8[165] = []byte("¥") // YEN SIGN
    converter.latinToUtf8[166] = []byte("¦") // BROKEN BAR
    converter.latinToUtf8[167] = []byte("§") // SECTION SIGN
    converter.latinToUtf8[168] = []byte("¨") // DIAERESIS
    converter.latinToUtf8[169] = []byte("©") // COPYRIGHT SIGN
    converter.latinToUtf8[170] = []byte("ª") // FEMININE ORDINAL INDICATOR
    converter.latinToUtf8[171] = []byte("«") // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
    converter.latinToUtf8[172] = []byte("¬") // NOT SIGN
    converter.latinToUtf8[173] = []byte("­") // SOFT HYPHEN
    converter.latinToUtf8[174] = []byte("®") // REGISTERED SIGN
    converter.latinToUtf8[175] = []byte("¯") // MACRON
    converter.latinToUtf8[176] = []byte("°") // DEGREE SIGN
    converter.latinToUtf8[177] = []byte("±") // PLUS-MINUS SIGN
    converter.latinToUtf8[178] = []byte("²") // SUPERSCRIPT TWO
    converter.latinToUtf8[179] = []byte("³") // SUPERSCRIPT THREE
    converter.latinToUtf8[180] = []byte("´") // ACUTE ACCENT
    converter.latinToUtf8[181] = []byte("µ") // MICRO SIGN
    converter.latinToUtf8[182] = []byte("¶") // PILCROW SIGN
    converter.latinToUtf8[183] = []byte("·") // MIDDLE DOT
    converter.latinToUtf8[184] = []byte("¸") // CEDILLA
    converter.latinToUtf8[185] = []byte("¹") // SUPERSCRIPT ONE
    converter.latinToUtf8[186] = []byte("º") // MASCULINE ORDINAL INDICATOR
    converter.latinToUtf8[187] = []byte("»") // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
    converter.latinToUtf8[188] = []byte("¼") // VULGAR FRACTION ONE QUARTER
    converter.latinToUtf8[189] = []byte("½") // VULGAR FRACTION ONE HALF
    converter.latinToUtf8[190] = []byte("¾") // VULGAR FRACTION THREE QUARTERS
    converter.latinToUtf8[191] = []byte("¿") // INVERTED QUESTION MARK
    converter.latinToUtf8[192] = []byte("À") // LATIN CAPITAL LETTER A WITH GRAVE
    converter.latinToUtf8[193] = []byte("Á") // LATIN CAPITAL LETTER A WITH ACUTE
    converter.latinToUtf8[194] = []byte("Â") // LATIN CAPITAL LETTER A WITH CIRCUMFLEX
    converter.latinToUtf8[195] = []byte("Ã") // LATIN CAPITAL LETTER A WITH TILDE
    converter.latinToUtf8[196] = []byte("Ä") // LATIN CAPITAL LETTER A WITH DIAERESIS
    converter.latinToUtf8[197] = []byte("Å") // LATIN CAPITAL LETTER A WITH RING ABOVE
    converter.latinToUtf8[198] = []byte("Æ") // LATIN CAPITAL LETTER AE
    converter.latinToUtf8[199] = []byte("Ç") // LATIN CAPITAL LETTER C WITH CEDILLA
    converter.latinToUtf8[200] = []byte("È") // LATIN CAPITAL LETTER E WITH GRAVE
    converter.latinToUtf8[201] = []byte("É") // LATIN CAPITAL LETTER E WITH ACUTE
    converter.latinToUtf8[202] = []byte("Ê") // LATIN CAPITAL LETTER E WITH CIRCUMFLEX
    converter.latinToUtf8[203] = []byte("Ë") // LATIN CAPITAL LETTER E WITH DIAERESIS
    converter.latinToUtf8[204] = []byte("Ì") // LATIN CAPITAL LETTER I WITH GRAVE
    converter.latinToUtf8[205] = []byte("Í") // LATIN CAPITAL LETTER I WITH ACUTE
    converter.latinToUtf8[206] = []byte("Î") // LATIN CAPITAL LETTER I WITH CIRCUMFLEX
    converter.latinToUtf8[207] = []byte("Ï") // LATIN CAPITAL LETTER I WITH DIAERESIS
    converter.latinToUtf8[208] = []byte("Ð") // LATIN CAPITAL LETTER ETH (Icelandic)
    converter.latinToUtf8[209] = []byte("Ñ") // LATIN CAPITAL LETTER N WITH TILDE
    converter.latinToUtf8[210] = []byte("Ò") // LATIN CAPITAL LETTER O WITH GRAVE
    converter.latinToUtf8[211] = []byte("Ó") // LATIN CAPITAL LETTER O WITH ACUTE
    converter.latinToUtf8[212] = []byte("Ô") // LATIN CAPITAL LETTER O WITH CIRCUMFLEX
    converter.latinToUtf8[213] = []byte("Õ") // LATIN CAPITAL LETTER O WITH TILDE
    converter.latinToUtf8[214] = []byte("Ö") // LATIN CAPITAL LETTER O WITH DIAERESIS
    converter.latinToUtf8[215] = []byte("×") // MULTIPLICATION SIGN
    converter.latinToUtf8[216] = []byte("Ø") // LATIN CAPITAL LETTER O WITH STROKE
    converter.latinToUtf8[217] = []byte("Ù") // LATIN CAPITAL LETTER U WITH GRAVE
    converter.latinToUtf8[218] = []byte("Ú") // LATIN CAPITAL LETTER U WITH ACUTE
    converter.latinToUtf8[219] = []byte("Û") // LATIN CAPITAL LETTER U WITH CIRCUMFLEX
    converter.latinToUtf8[220] = []byte("Ü") // LATIN CAPITAL LETTER U WITH DIAERESIS
    converter.latinToUtf8[221] = []byte("Ý") // LATIN CAPITAL LETTER Y WITH ACUTE
    converter.latinToUtf8[222] = []byte("Þ") // LATIN CAPITAL LETTER THORN (Icelandic)
    converter.latinToUtf8[223] = []byte("ß") // LATIN SMALL LETTER SHARP S (German)
    converter.latinToUtf8[224] = []byte("à") // LATIN SMALL LETTER A WITH GRAVE
    converter.latinToUtf8[225] = []byte("á") // LATIN SMALL LETTER A WITH ACUTE
    converter.latinToUtf8[226] = []byte("â") // LATIN SMALL LETTER A WITH CIRCUMFLEX
    converter.latinToUtf8[227] = []byte("ã") // LATIN SMALL LETTER A WITH TILDE
    converter.latinToUtf8[228] = []byte("ä") // LATIN SMALL LETTER A WITH DIAERESIS
    converter.latinToUtf8[229] = []byte("å") // LATIN SMALL LETTER A WITH RING ABOVE
    converter.latinToUtf8[230] = []byte("æ") // LATIN SMALL LETTER AE
    converter.latinToUtf8[231] = []byte("ç") // LATIN SMALL LETTER C WITH CEDILLA
    converter.latinToUtf8[232] = []byte("è") // LATIN SMALL LETTER E WITH GRAVE
    converter.latinToUtf8[233] = []byte("é") // LATIN SMALL LETTER E WITH ACUTE
    converter.latinToUtf8[234] = []byte("ê") // LATIN SMALL LETTER E WITH CIRCUMFLEX
    converter.latinToUtf8[235] = []byte("ë") // LATIN SMALL LETTER E WITH DIAERESIS
    converter.latinToUtf8[236] = []byte("ì") // LATIN SMALL LETTER I WITH GRAVE
    converter.latinToUtf8[237] = []byte("í") // LATIN SMALL LETTER I WITH ACUTE
    converter.latinToUtf8[238] = []byte("î") // LATIN SMALL LETTER I WITH CIRCUMFLEX
    converter.latinToUtf8[239] = []byte("ï") // LATIN SMALL LETTER I WITH DIAERESIS
    converter.latinToUtf8[240] = []byte("ð") // LATIN SMALL LETTER ETH (Icelandic)
    converter.latinToUtf8[241] = []byte("ñ") // LATIN SMALL LETTER N WITH TILDE
    converter.latinToUtf8[242] = []byte("ò") // LATIN SMALL LETTER O WITH GRAVE
    converter.latinToUtf8[243] = []byte("ó") // LATIN SMALL LETTER O WITH ACUTE
    converter.latinToUtf8[244] = []byte("ô") // LATIN SMALL LETTER O WITH CIRCUMFLEX
    converter.latinToUtf8[245] = []byte("õ") // LATIN SMALL LETTER O WITH TILDE
    converter.latinToUtf8[246] = []byte("ö") // LATIN SMALL LETTER O WITH DIAERESIS
    converter.latinToUtf8[247] = []byte("÷") // DIVISION SIGN
    converter.latinToUtf8[248] = []byte("ø") // LATIN SMALL LETTER O WITH STROKE
    converter.latinToUtf8[249] = []byte("ù") // LATIN SMALL LETTER U WITH GRAVE
    converter.latinToUtf8[250] = []byte("ú") // LATIN SMALL LETTER U WITH ACUTE
    converter.latinToUtf8[251] = []byte("û") // LATIN SMALL LETTER U WITH CIRCUMFLEX
    converter.latinToUtf8[252] = []byte("ü") // LATIN SMALL LETTER U WITH DIAERESIS
    converter.latinToUtf8[253] = []byte("ý") // LATIN SMALL LETTER Y WITH ACUTE
    converter.latinToUtf8[254] = []byte("þ") // LATIN SMALL LETTER THORN (Icelandic)
    converter.latinToUtf8[255] = []byte("ÿ") // LATIN SMALL LETTER Y WITH DIAERESIS


    converter.utf8ToLatin = make(map[int]byte)

    converter.utf8ToLatin[160] = 160 // [ ]  U+00A0 => A0
    converter.utf8ToLatin[161] = 161 // [¡]  U+00A1 => A1
    converter.utf8ToLatin[162] = 162 // [¢]  U+00A2 => A2
    converter.utf8ToLatin[163] = 163 // [£]  U+00A3 => A3
    converter.utf8ToLatin[164] = 164 // [¤]  U+00A4 => A4
    converter.utf8ToLatin[165] = 165 // [¥]  U+00A5 => A5
    converter.utf8ToLatin[166] = 166 // [¦]  U+00A6 => A6
    converter.utf8ToLatin[167] = 167 // [§]  U+00A7 => A7
    converter.utf8ToLatin[168] = 168 // [¨]  U+00A8 => A8
    converter.utf8ToLatin[169] = 169 // [©]  U+00A9 => A9
    converter.utf8ToLatin[170] = 170 // [ª]  U+00AA => AA
    converter.utf8ToLatin[171] = 171 // [«]  U+00AB => AB
    converter.utf8ToLatin[172] = 172 // [¬]  U+00AC => AC
    converter.utf8ToLatin[173] = 173 // [­]  U+00AD => AD
    converter.utf8ToLatin[174] = 174 // [®]  U+00AE => AE
    converter.utf8ToLatin[175] = 175 // [¯]  U+00AF => AF
    converter.utf8ToLatin[176] = 176 // [°]  U+00B0 => B0
    converter.utf8ToLatin[177] = 177 // [±]  U+00B1 => B1
    converter.utf8ToLatin[178] = 178 // [²]  U+00B2 => B2
    converter.utf8ToLatin[179] = 179 // [³]  U+00B3 => B3
    converter.utf8ToLatin[180] = 180 // [´]  U+00B4 => B4
    converter.utf8ToLatin[181] = 181 // [µ]  U+00B5 => B5
    converter.utf8ToLatin[182] = 182 // [¶]  U+00B6 => B6
    converter.utf8ToLatin[183] = 183 // [·]  U+00B7 => B7
    converter.utf8ToLatin[184] = 184 // [¸]  U+00B8 => B8
    converter.utf8ToLatin[185] = 185 // [¹]  U+00B9 => B9
    converter.utf8ToLatin[186] = 186 // [º]  U+00BA => BA
    converter.utf8ToLatin[187] = 187 // [»]  U+00BB => BB
    converter.utf8ToLatin[188] = 188 // [¼]  U+00BC => BC
    converter.utf8ToLatin[189] = 189 // [½]  U+00BD => BD
    converter.utf8ToLatin[190] = 190 // [¾]  U+00BE => BE
    converter.utf8ToLatin[191] = 191 // [¿]  U+00BF => BF
    converter.utf8ToLatin[192] = 192 // [À]  U+00C0 => C0
    converter.utf8ToLatin[193] = 193 // [Á]  U+00C1 => C1
    converter.utf8ToLatin[194] = 194 // [Â]  U+00C2 => C2
    converter.utf8ToLatin[195] = 195 // [Ã]  U+00C3 => C3
    converter.utf8ToLatin[196] = 196 // [Ä]  U+00C4 => C4
    converter.utf8ToLatin[197] = 197 // [Å]  U+00C5 => C5
    converter.utf8ToLatin[198] = 198 // [Æ]  U+00C6 => C6
    converter.utf8ToLatin[199] = 199 // [Ç]  U+00C7 => C7
    converter.utf8ToLatin[200] = 200 // [È]  U+00C8 => C8
    converter.utf8ToLatin[201] = 201 // [É]  U+00C9 => C9
    converter.utf8ToLatin[202] = 202 // [Ê]  U+00CA => CA
    converter.utf8ToLatin[203] = 203 // [Ë]  U+00CB => CB
    converter.utf8ToLatin[204] = 204 // [Ì]  U+00CC => CC
    converter.utf8ToLatin[205] = 205 // [Í]  U+00CD => CD
    converter.utf8ToLatin[206] = 206 // [Î]  U+00CE => CE
    converter.utf8ToLatin[207] = 207 // [Ï]  U+00CF => CF
    converter.utf8ToLatin[208] = 208 // [Ð]  U+00D0 => D0
    converter.utf8ToLatin[209] = 209 // [Ñ]  U+00D1 => D1
    converter.utf8ToLatin[210] = 210 // [Ò]  U+00D2 => D2
    converter.utf8ToLatin[211] = 211 // [Ó]  U+00D3 => D3
    converter.utf8ToLatin[212] = 212 // [Ô]  U+00D4 => D4
    converter.utf8ToLatin[213] = 213 // [Õ]  U+00D5 => D5
    converter.utf8ToLatin[214] = 214 // [Ö]  U+00D6 => D6
    converter.utf8ToLatin[215] = 215 // [×]  U+00D7 => D7
    converter.utf8ToLatin[216] = 216 // [Ø]  U+00D8 => D8
    converter.utf8ToLatin[217] = 217 // [Ù]  U+00D9 => D9
    converter.utf8ToLatin[218] = 218 // [Ú]  U+00DA => DA
    converter.utf8ToLatin[219] = 219 // [Û]  U+00DB => DB
    converter.utf8ToLatin[220] = 220 // [Ü]  U+00DC => DC
    converter.utf8ToLatin[221] = 221 // [Ý]  U+00DD => DD
    converter.utf8ToLatin[222] = 222 // [Þ]  U+00DE => DE
    converter.utf8ToLatin[223] = 223 // [ß]  U+00DF => DF
    converter.utf8ToLatin[224] = 224 // [à]  U+00E0 => E0
    converter.utf8ToLatin[225] = 225 // [á]  U+00E1 => E1
    converter.utf8ToLatin[226] = 226 // [â]  U+00E2 => E2
    converter.utf8ToLatin[227] = 227 // [ã]  U+00E3 => E3
    converter.utf8ToLatin[228] = 228 // [ä]  U+00E4 => E4
    converter.utf8ToLatin[229] = 229 // [å]  U+00E5 => E5
    converter.utf8ToLatin[230] = 230 // [æ]  U+00E6 => E6
    converter.utf8ToLatin[231] = 231 // [ç]  U+00E7 => E7
    converter.utf8ToLatin[232] = 232 // [è]  U+00E8 => E8
    converter.utf8ToLatin[233] = 233 // [é]  U+00E9 => E9
    converter.utf8ToLatin[234] = 234 // [ê]  U+00EA => EA
    converter.utf8ToLatin[235] = 235 // [ë]  U+00EB => EB
    converter.utf8ToLatin[236] = 236 // [ì]  U+00EC => EC
    converter.utf8ToLatin[237] = 237 // [í]  U+00ED => ED
    converter.utf8ToLatin[238] = 238 // [î]  U+00EE => EE
    converter.utf8ToLatin[239] = 239 // [ï]  U+00EF => EF
    converter.utf8ToLatin[240] = 240 // [ð]  U+00F0 => F0
    converter.utf8ToLatin[241] = 241 // [ñ]  U+00F1 => F1
    converter.utf8ToLatin[242] = 242 // [ò]  U+00F2 => F2
    converter.utf8ToLatin[243] = 243 // [ó]  U+00F3 => F3
    converter.utf8ToLatin[244] = 244 // [ô]  U+00F4 => F4
    converter.utf8ToLatin[245] = 245 // [õ]  U+00F5 => F5
    converter.utf8ToLatin[246] = 246 // [ö]  U+00F6 => F6
    converter.utf8ToLatin[247] = 247 // [÷]  U+00F7 => F7
    converter.utf8ToLatin[248] = 248 // [ø]  U+00F8 => F8
    converter.utf8ToLatin[249] = 249 // [ù]  U+00F9 => F9
    converter.utf8ToLatin[250] = 250 // [ú]  U+00FA => FA
    converter.utf8ToLatin[251] = 251 // [û]  U+00FB => FB
    converter.utf8ToLatin[252] = 252 // [ü]  U+00FC => FC
    converter.utf8ToLatin[253] = 253 // [ý]  U+00FD => FD
    converter.utf8ToLatin[254] = 254 // [þ]  U+00FE => FE
    converter.utf8ToLatin[255] = 255 // [ÿ]  U+00FF => FF

    return converter
}
