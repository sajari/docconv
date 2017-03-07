// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

// ISO 8859-11

// Generated from http://unicode.org/Public/MAPPINGS/ISO8859/8859-11.TXT

func newISO_8859_11Converter() *Converter {

    converter := new(Converter)
    converter.id = "Thai/ISO-8859-11"
    converter.latinToUtf8 = make(map[byte][]byte)

    converter.latinToUtf8[160] = []byte(" ") // NO-BREAK SPACE
    converter.latinToUtf8[161] = []byte("ก") // THAI CHARACTER KO KAI
    converter.latinToUtf8[162] = []byte("ข") // THAI CHARACTER KHO KHAI
    converter.latinToUtf8[163] = []byte("ฃ") // THAI CHARACTER KHO KHUAT
    converter.latinToUtf8[164] = []byte("ค") // THAI CHARACTER KHO KHWAI
    converter.latinToUtf8[165] = []byte("ฅ") // THAI CHARACTER KHO KHON
    converter.latinToUtf8[166] = []byte("ฆ") // THAI CHARACTER KHO RAKHANG
    converter.latinToUtf8[167] = []byte("ง") // THAI CHARACTER NGO NGU
    converter.latinToUtf8[168] = []byte("จ") // THAI CHARACTER CHO CHAN
    converter.latinToUtf8[169] = []byte("ฉ") // THAI CHARACTER CHO CHING
    converter.latinToUtf8[170] = []byte("ช") // THAI CHARACTER CHO CHANG
    converter.latinToUtf8[171] = []byte("ซ") // THAI CHARACTER SO SO
    converter.latinToUtf8[172] = []byte("ฌ") // THAI CHARACTER CHO CHOE
    converter.latinToUtf8[173] = []byte("ญ") // THAI CHARACTER YO YING
    converter.latinToUtf8[174] = []byte("ฎ") // THAI CHARACTER DO CHADA
    converter.latinToUtf8[175] = []byte("ฏ") // THAI CHARACTER TO PATAK
    converter.latinToUtf8[176] = []byte("ฐ") // THAI CHARACTER THO THAN
    converter.latinToUtf8[177] = []byte("ฑ") // THAI CHARACTER THO NANGMONTHO
    converter.latinToUtf8[178] = []byte("ฒ") // THAI CHARACTER THO PHUTHAO
    converter.latinToUtf8[179] = []byte("ณ") // THAI CHARACTER NO NEN
    converter.latinToUtf8[180] = []byte("ด") // THAI CHARACTER DO DEK
    converter.latinToUtf8[181] = []byte("ต") // THAI CHARACTER TO TAO
    converter.latinToUtf8[182] = []byte("ถ") // THAI CHARACTER THO THUNG
    converter.latinToUtf8[183] = []byte("ท") // THAI CHARACTER THO THAHAN
    converter.latinToUtf8[184] = []byte("ธ") // THAI CHARACTER THO THONG
    converter.latinToUtf8[185] = []byte("น") // THAI CHARACTER NO NU
    converter.latinToUtf8[186] = []byte("บ") // THAI CHARACTER BO BAIMAI
    converter.latinToUtf8[187] = []byte("ป") // THAI CHARACTER PO PLA
    converter.latinToUtf8[188] = []byte("ผ") // THAI CHARACTER PHO PHUNG
    converter.latinToUtf8[189] = []byte("ฝ") // THAI CHARACTER FO FA
    converter.latinToUtf8[190] = []byte("พ") // THAI CHARACTER PHO PHAN
    converter.latinToUtf8[191] = []byte("ฟ") // THAI CHARACTER FO FAN
    converter.latinToUtf8[192] = []byte("ภ") // THAI CHARACTER PHO SAMPHAO
    converter.latinToUtf8[193] = []byte("ม") // THAI CHARACTER MO MA
    converter.latinToUtf8[194] = []byte("ย") // THAI CHARACTER YO YAK
    converter.latinToUtf8[195] = []byte("ร") // THAI CHARACTER RO RUA
    converter.latinToUtf8[196] = []byte("ฤ") // THAI CHARACTER RU
    converter.latinToUtf8[197] = []byte("ล") // THAI CHARACTER LO LING
    converter.latinToUtf8[198] = []byte("ฦ") // THAI CHARACTER LU
    converter.latinToUtf8[199] = []byte("ว") // THAI CHARACTER WO WAEN
    converter.latinToUtf8[200] = []byte("ศ") // THAI CHARACTER SO SALA
    converter.latinToUtf8[201] = []byte("ษ") // THAI CHARACTER SO RUSI
    converter.latinToUtf8[202] = []byte("ส") // THAI CHARACTER SO SUA
    converter.latinToUtf8[203] = []byte("ห") // THAI CHARACTER HO HIP
    converter.latinToUtf8[204] = []byte("ฬ") // THAI CHARACTER LO CHULA
    converter.latinToUtf8[205] = []byte("อ") // THAI CHARACTER O ANG
    converter.latinToUtf8[206] = []byte("ฮ") // THAI CHARACTER HO NOKHUK
    converter.latinToUtf8[207] = []byte("ฯ") // THAI CHARACTER PAIYANNOI
    converter.latinToUtf8[208] = []byte("ะ") // THAI CHARACTER SARA A
    converter.latinToUtf8[209] = []byte("ั") // THAI CHARACTER MAI HAN-AKAT
    converter.latinToUtf8[210] = []byte("า") // THAI CHARACTER SARA AA
    converter.latinToUtf8[211] = []byte("ำ") // THAI CHARACTER SARA AM
    converter.latinToUtf8[212] = []byte("ิ") // THAI CHARACTER SARA I
    converter.latinToUtf8[213] = []byte("ี") // THAI CHARACTER SARA II
    converter.latinToUtf8[214] = []byte("ึ") // THAI CHARACTER SARA UE
    converter.latinToUtf8[215] = []byte("ื") // THAI CHARACTER SARA UEE
    converter.latinToUtf8[216] = []byte("ุ") // THAI CHARACTER SARA U
    converter.latinToUtf8[217] = []byte("ู") // THAI CHARACTER SARA UU
    converter.latinToUtf8[218] = []byte("ฺ") // THAI CHARACTER PHINTHU
    converter.latinToUtf8[223] = []byte("฿") // THAI CURRENCY SYMBOL BAHT
    converter.latinToUtf8[224] = []byte("เ") // THAI CHARACTER SARA E
    converter.latinToUtf8[225] = []byte("แ") // THAI CHARACTER SARA AE
    converter.latinToUtf8[226] = []byte("โ") // THAI CHARACTER SARA O
    converter.latinToUtf8[227] = []byte("ใ") // THAI CHARACTER SARA AI MAIMUAN
    converter.latinToUtf8[228] = []byte("ไ") // THAI CHARACTER SARA AI MAIMALAI
    converter.latinToUtf8[229] = []byte("ๅ") // THAI CHARACTER LAKKHANGYAO
    converter.latinToUtf8[230] = []byte("ๆ") // THAI CHARACTER MAIYAMOK
    converter.latinToUtf8[231] = []byte("็") // THAI CHARACTER MAITAIKHU
    converter.latinToUtf8[232] = []byte("่") // THAI CHARACTER MAI EK
    converter.latinToUtf8[233] = []byte("้") // THAI CHARACTER MAI THO
    converter.latinToUtf8[234] = []byte("๊") // THAI CHARACTER MAI TRI
    converter.latinToUtf8[235] = []byte("๋") // THAI CHARACTER MAI CHATTAWA
    converter.latinToUtf8[236] = []byte("์") // THAI CHARACTER THANTHAKHAT
    converter.latinToUtf8[237] = []byte("ํ") // THAI CHARACTER NIKHAHIT
    converter.latinToUtf8[238] = []byte("๎") // THAI CHARACTER YAMAKKAN
    converter.latinToUtf8[239] = []byte("๏") // THAI CHARACTER FONGMAN
    converter.latinToUtf8[240] = []byte("๐") // THAI DIGIT ZERO
    converter.latinToUtf8[241] = []byte("๑") // THAI DIGIT ONE
    converter.latinToUtf8[242] = []byte("๒") // THAI DIGIT TWO
    converter.latinToUtf8[243] = []byte("๓") // THAI DIGIT THREE
    converter.latinToUtf8[244] = []byte("๔") // THAI DIGIT FOUR
    converter.latinToUtf8[245] = []byte("๕") // THAI DIGIT FIVE
    converter.latinToUtf8[246] = []byte("๖") // THAI DIGIT SIX
    converter.latinToUtf8[247] = []byte("๗") // THAI DIGIT SEVEN
    converter.latinToUtf8[248] = []byte("๘") // THAI DIGIT EIGHT
    converter.latinToUtf8[249] = []byte("๙") // THAI DIGIT NINE
    converter.latinToUtf8[250] = []byte("๚") // THAI CHARACTER ANGKHANKHU
    converter.latinToUtf8[251] = []byte("๛") // THAI CHARACTER KHOMUT

    converter.utf8ToLatin = make(map[int]byte)

    converter.utf8ToLatin[160] = 160  // [ ]  U+00A0 => A0
    converter.utf8ToLatin[3585] = 161 // [ก]  U+0E01 => A1
    converter.utf8ToLatin[3586] = 162 // [ข]  U+0E02 => A2
    converter.utf8ToLatin[3587] = 163 // [ฃ]  U+0E03 => A3
    converter.utf8ToLatin[3588] = 164 // [ค]  U+0E04 => A4
    converter.utf8ToLatin[3589] = 165 // [ฅ]  U+0E05 => A5
    converter.utf8ToLatin[3590] = 166 // [ฆ]  U+0E06 => A6
    converter.utf8ToLatin[3591] = 167 // [ง]  U+0E07 => A7
    converter.utf8ToLatin[3592] = 168 // [จ]  U+0E08 => A8
    converter.utf8ToLatin[3593] = 169 // [ฉ]  U+0E09 => A9
    converter.utf8ToLatin[3594] = 170 // [ช]  U+0E0A => AA
    converter.utf8ToLatin[3595] = 171 // [ซ]  U+0E0B => AB
    converter.utf8ToLatin[3596] = 172 // [ฌ]  U+0E0C => AC
    converter.utf8ToLatin[3597] = 173 // [ญ]  U+0E0D => AD
    converter.utf8ToLatin[3598] = 174 // [ฎ]  U+0E0E => AE
    converter.utf8ToLatin[3599] = 175 // [ฏ]  U+0E0F => AF
    converter.utf8ToLatin[3600] = 176 // [ฐ]  U+0E10 => B0
    converter.utf8ToLatin[3601] = 177 // [ฑ]  U+0E11 => B1
    converter.utf8ToLatin[3602] = 178 // [ฒ]  U+0E12 => B2
    converter.utf8ToLatin[3603] = 179 // [ณ]  U+0E13 => B3
    converter.utf8ToLatin[3604] = 180 // [ด]  U+0E14 => B4
    converter.utf8ToLatin[3605] = 181 // [ต]  U+0E15 => B5
    converter.utf8ToLatin[3606] = 182 // [ถ]  U+0E16 => B6
    converter.utf8ToLatin[3607] = 183 // [ท]  U+0E17 => B7
    converter.utf8ToLatin[3608] = 184 // [ธ]  U+0E18 => B8
    converter.utf8ToLatin[3609] = 185 // [น]  U+0E19 => B9
    converter.utf8ToLatin[3610] = 186 // [บ]  U+0E1A => BA
    converter.utf8ToLatin[3611] = 187 // [ป]  U+0E1B => BB
    converter.utf8ToLatin[3612] = 188 // [ผ]  U+0E1C => BC
    converter.utf8ToLatin[3613] = 189 // [ฝ]  U+0E1D => BD
    converter.utf8ToLatin[3614] = 190 // [พ]  U+0E1E => BE
    converter.utf8ToLatin[3615] = 191 // [ฟ]  U+0E1F => BF
    converter.utf8ToLatin[3616] = 192 // [ภ]  U+0E20 => C0
    converter.utf8ToLatin[3617] = 193 // [ม]  U+0E21 => C1
    converter.utf8ToLatin[3618] = 194 // [ย]  U+0E22 => C2
    converter.utf8ToLatin[3619] = 195 // [ร]  U+0E23 => C3
    converter.utf8ToLatin[3620] = 196 // [ฤ]  U+0E24 => C4
    converter.utf8ToLatin[3621] = 197 // [ล]  U+0E25 => C5
    converter.utf8ToLatin[3622] = 198 // [ฦ]  U+0E26 => C6
    converter.utf8ToLatin[3623] = 199 // [ว]  U+0E27 => C7
    converter.utf8ToLatin[3624] = 200 // [ศ]  U+0E28 => C8
    converter.utf8ToLatin[3625] = 201 // [ษ]  U+0E29 => C9
    converter.utf8ToLatin[3626] = 202 // [ส]  U+0E2A => CA
    converter.utf8ToLatin[3627] = 203 // [ห]  U+0E2B => CB
    converter.utf8ToLatin[3628] = 204 // [ฬ]  U+0E2C => CC
    converter.utf8ToLatin[3629] = 205 // [อ]  U+0E2D => CD
    converter.utf8ToLatin[3630] = 206 // [ฮ]  U+0E2E => CE
    converter.utf8ToLatin[3631] = 207 // [ฯ]  U+0E2F => CF
    converter.utf8ToLatin[3632] = 208 // [ะ]  U+0E30 => D0
    converter.utf8ToLatin[3633] = 209 // [ั]  U+0E31 => D1
    converter.utf8ToLatin[3634] = 210 // [า]  U+0E32 => D2
    converter.utf8ToLatin[3635] = 211 // [ำ]  U+0E33 => D3
    converter.utf8ToLatin[3636] = 212 // [ิ]  U+0E34 => D4
    converter.utf8ToLatin[3637] = 213 // [ี]  U+0E35 => D5
    converter.utf8ToLatin[3638] = 214 // [ึ]  U+0E36 => D6
    converter.utf8ToLatin[3639] = 215 // [ื]  U+0E37 => D7
    converter.utf8ToLatin[3640] = 216 // [ุ]  U+0E38 => D8
    converter.utf8ToLatin[3641] = 217 // [ู]  U+0E39 => D9
    converter.utf8ToLatin[3642] = 218 // [ฺ]  U+0E3A => DA
    converter.utf8ToLatin[3647] = 223 // [฿]  U+0E3F => DF
    converter.utf8ToLatin[3648] = 224 // [เ]  U+0E40 => E0
    converter.utf8ToLatin[3649] = 225 // [แ]  U+0E41 => E1
    converter.utf8ToLatin[3650] = 226 // [โ]  U+0E42 => E2
    converter.utf8ToLatin[3651] = 227 // [ใ]  U+0E43 => E3
    converter.utf8ToLatin[3652] = 228 // [ไ]  U+0E44 => E4
    converter.utf8ToLatin[3653] = 229 // [ๅ]  U+0E45 => E5
    converter.utf8ToLatin[3654] = 230 // [ๆ]  U+0E46 => E6
    converter.utf8ToLatin[3655] = 231 // [็]  U+0E47 => E7
    converter.utf8ToLatin[3656] = 232 // [่]  U+0E48 => E8
    converter.utf8ToLatin[3657] = 233 // [้]  U+0E49 => E9
    converter.utf8ToLatin[3658] = 234 // [๊]  U+0E4A => EA
    converter.utf8ToLatin[3659] = 235 // [๋]  U+0E4B => EB
    converter.utf8ToLatin[3660] = 236 // [์]  U+0E4C => EC
    converter.utf8ToLatin[3661] = 237 // [ํ]  U+0E4D => ED
    converter.utf8ToLatin[3662] = 238 // [๎]  U+0E4E => EE
    converter.utf8ToLatin[3663] = 239 // [๏]  U+0E4F => EF
    converter.utf8ToLatin[3664] = 240 // [๐]  U+0E50 => F0
    converter.utf8ToLatin[3665] = 241 // [๑]  U+0E51 => F1
    converter.utf8ToLatin[3666] = 242 // [๒]  U+0E52 => F2
    converter.utf8ToLatin[3667] = 243 // [๓]  U+0E53 => F3
    converter.utf8ToLatin[3668] = 244 // [๔]  U+0E54 => F4
    converter.utf8ToLatin[3669] = 245 // [๕]  U+0E55 => F5
    converter.utf8ToLatin[3670] = 246 // [๖]  U+0E56 => F6
    converter.utf8ToLatin[3671] = 247 // [๗]  U+0E57 => F7
    converter.utf8ToLatin[3672] = 248 // [๘]  U+0E58 => F8
    converter.utf8ToLatin[3673] = 249 // [๙]  U+0E59 => F9
    converter.utf8ToLatin[3674] = 250 // [๚]  U+0E5A => FA
    converter.utf8ToLatin[3675] = 251 // [๛]  U+0E5B => FB

    return converter
}
