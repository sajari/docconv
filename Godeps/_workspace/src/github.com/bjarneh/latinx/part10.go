// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

// ISO 8859-10

// Generated from http://unicode.org/Public/MAPPINGS/ISO8859/8859-10.TXT

func newISO_8859_10Converter() *Converter {

    converter := new(Converter)
    converter.id = "Latin-6/ISO-8859-10"
    converter.latinToUtf8 = make(map[byte][]byte)

    converter.latinToUtf8[160] = []byte(" ") // NO-BREAK SPACE
    converter.latinToUtf8[161] = []byte("Ą") // LATIN CAPITAL LETTER A WITH OGONEK
    converter.latinToUtf8[162] = []byte("Ē") // LATIN CAPITAL LETTER E WITH MACRON
    converter.latinToUtf8[163] = []byte("Ģ") // LATIN CAPITAL LETTER G WITH CEDILLA
    converter.latinToUtf8[164] = []byte("Ī") // LATIN CAPITAL LETTER I WITH MACRON
    converter.latinToUtf8[165] = []byte("Ĩ") // LATIN CAPITAL LETTER I WITH TILDE
    converter.latinToUtf8[166] = []byte("Ķ") // LATIN CAPITAL LETTER K WITH CEDILLA
    converter.latinToUtf8[167] = []byte("§") // SECTION SIGN
    converter.latinToUtf8[168] = []byte("Ļ") // LATIN CAPITAL LETTER L WITH CEDILLA
    converter.latinToUtf8[169] = []byte("Đ") // LATIN CAPITAL LETTER D WITH STROKE
    converter.latinToUtf8[170] = []byte("Š") // LATIN CAPITAL LETTER S WITH CARON
    converter.latinToUtf8[171] = []byte("Ŧ") // LATIN CAPITAL LETTER T WITH STROKE
    converter.latinToUtf8[172] = []byte("Ž") // LATIN CAPITAL LETTER Z WITH CARON
    converter.latinToUtf8[173] = []byte("­") // SOFT HYPHEN
    converter.latinToUtf8[174] = []byte("Ū") // LATIN CAPITAL LETTER U WITH MACRON
    converter.latinToUtf8[175] = []byte("Ŋ") // LATIN CAPITAL LETTER ENG
    converter.latinToUtf8[176] = []byte("°") // DEGREE SIGN
    converter.latinToUtf8[177] = []byte("ą") // LATIN SMALL LETTER A WITH OGONEK
    converter.latinToUtf8[178] = []byte("ē") // LATIN SMALL LETTER E WITH MACRON
    converter.latinToUtf8[179] = []byte("ģ") // LATIN SMALL LETTER G WITH CEDILLA
    converter.latinToUtf8[180] = []byte("ī") // LATIN SMALL LETTER I WITH MACRON
    converter.latinToUtf8[181] = []byte("ĩ") // LATIN SMALL LETTER I WITH TILDE
    converter.latinToUtf8[182] = []byte("ķ") // LATIN SMALL LETTER K WITH CEDILLA
    converter.latinToUtf8[183] = []byte("·") // MIDDLE DOT
    converter.latinToUtf8[184] = []byte("ļ") // LATIN SMALL LETTER L WITH CEDILLA
    converter.latinToUtf8[185] = []byte("đ") // LATIN SMALL LETTER D WITH STROKE
    converter.latinToUtf8[186] = []byte("š") // LATIN SMALL LETTER S WITH CARON
    converter.latinToUtf8[187] = []byte("ŧ") // LATIN SMALL LETTER T WITH STROKE
    converter.latinToUtf8[188] = []byte("ž") // LATIN SMALL LETTER Z WITH CARON
    converter.latinToUtf8[189] = []byte("―") // HORIZONTAL BAR
    converter.latinToUtf8[190] = []byte("ū") // LATIN SMALL LETTER U WITH MACRON
    converter.latinToUtf8[191] = []byte("ŋ") // LATIN SMALL LETTER ENG
    converter.latinToUtf8[192] = []byte("Ā") // LATIN CAPITAL LETTER A WITH MACRON
    converter.latinToUtf8[193] = []byte("Á") // LATIN CAPITAL LETTER A WITH ACUTE
    converter.latinToUtf8[194] = []byte("Â") // LATIN CAPITAL LETTER A WITH CIRCUMFLEX
    converter.latinToUtf8[195] = []byte("Ã") // LATIN CAPITAL LETTER A WITH TILDE
    converter.latinToUtf8[196] = []byte("Ä") // LATIN CAPITAL LETTER A WITH DIAERESIS
    converter.latinToUtf8[197] = []byte("Å") // LATIN CAPITAL LETTER A WITH RING ABOVE
    converter.latinToUtf8[198] = []byte("Æ") // LATIN CAPITAL LETTER AE
    converter.latinToUtf8[199] = []byte("Į") // LATIN CAPITAL LETTER I WITH OGONEK
    converter.latinToUtf8[200] = []byte("Č") // LATIN CAPITAL LETTER C WITH CARON
    converter.latinToUtf8[201] = []byte("É") // LATIN CAPITAL LETTER E WITH ACUTE
    converter.latinToUtf8[202] = []byte("Ę") // LATIN CAPITAL LETTER E WITH OGONEK
    converter.latinToUtf8[203] = []byte("Ë") // LATIN CAPITAL LETTER E WITH DIAERESIS
    converter.latinToUtf8[204] = []byte("Ė") // LATIN CAPITAL LETTER E WITH DOT ABOVE
    converter.latinToUtf8[205] = []byte("Í") // LATIN CAPITAL LETTER I WITH ACUTE
    converter.latinToUtf8[206] = []byte("Î") // LATIN CAPITAL LETTER I WITH CIRCUMFLEX
    converter.latinToUtf8[207] = []byte("Ï") // LATIN CAPITAL LETTER I WITH DIAERESIS
    converter.latinToUtf8[208] = []byte("Ð") // LATIN CAPITAL LETTER ETH (Icelandic)
    converter.latinToUtf8[209] = []byte("Ņ") // LATIN CAPITAL LETTER N WITH CEDILLA
    converter.latinToUtf8[210] = []byte("Ō") // LATIN CAPITAL LETTER O WITH MACRON
    converter.latinToUtf8[211] = []byte("Ó") // LATIN CAPITAL LETTER O WITH ACUTE
    converter.latinToUtf8[212] = []byte("Ô") // LATIN CAPITAL LETTER O WITH CIRCUMFLEX
    converter.latinToUtf8[213] = []byte("Õ") // LATIN CAPITAL LETTER O WITH TILDE
    converter.latinToUtf8[214] = []byte("Ö") // LATIN CAPITAL LETTER O WITH DIAERESIS
    converter.latinToUtf8[215] = []byte("Ũ") // LATIN CAPITAL LETTER U WITH TILDE
    converter.latinToUtf8[216] = []byte("Ø") // LATIN CAPITAL LETTER O WITH STROKE
    converter.latinToUtf8[217] = []byte("Ų") // LATIN CAPITAL LETTER U WITH OGONEK
    converter.latinToUtf8[218] = []byte("Ú") // LATIN CAPITAL LETTER U WITH ACUTE
    converter.latinToUtf8[219] = []byte("Û") // LATIN CAPITAL LETTER U WITH CIRCUMFLEX
    converter.latinToUtf8[220] = []byte("Ü") // LATIN CAPITAL LETTER U WITH DIAERESIS
    converter.latinToUtf8[221] = []byte("Ý") // LATIN CAPITAL LETTER Y WITH ACUTE
    converter.latinToUtf8[222] = []byte("Þ") // LATIN CAPITAL LETTER THORN (Icelandic)
    converter.latinToUtf8[223] = []byte("ß") // LATIN SMALL LETTER SHARP S (German)
    converter.latinToUtf8[224] = []byte("ā") // LATIN SMALL LETTER A WITH MACRON
    converter.latinToUtf8[225] = []byte("á") // LATIN SMALL LETTER A WITH ACUTE
    converter.latinToUtf8[226] = []byte("â") // LATIN SMALL LETTER A WITH CIRCUMFLEX
    converter.latinToUtf8[227] = []byte("ã") // LATIN SMALL LETTER A WITH TILDE
    converter.latinToUtf8[228] = []byte("ä") // LATIN SMALL LETTER A WITH DIAERESIS
    converter.latinToUtf8[229] = []byte("å") // LATIN SMALL LETTER A WITH RING ABOVE
    converter.latinToUtf8[230] = []byte("æ") // LATIN SMALL LETTER AE
    converter.latinToUtf8[231] = []byte("į") // LATIN SMALL LETTER I WITH OGONEK
    converter.latinToUtf8[232] = []byte("č") // LATIN SMALL LETTER C WITH CARON
    converter.latinToUtf8[233] = []byte("é") // LATIN SMALL LETTER E WITH ACUTE
    converter.latinToUtf8[234] = []byte("ę") // LATIN SMALL LETTER E WITH OGONEK
    converter.latinToUtf8[235] = []byte("ë") // LATIN SMALL LETTER E WITH DIAERESIS
    converter.latinToUtf8[236] = []byte("ė") // LATIN SMALL LETTER E WITH DOT ABOVE
    converter.latinToUtf8[237] = []byte("í") // LATIN SMALL LETTER I WITH ACUTE
    converter.latinToUtf8[238] = []byte("î") // LATIN SMALL LETTER I WITH CIRCUMFLEX
    converter.latinToUtf8[239] = []byte("ï") // LATIN SMALL LETTER I WITH DIAERESIS
    converter.latinToUtf8[240] = []byte("ð") // LATIN SMALL LETTER ETH (Icelandic)
    converter.latinToUtf8[241] = []byte("ņ") // LATIN SMALL LETTER N WITH CEDILLA
    converter.latinToUtf8[242] = []byte("ō") // LATIN SMALL LETTER O WITH MACRON
    converter.latinToUtf8[243] = []byte("ó") // LATIN SMALL LETTER O WITH ACUTE
    converter.latinToUtf8[244] = []byte("ô") // LATIN SMALL LETTER O WITH CIRCUMFLEX
    converter.latinToUtf8[245] = []byte("õ") // LATIN SMALL LETTER O WITH TILDE
    converter.latinToUtf8[246] = []byte("ö") // LATIN SMALL LETTER O WITH DIAERESIS
    converter.latinToUtf8[247] = []byte("ũ") // LATIN SMALL LETTER U WITH TILDE
    converter.latinToUtf8[248] = []byte("ø") // LATIN SMALL LETTER O WITH STROKE
    converter.latinToUtf8[249] = []byte("ų") // LATIN SMALL LETTER U WITH OGONEK
    converter.latinToUtf8[250] = []byte("ú") // LATIN SMALL LETTER U WITH ACUTE
    converter.latinToUtf8[251] = []byte("û") // LATIN SMALL LETTER U WITH CIRCUMFLEX
    converter.latinToUtf8[252] = []byte("ü") // LATIN SMALL LETTER U WITH DIAERESIS
    converter.latinToUtf8[253] = []byte("ý") // LATIN SMALL LETTER Y WITH ACUTE
    converter.latinToUtf8[254] = []byte("þ") // LATIN SMALL LETTER THORN (Icelandic)
    converter.latinToUtf8[255] = []byte("ĸ") // LATIN SMALL LETTER KRA

    converter.utf8ToLatin = make(map[int]byte)

    converter.utf8ToLatin[160] = 160  // [ ]  U+00A0 => A0
    converter.utf8ToLatin[260] = 161  // [Ą]  U+0104 => A1
    converter.utf8ToLatin[274] = 162  // [Ē]  U+0112 => A2
    converter.utf8ToLatin[290] = 163  // [Ģ]  U+0122 => A3
    converter.utf8ToLatin[298] = 164  // [Ī]  U+012A => A4
    converter.utf8ToLatin[296] = 165  // [Ĩ]  U+0128 => A5
    converter.utf8ToLatin[310] = 166  // [Ķ]  U+0136 => A6
    converter.utf8ToLatin[167] = 167  // [§]  U+00A7 => A7
    converter.utf8ToLatin[315] = 168  // [Ļ]  U+013B => A8
    converter.utf8ToLatin[272] = 169  // [Đ]  U+0110 => A9
    converter.utf8ToLatin[352] = 170  // [Š]  U+0160 => AA
    converter.utf8ToLatin[358] = 171  // [Ŧ]  U+0166 => AB
    converter.utf8ToLatin[381] = 172  // [Ž]  U+017D => AC
    converter.utf8ToLatin[173] = 173  // [­]  U+00AD => AD
    converter.utf8ToLatin[362] = 174  // [Ū]  U+016A => AE
    converter.utf8ToLatin[330] = 175  // [Ŋ]  U+014A => AF
    converter.utf8ToLatin[176] = 176  // [°]  U+00B0 => B0
    converter.utf8ToLatin[261] = 177  // [ą]  U+0105 => B1
    converter.utf8ToLatin[275] = 178  // [ē]  U+0113 => B2
    converter.utf8ToLatin[291] = 179  // [ģ]  U+0123 => B3
    converter.utf8ToLatin[299] = 180  // [ī]  U+012B => B4
    converter.utf8ToLatin[297] = 181  // [ĩ]  U+0129 => B5
    converter.utf8ToLatin[311] = 182  // [ķ]  U+0137 => B6
    converter.utf8ToLatin[183] = 183  // [·]  U+00B7 => B7
    converter.utf8ToLatin[316] = 184  // [ļ]  U+013C => B8
    converter.utf8ToLatin[273] = 185  // [đ]  U+0111 => B9
    converter.utf8ToLatin[353] = 186  // [š]  U+0161 => BA
    converter.utf8ToLatin[359] = 187  // [ŧ]  U+0167 => BB
    converter.utf8ToLatin[382] = 188  // [ž]  U+017E => BC
    converter.utf8ToLatin[8213] = 189 // [―]  U+2015 => BD
    converter.utf8ToLatin[363] = 190  // [ū]  U+016B => BE
    converter.utf8ToLatin[331] = 191  // [ŋ]  U+014B => BF
    converter.utf8ToLatin[256] = 192  // [Ā]  U+0100 => C0
    converter.utf8ToLatin[193] = 193  // [Á]  U+00C1 => C1
    converter.utf8ToLatin[194] = 194  // [Â]  U+00C2 => C2
    converter.utf8ToLatin[195] = 195  // [Ã]  U+00C3 => C3
    converter.utf8ToLatin[196] = 196  // [Ä]  U+00C4 => C4
    converter.utf8ToLatin[197] = 197  // [Å]  U+00C5 => C5
    converter.utf8ToLatin[198] = 198  // [Æ]  U+00C6 => C6
    converter.utf8ToLatin[302] = 199  // [Į]  U+012E => C7
    converter.utf8ToLatin[268] = 200  // [Č]  U+010C => C8
    converter.utf8ToLatin[201] = 201  // [É]  U+00C9 => C9
    converter.utf8ToLatin[280] = 202  // [Ę]  U+0118 => CA
    converter.utf8ToLatin[203] = 203  // [Ë]  U+00CB => CB
    converter.utf8ToLatin[278] = 204  // [Ė]  U+0116 => CC
    converter.utf8ToLatin[205] = 205  // [Í]  U+00CD => CD
    converter.utf8ToLatin[206] = 206  // [Î]  U+00CE => CE
    converter.utf8ToLatin[207] = 207  // [Ï]  U+00CF => CF
    converter.utf8ToLatin[208] = 208  // [Ð]  U+00D0 => D0
    converter.utf8ToLatin[325] = 209  // [Ņ]  U+0145 => D1
    converter.utf8ToLatin[332] = 210  // [Ō]  U+014C => D2
    converter.utf8ToLatin[211] = 211  // [Ó]  U+00D3 => D3
    converter.utf8ToLatin[212] = 212  // [Ô]  U+00D4 => D4
    converter.utf8ToLatin[213] = 213  // [Õ]  U+00D5 => D5
    converter.utf8ToLatin[214] = 214  // [Ö]  U+00D6 => D6
    converter.utf8ToLatin[360] = 215  // [Ũ]  U+0168 => D7
    converter.utf8ToLatin[216] = 216  // [Ø]  U+00D8 => D8
    converter.utf8ToLatin[370] = 217  // [Ų]  U+0172 => D9
    converter.utf8ToLatin[218] = 218  // [Ú]  U+00DA => DA
    converter.utf8ToLatin[219] = 219  // [Û]  U+00DB => DB
    converter.utf8ToLatin[220] = 220  // [Ü]  U+00DC => DC
    converter.utf8ToLatin[221] = 221  // [Ý]  U+00DD => DD
    converter.utf8ToLatin[222] = 222  // [Þ]  U+00DE => DE
    converter.utf8ToLatin[223] = 223  // [ß]  U+00DF => DF
    converter.utf8ToLatin[257] = 224  // [ā]  U+0101 => E0
    converter.utf8ToLatin[225] = 225  // [á]  U+00E1 => E1
    converter.utf8ToLatin[226] = 226  // [â]  U+00E2 => E2
    converter.utf8ToLatin[227] = 227  // [ã]  U+00E3 => E3
    converter.utf8ToLatin[228] = 228  // [ä]  U+00E4 => E4
    converter.utf8ToLatin[229] = 229  // [å]  U+00E5 => E5
    converter.utf8ToLatin[230] = 230  // [æ]  U+00E6 => E6
    converter.utf8ToLatin[303] = 231  // [į]  U+012F => E7
    converter.utf8ToLatin[269] = 232  // [č]  U+010D => E8
    converter.utf8ToLatin[233] = 233  // [é]  U+00E9 => E9
    converter.utf8ToLatin[281] = 234  // [ę]  U+0119 => EA
    converter.utf8ToLatin[235] = 235  // [ë]  U+00EB => EB
    converter.utf8ToLatin[279] = 236  // [ė]  U+0117 => EC
    converter.utf8ToLatin[237] = 237  // [í]  U+00ED => ED
    converter.utf8ToLatin[238] = 238  // [î]  U+00EE => EE
    converter.utf8ToLatin[239] = 239  // [ï]  U+00EF => EF
    converter.utf8ToLatin[240] = 240  // [ð]  U+00F0 => F0
    converter.utf8ToLatin[326] = 241  // [ņ]  U+0146 => F1
    converter.utf8ToLatin[333] = 242  // [ō]  U+014D => F2
    converter.utf8ToLatin[243] = 243  // [ó]  U+00F3 => F3
    converter.utf8ToLatin[244] = 244  // [ô]  U+00F4 => F4
    converter.utf8ToLatin[245] = 245  // [õ]  U+00F5 => F5
    converter.utf8ToLatin[246] = 246  // [ö]  U+00F6 => F6
    converter.utf8ToLatin[361] = 247  // [ũ]  U+0169 => F7
    converter.utf8ToLatin[248] = 248  // [ø]  U+00F8 => F8
    converter.utf8ToLatin[371] = 249  // [ų]  U+0173 => F9
    converter.utf8ToLatin[250] = 250  // [ú]  U+00FA => FA
    converter.utf8ToLatin[251] = 251  // [û]  U+00FB => FB
    converter.utf8ToLatin[252] = 252  // [ü]  U+00FC => FC
    converter.utf8ToLatin[253] = 253  // [ý]  U+00FD => FD
    converter.utf8ToLatin[254] = 254  // [þ]  U+00FE => FE
    converter.utf8ToLatin[312] = 255  // [ĸ]  U+0138 => FF

    return converter
}
