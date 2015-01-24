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
const REPLACE = 57351
const FROM = 57352
const WHERE = 57353
const GROUP = 57354
const HAVING = 57355
const ORDER = 57356
const BY = 57357
const LIMIT = 57358
const FOR = 57359
const BEGIN = 57360
const COMMIT = 57361
const ROLLBACK = 57362
const START = 57363
const TRANSACTION = 57364
const ALL = 57365
const DISTINCT = 57366
const AS = 57367
const EXISTS = 57368
const IN = 57369
const IS = 57370
const LIKE = 57371
const BETWEEN = 57372
const NULL = 57373
const ASC = 57374
const DESC = 57375
const VALUES = 57376
const INTO = 57377
const DUPLICATE = 57378
const KEY = 57379
const DEFAULT = 57380
const SET = 57381
const LOCK = 57382
const ID = 57383
const STRING = 57384
const NUMBER = 57385
const VALUE_ARG = 57386
const LIST_ARG = 57387
const COMMENT = 57388
const LE = 57389
const GE = 57390
const NE = 57391
const NULL_SAFE_EQUAL = 57392
const NAMES = 57393
const GLOBAL = 57394
const SESSION = 57395
const UNION = 57396
const MINUS = 57397
const EXCEPT = 57398
const INTERSECT = 57399
const JOIN = 57400
const STRAIGHT_JOIN = 57401
const LEFT = 57402
const RIGHT = 57403
const INNER = 57404
const OUTER = 57405
const CROSS = 57406
const NATURAL = 57407
const USE = 57408
const FORCE = 57409
const ON = 57410
const OR = 57411
const AND = 57412
const NOT = 57413
const UNARY = 57414
const CASE = 57415
const WHEN = 57416
const THEN = 57417
const ELSE = 57418
const END = 57419
const CREATE = 57420
const ALTER = 57421
const DROP = 57422
const RENAME = 57423
const ANALYZE = 57424
const TABLE = 57425
const INDEX = 57426
const VIEW = 57427
const TO = 57428
const IGNORE = 57429
const IF = 57430
const UNIQUE = 57431
const USING = 57432
const SHOW = 57433
const DESCRIBE = 57434
const EXPLAIN = 57435

var yyToknames = []string{
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"REPLACE",
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
	"START",
	"TRANSACTION",
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
	"NAMES",
	"GLOBAL",
	"SESSION",
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

const yyNprod = 216
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 623

var yyAct = []int{

	109, 180, 177, 383, 106, 351, 74, 268, 107, 61,
	315, 305, 197, 217, 257, 228, 78, 95, 259, 179,
	3, 190, 205, 96, 392, 117, 279, 280, 281, 282,
	283, 91, 284, 285, 154, 153, 100, 37, 38, 39,
	40, 83, 392, 392, 62, 63, 148, 249, 80, 275,
	311, 85, 79, 142, 148, 88, 64, 148, 249, 92,
	54, 48, 55, 50, 57, 58, 59, 51, 247, 101,
	394, 336, 332, 334, 362, 361, 80, 80, 342, 87,
	79, 79, 131, 132, 360, 138, 60, 248, 393, 391,
	258, 84, 341, 338, 146, 139, 310, 211, 141, 150,
	300, 56, 333, 298, 250, 154, 153, 181, 290, 152,
	339, 182, 161, 162, 163, 164, 165, 166, 167, 168,
	344, 209, 137, 135, 212, 258, 189, 303, 80, 128,
	237, 80, 79, 153, 195, 79, 201, 200, 186, 86,
	176, 178, 166, 167, 168, 357, 202, 199, 193, 154,
	153, 101, 223, 201, 306, 271, 215, 359, 227, 145,
	358, 235, 236, 198, 239, 240, 241, 242, 243, 244,
	245, 246, 222, 221, 370, 371, 238, 208, 210, 207,
	326, 224, 230, 130, 130, 327, 251, 101, 101, 330,
	225, 226, 80, 80, 306, 80, 79, 264, 262, 79,
	324, 266, 253, 255, 329, 325, 272, 328, 267, 198,
	261, 249, 265, 261, 368, 277, 273, 18, 346, 134,
	161, 162, 163, 164, 165, 166, 167, 168, 90, 276,
	18, 291, 251, 289, 147, 270, 293, 294, 112, 37,
	38, 39, 40, 116, 231, 221, 122, 71, 292, 220,
	229, 378, 297, 81, 113, 114, 115, 101, 230, 219,
	377, 130, 376, 104, 367, 183, 220, 120, 187, 308,
	185, 302, 299, 184, 309, 314, 219, 312, 313, 161,
	162, 163, 164, 165, 166, 167, 168, 148, 93, 103,
	194, 322, 323, 118, 119, 304, 124, 86, 340, 389,
	123, 191, 192, 221, 221, 81, 343, 164, 165, 166,
	167, 168, 80, 192, 81, 121, 347, 125, 126, 349,
	352, 129, 390, 127, 288, 348, 337, 151, 353, 77,
	75, 76, 335, 319, 318, 18, 19, 21, 22, 20,
	287, 214, 363, 86, 213, 196, 143, 364, 25, 27,
	28, 26, 140, 136, 133, 366, 72, 94, 374, 251,
	89, 365, 372, 345, 70, 296, 69, 18, 396, 23,
	380, 352, 67, 203, 381, 144, 65, 384, 384, 384,
	80, 385, 386, 382, 79, 52, 387, 316, 254, 356,
	112, 317, 373, 397, 375, 116, 260, 398, 122, 399,
	232, 269, 233, 234, 355, 99, 113, 114, 115, 321,
	279, 280, 281, 282, 283, 104, 284, 285, 198, 120,
	73, 395, 379, 24, 29, 31, 30, 32, 161, 162,
	163, 164, 165, 166, 167, 168, 33, 34, 35, 18,
	42, 103, 204, 112, 49, 118, 119, 97, 116, 274,
	206, 122, 123, 53, 82, 263, 388, 369, 99, 113,
	114, 115, 350, 354, 320, 41, 112, 121, 104, 301,
	188, 116, 120, 252, 122, 256, 111, 18, 108, 110,
	307, 81, 113, 114, 115, 43, 44, 45, 46, 47,
	105, 104, 155, 102, 103, 120, 331, 218, 118, 119,
	97, 278, 216, 116, 98, 123, 122, 286, 149, 66,
	36, 68, 17, 81, 113, 114, 115, 103, 16, 12,
	121, 118, 119, 183, 11, 10, 116, 120, 123, 122,
	15, 14, 13, 9, 8, 5, 81, 113, 114, 115,
	7, 6, 4, 121, 2, 1, 183, 0, 0, 0,
	120, 0, 0, 118, 119, 0, 0, 0, 0, 0,
	123, 0, 0, 0, 156, 160, 158, 159, 0, 0,
	0, 0, 0, 0, 0, 121, 118, 119, 0, 0,
	0, 0, 0, 123, 172, 173, 174, 175, 0, 169,
	170, 171, 0, 0, 0, 0, 0, 295, 121, 161,
	162, 163, 164, 165, 166, 167, 168, 0, 0, 0,
	0, 0, 0, 0, 157, 161, 162, 163, 164, 165,
	166, 167, 168,
}
var yyPact = []int{

	330, -1000, -1000, 180, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -37, -1000, 363, -1000, -1000, -40,
	3, -34, -12, -1000, -1000, -1000, 434, 353, -1000, -1000,
	-1000, 348, -1000, 331, 329, 315, 410, 273, -62, -8,
	256, -1000, -1000, -19, 256, -1000, 319, -72, 256, -72,
	316, -1000, -1000, -1000, -1000, -1000, 417, -1000, 250, 315,
	315, 284, 43, 315, 121, 264, 264, 313, -1000, 167,
	-1000, 37, 312, 45, 256, -1000, -1000, 311, -1000, -48,
	305, 349, 85, 256, -1000, 224, -1000, -1000, 302, 23,
	74, 537, -1000, 440, 212, -1000, -1000, -1000, 495, 222,
	219, -1000, 217, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 495, -1000, 262, 251, 264, 304, 407,
	264, 121, 121, -1000, 495, 256, -1000, 347, -83, -1000,
	83, -1000, 303, -1000, -1000, 300, -1000, 208, 417, -1000,
	-1000, 256, 98, 440, 440, 495, 199, 373, 495, 495,
	99, 495, 495, 495, 495, 495, 495, 495, 495, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 537, -41, -22,
	-5, 537, -1000, 472, 364, 417, -1000, 434, 1, 350,
	362, 264, 264, 362, 264, 198, -1000, 387, 440, -1000,
	350, -1000, -1000, -1000, 81, 256, -1000, -52, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 152, 346, 299, 225,
	22, -1000, -1000, -1000, -1000, -1000, 57, 350, -1000, 472,
	-1000, -1000, 199, 495, 495, 350, 521, -1000, 334, 226,
	226, 226, 59, 59, -1000, -1000, -1000, -1000, -1000, 495,
	-1000, 350, -1000, -6, 417, -9, 36, -1000, 440, 80,
	214, 180, 120, -13, -1000, 80, 120, 387, 371, 376,
	74, 293, -1000, -1000, 292, -1000, 397, 208, 208, -1000,
	-1000, 136, 116, 143, 140, 125, 0, -1000, 291, -38,
	285, -16, -1000, 350, 34, 495, -1000, 350, -1000, -17,
	-1000, -14, -1000, 495, 30, -1000, 327, 155, -1000, -1000,
	-1000, 264, -1000, -1000, 371, -1000, 495, 495, -1000, -1000,
	391, 374, 346, 71, -1000, 96, -1000, 93, -1000, -1000,
	-1000, -1000, -15, -24, -25, -1000, -1000, -1000, -1000, 495,
	350, -1000, -1000, 350, 495, 324, 214, -1000, -1000, 201,
	151, -1000, 142, -1000, 387, 440, 495, 440, -1000, -1000,
	211, 209, 200, 350, 350, 415, -1000, 495, 495, -1000,
	-1000, -1000, 371, 74, 148, 74, 256, 256, 256, 264,
	350, -1000, 282, -20, -1000, -21, -39, 121, -1000, 414,
	341, -1000, 256, -1000, -1000, -1000, 256, -1000, 256, -1000,
}
var yyPgo = []int{

	0, 545, 544, 19, 542, 541, 540, 535, 534, 533,
	532, 531, 530, 525, 524, 519, 518, 512, 465, 511,
	510, 509, 17, 23, 508, 507, 504, 502, 13, 501,
	497, 247, 496, 3, 12, 36, 493, 492, 18, 490,
	2, 15, 1, 480, 8, 479, 25, 478, 4, 476,
	475, 14, 470, 469, 464, 463, 7, 462, 5, 457,
	10, 456, 21, 455, 11, 6, 16, 228, 454, 453,
	450, 449, 444, 442, 0, 9, 440,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 7, 7, 5, 6, 8, 8, 8, 8,
	9, 9, 9, 13, 13, 14, 15, 10, 10, 10,
	11, 12, 12, 12, 16, 17, 17, 17, 76, 18,
	19, 19, 20, 20, 20, 20, 20, 21, 21, 22,
	22, 23, 23, 23, 26, 26, 24, 24, 24, 27,
	27, 28, 28, 28, 28, 25, 25, 25, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 30, 30, 30,
	31, 31, 32, 32, 32, 32, 33, 33, 34, 34,
	35, 35, 35, 35, 35, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 37, 37, 37, 37, 37,
	37, 37, 41, 41, 41, 46, 42, 42, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 45, 45, 47, 47, 47,
	49, 52, 52, 50, 50, 51, 53, 53, 48, 48,
	39, 39, 39, 39, 54, 54, 55, 55, 56, 56,
	57, 57, 58, 59, 59, 59, 60, 60, 60, 61,
	61, 61, 62, 62, 63, 63, 64, 64, 38, 38,
	43, 43, 44, 44, 65, 65, 66, 67, 67, 68,
	68, 69, 69, 70, 70, 70, 70, 70, 71, 71,
	72, 72, 73, 73, 74, 75,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 7, 7, 8, 7, 3, 4, 4, 4,
	5, 8, 4, 1, 2, 1, 1, 6, 7, 4,
	5, 4, 5, 5, 3, 2, 2, 2, 0, 2,
	0, 2, 1, 2, 1, 1, 1, 0, 1, 1,
	3, 1, 2, 3, 1, 1, 0, 1, 2, 1,
	3, 3, 3, 3, 5, 0, 1, 2, 1, 1,
	2, 3, 2, 3, 2, 2, 2, 1, 3, 1,
	1, 3, 0, 5, 5, 5, 1, 3, 0, 2,
	1, 3, 3, 2, 3, 3, 3, 4, 3, 4,
	5, 6, 3, 4, 2, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 1, 3, 1, 3, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	3, 4, 5, 4, 1, 1, 1, 1, 1, 1,
	5, 0, 1, 1, 2, 4, 0, 2, 1, 3,
	1, 1, 1, 1, 0, 3, 0, 2, 0, 3,
	1, 3, 2, 0, 1, 1, 0, 2, 4, 0,
	2, 4, 0, 3, 1, 3, 0, 5, 2, 1,
	1, 3, 3, 1, 1, 3, 3, 0, 2, 0,
	3, 0, 1, 1, 1, 1, 1, 1, 0, 1,
	0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -7, -5, -6, -8, -9,
	-13, -14, -15, -10, -11, -12, -16, -17, 5, 6,
	9, 7, 8, 39, 93, 18, 21, 19, 20, 94,
	96, 95, 97, 106, 107, 108, -20, 59, 60, 61,
	62, -18, -76, -18, -18, -18, -18, -18, 98, -72,
	100, 104, 22, -69, 100, 102, 98, 98, 99, 100,
	98, -75, -75, -75, -3, 23, -21, 24, -19, 35,
	35, -31, 41, 10, -65, 57, 58, 56, -66, -48,
	-74, 41, -68, 103, 99, -74, 41, 98, -74, 41,
	-67, 103, -74, -67, 41, -22, -23, 83, -26, 41,
	-35, -40, -36, 77, 51, -39, -48, -44, -47, -74,
	-45, -49, 26, 42, 43, 44, 31, -46, 81, 82,
	55, 103, 34, 88, 46, -31, -31, 39, 86, -31,
	63, -65, -65, 41, 52, 86, 41, 77, -74, -75,
	41, -75, 101, 41, 26, 74, -74, 10, 63, -24,
	-74, 25, 86, 76, 75, -37, 27, 77, 29, 30,
	28, 78, 79, 80, 81, 82, 83, 84, 85, 52,
	53, 54, 47, 48, 49, 50, -35, -40, -35, -3,
	-42, -40, -40, 51, 51, 51, -46, 51, -52, -40,
	-62, 39, 51, -62, 39, -65, 41, -34, 11, -66,
	-40, -74, -75, 26, -73, 105, -70, 96, 94, 38,
	95, 14, 41, 41, 41, -75, -27, -28, -30, 51,
	41, -46, -23, -74, 83, -35, -35, -40, -41, 51,
	-46, 45, 27, 29, 30, -40, -40, 31, 77, -40,
	-40, -40, -40, -40, -40, -40, -40, 109, 109, 63,
	109, -40, 109, -22, 24, -22, -50, -51, 89, -38,
	34, -3, -65, -63, -48, -38, -65, -34, -56, 14,
	-35, 74, -74, -75, -71, 101, -34, 63, -29, 64,
	65, 66, 67, 68, 70, 71, -25, 41, 25, -28,
	86, -42, -41, -40, -40, 76, 31, -40, 109, -22,
	109, -53, -51, 91, -35, -64, 74, -43, -44, -64,
	109, 63, -64, -64, -56, -60, 16, 15, 41, 41,
	-54, 12, -28, -28, 64, 69, 64, 69, 64, 64,
	64, -32, 72, 102, 73, 41, 109, 41, 109, 76,
	-40, 109, 92, -40, 90, 36, 63, -48, -60, -40,
	-57, -58, -40, -75, -55, 13, 15, 74, 64, 64,
	99, 99, 99, -40, -40, 37, -44, 63, 63, -59,
	32, 33, -56, -35, -42, -35, 51, 51, 51, 7,
	-40, -58, -60, -33, -74, -33, -33, -65, -61, 17,
	40, 109, 63, 109, 109, 7, 27, -74, -74, -74,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 48, 48,
	48, 48, 48, 48, 210, 33, 0, 35, 36, 201,
	0, 0, 0, 215, 215, 215, 0, 52, 54, 55,
	56, 57, 50, 0, 0, 0, 0, 0, 199, 0,
	0, 211, 34, 0, 0, 202, 0, 197, 0, 197,
	0, 45, 46, 47, 19, 53, 0, 58, 49, 0,
	0, 0, 90, 0, 26, 0, 0, 0, 194, 0,
	158, 214, 0, 0, 0, 215, 214, 0, 215, 0,
	0, 0, 0, 0, 44, 17, 59, 61, 66, 214,
	64, 65, 100, 0, 0, 128, 129, 130, 0, 158,
	0, 144, 0, 160, 161, 162, 163, 193, 147, 148,
	149, 145, 146, 151, 51, 182, 182, 0, 0, 98,
	0, 27, 28, 29, 0, 0, 215, 0, 212, 32,
	0, 39, 0, 41, 198, 0, 215, 0, 0, 62,
	67, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 115,
	116, 117, 118, 119, 120, 121, 103, 0, 0, 0,
	0, 126, 139, 0, 0, 0, 114, 0, 0, 152,
	0, 0, 0, 0, 0, 98, 91, 168, 0, 195,
	196, 159, 30, 200, 0, 0, 215, 208, 203, 204,
	205, 206, 207, 40, 42, 43, 98, 69, 75, 0,
	87, 89, 60, 68, 63, 101, 102, 105, 106, 0,
	123, 124, 0, 0, 0, 108, 0, 112, 0, 131,
	132, 133, 134, 135, 136, 137, 138, 104, 125, 0,
	192, 126, 140, 0, 0, 0, 156, 153, 0, 186,
	0, 189, 186, 0, 184, 186, 186, 168, 176, 0,
	99, 0, 213, 37, 0, 209, 164, 0, 0, 78,
	79, 0, 0, 0, 0, 0, 92, 76, 0, 0,
	0, 0, 107, 109, 0, 0, 113, 127, 141, 0,
	143, 0, 154, 0, 0, 20, 0, 188, 190, 21,
	183, 0, 22, 23, 176, 25, 0, 0, 215, 38,
	166, 0, 70, 73, 80, 0, 82, 0, 84, 85,
	86, 71, 0, 0, 0, 77, 72, 88, 122, 0,
	110, 142, 150, 157, 0, 0, 0, 185, 24, 177,
	169, 170, 173, 31, 168, 0, 0, 0, 81, 83,
	0, 0, 0, 111, 155, 0, 191, 0, 0, 172,
	174, 175, 176, 167, 165, 74, 0, 0, 0, 0,
	178, 171, 179, 0, 96, 0, 0, 187, 18, 0,
	0, 93, 0, 94, 95, 180, 0, 97, 0, 181,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 85, 78, 3,
	51, 109, 83, 81, 63, 82, 86, 84, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	53, 52, 54, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 80, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 79, 3, 55,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 56,
	57, 58, 59, 60, 61, 62, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 73, 74, 75, 76, 77,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108,
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
		//line sql.y:157
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:163
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
		yyVAL.statement = yyS[yypt-0].statement
	case 17:
		//line sql.y:183
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 18:
		//line sql.y:187
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 19:
		//line sql.y:191
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 20:
		//line sql.y:197
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:201
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 22:
		//line sql.y:214
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 23:
		//line sql.y:218
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 24:
		//line sql.y:231
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 25:
		//line sql.y:237
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 26:
		//line sql.y:243
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 27:
		//line sql.y:247
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: yyS[yypt-0].updateExprs, Scope: "global"}
		}
	case 28:
		//line sql.y:251
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: yyS[yypt-0].updateExprs, Scope: "session"}
		}
	case 29:
		//line sql.y:255
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 30:
		//line sql.y:261
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 31:
		//line sql.y:265
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 32:
		//line sql.y:270
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 33:
		//line sql.y:276
		{
			yyVAL.statement = &Begin{}
		}
	case 34:
		//line sql.y:280
		{
			yyVAL.statement = &Begin{}
		}
	case 35:
		//line sql.y:286
		{
			yyVAL.statement = &Commit{}
		}
	case 36:
		//line sql.y:292
		{
			yyVAL.statement = &Rollback{}
		}
	case 37:
		//line sql.y:299
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 38:
		//line sql.y:303
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 39:
		//line sql.y:308
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 40:
		//line sql.y:314
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 41:
		//line sql.y:320
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 42:
		//line sql.y:324
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 43:
		//line sql.y:329
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 44:
		//line sql.y:335
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 45:
		//line sql.y:341
		{
			yyVAL.statement = &Other{}
		}
	case 46:
		//line sql.y:345
		{
			yyVAL.statement = &Other{}
		}
	case 47:
		//line sql.y:349
		{
			yyVAL.statement = &Other{}
		}
	case 48:
		//line sql.y:354
		{
			SetAllowComments(yylex, true)
		}
	case 49:
		//line sql.y:358
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 50:
		//line sql.y:364
		{
			yyVAL.bytes2 = nil
		}
	case 51:
		//line sql.y:368
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 52:
		//line sql.y:374
		{
			yyVAL.str = AST_UNION
		}
	case 53:
		//line sql.y:378
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 54:
		//line sql.y:382
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 55:
		//line sql.y:386
		{
			yyVAL.str = AST_EXCEPT
		}
	case 56:
		//line sql.y:390
		{
			yyVAL.str = AST_INTERSECT
		}
	case 57:
		//line sql.y:395
		{
			yyVAL.str = ""
		}
	case 58:
		//line sql.y:399
		{
			yyVAL.str = AST_DISTINCT
		}
	case 59:
		//line sql.y:405
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 60:
		//line sql.y:409
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 61:
		//line sql.y:415
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 62:
		//line sql.y:419
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 63:
		//line sql.y:423
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 64:
		//line sql.y:429
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 65:
		//line sql.y:433
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 66:
		//line sql.y:438
		{
			yyVAL.bytes = nil
		}
	case 67:
		//line sql.y:442
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 68:
		//line sql.y:446
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 69:
		//line sql.y:452
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 70:
		//line sql.y:456
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 71:
		//line sql.y:462
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 72:
		//line sql.y:466
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 73:
		//line sql.y:470
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 74:
		//line sql.y:474
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 75:
		//line sql.y:479
		{
			yyVAL.bytes = nil
		}
	case 76:
		//line sql.y:483
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 77:
		//line sql.y:487
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 78:
		//line sql.y:493
		{
			yyVAL.str = AST_JOIN
		}
	case 79:
		//line sql.y:497
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 80:
		//line sql.y:501
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 81:
		//line sql.y:505
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 82:
		//line sql.y:509
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 83:
		//line sql.y:513
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 84:
		//line sql.y:517
		{
			yyVAL.str = AST_JOIN
		}
	case 85:
		//line sql.y:521
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 86:
		//line sql.y:525
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 87:
		//line sql.y:531
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 88:
		//line sql.y:535
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 89:
		//line sql.y:539
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 90:
		//line sql.y:545
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 91:
		//line sql.y:549
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 92:
		//line sql.y:554
		{
			yyVAL.indexHints = nil
		}
	case 93:
		//line sql.y:558
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 94:
		//line sql.y:562
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 95:
		//line sql.y:566
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 96:
		//line sql.y:572
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 97:
		//line sql.y:576
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 98:
		//line sql.y:581
		{
			yyVAL.boolExpr = nil
		}
	case 99:
		//line sql.y:585
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 100:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 101:
		//line sql.y:592
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 102:
		//line sql.y:596
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 103:
		//line sql.y:600
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 104:
		//line sql.y:604
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 105:
		//line sql.y:610
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 106:
		//line sql.y:614
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].colTuple}
		}
	case 107:
		//line sql.y:618
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].colTuple}
		}
	case 108:
		//line sql.y:622
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 109:
		//line sql.y:626
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 110:
		//line sql.y:630
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 111:
		//line sql.y:634
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 112:
		//line sql.y:638
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 113:
		//line sql.y:642
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 114:
		//line sql.y:646
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 115:
		//line sql.y:652
		{
			yyVAL.str = AST_EQ
		}
	case 116:
		//line sql.y:656
		{
			yyVAL.str = AST_LT
		}
	case 117:
		//line sql.y:660
		{
			yyVAL.str = AST_GT
		}
	case 118:
		//line sql.y:664
		{
			yyVAL.str = AST_LE
		}
	case 119:
		//line sql.y:668
		{
			yyVAL.str = AST_GE
		}
	case 120:
		//line sql.y:672
		{
			yyVAL.str = AST_NE
		}
	case 121:
		//line sql.y:676
		{
			yyVAL.str = AST_NSE
		}
	case 122:
		//line sql.y:682
		{
			yyVAL.colTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 123:
		//line sql.y:686
		{
			yyVAL.colTuple = yyS[yypt-0].subquery
		}
	case 124:
		//line sql.y:690
		{
			yyVAL.colTuple = ListArg(yyS[yypt-0].bytes)
		}
	case 125:
		//line sql.y:696
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 126:
		//line sql.y:702
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 127:
		//line sql.y:706
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 128:
		//line sql.y:712
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 129:
		//line sql.y:716
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 130:
		//line sql.y:720
		{
			yyVAL.valExpr = yyS[yypt-0].rowTuple
		}
	case 131:
		//line sql.y:724
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:728
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 133:
		//line sql.y:732
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 134:
		//line sql.y:736
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 135:
		//line sql.y:740
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 136:
		//line sql.y:744
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 137:
		//line sql.y:748
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 138:
		//line sql.y:752
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 139:
		//line sql.y:756
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
	case 140:
		//line sql.y:771
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 141:
		//line sql.y:775
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 142:
		//line sql.y:779
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 143:
		//line sql.y:783
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 144:
		//line sql.y:787
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 145:
		//line sql.y:793
		{
			yyVAL.bytes = IF_BYTES
		}
	case 146:
		//line sql.y:797
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 147:
		//line sql.y:803
		{
			yyVAL.byt = AST_UPLUS
		}
	case 148:
		//line sql.y:807
		{
			yyVAL.byt = AST_UMINUS
		}
	case 149:
		//line sql.y:811
		{
			yyVAL.byt = AST_TILDA
		}
	case 150:
		//line sql.y:817
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 151:
		//line sql.y:822
		{
			yyVAL.valExpr = nil
		}
	case 152:
		//line sql.y:826
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 153:
		//line sql.y:832
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 154:
		//line sql.y:836
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 155:
		//line sql.y:842
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 156:
		//line sql.y:847
		{
			yyVAL.valExpr = nil
		}
	case 157:
		//line sql.y:851
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 158:
		//line sql.y:857
		{
			yyVAL.colName = &ColName{Name: bytes.ToLower(yyS[yypt-0].bytes)}
		}
	case 159:
		//line sql.y:861
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: bytes.ToLower(yyS[yypt-0].bytes)}
		}
	case 160:
		//line sql.y:867
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 161:
		//line sql.y:871
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 162:
		//line sql.y:875
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 163:
		//line sql.y:879
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 164:
		//line sql.y:884
		{
			yyVAL.valExprs = nil
		}
	case 165:
		//line sql.y:888
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 166:
		//line sql.y:893
		{
			yyVAL.boolExpr = nil
		}
	case 167:
		//line sql.y:897
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 168:
		//line sql.y:902
		{
			yyVAL.orderBy = nil
		}
	case 169:
		//line sql.y:906
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 170:
		//line sql.y:912
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 171:
		//line sql.y:916
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 172:
		//line sql.y:922
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 173:
		//line sql.y:927
		{
			yyVAL.str = AST_ASC
		}
	case 174:
		//line sql.y:931
		{
			yyVAL.str = AST_ASC
		}
	case 175:
		//line sql.y:935
		{
			yyVAL.str = AST_DESC
		}
	case 176:
		//line sql.y:940
		{
			yyVAL.limit = nil
		}
	case 177:
		//line sql.y:944
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 178:
		//line sql.y:948
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 179:
		//line sql.y:953
		{
			yyVAL.str = ""
		}
	case 180:
		//line sql.y:957
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 181:
		//line sql.y:961
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
	case 182:
		//line sql.y:974
		{
			yyVAL.columns = nil
		}
	case 183:
		//line sql.y:978
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 184:
		//line sql.y:984
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 185:
		//line sql.y:988
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 186:
		//line sql.y:993
		{
			yyVAL.updateExprs = nil
		}
	case 187:
		//line sql.y:997
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 188:
		//line sql.y:1003
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 189:
		//line sql.y:1007
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 190:
		//line sql.y:1013
		{
			yyVAL.values = Values{yyS[yypt-0].rowTuple}
		}
	case 191:
		//line sql.y:1017
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].rowTuple)
		}
	case 192:
		//line sql.y:1023
		{
			yyVAL.rowTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 193:
		//line sql.y:1027
		{
			yyVAL.rowTuple = yyS[yypt-0].subquery
		}
	case 194:
		//line sql.y:1033
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 195:
		//line sql.y:1037
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 196:
		//line sql.y:1043
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 197:
		//line sql.y:1048
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1050
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1053
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1055
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1058
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1060
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1064
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1066
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1068
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1070
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1072
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		//line sql.y:1075
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		//line sql.y:1077
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		//line sql.y:1080
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		//line sql.y:1082
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		//line sql.y:1085
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		//line sql.y:1087
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		//line sql.y:1091
		{
			//comment by liuqi    $$ = bytes.ToLower($1)
			yyVAL.bytes = yyS[yypt-0].bytes //add by liuqi
		}
	case 215:
		//line sql.y:1097
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
