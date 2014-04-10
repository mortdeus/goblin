
//line cc.y:2
package cc
import __yyfmt__ "fmt"
//line cc.y:2
		
//line cc.y:5
type yySymType struct{
	yys int
	node *Node
	sym *Sym
	type1 *Type
	tycl struct{
		t *Type
		c byte
	}
	tyty struct{
		t1,
		t2,
		t3 *Type
		c byte
	}
	sval struct{
		s string
		l int
	}
	lval int32
	dval double
	vval int64
}

const LPE = 57346
const LME = 57347
const LMLE = 57348
const LDVE = 57349
const LMDE = 57350
const LRSHE = 57351
const LLSHE = 57352
const LANDE = 57353
const LXORE = 57354
const LORE = 57355
const LOROR = 57356
const LANDAND = 57357
const LEQ = 57358
const LNE = 57359
const LLE = 57360
const LGE = 57361
const LLSH = 57362
const LRSH = 57363
const LMM = 57364
const LPP = 57365
const LMG = 57366
const LNAME = 57367
const LTYPE = 57368
const LFCONST = 57369
const LDCONST = 57370
const LCONST = 57371
const LLCONST = 57372
const LUCONST = 57373
const LULCONST = 57374
const LVLCONST = 57375
const LUVLCONST = 57376
const LSTRING = 57377
const LLSTRING = 57378
const LAUTO = 57379
const LBREAK = 57380
const LCASE = 57381
const LCHAR = 57382
const LCONTINUE = 57383
const LDEFAULT = 57384
const LDO = 57385
const LDOUBLE = 57386
const LELSE = 57387
const LEXTERN = 57388
const LFLOAT = 57389
const LFOR = 57390
const LGOTO = 57391
const LIF = 57392
const LINT = 57393
const LLONG = 57394
const LPREFETCH = 57395
const LREGISTER = 57396
const LRETURN = 57397
const LSHORT = 57398
const LSIZEOF = 57399
const LUSED = 57400
const LSTATIC = 57401
const LSTRUCT = 57402
const LSWITCH = 57403
const LTYPEDEF = 57404
const LTYPESTR = 57405
const LUNION = 57406
const LUNSIGNED = 57407
const LWHILE = 57408
const LVOID = 57409
const LENUM = 57410
const LSIGNED = 57411
const LCONSTNT = 57412
const LVOLATILE = 57413
const LSET = 57414
const LSIGNOF = 57415
const LRESTRICT = 57416
const LINLINE = 57417

var yyToknames = []string{
	" ;",
	" ,",
	" =",
	"LPE",
	"LME",
	"LMLE",
	"LDVE",
	"LMDE",
	"LRSHE",
	"LLSHE",
	"LANDE",
	"LXORE",
	"LORE",
	" ?",
	" :",
	"LOROR",
	"LANDAND",
	" |",
	" ^",
	" &",
	"LEQ",
	"LNE",
	" <",
	" >",
	"LLE",
	"LGE",
	"LLSH",
	"LRSH",
	" +",
	" -",
	" *",
	" /",
	" %",
	"LMM",
	"LPP",
	"LMG",
	" .",
	" [",
	" (",
	"LNAME",
	"LTYPE",
	"LFCONST",
	"LDCONST",
	"LCONST",
	"LLCONST",
	"LUCONST",
	"LULCONST",
	"LVLCONST",
	"LUVLCONST",
	"LSTRING",
	"LLSTRING",
	"LAUTO",
	"LBREAK",
	"LCASE",
	"LCHAR",
	"LCONTINUE",
	"LDEFAULT",
	"LDO",
	"LDOUBLE",
	"LELSE",
	"LEXTERN",
	"LFLOAT",
	"LFOR",
	"LGOTO",
	"LIF",
	"LINT",
	"LLONG",
	"LPREFETCH",
	"LREGISTER",
	"LRETURN",
	"LSHORT",
	"LSIZEOF",
	"LUSED",
	"LSTATIC",
	"LSTRUCT",
	"LSWITCH",
	"LTYPEDEF",
	"LTYPESTR",
	"LUNION",
	"LUNSIGNED",
	"LWHILE",
	"LVOID",
	"LENUM",
	"LSIGNED",
	"LCONSTNT",
	"LVOLATILE",
	"LSET",
	"LSIGNOF",
	"LRESTRICT",
	"LINLINE",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line cc.y:1199


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 182,
	-1, 37,
	4, 8,
	5, 8,
	6, 9,
	-2, 5,
	-1, 54,
	96, 195,
	-2, 194,
	-1, 57,
	96, 199,
	-2, 198,
	-1, 59,
	96, 203,
	-2, 202,
	-1, 78,
	6, 9,
	-2, 8,
	-1, 276,
	4, 100,
	66, 88,
	96, 84,
	-2, 0,
	-1, 312,
	66, 88,
	96, 84,
	-2, 100,
	-1, 318,
	4, 100,
	66, 88,
	96, 84,
	-2, 0,
	-1, 348,
	6, 21,
	-2, 20,
	-1, 387,
	4, 100,
	66, 88,
	96, 84,
	-2, 0,
	-1, 391,
	4, 100,
	66, 88,
	96, 84,
	-2, 0,
	-1, 393,
	4, 100,
	66, 88,
	96, 84,
	-2, 0,
	-1, 407,
	4, 100,
	66, 88,
	96, 84,
	-2, 0,
	-1, 414,
	4, 100,
	66, 88,
	96, 84,
	-2, 0,
}

const yyNprod = 247
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1216

var yyAct = []int{

	177, 308, 313, 347, 211, 209, 87, 4, 5, 42,
	311, 205, 89, 328, 258, 327, 268, 259, 81, 54,
	57, 59, 91, 23, 40, 67, 207, 267, 212, 126,
	48, 48, 92, 138, 135, 203, 203, 137, 88, 384,
	142, 140, 43, 44, 138, 82, 270, 105, 43, 44,
	37, 142, 140, 43, 44, 43, 44, 254, 279, 206,
	125, 56, 336, 334, 142, 255, 290, 90, 144, 48,
	414, 396, 395, 394, 48, 342, 48, 70, 341, 335,
	295, 131, 292, 289, 133, 130, 5, 129, 68, 120,
	376, 119, 355, 219, 119, 48, 407, 197, 25, 26,
	256, 60, 27, 196, 256, 127, 256, 176, 56, 78,
	43, 44, 84, 83, 118, 256, 392, 219, 184, 185,
	186, 187, 188, 189, 190, 191, 374, 305, 202, 256,
	39, 136, 256, 131, 36, 365, 192, 194, 41, 43,
	44, 25, 26, 90, 364, 27, 223, 224, 225, 226,
	227, 228, 229, 230, 231, 232, 233, 234, 235, 236,
	237, 238, 239, 240, 216, 242, 243, 244, 245, 246,
	247, 248, 249, 250, 251, 252, 208, 241, 220, 260,
	215, 221, 83, 363, 25, 26, 123, 68, 27, 409,
	262, 263, 254, 261, 297, 393, 359, 77, 356, 142,
	255, 55, 218, 217, 391, 275, 175, 176, 45, 176,
	253, 131, 58, 354, 90, 280, 50, 222, 387, 90,
	138, 257, 66, 65, 285, 271, 284, 142, 140, 43,
	44, 273, 143, 274, 28, 198, 287, 13, 368, 281,
	119, 20, 256, 30, 19, 367, 25, 26, 15, 16,
	27, 33, 46, 14, 286, 291, 29, 288, 303, 31,
	32, 71, 18, 118, 21, 83, 17, 25, 26, 39,
	294, 27, 34, 90, 121, 386, 124, 41, 43, 44,
	293, 283, 254, 5, 309, 304, 370, 371, 24, 142,
	255, 271, 337, 220, 300, 49, 49, 260, 69, 298,
	299, 204, 7, 69, 90, 332, 264, 371, 265, 47,
	47, 52, 346, 80, 340, 338, 345, 119, 256, 344,
	357, 333, 358, 351, 69, 208, 353, 271, 350, 6,
	366, 286, 302, 362, 49, 145, 146, 147, 51, 49,
	131, 49, 39, 277, 278, 369, 35, 412, 47, 22,
	41, 43, 44, 47, 296, 47, 348, 282, 53, 134,
	49, 61, 62, 411, 260, 260, 260, 406, 373, 405,
	375, 404, 399, 378, 47, 385, 39, 389, 380, 381,
	382, 5, 390, 379, 41, 43, 44, 131, 377, 398,
	361, 397, 360, 401, 400, 403, 272, 352, 310, 349,
	76, 343, 408, 301, 201, 75, 402, 74, 72, 410,
	73, 316, 39, 314, 413, 266, 415, 200, 348, 96,
	41, 43, 44, 148, 149, 145, 146, 147, 97, 98,
	95, 122, 372, 102, 101, 64, 128, 348, 93, 331,
	12, 111, 110, 106, 107, 108, 109, 112, 113, 116,
	117, 28, 321, 329, 13, 322, 330, 318, 20, 79,
	30, 19, 63, 323, 315, 15, 16, 325, 33, 319,
	14, 103, 324, 29, 9, 320, 31, 32, 10, 18,
	317, 21, 11, 17, 25, 26, 326, 104, 27, 34,
	96, 3, 2, 306, 99, 100, 1, 388, 141, 97,
	98, 95, 139, 210, 102, 101, 269, 312, 38, 93,
	86, 12, 111, 110, 106, 107, 108, 109, 112, 113,
	116, 117, 28, 115, 114, 13, 276, 307, 94, 20,
	8, 30, 19, 0, 0, 0, 15, 16, 0, 33,
	310, 14, 103, 0, 29, 9, 0, 31, 32, 10,
	18, 0, 21, 11, 17, 25, 26, 0, 104, 27,
	34, 96, 0, 0, 0, 99, 100, 0, 0, 0,
	97, 98, 95, 0, 0, 102, 101, 0, 0, 0,
	93, 331, 0, 111, 110, 106, 107, 108, 109, 112,
	113, 116, 117, 0, 321, 329, 0, 322, 330, 318,
	0, 0, 0, 0, 0, 323, 315, 0, 0, 325,
	0, 319, 0, 103, 324, 0, 0, 320, 0, 0,
	0, 0, 317, 0, 96, 0, 0, 0, 326, 104,
	0, 0, 0, 97, 98, 95, 99, 100, 102, 101,
	0, 0, 0, 93, 331, 0, 111, 110, 106, 107,
	108, 109, 112, 113, 116, 117, 0, 321, 329, 0,
	322, 330, 318, 0, 0, 0, 0, 0, 323, 315,
	0, 0, 325, 0, 319, 0, 103, 324, 0, 0,
	320, 0, 0, 0, 0, 317, 0, 96, 0, 0,
	0, 326, 104, 0, 0, 0, 97, 98, 95, 99,
	100, 102, 101, 0, 214, 213, 93, 86, 0, 111,
	110, 106, 107, 108, 109, 112, 113, 116, 117, 0,
	0, 96, 183, 182, 180, 181, 179, 178, 0, 0,
	97, 98, 95, 0, 0, 102, 101, 0, 0, 103,
	93, 86, 0, 111, 110, 106, 107, 108, 109, 112,
	113, 116, 117, 0, 0, 104, 96, 0, 0, 0,
	132, 0, 99, 100, 0, 97, 98, 95, 0, 0,
	102, 101, 0, 103, 0, 93, 86, 0, 111, 110,
	106, 107, 108, 109, 112, 113, 116, 117, 0, 104,
	96, 0, 0, 0, 132, 0, 99, 100, 0, 97,
	98, 95, 0, 0, 102, 101, 0, 0, 103, 93,
	86, 0, 111, 110, 106, 107, 108, 109, 112, 113,
	116, 117, 0, 0, 104, 0, 0, 0, 0, 339,
	0, 99, 100, 0, 12, 151, 150, 148, 149, 145,
	146, 147, 103, 0, 0, 28, 0, 0, 13, 0,
	0, 0, 20, 0, 30, 19, 0, 0, 104, 15,
	16, 0, 33, 0, 14, 99, 100, 29, 9, 0,
	31, 32, 10, 18, 0, 21, 11, 17, 25, 26,
	96, 0, 27, 34, 0, 0, 0, 199, 0, 97,
	98, 95, 0, 0, 102, 101, 0, 0, 0, 195,
	86, 0, 111, 110, 106, 107, 108, 109, 112, 113,
	116, 117, 0, 0, 96, 0, 0, 0, 0, 0,
	0, 0, 0, 97, 98, 95, 0, 0, 102, 101,
	0, 0, 103, 193, 86, 0, 111, 110, 106, 107,
	108, 109, 112, 113, 116, 117, 0, 0, 104, 0,
	0, 0, 0, 0, 85, 99, 100, 86, 12, 0,
	0, 0, 0, 0, 0, 0, 103, 0, 0, 28,
	0, 0, 13, 0, 0, 0, 20, 0, 30, 19,
	0, 0, 104, 15, 16, 0, 33, 0, 14, 99,
	100, 29, 9, 0, 31, 32, 10, 18, 12, 21,
	11, 17, 25, 26, 0, 0, 27, 34, 0, 28,
	0, 0, 13, 0, 0, 0, 20, 0, 30, 19,
	0, 0, 0, 15, 16, 0, 33, 0, 14, 0,
	0, 29, 9, 0, 31, 32, 10, 18, 0, 21,
	11, 17, 25, 26, 0, 0, 27, 34, 164, 165,
	166, 167, 168, 169, 171, 170, 172, 173, 174, 163,
	383, 162, 161, 160, 159, 158, 156, 157, 152, 153,
	154, 155, 151, 150, 148, 149, 145, 146, 147, 164,
	165, 166, 167, 168, 169, 171, 170, 172, 173, 174,
	163, 0, 162, 161, 160, 159, 158, 156, 157, 152,
	153, 154, 155, 151, 150, 148, 149, 145, 146, 147,
	163, 0, 162, 161, 160, 159, 158, 156, 157, 152,
	153, 154, 155, 151, 150, 148, 149, 145, 146, 147,
	161, 160, 159, 158, 156, 157, 152, 153, 154, 155,
	151, 150, 148, 149, 145, 146, 147, 160, 159, 158,
	156, 157, 152, 153, 154, 155, 151, 150, 148, 149,
	145, 146, 147, 159, 158, 156, 157, 152, 153, 154,
	155, 151, 150, 148, 149, 145, 146, 147, 158, 156,
	157, 152, 153, 154, 155, 151, 150, 148, 149, 145,
	146, 147, 156, 157, 152, 153, 154, 155, 151, 150,
	148, 149, 145, 146, 147, 152, 153, 154, 155, 151,
	150, 148, 149, 145, 146, 147,
}
var yyPact = []int{

	-1000, 954, -1000, 342, -1000, -1000, 179, 179, 954, 12,
	12, 5, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 357, -1000, 181, -1000,
	-1000, 235, -1000, -1000, -1000, 179, -1000, -1000, -1000, -1000,
	179, -1000, 179, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 235, -1000, 307, 914, 767, 96, -5, -1000,
	53, 179, -35, 954, -35, -36, 62, -1000, -1000, 954,
	698, -10, 354, -1000, 186, 192, -1000, -1000, -27, -1000,
	1073, -1000, -1000, 467, 685, 767, 767, 767, 767, 767,
	767, 767, 767, 891, 857, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 50, 43, -1000, -1000, -1000, -1000,
	-1000, -1000, 790, -1000, -1000, -1000, 31, 295, -37, 235,
	-1000, 1073, 664, -1000, 914, -1000, -1000, -1000, -1000, 161,
	-1, -1000, 767, 177, -1000, 767, 767, 767, 767, 767,
	767, 767, 767, 767, 767, 767, 767, 767, 767, 767,
	767, 767, 767, 767, 767, 767, 767, 767, 767, 767,
	767, 767, 767, 767, 767, 248, 127, 1073, 767, 767,
	67, 67, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 467, -1000, 467, -1000, -1000, -1000, -1000,
	378, 62, -1000, 62, 767, -1000, -1000, 339, -1000, -39,
	664, 352, 275, 767, 67, -1000, 10, 914, 767, -1000,
	-11, -29, -1000, -1000, -1000, -1000, 301, 301, 391, 391,
	805, 805, 805, 805, 1179, 1179, 1168, 1155, 1141, 1126,
	1110, 237, 1073, 1073, 1073, 1073, 1073, 1073, 1073, 1073,
	1073, 1073, 1073, -12, -1000, 23, 767, -1000, -14, 349,
	1073, 99, -1000, -1000, 248, 248, 378, 399, 327, -1000,
	-1000, 240, 767, 30, -1000, 1073, 396, -1000, 235, -1000,
	316, 275, -1000, -1000, -32, -1000, -1000, -15, -33, -1000,
	-1000, 767, 733, 158, -1000, -1000, 767, -1000, -16, -19,
	397, -1000, 378, 767, -1000, -1000, -1000, -1000, -1000, 308,
	395, -1000, 601, 393, -37, 171, 26, 156, 538, 767,
	154, 388, 386, 67, 141, 102, 93, -1000, 313, 767,
	227, 220, -1000, -1000, -1000, -1000, -1000, 1093, -1000, 664,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 282, -1000, -1000,
	-1000, -1000, -1000, -1000, 767, 84, 767, 6, 384, 767,
	-1000, -1000, 379, 767, 767, 767, 1042, -1000, -1000, -58,
	-1000, 235, 269, 124, 467, 110, 74, -1000, 101, -1000,
	-21, -22, -23, -1000, -1000, -1000, 698, 538, 368, -1000,
	235, 538, 767, 538, 367, 365, 363, -1000, 33, 767,
	302, -1000, 95, -1000, -1000, -1000, -1000, 538, 359, 343,
	-1000, 767, -1000, -24, 538, -1000,
}
var yyPgo = []int{

	0, 9, 252, 349, 288, 23, 302, 208, 530, 25,
	112, 201, 329, 6, 18, 45, 2, 47, 11, 1,
	13, 0, 22, 528, 14, 17, 527, 526, 32, 524,
	523, 46, 508, 507, 15, 10, 3, 506, 24, 28,
	503, 34, 37, 502, 498, 38, 12, 4, 5, 497,
	496, 492, 491, 134, 462, 459, 436, 435, 7, 432,
	26, 431, 417, 27, 415, 16, 413, 411, 410, 408,
	407, 405, 404, 29, 400,
}
var yyR1 = []int{

	0, 50, 50, 51, 51, 54, 56, 51, 53, 57,
	53, 53, 31, 31, 32, 32, 32, 32, 26, 26,
	36, 59, 36, 36, 55, 55, 60, 60, 62, 61,
	64, 61, 63, 63, 65, 65, 37, 37, 37, 41,
	41, 42, 42, 42, 43, 43, 43, 44, 44, 44,
	47, 47, 39, 39, 39, 40, 40, 40, 40, 48,
	48, 48, 14, 14, 15, 15, 15, 15, 15, 18,
	27, 27, 27, 33, 33, 34, 34, 34, 19, 19,
	19, 49, 49, 35, 66, 35, 35, 35, 67, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	16, 16, 45, 45, 46, 20, 20, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 22, 22,
	22, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 29, 29, 30, 30, 24, 24, 25, 25,
	68, 11, 52, 52, 13, 13, 13, 13, 13, 13,
	13, 13, 10, 58, 12, 69, 12, 12, 12, 70,
	12, 12, 12, 71, 72, 12, 74, 12, 12, 7,
	7, 9, 9, 2, 2, 2, 8, 8, 3, 3,
	73, 73, 73, 73, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 4, 4, 4, 4, 4, 4, 4,
	5, 5, 5, 17, 38, 1, 1,
}
var yyR2 = []int{

	0, 0, 2, 2, 3, 0, 0, 6, 1, 0,
	4, 3, 1, 3, 1, 3, 4, 4, 2, 3,
	1, 0, 4, 3, 0, 4, 1, 3, 0, 4,
	0, 5, 0, 1, 1, 3, 1, 3, 2, 0,
	1, 2, 3, 1, 1, 4, 4, 2, 3, 3,
	1, 3, 3, 2, 2, 2, 3, 1, 2, 1,
	1, 2, 0, 1, 1, 2, 2, 3, 3, 3,
	0, 2, 2, 1, 2, 3, 2, 2, 2, 1,
	2, 1, 2, 2, 0, 2, 5, 7, 0, 10,
	5, 7, 3, 5, 2, 2, 3, 5, 5, 5,
	0, 1, 0, 1, 1, 1, 3, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 5, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 1, 5,
	7, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 3, 5, 5, 4, 4, 3, 3, 2,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 1, 2, 0, 1, 1, 3,
	0, 4, 0, 1, 1, 1, 1, 2, 2, 3,
	2, 3, 1, 1, 2, 0, 4, 2, 2, 0,
	4, 2, 2, 0, 0, 7, 0, 5, 1, 1,
	2, 0, 2, 1, 1, 1, 1, 2, 1, 1,
	1, 3, 2, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1,
}
var yyChk = []int{

	-1000, -50, -51, -52, -58, -13, -12, -6, -8, 78,
	82, 86, 44, 58, 74, 69, 70, 87, 83, 65,
	62, 85, -3, -5, -4, 88, 89, 92, 55, 77,
	64, 80, 81, 72, 93, 4, -53, -31, -32, 34,
	-38, 42, -1, 43, 44, -7, -2, -6, -5, -4,
	-7, -12, -6, -3, -1, -11, 96, -1, -11, -1,
	96, 4, 5, -54, -57, 42, 41, -9, -31, -2,
	-9, -7, -69, -68, -70, -71, -74, -53, -31, -55,
	6, -14, -15, -17, -10, 40, 43, -13, -45, -46,
	-21, -22, -28, 42, -23, 34, 23, 32, 33, 98,
	99, 38, 37, 75, 91, -17, 47, 48, 49, 50,
	46, 45, 51, 52, -29, -30, 53, 54, -31, -5,
	94, -11, -61, -10, -11, 96, -73, 43, -56, -58,
	-47, -21, 96, 94, 5, -41, -31, -42, 34, -43,
	42, -44, 41, 40, 95, 34, 35, 36, 32, 33,
	31, 30, 26, 27, 28, 29, 24, 25, 23, 22,
	21, 20, 19, 17, 6, 7, 8, 9, 10, 11,
	13, 12, 14, 15, 16, -10, -20, -21, 42, 41,
	39, 40, 38, 37, -22, -22, -22, -22, -22, -22,
	-22, -22, -28, 42, -28, 42, 53, 54, -10, 97,
	-62, -72, 97, 5, 6, -18, 96, -60, -31, -48,
	-40, -47, -39, 41, 40, -15, -9, 42, 41, 94,
	-42, -45, 40, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -20, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -41, 34, 42, 5, 94, -24, -25,
	-21, -20, -1, -1, -10, -10, -64, -63, -65, -37,
	-31, -38, 18, -73, -73, -21, -27, 4, 5, 97,
	-47, -39, 5, 6, -46, -1, -42, -14, -45, 94,
	95, 18, 94, -9, -20, 94, 5, 95, -41, -41,
	-63, 4, 5, 18, -46, 97, 97, -26, -19, -58,
	2, -35, -33, -16, -66, 68, -67, 84, 61, 73,
	79, 56, 59, 67, 76, 71, 90, -34, -20, 57,
	60, 43, -60, 5, 95, 94, 95, -21, -22, 96,
	-25, 94, 94, 4, -65, -46, 4, -36, -31, 4,
	-34, -35, 4, -18, 42, 66, 42, -19, -16, 42,
	4, 4, -1, 42, 42, 42, -21, 18, 18, -48,
	4, 5, -59, -20, 42, -20, 84, 4, -20, 4,
	-24, -24, -24, 18, 97, -36, 6, 94, -49, -16,
	-58, 94, 42, 94, 94, 94, 94, -47, -19, 4,
	-36, -19, -20, -19, 4, 4, 4, 63, -16, 94,
	-19, 4, 4, -16, 94, -19,
}
var yyDef = []int{

	1, -2, 2, 0, 183, 193, 184, 185, 186, 0,
	0, 0, 208, 224, 225, 226, 227, 228, 229, 230,
	231, 232, 216, 218, 219, 240, 241, 242, 233, 234,
	235, 236, 237, 238, 239, 3, 0, -2, 12, 211,
	14, 0, 244, 245, 246, 187, 209, 213, 214, 215,
	188, 211, 190, 217, -2, 197, 180, -2, 201, -2,
	206, 4, 0, 24, 0, 62, 102, 0, 0, 210,
	189, 191, 0, 0, 0, 0, 0, 11, -2, 6,
	0, 0, 63, 64, 39, 0, 243, 192, 0, 103,
	104, 107, 138, 0, 141, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 161, 162, 163, 164, 165,
	166, 167, 168, 169, 170, 171, 172, 174, 13, 212,
	15, 196, 0, 28, 200, 204, 0, 220, 0, 0,
	10, 50, 0, 16, 0, 65, 66, 40, 211, 43,
	0, 44, 102, 0, 17, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 39, 0, 105, 176, 0,
	0, 0, 159, 160, 142, 143, 144, 145, 146, 147,
	148, 149, 150, 0, 151, 0, 173, 175, 30, 181,
	32, 0, 207, 222, 0, 7, 70, 0, 26, 0,
	59, 60, 57, 0, 0, 68, 41, 62, 102, 47,
	0, 0, 67, 108, 109, 110, 111, 112, 113, 114,
	115, 116, 117, 118, 119, 120, 121, 122, 123, 124,
	125, 0, 127, 128, 129, 130, 131, 132, 133, 134,
	135, 136, 137, 0, 211, 0, 0, 152, 0, 177,
	178, 0, 157, 158, 39, 39, 32, 0, 33, 34,
	36, 14, 0, 0, 223, 221, -2, 25, 0, 51,
	61, 58, 55, 54, 0, 53, 42, 0, 0, 49,
	48, 0, 0, 41, 106, 155, 0, 156, 0, 0,
	0, 29, 0, 0, 38, 205, 69, 71, 72, 0,
	0, 79, -2, 0, 0, 0, 0, 0, -2, 100,
	0, 0, 0, 0, 0, 0, 0, 73, 101, 0,
	0, 243, 27, 56, 52, 45, 46, 126, 139, 0,
	179, 153, 154, 31, 35, 37, 18, 0, -2, 78,
	74, 80, 83, 85, 0, 0, 0, 0, 0, 0,
	94, 95, 0, 176, 176, 176, 0, 76, 77, 0,
	19, 0, 0, 0, 100, 0, 0, 92, 0, 96,
	0, 0, 0, 75, 140, 23, 0, -2, 0, 81,
	0, -2, 0, -2, 0, 0, 0, 22, 86, 100,
	82, 90, 0, 93, 97, 98, 99, -2, 0, 0,
	87, 100, 91, 0, -2, 89,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 98, 3, 3, 3, 36, 23, 3,
	42, 94, 34, 32, 5, 33, 40, 35, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 18, 4,
	26, 6, 27, 17, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 41, 3, 95, 22, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 96, 21, 97, 99,
}
var yyTok2 = []int{

	2, 3, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 16, 19, 20, 24, 25, 28, 29, 30, 31,
	37, 38, 39, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
	90, 91, 92, 93,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 3:
		//line cc.y:78
		{
			dodecl(xdecl, lastclass, lasttype, new(Node))
		}
	case 5:
		//line cc.y:83
		{
			lastdcl = new(Type)
			firstarg = new(Sym)
			dodecl(xdecl, lastclass, lasttype, yyS[yypt-0].node)
			if lastdcl == nil || lastdcl.Etype != TFUNC {
				diag(yyS[yypt-0].node, "not a function")
				lastdcl = Types[TFUNC];
			}
			thisfn = lastdcl;
			markdcl()
			firstdcl = dclstack;
			argmark(yyS[yypt-0].node, 0)
		}
	case 6:
		//line cc.y:97
		{
			argmark(yyS[yypt-2].node, 1)
		}
	case 7:
		//line cc.y:101
		{
			n := revertdcl()
			if n{
				yyS[yypt-0].node = Node.new(OLIST, n, yyS[yypt-0].node)
			}
			if !debug['a'] && !debug['Z']{
				codgen(yyS[yypt-0].node, yyS[yypt-4].node)
			}
		}
	case 8:
		//line cc.y:113
		{
			dodecl(xdecl, lastclass, lasttype, yyS[yypt-0].node)
		}
	case 9:
		//line cc.y:117
		{
			yyS[yypt-0].node = dodecl(xdecl, lastclass, lasttype, yyS[yypt-0].node)
		}
	case 10:
		//line cc.y:121
		{
			doinit(yyS[yypt-3].node.sym, yyS[yypt-3].node.Type, 0, yyS[yypt-0].node)
		}
	case 12:
		yyVAL.node = yyS[yypt-0].node
	case 13:
		//line cc.y:129
		{
			yyVAL.node = Node.new(OIND, yyS[yypt-0].node, new(Node))
			yyVAL.node.garb = simpleg(yyS[yypt-1].lval)
		}
	case 14:
		yyVAL.node = yyS[yypt-0].node
	case 15:
		//line cc.y:137
		{
			yyVAL.node = yyS[yypt-1].node;
		}
	case 16:
		//line cc.y:141
		{
			yyVAL.node = Node.new(OFUNC, yyS[yypt-3].node, yyS[yypt-1].node)
		}
	case 17:
		//line cc.y:145
		{
			yyVAL.node = Node.new(OARRAY, yyS[yypt-3].node, yyS[yypt-1].node)
		}
	case 18:
		//line cc.y:154
		{
			yyVAL.node = dodecl(adecl, lastclass, lasttype, new(Node))
		}
	case 19:
		//line cc.y:158
		{
			yyVAL.node = yyS[yypt-1].node;
		}
	case 20:
		//line cc.y:164
		{
			dodecl(adecl, lastclass, lasttype, yyS[yypt-0].node)
			yyVAL.node = new(Node)
		}
	case 21:
		//line cc.y:169
		{
			yyS[yypt-0].node = dodecl(adecl, lastclass, lasttype, yyS[yypt-0].node)
		}
	case 22:
		//line cc.y:173
		{
			w := yyS[yypt-3].node.sym.Type.width
			yyVAL.node = doinit(yyS[yypt-3].node.sym, yyS[yypt-3].node.Type, 0, yyS[yypt-0].node)
			yyVAL.node = contig(yyS[yypt-3].node.sym, yyVAL.node, w)
		}
	case 23:
		//line cc.y:179
		{
			yyVAL.node = yyS[yypt-2].node;
			if yyS[yypt-0].node != nil{ 
				yyVAL.node = yyS[yypt-0].node;
			
			if yyS[yypt-2].node != nil{
				yyVAL.node = Node.new(OLIST, yyS[yypt-2].node, yyS[yypt-0].node)
				}
			}
		}
	case 26:
		//line cc.y:198
		{
			dodecl(pdecl, lastclass, lasttype, yyS[yypt-0].node)
		}
	case 28:
		//line cc.y:208
		{
			lasttype = yyS[yypt-0].ytype;
		}
	case 30:
		//line cc.y:213
		{
			lasttype = yyS[yypt-0].ytype;
		}
	case 32:
		//line cc.y:219
		{
			lastfield = 0;
			edecl(CXXX, lasttype, S)
		}
	case 34:
		//line cc.y:227
		{
			dodecl(edecl, CXXX, lasttype, yyS[yypt-0].node)
		}
	case 36:
		//line cc.y:234
		{
			lastbit = 0;
			firstbit = 1;
		}
	case 37:
		//line cc.y:239
		{
			yyVAL.node = Node.new(OBIT, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 38:
		//line cc.y:243
		{
			yyVAL.node = Node.new(OBIT, new(Node), yyS[yypt-0].node)
		}
	case 39:
		//line cc.y:251
		{
			yyVAL.node = new(Node)
		}
	case 40:
		yyVAL.node = yyS[yypt-0].node
	case 41:
		//line cc.y:258
		{
			yyVAL.node = Node.new(OIND, new(Node), new(Node))
			yyVAL.node.garb = simpleg(yyS[yypt-0].lval)
		}
	case 42:
		//line cc.y:263
		{
			yyVAL.node = Node.new(OIND, yyS[yypt-0].node, new(Node))
			yyVAL.node.garb = simpleg(yyS[yypt-1].lval)
		}
	case 43:
		yyVAL.node = yyS[yypt-0].node
	case 44:
		yyVAL.node = yyS[yypt-0].node
	case 45:
		//line cc.y:272
		{
			yyVAL.node = Node.new(OFUNC, yyS[yypt-3].node, yyS[yypt-1].node)
		}
	case 46:
		//line cc.y:276
		{
			yyVAL.node = Node.new(OARRAY, yyS[yypt-3].node, yyS[yypt-1].node)
		}
	case 47:
		//line cc.y:282
		{
			yyVAL.node = Node.new(OFUNC, new(Node), new(Node))
		}
	case 48:
		//line cc.y:286
		{
			yyVAL.node = Node.new(OARRAY, new(Node), yyS[yypt-1].node)
		}
	case 49:
		//line cc.y:290
		{
			yyVAL.node = yyS[yypt-1].node;
		}
	case 50:
		yyVAL.node = yyS[yypt-0].node
	case 51:
		//line cc.y:297
		{
			yyVAL.node = Node.new(OINIT, invert(yyS[yypt-1].node), new(Node))
		}
	case 52:
		//line cc.y:303
		{
			yyVAL.node = Node.new(OARRAY, yyS[yypt-1].node, new(Node))
		}
	case 53:
		//line cc.y:307
		{
			yyVAL.node = Node.new(OELEM, new(Node), new(Node))
			yyVAL.node.sym = yyS[yypt-0].sym;
		}
	case 54:
		yyVAL.node = yyS[yypt-0].node
	case 55:
		yyVAL.node = yyS[yypt-0].node
	case 56:
		//line cc.y:316
		{
			yyVAL.node = Node.new(OLIST, yyS[yypt-2].node, yyS[yypt-1].node)
		}
	case 57:
		yyVAL.node = yyS[yypt-0].node
	case 58:
		//line cc.y:321
		{
			yyVAL.node = Node.new(OLIST, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 59:
		yyVAL.node = yyS[yypt-0].node
	case 60:
		yyVAL.node = yyS[yypt-0].node
	case 61:
		//line cc.y:329
		{
			yyVAL.node = Node.new(OLIST, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 62:
		//line cc.y:334
		{
			yyVAL.node = new(Node);
		}
	case 63:
		//line cc.y:338
		{
			yyVAL.node = invert(yyS[yypt-0].node)
		}
	case 64:
		yyVAL.node = yyS[yypt-0].node
	case 65:
		//line cc.y:346
		{
			yyVAL.node = Node.new(OPROTO, yyS[yypt-0].node, new(Node))
			yyVAL.node.Type = yyS[yypt-1].ytype;
		}
	case 66:
		//line cc.y:351
		{
			yyVAL.node = Node.new(OPROTO, yyS[yypt-0].node, new(Node))
			yyVAL.node.Type = yyS[yypt-1].ytype;
		}
	case 67:
		//line cc.y:356
		{
			yyVAL.node = Node.new(ODOTDOT, new(Node), new(Node))
		}
	case 68:
		//line cc.y:360
		{
			yyVAL.node = Node.new(OLIST, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 69:
		//line cc.y:366
		{
			yyVAL.node = invert(yyS[yypt-1].node)
		//	if $2 != nil
	//		$$ = Node.new(OLIST, $2, $$)
		if yyVAL.node == nil{
				yyVAL.node = Node.new(OLIST, new(Node), new(Node))
			}
		}
	case 70:
		//line cc.y:376
		{
			yyVAL.node = new(Node)
		}
	case 71:
		//line cc.y:380
		{
			yyVAL.node = Node.new(OLIST, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 72:
		//line cc.y:384
		{
			yyVAL.node = Node.new(OLIST, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 73:
		yyVAL.node = yyS[yypt-0].node
	case 74:
		//line cc.y:391
		{
			yyVAL.node = Node.new(OLIST, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 75:
		//line cc.y:397
		{
			yyVAL.node = Node.new(OCASE, yyS[yypt-1].node, new(Node))
		}
	case 76:
		//line cc.y:401
		{
			yyVAL.node = Node.new(OCASE, new(Node), new(Node))
		}
	case 77:
		//line cc.y:405
		{
			yyVAL.node = Node.new(OLABEL, dcllabel(yyS[yypt-1].sym, 1), new(Node))
		}
	case 78:
		//line cc.y:411
		{
			yyVAL.node = new(Node)
		}
	case 79:
		yyVAL.node = yyS[yypt-0].node
	case 80:
		//line cc.y:416
		{
			yyVAL.node = Node.new(OLIST, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 81:
		yyVAL.node = yyS[yypt-0].node
	case 82:
		//line cc.y:423
		{
			yyVAL.node = yyS[yypt-0].node;
		}
	case 83:
		yyVAL.node = yyS[yypt-0].node
	case 84:
		//line cc.y:429
		{
			markdcl()
		}
	case 85:
		//line cc.y:433
		{
			yyVAL.node = revertdcl()
			if yyVAL.node{
				yyVAL.node = Node.new(OLIST, yyVAL.node, yyS[yypt-0].node)
			}else{
				yyVAL.node = yyS[yypt-0].node;
			}
		}
	case 86:
		//line cc.y:442
		{
			yyVAL.node = Node.new(OIF, yyS[yypt-2].node, new(OLIST, yyS[yypt-0].node, new(Node)))
			if yyS[yypt-0].node == nil{
				warn(yyS[yypt-2].node, "empty if body")
			}
		}
	case 87:
		//line cc.y:449
		{
			yyVAL.node = Node.new(OIF, yyS[yypt-4].node, new(OLIST, yyS[yypt-2].node, yyS[yypt-0].node))
			if yyS[yypt-2].node == nil{
				warn(yyS[yypt-4].node, "empty if body")
			}
			if yyS[yypt-0].node == nil{
				warn(yyS[yypt-4].node, "empty else body")
			}
		}
	case 88:
		//line cc.y:458
		{ markdcl() }
	case 89:
		//line cc.y:459
		{
			yyVAL.node = revertdcl()
			if yyVAL.node{
				if yyS[yypt-6].node {
					yyS[yypt-6].node = Node.new(OLIST, yyVAL.node, yyS[yypt-6].node)
				}else{
					yyS[yypt-6].node = yyVAL.node
				}
			}
			yyVAL.node = Node.new(OFOR, Node.new(OLIST, yyS[yypt-4].node, Node.new(OLIST, yyS[yypt-6].node, yyS[yypt-2].node)), yyS[yypt-0].node)
		}
	case 90:
		//line cc.y:471
		{
			yyVAL.node = Node.new(OWHILE, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 91:
		//line cc.y:475
		{
			yyVAL.node = Node.new(ODWHILE, yyS[yypt-2].node, yyS[yypt-5].node)
		}
	case 92:
		//line cc.y:479
		{
			yyVAL.node = Node.new(ORETURN, yyS[yypt-1].node, new(Node))
			yyVAL.node.Type = thisfn.link;
		}
	case 93:
		//line cc.y:484
		{
			yyVAL.node = Node.new(OCONST, new(Node), new(Node))
			yyVAL.node.vconst = 0;
			yyVAL.node.Type = types[TINT];
			yyS[yypt-2].node = Node.new(OSUB, yyVAL.node, yyS[yypt-2].node)
	
			yyVAL.node = Node.new(OCONST, new(Node), new(Node))
			yyVAL.node.vconst = 0;
			yyVAL.node.Type = types[TINT];
			yyS[yypt-2].node = Node.new(OSUB, yyVAL.node, yyS[yypt-2].node)
	
			yyVAL.node = Node.new(OSWITCH, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 94:
		//line cc.y:498
		{
			yyVAL.node = Node.new(OBREAK, new(Node), new(Node))
		}
	case 95:
		//line cc.y:502
		{
			yyVAL.node = Node.new(OCONTINUE, new(Node), new(Node))
		}
	case 96:
		//line cc.y:506
		{
			yyVAL.node = Node.new(OGOTO, dcllabel(yyS[yypt-1].sym, 0), new(Node))
		}
	case 97:
		//line cc.y:510
		{
			yyVAL.node = Node.new(OUSED, yyS[yypt-2].node, new(Node))
		}
	case 98:
		//line cc.y:514
		{
			yyVAL.node = Node.new(OPREFETCH, yyS[yypt-2].node, new(Node))
		}
	case 99:
		//line cc.y:518
		{
			yyVAL.node = Node.new(OSET, yyS[yypt-2].node, new(Node))
		}
	case 100:
		//line cc.y:523
		{
			yyVAL.node = new(Node)
		}
	case 101:
		yyVAL.node = yyS[yypt-0].node
	case 102:
		//line cc.y:529
		{
			yyVAL.node = new(Node)
		}
	case 103:
		yyVAL.node = yyS[yypt-0].node
	case 104:
		//line cc.y:536
		{
			yyVAL.node = Node.new(OCAST, yyS[yypt-0].node, new(Node))
			yyVAL.node.Type = types[TLONG];
		}
	case 105:
		yyVAL.node = yyS[yypt-0].node
	case 106:
		//line cc.y:544
		{
			yyVAL.node = Node.new(OCOMMA, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 107:
		yyVAL.node = yyS[yypt-0].node
	case 108:
		//line cc.y:551
		{
			yyVAL.node = Node.new(OMUL, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 109:
		//line cc.y:555
		{
			yyVAL.node = Node.new(ODIV, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 110:
		//line cc.y:559
		{
			yyVAL.node = Node.new(OMOD, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 111:
		//line cc.y:563
		{
			yyVAL.node = Node.new(OADD, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 112:
		//line cc.y:567
		{
			yyVAL.node = Node.new(OSUB, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 113:
		//line cc.y:571
		{
			yyVAL.node = Node.new(OASHR, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 114:
		//line cc.y:575
		{
			yyVAL.node = Node.new(OASHL, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 115:
		//line cc.y:579
		{
			yyVAL.node = Node.new(OLT, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 116:
		//line cc.y:583
		{
			yyVAL.node = Node.new(OGT, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 117:
		//line cc.y:587
		{
			yyVAL.node = Node.new(OLE, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 118:
		//line cc.y:591
		{
			yyVAL.node = Node.new(OGE, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 119:
		//line cc.y:595
		{
			yyVAL.node = Node.new(OEQ, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 120:
		//line cc.y:599
		{
			yyVAL.node = Node.new(ONE, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 121:
		//line cc.y:603
		{
			yyVAL.node = Node.new(OAND, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 122:
		//line cc.y:607
		{
			yyVAL.node = Node.new(OXOR, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 123:
		//line cc.y:611
		{
			yyVAL.node = Node.new(OOR, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 124:
		//line cc.y:615
		{
			yyVAL.node = Node.new(OANDAND, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 125:
		//line cc.y:619
		{
			yyVAL.node = Node.new(OOROR, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 126:
		//line cc.y:623
		{
			yyVAL.node = Node.new(OCOND, yyS[yypt-4].node, Node.new(OLIST, yyS[yypt-2].node, yyS[yypt-0].node))
		}
	case 127:
		//line cc.y:627
		{
			yyVAL.node = Node.new(OAS, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 128:
		//line cc.y:631
		{
			yyVAL.node = Node.new(OASADD, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 129:
		//line cc.y:635
		{
			yyVAL.node = Node.new(OASSUB, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 130:
		//line cc.y:639
		{
			yyVAL.node = Node.new(OASMUL, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 131:
		//line cc.y:643
		{
			yyVAL.node = Node.new(OASDIV, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 132:
		//line cc.y:647
		{
			yyVAL.node = Node.new(OASMOD, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 133:
		//line cc.y:651
		{
			yyVAL.node = Node.new(OASASHL, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 134:
		//line cc.y:655
		{
			yyVAL.node = Node.new(OASASHR, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 135:
		//line cc.y:659
		{
			yyVAL.node = Node.new(OASAND, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 136:
		//line cc.y:663
		{
			yyVAL.node = Node.new(OASXOR, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 137:
		//line cc.y:667
		{
			yyVAL.node = Node.new(OASOR, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 138:
		yyVAL.node = yyS[yypt-0].node
	case 139:
		//line cc.y:674
		{
			yyVAL.node = Node.new(OCAST, yyS[yypt-0].node, new(Node))
			dodecl(NODECL, CXXX, yyS[yypt-3].ytype, yyS[yypt-2].node)
			yyVAL.node.Type = lastdcl;
			yyVAL.node.xcast = 1;
		}
	case 140:
		//line cc.y:681
		{
			yyVAL.node = Node.new(OSTRUCT, yyS[yypt-1].node, new(Node))
			dodecl(NODECL, CXXX, yyS[yypt-5].ytype, yyS[yypt-4].node)
			yyVAL.node.Type = lastdcl;
		}
	case 141:
		yyVAL.node = yyS[yypt-0].node
	case 142:
		//line cc.y:690
		{
			yyVAL.node = Node.new(OIND, yyS[yypt-0].node, new(Node))
		}
	case 143:
		//line cc.y:694
		{
			yyVAL.node = Node.new(OADDR, yyS[yypt-0].node, new(Node))
		}
	case 144:
		//line cc.y:698
		{
			yyVAL.node = Node.new(OPOS, yyS[yypt-0].node, new(Node))
		}
	case 145:
		//line cc.y:702
		{
			yyVAL.node = Node.new(ONEG, yyS[yypt-0].node, new(Node))
		}
	case 146:
		//line cc.y:706
		{
			yyVAL.node = Node.new(ONOT, yyS[yypt-0].node, new(Node))
		}
	case 147:
		//line cc.y:710
		{
			yyVAL.node = Node.new(OCOM, yyS[yypt-0].node, new(Node))
		}
	case 148:
		//line cc.y:714
		{
			yyVAL.node = Node.new(OPREINC, yyS[yypt-0].node, new(Node))
		}
	case 149:
		//line cc.y:718
		{
			yyVAL.node = Node.new(OPREDEC, yyS[yypt-0].node, new(Node))
		}
	case 150:
		//line cc.y:722
		{
			yyVAL.node = Node.new(OSIZE, yyS[yypt-0].node, new(Node))
		}
	case 151:
		//line cc.y:726
		{
			yyVAL.node = Node.new(OSIGN, yyS[yypt-0].node, new(Node))
		}
	case 152:
		//line cc.y:732
		{
			yyVAL.node = yyS[yypt-1].node;
		}
	case 153:
		//line cc.y:736
		{
			yyVAL.node = Node.new(OSIZE, new(Node), new(Node))
			dodecl(NODECL, CXXX, yyS[yypt-2].ytype, yyS[yypt-1].node)
			yyVAL.node.Type = lastdcl;
		}
	case 154:
		//line cc.y:742
		{
			yyVAL.node = Node.new(OSIGN, new(Node), new(Node))
			dodecl(NODECL, CXXX, yyS[yypt-2].ytype, yyS[yypt-1].node)
			yyVAL.node.Type = lastdcl;
		}
	case 155:
		//line cc.y:748
		{
			yyVAL.node = Node.new(OFUNC, yyS[yypt-3].node, new(Node))
			if yyS[yypt-3].node.op == ONAME{
				if yyS[yypt-3].node.Type == nil{
					dodecl(xdecl, CXXX, types[TINT], yyVAL.node)
				}
			}
			yyVAL.node.right = invert(yyS[yypt-1].node)
		}
	case 156:
		//line cc.y:758
		{
			yyVAL.node = Node.new(OIND, Node.new(OADD, yyS[yypt-3].node, yyS[yypt-1].node), new(Node))
		}
	case 157:
		//line cc.y:762
		{
			yyVAL.node = Node.new(ODOT, Node.new(OIND, yyS[yypt-2].node, new(Node)), new(Node))
			yyVAL.node.sym = yyS[yypt-0].sym;
		}
	case 158:
		//line cc.y:767
		{
			yyVAL.node = Node.new(ODOT, yyS[yypt-2].node, new(Node))
			yyVAL.node.sym = yyS[yypt-0].sym;
		}
	case 159:
		//line cc.y:772
		{
			yyVAL.node = Node.new(OPOSTINC, yyS[yypt-1].node, new(Node))
		}
	case 160:
		//line cc.y:776
		{
			yyVAL.node = Node.new(OPOSTDEC, yyS[yypt-1].node, new(Node))
		}
	case 161:
		yyVAL.node = yyS[yypt-0].node
	case 162:
		//line cc.y:781
		{
			yyVAL.node = Node.new(OCONST, new(Node), new(Node))
			yyVAL.node.Type = types[TINT]
			yyVAL.node.vconst = yyS[yypt-0].vval
			yyVAL.node.cstring = symb
		}
	case 163:
		//line cc.y:788
		{
			yyVAL.node = Node.new(OCONST, new(Node), new(Node))
			yyVAL.node.Type = types[TLONG]
			yyVAL.node.vconst = yyS[yypt-0].vval
			yyVAL.node.cstring = symb
		}
	case 164:
		//line cc.y:795
		{
			yyVAL.node = Node.new(OCONST, new(Node), new(Node))
			yyVAL.node.Type = types[TUINT]
			yyVAL.node.vconst = yyS[yypt-0].vval
			yyVAL.node.cstring = symb
		}
	case 165:
		//line cc.y:802
		{
			yyVAL.node = Node.new(OCONST, new(Node), new(Node))
			yyVAL.node.Type = types[TULONG]
			yyVAL.node.vconst = yyS[yypt-0].vval
			yyVAL.node.cstring = symb
		}
	case 166:
		//line cc.y:809
		{
			yyVAL.node = Node.new(OCONST, new(Node), new(Node))
			yyVAL.node.Type = types[TDOUBLE]
			yyVAL.node.fconst = yyS[yypt-0].dval
			yyVAL.node.cstring = symb
		}
	case 167:
		//line cc.y:816
		{
			yyVAL.node = Node.new(OCONST, new(Node), new(Node))
			yyVAL.node.Type = types[TFLOAT]
			yyVAL.node.fconst = yyS[yypt-0].dval
			yyVAL.node.cstring = symb
		}
	case 168:
		//line cc.y:823
		{
			yyVAL.node = Node.new(OCONST, new(Node), new(Node))
			yyVAL.node.Type = types[TVLONG]
			yyVAL.node.vconst = yyS[yypt-0].vval
			yyVAL.node.cstring = symb
		}
	case 169:
		//line cc.y:830
		{
			yyVAL.node = Node.new(OCONST, new(Node), new(Node))
			yyVAL.node.Type = types[TUVLONG]
			yyVAL.node.vconst = yyS[yypt-0].vval
			yyVAL.node.cstring = symb
		}
	case 170:
		yyVAL.node = yyS[yypt-0].node
	case 171:
		yyVAL.node = yyS[yypt-0].node
	case 172:
		//line cc.y:841
		{
			yyVAL.node = Node.new(OSTRING, new(Node), new(Node))
			yyVAL.node.Type = typ(TARRAY, types[TCHAR])
			yyVAL.node.Type.width = yyS[yypt-0].sval.l + 1
			yyVAL.node.cstring = yyS[yypt-0].sval.s
			yyVAL.node.sym = symstring
			yyVAL.node.etype = TARRAY
			yyVAL.node.class = CSTATIC
		}
	case 173:
		//line cc.y:851
		{
			n := yyS[yypt-1].node.Type.width - 1
			s := alloc(n+yyS[yypt-0].sval.l+MAXALIGN)
	
			memcpy(s, yyS[yypt-1].node.cstring, n)
			memcpy(s+n, yyS[yypt-0].sval.s, yyS[yypt-0].sval.l)
			s[n+yyS[yypt-0].sval.l] = 0
	
			yyVAL.node = yyS[yypt-1].node
			yyVAL.node.Type.width += yyS[yypt-0].sval.l
			yyVAL.node.cstring = s
		}
	case 174:
		//line cc.y:866
		{
			yyVAL.node = Node.new(OLSTRING, new(Node), new(Node))
			yyVAL.node.Type = typ(TARRAY, types[TUSHORT])
			yyVAL.node.Type.width = yyS[yypt-0].sval.l + sizeof(ushort)
			yyVAL.node.rstring = *ushort(yyS[yypt-0].sval.s)
			yyVAL.node.sym = symstring
			yyVAL.node.etype = TARRAY
			yyVAL.node.class = CSTATIC
		}
	case 175:
		//line cc.y:876
		{
			n := yyS[yypt-1].node.Type.width - sizeof(ushort)
			s := alloc(n+yyS[yypt-0].sval.l+MAXALIGN)
	
			memcpy(s, yyS[yypt-1].node.rstring, n)
			memcpy(s+n, yyS[yypt-0].sval.s, yyS[yypt-0].sval.l)
			*(*ushort)(s+n+yyS[yypt-0].sval.l) = 0
	
			yyVAL.node = yyS[yypt-1].node
			yyVAL.node.Type.width += yyS[yypt-0].sval.l
			yyVAL.node.rstring = *ushort(s)
		}
	case 176:
		//line cc.y:890
		{
			yyVAL.node = new(Node)
		}
	case 177:
		yyVAL.node = yyS[yypt-0].node
	case 178:
		yyVAL.node = yyS[yypt-0].node
	case 179:
		//line cc.y:898
		{
			yyVAL.node = Node.new(OLIST, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 180:
		//line cc.y:904
		{
			yyVAL.tyty.t1 = strf
			yyVAL.tyty.t2 = strl
			yyVAL.tyty.t3 = lasttype
			yyVAL.tyty.c = lastclass
			strf = new(Type)
			strl = new(Type)
			lastbit = 0
			firstbit = 1
			lastclass = CXXX
			lasttype = new(Type)
		}
	case 181:
		//line cc.y:917
		{
			yyVAL.ytype = strf
			strf = yyS[yypt-2].tyty.t1
			strl = yyS[yypt-2].tyty.t2
			lasttype = yyS[yypt-2].tyty.t3
			lastclass = yyS[yypt-2].tyty.c
		}
	case 182:
		//line cc.y:926
		{
			lastclass = CXXX
			lasttype = types[TINT]
		}
	case 184:
		//line cc.y:934
		{
			yyVAL.tycl.t = yyS[yypt-0].ytype
			yyVAL.tycl.c = CXXX
		}
	case 185:
		//line cc.y:939
		{
			yyVAL.tycl.t = simplet(yyS[yypt-0].lval)
			yyVAL.tycl.c = CXXX
		}
	case 186:
		//line cc.y:944
		{
			yyVAL.tycl.t = simplet(yyS[yypt-0].lval)
			yyVAL.tycl.c = simplec(yyS[yypt-0].lval)
			yyVAL.tycl.t = garbt(yyVAL.tycl.t, yyS[yypt-0].lval)
		}
	case 187:
		//line cc.y:950
		{
			yyVAL.tycl.t = yyS[yypt-1].ytype
			yyVAL.tycl.c = simplec(yyS[yypt-0].lval)
			yyVAL.tycl.t = garbt(yyVAL.tycl.t, yyS[yypt-0].lval)
			if yyS[yypt-0].lval & ^BCLASS & ^BGARB {
				diag(new(Node), "duplicate types given: %T and %Q", yyS[yypt-1].ytype, yyS[yypt-0].lval)
			}
		}
	case 188:
		//line cc.y:959
		{
			yyVAL.tycl.t = simplet(typebitor(yyS[yypt-1].lval, yyS[yypt-0].lval))
			yyVAL.tycl.c = simplec(yyS[yypt-0].lval)
			yyVAL.tycl.t = garbt(yyVAL.tycl.t, yyS[yypt-0].lval)
		}
	case 189:
		//line cc.y:965
		{
			yyVAL.tycl.t = yyS[yypt-1].ytype
			yyVAL.tycl.c = simplec(yyS[yypt-2].lval)
			yyVAL.tycl.t = garbt(yyVAL.tycl.t, yyS[yypt-2].lval|yyS[yypt-0].lval)
		}
	case 190:
		//line cc.y:971
		{
			yyVAL.tycl.t = simplet(yyS[yypt-0].lval)
			yyVAL.tycl.c = simplec(yyS[yypt-1].lval)
			yyVAL.tycl.t = garbt(yyVAL.tycl.t, yyS[yypt-1].lval)
		}
	case 191:
		//line cc.y:977
		{
			yyVAL.tycl.t = simplet(typebitor(yyS[yypt-1].lval, yyS[yypt-0].lval))
			yyVAL.tycl.c = simplec(yyS[yypt-2].lval|yyS[yypt-0].lval)
			yyVAL.tycl.t = garbt(yyVAL.tycl.t, yyS[yypt-2].lval|yyS[yypt-0].lval)
		}
	case 192:
		//line cc.y:985
		{
			yyVAL.ytype = yyS[yypt-0].tycl.t
			if yyS[yypt-0].tycl.c != CXXX{
				diag(new(Node), "illegal combination of class 4: %s", cnames[yyS[yypt-0].tycl.c])
			}
		}
	case 193:
		//line cc.y:994
		{
			lasttype = yyS[yypt-0].tycl.t
			lastclass = yyS[yypt-0].tycl.c
		}
	case 194:
		//line cc.y:1001
		{
			dotag(yyS[yypt-0].sym, TSTRUCT, 0)
			yyVAL.ytype = yyS[yypt-0].sym.suetag
		}
	case 195:
		//line cc.y:1006
		{
			dotag(yyS[yypt-0].sym, TSTRUCT, autobn)
		}
	case 196:
		//line cc.y:1010
		{
			yyVAL.ytype = yyS[yypt-2].sym.suetag
			if yyVAL.ytype.link != nil{
				diag(new(Node), "redeclare tag: %s", yyS[yypt-2].sym.name)
			}
			yyVAL.ytype.link = yyS[yypt-0].ytype
			sualign(yyVAL.ytype)
		}
	case 197:
		//line cc.y:1019
		{
			taggen++
			sprint(symb, "_%d_", taggen)
			yyVAL.ytype = dotag(lookup(), TSTRUCT, autobn)
			yyVAL.ytype.link = yyS[yypt-0].ytype
			sualign(yyVAL.ytype)
		}
	case 198:
		//line cc.y:1027
		{
			dotag(yyS[yypt-0].sym, TUNION, 0)
			yyVAL.ytype = yyS[yypt-0].sym.suetag
		}
	case 199:
		//line cc.y:1032
		{
			dotag(yyS[yypt-0].sym, TUNION, autobn)
		}
	case 200:
		//line cc.y:1036
		{
			yyVAL.ytype = yyS[yypt-2].sym.suetag
			if yyVAL.ytype.link != nil{
				diag(new(Node), "redeclare tag: %s", yyS[yypt-2].sym.name)
			}
			yyVAL.ytype.link = yyS[yypt-0].ytype
			sualign(yyVAL.ytype)
		}
	case 201:
		//line cc.y:1045
		{
			taggen++
			sprint(symb, "_%d_", taggen)
			yyVAL.ytype = dotag(lookup(), TUNION, autobn)
			yyVAL.ytype.link = yyS[yypt-0].ytype
			sualign(yyVAL.ytype)
		}
	case 202:
		//line cc.y:1053
		{
			dotag(yyS[yypt-0].sym, TENUM, 0)
			yyVAL.ytype = yyS[yypt-0].sym.suetag
			if yyVAL.ytype.link == nil{
				yyVAL.ytype.link = types[TINT]
			}
			yyVAL.ytype = yyVAL.ytype.link
		}
	case 203:
		//line cc.y:1062
		{
			dotag(yyS[yypt-0].sym, TENUM, autobn)
		}
	case 204:
		//line cc.y:1066
		{
			en.tenum = new(Type)
			en.cenum = new(Type)
		}
	case 205:
		//line cc.y:1071
		{
			yyVAL.ytype = yyS[yypt-5].sym.suetag
			if yyVAL.ytype.link != nil{
				diag(new(Node), "redeclare tag: %s", yyS[yypt-5].sym.name)
			}
			if en.tenum == nil{ 
				diag(new(Node), "enum type ambiguous: %s", yyS[yypt-5].sym.name)
				en.tenum = types[TINT]
			}
			yyVAL.ytype.link = en.tenum
			yyVAL.ytype = en.tenum
		}
	case 206:
		//line cc.y:1084
		{
			en.tenum = T
			en.cenum = T
		}
	case 207:
		//line cc.y:1089
		{
			yyVAL.ytype = en.tenum
		}
	case 208:
		//line cc.y:1093
		{
			yyVAL.ytype = tcopy(yyS[yypt-0].sym.Type)
		}
	case 209:
		yyVAL.lval = yyS[yypt-0].lval
	case 210:
		//line cc.y:1100
		{
			yyVAL.lval = typebitor(yyS[yypt-1].lval, yyS[yypt-0].lval)
		}
	case 211:
		//line cc.y:1105
		{
			yyVAL.lval = 0
		}
	case 212:
		//line cc.y:1109
		{
			yyVAL.lval = typebitor(yyS[yypt-1].lval, yyS[yypt-0].lval)
		}
	case 213:
		yyVAL.lval = yyS[yypt-0].lval
	case 214:
		yyVAL.lval = yyS[yypt-0].lval
	case 215:
		yyVAL.lval = yyS[yypt-0].lval
	case 216:
		yyVAL.lval = yyS[yypt-0].lval
	case 217:
		//line cc.y:1121
		{
			yyVAL.lval = typebitor(yyS[yypt-1].lval, yyS[yypt-0].lval)
		}
	case 218:
		yyVAL.lval = yyS[yypt-0].lval
	case 219:
		yyVAL.lval = yyS[yypt-0].lval
	case 220:
		//line cc.y:1131
		{
			doenum(yyS[yypt-0].sym, new(Node))
		}
	case 221:
		//line cc.y:1135
		{
			doenum(yyS[yypt-2].sym, yyS[yypt-0].node)
		}
	case 224:
		//line cc.y:1142
		{ yyVAL.lval = BCHAR }
	case 225:
		//line cc.y:1143
		{ yyVAL.lval = BSHORT }
	case 226:
		//line cc.y:1144
		{ yyVAL.lval = BINT }
	case 227:
		//line cc.y:1145
		{ yyVAL.lval = BLONG }
	case 228:
		//line cc.y:1146
		{ yyVAL.lval = BSIGNED }
	case 229:
		//line cc.y:1147
		{ yyVAL.lval = BUNSIGNED }
	case 230:
		//line cc.y:1148
		{ yyVAL.lval = BFLOAT }
	case 231:
		//line cc.y:1149
		{ yyVAL.lval = BDOUBLE }
	case 232:
		//line cc.y:1150
		{ yyVAL.lval = BVOID }
	case 233:
		//line cc.y:1153
		{ yyVAL.lval = BAUTO }
	case 234:
		//line cc.y:1154
		{ yyVAL.lval = BSTATIC }
	case 235:
		//line cc.y:1155
		{ yyVAL.lval = BEXTERN }
	case 236:
		//line cc.y:1156
		{ yyVAL.lval = BTYPEDEF }
	case 237:
		//line cc.y:1157
		{ yyVAL.lval = BTYPESTR }
	case 238:
		//line cc.y:1158
		{ yyVAL.lval = BREGISTER }
	case 239:
		//line cc.y:1159
		{ yyVAL.lval = 0 }
	case 240:
		//line cc.y:1162
		{ yyVAL.lval = BCONSTNT }
	case 241:
		//line cc.y:1163
		{ yyVAL.lval = BVOLATILE }
	case 242:
		//line cc.y:1164
		{ yyVAL.lval = 0 }
	case 243:
		//line cc.y:1168
		{
			yyVAL.node = Node.new(ONAME, new(Node), new(Node))
			if yyS[yypt-0].sym.class == CLOCAL{
				yyS[yypt-0].sym = mkstatic(yyS[yypt-0].sym)
			}
			yyVAL.node.sym = yyS[yypt-0].sym
			yyVAL.node.Type = yyS[yypt-0].sym.Type
			yyVAL.node.etype = TVOID
			if yyVAL.node.Type != nil{
				yyVAL.node.etype = yyVAL.node.Type.etype
			}
			yyVAL.node.xoffset = yyS[yypt-0].sym.offset
			yyVAL.node.class = yyS[yypt-0].sym.class
			yyS[yypt-0].sym.aused = 1
		}
	case 244:
		//line cc.y:1185
		{
			yyVAL.node = Node.new(ONAME, new(Node), new(Node))
			yyVAL.node.sym = yyS[yypt-0].sym
			yyVAL.node.Type = yyS[yypt-0].sym.Type
			yyVAL.node.etype = TVOID
			if yyVAL.node.Type != nil{
				yyVAL.node.etype = yyVAL.node.Type.etype
			}
			yyVAL.node.xoffset = yyS[yypt-0].sym.offset;
			yyVAL.node.class = yyS[yypt-0].sym.class
		}
	case 245:
		yyVAL.sym = yyS[yypt-0].sym
	case 246:
		yyVAL.sym = yyS[yypt-0].sym
	}
	goto yystack /* stack new state and value */
}
