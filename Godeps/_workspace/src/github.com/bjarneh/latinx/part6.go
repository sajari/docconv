// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

// ISO 8859-6

// Generated from http://unicode.org/Public/MAPPINGS/ISO8859/8859-6.TXT

func newISO_8859_6Converter() *Converter {

    converter := new(Converter)
    converter.id = "Arabic/ISO-8859-6"
    converter.latinToUtf8 = make(map[byte][]byte)

    converter.latinToUtf8[160] = []byte(" ") // NO-BREAK SPACE
    converter.latinToUtf8[164] = []byte("¤") // CURRENCY SIGN
    converter.latinToUtf8[172] = []byte("،") // ARABIC COMMA
    converter.latinToUtf8[173] = []byte("­") // SOFT HYPHEN
    converter.latinToUtf8[187] = []byte("؛") // ARABIC SEMICOLON
    converter.latinToUtf8[191] = []byte("؟") // ARABIC QUESTION MARK
    converter.latinToUtf8[193] = []byte("ء") // ARABIC LETTER HAMZA
    converter.latinToUtf8[194] = []byte("آ") // ARABIC LETTER ALEF WITH MADDA ABOVE
    converter.latinToUtf8[195] = []byte("أ") // ARABIC LETTER ALEF WITH HAMZA ABOVE
    converter.latinToUtf8[196] = []byte("ؤ") // ARABIC LETTER WAW WITH HAMZA ABOVE
    converter.latinToUtf8[197] = []byte("إ") // ARABIC LETTER ALEF WITH HAMZA BELOW
    converter.latinToUtf8[198] = []byte("ئ") // ARABIC LETTER YEH WITH HAMZA ABOVE
    converter.latinToUtf8[199] = []byte("ا") // ARABIC LETTER ALEF
    converter.latinToUtf8[200] = []byte("ب") // ARABIC LETTER BEH
    converter.latinToUtf8[201] = []byte("ة") // ARABIC LETTER TEH MARBUTA
    converter.latinToUtf8[202] = []byte("ت") // ARABIC LETTER TEH
    converter.latinToUtf8[203] = []byte("ث") // ARABIC LETTER THEH
    converter.latinToUtf8[204] = []byte("ج") // ARABIC LETTER JEEM
    converter.latinToUtf8[205] = []byte("ح") // ARABIC LETTER HAH
    converter.latinToUtf8[206] = []byte("خ") // ARABIC LETTER KHAH
    converter.latinToUtf8[207] = []byte("د") // ARABIC LETTER DAL
    converter.latinToUtf8[208] = []byte("ذ") // ARABIC LETTER THAL
    converter.latinToUtf8[209] = []byte("ر") // ARABIC LETTER REH
    converter.latinToUtf8[210] = []byte("ز") // ARABIC LETTER ZAIN
    converter.latinToUtf8[211] = []byte("س") // ARABIC LETTER SEEN
    converter.latinToUtf8[212] = []byte("ش") // ARABIC LETTER SHEEN
    converter.latinToUtf8[213] = []byte("ص") // ARABIC LETTER SAD
    converter.latinToUtf8[214] = []byte("ض") // ARABIC LETTER DAD
    converter.latinToUtf8[215] = []byte("ط") // ARABIC LETTER TAH
    converter.latinToUtf8[216] = []byte("ظ") // ARABIC LETTER ZAH
    converter.latinToUtf8[217] = []byte("ع") // ARABIC LETTER AIN
    converter.latinToUtf8[218] = []byte("غ") // ARABIC LETTER GHAIN
    converter.latinToUtf8[224] = []byte("ـ") // ARABIC TATWEEL
    converter.latinToUtf8[225] = []byte("ف") // ARABIC LETTER FEH
    converter.latinToUtf8[226] = []byte("ق") // ARABIC LETTER QAF
    converter.latinToUtf8[227] = []byte("ك") // ARABIC LETTER KAF
    converter.latinToUtf8[228] = []byte("ل") // ARABIC LETTER LAM
    converter.latinToUtf8[229] = []byte("م") // ARABIC LETTER MEEM
    converter.latinToUtf8[230] = []byte("ن") // ARABIC LETTER NOON
    converter.latinToUtf8[231] = []byte("ه") // ARABIC LETTER HEH
    converter.latinToUtf8[232] = []byte("و") // ARABIC LETTER WAW
    converter.latinToUtf8[233] = []byte("ى") // ARABIC LETTER ALEF MAKSURA
    converter.latinToUtf8[234] = []byte("ي") // ARABIC LETTER YEH
    converter.latinToUtf8[235] = []byte("ً") // ARABIC FATHATAN
    converter.latinToUtf8[236] = []byte("ٌ") // ARABIC DAMMATAN
    converter.latinToUtf8[237] = []byte("ٍ") // ARABIC KASRATAN
    converter.latinToUtf8[238] = []byte("َ") // ARABIC FATHA
    converter.latinToUtf8[239] = []byte("ُ") // ARABIC DAMMA
    converter.latinToUtf8[240] = []byte("ِ") // ARABIC KASRA
    converter.latinToUtf8[241] = []byte("ّ") // ARABIC SHADDA
    converter.latinToUtf8[242] = []byte("ْ") // ARABIC SUKUN

    converter.utf8ToLatin = make(map[int]byte)

    converter.utf8ToLatin[160] = 160  // [ ]  U+00A0 => A0
    converter.utf8ToLatin[164] = 164  // [¤]  U+00A4 => A4
    converter.utf8ToLatin[1548] = 172 // [،]  U+060C => AC
    converter.utf8ToLatin[173] = 173  // [­]  U+00AD => AD
    converter.utf8ToLatin[1563] = 187 // [؛]  U+061B => BB
    converter.utf8ToLatin[1567] = 191 // [؟]  U+061F => BF
    converter.utf8ToLatin[1569] = 193 // [ء]  U+0621 => C1
    converter.utf8ToLatin[1570] = 194 // [آ]  U+0622 => C2
    converter.utf8ToLatin[1571] = 195 // [أ]  U+0623 => C3
    converter.utf8ToLatin[1572] = 196 // [ؤ]  U+0624 => C4
    converter.utf8ToLatin[1573] = 197 // [إ]  U+0625 => C5
    converter.utf8ToLatin[1574] = 198 // [ئ]  U+0626 => C6
    converter.utf8ToLatin[1575] = 199 // [ا]  U+0627 => C7
    converter.utf8ToLatin[1576] = 200 // [ب]  U+0628 => C8
    converter.utf8ToLatin[1577] = 201 // [ة]  U+0629 => C9
    converter.utf8ToLatin[1578] = 202 // [ت]  U+062A => CA
    converter.utf8ToLatin[1579] = 203 // [ث]  U+062B => CB
    converter.utf8ToLatin[1580] = 204 // [ج]  U+062C => CC
    converter.utf8ToLatin[1581] = 205 // [ح]  U+062D => CD
    converter.utf8ToLatin[1582] = 206 // [خ]  U+062E => CE
    converter.utf8ToLatin[1583] = 207 // [د]  U+062F => CF
    converter.utf8ToLatin[1584] = 208 // [ذ]  U+0630 => D0
    converter.utf8ToLatin[1585] = 209 // [ر]  U+0631 => D1
    converter.utf8ToLatin[1586] = 210 // [ز]  U+0632 => D2
    converter.utf8ToLatin[1587] = 211 // [س]  U+0633 => D3
    converter.utf8ToLatin[1588] = 212 // [ش]  U+0634 => D4
    converter.utf8ToLatin[1589] = 213 // [ص]  U+0635 => D5
    converter.utf8ToLatin[1590] = 214 // [ض]  U+0636 => D6
    converter.utf8ToLatin[1591] = 215 // [ط]  U+0637 => D7
    converter.utf8ToLatin[1592] = 216 // [ظ]  U+0638 => D8
    converter.utf8ToLatin[1593] = 217 // [ع]  U+0639 => D9
    converter.utf8ToLatin[1594] = 218 // [غ]  U+063A => DA
    converter.utf8ToLatin[1600] = 224 // [ـ]  U+0640 => E0
    converter.utf8ToLatin[1601] = 225 // [ف]  U+0641 => E1
    converter.utf8ToLatin[1602] = 226 // [ق]  U+0642 => E2
    converter.utf8ToLatin[1603] = 227 // [ك]  U+0643 => E3
    converter.utf8ToLatin[1604] = 228 // [ل]  U+0644 => E4
    converter.utf8ToLatin[1605] = 229 // [م]  U+0645 => E5
    converter.utf8ToLatin[1606] = 230 // [ن]  U+0646 => E6
    converter.utf8ToLatin[1607] = 231 // [ه]  U+0647 => E7
    converter.utf8ToLatin[1608] = 232 // [و]  U+0648 => E8
    converter.utf8ToLatin[1609] = 233 // [ى]  U+0649 => E9
    converter.utf8ToLatin[1610] = 234 // [ي]  U+064A => EA
    converter.utf8ToLatin[1611] = 235 // [ً]  U+064B => EB
    converter.utf8ToLatin[1612] = 236 // [ٌ]  U+064C => EC
    converter.utf8ToLatin[1613] = 237 // [ٍ]  U+064D => ED
    converter.utf8ToLatin[1614] = 238 // [َ]  U+064E => EE
    converter.utf8ToLatin[1615] = 239 // [ُ]  U+064F => EF
    converter.utf8ToLatin[1616] = 240 // [ِ]  U+0650 => F0
    converter.utf8ToLatin[1617] = 241 // [ّ]  U+0651 => F1
    converter.utf8ToLatin[1618] = 242 // [ْ]  U+0652 => F2

    return converter
}
