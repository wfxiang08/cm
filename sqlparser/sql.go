//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "sync"

import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const BEGIN = 57359
const COMMIT = 57360
const ROLLBACK = 57361
const ALL = 57362
const DISTINCT = 57363
const AS = 57364
const EXISTS = 57365
const IN = 57366
const IS = 57367
const LIKE = 57368
const BETWEEN = 57369
const NULL = 57370
const ASC = 57371
const DESC = 57372
const VALUES = 57373
const INTO = 57374
const DUPLICATE = 57375
const KEY = 57376
const DEFAULT = 57377
const SET = 57378
const LOCK = 57379
const ID = 57380
const STRING = 57381
const NUMBER = 57382
const VALUE_ARG = 57383
const LIST_ARG = 57384
const COMMENT = 57385
const LE = 57386
const GE = 57387
const NE = 57388
const NULL_SAFE_EQUAL = 57389
const UNION = 57390
const MINUS = 57391
const EXCEPT = 57392
const INTERSECT = 57393
const JOIN = 57394
const STRAIGHT_JOIN = 57395
const LEFT = 57396
const RIGHT = 57397
const INNER = 57398
const OUTER = 57399
const CROSS = 57400
const NATURAL = 57401
const USE = 57402
const FORCE = 57403
const ON = 57404
const OR = 57405
const AND = 57406
const NOT = 57407
const UNARY = 57408
const CASE = 57409
const WHEN = 57410
const THEN = 57411
const ELSE = 57412
const END = 57413
const CREATE = 57414
const ALTER = 57415
const DROP = 57416
const RENAME = 57417
const ANALYZE = 57418
const TABLE = 57419
const INDEX = 57420
const VIEW = 57421
const TO = 57422
const IGNORE = 57423
const IF = 57424
const UNIQUE = 57425
const USING = 57426
const SHOW = 57427
const DESCRIBE = 57428
const EXPLAIN = 57429

var yyToknames = []string{
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"BEGIN",
	"COMMIT",
	"ROLLBACK",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	" (",
	" =",
	" <",
	" >",
	" ~",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	" ,",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	" &",
	" |",
	" ^",
	" +",
	" -",
	" *",
	" /",
	" %",
	" .",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 209
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 583

var yyAct = []int{

	100, 296, 164, 364, 97, 91, 332, 251, 68, 167,
	98, 288, 242, 86, 202, 213, 69, 166, 3, 141,
	140, 182, 373, 373, 373, 87, 56, 262, 263, 264,
	265, 266, 135, 267, 268, 190, 34, 35, 36, 37,
	82, 234, 294, 135, 71, 135, 234, 76, 70, 74,
	79, 59, 258, 232, 83, 108, 129, 343, 57, 58,
	44, 273, 46, 342, 92, 49, 47, 50, 375, 374,
	372, 341, 317, 75, 313, 315, 125, 323, 322, 78,
	243, 17, 18, 19, 20, 133, 233, 319, 293, 283,
	137, 281, 235, 23, 24, 25, 55, 51, 168, 139,
	163, 165, 169, 126, 314, 122, 128, 52, 53, 54,
	118, 124, 21, 348, 243, 140, 286, 176, 71, 222,
	338, 71, 70, 186, 185, 70, 180, 289, 148, 149,
	150, 151, 152, 153, 154, 155, 254, 184, 92, 208,
	186, 153, 154, 155, 340, 212, 210, 211, 220, 221,
	187, 224, 225, 226, 227, 228, 229, 230, 231, 173,
	200, 207, 223, 22, 26, 28, 27, 29, 141, 140,
	120, 196, 77, 236, 92, 92, 30, 31, 32, 71,
	71, 289, 132, 70, 249, 238, 240, 247, 339, 253,
	206, 255, 134, 194, 311, 246, 197, 141, 140, 215,
	65, 320, 250, 148, 149, 150, 151, 152, 153, 154,
	155, 209, 325, 310, 309, 307, 183, 236, 256, 272,
	308, 276, 277, 259, 274, 148, 149, 150, 151, 152,
	153, 154, 155, 275, 305, 120, 234, 280, 349, 306,
	135, 183, 92, 327, 17, 178, 193, 195, 192, 287,
	121, 205, 81, 282, 285, 115, 291, 179, 295, 292,
	206, 204, 103, 260, 359, 116, 216, 107, 119, 358,
	113, 357, 214, 215, 170, 303, 304, 72, 104, 105,
	106, 321, 151, 152, 153, 154, 155, 95, 120, 324,
	174, 111, 17, 172, 171, 71, 271, 329, 77, 328,
	330, 333, 262, 263, 264, 265, 266, 84, 267, 268,
	94, 138, 270, 370, 109, 110, 206, 206, 72, 318,
	316, 114, 300, 344, 299, 205, 334, 77, 345, 34,
	35, 36, 37, 199, 371, 204, 112, 198, 347, 181,
	236, 130, 354, 353, 356, 127, 123, 355, 66, 85,
	80, 361, 333, 117, 346, 363, 362, 326, 365, 365,
	365, 71, 366, 367, 17, 70, 239, 64, 103, 368,
	279, 377, 188, 107, 378, 131, 113, 62, 379, 217,
	380, 218, 219, 90, 104, 105, 106, 38, 60, 297,
	245, 107, 337, 95, 113, 298, 252, 111, 336, 302,
	183, 72, 104, 105, 106, 67, 40, 41, 42, 43,
	376, 170, 360, 17, 39, 111, 94, 189, 103, 45,
	109, 110, 88, 107, 257, 191, 113, 114, 48, 73,
	248, 177, 369, 90, 104, 105, 106, 350, 109, 110,
	331, 335, 112, 95, 301, 114, 284, 111, 237, 175,
	241, 102, 99, 101, 290, 17, 96, 244, 142, 93,
	112, 312, 203, 261, 201, 89, 94, 269, 103, 136,
	109, 110, 88, 107, 61, 33, 113, 114, 107, 63,
	16, 113, 15, 72, 104, 105, 106, 11, 72, 104,
	105, 106, 112, 95, 10, 9, 14, 111, 170, 13,
	12, 8, 111, 7, 6, 5, 4, 2, 143, 147,
	145, 146, 1, 0, 0, 0, 94, 0, 0, 0,
	109, 110, 351, 352, 0, 109, 110, 114, 159, 160,
	161, 162, 114, 156, 157, 158, 0, 0, 0, 0,
	0, 0, 112, 0, 0, 0, 0, 112, 0, 0,
	0, 0, 0, 0, 0, 144, 148, 149, 150, 151,
	152, 153, 154, 155, 0, 148, 149, 150, 151, 152,
	153, 154, 155, 278, 0, 148, 149, 150, 151, 152,
	153, 154, 155,
}
var yyPact = []int{

	76, -1000, -1000, 276, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -32, -1000, -1000, -1000, -29, 5, 15, 4,
	-1000, -1000, -1000, 408, 368, -1000, -1000, -1000, 356, -1000,
	335, 310, 396, 280, -48, -20, 260, -1000, -13, 260,
	-1000, 312, -57, 260, -57, 311, -1000, -1000, -1000, -1000,
	-1000, 395, -1000, 212, 310, 317, 30, 310, 178, -1000,
	201, -1000, 25, 308, 40, 260, -1000, -1000, 307, -1000,
	-39, 303, 352, 114, 260, -1000, 183, -1000, -1000, 289,
	19, 99, 484, -1000, 445, 239, -1000, -1000, -1000, 363,
	246, 245, -1000, 242, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 363, -1000, 209, 280, 301, 390,
	280, 363, 260, -1000, 349, -64, -1000, 158, -1000, 299,
	-1000, -1000, 295, -1000, 213, 395, -1000, -1000, 260, 134,
	445, 445, 363, 224, 355, 363, 363, 91, 363, 363,
	363, 363, 363, 363, 363, 363, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 484, -50, -17, -11, 484, -1000,
	450, 345, 395, -1000, 408, -3, 153, 359, 280, 280,
	231, -1000, 383, 445, -1000, 153, -1000, -1000, -1000, 68,
	260, -1000, -43, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 206, 244, 274, 287, -19, -1000, -1000, -1000, -1000,
	-1000, 45, 153, -1000, 450, -1000, -1000, 224, 363, 363,
	153, 503, -1000, 342, 207, 207, 207, 64, 64, -1000,
	-1000, -1000, -1000, -1000, 363, -1000, 153, -1000, -12, 395,
	-14, 31, -1000, 445, 59, 226, 276, 113, -15, -1000,
	383, 374, 381, 99, 286, -1000, -1000, 284, -1000, 388,
	213, 213, -1000, -1000, 176, 157, 156, 155, 136, 8,
	-1000, 282, -31, 281, -16, -1000, 153, 131, 363, -1000,
	153, -1000, -25, -1000, -9, -1000, 363, 128, -1000, 324,
	186, -1000, -1000, -1000, 280, 374, -1000, 363, 363, -1000,
	-1000, 386, 378, 244, 52, -1000, 130, -1000, 86, -1000,
	-1000, -1000, -1000, -22, -30, -36, -1000, -1000, -1000, -1000,
	363, 153, -1000, -1000, 153, 363, 320, 226, -1000, -1000,
	56, 181, -1000, 493, -1000, 383, 445, 363, 445, -1000,
	-1000, 223, 221, 216, 153, 153, 405, -1000, 363, 363,
	-1000, -1000, -1000, 374, 99, 179, 99, 260, 260, 260,
	280, 153, -1000, 297, -33, -1000, -34, -35, 178, -1000,
	403, 347, -1000, 260, -1000, -1000, -1000, 260, -1000, 260,
	-1000,
}
var yyPgo = []int{

	0, 512, 507, 17, 506, 505, 504, 503, 501, 500,
	499, 496, 495, 494, 487, 482, 480, 387, 479, 475,
	474, 13, 25, 469, 467, 465, 464, 14, 463, 462,
	200, 461, 3, 21, 5, 459, 458, 457, 456, 2,
	15, 9, 454, 10, 453, 55, 452, 4, 451, 450,
	12, 449, 446, 444, 441, 7, 440, 6, 437, 1,
	432, 431, 430, 11, 8, 16, 252, 429, 428, 425,
	424, 419, 417, 0, 26, 414,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 3, 3, 3, 4,
	4, 5, 6, 7, 8, 8, 8, 12, 13, 14,
	9, 9, 9, 10, 11, 11, 11, 15, 16, 16,
	16, 75, 17, 18, 18, 19, 19, 19, 19, 19,
	20, 20, 21, 21, 22, 22, 22, 25, 25, 23,
	23, 23, 26, 26, 27, 27, 27, 27, 24, 24,
	24, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	29, 29, 29, 30, 30, 31, 31, 31, 31, 32,
	32, 33, 33, 34, 34, 34, 34, 34, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 36, 36,
	36, 36, 36, 36, 36, 40, 40, 40, 45, 41,
	41, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 44, 44,
	46, 46, 46, 48, 51, 51, 49, 49, 50, 52,
	52, 47, 47, 38, 38, 38, 38, 53, 53, 54,
	54, 55, 55, 56, 56, 57, 58, 58, 58, 59,
	59, 59, 60, 60, 60, 61, 61, 62, 62, 63,
	63, 37, 37, 42, 42, 43, 43, 64, 64, 65,
	66, 66, 67, 67, 68, 68, 69, 69, 69, 69,
	69, 70, 70, 71, 71, 72, 72, 73, 74,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 12, 3, 7,
	7, 8, 7, 3, 5, 8, 4, 1, 1, 1,
	6, 7, 4, 5, 4, 5, 5, 3, 2, 2,
	2, 0, 2, 0, 2, 1, 2, 1, 1, 1,
	0, 1, 1, 3, 1, 2, 3, 1, 1, 0,
	1, 2, 1, 3, 3, 3, 3, 5, 0, 1,
	2, 1, 1, 2, 3, 2, 3, 2, 2, 2,
	1, 3, 1, 1, 3, 0, 5, 5, 5, 1,
	3, 0, 2, 1, 3, 3, 2, 3, 3, 3,
	4, 3, 4, 5, 6, 3, 4, 2, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 1, 3, 1,
	3, 1, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 3, 4, 5, 4, 1, 1, 1,
	1, 1, 1, 5, 0, 1, 1, 2, 4, 0,
	2, 1, 3, 1, 1, 1, 1, 0, 3, 0,
	2, 0, 3, 1, 3, 2, 0, 1, 1, 0,
	2, 4, 0, 2, 4, 0, 3, 1, 3, 0,
	5, 2, 1, 1, 3, 3, 1, 1, 3, 3,
	0, 2, 0, 3, 0, 1, 1, 1, 1, 1,
	1, 0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -12,
	-13, -14, -9, -10, -11, -15, -16, 5, 6, 7,
	8, 36, 87, 17, 18, 19, 88, 90, 89, 91,
	100, 101, 102, -19, 53, 54, 55, 56, -17, -75,
	-17, -17, -17, -17, 92, -71, 94, 98, -68, 94,
	96, 92, 92, 93, 94, 92, -74, -74, -74, -3,
	20, -20, 21, -18, 32, -30, 38, 9, -64, -65,
	-47, -73, 38, -67, 97, 93, -73, 38, 92, -73,
	38, -66, 97, -73, -66, 38, -21, -22, 77, -25,
	38, -34, -39, -35, 71, 48, -38, -47, -43, -46,
	-73, -44, -48, 23, 39, 40, 41, 28, -45, 75,
	76, 52, 97, 31, 82, 43, -30, 36, 80, -30,
	57, 49, 80, 38, 71, -73, -74, 38, -74, 95,
	38, 23, 68, -73, 9, 57, -23, -73, 22, 80,
	70, 69, -36, 24, 71, 26, 27, 25, 72, 73,
	74, 75, 76, 77, 78, 79, 49, 50, 51, 44,
	45, 46, 47, -34, -39, -34, -3, -41, -39, -39,
	48, 48, 48, -45, 48, -51, -39, -61, 36, 48,
	-64, 38, -33, 10, -65, -39, -73, -74, 23, -72,
	99, -69, 90, 88, 35, 89, 13, 38, 38, 38,
	-74, -26, -27, -29, 48, 38, -45, -22, -73, 77,
	-34, -34, -39, -40, 48, -45, 42, 24, 26, 27,
	-39, -39, 28, 71, -39, -39, -39, -39, -39, -39,
	-39, -39, 103, 103, 57, 103, -39, 103, -21, 21,
	-21, -49, -50, 83, -37, 31, -3, -64, -62, -47,
	-33, -55, 13, -34, 68, -73, -74, -70, 95, -33,
	57, -28, 58, 59, 60, 61, 62, 64, 65, -24,
	38, 22, -27, 80, -41, -40, -39, -39, 70, 28,
	-39, 103, -21, 103, -52, -50, 85, -34, -63, 68,
	-42, -43, -63, 103, 57, -55, -59, 15, 14, 38,
	38, -53, 11, -27, -27, 58, 63, 58, 63, 58,
	58, 58, -31, 66, 96, 67, 38, 103, 38, 103,
	70, -39, 103, 86, -39, 84, 33, 57, -47, -59,
	-39, -56, -57, -39, -74, -54, 12, 14, 68, 58,
	58, 93, 93, 93, -39, -39, 34, -43, 57, 57,
	-58, 29, 30, -55, -34, -41, -34, 48, 48, 48,
	7, -39, -57, -59, -32, -73, -32, -32, -64, -60,
	16, 37, 103, 57, 103, 103, 7, 24, -73, -73,
	-73,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 41, 41, 41,
	41, 41, 203, 27, 28, 29, 194, 0, 0, 0,
	208, 208, 208, 0, 45, 47, 48, 49, 50, 43,
	0, 0, 0, 0, 192, 0, 0, 204, 0, 0,
	195, 0, 190, 0, 190, 0, 38, 39, 40, 18,
	46, 0, 51, 42, 0, 0, 83, 0, 23, 187,
	0, 151, 207, 0, 0, 0, 208, 207, 0, 208,
	0, 0, 0, 0, 0, 37, 16, 52, 54, 59,
	207, 57, 58, 93, 0, 0, 121, 122, 123, 0,
	151, 0, 137, 0, 153, 154, 155, 156, 186, 140,
	141, 142, 138, 139, 144, 44, 175, 0, 0, 91,
	0, 0, 0, 208, 0, 205, 26, 0, 32, 0,
	34, 191, 0, 208, 0, 0, 55, 60, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 108, 109, 110, 111,
	112, 113, 114, 96, 0, 0, 0, 0, 119, 132,
	0, 0, 0, 107, 0, 0, 145, 0, 0, 0,
	91, 84, 161, 0, 188, 189, 152, 24, 193, 0,
	0, 208, 201, 196, 197, 198, 199, 200, 33, 35,
	36, 91, 62, 68, 0, 80, 82, 53, 61, 56,
	94, 95, 98, 99, 0, 116, 117, 0, 0, 0,
	101, 0, 105, 0, 124, 125, 126, 127, 128, 129,
	130, 131, 97, 118, 0, 185, 119, 133, 0, 0,
	0, 149, 146, 0, 179, 0, 182, 179, 0, 177,
	161, 169, 0, 92, 0, 206, 30, 0, 202, 157,
	0, 0, 71, 72, 0, 0, 0, 0, 0, 85,
	69, 0, 0, 0, 0, 100, 102, 0, 0, 106,
	120, 134, 0, 136, 0, 147, 0, 0, 19, 0,
	181, 183, 20, 176, 0, 169, 22, 0, 0, 208,
	31, 159, 0, 63, 66, 73, 0, 75, 0, 77,
	78, 79, 64, 0, 0, 0, 70, 65, 81, 115,
	0, 103, 135, 143, 150, 0, 0, 0, 178, 21,
	170, 162, 163, 166, 25, 161, 0, 0, 0, 74,
	76, 0, 0, 0, 104, 148, 0, 184, 0, 0,
	165, 167, 168, 169, 160, 158, 67, 0, 0, 0,
	0, 171, 164, 172, 0, 89, 0, 0, 180, 17,
	0, 0, 86, 0, 87, 88, 173, 0, 90, 0,
	174,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 79, 72, 3,
	48, 103, 77, 75, 57, 76, 80, 78, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	50, 49, 51, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 74, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 73, 3, 52,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 53, 54, 55, 56,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/
var yyPool = sync.Pool{
	New: func() interface{} {
		return make([]yySymType, yyMaxDepth)
	},
}

func allocyyS() []yySymType {
	return yyPool.Get().([]yySymType)
}

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
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var flag bool
	yyS := allocyyS()
	yysOrg := yyS

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	yyPool.Put(yysOrg)
	return 0

ret1:
	yyPool.Put(yysOrg)
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
	if flag {
		yyS[yyp] = yylval
		flag = false
	} else {
		yyS[yyp] = yyVAL
	}
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
		//yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		flag = true
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
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
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

	case 1:
		//line sql.y:153
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:159
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		yyVAL.statement = yyS[yypt-0].statement
	case 4:
		yyVAL.statement = yyS[yypt-0].statement
	case 5:
		yyVAL.statement = yyS[yypt-0].statement
	case 6:
		yyVAL.statement = yyS[yypt-0].statement
	case 7:
		yyVAL.statement = yyS[yypt-0].statement
	case 8:
		yyVAL.statement = yyS[yypt-0].statement
	case 9:
		yyVAL.statement = yyS[yypt-0].statement
	case 10:
		yyVAL.statement = yyS[yypt-0].statement
	case 11:
		yyVAL.statement = yyS[yypt-0].statement
	case 12:
		yyVAL.statement = yyS[yypt-0].statement
	case 13:
		yyVAL.statement = yyS[yypt-0].statement
	case 14:
		yyVAL.statement = yyS[yypt-0].statement
	case 15:
		yyVAL.statement = yyS[yypt-0].statement
	case 16:
		//line sql.y:178
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 17:
		//line sql.y:182
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 18:
		//line sql.y:186
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 19:
		//line sql.y:192
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 20:
		//line sql.y:196
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:208
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 22:
		//line sql.y:214
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 23:
		//line sql.y:220
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 24:
		//line sql.y:226
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 25:
		//line sql.y:230
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 26:
		//line sql.y:235
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 27:
		//line sql.y:241
		{
			yyVAL.statement = &Begin{}
		}
	case 28:
		//line sql.y:247
		{
			yyVAL.statement = &Commit{}
		}
	case 29:
		//line sql.y:253
		{
			yyVAL.statement = &Rollback{}
		}
	case 30:
		//line sql.y:260
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 31:
		//line sql.y:264
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 32:
		//line sql.y:269
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 33:
		//line sql.y:275
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 34:
		//line sql.y:281
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 35:
		//line sql.y:285
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 36:
		//line sql.y:290
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 37:
		//line sql.y:296
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 38:
		//line sql.y:302
		{
			yyVAL.statement = &Other{}
		}
	case 39:
		//line sql.y:306
		{
			yyVAL.statement = &Other{}
		}
	case 40:
		//line sql.y:310
		{
			yyVAL.statement = &Other{}
		}
	case 41:
		//line sql.y:315
		{
			SetAllowComments(yylex, true)
		}
	case 42:
		//line sql.y:319
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 43:
		//line sql.y:325
		{
			yyVAL.bytes2 = nil
		}
	case 44:
		//line sql.y:329
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 45:
		//line sql.y:335
		{
			yyVAL.str = AST_UNION
		}
	case 46:
		//line sql.y:339
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 47:
		//line sql.y:343
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 48:
		//line sql.y:347
		{
			yyVAL.str = AST_EXCEPT
		}
	case 49:
		//line sql.y:351
		{
			yyVAL.str = AST_INTERSECT
		}
	case 50:
		//line sql.y:356
		{
			yyVAL.str = ""
		}
	case 51:
		//line sql.y:360
		{
			yyVAL.str = AST_DISTINCT
		}
	case 52:
		//line sql.y:366
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 53:
		//line sql.y:370
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 54:
		//line sql.y:376
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 55:
		//line sql.y:380
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 56:
		//line sql.y:384
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 57:
		//line sql.y:390
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 58:
		//line sql.y:394
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 59:
		//line sql.y:399
		{
			yyVAL.bytes = nil
		}
	case 60:
		//line sql.y:403
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 61:
		//line sql.y:407
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 62:
		//line sql.y:413
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 63:
		//line sql.y:417
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 64:
		//line sql.y:423
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 65:
		//line sql.y:427
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 66:
		//line sql.y:431
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 67:
		//line sql.y:435
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 68:
		//line sql.y:440
		{
			yyVAL.bytes = nil
		}
	case 69:
		//line sql.y:444
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 70:
		//line sql.y:448
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 71:
		//line sql.y:454
		{
			yyVAL.str = AST_JOIN
		}
	case 72:
		//line sql.y:458
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 73:
		//line sql.y:462
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 74:
		//line sql.y:466
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 75:
		//line sql.y:470
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 76:
		//line sql.y:474
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 77:
		//line sql.y:478
		{
			yyVAL.str = AST_JOIN
		}
	case 78:
		//line sql.y:482
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 79:
		//line sql.y:486
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 80:
		//line sql.y:492
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 81:
		//line sql.y:496
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 82:
		//line sql.y:500
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 83:
		//line sql.y:506
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 84:
		//line sql.y:510
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 85:
		//line sql.y:515
		{
			yyVAL.indexHints = nil
		}
	case 86:
		//line sql.y:519
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 87:
		//line sql.y:523
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 88:
		//line sql.y:527
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 89:
		//line sql.y:533
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 90:
		//line sql.y:537
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 91:
		//line sql.y:542
		{
			yyVAL.boolExpr = nil
		}
	case 92:
		//line sql.y:546
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 93:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 94:
		//line sql.y:553
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 95:
		//line sql.y:557
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 96:
		//line sql.y:561
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 97:
		//line sql.y:565
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 98:
		//line sql.y:571
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 99:
		//line sql.y:575
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].colTuple}
		}
	case 100:
		//line sql.y:579
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].colTuple}
		}
	case 101:
		//line sql.y:583
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 102:
		//line sql.y:587
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 103:
		//line sql.y:591
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 104:
		//line sql.y:595
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 105:
		//line sql.y:599
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 106:
		//line sql.y:603
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 107:
		//line sql.y:607
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 108:
		//line sql.y:613
		{
			yyVAL.str = AST_EQ
		}
	case 109:
		//line sql.y:617
		{
			yyVAL.str = AST_LT
		}
	case 110:
		//line sql.y:621
		{
			yyVAL.str = AST_GT
		}
	case 111:
		//line sql.y:625
		{
			yyVAL.str = AST_LE
		}
	case 112:
		//line sql.y:629
		{
			yyVAL.str = AST_GE
		}
	case 113:
		//line sql.y:633
		{
			yyVAL.str = AST_NE
		}
	case 114:
		//line sql.y:637
		{
			yyVAL.str = AST_NSE
		}
	case 115:
		//line sql.y:643
		{
			yyVAL.colTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 116:
		//line sql.y:647
		{
			yyVAL.colTuple = yyS[yypt-0].subquery
		}
	case 117:
		//line sql.y:651
		{
			yyVAL.colTuple = ListArg(yyS[yypt-0].bytes)
		}
	case 118:
		//line sql.y:657
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 119:
		//line sql.y:663
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 120:
		//line sql.y:667
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 121:
		//line sql.y:673
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 122:
		//line sql.y:677
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 123:
		//line sql.y:681
		{
			yyVAL.valExpr = yyS[yypt-0].rowTuple
		}
	case 124:
		//line sql.y:685
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 125:
		//line sql.y:689
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 126:
		//line sql.y:693
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 127:
		//line sql.y:697
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 128:
		//line sql.y:701
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 129:
		//line sql.y:705
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:709
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 131:
		//line sql.y:713
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:717
		{
			if num, ok := yyS[yypt-0].valExpr.(NumVal); ok {
				switch yyS[yypt-1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
			}
		}
	case 133:
		//line sql.y:732
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 134:
		//line sql.y:736
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 135:
		//line sql.y:740
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 136:
		//line sql.y:744
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 137:
		//line sql.y:748
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 138:
		//line sql.y:754
		{
			yyVAL.bytes = IF_BYTES
		}
	case 139:
		//line sql.y:758
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 140:
		//line sql.y:764
		{
			yyVAL.byt = AST_UPLUS
		}
	case 141:
		//line sql.y:768
		{
			yyVAL.byt = AST_UMINUS
		}
	case 142:
		//line sql.y:772
		{
			yyVAL.byt = AST_TILDA
		}
	case 143:
		//line sql.y:778
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 144:
		//line sql.y:783
		{
			yyVAL.valExpr = nil
		}
	case 145:
		//line sql.y:787
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 146:
		//line sql.y:793
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 147:
		//line sql.y:797
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 148:
		//line sql.y:803
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 149:
		//line sql.y:808
		{
			yyVAL.valExpr = nil
		}
	case 150:
		//line sql.y:812
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 151:
		//line sql.y:818
		{
			yyVAL.colName = &ColName{Name: bytes.ToLower(yyS[yypt-0].bytes)}
		}
	case 152:
		//line sql.y:822
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: bytes.ToLower(yyS[yypt-0].bytes)}
		}
	case 153:
		//line sql.y:828
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 154:
		//line sql.y:832
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 155:
		//line sql.y:836
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 156:
		//line sql.y:840
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 157:
		//line sql.y:845
		{
			yyVAL.valExprs = nil
		}
	case 158:
		//line sql.y:849
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 159:
		//line sql.y:854
		{
			yyVAL.boolExpr = nil
		}
	case 160:
		//line sql.y:858
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 161:
		//line sql.y:863
		{
			yyVAL.orderBy = nil
		}
	case 162:
		//line sql.y:867
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 163:
		//line sql.y:873
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 164:
		//line sql.y:877
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 165:
		//line sql.y:883
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 166:
		//line sql.y:888
		{
			yyVAL.str = AST_ASC
		}
	case 167:
		//line sql.y:892
		{
			yyVAL.str = AST_ASC
		}
	case 168:
		//line sql.y:896
		{
			yyVAL.str = AST_DESC
		}
	case 169:
		//line sql.y:901
		{
			yyVAL.limit = nil
		}
	case 170:
		//line sql.y:905
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 171:
		//line sql.y:909
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 172:
		//line sql.y:914
		{
			yyVAL.str = ""
		}
	case 173:
		//line sql.y:918
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 174:
		//line sql.y:922
		{
			if !bytes.Equal(yyS[yypt-1].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyS[yypt-0].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 175:
		//line sql.y:935
		{
			yyVAL.columns = nil
		}
	case 176:
		//line sql.y:939
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 177:
		//line sql.y:945
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 178:
		//line sql.y:949
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 179:
		//line sql.y:954
		{
			yyVAL.updateExprs = nil
		}
	case 180:
		//line sql.y:958
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 181:
		//line sql.y:964
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 182:
		//line sql.y:968
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 183:
		//line sql.y:974
		{
			yyVAL.values = Values{yyS[yypt-0].rowTuple}
		}
	case 184:
		//line sql.y:978
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].rowTuple)
		}
	case 185:
		//line sql.y:984
		{
			yyVAL.rowTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 186:
		//line sql.y:988
		{
			yyVAL.rowTuple = yyS[yypt-0].subquery
		}
	case 187:
		//line sql.y:994
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 188:
		//line sql.y:998
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 189:
		//line sql.y:1004
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 190:
		//line sql.y:1009
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		//line sql.y:1011
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:1014
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:1016
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1019
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1021
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1025
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1027
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1029
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1031
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1033
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1036
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1038
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1041
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1043
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1046
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1048
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1052
		{
			//comment by liuqi    $$ = bytes.ToLower($1)
			yyVAL.bytes = yyS[yypt-0].bytes //add by liuqi
		}
	case 208:
		//line sql.y:1058
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
