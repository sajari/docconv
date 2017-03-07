// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

// ISO 8859-8

// Generated from http://unicode.org/Public/MAPPINGS/ISO8859/8859-8.TXT

func newISO_8859_8Converter() *Converter {

    converter := new(Converter)
    converter.id = "Hebrew/ISO-8859-8"
    converter.latinToUtf8 = make(map[byte][]byte)

    converter.latinToUtf8[160] = []byte(" ") // NO-BREAK SPACE
    converter.latinToUtf8[162] = []byte("¢") // CENT SIGN
    converter.latinToUtf8[163] = []byte("£") // POUND SIGN
    converter.latinToUtf8[164] = []byte("¤") // CURRENCY SIGN
    converter.latinToUtf8[165] = []byte("¥") // YEN SIGN
    converter.latinToUtf8[166] = []byte("¦") // BROKEN BAR
    converter.latinToUtf8[167] = []byte("§") // SECTION SIGN
    converter.latinToUtf8[168] = []byte("¨") // DIAERESIS
    converter.latinToUtf8[169] = []byte("©") // COPYRIGHT SIGN
    converter.latinToUtf8[170] = []byte("×") // MULTIPLICATION SIGN
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
    converter.latinToUtf8[186] = []byte("÷") // DIVISION SIGN
    converter.latinToUtf8[187] = []byte("»") // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
    converter.latinToUtf8[188] = []byte("¼") // VULGAR FRACTION ONE QUARTER
    converter.latinToUtf8[189] = []byte("½") // VULGAR FRACTION ONE HALF
    converter.latinToUtf8[190] = []byte("¾") // VULGAR FRACTION THREE QUARTERS
    converter.latinToUtf8[223] = []byte("‗") // DOUBLE LOW LINE
    converter.latinToUtf8[224] = []byte("א") // HEBREW LETTER ALEF
    converter.latinToUtf8[225] = []byte("ב") // HEBREW LETTER BET
    converter.latinToUtf8[226] = []byte("ג") // HEBREW LETTER GIMEL
    converter.latinToUtf8[227] = []byte("ד") // HEBREW LETTER DALET
    converter.latinToUtf8[228] = []byte("ה") // HEBREW LETTER HE
    converter.latinToUtf8[229] = []byte("ו") // HEBREW LETTER VAV
    converter.latinToUtf8[230] = []byte("ז") // HEBREW LETTER ZAYIN
    converter.latinToUtf8[231] = []byte("ח") // HEBREW LETTER HET
    converter.latinToUtf8[232] = []byte("ט") // HEBREW LETTER TET
    converter.latinToUtf8[233] = []byte("י") // HEBREW LETTER YOD
    converter.latinToUtf8[234] = []byte("ך") // HEBREW LETTER FINAL KAF
    converter.latinToUtf8[235] = []byte("כ") // HEBREW LETTER KAF
    converter.latinToUtf8[236] = []byte("ל") // HEBREW LETTER LAMED
    converter.latinToUtf8[237] = []byte("ם") // HEBREW LETTER FINAL MEM
    converter.latinToUtf8[238] = []byte("מ") // HEBREW LETTER MEM
    converter.latinToUtf8[239] = []byte("ן") // HEBREW LETTER FINAL NUN
    converter.latinToUtf8[240] = []byte("נ") // HEBREW LETTER NUN
    converter.latinToUtf8[241] = []byte("ס") // HEBREW LETTER SAMEKH
    converter.latinToUtf8[242] = []byte("ע") // HEBREW LETTER AYIN
    converter.latinToUtf8[243] = []byte("ף") // HEBREW LETTER FINAL PE
    converter.latinToUtf8[244] = []byte("פ") // HEBREW LETTER PE
    converter.latinToUtf8[245] = []byte("ץ") // HEBREW LETTER FINAL TSADI
    converter.latinToUtf8[246] = []byte("צ") // HEBREW LETTER TSADI
    converter.latinToUtf8[247] = []byte("ק") // HEBREW LETTER QOF
    converter.latinToUtf8[248] = []byte("ר") // HEBREW LETTER RESH
    converter.latinToUtf8[249] = []byte("ש") // HEBREW LETTER SHIN
    converter.latinToUtf8[250] = []byte("ת") // HEBREW LETTER TAV
    converter.latinToUtf8[253] = []byte("‎") // LEFT-TO-RIGHT MARK
    converter.latinToUtf8[254] = []byte("‏") // RIGHT-TO-LEFT MARK


    converter.utf8ToLatin = make(map[int]byte)

    //[Latin-8/ISO-8859-8]

    converter.utf8ToLatin[160] = 160  // [ ]  U+00A0 => A0
    converter.utf8ToLatin[162] = 162  // [¢]  U+00A2 => A2
    converter.utf8ToLatin[163] = 163  // [£]  U+00A3 => A3
    converter.utf8ToLatin[164] = 164  // [¤]  U+00A4 => A4
    converter.utf8ToLatin[165] = 165  // [¥]  U+00A5 => A5
    converter.utf8ToLatin[166] = 166  // [¦]  U+00A6 => A6
    converter.utf8ToLatin[167] = 167  // [§]  U+00A7 => A7
    converter.utf8ToLatin[168] = 168  // [¨]  U+00A8 => A8
    converter.utf8ToLatin[169] = 169  // [©]  U+00A9 => A9
    converter.utf8ToLatin[215] = 170  // [×]  U+00D7 => AA
    converter.utf8ToLatin[171] = 171  // [«]  U+00AB => AB
    converter.utf8ToLatin[172] = 172  // [¬]  U+00AC => AC
    converter.utf8ToLatin[173] = 173  // [­]  U+00AD => AD
    converter.utf8ToLatin[174] = 174  // [®]  U+00AE => AE
    converter.utf8ToLatin[175] = 175  // [¯]  U+00AF => AF
    converter.utf8ToLatin[176] = 176  // [°]  U+00B0 => B0
    converter.utf8ToLatin[177] = 177  // [±]  U+00B1 => B1
    converter.utf8ToLatin[178] = 178  // [²]  U+00B2 => B2
    converter.utf8ToLatin[179] = 179  // [³]  U+00B3 => B3
    converter.utf8ToLatin[180] = 180  // [´]  U+00B4 => B4
    converter.utf8ToLatin[181] = 181  // [µ]  U+00B5 => B5
    converter.utf8ToLatin[182] = 182  // [¶]  U+00B6 => B6
    converter.utf8ToLatin[183] = 183  // [·]  U+00B7 => B7
    converter.utf8ToLatin[184] = 184  // [¸]  U+00B8 => B8
    converter.utf8ToLatin[185] = 185  // [¹]  U+00B9 => B9
    converter.utf8ToLatin[247] = 186  // [÷]  U+00F7 => BA
    converter.utf8ToLatin[187] = 187  // [»]  U+00BB => BB
    converter.utf8ToLatin[188] = 188  // [¼]  U+00BC => BC
    converter.utf8ToLatin[189] = 189  // [½]  U+00BD => BD
    converter.utf8ToLatin[190] = 190  // [¾]  U+00BE => BE
    converter.utf8ToLatin[8215] = 223 // [‗]  U+2017 => DF
    converter.utf8ToLatin[1488] = 224 // [א]  U+05D0 => E0
    converter.utf8ToLatin[1489] = 225 // [ב]  U+05D1 => E1
    converter.utf8ToLatin[1490] = 226 // [ג]  U+05D2 => E2
    converter.utf8ToLatin[1491] = 227 // [ד]  U+05D3 => E3
    converter.utf8ToLatin[1492] = 228 // [ה]  U+05D4 => E4
    converter.utf8ToLatin[1493] = 229 // [ו]  U+05D5 => E5
    converter.utf8ToLatin[1494] = 230 // [ז]  U+05D6 => E6
    converter.utf8ToLatin[1495] = 231 // [ח]  U+05D7 => E7
    converter.utf8ToLatin[1496] = 232 // [ט]  U+05D8 => E8
    converter.utf8ToLatin[1497] = 233 // [י]  U+05D9 => E9
    converter.utf8ToLatin[1498] = 234 // [ך]  U+05DA => EA
    converter.utf8ToLatin[1499] = 235 // [כ]  U+05DB => EB
    converter.utf8ToLatin[1500] = 236 // [ל]  U+05DC => EC
    converter.utf8ToLatin[1501] = 237 // [ם]  U+05DD => ED
    converter.utf8ToLatin[1502] = 238 // [מ]  U+05DE => EE
    converter.utf8ToLatin[1503] = 239 // [ן]  U+05DF => EF
    converter.utf8ToLatin[1504] = 240 // [נ]  U+05E0 => F0
    converter.utf8ToLatin[1505] = 241 // [ס]  U+05E1 => F1
    converter.utf8ToLatin[1506] = 242 // [ע]  U+05E2 => F2
    converter.utf8ToLatin[1507] = 243 // [ף]  U+05E3 => F3
    converter.utf8ToLatin[1508] = 244 // [פ]  U+05E4 => F4
    converter.utf8ToLatin[1509] = 245 // [ץ]  U+05E5 => F5
    converter.utf8ToLatin[1510] = 246 // [צ]  U+05E6 => F6
    converter.utf8ToLatin[1511] = 247 // [ק]  U+05E7 => F7
    converter.utf8ToLatin[1512] = 248 // [ר]  U+05E8 => F8
    converter.utf8ToLatin[1513] = 249 // [ש]  U+05E9 => F9
    converter.utf8ToLatin[1514] = 250 // [ת]  U+05EA => FA
    converter.utf8ToLatin[8206] = 253 // [‎]  U+200E => FD
    converter.utf8ToLatin[8207] = 254 // [‏]  U+200F => FE

    return converter
}
