// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

// ISO 8859-7

// Generated from http://unicode.org/Public/MAPPINGS/ISO8859/8859-7.TXT

func newISO_8859_7Converter() *Converter {

    converter := new(Converter)
    converter.id = "Greek/ISO-8859-7"
    converter.latinToUtf8 = make(map[byte][]byte)

    converter.latinToUtf8[160] = []byte(" ") // NO-BREAK SPACE
    converter.latinToUtf8[161] = []byte("‘") // LEFT SINGLE QUOTATION MARK
    converter.latinToUtf8[162] = []byte("’") // RIGHT SINGLE QUOTATION MARK
    converter.latinToUtf8[163] = []byte("£") // POUND SIGN
    converter.latinToUtf8[164] = []byte("€") // EURO SIGN
    converter.latinToUtf8[165] = []byte("₯") // DRACHMA SIGN
    converter.latinToUtf8[166] = []byte("¦") // BROKEN BAR
    converter.latinToUtf8[167] = []byte("§") // SECTION SIGN
    converter.latinToUtf8[168] = []byte("¨") // DIAERESIS
    converter.latinToUtf8[169] = []byte("©") // COPYRIGHT SIGN
    converter.latinToUtf8[170] = []byte("ͺ") // GREEK YPOGEGRAMMENI
    converter.latinToUtf8[171] = []byte("«") // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
    converter.latinToUtf8[172] = []byte("¬") // NOT SIGN
    converter.latinToUtf8[173] = []byte("­") // SOFT HYPHEN
    converter.latinToUtf8[175] = []byte("―") // HORIZONTAL BAR
    converter.latinToUtf8[176] = []byte("°") // DEGREE SIGN
    converter.latinToUtf8[177] = []byte("±") // PLUS-MINUS SIGN
    converter.latinToUtf8[178] = []byte("²") // SUPERSCRIPT TWO
    converter.latinToUtf8[179] = []byte("³") // SUPERSCRIPT THREE
    converter.latinToUtf8[180] = []byte("΄") // GREEK TONOS
    converter.latinToUtf8[181] = []byte("΅") // GREEK DIALYTIKA TONOS
    converter.latinToUtf8[182] = []byte("Ά") // GREEK CAPITAL LETTER ALPHA WITH TONOS
    converter.latinToUtf8[183] = []byte("·") // MIDDLE DOT
    converter.latinToUtf8[184] = []byte("Έ") // GREEK CAPITAL LETTER EPSILON WITH TONOS
    converter.latinToUtf8[185] = []byte("Ή") // GREEK CAPITAL LETTER ETA WITH TONOS
    converter.latinToUtf8[186] = []byte("Ί") // GREEK CAPITAL LETTER IOTA WITH TONOS
    converter.latinToUtf8[187] = []byte("»") // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
    converter.latinToUtf8[188] = []byte("Ό") // GREEK CAPITAL LETTER OMICRON WITH TONOS
    converter.latinToUtf8[189] = []byte("½") // VULGAR FRACTION ONE HALF
    converter.latinToUtf8[190] = []byte("Ύ") // GREEK CAPITAL LETTER UPSILON WITH TONOS
    converter.latinToUtf8[191] = []byte("Ώ") // GREEK CAPITAL LETTER OMEGA WITH TONOS
    converter.latinToUtf8[192] = []byte("ΐ") // GREEK SMALL LETTER IOTA WITH DIALYTIKA AND TONOS
    converter.latinToUtf8[193] = []byte("Α") // GREEK CAPITAL LETTER ALPHA
    converter.latinToUtf8[194] = []byte("Β") // GREEK CAPITAL LETTER BETA
    converter.latinToUtf8[195] = []byte("Γ") // GREEK CAPITAL LETTER GAMMA
    converter.latinToUtf8[196] = []byte("Δ") // GREEK CAPITAL LETTER DELTA
    converter.latinToUtf8[197] = []byte("Ε") // GREEK CAPITAL LETTER EPSILON
    converter.latinToUtf8[198] = []byte("Ζ") // GREEK CAPITAL LETTER ZETA
    converter.latinToUtf8[199] = []byte("Η") // GREEK CAPITAL LETTER ETA
    converter.latinToUtf8[200] = []byte("Θ") // GREEK CAPITAL LETTER THETA
    converter.latinToUtf8[201] = []byte("Ι") // GREEK CAPITAL LETTER IOTA
    converter.latinToUtf8[202] = []byte("Κ") // GREEK CAPITAL LETTER KAPPA
    converter.latinToUtf8[203] = []byte("Λ") // GREEK CAPITAL LETTER LAMDA
    converter.latinToUtf8[204] = []byte("Μ") // GREEK CAPITAL LETTER MU
    converter.latinToUtf8[205] = []byte("Ν") // GREEK CAPITAL LETTER NU
    converter.latinToUtf8[206] = []byte("Ξ") // GREEK CAPITAL LETTER XI
    converter.latinToUtf8[207] = []byte("Ο") // GREEK CAPITAL LETTER OMICRON
    converter.latinToUtf8[208] = []byte("Π") // GREEK CAPITAL LETTER PI
    converter.latinToUtf8[209] = []byte("Ρ") // GREEK CAPITAL LETTER RHO
    converter.latinToUtf8[211] = []byte("Σ") // GREEK CAPITAL LETTER SIGMA
    converter.latinToUtf8[212] = []byte("Τ") // GREEK CAPITAL LETTER TAU
    converter.latinToUtf8[213] = []byte("Υ") // GREEK CAPITAL LETTER UPSILON
    converter.latinToUtf8[214] = []byte("Φ") // GREEK CAPITAL LETTER PHI
    converter.latinToUtf8[215] = []byte("Χ") // GREEK CAPITAL LETTER CHI
    converter.latinToUtf8[216] = []byte("Ψ") // GREEK CAPITAL LETTER PSI
    converter.latinToUtf8[217] = []byte("Ω") // GREEK CAPITAL LETTER OMEGA
    converter.latinToUtf8[218] = []byte("Ϊ") // GREEK CAPITAL LETTER IOTA WITH DIALYTIKA
    converter.latinToUtf8[219] = []byte("Ϋ") // GREEK CAPITAL LETTER UPSILON WITH DIALYTIKA
    converter.latinToUtf8[220] = []byte("ά") // GREEK SMALL LETTER ALPHA WITH TONOS
    converter.latinToUtf8[221] = []byte("έ") // GREEK SMALL LETTER EPSILON WITH TONOS
    converter.latinToUtf8[222] = []byte("ή") // GREEK SMALL LETTER ETA WITH TONOS
    converter.latinToUtf8[223] = []byte("ί") // GREEK SMALL LETTER IOTA WITH TONOS
    converter.latinToUtf8[224] = []byte("ΰ") // GREEK SMALL LETTER UPSILON WITH DIALYTIKA AND TONOS
    converter.latinToUtf8[225] = []byte("α") // GREEK SMALL LETTER ALPHA
    converter.latinToUtf8[226] = []byte("β") // GREEK SMALL LETTER BETA
    converter.latinToUtf8[227] = []byte("γ") // GREEK SMALL LETTER GAMMA
    converter.latinToUtf8[228] = []byte("δ") // GREEK SMALL LETTER DELTA
    converter.latinToUtf8[229] = []byte("ε") // GREEK SMALL LETTER EPSILON
    converter.latinToUtf8[230] = []byte("ζ") // GREEK SMALL LETTER ZETA
    converter.latinToUtf8[231] = []byte("η") // GREEK SMALL LETTER ETA
    converter.latinToUtf8[232] = []byte("θ") // GREEK SMALL LETTER THETA
    converter.latinToUtf8[233] = []byte("ι") // GREEK SMALL LETTER IOTA
    converter.latinToUtf8[234] = []byte("κ") // GREEK SMALL LETTER KAPPA
    converter.latinToUtf8[235] = []byte("λ") // GREEK SMALL LETTER LAMDA
    converter.latinToUtf8[236] = []byte("μ") // GREEK SMALL LETTER MU
    converter.latinToUtf8[237] = []byte("ν") // GREEK SMALL LETTER NU
    converter.latinToUtf8[238] = []byte("ξ") // GREEK SMALL LETTER XI
    converter.latinToUtf8[239] = []byte("ο") // GREEK SMALL LETTER OMICRON
    converter.latinToUtf8[240] = []byte("π") // GREEK SMALL LETTER PI
    converter.latinToUtf8[241] = []byte("ρ") // GREEK SMALL LETTER RHO
    converter.latinToUtf8[242] = []byte("ς") // GREEK SMALL LETTER FINAL SIGMA
    converter.latinToUtf8[243] = []byte("σ") // GREEK SMALL LETTER SIGMA
    converter.latinToUtf8[244] = []byte("τ") // GREEK SMALL LETTER TAU
    converter.latinToUtf8[245] = []byte("υ") // GREEK SMALL LETTER UPSILON
    converter.latinToUtf8[246] = []byte("φ") // GREEK SMALL LETTER PHI
    converter.latinToUtf8[247] = []byte("χ") // GREEK SMALL LETTER CHI
    converter.latinToUtf8[248] = []byte("ψ") // GREEK SMALL LETTER PSI
    converter.latinToUtf8[249] = []byte("ω") // GREEK SMALL LETTER OMEGA
    converter.latinToUtf8[250] = []byte("ϊ") // GREEK SMALL LETTER IOTA WITH DIALYTIKA
    converter.latinToUtf8[251] = []byte("ϋ") // GREEK SMALL LETTER UPSILON WITH DIALYTIKA
    converter.latinToUtf8[252] = []byte("ό") // GREEK SMALL LETTER OMICRON WITH TONOS
    converter.latinToUtf8[253] = []byte("ύ") // GREEK SMALL LETTER UPSILON WITH TONOS
    converter.latinToUtf8[254] = []byte("ώ") // GREEK SMALL LETTER OMEGA WITH TONOS


    converter.utf8ToLatin = make(map[int]byte)

    converter.utf8ToLatin[160] = 160  // [ ]  U+00A0 => A0
    converter.utf8ToLatin[8216] = 161 // [‘]  U+2018 => A1
    converter.utf8ToLatin[8217] = 162 // [’]  U+2019 => A2
    converter.utf8ToLatin[163] = 163  // [£]  U+00A3 => A3
    converter.utf8ToLatin[8364] = 164 // [€]  U+20AC => A4
    converter.utf8ToLatin[8367] = 165 // [₯]  U+20AF => A5
    converter.utf8ToLatin[166] = 166  // [¦]  U+00A6 => A6
    converter.utf8ToLatin[167] = 167  // [§]  U+00A7 => A7
    converter.utf8ToLatin[168] = 168  // [¨]  U+00A8 => A8
    converter.utf8ToLatin[169] = 169  // [©]  U+00A9 => A9
    converter.utf8ToLatin[890] = 170  // [ͺ]  U+037A => AA
    converter.utf8ToLatin[171] = 171  // [«]  U+00AB => AB
    converter.utf8ToLatin[172] = 172  // [¬]  U+00AC => AC
    converter.utf8ToLatin[173] = 173  // [­]  U+00AD => AD
    converter.utf8ToLatin[8213] = 175 // [―]  U+2015 => AF
    converter.utf8ToLatin[176] = 176  // [°]  U+00B0 => B0
    converter.utf8ToLatin[177] = 177  // [±]  U+00B1 => B1
    converter.utf8ToLatin[178] = 178  // [²]  U+00B2 => B2
    converter.utf8ToLatin[179] = 179  // [³]  U+00B3 => B3
    converter.utf8ToLatin[900] = 180  // [΄]  U+0384 => B4
    converter.utf8ToLatin[901] = 181  // [΅]  U+0385 => B5
    converter.utf8ToLatin[902] = 182  // [Ά]  U+0386 => B6
    converter.utf8ToLatin[183] = 183  // [·]  U+00B7 => B7
    converter.utf8ToLatin[904] = 184  // [Έ]  U+0388 => B8
    converter.utf8ToLatin[905] = 185  // [Ή]  U+0389 => B9
    converter.utf8ToLatin[906] = 186  // [Ί]  U+038A => BA
    converter.utf8ToLatin[187] = 187  // [»]  U+00BB => BB
    converter.utf8ToLatin[908] = 188  // [Ό]  U+038C => BC
    converter.utf8ToLatin[189] = 189  // [½]  U+00BD => BD
    converter.utf8ToLatin[910] = 190  // [Ύ]  U+038E => BE
    converter.utf8ToLatin[911] = 191  // [Ώ]  U+038F => BF
    converter.utf8ToLatin[912] = 192  // [ΐ]  U+0390 => C0
    converter.utf8ToLatin[913] = 193  // [Α]  U+0391 => C1
    converter.utf8ToLatin[914] = 194  // [Β]  U+0392 => C2
    converter.utf8ToLatin[915] = 195  // [Γ]  U+0393 => C3
    converter.utf8ToLatin[916] = 196  // [Δ]  U+0394 => C4
    converter.utf8ToLatin[917] = 197  // [Ε]  U+0395 => C5
    converter.utf8ToLatin[918] = 198  // [Ζ]  U+0396 => C6
    converter.utf8ToLatin[919] = 199  // [Η]  U+0397 => C7
    converter.utf8ToLatin[920] = 200  // [Θ]  U+0398 => C8
    converter.utf8ToLatin[921] = 201  // [Ι]  U+0399 => C9
    converter.utf8ToLatin[922] = 202  // [Κ]  U+039A => CA
    converter.utf8ToLatin[923] = 203  // [Λ]  U+039B => CB
    converter.utf8ToLatin[924] = 204  // [Μ]  U+039C => CC
    converter.utf8ToLatin[925] = 205  // [Ν]  U+039D => CD
    converter.utf8ToLatin[926] = 206  // [Ξ]  U+039E => CE
    converter.utf8ToLatin[927] = 207  // [Ο]  U+039F => CF
    converter.utf8ToLatin[928] = 208  // [Π]  U+03A0 => D0
    converter.utf8ToLatin[929] = 209  // [Ρ]  U+03A1 => D1
    converter.utf8ToLatin[931] = 211  // [Σ]  U+03A3 => D3
    converter.utf8ToLatin[932] = 212  // [Τ]  U+03A4 => D4
    converter.utf8ToLatin[933] = 213  // [Υ]  U+03A5 => D5
    converter.utf8ToLatin[934] = 214  // [Φ]  U+03A6 => D6
    converter.utf8ToLatin[935] = 215  // [Χ]  U+03A7 => D7
    converter.utf8ToLatin[936] = 216  // [Ψ]  U+03A8 => D8
    converter.utf8ToLatin[937] = 217  // [Ω]  U+03A9 => D9
    converter.utf8ToLatin[938] = 218  // [Ϊ]  U+03AA => DA
    converter.utf8ToLatin[939] = 219  // [Ϋ]  U+03AB => DB
    converter.utf8ToLatin[940] = 220  // [ά]  U+03AC => DC
    converter.utf8ToLatin[941] = 221  // [έ]  U+03AD => DD
    converter.utf8ToLatin[942] = 222  // [ή]  U+03AE => DE
    converter.utf8ToLatin[943] = 223  // [ί]  U+03AF => DF
    converter.utf8ToLatin[944] = 224  // [ΰ]  U+03B0 => E0
    converter.utf8ToLatin[945] = 225  // [α]  U+03B1 => E1
    converter.utf8ToLatin[946] = 226  // [β]  U+03B2 => E2
    converter.utf8ToLatin[947] = 227  // [γ]  U+03B3 => E3
    converter.utf8ToLatin[948] = 228  // [δ]  U+03B4 => E4
    converter.utf8ToLatin[949] = 229  // [ε]  U+03B5 => E5
    converter.utf8ToLatin[950] = 230  // [ζ]  U+03B6 => E6
    converter.utf8ToLatin[951] = 231  // [η]  U+03B7 => E7
    converter.utf8ToLatin[952] = 232  // [θ]  U+03B8 => E8
    converter.utf8ToLatin[953] = 233  // [ι]  U+03B9 => E9
    converter.utf8ToLatin[954] = 234  // [κ]  U+03BA => EA
    converter.utf8ToLatin[955] = 235  // [λ]  U+03BB => EB
    converter.utf8ToLatin[956] = 236  // [μ]  U+03BC => EC
    converter.utf8ToLatin[957] = 237  // [ν]  U+03BD => ED
    converter.utf8ToLatin[958] = 238  // [ξ]  U+03BE => EE
    converter.utf8ToLatin[959] = 239  // [ο]  U+03BF => EF
    converter.utf8ToLatin[960] = 240  // [π]  U+03C0 => F0
    converter.utf8ToLatin[961] = 241  // [ρ]  U+03C1 => F1
    converter.utf8ToLatin[962] = 242  // [ς]  U+03C2 => F2
    converter.utf8ToLatin[963] = 243  // [σ]  U+03C3 => F3
    converter.utf8ToLatin[964] = 244  // [τ]  U+03C4 => F4
    converter.utf8ToLatin[965] = 245  // [υ]  U+03C5 => F5
    converter.utf8ToLatin[966] = 246  // [φ]  U+03C6 => F6
    converter.utf8ToLatin[967] = 247  // [χ]  U+03C7 => F7
    converter.utf8ToLatin[968] = 248  // [ψ]  U+03C8 => F8
    converter.utf8ToLatin[969] = 249  // [ω]  U+03C9 => F9
    converter.utf8ToLatin[970] = 250  // [ϊ]  U+03CA => FA
    converter.utf8ToLatin[971] = 251  // [ϋ]  U+03CB => FB
    converter.utf8ToLatin[972] = 252  // [ό]  U+03CC => FC
    converter.utf8ToLatin[973] = 253  // [ύ]  U+03CD => FD
    converter.utf8ToLatin[974] = 254  // [ώ]  U+03CE => FE

    return converter
}
