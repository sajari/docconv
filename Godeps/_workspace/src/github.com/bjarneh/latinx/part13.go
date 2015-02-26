// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

// ISO 8859-13

// Generated from http://unicode.org/Public/MAPPINGS/ISO8859/8859-13.TXT

func newISO_8859_13Converter() *Converter {

    converter := new(Converter)
    converter.id = "Latin-7/ISO-8859-13"
    converter.latinToUtf8 = make(map[byte][]byte)

    converter.latinToUtf8[160] = []byte(" ") // NO-BREAK SPACE
    converter.latinToUtf8[161] = []byte("”") // RIGHT DOUBLE QUOTATION MARK
    converter.latinToUtf8[162] = []byte("¢") // CENT SIGN
    converter.latinToUtf8[163] = []byte("£") // POUND SIGN
    converter.latinToUtf8[164] = []byte("¤") // CURRENCY SIGN
    converter.latinToUtf8[165] = []byte("„") // DOUBLE LOW-9 QUOTATION MARK
    converter.latinToUtf8[166] = []byte("¦") // BROKEN BAR
    converter.latinToUtf8[167] = []byte("§") // SECTION SIGN
    converter.latinToUtf8[168] = []byte("Ø") // LATIN CAPITAL LETTER O WITH STROKE
    converter.latinToUtf8[169] = []byte("©") // COPYRIGHT SIGN
    converter.latinToUtf8[170] = []byte("Ŗ") // LATIN CAPITAL LETTER R WITH CEDILLA
    converter.latinToUtf8[171] = []byte("«") // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
    converter.latinToUtf8[172] = []byte("¬") // NOT SIGN
    converter.latinToUtf8[173] = []byte("­") // SOFT HYPHEN
    converter.latinToUtf8[174] = []byte("®") // REGISTERED SIGN
    converter.latinToUtf8[175] = []byte("Æ") // LATIN CAPITAL LETTER AE
    converter.latinToUtf8[176] = []byte("°") // DEGREE SIGN
    converter.latinToUtf8[177] = []byte("±") // PLUS-MINUS SIGN
    converter.latinToUtf8[178] = []byte("²") // SUPERSCRIPT TWO
    converter.latinToUtf8[179] = []byte("³") // SUPERSCRIPT THREE
    converter.latinToUtf8[180] = []byte("“") // LEFT DOUBLE QUOTATION MARK
    converter.latinToUtf8[181] = []byte("µ") // MICRO SIGN
    converter.latinToUtf8[182] = []byte("¶") // PILCROW SIGN
    converter.latinToUtf8[183] = []byte("·") // MIDDLE DOT
    converter.latinToUtf8[184] = []byte("ø") // LATIN SMALL LETTER O WITH STROKE
    converter.latinToUtf8[185] = []byte("¹") // SUPERSCRIPT ONE
    converter.latinToUtf8[186] = []byte("ŗ") // LATIN SMALL LETTER R WITH CEDILLA
    converter.latinToUtf8[187] = []byte("»") // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
    converter.latinToUtf8[188] = []byte("¼") // VULGAR FRACTION ONE QUARTER
    converter.latinToUtf8[189] = []byte("½") // VULGAR FRACTION ONE HALF
    converter.latinToUtf8[190] = []byte("¾") // VULGAR FRACTION THREE QUARTERS
    converter.latinToUtf8[191] = []byte("æ") // LATIN SMALL LETTER AE
    converter.latinToUtf8[192] = []byte("Ą") // LATIN CAPITAL LETTER A WITH OGONEK
    converter.latinToUtf8[193] = []byte("Į") // LATIN CAPITAL LETTER I WITH OGONEK
    converter.latinToUtf8[194] = []byte("Ā") // LATIN CAPITAL LETTER A WITH MACRON
    converter.latinToUtf8[195] = []byte("Ć") // LATIN CAPITAL LETTER C WITH ACUTE
    converter.latinToUtf8[196] = []byte("Ä") // LATIN CAPITAL LETTER A WITH DIAERESIS
    converter.latinToUtf8[197] = []byte("Å") // LATIN CAPITAL LETTER A WITH RING ABOVE
    converter.latinToUtf8[198] = []byte("Ę") // LATIN CAPITAL LETTER E WITH OGONEK
    converter.latinToUtf8[199] = []byte("Ē") // LATIN CAPITAL LETTER E WITH MACRON
    converter.latinToUtf8[200] = []byte("Č") // LATIN CAPITAL LETTER C WITH CARON
    converter.latinToUtf8[201] = []byte("É") // LATIN CAPITAL LETTER E WITH ACUTE
    converter.latinToUtf8[202] = []byte("Ź") // LATIN CAPITAL LETTER Z WITH ACUTE
    converter.latinToUtf8[203] = []byte("Ė") // LATIN CAPITAL LETTER E WITH DOT ABOVE
    converter.latinToUtf8[204] = []byte("Ģ") // LATIN CAPITAL LETTER G WITH CEDILLA
    converter.latinToUtf8[205] = []byte("Ķ") // LATIN CAPITAL LETTER K WITH CEDILLA
    converter.latinToUtf8[206] = []byte("Ī") // LATIN CAPITAL LETTER I WITH MACRON
    converter.latinToUtf8[207] = []byte("Ļ") // LATIN CAPITAL LETTER L WITH CEDILLA
    converter.latinToUtf8[208] = []byte("Š") // LATIN CAPITAL LETTER S WITH CARON
    converter.latinToUtf8[209] = []byte("Ń") // LATIN CAPITAL LETTER N WITH ACUTE
    converter.latinToUtf8[210] = []byte("Ņ") // LATIN CAPITAL LETTER N WITH CEDILLA
    converter.latinToUtf8[211] = []byte("Ó") // LATIN CAPITAL LETTER O WITH ACUTE
    converter.latinToUtf8[212] = []byte("Ō") // LATIN CAPITAL LETTER O WITH MACRON
    converter.latinToUtf8[213] = []byte("Õ") // LATIN CAPITAL LETTER O WITH TILDE
    converter.latinToUtf8[214] = []byte("Ö") // LATIN CAPITAL LETTER O WITH DIAERESIS
    converter.latinToUtf8[215] = []byte("×") // MULTIPLICATION SIGN
    converter.latinToUtf8[216] = []byte("Ų") // LATIN CAPITAL LETTER U WITH OGONEK
    converter.latinToUtf8[217] = []byte("Ł") // LATIN CAPITAL LETTER L WITH STROKE
    converter.latinToUtf8[218] = []byte("Ś") // LATIN CAPITAL LETTER S WITH ACUTE
    converter.latinToUtf8[219] = []byte("Ū") // LATIN CAPITAL LETTER U WITH MACRON
    converter.latinToUtf8[220] = []byte("Ü") // LATIN CAPITAL LETTER U WITH DIAERESIS
    converter.latinToUtf8[221] = []byte("Ż") // LATIN CAPITAL LETTER Z WITH DOT ABOVE
    converter.latinToUtf8[222] = []byte("Ž") // LATIN CAPITAL LETTER Z WITH CARON
    converter.latinToUtf8[223] = []byte("ß") // LATIN SMALL LETTER SHARP S (German)
    converter.latinToUtf8[224] = []byte("ą") // LATIN SMALL LETTER A WITH OGONEK
    converter.latinToUtf8[225] = []byte("į") // LATIN SMALL LETTER I WITH OGONEK
    converter.latinToUtf8[226] = []byte("ā") // LATIN SMALL LETTER A WITH MACRON
    converter.latinToUtf8[227] = []byte("ć") // LATIN SMALL LETTER C WITH ACUTE
    converter.latinToUtf8[228] = []byte("ä") // LATIN SMALL LETTER A WITH DIAERESIS
    converter.latinToUtf8[229] = []byte("å") // LATIN SMALL LETTER A WITH RING ABOVE
    converter.latinToUtf8[230] = []byte("ę") // LATIN SMALL LETTER E WITH OGONEK
    converter.latinToUtf8[231] = []byte("ē") // LATIN SMALL LETTER E WITH MACRON
    converter.latinToUtf8[232] = []byte("č") // LATIN SMALL LETTER C WITH CARON
    converter.latinToUtf8[233] = []byte("é") // LATIN SMALL LETTER E WITH ACUTE
    converter.latinToUtf8[234] = []byte("ź") // LATIN SMALL LETTER Z WITH ACUTE
    converter.latinToUtf8[235] = []byte("ė") // LATIN SMALL LETTER E WITH DOT ABOVE
    converter.latinToUtf8[236] = []byte("ģ") // LATIN SMALL LETTER G WITH CEDILLA
    converter.latinToUtf8[237] = []byte("ķ") // LATIN SMALL LETTER K WITH CEDILLA
    converter.latinToUtf8[238] = []byte("ī") // LATIN SMALL LETTER I WITH MACRON
    converter.latinToUtf8[239] = []byte("ļ") // LATIN SMALL LETTER L WITH CEDILLA
    converter.latinToUtf8[240] = []byte("š") // LATIN SMALL LETTER S WITH CARON
    converter.latinToUtf8[241] = []byte("ń") // LATIN SMALL LETTER N WITH ACUTE
    converter.latinToUtf8[242] = []byte("ņ") // LATIN SMALL LETTER N WITH CEDILLA
    converter.latinToUtf8[243] = []byte("ó") // LATIN SMALL LETTER O WITH ACUTE
    converter.latinToUtf8[244] = []byte("ō") // LATIN SMALL LETTER O WITH MACRON
    converter.latinToUtf8[245] = []byte("õ") // LATIN SMALL LETTER O WITH TILDE
    converter.latinToUtf8[246] = []byte("ö") // LATIN SMALL LETTER O WITH DIAERESIS
    converter.latinToUtf8[247] = []byte("÷") // DIVISION SIGN
    converter.latinToUtf8[248] = []byte("ų") // LATIN SMALL LETTER U WITH OGONEK
    converter.latinToUtf8[249] = []byte("ł") // LATIN SMALL LETTER L WITH STROKE
    converter.latinToUtf8[250] = []byte("ś") // LATIN SMALL LETTER S WITH ACUTE
    converter.latinToUtf8[251] = []byte("ū") // LATIN SMALL LETTER U WITH MACRON
    converter.latinToUtf8[252] = []byte("ü") // LATIN SMALL LETTER U WITH DIAERESIS
    converter.latinToUtf8[253] = []byte("ż") // LATIN SMALL LETTER Z WITH DOT ABOVE
    converter.latinToUtf8[254] = []byte("ž") // LATIN SMALL LETTER Z WITH CARON
    converter.latinToUtf8[255] = []byte("’") // RIGHT SINGLE QUOTATION MARK

    converter.utf8ToLatin = make(map[int]byte)

    converter.utf8ToLatin[160] = 160  // [ ]  U+00A0 => A0
    converter.utf8ToLatin[8221] = 161 // [”]  U+201D => A1
    converter.utf8ToLatin[162] = 162  // [¢]  U+00A2 => A2
    converter.utf8ToLatin[163] = 163  // [£]  U+00A3 => A3
    converter.utf8ToLatin[164] = 164  // [¤]  U+00A4 => A4
    converter.utf8ToLatin[8222] = 165 // [„]  U+201E => A5
    converter.utf8ToLatin[166] = 166  // [¦]  U+00A6 => A6
    converter.utf8ToLatin[167] = 167  // [§]  U+00A7 => A7
    converter.utf8ToLatin[216] = 168  // [Ø]  U+00D8 => A8
    converter.utf8ToLatin[169] = 169  // [©]  U+00A9 => A9
    converter.utf8ToLatin[342] = 170  // [Ŗ]  U+0156 => AA
    converter.utf8ToLatin[171] = 171  // [«]  U+00AB => AB
    converter.utf8ToLatin[172] = 172  // [¬]  U+00AC => AC
    converter.utf8ToLatin[173] = 173  // [­]  U+00AD => AD
    converter.utf8ToLatin[174] = 174  // [®]  U+00AE => AE
    converter.utf8ToLatin[198] = 175  // [Æ]  U+00C6 => AF
    converter.utf8ToLatin[176] = 176  // [°]  U+00B0 => B0
    converter.utf8ToLatin[177] = 177  // [±]  U+00B1 => B1
    converter.utf8ToLatin[178] = 178  // [²]  U+00B2 => B2
    converter.utf8ToLatin[179] = 179  // [³]  U+00B3 => B3
    converter.utf8ToLatin[8220] = 180 // [“]  U+201C => B4
    converter.utf8ToLatin[181] = 181  // [µ]  U+00B5 => B5
    converter.utf8ToLatin[182] = 182  // [¶]  U+00B6 => B6
    converter.utf8ToLatin[183] = 183  // [·]  U+00B7 => B7
    converter.utf8ToLatin[248] = 184  // [ø]  U+00F8 => B8
    converter.utf8ToLatin[185] = 185  // [¹]  U+00B9 => B9
    converter.utf8ToLatin[343] = 186  // [ŗ]  U+0157 => BA
    converter.utf8ToLatin[187] = 187  // [»]  U+00BB => BB
    converter.utf8ToLatin[188] = 188  // [¼]  U+00BC => BC
    converter.utf8ToLatin[189] = 189  // [½]  U+00BD => BD
    converter.utf8ToLatin[190] = 190  // [¾]  U+00BE => BE
    converter.utf8ToLatin[230] = 191  // [æ]  U+00E6 => BF
    converter.utf8ToLatin[260] = 192  // [Ą]  U+0104 => C0
    converter.utf8ToLatin[302] = 193  // [Į]  U+012E => C1
    converter.utf8ToLatin[256] = 194  // [Ā]  U+0100 => C2
    converter.utf8ToLatin[262] = 195  // [Ć]  U+0106 => C3
    converter.utf8ToLatin[196] = 196  // [Ä]  U+00C4 => C4
    converter.utf8ToLatin[197] = 197  // [Å]  U+00C5 => C5
    converter.utf8ToLatin[280] = 198  // [Ę]  U+0118 => C6
    converter.utf8ToLatin[274] = 199  // [Ē]  U+0112 => C7
    converter.utf8ToLatin[268] = 200  // [Č]  U+010C => C8
    converter.utf8ToLatin[201] = 201  // [É]  U+00C9 => C9
    converter.utf8ToLatin[377] = 202  // [Ź]  U+0179 => CA
    converter.utf8ToLatin[278] = 203  // [Ė]  U+0116 => CB
    converter.utf8ToLatin[290] = 204  // [Ģ]  U+0122 => CC
    converter.utf8ToLatin[310] = 205  // [Ķ]  U+0136 => CD
    converter.utf8ToLatin[298] = 206  // [Ī]  U+012A => CE
    converter.utf8ToLatin[315] = 207  // [Ļ]  U+013B => CF
    converter.utf8ToLatin[352] = 208  // [Š]  U+0160 => D0
    converter.utf8ToLatin[323] = 209  // [Ń]  U+0143 => D1
    converter.utf8ToLatin[325] = 210  // [Ņ]  U+0145 => D2
    converter.utf8ToLatin[211] = 211  // [Ó]  U+00D3 => D3
    converter.utf8ToLatin[332] = 212  // [Ō]  U+014C => D4
    converter.utf8ToLatin[213] = 213  // [Õ]  U+00D5 => D5
    converter.utf8ToLatin[214] = 214  // [Ö]  U+00D6 => D6
    converter.utf8ToLatin[215] = 215  // [×]  U+00D7 => D7
    converter.utf8ToLatin[370] = 216  // [Ų]  U+0172 => D8
    converter.utf8ToLatin[321] = 217  // [Ł]  U+0141 => D9
    converter.utf8ToLatin[346] = 218  // [Ś]  U+015A => DA
    converter.utf8ToLatin[362] = 219  // [Ū]  U+016A => DB
    converter.utf8ToLatin[220] = 220  // [Ü]  U+00DC => DC
    converter.utf8ToLatin[379] = 221  // [Ż]  U+017B => DD
    converter.utf8ToLatin[381] = 222  // [Ž]  U+017D => DE
    converter.utf8ToLatin[223] = 223  // [ß]  U+00DF => DF
    converter.utf8ToLatin[261] = 224  // [ą]  U+0105 => E0
    converter.utf8ToLatin[303] = 225  // [į]  U+012F => E1
    converter.utf8ToLatin[257] = 226  // [ā]  U+0101 => E2
    converter.utf8ToLatin[263] = 227  // [ć]  U+0107 => E3
    converter.utf8ToLatin[228] = 228  // [ä]  U+00E4 => E4
    converter.utf8ToLatin[229] = 229  // [å]  U+00E5 => E5
    converter.utf8ToLatin[281] = 230  // [ę]  U+0119 => E6
    converter.utf8ToLatin[275] = 231  // [ē]  U+0113 => E7
    converter.utf8ToLatin[269] = 232  // [č]  U+010D => E8
    converter.utf8ToLatin[233] = 233  // [é]  U+00E9 => E9
    converter.utf8ToLatin[378] = 234  // [ź]  U+017A => EA
    converter.utf8ToLatin[279] = 235  // [ė]  U+0117 => EB
    converter.utf8ToLatin[291] = 236  // [ģ]  U+0123 => EC
    converter.utf8ToLatin[311] = 237  // [ķ]  U+0137 => ED
    converter.utf8ToLatin[299] = 238  // [ī]  U+012B => EE
    converter.utf8ToLatin[316] = 239  // [ļ]  U+013C => EF
    converter.utf8ToLatin[353] = 240  // [š]  U+0161 => F0
    converter.utf8ToLatin[324] = 241  // [ń]  U+0144 => F1
    converter.utf8ToLatin[326] = 242  // [ņ]  U+0146 => F2
    converter.utf8ToLatin[243] = 243  // [ó]  U+00F3 => F3
    converter.utf8ToLatin[333] = 244  // [ō]  U+014D => F4
    converter.utf8ToLatin[245] = 245  // [õ]  U+00F5 => F5
    converter.utf8ToLatin[246] = 246  // [ö]  U+00F6 => F6
    converter.utf8ToLatin[247] = 247  // [÷]  U+00F7 => F7
    converter.utf8ToLatin[371] = 248  // [ų]  U+0173 => F8
    converter.utf8ToLatin[322] = 249  // [ł]  U+0142 => F9
    converter.utf8ToLatin[347] = 250  // [ś]  U+015B => FA
    converter.utf8ToLatin[363] = 251  // [ū]  U+016B => FB
    converter.utf8ToLatin[252] = 252  // [ü]  U+00FC => FC
    converter.utf8ToLatin[380] = 253  // [ż]  U+017C => FD
    converter.utf8ToLatin[382] = 254  // [ž]  U+017E => FE
    converter.utf8ToLatin[8217] = 255 // [’]  U+2019 => FF

    return converter
}
