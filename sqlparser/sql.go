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
const START = 57362
const TRANSACTION = 57363
const ALL = 57364
const DISTINCT = 57365
const AS = 57366
const EXISTS = 57367
const IN = 57368
const IS = 57369
const LIKE = 57370
const BETWEEN = 57371
const NULL = 57372
const ASC = 57373
const DESC = 57374
const VALUES = 57375
const INTO = 57376
const DUPLICATE = 57377
const KEY = 57378
const DEFAULT = 57379
const SET = 57380
const LOCK = 57381
const ID = 57382
const STRING = 57383
const NUMBER = 57384
const VALUE_ARG = 57385
const LIST_ARG = 57386
const COMMENT = 57387
const LE = 57388
const GE = 57389
const NE = 57390
const NULL_SAFE_EQUAL = 57391
const UNION = 57392
const MINUS = 57393
const EXCEPT = 57394
const INTERSECT = 57395
const JOIN = 57396
const STRAIGHT_JOIN = 57397
const LEFT = 57398
const RIGHT = 57399
const INNER = 57400
const OUTER = 57401
const CROSS = 57402
const NATURAL = 57403
const USE = 57404
const FORCE = 57405
const ON = 57406
const OR = 57407
const AND = 57408
const NOT = 57409
const UNARY = 57410
const CASE = 57411
const WHEN = 57412
const THEN = 57413
const ELSE = 57414
const END = 57415
const CREATE = 57416
const ALTER = 57417
const DROP = 57418
const RENAME = 57419
const ANALYZE = 57420
const TABLE = 57421
const INDEX = 57422
const VIEW = 57423
const TO = 57424
const IGNORE = 57425
const IF = 57426
const UNIQUE = 57427
const USING = 57428
const SHOW = 57429
const DESCRIBE = 57430
const EXPLAIN = 57431

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

const yyNprod = 210
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 590

var yyAct = []int{

	102, 298, 166, 366, 99, 93, 334, 253, 70, 169,
	100, 290, 244, 88, 204, 215, 168, 3, 375, 375,
	71, 184, 375, 143, 142, 89, 58, 264, 265, 266,
	267, 268, 192, 269, 270, 35, 36, 37, 38, 84,
	137, 76, 236, 296, 137, 73, 315, 317, 78, 72,
	260, 61, 81, 137, 131, 110, 85, 234, 80, 59,
	60, 236, 325, 345, 377, 376, 94, 45, 374, 47,
	344, 343, 319, 48, 245, 51, 316, 52, 127, 275,
	57, 17, 18, 19, 20, 235, 324, 135, 321, 295,
	285, 77, 139, 23, 25, 26, 24, 53, 141, 283,
	170, 124, 165, 167, 171, 128, 120, 237, 130, 54,
	55, 56, 143, 142, 21, 350, 245, 126, 288, 178,
	73, 224, 142, 73, 72, 188, 187, 72, 182, 340,
	150, 151, 152, 153, 154, 155, 156, 157, 291, 256,
	94, 210, 188, 186, 155, 156, 157, 214, 212, 213,
	222, 223, 189, 226, 227, 228, 229, 230, 231, 232,
	233, 175, 202, 209, 225, 22, 27, 29, 28, 30,
	198, 134, 122, 67, 342, 238, 94, 94, 31, 32,
	33, 73, 73, 143, 142, 72, 251, 240, 242, 249,
	309, 255, 208, 257, 196, 310, 248, 199, 327, 79,
	341, 217, 313, 322, 252, 150, 151, 152, 153, 154,
	155, 156, 157, 153, 154, 155, 156, 157, 312, 238,
	258, 274, 311, 278, 279, 261, 276, 150, 151, 152,
	153, 154, 155, 156, 157, 277, 185, 307, 211, 282,
	118, 122, 308, 121, 94, 185, 136, 195, 197, 194,
	236, 289, 291, 351, 329, 284, 287, 83, 293, 180,
	297, 294, 208, 207, 105, 35, 36, 37, 38, 109,
	123, 181, 115, 206, 361, 217, 117, 305, 306, 74,
	106, 107, 108, 323, 218, 262, 17, 79, 360, 97,
	216, 326, 359, 113, 122, 172, 137, 73, 176, 331,
	174, 330, 332, 335, 264, 265, 266, 267, 268, 173,
	269, 270, 96, 273, 86, 74, 111, 112, 208, 208,
	140, 207, 320, 116, 372, 346, 318, 302, 336, 272,
	347, 206, 301, 201, 200, 183, 79, 132, 114, 129,
	349, 125, 238, 68, 356, 355, 358, 373, 348, 357,
	87, 82, 119, 363, 335, 328, 17, 365, 364, 66,
	367, 367, 367, 73, 368, 369, 281, 72, 241, 379,
	105, 370, 190, 133, 62, 109, 380, 64, 115, 49,
	381, 299, 382, 339, 247, 92, 106, 107, 108, 39,
	300, 254, 338, 109, 304, 97, 115, 185, 219, 113,
	220, 221, 69, 74, 106, 107, 108, 378, 41, 42,
	43, 44, 362, 172, 17, 40, 191, 113, 96, 46,
	105, 259, 111, 112, 90, 109, 193, 50, 115, 116,
	75, 250, 179, 371, 352, 92, 106, 107, 108, 333,
	111, 112, 337, 303, 114, 97, 286, 116, 177, 113,
	239, 243, 104, 101, 103, 17, 292, 98, 246, 144,
	17, 95, 114, 314, 205, 263, 203, 91, 96, 271,
	138, 63, 111, 112, 90, 105, 34, 65, 16, 116,
	109, 15, 11, 115, 10, 109, 9, 14, 115, 13,
	74, 106, 107, 108, 114, 74, 106, 107, 108, 12,
	97, 8, 7, 6, 113, 172, 5, 4, 2, 113,
	1, 0, 0, 0, 0, 145, 149, 147, 148, 0,
	0, 0, 0, 96, 0, 0, 0, 111, 112, 353,
	354, 0, 111, 112, 116, 161, 162, 163, 164, 116,
	158, 159, 160, 0, 0, 0, 0, 0, 0, 114,
	0, 0, 0, 0, 114, 0, 0, 0, 0, 0,
	0, 0, 146, 150, 151, 152, 153, 154, 155, 156,
	157, 0, 150, 151, 152, 153, 154, 155, 156, 157,
	280, 0, 150, 151, 152, 153, 154, 155, 156, 157,
}
var yyPact = []int{

	76, -1000, -1000, 210, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -27, -1000, 358, -1000, -1000, -21, 3, 15,
	-14, -1000, -1000, -1000, 409, 352, -1000, -1000, -1000, 354,
	-1000, 325, 303, 393, 275, -58, -4, 247, -1000, -1000,
	-36, 247, -1000, 311, -60, 247, -60, 310, -1000, -1000,
	-1000, -1000, -1000, 395, -1000, 231, 303, 314, 24, 303,
	113, -1000, 219, -1000, 19, 301, 44, 247, -1000, -1000,
	299, -1000, -43, 297, 348, 101, 247, -1000, 237, -1000,
	-1000, 296, 16, 41, 489, -1000, 239, 450, -1000, -1000,
	-1000, 363, 259, 250, -1000, 248, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 363, -1000, 221, 275,
	295, 387, 275, 363, 247, -1000, 347, -69, -1000, 157,
	-1000, 294, -1000, -1000, 293, -1000, 223, 395, -1000, -1000,
	247, 159, 239, 239, 363, 240, 372, 363, 363, 91,
	363, 363, 363, 363, 363, 363, 363, 363, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 489, -48, -20, 2,
	489, -1000, 455, 345, 395, -1000, 409, -11, 153, 351,
	275, 275, 235, -1000, 378, 239, -1000, 153, -1000, -1000,
	-1000, 69, 247, -1000, -47, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 226, 244, 289, 281, -3, -1000, -1000,
	-1000, -1000, -1000, 50, 153, -1000, 455, -1000, -1000, 240,
	363, 363, 153, 508, -1000, 336, 136, 136, 136, 65,
	65, -1000, -1000, -1000, -1000, -1000, 363, -1000, 153, -1000,
	-6, 395, -15, 31, -1000, 239, 68, 245, 210, 182,
	-16, -1000, 378, 366, 376, 41, 292, -1000, -1000, 287,
	-1000, 383, 223, 223, -1000, -1000, 177, 130, 162, 158,
	142, -22, -1000, 286, -33, 282, -17, -1000, 153, 131,
	363, -1000, 153, -1000, -19, -1000, -26, -1000, 363, 112,
	-1000, 320, 195, -1000, -1000, -1000, 275, 366, -1000, 363,
	363, -1000, -1000, 380, 369, 244, 59, -1000, 140, -1000,
	114, -1000, -1000, -1000, -1000, -24, -25, -32, -1000, -1000,
	-1000, -1000, 363, 153, -1000, -1000, 153, 363, 312, 245,
	-1000, -1000, 56, 194, -1000, 498, -1000, 378, 239, 363,
	239, -1000, -1000, 242, 238, 224, 153, 153, 405, -1000,
	363, 363, -1000, -1000, -1000, 366, 41, 191, 41, 247,
	247, 247, 275, 153, -1000, 308, -37, -1000, -40, -41,
	113, -1000, 400, 343, -1000, 247, -1000, -1000, -1000, 247,
	-1000, 247, -1000,
}
var yyPgo = []int{

	0, 510, 508, 16, 507, 506, 503, 502, 501, 499,
	489, 487, 486, 484, 482, 481, 478, 389, 477, 476,
	471, 13, 25, 470, 469, 467, 466, 14, 465, 464,
	173, 463, 3, 21, 5, 461, 459, 458, 457, 2,
	15, 9, 456, 10, 454, 55, 453, 4, 452, 451,
	12, 448, 446, 443, 442, 7, 439, 6, 434, 1,
	433, 432, 431, 11, 8, 20, 257, 430, 427, 426,
	421, 419, 416, 0, 26, 415,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 3, 3, 3, 4,
	4, 5, 6, 7, 8, 8, 8, 12, 12, 13,
	14, 9, 9, 9, 10, 11, 11, 11, 15, 16,
	16, 16, 75, 17, 18, 18, 19, 19, 19, 19,
	19, 20, 20, 21, 21, 22, 22, 22, 25, 25,
	23, 23, 23, 26, 26, 27, 27, 27, 27, 24,
	24, 24, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 29, 29, 29, 30, 30, 31, 31, 31, 31,
	32, 32, 33, 33, 34, 34, 34, 34, 34, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 36,
	36, 36, 36, 36, 36, 36, 40, 40, 40, 45,
	41, 41, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 44,
	44, 46, 46, 46, 48, 51, 51, 49, 49, 50,
	52, 52, 47, 47, 38, 38, 38, 38, 53, 53,
	54, 54, 55, 55, 56, 56, 57, 58, 58, 58,
	59, 59, 59, 60, 60, 60, 61, 61, 62, 62,
	63, 63, 37, 37, 42, 42, 43, 43, 64, 64,
	65, 66, 66, 67, 67, 68, 68, 69, 69, 69,
	69, 69, 70, 70, 71, 71, 72, 72, 73, 74,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 12, 3, 7,
	7, 8, 7, 3, 5, 8, 4, 1, 2, 1,
	1, 6, 7, 4, 5, 4, 5, 5, 3, 2,
	2, 2, 0, 2, 0, 2, 1, 2, 1, 1,
	1, 0, 1, 1, 3, 1, 2, 3, 1, 1,
	0, 1, 2, 1, 3, 3, 3, 3, 5, 0,
	1, 2, 1, 1, 2, 3, 2, 3, 2, 2,
	2, 1, 3, 1, 1, 3, 0, 5, 5, 5,
	1, 3, 0, 2, 1, 3, 3, 2, 3, 3,
	3, 4, 3, 4, 5, 6, 3, 4, 2, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 1, 3,
	1, 3, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 3, 4, 5, 4, 1, 1,
	1, 1, 1, 1, 5, 0, 1, 1, 2, 4,
	0, 2, 1, 3, 1, 1, 1, 1, 0, 3,
	0, 2, 0, 3, 1, 3, 2, 0, 1, 1,
	0, 2, 4, 0, 2, 4, 0, 3, 1, 3,
	0, 5, 2, 1, 1, 3, 3, 1, 1, 3,
	3, 0, 2, 0, 3, 0, 1, 1, 1, 1,
	1, 1, 0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -12,
	-13, -14, -9, -10, -11, -15, -16, 5, 6, 7,
	8, 38, 89, 17, 20, 18, 19, 90, 92, 91,
	93, 102, 103, 104, -19, 55, 56, 57, 58, -17,
	-75, -17, -17, -17, -17, 94, -71, 96, 100, 21,
	-68, 96, 98, 94, 94, 95, 96, 94, -74, -74,
	-74, -3, 22, -20, 23, -18, 34, -30, 40, 9,
	-64, -65, -47, -73, 40, -67, 99, 95, -73, 40,
	94, -73, 40, -66, 99, -73, -66, 40, -21, -22,
	79, -25, 40, -34, -39, -35, 73, 50, -38, -47,
	-43, -46, -73, -44, -48, 25, 41, 42, 43, 30,
	-45, 77, 78, 54, 99, 33, 84, 45, -30, 38,
	82, -30, 59, 51, 82, 40, 73, -73, -74, 40,
	-74, 97, 40, 25, 70, -73, 9, 59, -23, -73,
	24, 82, 72, 71, -36, 26, 73, 28, 29, 27,
	74, 75, 76, 77, 78, 79, 80, 81, 51, 52,
	53, 46, 47, 48, 49, -34, -39, -34, -3, -41,
	-39, -39, 50, 50, 50, -45, 50, -51, -39, -61,
	38, 50, -64, 40, -33, 10, -65, -39, -73, -74,
	25, -72, 101, -69, 92, 90, 37, 91, 13, 40,
	40, 40, -74, -26, -27, -29, 50, 40, -45, -22,
	-73, 79, -34, -34, -39, -40, 50, -45, 44, 26,
	28, 29, -39, -39, 30, 73, -39, -39, -39, -39,
	-39, -39, -39, -39, 105, 105, 59, 105, -39, 105,
	-21, 23, -21, -49, -50, 85, -37, 33, -3, -64,
	-62, -47, -33, -55, 13, -34, 70, -73, -74, -70,
	97, -33, 59, -28, 60, 61, 62, 63, 64, 66,
	67, -24, 40, 24, -27, 82, -41, -40, -39, -39,
	72, 30, -39, 105, -21, 105, -52, -50, 87, -34,
	-63, 70, -42, -43, -63, 105, 59, -55, -59, 15,
	14, 40, 40, -53, 11, -27, -27, 60, 65, 60,
	65, 60, 60, 60, -31, 68, 98, 69, 40, 105,
	40, 105, 72, -39, 105, 88, -39, 86, 35, 59,
	-47, -59, -39, -56, -57, -39, -74, -54, 12, 14,
	70, 60, 60, 95, 95, 95, -39, -39, 36, -43,
	59, 59, -58, 31, 32, -55, -34, -41, -34, 50,
	50, 50, 7, -39, -57, -59, -32, -73, -32, -32,
	-64, -60, 16, 39, 105, 59, 105, 105, 7, 26,
	-73, -73, -73,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 42, 42, 42,
	42, 42, 204, 27, 0, 29, 30, 195, 0, 0,
	0, 209, 209, 209, 0, 46, 48, 49, 50, 51,
	44, 0, 0, 0, 0, 193, 0, 0, 205, 28,
	0, 0, 196, 0, 191, 0, 191, 0, 39, 40,
	41, 18, 47, 0, 52, 43, 0, 0, 84, 0,
	23, 188, 0, 152, 208, 0, 0, 0, 209, 208,
	0, 209, 0, 0, 0, 0, 0, 38, 16, 53,
	55, 60, 208, 58, 59, 94, 0, 0, 122, 123,
	124, 0, 152, 0, 138, 0, 154, 155, 156, 157,
	187, 141, 142, 143, 139, 140, 145, 45, 176, 0,
	0, 92, 0, 0, 0, 209, 0, 206, 26, 0,
	33, 0, 35, 192, 0, 209, 0, 0, 56, 61,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 109, 110,
	111, 112, 113, 114, 115, 97, 0, 0, 0, 0,
	120, 133, 0, 0, 0, 108, 0, 0, 146, 0,
	0, 0, 92, 85, 162, 0, 189, 190, 153, 24,
	194, 0, 0, 209, 202, 197, 198, 199, 200, 201,
	34, 36, 37, 92, 63, 69, 0, 81, 83, 54,
	62, 57, 95, 96, 99, 100, 0, 117, 118, 0,
	0, 0, 102, 0, 106, 0, 125, 126, 127, 128,
	129, 130, 131, 132, 98, 119, 0, 186, 120, 134,
	0, 0, 0, 150, 147, 0, 180, 0, 183, 180,
	0, 178, 162, 170, 0, 93, 0, 207, 31, 0,
	203, 158, 0, 0, 72, 73, 0, 0, 0, 0,
	0, 86, 70, 0, 0, 0, 0, 101, 103, 0,
	0, 107, 121, 135, 0, 137, 0, 148, 0, 0,
	19, 0, 182, 184, 20, 177, 0, 170, 22, 0,
	0, 209, 32, 160, 0, 64, 67, 74, 0, 76,
	0, 78, 79, 80, 65, 0, 0, 0, 71, 66,
	82, 116, 0, 104, 136, 144, 151, 0, 0, 0,
	179, 21, 171, 163, 164, 167, 25, 162, 0, 0,
	0, 75, 77, 0, 0, 0, 105, 149, 0, 185,
	0, 0, 166, 168, 169, 170, 161, 159, 68, 0,
	0, 0, 0, 172, 165, 173, 0, 90, 0, 0,
	181, 17, 0, 0, 87, 0, 88, 89, 174, 0,
	91, 0, 175,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 81, 74, 3,
	50, 105, 79, 77, 59, 78, 82, 80, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	52, 51, 53, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 76, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 75, 3, 54,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 55, 56,
	57, 58, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 73, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104,
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
		//line sql.y:245
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		//line sql.y:251
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		//line sql.y:257
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		//line sql.y:264
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 32:
		//line sql.y:268
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 33:
		//line sql.y:273
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 34:
		//line sql.y:279
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 35:
		//line sql.y:285
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 36:
		//line sql.y:289
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 37:
		//line sql.y:294
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 38:
		//line sql.y:300
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
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
		//line sql.y:314
		{
			yyVAL.statement = &Other{}
		}
	case 42:
		//line sql.y:319
		{
			SetAllowComments(yylex, true)
		}
	case 43:
		//line sql.y:323
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 44:
		//line sql.y:329
		{
			yyVAL.bytes2 = nil
		}
	case 45:
		//line sql.y:333
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 46:
		//line sql.y:339
		{
			yyVAL.str = AST_UNION
		}
	case 47:
		//line sql.y:343
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 48:
		//line sql.y:347
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 49:
		//line sql.y:351
		{
			yyVAL.str = AST_EXCEPT
		}
	case 50:
		//line sql.y:355
		{
			yyVAL.str = AST_INTERSECT
		}
	case 51:
		//line sql.y:360
		{
			yyVAL.str = ""
		}
	case 52:
		//line sql.y:364
		{
			yyVAL.str = AST_DISTINCT
		}
	case 53:
		//line sql.y:370
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 54:
		//line sql.y:374
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 55:
		//line sql.y:380
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 56:
		//line sql.y:384
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 57:
		//line sql.y:388
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 58:
		//line sql.y:394
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 59:
		//line sql.y:398
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 60:
		//line sql.y:403
		{
			yyVAL.bytes = nil
		}
	case 61:
		//line sql.y:407
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 62:
		//line sql.y:411
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 63:
		//line sql.y:417
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 64:
		//line sql.y:421
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 65:
		//line sql.y:427
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 66:
		//line sql.y:431
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 67:
		//line sql.y:435
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 68:
		//line sql.y:439
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 69:
		//line sql.y:444
		{
			yyVAL.bytes = nil
		}
	case 70:
		//line sql.y:448
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 71:
		//line sql.y:452
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 72:
		//line sql.y:458
		{
			yyVAL.str = AST_JOIN
		}
	case 73:
		//line sql.y:462
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 74:
		//line sql.y:466
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 75:
		//line sql.y:470
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 76:
		//line sql.y:474
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 77:
		//line sql.y:478
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 78:
		//line sql.y:482
		{
			yyVAL.str = AST_JOIN
		}
	case 79:
		//line sql.y:486
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 80:
		//line sql.y:490
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 81:
		//line sql.y:496
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 82:
		//line sql.y:500
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 83:
		//line sql.y:504
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 84:
		//line sql.y:510
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 85:
		//line sql.y:514
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 86:
		//line sql.y:519
		{
			yyVAL.indexHints = nil
		}
	case 87:
		//line sql.y:523
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 88:
		//line sql.y:527
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 89:
		//line sql.y:531
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 90:
		//line sql.y:537
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 91:
		//line sql.y:541
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 92:
		//line sql.y:546
		{
			yyVAL.boolExpr = nil
		}
	case 93:
		//line sql.y:550
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 94:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 95:
		//line sql.y:557
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 96:
		//line sql.y:561
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 97:
		//line sql.y:565
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 98:
		//line sql.y:569
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 99:
		//line sql.y:575
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 100:
		//line sql.y:579
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].colTuple}
		}
	case 101:
		//line sql.y:583
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].colTuple}
		}
	case 102:
		//line sql.y:587
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 103:
		//line sql.y:591
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 104:
		//line sql.y:595
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 105:
		//line sql.y:599
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 106:
		//line sql.y:603
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 107:
		//line sql.y:607
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 108:
		//line sql.y:611
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 109:
		//line sql.y:617
		{
			yyVAL.str = AST_EQ
		}
	case 110:
		//line sql.y:621
		{
			yyVAL.str = AST_LT
		}
	case 111:
		//line sql.y:625
		{
			yyVAL.str = AST_GT
		}
	case 112:
		//line sql.y:629
		{
			yyVAL.str = AST_LE
		}
	case 113:
		//line sql.y:633
		{
			yyVAL.str = AST_GE
		}
	case 114:
		//line sql.y:637
		{
			yyVAL.str = AST_NE
		}
	case 115:
		//line sql.y:641
		{
			yyVAL.str = AST_NSE
		}
	case 116:
		//line sql.y:647
		{
			yyVAL.colTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 117:
		//line sql.y:651
		{
			yyVAL.colTuple = yyS[yypt-0].subquery
		}
	case 118:
		//line sql.y:655
		{
			yyVAL.colTuple = ListArg(yyS[yypt-0].bytes)
		}
	case 119:
		//line sql.y:661
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 120:
		//line sql.y:667
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 121:
		//line sql.y:671
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 122:
		//line sql.y:677
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 123:
		//line sql.y:681
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 124:
		//line sql.y:685
		{
			yyVAL.valExpr = yyS[yypt-0].rowTuple
		}
	case 125:
		//line sql.y:689
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 126:
		//line sql.y:693
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 127:
		//line sql.y:697
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 128:
		//line sql.y:701
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 129:
		//line sql.y:705
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:709
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 131:
		//line sql.y:713
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:717
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 133:
		//line sql.y:721
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
	case 134:
		//line sql.y:736
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 135:
		//line sql.y:740
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 136:
		//line sql.y:744
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 137:
		//line sql.y:748
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 138:
		//line sql.y:752
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 139:
		//line sql.y:758
		{
			yyVAL.bytes = IF_BYTES
		}
	case 140:
		//line sql.y:762
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 141:
		//line sql.y:768
		{
			yyVAL.byt = AST_UPLUS
		}
	case 142:
		//line sql.y:772
		{
			yyVAL.byt = AST_UMINUS
		}
	case 143:
		//line sql.y:776
		{
			yyVAL.byt = AST_TILDA
		}
	case 144:
		//line sql.y:782
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 145:
		//line sql.y:787
		{
			yyVAL.valExpr = nil
		}
	case 146:
		//line sql.y:791
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 147:
		//line sql.y:797
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 148:
		//line sql.y:801
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 149:
		//line sql.y:807
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 150:
		//line sql.y:812
		{
			yyVAL.valExpr = nil
		}
	case 151:
		//line sql.y:816
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 152:
		//line sql.y:822
		{
			yyVAL.colName = &ColName{Name: bytes.ToLower(yyS[yypt-0].bytes)}
		}
	case 153:
		//line sql.y:826
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: bytes.ToLower(yyS[yypt-0].bytes)}
		}
	case 154:
		//line sql.y:832
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 155:
		//line sql.y:836
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 156:
		//line sql.y:840
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 157:
		//line sql.y:844
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 158:
		//line sql.y:849
		{
			yyVAL.valExprs = nil
		}
	case 159:
		//line sql.y:853
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 160:
		//line sql.y:858
		{
			yyVAL.boolExpr = nil
		}
	case 161:
		//line sql.y:862
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 162:
		//line sql.y:867
		{
			yyVAL.orderBy = nil
		}
	case 163:
		//line sql.y:871
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 164:
		//line sql.y:877
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 165:
		//line sql.y:881
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 166:
		//line sql.y:887
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 167:
		//line sql.y:892
		{
			yyVAL.str = AST_ASC
		}
	case 168:
		//line sql.y:896
		{
			yyVAL.str = AST_ASC
		}
	case 169:
		//line sql.y:900
		{
			yyVAL.str = AST_DESC
		}
	case 170:
		//line sql.y:905
		{
			yyVAL.limit = nil
		}
	case 171:
		//line sql.y:909
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 172:
		//line sql.y:913
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 173:
		//line sql.y:918
		{
			yyVAL.str = ""
		}
	case 174:
		//line sql.y:922
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 175:
		//line sql.y:926
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
	case 176:
		//line sql.y:939
		{
			yyVAL.columns = nil
		}
	case 177:
		//line sql.y:943
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 178:
		//line sql.y:949
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 179:
		//line sql.y:953
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 180:
		//line sql.y:958
		{
			yyVAL.updateExprs = nil
		}
	case 181:
		//line sql.y:962
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 182:
		//line sql.y:968
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 183:
		//line sql.y:972
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 184:
		//line sql.y:978
		{
			yyVAL.values = Values{yyS[yypt-0].rowTuple}
		}
	case 185:
		//line sql.y:982
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].rowTuple)
		}
	case 186:
		//line sql.y:988
		{
			yyVAL.rowTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 187:
		//line sql.y:992
		{
			yyVAL.rowTuple = yyS[yypt-0].subquery
		}
	case 188:
		//line sql.y:998
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 189:
		//line sql.y:1002
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 190:
		//line sql.y:1008
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 191:
		//line sql.y:1013
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:1015
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:1018
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1020
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1023
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1025
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1029
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1031
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1033
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1035
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1037
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1040
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1042
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1045
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1047
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1050
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1052
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		//line sql.y:1056
		{
			//comment by liuqi    $$ = bytes.ToLower($1)
			yyVAL.bytes = yyS[yypt-0].bytes //add by liuqi
		}
	case 209:
		//line sql.y:1062
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
