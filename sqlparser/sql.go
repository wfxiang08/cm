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
const UNION = 57394
const MINUS = 57395
const EXCEPT = 57396
const INTERSECT = 57397
const JOIN = 57398
const STRAIGHT_JOIN = 57399
const LEFT = 57400
const RIGHT = 57401
const INNER = 57402
const OUTER = 57403
const CROSS = 57404
const NATURAL = 57405
const USE = 57406
const FORCE = 57407
const ON = 57408
const OR = 57409
const AND = 57410
const NOT = 57411
const UNARY = 57412
const CASE = 57413
const WHEN = 57414
const THEN = 57415
const ELSE = 57416
const END = 57417
const CREATE = 57418
const ALTER = 57419
const DROP = 57420
const RENAME = 57421
const ANALYZE = 57422
const TABLE = 57423
const INDEX = 57424
const VIEW = 57425
const TO = 57426
const IGNORE = 57427
const IF = 57428
const UNIQUE = 57429
const USING = 57430
const SHOW = 57431
const DESCRIBE = 57432
const EXPLAIN = 57433

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

const yyNprod = 214
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 591

var yyAct = []int{

	107, 311, 173, 379, 74, 104, 347, 264, 176, 105,
	213, 253, 224, 301, 193, 255, 93, 175, 3, 186,
	388, 76, 388, 201, 388, 94, 61, 110, 89, 150,
	149, 81, 114, 144, 271, 120, 98, 138, 37, 38,
	39, 40, 97, 111, 112, 113, 358, 245, 78, 307,
	144, 83, 102, 77, 64, 86, 118, 115, 144, 90,
	245, 62, 63, 243, 328, 330, 390, 357, 389, 99,
	387, 275, 276, 277, 278, 279, 101, 280, 281, 337,
	116, 117, 95, 134, 356, 338, 85, 121, 244, 54,
	82, 55, 142, 334, 329, 306, 296, 146, 57, 58,
	59, 48, 119, 50, 294, 177, 246, 51, 60, 178,
	135, 56, 254, 137, 299, 254, 332, 162, 163, 164,
	150, 149, 286, 148, 185, 131, 78, 126, 133, 78,
	191, 77, 197, 196, 77, 340, 150, 149, 172, 174,
	149, 128, 353, 302, 189, 84, 233, 99, 219, 197,
	195, 267, 302, 141, 223, 355, 354, 231, 232, 198,
	235, 236, 237, 238, 239, 240, 241, 242, 182, 211,
	218, 366, 367, 71, 326, 325, 207, 160, 161, 162,
	163, 164, 247, 99, 99, 220, 221, 222, 78, 78,
	234, 78, 258, 77, 260, 262, 77, 249, 251, 324,
	205, 217, 268, 208, 257, 261, 263, 257, 322, 320,
	226, 128, 194, 323, 321, 157, 158, 159, 160, 161,
	162, 163, 164, 88, 194, 245, 285, 272, 247, 269,
	143, 266, 289, 290, 287, 364, 18, 19, 21, 22,
	20, 288, 342, 123, 124, 130, 79, 127, 293, 25,
	27, 28, 26, 99, 204, 206, 203, 37, 38, 39,
	40, 75, 273, 227, 298, 190, 304, 295, 216, 225,
	23, 310, 305, 217, 128, 308, 309, 188, 215, 374,
	187, 144, 18, 91, 318, 319, 226, 373, 372, 179,
	183, 300, 188, 181, 336, 180, 122, 275, 276, 277,
	278, 279, 339, 280, 281, 284, 84, 385, 78, 125,
	79, 147, 344, 343, 333, 345, 348, 331, 216, 315,
	314, 283, 24, 29, 31, 30, 32, 84, 215, 363,
	386, 217, 217, 210, 209, 33, 34, 35, 359, 192,
	139, 349, 136, 360, 157, 158, 159, 160, 161, 162,
	163, 164, 362, 132, 129, 247, 72, 92, 368, 87,
	361, 370, 341, 18, 70, 69, 376, 348, 292, 392,
	378, 377, 199, 380, 380, 380, 78, 381, 382, 140,
	383, 77, 65, 18, 250, 67, 110, 52, 369, 393,
	371, 114, 256, 394, 120, 395, 228, 312, 229, 230,
	352, 97, 111, 112, 113, 313, 265, 351, 317, 114,
	194, 102, 120, 73, 391, 118, 18, 375, 18, 79,
	111, 112, 113, 42, 200, 49, 270, 41, 202, 179,
	53, 80, 259, 118, 384, 101, 365, 110, 346, 116,
	117, 95, 114, 350, 316, 120, 121, 43, 44, 45,
	46, 47, 79, 111, 112, 113, 297, 116, 117, 184,
	110, 119, 102, 252, 121, 114, 118, 248, 120, 109,
	106, 108, 303, 103, 151, 79, 111, 112, 113, 119,
	100, 327, 214, 274, 212, 102, 101, 96, 282, 118,
	116, 117, 145, 66, 36, 114, 68, 121, 120, 17,
	16, 12, 11, 10, 15, 79, 111, 112, 113, 101,
	14, 13, 119, 116, 117, 179, 9, 8, 5, 118,
	121, 7, 6, 4, 2, 1, 152, 156, 154, 155,
	0, 0, 0, 0, 0, 119, 0, 0, 0, 0,
	0, 0, 0, 116, 117, 0, 168, 169, 170, 171,
	121, 165, 166, 167, 335, 0, 157, 158, 159, 160,
	161, 162, 163, 164, 291, 119, 157, 158, 159, 160,
	161, 162, 163, 164, 153, 157, 158, 159, 160, 161,
	162, 163, 164, 157, 158, 159, 160, 161, 162, 163,
	164,
}
var yyPact = []int{

	231, -1000, -1000, 200, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 5, -1000, 365, -1000, -1000, -9,
	15, 2, 12, -1000, -1000, -1000, 413, 359, -1000, -1000,
	-1000, 361, -1000, 330, 329, 315, 403, 205, -70, -7,
	265, -1000, -1000, -10, 265, -1000, 318, -73, 265, -73,
	316, -1000, -1000, -1000, -1000, -1000, 1, -1000, 250, 315,
	315, 270, 43, 315, 150, 313, -1000, 193, -1000, 41,
	312, 53, 265, -1000, -1000, 301, -1000, -62, 299, 353,
	81, 265, -1000, 220, -1000, -1000, 286, 39, 63, 499,
	-1000, 434, 411, -1000, -1000, -1000, 464, 244, 242, -1000,
	239, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 464, -1000, 241, 226, 269, 298, 399, 269, -1000,
	464, 265, -1000, 346, -80, -1000, 162, -1000, 293, -1000,
	-1000, 292, -1000, 227, 1, -1000, -1000, 265, 104, 434,
	434, 464, 218, 369, 464, 464, 115, 464, 464, 464,
	464, 464, 464, 464, 464, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 499, -44, -19, -1, 499, -1000, 378,
	360, 1, -1000, 413, 28, 507, 358, 269, 269, 358,
	269, 213, -1000, 392, 434, -1000, 507, -1000, -1000, -1000,
	79, 265, -1000, -65, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 201, 235, 280, 277, 38, -1000, -1000, -1000,
	-1000, -1000, 66, 507, -1000, 378, -1000, -1000, 218, 464,
	464, 507, 490, -1000, 337, 98, 98, 98, 36, 36,
	-1000, -1000, -1000, -1000, -1000, 464, -1000, 507, -1000, -3,
	1, -11, 25, -1000, 434, 71, 238, 200, 80, -12,
	-1000, 71, 80, 392, 381, 390, 63, 279, -1000, -1000,
	278, -1000, 396, 227, 227, -1000, -1000, 147, 146, 137,
	113, 112, -6, -1000, 276, 9, 273, -14, -1000, 507,
	480, 464, -1000, 507, -1000, -28, -1000, -5, -1000, 464,
	47, -1000, 326, 181, -1000, -1000, -1000, 269, -1000, -1000,
	381, -1000, 464, 464, -1000, -1000, 394, 385, 235, 70,
	-1000, 94, -1000, 93, -1000, -1000, -1000, -1000, -13, -30,
	-51, -1000, -1000, -1000, -1000, 464, 507, -1000, -1000, 507,
	464, 323, 238, -1000, -1000, 268, 174, -1000, 139, -1000,
	392, 434, 464, 434, -1000, -1000, 237, 236, 228, 507,
	507, 410, -1000, 464, 464, -1000, -1000, -1000, 381, 63,
	164, 63, 265, 265, 265, 269, 507, -1000, 290, -37,
	-1000, -39, -41, 150, -1000, 407, 342, -1000, 265, -1000,
	-1000, -1000, 265, -1000, 265, -1000,
}
var yyPgo = []int{

	0, 525, 524, 17, 523, 522, 521, 518, 517, 516,
	511, 510, 504, 503, 502, 501, 500, 499, 427, 496,
	494, 493, 16, 25, 492, 488, 487, 484, 10, 483,
	482, 173, 481, 3, 14, 36, 480, 474, 15, 473,
	2, 12, 8, 472, 9, 471, 57, 470, 5, 469,
	463, 11, 459, 456, 444, 443, 7, 438, 6, 436,
	1, 434, 19, 432, 13, 4, 21, 223, 431, 430,
	428, 426, 425, 424, 0, 26, 423,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 7, 7, 5, 6, 8, 8, 9, 9,
	9, 13, 13, 14, 15, 10, 10, 10, 11, 12,
	12, 12, 16, 17, 17, 17, 76, 18, 19, 19,
	20, 20, 20, 20, 20, 21, 21, 22, 22, 23,
	23, 23, 26, 26, 24, 24, 24, 27, 27, 28,
	28, 28, 28, 25, 25, 25, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 30, 30, 30, 31, 31,
	32, 32, 32, 32, 33, 33, 34, 34, 35, 35,
	35, 35, 35, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 37, 37, 37, 37, 37, 37, 37,
	41, 41, 41, 46, 42, 42, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 45, 45, 47, 47, 47, 49, 52,
	52, 50, 50, 51, 53, 53, 48, 48, 39, 39,
	39, 39, 54, 54, 55, 55, 56, 56, 57, 57,
	58, 59, 59, 59, 60, 60, 60, 61, 61, 61,
	62, 62, 63, 63, 64, 64, 38, 38, 43, 43,
	44, 44, 65, 65, 66, 67, 67, 68, 68, 69,
	69, 70, 70, 70, 70, 70, 71, 71, 72, 72,
	73, 73, 74, 75,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 7, 7, 8, 7, 3, 4, 5, 8,
	4, 1, 2, 1, 1, 6, 7, 4, 5, 4,
	5, 5, 3, 2, 2, 2, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 1, 3,
	3, 2, 3, 3, 3, 4, 3, 4, 5, 6,
	3, 4, 2, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 1, 3, 1, 3, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 3, 4,
	5, 4, 1, 1, 1, 1, 1, 1, 5, 0,
	1, 1, 2, 4, 0, 2, 1, 3, 1, 1,
	1, 1, 0, 3, 0, 2, 0, 3, 1, 3,
	2, 0, 1, 1, 0, 2, 4, 0, 2, 4,
	0, 3, 1, 3, 0, 5, 2, 1, 1, 3,
	3, 1, 1, 3, 3, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -7, -5, -6, -8, -9,
	-13, -14, -15, -10, -11, -12, -16, -17, 5, 6,
	9, 7, 8, 39, 91, 18, 21, 19, 20, 92,
	94, 93, 95, 104, 105, 106, -20, 57, 58, 59,
	60, -18, -76, -18, -18, -18, -18, -18, 96, -72,
	98, 102, 22, -69, 98, 100, 96, 96, 97, 98,
	96, -75, -75, -75, -3, 23, -21, 24, -19, 35,
	35, -31, 41, 10, -65, 56, -66, -48, -74, 41,
	-68, 101, 97, -74, 41, 96, -74, 41, -67, 101,
	-74, -67, 41, -22, -23, 81, -26, 41, -35, -40,
	-36, 75, 51, -39, -48, -44, -47, -74, -45, -49,
	26, 42, 43, 44, 31, -46, 79, 80, 55, 101,
	34, 86, 46, -31, -31, 39, 84, -31, 61, 41,
	52, 84, 41, 75, -74, -75, 41, -75, 99, 41,
	26, 72, -74, 10, 61, -24, -74, 25, 84, 74,
	73, -37, 27, 75, 29, 30, 28, 76, 77, 78,
	79, 80, 81, 82, 83, 52, 53, 54, 47, 48,
	49, 50, -35, -40, -35, -3, -42, -40, -40, 51,
	51, 51, -46, 51, -52, -40, -62, 39, 51, -62,
	39, -65, 41, -34, 11, -66, -40, -74, -75, 26,
	-73, 103, -70, 94, 92, 38, 93, 14, 41, 41,
	41, -75, -27, -28, -30, 51, 41, -46, -23, -74,
	81, -35, -35, -40, -41, 51, -46, 45, 27, 29,
	30, -40, -40, 31, 75, -40, -40, -40, -40, -40,
	-40, -40, -40, 107, 107, 61, 107, -40, 107, -22,
	24, -22, -50, -51, 87, -38, 34, -3, -65, -63,
	-48, -38, -65, -34, -56, 14, -35, 72, -74, -75,
	-71, 99, -34, 61, -29, 62, 63, 64, 65, 66,
	68, 69, -25, 41, 25, -28, 84, -42, -41, -40,
	-40, 74, 31, -40, 107, -22, 107, -53, -51, 89,
	-35, -64, 72, -43, -44, -64, 107, 61, -64, -64,
	-56, -60, 16, 15, 41, 41, -54, 12, -28, -28,
	62, 67, 62, 67, 62, 62, 62, -32, 70, 100,
	71, 41, 107, 41, 107, 74, -40, 107, 90, -40,
	88, 36, 61, -48, -60, -40, -57, -58, -40, -75,
	-55, 13, 15, 72, 62, 62, 97, 97, 97, -40,
	-40, 37, -44, 61, 61, -59, 32, 33, -56, -35,
	-42, -35, 51, 51, 51, 7, -40, -58, -60, -33,
	-74, -33, -33, -65, -61, 17, 40, 107, 61, 107,
	107, 7, 27, -74, -74, -74,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 46, 46,
	46, 46, 46, 46, 208, 31, 0, 33, 34, 199,
	0, 0, 0, 213, 213, 213, 0, 50, 52, 53,
	54, 55, 48, 0, 0, 0, 0, 0, 197, 0,
	0, 209, 32, 0, 0, 200, 0, 195, 0, 195,
	0, 43, 44, 45, 19, 51, 0, 56, 47, 0,
	0, 0, 88, 0, 26, 0, 192, 0, 156, 212,
	0, 0, 0, 213, 212, 0, 213, 0, 0, 0,
	0, 0, 42, 17, 57, 59, 64, 212, 62, 63,
	98, 0, 0, 126, 127, 128, 0, 156, 0, 142,
	0, 158, 159, 160, 161, 191, 145, 146, 147, 143,
	144, 149, 49, 180, 180, 0, 0, 96, 0, 27,
	0, 0, 213, 0, 210, 30, 0, 37, 0, 39,
	196, 0, 213, 0, 0, 60, 65, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 113, 114, 115, 116, 117,
	118, 119, 101, 0, 0, 0, 0, 124, 137, 0,
	0, 0, 112, 0, 0, 150, 0, 0, 0, 0,
	0, 96, 89, 166, 0, 193, 194, 157, 28, 198,
	0, 0, 213, 206, 201, 202, 203, 204, 205, 38,
	40, 41, 96, 67, 73, 0, 85, 87, 58, 66,
	61, 99, 100, 103, 104, 0, 121, 122, 0, 0,
	0, 106, 0, 110, 0, 129, 130, 131, 132, 133,
	134, 135, 136, 102, 123, 0, 190, 124, 138, 0,
	0, 0, 154, 151, 0, 184, 0, 187, 184, 0,
	182, 184, 184, 166, 174, 0, 97, 0, 211, 35,
	0, 207, 162, 0, 0, 76, 77, 0, 0, 0,
	0, 0, 90, 74, 0, 0, 0, 0, 105, 107,
	0, 0, 111, 125, 139, 0, 141, 0, 152, 0,
	0, 20, 0, 186, 188, 21, 181, 0, 22, 23,
	174, 25, 0, 0, 213, 36, 164, 0, 68, 71,
	78, 0, 80, 0, 82, 83, 84, 69, 0, 0,
	0, 75, 70, 86, 120, 0, 108, 140, 148, 155,
	0, 0, 0, 183, 24, 175, 167, 168, 171, 29,
	166, 0, 0, 0, 79, 81, 0, 0, 0, 109,
	153, 0, 189, 0, 0, 170, 172, 173, 174, 165,
	163, 72, 0, 0, 0, 0, 176, 169, 177, 0,
	94, 0, 0, 185, 18, 0, 0, 91, 0, 92,
	93, 178, 0, 95, 0, 179,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 83, 76, 3,
	51, 107, 81, 79, 61, 80, 84, 82, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	53, 52, 54, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 78, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 77, 3, 55,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 56,
	57, 58, 59, 60, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 73, 74, 75, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
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
		//line sql.y:155
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:161
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
		//line sql.y:181
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 18:
		//line sql.y:185
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 19:
		//line sql.y:189
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 20:
		//line sql.y:195
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:199
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
		//line sql.y:212
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 23:
		//line sql.y:216
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
		//line sql.y:229
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 25:
		//line sql.y:235
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 26:
		//line sql.y:241
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 27:
		//line sql.y:245
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 28:
		//line sql.y:251
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 29:
		//line sql.y:255
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 30:
		//line sql.y:260
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 31:
		//line sql.y:266
		{
			yyVAL.statement = &Begin{}
		}
	case 32:
		//line sql.y:270
		{
			yyVAL.statement = &Begin{}
		}
	case 33:
		//line sql.y:276
		{
			yyVAL.statement = &Commit{}
		}
	case 34:
		//line sql.y:282
		{
			yyVAL.statement = &Rollback{}
		}
	case 35:
		//line sql.y:289
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 36:
		//line sql.y:293
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 37:
		//line sql.y:298
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 38:
		//line sql.y:304
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 39:
		//line sql.y:310
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 40:
		//line sql.y:314
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 41:
		//line sql.y:319
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 42:
		//line sql.y:325
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 43:
		//line sql.y:331
		{
			yyVAL.statement = &Other{}
		}
	case 44:
		//line sql.y:335
		{
			yyVAL.statement = &Other{}
		}
	case 45:
		//line sql.y:339
		{
			yyVAL.statement = &Other{}
		}
	case 46:
		//line sql.y:344
		{
			SetAllowComments(yylex, true)
		}
	case 47:
		//line sql.y:348
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 48:
		//line sql.y:354
		{
			yyVAL.bytes2 = nil
		}
	case 49:
		//line sql.y:358
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 50:
		//line sql.y:364
		{
			yyVAL.str = AST_UNION
		}
	case 51:
		//line sql.y:368
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 52:
		//line sql.y:372
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 53:
		//line sql.y:376
		{
			yyVAL.str = AST_EXCEPT
		}
	case 54:
		//line sql.y:380
		{
			yyVAL.str = AST_INTERSECT
		}
	case 55:
		//line sql.y:385
		{
			yyVAL.str = ""
		}
	case 56:
		//line sql.y:389
		{
			yyVAL.str = AST_DISTINCT
		}
	case 57:
		//line sql.y:395
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 58:
		//line sql.y:399
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 59:
		//line sql.y:405
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 60:
		//line sql.y:409
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 61:
		//line sql.y:413
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 62:
		//line sql.y:419
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 63:
		//line sql.y:423
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 64:
		//line sql.y:428
		{
			yyVAL.bytes = nil
		}
	case 65:
		//line sql.y:432
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 66:
		//line sql.y:436
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 67:
		//line sql.y:442
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 68:
		//line sql.y:446
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 69:
		//line sql.y:452
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 70:
		//line sql.y:456
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 71:
		//line sql.y:460
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 72:
		//line sql.y:464
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 73:
		//line sql.y:469
		{
			yyVAL.bytes = nil
		}
	case 74:
		//line sql.y:473
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 75:
		//line sql.y:477
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 76:
		//line sql.y:483
		{
			yyVAL.str = AST_JOIN
		}
	case 77:
		//line sql.y:487
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 78:
		//line sql.y:491
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 79:
		//line sql.y:495
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 80:
		//line sql.y:499
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 81:
		//line sql.y:503
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 82:
		//line sql.y:507
		{
			yyVAL.str = AST_JOIN
		}
	case 83:
		//line sql.y:511
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 84:
		//line sql.y:515
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 85:
		//line sql.y:521
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 86:
		//line sql.y:525
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 87:
		//line sql.y:529
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 88:
		//line sql.y:535
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 89:
		//line sql.y:539
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 90:
		//line sql.y:544
		{
			yyVAL.indexHints = nil
		}
	case 91:
		//line sql.y:548
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 92:
		//line sql.y:552
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 93:
		//line sql.y:556
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 94:
		//line sql.y:562
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 95:
		//line sql.y:566
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 96:
		//line sql.y:571
		{
			yyVAL.boolExpr = nil
		}
	case 97:
		//line sql.y:575
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 98:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 99:
		//line sql.y:582
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 100:
		//line sql.y:586
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 101:
		//line sql.y:590
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 102:
		//line sql.y:594
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 103:
		//line sql.y:600
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 104:
		//line sql.y:604
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].colTuple}
		}
	case 105:
		//line sql.y:608
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].colTuple}
		}
	case 106:
		//line sql.y:612
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 107:
		//line sql.y:616
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 108:
		//line sql.y:620
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 109:
		//line sql.y:624
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 110:
		//line sql.y:628
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 111:
		//line sql.y:632
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 112:
		//line sql.y:636
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 113:
		//line sql.y:642
		{
			yyVAL.str = AST_EQ
		}
	case 114:
		//line sql.y:646
		{
			yyVAL.str = AST_LT
		}
	case 115:
		//line sql.y:650
		{
			yyVAL.str = AST_GT
		}
	case 116:
		//line sql.y:654
		{
			yyVAL.str = AST_LE
		}
	case 117:
		//line sql.y:658
		{
			yyVAL.str = AST_GE
		}
	case 118:
		//line sql.y:662
		{
			yyVAL.str = AST_NE
		}
	case 119:
		//line sql.y:666
		{
			yyVAL.str = AST_NSE
		}
	case 120:
		//line sql.y:672
		{
			yyVAL.colTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 121:
		//line sql.y:676
		{
			yyVAL.colTuple = yyS[yypt-0].subquery
		}
	case 122:
		//line sql.y:680
		{
			yyVAL.colTuple = ListArg(yyS[yypt-0].bytes)
		}
	case 123:
		//line sql.y:686
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 124:
		//line sql.y:692
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 125:
		//line sql.y:696
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 126:
		//line sql.y:702
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 127:
		//line sql.y:706
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 128:
		//line sql.y:710
		{
			yyVAL.valExpr = yyS[yypt-0].rowTuple
		}
	case 129:
		//line sql.y:714
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:718
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 131:
		//line sql.y:722
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:726
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 133:
		//line sql.y:730
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 134:
		//line sql.y:734
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 135:
		//line sql.y:738
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 136:
		//line sql.y:742
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 137:
		//line sql.y:746
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
	case 138:
		//line sql.y:761
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 139:
		//line sql.y:765
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 140:
		//line sql.y:769
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 141:
		//line sql.y:773
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 142:
		//line sql.y:777
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 143:
		//line sql.y:783
		{
			yyVAL.bytes = IF_BYTES
		}
	case 144:
		//line sql.y:787
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 145:
		//line sql.y:793
		{
			yyVAL.byt = AST_UPLUS
		}
	case 146:
		//line sql.y:797
		{
			yyVAL.byt = AST_UMINUS
		}
	case 147:
		//line sql.y:801
		{
			yyVAL.byt = AST_TILDA
		}
	case 148:
		//line sql.y:807
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 149:
		//line sql.y:812
		{
			yyVAL.valExpr = nil
		}
	case 150:
		//line sql.y:816
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 151:
		//line sql.y:822
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 152:
		//line sql.y:826
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 153:
		//line sql.y:832
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 154:
		//line sql.y:837
		{
			yyVAL.valExpr = nil
		}
	case 155:
		//line sql.y:841
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 156:
		//line sql.y:847
		{
			yyVAL.colName = &ColName{Name: bytes.ToLower(yyS[yypt-0].bytes)}
		}
	case 157:
		//line sql.y:851
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: bytes.ToLower(yyS[yypt-0].bytes)}
		}
	case 158:
		//line sql.y:857
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 159:
		//line sql.y:861
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 160:
		//line sql.y:865
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 161:
		//line sql.y:869
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 162:
		//line sql.y:874
		{
			yyVAL.valExprs = nil
		}
	case 163:
		//line sql.y:878
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 164:
		//line sql.y:883
		{
			yyVAL.boolExpr = nil
		}
	case 165:
		//line sql.y:887
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 166:
		//line sql.y:892
		{
			yyVAL.orderBy = nil
		}
	case 167:
		//line sql.y:896
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 168:
		//line sql.y:902
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 169:
		//line sql.y:906
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 170:
		//line sql.y:912
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 171:
		//line sql.y:917
		{
			yyVAL.str = AST_ASC
		}
	case 172:
		//line sql.y:921
		{
			yyVAL.str = AST_ASC
		}
	case 173:
		//line sql.y:925
		{
			yyVAL.str = AST_DESC
		}
	case 174:
		//line sql.y:930
		{
			yyVAL.limit = nil
		}
	case 175:
		//line sql.y:934
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 176:
		//line sql.y:938
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 177:
		//line sql.y:943
		{
			yyVAL.str = ""
		}
	case 178:
		//line sql.y:947
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 179:
		//line sql.y:951
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
	case 180:
		//line sql.y:964
		{
			yyVAL.columns = nil
		}
	case 181:
		//line sql.y:968
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 182:
		//line sql.y:974
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 183:
		//line sql.y:978
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 184:
		//line sql.y:983
		{
			yyVAL.updateExprs = nil
		}
	case 185:
		//line sql.y:987
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 186:
		//line sql.y:993
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 187:
		//line sql.y:997
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 188:
		//line sql.y:1003
		{
			yyVAL.values = Values{yyS[yypt-0].rowTuple}
		}
	case 189:
		//line sql.y:1007
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].rowTuple)
		}
	case 190:
		//line sql.y:1013
		{
			yyVAL.rowTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 191:
		//line sql.y:1017
		{
			yyVAL.rowTuple = yyS[yypt-0].subquery
		}
	case 192:
		//line sql.y:1023
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 193:
		//line sql.y:1027
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 194:
		//line sql.y:1033
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 195:
		//line sql.y:1038
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1040
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1043
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1045
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1048
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1050
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1054
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1056
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1058
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1060
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1062
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1065
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1067
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		//line sql.y:1070
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		//line sql.y:1072
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		//line sql.y:1075
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		//line sql.y:1077
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		//line sql.y:1081
		{
			//comment by liuqi    $$ = bytes.ToLower($1)
			yyVAL.bytes = yyS[yypt-0].bytes //add by liuqi
		}
	case 213:
		//line sql.y:1087
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
