// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

// ISO 8859-5

// Generated from http://unicode.org/Public/MAPPINGS/ISO8859/8859-5.TXT

func newISO_8859_5Converter() *Converter {

    converter := new(Converter)
    converter.id = "Cyrillic/ISO-8859-5"
    converter.latinToUtf8 = make(map[byte][]byte)

    converter.latinToUtf8[160] = []byte(" ") // NO-BREAK SPACE
    converter.latinToUtf8[161] = []byte("Ё") // CYRILLIC CAPITAL LETTER IO
    converter.latinToUtf8[162] = []byte("Ђ") // CYRILLIC CAPITAL LETTER DJE
    converter.latinToUtf8[163] = []byte("Ѓ") // CYRILLIC CAPITAL LETTER GJE
    converter.latinToUtf8[164] = []byte("Є") // CYRILLIC CAPITAL LETTER UKRAINIAN IE
    converter.latinToUtf8[165] = []byte("Ѕ") // CYRILLIC CAPITAL LETTER DZE
    converter.latinToUtf8[166] = []byte("І") // CYRILLIC CAPITAL LETTER BYELORUSSIAN-UKRAINIAN I
    converter.latinToUtf8[167] = []byte("Ї") // CYRILLIC CAPITAL LETTER YI
    converter.latinToUtf8[168] = []byte("Ј") // CYRILLIC CAPITAL LETTER JE
    converter.latinToUtf8[169] = []byte("Љ") // CYRILLIC CAPITAL LETTER LJE
    converter.latinToUtf8[170] = []byte("Њ") // CYRILLIC CAPITAL LETTER NJE
    converter.latinToUtf8[171] = []byte("Ћ") // CYRILLIC CAPITAL LETTER TSHE
    converter.latinToUtf8[172] = []byte("Ќ") // CYRILLIC CAPITAL LETTER KJE
    converter.latinToUtf8[173] = []byte("­") // SOFT HYPHEN
    converter.latinToUtf8[174] = []byte("Ў") // CYRILLIC CAPITAL LETTER SHORT U
    converter.latinToUtf8[175] = []byte("Џ") // CYRILLIC CAPITAL LETTER DZHE
    converter.latinToUtf8[176] = []byte("А") // CYRILLIC CAPITAL LETTER A
    converter.latinToUtf8[177] = []byte("Б") // CYRILLIC CAPITAL LETTER BE
    converter.latinToUtf8[178] = []byte("В") // CYRILLIC CAPITAL LETTER VE
    converter.latinToUtf8[179] = []byte("Г") // CYRILLIC CAPITAL LETTER GHE
    converter.latinToUtf8[180] = []byte("Д") // CYRILLIC CAPITAL LETTER DE
    converter.latinToUtf8[181] = []byte("Е") // CYRILLIC CAPITAL LETTER IE
    converter.latinToUtf8[182] = []byte("Ж") // CYRILLIC CAPITAL LETTER ZHE
    converter.latinToUtf8[183] = []byte("З") // CYRILLIC CAPITAL LETTER ZE
    converter.latinToUtf8[184] = []byte("И") // CYRILLIC CAPITAL LETTER I
    converter.latinToUtf8[185] = []byte("Й") // CYRILLIC CAPITAL LETTER SHORT I
    converter.latinToUtf8[186] = []byte("К") // CYRILLIC CAPITAL LETTER KA
    converter.latinToUtf8[187] = []byte("Л") // CYRILLIC CAPITAL LETTER EL
    converter.latinToUtf8[188] = []byte("М") // CYRILLIC CAPITAL LETTER EM
    converter.latinToUtf8[189] = []byte("Н") // CYRILLIC CAPITAL LETTER EN
    converter.latinToUtf8[190] = []byte("О") // CYRILLIC CAPITAL LETTER O
    converter.latinToUtf8[191] = []byte("П") // CYRILLIC CAPITAL LETTER PE
    converter.latinToUtf8[192] = []byte("Р") // CYRILLIC CAPITAL LETTER ER
    converter.latinToUtf8[193] = []byte("С") // CYRILLIC CAPITAL LETTER ES
    converter.latinToUtf8[194] = []byte("Т") // CYRILLIC CAPITAL LETTER TE
    converter.latinToUtf8[195] = []byte("У") // CYRILLIC CAPITAL LETTER U
    converter.latinToUtf8[196] = []byte("Ф") // CYRILLIC CAPITAL LETTER EF
    converter.latinToUtf8[197] = []byte("Х") // CYRILLIC CAPITAL LETTER HA
    converter.latinToUtf8[198] = []byte("Ц") // CYRILLIC CAPITAL LETTER TSE
    converter.latinToUtf8[199] = []byte("Ч") // CYRILLIC CAPITAL LETTER CHE
    converter.latinToUtf8[200] = []byte("Ш") // CYRILLIC CAPITAL LETTER SHA
    converter.latinToUtf8[201] = []byte("Щ") // CYRILLIC CAPITAL LETTER SHCHA
    converter.latinToUtf8[202] = []byte("Ъ") // CYRILLIC CAPITAL LETTER HARD SIGN
    converter.latinToUtf8[203] = []byte("Ы") // CYRILLIC CAPITAL LETTER YERU
    converter.latinToUtf8[204] = []byte("Ь") // CYRILLIC CAPITAL LETTER SOFT SIGN
    converter.latinToUtf8[205] = []byte("Э") // CYRILLIC CAPITAL LETTER E
    converter.latinToUtf8[206] = []byte("Ю") // CYRILLIC CAPITAL LETTER YU
    converter.latinToUtf8[207] = []byte("Я") // CYRILLIC CAPITAL LETTER YA
    converter.latinToUtf8[208] = []byte("а") // CYRILLIC SMALL LETTER A
    converter.latinToUtf8[209] = []byte("б") // CYRILLIC SMALL LETTER BE
    converter.latinToUtf8[210] = []byte("в") // CYRILLIC SMALL LETTER VE
    converter.latinToUtf8[211] = []byte("г") // CYRILLIC SMALL LETTER GHE
    converter.latinToUtf8[212] = []byte("д") // CYRILLIC SMALL LETTER DE
    converter.latinToUtf8[213] = []byte("е") // CYRILLIC SMALL LETTER IE
    converter.latinToUtf8[214] = []byte("ж") // CYRILLIC SMALL LETTER ZHE
    converter.latinToUtf8[215] = []byte("з") // CYRILLIC SMALL LETTER ZE
    converter.latinToUtf8[216] = []byte("и") // CYRILLIC SMALL LETTER I
    converter.latinToUtf8[217] = []byte("й") // CYRILLIC SMALL LETTER SHORT I
    converter.latinToUtf8[218] = []byte("к") // CYRILLIC SMALL LETTER KA
    converter.latinToUtf8[219] = []byte("л") // CYRILLIC SMALL LETTER EL
    converter.latinToUtf8[220] = []byte("м") // CYRILLIC SMALL LETTER EM
    converter.latinToUtf8[221] = []byte("н") // CYRILLIC SMALL LETTER EN
    converter.latinToUtf8[222] = []byte("о") // CYRILLIC SMALL LETTER O
    converter.latinToUtf8[223] = []byte("п") // CYRILLIC SMALL LETTER PE
    converter.latinToUtf8[224] = []byte("р") // CYRILLIC SMALL LETTER ER
    converter.latinToUtf8[225] = []byte("с") // CYRILLIC SMALL LETTER ES
    converter.latinToUtf8[226] = []byte("т") // CYRILLIC SMALL LETTER TE
    converter.latinToUtf8[227] = []byte("у") // CYRILLIC SMALL LETTER U
    converter.latinToUtf8[228] = []byte("ф") // CYRILLIC SMALL LETTER EF
    converter.latinToUtf8[229] = []byte("х") // CYRILLIC SMALL LETTER HA
    converter.latinToUtf8[230] = []byte("ц") // CYRILLIC SMALL LETTER TSE
    converter.latinToUtf8[231] = []byte("ч") // CYRILLIC SMALL LETTER CHE
    converter.latinToUtf8[232] = []byte("ш") // CYRILLIC SMALL LETTER SHA
    converter.latinToUtf8[233] = []byte("щ") // CYRILLIC SMALL LETTER SHCHA
    converter.latinToUtf8[234] = []byte("ъ") // CYRILLIC SMALL LETTER HARD SIGN
    converter.latinToUtf8[235] = []byte("ы") // CYRILLIC SMALL LETTER YERU
    converter.latinToUtf8[236] = []byte("ь") // CYRILLIC SMALL LETTER SOFT SIGN
    converter.latinToUtf8[237] = []byte("э") // CYRILLIC SMALL LETTER E
    converter.latinToUtf8[238] = []byte("ю") // CYRILLIC SMALL LETTER YU
    converter.latinToUtf8[239] = []byte("я") // CYRILLIC SMALL LETTER YA
    converter.latinToUtf8[240] = []byte("№") // NUMERO SIGN
    converter.latinToUtf8[241] = []byte("ё") // CYRILLIC SMALL LETTER IO
    converter.latinToUtf8[242] = []byte("ђ") // CYRILLIC SMALL LETTER DJE
    converter.latinToUtf8[243] = []byte("ѓ") // CYRILLIC SMALL LETTER GJE
    converter.latinToUtf8[244] = []byte("є") // CYRILLIC SMALL LETTER UKRAINIAN IE
    converter.latinToUtf8[245] = []byte("ѕ") // CYRILLIC SMALL LETTER DZE
    converter.latinToUtf8[246] = []byte("і") // CYRILLIC SMALL LETTER BYELORUSSIAN-UKRAINIAN I
    converter.latinToUtf8[247] = []byte("ї") // CYRILLIC SMALL LETTER YI
    converter.latinToUtf8[248] = []byte("ј") // CYRILLIC SMALL LETTER JE
    converter.latinToUtf8[249] = []byte("љ") // CYRILLIC SMALL LETTER LJE
    converter.latinToUtf8[250] = []byte("њ") // CYRILLIC SMALL LETTER NJE
    converter.latinToUtf8[251] = []byte("ћ") // CYRILLIC SMALL LETTER TSHE
    converter.latinToUtf8[252] = []byte("ќ") // CYRILLIC SMALL LETTER KJE
    converter.latinToUtf8[253] = []byte("§") // SECTION SIGN
    converter.latinToUtf8[254] = []byte("ў") // CYRILLIC SMALL LETTER SHORT U
    converter.latinToUtf8[255] = []byte("џ") // CYRILLIC SMALL LETTER DZHE

    converter.utf8ToLatin = make(map[int]byte)

    converter.utf8ToLatin[160] = 160  // [ ]  U+00A0 => A0
    converter.utf8ToLatin[1025] = 161 // [Ё]  U+0401 => A1
    converter.utf8ToLatin[1026] = 162 // [Ђ]  U+0402 => A2
    converter.utf8ToLatin[1027] = 163 // [Ѓ]  U+0403 => A3
    converter.utf8ToLatin[1028] = 164 // [Є]  U+0404 => A4
    converter.utf8ToLatin[1029] = 165 // [Ѕ]  U+0405 => A5
    converter.utf8ToLatin[1030] = 166 // [І]  U+0406 => A6
    converter.utf8ToLatin[1031] = 167 // [Ї]  U+0407 => A7
    converter.utf8ToLatin[1032] = 168 // [Ј]  U+0408 => A8
    converter.utf8ToLatin[1033] = 169 // [Љ]  U+0409 => A9
    converter.utf8ToLatin[1034] = 170 // [Њ]  U+040A => AA
    converter.utf8ToLatin[1035] = 171 // [Ћ]  U+040B => AB
    converter.utf8ToLatin[1036] = 172 // [Ќ]  U+040C => AC
    converter.utf8ToLatin[173] = 173  // [­]  U+00AD => AD
    converter.utf8ToLatin[1038] = 174 // [Ў]  U+040E => AE
    converter.utf8ToLatin[1039] = 175 // [Џ]  U+040F => AF
    converter.utf8ToLatin[1040] = 176 // [А]  U+0410 => B0
    converter.utf8ToLatin[1041] = 177 // [Б]  U+0411 => B1
    converter.utf8ToLatin[1042] = 178 // [В]  U+0412 => B2
    converter.utf8ToLatin[1043] = 179 // [Г]  U+0413 => B3
    converter.utf8ToLatin[1044] = 180 // [Д]  U+0414 => B4
    converter.utf8ToLatin[1045] = 181 // [Е]  U+0415 => B5
    converter.utf8ToLatin[1046] = 182 // [Ж]  U+0416 => B6
    converter.utf8ToLatin[1047] = 183 // [З]  U+0417 => B7
    converter.utf8ToLatin[1048] = 184 // [И]  U+0418 => B8
    converter.utf8ToLatin[1049] = 185 // [Й]  U+0419 => B9
    converter.utf8ToLatin[1050] = 186 // [К]  U+041A => BA
    converter.utf8ToLatin[1051] = 187 // [Л]  U+041B => BB
    converter.utf8ToLatin[1052] = 188 // [М]  U+041C => BC
    converter.utf8ToLatin[1053] = 189 // [Н]  U+041D => BD
    converter.utf8ToLatin[1054] = 190 // [О]  U+041E => BE
    converter.utf8ToLatin[1055] = 191 // [П]  U+041F => BF
    converter.utf8ToLatin[1056] = 192 // [Р]  U+0420 => C0
    converter.utf8ToLatin[1057] = 193 // [С]  U+0421 => C1
    converter.utf8ToLatin[1058] = 194 // [Т]  U+0422 => C2
    converter.utf8ToLatin[1059] = 195 // [У]  U+0423 => C3
    converter.utf8ToLatin[1060] = 196 // [Ф]  U+0424 => C4
    converter.utf8ToLatin[1061] = 197 // [Х]  U+0425 => C5
    converter.utf8ToLatin[1062] = 198 // [Ц]  U+0426 => C6
    converter.utf8ToLatin[1063] = 199 // [Ч]  U+0427 => C7
    converter.utf8ToLatin[1064] = 200 // [Ш]  U+0428 => C8
    converter.utf8ToLatin[1065] = 201 // [Щ]  U+0429 => C9
    converter.utf8ToLatin[1066] = 202 // [Ъ]  U+042A => CA
    converter.utf8ToLatin[1067] = 203 // [Ы]  U+042B => CB
    converter.utf8ToLatin[1068] = 204 // [Ь]  U+042C => CC
    converter.utf8ToLatin[1069] = 205 // [Э]  U+042D => CD
    converter.utf8ToLatin[1070] = 206 // [Ю]  U+042E => CE
    converter.utf8ToLatin[1071] = 207 // [Я]  U+042F => CF
    converter.utf8ToLatin[1072] = 208 // [а]  U+0430 => D0
    converter.utf8ToLatin[1073] = 209 // [б]  U+0431 => D1
    converter.utf8ToLatin[1074] = 210 // [в]  U+0432 => D2
    converter.utf8ToLatin[1075] = 211 // [г]  U+0433 => D3
    converter.utf8ToLatin[1076] = 212 // [д]  U+0434 => D4
    converter.utf8ToLatin[1077] = 213 // [е]  U+0435 => D5
    converter.utf8ToLatin[1078] = 214 // [ж]  U+0436 => D6
    converter.utf8ToLatin[1079] = 215 // [з]  U+0437 => D7
    converter.utf8ToLatin[1080] = 216 // [и]  U+0438 => D8
    converter.utf8ToLatin[1081] = 217 // [й]  U+0439 => D9
    converter.utf8ToLatin[1082] = 218 // [к]  U+043A => DA
    converter.utf8ToLatin[1083] = 219 // [л]  U+043B => DB
    converter.utf8ToLatin[1084] = 220 // [м]  U+043C => DC
    converter.utf8ToLatin[1085] = 221 // [н]  U+043D => DD
    converter.utf8ToLatin[1086] = 222 // [о]  U+043E => DE
    converter.utf8ToLatin[1087] = 223 // [п]  U+043F => DF
    converter.utf8ToLatin[1088] = 224 // [р]  U+0440 => E0
    converter.utf8ToLatin[1089] = 225 // [с]  U+0441 => E1
    converter.utf8ToLatin[1090] = 226 // [т]  U+0442 => E2
    converter.utf8ToLatin[1091] = 227 // [у]  U+0443 => E3
    converter.utf8ToLatin[1092] = 228 // [ф]  U+0444 => E4
    converter.utf8ToLatin[1093] = 229 // [х]  U+0445 => E5
    converter.utf8ToLatin[1094] = 230 // [ц]  U+0446 => E6
    converter.utf8ToLatin[1095] = 231 // [ч]  U+0447 => E7
    converter.utf8ToLatin[1096] = 232 // [ш]  U+0448 => E8
    converter.utf8ToLatin[1097] = 233 // [щ]  U+0449 => E9
    converter.utf8ToLatin[1098] = 234 // [ъ]  U+044A => EA
    converter.utf8ToLatin[1099] = 235 // [ы]  U+044B => EB
    converter.utf8ToLatin[1100] = 236 // [ь]  U+044C => EC
    converter.utf8ToLatin[1101] = 237 // [э]  U+044D => ED
    converter.utf8ToLatin[1102] = 238 // [ю]  U+044E => EE
    converter.utf8ToLatin[1103] = 239 // [я]  U+044F => EF
    converter.utf8ToLatin[8470] = 240 // [№]  U+2116 => F0
    converter.utf8ToLatin[1105] = 241 // [ё]  U+0451 => F1
    converter.utf8ToLatin[1106] = 242 // [ђ]  U+0452 => F2
    converter.utf8ToLatin[1107] = 243 // [ѓ]  U+0453 => F3
    converter.utf8ToLatin[1108] = 244 // [є]  U+0454 => F4
    converter.utf8ToLatin[1109] = 245 // [ѕ]  U+0455 => F5
    converter.utf8ToLatin[1110] = 246 // [і]  U+0456 => F6
    converter.utf8ToLatin[1111] = 247 // [ї]  U+0457 => F7
    converter.utf8ToLatin[1112] = 248 // [ј]  U+0458 => F8
    converter.utf8ToLatin[1113] = 249 // [љ]  U+0459 => F9
    converter.utf8ToLatin[1114] = 250 // [њ]  U+045A => FA
    converter.utf8ToLatin[1115] = 251 // [ћ]  U+045B => FB
    converter.utf8ToLatin[1116] = 252 // [ќ]  U+045C => FC
    converter.utf8ToLatin[167] = 253  // [§]  U+00A7 => FD
    converter.utf8ToLatin[1118] = 254 // [ў]  U+045E => FE
    converter.utf8ToLatin[1119] = 255 // [џ]  U+045F => FF

    return converter
}
