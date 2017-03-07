// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

// ISO 8859-2

// Generated from http://unicode.org/Public/MAPPINGS/ISO8859/8859-2.TXT

func newISO_8859_2Converter() *Converter {

    converter := new(Converter)
    converter.id = "Latin-2/ISO-8859-2"
    converter.latinToUtf8 = make(map[byte][]byte)

    converter.latinToUtf8[160] = []byte(" ") // NO-BREAK SPACE
    converter.latinToUtf8[161] = []byte("Ą") // LATIN CAPITAL LETTER A WITH OGONEK
    converter.latinToUtf8[162] = []byte("˘") // BREVE
    converter.latinToUtf8[163] = []byte("Ł") // LATIN CAPITAL LETTER L WITH STROKE
    converter.latinToUtf8[164] = []byte("¤") // CURRENCY SIGN
    converter.latinToUtf8[165] = []byte("Ľ") // LATIN CAPITAL LETTER L WITH CARON
    converter.latinToUtf8[166] = []byte("Ś") // LATIN CAPITAL LETTER S WITH ACUTE
    converter.latinToUtf8[167] = []byte("§") // SECTION SIGN
    converter.latinToUtf8[168] = []byte("¨") // DIAERESIS
    converter.latinToUtf8[169] = []byte("Š") // LATIN CAPITAL LETTER S WITH CARON
    converter.latinToUtf8[170] = []byte("Ş") // LATIN CAPITAL LETTER S WITH CEDILLA
    converter.latinToUtf8[171] = []byte("Ť") // LATIN CAPITAL LETTER T WITH CARON
    converter.latinToUtf8[172] = []byte("Ź") // LATIN CAPITAL LETTER Z WITH ACUTE
    converter.latinToUtf8[173] = []byte("­") // SOFT HYPHEN
    converter.latinToUtf8[174] = []byte("Ž") // LATIN CAPITAL LETTER Z WITH CARON
    converter.latinToUtf8[175] = []byte("Ż") // LATIN CAPITAL LETTER Z WITH DOT ABOVE
    converter.latinToUtf8[176] = []byte("°") // DEGREE SIGN
    converter.latinToUtf8[177] = []byte("ą") // LATIN SMALL LETTER A WITH OGONEK
    converter.latinToUtf8[178] = []byte("˛") // OGONEK
    converter.latinToUtf8[179] = []byte("ł") // LATIN SMALL LETTER L WITH STROKE
    converter.latinToUtf8[180] = []byte("´") // ACUTE ACCENT
    converter.latinToUtf8[181] = []byte("ľ") // LATIN SMALL LETTER L WITH CARON
    converter.latinToUtf8[182] = []byte("ś") // LATIN SMALL LETTER S WITH ACUTE
    converter.latinToUtf8[183] = []byte("ˇ") // CARON
    converter.latinToUtf8[184] = []byte("¸") // CEDILLA
    converter.latinToUtf8[185] = []byte("š") // LATIN SMALL LETTER S WITH CARON
    converter.latinToUtf8[186] = []byte("ş") // LATIN SMALL LETTER S WITH CEDILLA
    converter.latinToUtf8[187] = []byte("ť") // LATIN SMALL LETTER T WITH CARON
    converter.latinToUtf8[188] = []byte("ź") // LATIN SMALL LETTER Z WITH ACUTE
    converter.latinToUtf8[189] = []byte("˝") // DOUBLE ACUTE ACCENT
    converter.latinToUtf8[190] = []byte("ž") // LATIN SMALL LETTER Z WITH CARON
    converter.latinToUtf8[191] = []byte("ż") // LATIN SMALL LETTER Z WITH DOT ABOVE
    converter.latinToUtf8[192] = []byte("Ŕ") // LATIN CAPITAL LETTER R WITH ACUTE
    converter.latinToUtf8[193] = []byte("Á") // LATIN CAPITAL LETTER A WITH ACUTE
    converter.latinToUtf8[194] = []byte("Â") // LATIN CAPITAL LETTER A WITH CIRCUMFLEX
    converter.latinToUtf8[195] = []byte("Ă") // LATIN CAPITAL LETTER A WITH BREVE
    converter.latinToUtf8[196] = []byte("Ä") // LATIN CAPITAL LETTER A WITH DIAERESIS
    converter.latinToUtf8[197] = []byte("Ĺ") // LATIN CAPITAL LETTER L WITH ACUTE
    converter.latinToUtf8[198] = []byte("Ć") // LATIN CAPITAL LETTER C WITH ACUTE
    converter.latinToUtf8[199] = []byte("Ç") // LATIN CAPITAL LETTER C WITH CEDILLA
    converter.latinToUtf8[200] = []byte("Č") // LATIN CAPITAL LETTER C WITH CARON
    converter.latinToUtf8[201] = []byte("É") // LATIN CAPITAL LETTER E WITH ACUTE
    converter.latinToUtf8[202] = []byte("Ę") // LATIN CAPITAL LETTER E WITH OGONEK
    converter.latinToUtf8[203] = []byte("Ë") // LATIN CAPITAL LETTER E WITH DIAERESIS
    converter.latinToUtf8[204] = []byte("Ě") // LATIN CAPITAL LETTER E WITH CARON
    converter.latinToUtf8[205] = []byte("Í") // LATIN CAPITAL LETTER I WITH ACUTE
    converter.latinToUtf8[206] = []byte("Î") // LATIN CAPITAL LETTER I WITH CIRCUMFLEX
    converter.latinToUtf8[207] = []byte("Ď") // LATIN CAPITAL LETTER D WITH CARON
    converter.latinToUtf8[208] = []byte("Đ") // LATIN CAPITAL LETTER D WITH STROKE
    converter.latinToUtf8[209] = []byte("Ń") // LATIN CAPITAL LETTER N WITH ACUTE
    converter.latinToUtf8[210] = []byte("Ň") // LATIN CAPITAL LETTER N WITH CARON
    converter.latinToUtf8[211] = []byte("Ó") // LATIN CAPITAL LETTER O WITH ACUTE
    converter.latinToUtf8[212] = []byte("Ô") // LATIN CAPITAL LETTER O WITH CIRCUMFLEX
    converter.latinToUtf8[213] = []byte("Ő") // LATIN CAPITAL LETTER O WITH DOUBLE ACUTE
    converter.latinToUtf8[214] = []byte("Ö") // LATIN CAPITAL LETTER O WITH DIAERESIS
    converter.latinToUtf8[215] = []byte("×") // MULTIPLICATION SIGN
    converter.latinToUtf8[216] = []byte("Ř") // LATIN CAPITAL LETTER R WITH CARON
    converter.latinToUtf8[217] = []byte("Ů") // LATIN CAPITAL LETTER U WITH RING ABOVE
    converter.latinToUtf8[218] = []byte("Ú") // LATIN CAPITAL LETTER U WITH ACUTE
    converter.latinToUtf8[219] = []byte("Ű") // LATIN CAPITAL LETTER U WITH DOUBLE ACUTE
    converter.latinToUtf8[220] = []byte("Ü") // LATIN CAPITAL LETTER U WITH DIAERESIS
    converter.latinToUtf8[221] = []byte("Ý") // LATIN CAPITAL LETTER Y WITH ACUTE
    converter.latinToUtf8[222] = []byte("Ţ") // LATIN CAPITAL LETTER T WITH CEDILLA
    converter.latinToUtf8[223] = []byte("ß") // LATIN SMALL LETTER SHARP S
    converter.latinToUtf8[224] = []byte("ŕ") // LATIN SMALL LETTER R WITH ACUTE
    converter.latinToUtf8[225] = []byte("á") // LATIN SMALL LETTER A WITH ACUTE
    converter.latinToUtf8[226] = []byte("â") // LATIN SMALL LETTER A WITH CIRCUMFLEX
    converter.latinToUtf8[227] = []byte("ă") // LATIN SMALL LETTER A WITH BREVE
    converter.latinToUtf8[228] = []byte("ä") // LATIN SMALL LETTER A WITH DIAERESIS
    converter.latinToUtf8[229] = []byte("ĺ") // LATIN SMALL LETTER L WITH ACUTE
    converter.latinToUtf8[230] = []byte("ć") // LATIN SMALL LETTER C WITH ACUTE
    converter.latinToUtf8[231] = []byte("ç") // LATIN SMALL LETTER C WITH CEDILLA
    converter.latinToUtf8[232] = []byte("č") // LATIN SMALL LETTER C WITH CARON
    converter.latinToUtf8[233] = []byte("é") // LATIN SMALL LETTER E WITH ACUTE
    converter.latinToUtf8[234] = []byte("ę") // LATIN SMALL LETTER E WITH OGONEK
    converter.latinToUtf8[235] = []byte("ë") // LATIN SMALL LETTER E WITH DIAERESIS
    converter.latinToUtf8[236] = []byte("ě") // LATIN SMALL LETTER E WITH CARON
    converter.latinToUtf8[237] = []byte("í") // LATIN SMALL LETTER I WITH ACUTE
    converter.latinToUtf8[238] = []byte("î") // LATIN SMALL LETTER I WITH CIRCUMFLEX
    converter.latinToUtf8[239] = []byte("ď") // LATIN SMALL LETTER D WITH CARON
    converter.latinToUtf8[240] = []byte("đ") // LATIN SMALL LETTER D WITH STROKE
    converter.latinToUtf8[241] = []byte("ń") // LATIN SMALL LETTER N WITH ACUTE
    converter.latinToUtf8[242] = []byte("ň") // LATIN SMALL LETTER N WITH CARON
    converter.latinToUtf8[243] = []byte("ó") // LATIN SMALL LETTER O WITH ACUTE
    converter.latinToUtf8[244] = []byte("ô") // LATIN SMALL LETTER O WITH CIRCUMFLEX
    converter.latinToUtf8[245] = []byte("ő") // LATIN SMALL LETTER O WITH DOUBLE ACUTE
    converter.latinToUtf8[246] = []byte("ö") // LATIN SMALL LETTER O WITH DIAERESIS
    converter.latinToUtf8[247] = []byte("÷") // DIVISION SIGN
    converter.latinToUtf8[248] = []byte("ř") // LATIN SMALL LETTER R WITH CARON
    converter.latinToUtf8[249] = []byte("ů") // LATIN SMALL LETTER U WITH RING ABOVE
    converter.latinToUtf8[250] = []byte("ú") // LATIN SMALL LETTER U WITH ACUTE
    converter.latinToUtf8[251] = []byte("ű") // LATIN SMALL LETTER U WITH DOUBLE ACUTE
    converter.latinToUtf8[252] = []byte("ü") // LATIN SMALL LETTER U WITH DIAERESIS
    converter.latinToUtf8[253] = []byte("ý") // LATIN SMALL LETTER Y WITH ACUTE
    converter.latinToUtf8[254] = []byte("ţ") // LATIN SMALL LETTER T WITH CEDILLA
    converter.latinToUtf8[255] = []byte("˙") // DOT ABOVE


    converter.utf8ToLatin = make(map[int]byte)

    converter.utf8ToLatin[160] = 160 // [ ]  U+00A0 => A0
    converter.utf8ToLatin[260] = 161 // [Ą]  U+0104 => A1
    converter.utf8ToLatin[728] = 162 // [˘]  U+02D8 => A2
    converter.utf8ToLatin[321] = 163 // [Ł]  U+0141 => A3
    converter.utf8ToLatin[164] = 164 // [¤]  U+00A4 => A4
    converter.utf8ToLatin[317] = 165 // [Ľ]  U+013D => A5
    converter.utf8ToLatin[346] = 166 // [Ś]  U+015A => A6
    converter.utf8ToLatin[167] = 167 // [§]  U+00A7 => A7
    converter.utf8ToLatin[168] = 168 // [¨]  U+00A8 => A8
    converter.utf8ToLatin[352] = 169 // [Š]  U+0160 => A9
    converter.utf8ToLatin[350] = 170 // [Ş]  U+015E => AA
    converter.utf8ToLatin[356] = 171 // [Ť]  U+0164 => AB
    converter.utf8ToLatin[377] = 172 // [Ź]  U+0179 => AC
    converter.utf8ToLatin[173] = 173 // [­]  U+00AD => AD
    converter.utf8ToLatin[381] = 174 // [Ž]  U+017D => AE
    converter.utf8ToLatin[379] = 175 // [Ż]  U+017B => AF
    converter.utf8ToLatin[176] = 176 // [°]  U+00B0 => B0
    converter.utf8ToLatin[261] = 177 // [ą]  U+0105 => B1
    converter.utf8ToLatin[731] = 178 // [˛]  U+02DB => B2
    converter.utf8ToLatin[322] = 179 // [ł]  U+0142 => B3
    converter.utf8ToLatin[180] = 180 // [´]  U+00B4 => B4
    converter.utf8ToLatin[318] = 181 // [ľ]  U+013E => B5
    converter.utf8ToLatin[347] = 182 // [ś]  U+015B => B6
    converter.utf8ToLatin[711] = 183 // [ˇ]  U+02C7 => B7
    converter.utf8ToLatin[184] = 184 // [¸]  U+00B8 => B8
    converter.utf8ToLatin[353] = 185 // [š]  U+0161 => B9
    converter.utf8ToLatin[351] = 186 // [ş]  U+015F => BA
    converter.utf8ToLatin[357] = 187 // [ť]  U+0165 => BB
    converter.utf8ToLatin[378] = 188 // [ź]  U+017A => BC
    converter.utf8ToLatin[733] = 189 // [˝]  U+02DD => BD
    converter.utf8ToLatin[382] = 190 // [ž]  U+017E => BE
    converter.utf8ToLatin[380] = 191 // [ż]  U+017C => BF
    converter.utf8ToLatin[340] = 192 // [Ŕ]  U+0154 => C0
    converter.utf8ToLatin[193] = 193 // [Á]  U+00C1 => C1
    converter.utf8ToLatin[194] = 194 // [Â]  U+00C2 => C2
    converter.utf8ToLatin[258] = 195 // [Ă]  U+0102 => C3
    converter.utf8ToLatin[196] = 196 // [Ä]  U+00C4 => C4
    converter.utf8ToLatin[313] = 197 // [Ĺ]  U+0139 => C5
    converter.utf8ToLatin[262] = 198 // [Ć]  U+0106 => C6
    converter.utf8ToLatin[199] = 199 // [Ç]  U+00C7 => C7
    converter.utf8ToLatin[268] = 200 // [Č]  U+010C => C8
    converter.utf8ToLatin[201] = 201 // [É]  U+00C9 => C9
    converter.utf8ToLatin[280] = 202 // [Ę]  U+0118 => CA
    converter.utf8ToLatin[203] = 203 // [Ë]  U+00CB => CB
    converter.utf8ToLatin[282] = 204 // [Ě]  U+011A => CC
    converter.utf8ToLatin[205] = 205 // [Í]  U+00CD => CD
    converter.utf8ToLatin[206] = 206 // [Î]  U+00CE => CE
    converter.utf8ToLatin[270] = 207 // [Ď]  U+010E => CF
    converter.utf8ToLatin[272] = 208 // [Đ]  U+0110 => D0
    converter.utf8ToLatin[323] = 209 // [Ń]  U+0143 => D1
    converter.utf8ToLatin[327] = 210 // [Ň]  U+0147 => D2
    converter.utf8ToLatin[211] = 211 // [Ó]  U+00D3 => D3
    converter.utf8ToLatin[212] = 212 // [Ô]  U+00D4 => D4
    converter.utf8ToLatin[336] = 213 // [Ő]  U+0150 => D5
    converter.utf8ToLatin[214] = 214 // [Ö]  U+00D6 => D6
    converter.utf8ToLatin[215] = 215 // [×]  U+00D7 => D7
    converter.utf8ToLatin[344] = 216 // [Ř]  U+0158 => D8
    converter.utf8ToLatin[366] = 217 // [Ů]  U+016E => D9
    converter.utf8ToLatin[218] = 218 // [Ú]  U+00DA => DA
    converter.utf8ToLatin[368] = 219 // [Ű]  U+0170 => DB
    converter.utf8ToLatin[220] = 220 // [Ü]  U+00DC => DC
    converter.utf8ToLatin[221] = 221 // [Ý]  U+00DD => DD
    converter.utf8ToLatin[354] = 222 // [Ţ]  U+0162 => DE
    converter.utf8ToLatin[223] = 223 // [ß]  U+00DF => DF
    converter.utf8ToLatin[341] = 224 // [ŕ]  U+0155 => E0
    converter.utf8ToLatin[225] = 225 // [á]  U+00E1 => E1
    converter.utf8ToLatin[226] = 226 // [â]  U+00E2 => E2
    converter.utf8ToLatin[259] = 227 // [ă]  U+0103 => E3
    converter.utf8ToLatin[228] = 228 // [ä]  U+00E4 => E4
    converter.utf8ToLatin[314] = 229 // [ĺ]  U+013A => E5
    converter.utf8ToLatin[263] = 230 // [ć]  U+0107 => E6
    converter.utf8ToLatin[231] = 231 // [ç]  U+00E7 => E7
    converter.utf8ToLatin[269] = 232 // [č]  U+010D => E8
    converter.utf8ToLatin[233] = 233 // [é]  U+00E9 => E9
    converter.utf8ToLatin[281] = 234 // [ę]  U+0119 => EA
    converter.utf8ToLatin[235] = 235 // [ë]  U+00EB => EB
    converter.utf8ToLatin[283] = 236 // [ě]  U+011B => EC
    converter.utf8ToLatin[237] = 237 // [í]  U+00ED => ED
    converter.utf8ToLatin[238] = 238 // [î]  U+00EE => EE
    converter.utf8ToLatin[271] = 239 // [ď]  U+010F => EF
    converter.utf8ToLatin[273] = 240 // [đ]  U+0111 => F0
    converter.utf8ToLatin[324] = 241 // [ń]  U+0144 => F1
    converter.utf8ToLatin[328] = 242 // [ň]  U+0148 => F2
    converter.utf8ToLatin[243] = 243 // [ó]  U+00F3 => F3
    converter.utf8ToLatin[244] = 244 // [ô]  U+00F4 => F4
    converter.utf8ToLatin[337] = 245 // [ő]  U+0151 => F5
    converter.utf8ToLatin[246] = 246 // [ö]  U+00F6 => F6
    converter.utf8ToLatin[247] = 247 // [÷]  U+00F7 => F7
    converter.utf8ToLatin[345] = 248 // [ř]  U+0159 => F8
    converter.utf8ToLatin[367] = 249 // [ů]  U+016F => F9
    converter.utf8ToLatin[250] = 250 // [ú]  U+00FA => FA
    converter.utf8ToLatin[369] = 251 // [ű]  U+0171 => FB
    converter.utf8ToLatin[252] = 252 // [ü]  U+00FC => FC
    converter.utf8ToLatin[253] = 253 // [ý]  U+00FD => FD
    converter.utf8ToLatin[355] = 254 // [ţ]  U+0163 => FE
    converter.utf8ToLatin[729] = 255 // [˙]  U+02D9 => FF

    return converter
}
