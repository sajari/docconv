// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

// ISO 8859-16

// Generated from http://unicode.org/Public/MAPPINGS/ISO8859/8859-16.TXT

func newISO_8859_16Converter() *Converter {

    converter := new(Converter)
    converter.id = "Latin-10/ISO-8859-16"
    converter.latinToUtf8 = make(map[byte][]byte)

    converter.latinToUtf8[160] = []byte(" ") // NO-BREAK SPACE
    converter.latinToUtf8[161] = []byte("Ą") // LATIN CAPITAL LETTER A WITH OGONEK
    converter.latinToUtf8[162] = []byte("ą") // LATIN SMALL LETTER A WITH OGONEK
    converter.latinToUtf8[163] = []byte("Ł") // LATIN CAPITAL LETTER L WITH STROKE
    converter.latinToUtf8[164] = []byte("€") // EURO SIGN
    converter.latinToUtf8[165] = []byte("„") // DOUBLE LOW-9 QUOTATION MARK
    converter.latinToUtf8[166] = []byte("Š") // LATIN CAPITAL LETTER S WITH CARON
    converter.latinToUtf8[167] = []byte("§") // SECTION SIGN
    converter.latinToUtf8[168] = []byte("š") // LATIN SMALL LETTER S WITH CARON
    converter.latinToUtf8[169] = []byte("©") // COPYRIGHT SIGN
    converter.latinToUtf8[170] = []byte("Ș") // LATIN CAPITAL LETTER S WITH COMMA BELOW
    converter.latinToUtf8[171] = []byte("«") // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
    converter.latinToUtf8[172] = []byte("Ź") // LATIN CAPITAL LETTER Z WITH ACUTE
    converter.latinToUtf8[173] = []byte("­") // SOFT HYPHEN
    converter.latinToUtf8[174] = []byte("ź") // LATIN SMALL LETTER Z WITH ACUTE
    converter.latinToUtf8[175] = []byte("Ż") // LATIN CAPITAL LETTER Z WITH DOT ABOVE
    converter.latinToUtf8[176] = []byte("°") // DEGREE SIGN
    converter.latinToUtf8[177] = []byte("±") // PLUS-MINUS SIGN
    converter.latinToUtf8[178] = []byte("Č") // LATIN CAPITAL LETTER C WITH CARON
    converter.latinToUtf8[179] = []byte("ł") // LATIN SMALL LETTER L WITH STROKE
    converter.latinToUtf8[180] = []byte("Ž") // LATIN CAPITAL LETTER Z WITH CARON
    converter.latinToUtf8[181] = []byte("”") // RIGHT DOUBLE QUOTATION MARK
    converter.latinToUtf8[182] = []byte("¶") // PILCROW SIGN
    converter.latinToUtf8[183] = []byte("·") // MIDDLE DOT
    converter.latinToUtf8[184] = []byte("ž") // LATIN SMALL LETTER Z WITH CARON
    converter.latinToUtf8[185] = []byte("č") // LATIN SMALL LETTER C WITH CARON
    converter.latinToUtf8[186] = []byte("ș") // LATIN SMALL LETTER S WITH COMMA BELOW
    converter.latinToUtf8[187] = []byte("»") // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
    converter.latinToUtf8[188] = []byte("Œ") // LATIN CAPITAL LIGATURE OE
    converter.latinToUtf8[189] = []byte("œ") // LATIN SMALL LIGATURE OE
    converter.latinToUtf8[190] = []byte("Ÿ") // LATIN CAPITAL LETTER Y WITH DIAERESIS
    converter.latinToUtf8[191] = []byte("ż") // LATIN SMALL LETTER Z WITH DOT ABOVE
    converter.latinToUtf8[192] = []byte("À") // LATIN CAPITAL LETTER A WITH GRAVE
    converter.latinToUtf8[193] = []byte("Á") // LATIN CAPITAL LETTER A WITH ACUTE
    converter.latinToUtf8[194] = []byte("Â") // LATIN CAPITAL LETTER A WITH CIRCUMFLEX
    converter.latinToUtf8[195] = []byte("Ă") // LATIN CAPITAL LETTER A WITH BREVE
    converter.latinToUtf8[196] = []byte("Ä") // LATIN CAPITAL LETTER A WITH DIAERESIS
    converter.latinToUtf8[197] = []byte("Ć") // LATIN CAPITAL LETTER C WITH ACUTE
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
    converter.latinToUtf8[208] = []byte("Đ") // LATIN CAPITAL LETTER D WITH STROKE
    converter.latinToUtf8[209] = []byte("Ń") // LATIN CAPITAL LETTER N WITH ACUTE
    converter.latinToUtf8[210] = []byte("Ò") // LATIN CAPITAL LETTER O WITH GRAVE
    converter.latinToUtf8[211] = []byte("Ó") // LATIN CAPITAL LETTER O WITH ACUTE
    converter.latinToUtf8[212] = []byte("Ô") // LATIN CAPITAL LETTER O WITH CIRCUMFLEX
    converter.latinToUtf8[213] = []byte("Ő") // LATIN CAPITAL LETTER O WITH DOUBLE ACUTE
    converter.latinToUtf8[214] = []byte("Ö") // LATIN CAPITAL LETTER O WITH DIAERESIS
    converter.latinToUtf8[215] = []byte("Ś") // LATIN CAPITAL LETTER S WITH ACUTE
    converter.latinToUtf8[216] = []byte("Ű") // LATIN CAPITAL LETTER U WITH DOUBLE ACUTE
    converter.latinToUtf8[217] = []byte("Ù") // LATIN CAPITAL LETTER U WITH GRAVE
    converter.latinToUtf8[218] = []byte("Ú") // LATIN CAPITAL LETTER U WITH ACUTE
    converter.latinToUtf8[219] = []byte("Û") // LATIN CAPITAL LETTER U WITH CIRCUMFLEX
    converter.latinToUtf8[220] = []byte("Ü") // LATIN CAPITAL LETTER U WITH DIAERESIS
    converter.latinToUtf8[221] = []byte("Ę") // LATIN CAPITAL LETTER E WITH OGONEK
    converter.latinToUtf8[222] = []byte("Ț") // LATIN CAPITAL LETTER T WITH COMMA BELOW
    converter.latinToUtf8[223] = []byte("ß") // LATIN SMALL LETTER SHARP S
    converter.latinToUtf8[224] = []byte("à") // LATIN SMALL LETTER A WITH GRAVE
    converter.latinToUtf8[225] = []byte("á") // LATIN SMALL LETTER A WITH ACUTE
    converter.latinToUtf8[226] = []byte("â") // LATIN SMALL LETTER A WITH CIRCUMFLEX
    converter.latinToUtf8[227] = []byte("ă") // LATIN SMALL LETTER A WITH BREVE
    converter.latinToUtf8[228] = []byte("ä") // LATIN SMALL LETTER A WITH DIAERESIS
    converter.latinToUtf8[229] = []byte("ć") // LATIN SMALL LETTER C WITH ACUTE
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
    converter.latinToUtf8[240] = []byte("đ") // LATIN SMALL LETTER D WITH STROKE
    converter.latinToUtf8[241] = []byte("ń") // LATIN SMALL LETTER N WITH ACUTE
    converter.latinToUtf8[242] = []byte("ò") // LATIN SMALL LETTER O WITH GRAVE
    converter.latinToUtf8[243] = []byte("ó") // LATIN SMALL LETTER O WITH ACUTE
    converter.latinToUtf8[244] = []byte("ô") // LATIN SMALL LETTER O WITH CIRCUMFLEX
    converter.latinToUtf8[245] = []byte("ő") // LATIN SMALL LETTER O WITH DOUBLE ACUTE
    converter.latinToUtf8[246] = []byte("ö") // LATIN SMALL LETTER O WITH DIAERESIS
    converter.latinToUtf8[247] = []byte("ś") // LATIN SMALL LETTER S WITH ACUTE
    converter.latinToUtf8[248] = []byte("ű") // LATIN SMALL LETTER U WITH DOUBLE ACUTE
    converter.latinToUtf8[249] = []byte("ù") // LATIN SMALL LETTER U WITH GRAVE
    converter.latinToUtf8[250] = []byte("ú") // LATIN SMALL LETTER U WITH ACUTE
    converter.latinToUtf8[251] = []byte("û") // LATIN SMALL LETTER U WITH CIRCUMFLEX
    converter.latinToUtf8[252] = []byte("ü") // LATIN SMALL LETTER U WITH DIAERESIS
    converter.latinToUtf8[253] = []byte("ę") // LATIN SMALL LETTER E WITH OGONEK
    converter.latinToUtf8[254] = []byte("ț") // LATIN SMALL LETTER T WITH COMMA BELOW
    converter.latinToUtf8[255] = []byte("ÿ") // LATIN SMALL LETTER Y WITH DIAERESIS

    converter.utf8ToLatin = make(map[int]byte)

    converter.utf8ToLatin[160] = 160  // [ ]  U+00A0 => A0
    converter.utf8ToLatin[260] = 161  // [Ą]  U+0104 => A1
    converter.utf8ToLatin[261] = 162  // [ą]  U+0105 => A2
    converter.utf8ToLatin[321] = 163  // [Ł]  U+0141 => A3
    converter.utf8ToLatin[8364] = 164 // [€]  U+20AC => A4
    converter.utf8ToLatin[8222] = 165 // [„]  U+201E => A5
    converter.utf8ToLatin[352] = 166  // [Š]  U+0160 => A6
    converter.utf8ToLatin[167] = 167  // [§]  U+00A7 => A7
    converter.utf8ToLatin[353] = 168  // [š]  U+0161 => A8
    converter.utf8ToLatin[169] = 169  // [©]  U+00A9 => A9
    converter.utf8ToLatin[536] = 170  // [Ș]  U+0218 => AA
    converter.utf8ToLatin[171] = 171  // [«]  U+00AB => AB
    converter.utf8ToLatin[377] = 172  // [Ź]  U+0179 => AC
    converter.utf8ToLatin[173] = 173  // [­]  U+00AD => AD
    converter.utf8ToLatin[378] = 174  // [ź]  U+017A => AE
    converter.utf8ToLatin[379] = 175  // [Ż]  U+017B => AF
    converter.utf8ToLatin[176] = 176  // [°]  U+00B0 => B0
    converter.utf8ToLatin[177] = 177  // [±]  U+00B1 => B1
    converter.utf8ToLatin[268] = 178  // [Č]  U+010C => B2
    converter.utf8ToLatin[322] = 179  // [ł]  U+0142 => B3
    converter.utf8ToLatin[381] = 180  // [Ž]  U+017D => B4
    converter.utf8ToLatin[8221] = 181 // [”]  U+201D => B5
    converter.utf8ToLatin[182] = 182  // [¶]  U+00B6 => B6
    converter.utf8ToLatin[183] = 183  // [·]  U+00B7 => B7
    converter.utf8ToLatin[382] = 184  // [ž]  U+017E => B8
    converter.utf8ToLatin[269] = 185  // [č]  U+010D => B9
    converter.utf8ToLatin[537] = 186  // [ș]  U+0219 => BA
    converter.utf8ToLatin[187] = 187  // [»]  U+00BB => BB
    converter.utf8ToLatin[338] = 188  // [Œ]  U+0152 => BC
    converter.utf8ToLatin[339] = 189  // [œ]  U+0153 => BD
    converter.utf8ToLatin[376] = 190  // [Ÿ]  U+0178 => BE
    converter.utf8ToLatin[380] = 191  // [ż]  U+017C => BF
    converter.utf8ToLatin[192] = 192  // [À]  U+00C0 => C0
    converter.utf8ToLatin[193] = 193  // [Á]  U+00C1 => C1
    converter.utf8ToLatin[194] = 194  // [Â]  U+00C2 => C2
    converter.utf8ToLatin[258] = 195  // [Ă]  U+0102 => C3
    converter.utf8ToLatin[196] = 196  // [Ä]  U+00C4 => C4
    converter.utf8ToLatin[262] = 197  // [Ć]  U+0106 => C5
    converter.utf8ToLatin[198] = 198  // [Æ]  U+00C6 => C6
    converter.utf8ToLatin[199] = 199  // [Ç]  U+00C7 => C7
    converter.utf8ToLatin[200] = 200  // [È]  U+00C8 => C8
    converter.utf8ToLatin[201] = 201  // [É]  U+00C9 => C9
    converter.utf8ToLatin[202] = 202  // [Ê]  U+00CA => CA
    converter.utf8ToLatin[203] = 203  // [Ë]  U+00CB => CB
    converter.utf8ToLatin[204] = 204  // [Ì]  U+00CC => CC
    converter.utf8ToLatin[205] = 205  // [Í]  U+00CD => CD
    converter.utf8ToLatin[206] = 206  // [Î]  U+00CE => CE
    converter.utf8ToLatin[207] = 207  // [Ï]  U+00CF => CF
    converter.utf8ToLatin[272] = 208  // [Đ]  U+0110 => D0
    converter.utf8ToLatin[323] = 209  // [Ń]  U+0143 => D1
    converter.utf8ToLatin[210] = 210  // [Ò]  U+00D2 => D2
    converter.utf8ToLatin[211] = 211  // [Ó]  U+00D3 => D3
    converter.utf8ToLatin[212] = 212  // [Ô]  U+00D4 => D4
    converter.utf8ToLatin[336] = 213  // [Ő]  U+0150 => D5
    converter.utf8ToLatin[214] = 214  // [Ö]  U+00D6 => D6
    converter.utf8ToLatin[346] = 215  // [Ś]  U+015A => D7
    converter.utf8ToLatin[368] = 216  // [Ű]  U+0170 => D8
    converter.utf8ToLatin[217] = 217  // [Ù]  U+00D9 => D9
    converter.utf8ToLatin[218] = 218  // [Ú]  U+00DA => DA
    converter.utf8ToLatin[219] = 219  // [Û]  U+00DB => DB
    converter.utf8ToLatin[220] = 220  // [Ü]  U+00DC => DC
    converter.utf8ToLatin[280] = 221  // [Ę]  U+0118 => DD
    converter.utf8ToLatin[538] = 222  // [Ț]  U+021A => DE
    converter.utf8ToLatin[223] = 223  // [ß]  U+00DF => DF
    converter.utf8ToLatin[224] = 224  // [à]  U+00E0 => E0
    converter.utf8ToLatin[225] = 225  // [á]  U+00E1 => E1
    converter.utf8ToLatin[226] = 226  // [â]  U+00E2 => E2
    converter.utf8ToLatin[259] = 227  // [ă]  U+0103 => E3
    converter.utf8ToLatin[228] = 228  // [ä]  U+00E4 => E4
    converter.utf8ToLatin[263] = 229  // [ć]  U+0107 => E5
    converter.utf8ToLatin[230] = 230  // [æ]  U+00E6 => E6
    converter.utf8ToLatin[231] = 231  // [ç]  U+00E7 => E7
    converter.utf8ToLatin[232] = 232  // [è]  U+00E8 => E8
    converter.utf8ToLatin[233] = 233  // [é]  U+00E9 => E9
    converter.utf8ToLatin[234] = 234  // [ê]  U+00EA => EA
    converter.utf8ToLatin[235] = 235  // [ë]  U+00EB => EB
    converter.utf8ToLatin[236] = 236  // [ì]  U+00EC => EC
    converter.utf8ToLatin[237] = 237  // [í]  U+00ED => ED
    converter.utf8ToLatin[238] = 238  // [î]  U+00EE => EE
    converter.utf8ToLatin[239] = 239  // [ï]  U+00EF => EF
    converter.utf8ToLatin[273] = 240  // [đ]  U+0111 => F0
    converter.utf8ToLatin[324] = 241  // [ń]  U+0144 => F1
    converter.utf8ToLatin[242] = 242  // [ò]  U+00F2 => F2
    converter.utf8ToLatin[243] = 243  // [ó]  U+00F3 => F3
    converter.utf8ToLatin[244] = 244  // [ô]  U+00F4 => F4
    converter.utf8ToLatin[337] = 245  // [ő]  U+0151 => F5
    converter.utf8ToLatin[246] = 246  // [ö]  U+00F6 => F6
    converter.utf8ToLatin[347] = 247  // [ś]  U+015B => F7
    converter.utf8ToLatin[369] = 248  // [ű]  U+0171 => F8
    converter.utf8ToLatin[249] = 249  // [ù]  U+00F9 => F9
    converter.utf8ToLatin[250] = 250  // [ú]  U+00FA => FA
    converter.utf8ToLatin[251] = 251  // [û]  U+00FB => FB
    converter.utf8ToLatin[252] = 252  // [ü]  U+00FC => FC
    converter.utf8ToLatin[281] = 253  // [ę]  U+0119 => FD
    converter.utf8ToLatin[539] = 254  // [ț]  U+021B => FE
    converter.utf8ToLatin[255] = 255  // [ÿ]  U+00FF => FF

    return converter
}
