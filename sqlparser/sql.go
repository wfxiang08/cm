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
const UNION = 57393
const MINUS = 57394
const EXCEPT = 57395
const INTERSECT = 57396
const JOIN = 57397
const STRAIGHT_JOIN = 57398
const LEFT = 57399
const RIGHT = 57400
const INNER = 57401
const OUTER = 57402
const CROSS = 57403
const NATURAL = 57404
const USE = 57405
const FORCE = 57406
const ON = 57407
const OR = 57408
const AND = 57409
const NOT = 57410
const UNARY = 57411
const CASE = 57412
const WHEN = 57413
const THEN = 57414
const ELSE = 57415
const END = 57416
const CREATE = 57417
const ALTER = 57418
const DROP = 57419
const RENAME = 57420
const ANALYZE = 57421
const TABLE = 57422
const INDEX = 57423
const VIEW = 57424
const TO = 57425
const IGNORE = 57426
const IF = 57427
const UNIQUE = 57428
const USING = 57429
const SHOW = 57430
const DESCRIBE = 57431
const EXPLAIN = 57432

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

const yyNprod = 213
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 607

var yyAct = []int{

	106, 174, 171, 377, 103, 345, 262, 104, 74, 61,
	309, 222, 299, 211, 251, 253, 75, 191, 184, 386,
	173, 3, 92, 93, 109, 114, 199, 386, 386, 113,
	88, 80, 119, 142, 148, 147, 97, 269, 136, 96,
	110, 111, 112, 243, 62, 63, 356, 305, 77, 101,
	355, 82, 76, 117, 142, 85, 142, 64, 84, 89,
	273, 274, 275, 276, 277, 388, 278, 279, 241, 98,
	243, 354, 100, 387, 385, 81, 115, 116, 94, 335,
	326, 328, 132, 120, 37, 38, 39, 40, 48, 332,
	50, 140, 133, 304, 51, 135, 144, 54, 118, 55,
	294, 60, 292, 56, 175, 330, 148, 147, 176, 336,
	327, 57, 58, 59, 284, 252, 244, 297, 252, 146,
	129, 338, 125, 183, 131, 77, 147, 231, 77, 76,
	195, 194, 76, 189, 242, 180, 83, 170, 172, 351,
	196, 300, 187, 127, 193, 98, 217, 195, 148, 147,
	209, 265, 221, 139, 300, 229, 230, 127, 233, 234,
	235, 236, 237, 238, 239, 240, 216, 215, 364, 365,
	232, 205, 160, 161, 162, 218, 224, 353, 320, 352,
	245, 98, 98, 321, 219, 220, 77, 77, 324, 77,
	76, 258, 323, 76, 256, 203, 71, 260, 206, 318,
	266, 247, 249, 259, 319, 255, 322, 261, 255, 192,
	267, 155, 156, 157, 158, 159, 160, 161, 162, 158,
	159, 160, 161, 162, 243, 285, 245, 283, 270, 264,
	287, 288, 37, 38, 39, 40, 87, 141, 286, 215,
	18, 19, 21, 22, 20, 192, 291, 362, 202, 204,
	201, 98, 224, 25, 27, 28, 26, 340, 271, 128,
	361, 372, 302, 18, 371, 296, 122, 123, 308, 303,
	126, 293, 306, 307, 23, 155, 156, 157, 158, 159,
	160, 161, 162, 214, 225, 316, 317, 142, 370, 298,
	223, 121, 334, 213, 127, 177, 90, 215, 215, 214,
	337, 273, 274, 275, 276, 277, 77, 278, 279, 213,
	341, 188, 185, 343, 346, 181, 179, 178, 282, 342,
	83, 145, 347, 186, 186, 24, 29, 31, 30, 32,
	78, 331, 329, 313, 281, 312, 357, 83, 33, 34,
	35, 358, 208, 383, 207, 190, 137, 134, 360, 130,
	72, 91, 368, 245, 333, 366, 155, 156, 157, 158,
	159, 160, 161, 162, 374, 346, 384, 359, 375, 86,
	124, 378, 378, 378, 77, 379, 380, 376, 76, 339,
	70, 69, 381, 290, 67, 390, 367, 391, 369, 18,
	248, 392, 109, 393, 197, 65, 226, 113, 227, 228,
	119, 138, 113, 52, 310, 119, 350, 96, 110, 111,
	112, 311, 78, 110, 111, 112, 263, 101, 254, 349,
	315, 117, 177, 192, 73, 389, 117, 373, 18, 42,
	289, 18, 155, 156, 157, 158, 159, 160, 161, 162,
	100, 41, 198, 49, 115, 116, 94, 268, 200, 115,
	116, 120, 109, 53, 79, 257, 120, 113, 382, 363,
	119, 43, 44, 45, 46, 47, 118, 78, 110, 111,
	112, 118, 246, 344, 348, 314, 295, 101, 182, 250,
	108, 117, 105, 107, 301, 102, 18, 149, 99, 325,
	212, 109, 272, 210, 95, 280, 113, 143, 66, 119,
	100, 36, 68, 17, 115, 116, 78, 110, 111, 112,
	16, 120, 113, 12, 11, 119, 101, 10, 15, 14,
	117, 13, 78, 110, 111, 112, 118, 9, 8, 5,
	7, 6, 177, 4, 2, 1, 117, 0, 0, 100,
	0, 0, 0, 115, 116, 0, 0, 0, 0, 0,
	120, 150, 154, 152, 153, 0, 0, 0, 0, 115,
	116, 0, 0, 0, 0, 118, 120, 0, 0, 0,
	0, 166, 167, 168, 169, 0, 163, 164, 165, 0,
	0, 118, 155, 156, 157, 158, 159, 160, 161, 162,
	0, 0, 0, 0, 0, 0, 0, 0, 151, 155,
	156, 157, 158, 159, 160, 161, 162,
}
var yyPact = []int{

	235, -1000, -1000, 176, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -7, -1000, 381, -1000, -1000, 0,
	8, 16, 6, -1000, -1000, -1000, 423, 372, -1000, -1000,
	-1000, 360, -1000, 346, 345, 309, 414, 289, -69, -21,
	279, -1000, -1000, -37, 279, -1000, 328, -70, 279, -70,
	310, -1000, -1000, -1000, -1000, -1000, -2, -1000, 245, 309,
	309, 331, 39, 309, 97, -1000, 207, -1000, 37, 308,
	50, 279, -1000, -1000, 306, -1000, -60, 305, 375, 82,
	279, -1000, 227, -1000, -1000, 296, 36, 76, 524, -1000,
	465, 426, -1000, -1000, -1000, 371, 266, 265, -1000, 264,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	371, -1000, 273, 272, 289, 304, 412, 289, 371, 279,
	-1000, 368, -76, -1000, 157, -1000, 303, -1000, -1000, 301,
	-1000, 242, -2, -1000, -1000, 279, 95, 465, 465, 371,
	239, 369, 371, 371, 96, 371, 371, 371, 371, 371,
	371, 371, 371, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 524, -38, 28, 10, 524, -1000, 481, 366, -2,
	-1000, 423, 32, 507, 384, 289, 289, 384, 289, 234,
	-1000, 402, 465, -1000, 507, -1000, -1000, -1000, 80, 279,
	-1000, -61, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	198, 240, 293, 258, 31, -1000, -1000, -1000, -1000, -1000,
	53, 507, -1000, 481, -1000, -1000, 239, 371, 371, 507,
	357, -1000, 352, 141, 141, 141, 92, 92, -1000, -1000,
	-1000, -1000, -1000, 371, -1000, 507, -1000, -4, -2, -6,
	29, -1000, 465, 70, 244, 176, 83, -13, -1000, 70,
	83, 402, 388, 396, 76, 294, -1000, -1000, 292, -1000,
	408, 242, 242, -1000, -1000, 138, 117, 145, 131, 127,
	11, -1000, 291, -1, 290, -17, -1000, 507, 281, 371,
	-1000, 507, -1000, -27, -1000, 20, -1000, 371, 34, -1000,
	343, 197, -1000, -1000, -1000, 289, -1000, -1000, 388, -1000,
	371, 371, -1000, -1000, 406, 391, 240, 68, -1000, 118,
	-1000, 116, -1000, -1000, -1000, -1000, -25, -46, -50, -1000,
	-1000, -1000, -1000, 371, 507, -1000, -1000, 507, 371, 330,
	244, -1000, -1000, 200, 187, -1000, 136, -1000, 402, 465,
	371, 465, -1000, -1000, 237, 213, 210, 507, 507, 420,
	-1000, 371, 371, -1000, -1000, -1000, 388, 76, 164, 76,
	279, 279, 279, 289, 507, -1000, 326, -32, -1000, -33,
	-41, 97, -1000, 418, 358, -1000, 279, -1000, -1000, -1000,
	279, -1000, 279, -1000,
}
var yyPgo = []int{

	0, 535, 534, 20, 533, 531, 530, 529, 528, 527,
	521, 519, 518, 517, 514, 513, 510, 503, 441, 502,
	501, 498, 22, 23, 497, 495, 494, 493, 13, 492,
	490, 196, 489, 3, 17, 36, 488, 487, 15, 485,
	2, 11, 1, 484, 7, 483, 25, 482, 4, 480,
	479, 14, 478, 476, 475, 474, 6, 473, 5, 459,
	10, 458, 18, 455, 12, 8, 16, 236, 454, 453,
	448, 447, 443, 442, 0, 9, 429,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 7, 7, 5, 6, 8, 9, 9, 9,
	13, 13, 14, 15, 10, 10, 10, 11, 12, 12,
	12, 16, 17, 17, 17, 76, 18, 19, 19, 20,
	20, 20, 20, 20, 21, 21, 22, 22, 23, 23,
	23, 26, 26, 24, 24, 24, 27, 27, 28, 28,
	28, 28, 25, 25, 25, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 30, 30, 30, 31, 31, 32,
	32, 32, 32, 33, 33, 34, 34, 35, 35, 35,
	35, 35, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 37, 37, 37, 37, 37, 37, 37, 41,
	41, 41, 46, 42, 42, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 45, 45, 47, 47, 47, 49, 52, 52,
	50, 50, 51, 53, 53, 48, 48, 39, 39, 39,
	39, 54, 54, 55, 55, 56, 56, 57, 57, 58,
	59, 59, 59, 60, 60, 60, 61, 61, 61, 62,
	62, 63, 63, 64, 64, 38, 38, 43, 43, 44,
	44, 65, 65, 66, 67, 67, 68, 68, 69, 69,
	70, 70, 70, 70, 70, 71, 71, 72, 72, 73,
	73, 74, 75,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 7, 7, 8, 7, 3, 5, 8, 4,
	1, 2, 1, 1, 6, 7, 4, 5, 4, 5,
	5, 3, 2, 2, 2, 0, 2, 0, 2, 1,
	2, 1, 1, 1, 0, 1, 1, 3, 1, 2,
	3, 1, 1, 0, 1, 2, 1, 3, 3, 3,
	3, 5, 0, 1, 2, 1, 1, 2, 3, 2,
	3, 2, 2, 2, 1, 3, 1, 1, 3, 0,
	5, 5, 5, 1, 3, 0, 2, 1, 3, 3,
	2, 3, 3, 3, 4, 3, 4, 5, 6, 3,
	4, 2, 1, 1, 1, 1, 1, 1, 1, 3,
	1, 1, 3, 1, 3, 1, 1, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 3, 4, 5,
	4, 1, 1, 1, 1, 1, 1, 5, 0, 1,
	1, 2, 4, 0, 2, 1, 3, 1, 1, 1,
	1, 0, 3, 0, 2, 0, 3, 1, 3, 2,
	0, 1, 1, 0, 2, 4, 0, 2, 4, 0,
	3, 1, 3, 0, 5, 2, 1, 1, 3, 3,
	1, 1, 3, 3, 0, 2, 0, 3, 0, 1,
	1, 1, 1, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -7, -5, -6, -8, -9,
	-13, -14, -15, -10, -11, -12, -16, -17, 5, 6,
	9, 7, 8, 39, 90, 18, 21, 19, 20, 91,
	93, 92, 94, 103, 104, 105, -20, 56, 57, 58,
	59, -18, -76, -18, -18, -18, -18, -18, 95, -72,
	97, 101, 22, -69, 97, 99, 95, 95, 96, 97,
	95, -75, -75, -75, -3, 23, -21, 24, -19, 35,
	35, -31, 41, 10, -65, -66, -48, -74, 41, -68,
	100, 96, -74, 41, 95, -74, 41, -67, 100, -74,
	-67, 41, -22, -23, 80, -26, 41, -35, -40, -36,
	74, 51, -39, -48, -44, -47, -74, -45, -49, 26,
	42, 43, 44, 31, -46, 78, 79, 55, 100, 34,
	85, 46, -31, -31, 39, 83, -31, 60, 52, 83,
	41, 74, -74, -75, 41, -75, 98, 41, 26, 71,
	-74, 10, 60, -24, -74, 25, 83, 73, 72, -37,
	27, 74, 29, 30, 28, 75, 76, 77, 78, 79,
	80, 81, 82, 52, 53, 54, 47, 48, 49, 50,
	-35, -40, -35, -3, -42, -40, -40, 51, 51, 51,
	-46, 51, -52, -40, -62, 39, 51, -62, 39, -65,
	41, -34, 11, -66, -40, -74, -75, 26, -73, 102,
	-70, 93, 91, 38, 92, 14, 41, 41, 41, -75,
	-27, -28, -30, 51, 41, -46, -23, -74, 80, -35,
	-35, -40, -41, 51, -46, 45, 27, 29, 30, -40,
	-40, 31, 74, -40, -40, -40, -40, -40, -40, -40,
	-40, 106, 106, 60, 106, -40, 106, -22, 24, -22,
	-50, -51, 86, -38, 34, -3, -65, -63, -48, -38,
	-65, -34, -56, 14, -35, 71, -74, -75, -71, 98,
	-34, 60, -29, 61, 62, 63, 64, 65, 67, 68,
	-25, 41, 25, -28, 83, -42, -41, -40, -40, 73,
	31, -40, 106, -22, 106, -53, -51, 88, -35, -64,
	71, -43, -44, -64, 106, 60, -64, -64, -56, -60,
	16, 15, 41, 41, -54, 12, -28, -28, 61, 66,
	61, 66, 61, 61, 61, -32, 69, 99, 70, 41,
	106, 41, 106, 73, -40, 106, 89, -40, 87, 36,
	60, -48, -60, -40, -57, -58, -40, -75, -55, 13,
	15, 71, 61, 61, 96, 96, 96, -40, -40, 37,
	-44, 60, 60, -59, 32, 33, -56, -35, -42, -35,
	51, 51, 51, 7, -40, -58, -60, -33, -74, -33,
	-33, -65, -61, 17, 40, 106, 60, 106, 106, 7,
	27, -74, -74, -74,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 45, 45,
	45, 45, 45, 45, 207, 30, 0, 32, 33, 198,
	0, 0, 0, 212, 212, 212, 0, 49, 51, 52,
	53, 54, 47, 0, 0, 0, 0, 0, 196, 0,
	0, 208, 31, 0, 0, 199, 0, 194, 0, 194,
	0, 42, 43, 44, 19, 50, 0, 55, 46, 0,
	0, 0, 87, 0, 26, 191, 0, 155, 211, 0,
	0, 0, 212, 211, 0, 212, 0, 0, 0, 0,
	0, 41, 17, 56, 58, 63, 211, 61, 62, 97,
	0, 0, 125, 126, 127, 0, 155, 0, 141, 0,
	157, 158, 159, 160, 190, 144, 145, 146, 142, 143,
	148, 48, 179, 179, 0, 0, 95, 0, 0, 0,
	212, 0, 209, 29, 0, 36, 0, 38, 195, 0,
	212, 0, 0, 59, 64, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 112, 113, 114, 115, 116, 117, 118,
	100, 0, 0, 0, 0, 123, 136, 0, 0, 0,
	111, 0, 0, 149, 0, 0, 0, 0, 0, 95,
	88, 165, 0, 192, 193, 156, 27, 197, 0, 0,
	212, 205, 200, 201, 202, 203, 204, 37, 39, 40,
	95, 66, 72, 0, 84, 86, 57, 65, 60, 98,
	99, 102, 103, 0, 120, 121, 0, 0, 0, 105,
	0, 109, 0, 128, 129, 130, 131, 132, 133, 134,
	135, 101, 122, 0, 189, 123, 137, 0, 0, 0,
	153, 150, 0, 183, 0, 186, 183, 0, 181, 183,
	183, 165, 173, 0, 96, 0, 210, 34, 0, 206,
	161, 0, 0, 75, 76, 0, 0, 0, 0, 0,
	89, 73, 0, 0, 0, 0, 104, 106, 0, 0,
	110, 124, 138, 0, 140, 0, 151, 0, 0, 20,
	0, 185, 187, 21, 180, 0, 22, 23, 173, 25,
	0, 0, 212, 35, 163, 0, 67, 70, 77, 0,
	79, 0, 81, 82, 83, 68, 0, 0, 0, 74,
	69, 85, 119, 0, 107, 139, 147, 154, 0, 0,
	0, 182, 24, 174, 166, 167, 170, 28, 165, 0,
	0, 0, 78, 80, 0, 0, 0, 108, 152, 0,
	188, 0, 0, 169, 171, 172, 173, 164, 162, 71,
	0, 0, 0, 0, 175, 168, 176, 0, 93, 0,
	0, 184, 18, 0, 0, 90, 0, 91, 92, 177,
	0, 94, 0, 178,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 82, 75, 3,
	51, 106, 80, 78, 60, 79, 83, 81, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	53, 52, 54, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 77, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 76, 3, 55,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 56,
	57, 58, 59, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 73, 74, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105,
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
		yyVAL.statement = yyS[yypt-0].statement
	case 17:
		//line sql.y:179
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 18:
		//line sql.y:183
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 19:
		//line sql.y:187
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 20:
		//line sql.y:193
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:197
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
		//line sql.y:210
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 23:
		//line sql.y:214
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
		//line sql.y:227
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 25:
		//line sql.y:233
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 26:
		//line sql.y:239
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 27:
		//line sql.y:245
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 28:
		//line sql.y:249
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 29:
		//line sql.y:254
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 30:
		//line sql.y:260
		{
			yyVAL.statement = &Begin{}
		}
	case 31:
		//line sql.y:264
		{
			yyVAL.statement = &Begin{}
		}
	case 32:
		//line sql.y:270
		{
			yyVAL.statement = &Commit{}
		}
	case 33:
		//line sql.y:276
		{
			yyVAL.statement = &Rollback{}
		}
	case 34:
		//line sql.y:283
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 35:
		//line sql.y:287
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 36:
		//line sql.y:292
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 37:
		//line sql.y:298
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 38:
		//line sql.y:304
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 39:
		//line sql.y:308
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 40:
		//line sql.y:313
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 41:
		//line sql.y:319
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 42:
		//line sql.y:325
		{
			yyVAL.statement = &Other{}
		}
	case 43:
		//line sql.y:329
		{
			yyVAL.statement = &Other{}
		}
	case 44:
		//line sql.y:333
		{
			yyVAL.statement = &Other{}
		}
	case 45:
		//line sql.y:338
		{
			SetAllowComments(yylex, true)
		}
	case 46:
		//line sql.y:342
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 47:
		//line sql.y:348
		{
			yyVAL.bytes2 = nil
		}
	case 48:
		//line sql.y:352
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 49:
		//line sql.y:358
		{
			yyVAL.str = AST_UNION
		}
	case 50:
		//line sql.y:362
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 51:
		//line sql.y:366
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 52:
		//line sql.y:370
		{
			yyVAL.str = AST_EXCEPT
		}
	case 53:
		//line sql.y:374
		{
			yyVAL.str = AST_INTERSECT
		}
	case 54:
		//line sql.y:379
		{
			yyVAL.str = ""
		}
	case 55:
		//line sql.y:383
		{
			yyVAL.str = AST_DISTINCT
		}
	case 56:
		//line sql.y:389
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 57:
		//line sql.y:393
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 58:
		//line sql.y:399
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 59:
		//line sql.y:403
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 60:
		//line sql.y:407
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 61:
		//line sql.y:413
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 62:
		//line sql.y:417
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 63:
		//line sql.y:422
		{
			yyVAL.bytes = nil
		}
	case 64:
		//line sql.y:426
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 65:
		//line sql.y:430
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 66:
		//line sql.y:436
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 67:
		//line sql.y:440
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 68:
		//line sql.y:446
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 69:
		//line sql.y:450
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 70:
		//line sql.y:454
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 71:
		//line sql.y:458
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 72:
		//line sql.y:463
		{
			yyVAL.bytes = nil
		}
	case 73:
		//line sql.y:467
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 74:
		//line sql.y:471
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 75:
		//line sql.y:477
		{
			yyVAL.str = AST_JOIN
		}
	case 76:
		//line sql.y:481
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 77:
		//line sql.y:485
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 78:
		//line sql.y:489
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 79:
		//line sql.y:493
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 80:
		//line sql.y:497
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 81:
		//line sql.y:501
		{
			yyVAL.str = AST_JOIN
		}
	case 82:
		//line sql.y:505
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 83:
		//line sql.y:509
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 84:
		//line sql.y:515
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 85:
		//line sql.y:519
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 86:
		//line sql.y:523
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 87:
		//line sql.y:529
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 88:
		//line sql.y:533
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 89:
		//line sql.y:538
		{
			yyVAL.indexHints = nil
		}
	case 90:
		//line sql.y:542
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 91:
		//line sql.y:546
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 92:
		//line sql.y:550
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 93:
		//line sql.y:556
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 94:
		//line sql.y:560
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 95:
		//line sql.y:565
		{
			yyVAL.boolExpr = nil
		}
	case 96:
		//line sql.y:569
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 97:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 98:
		//line sql.y:576
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 99:
		//line sql.y:580
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 100:
		//line sql.y:584
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 101:
		//line sql.y:588
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 102:
		//line sql.y:594
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 103:
		//line sql.y:598
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].colTuple}
		}
	case 104:
		//line sql.y:602
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].colTuple}
		}
	case 105:
		//line sql.y:606
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 106:
		//line sql.y:610
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 107:
		//line sql.y:614
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 108:
		//line sql.y:618
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 109:
		//line sql.y:622
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 110:
		//line sql.y:626
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 111:
		//line sql.y:630
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 112:
		//line sql.y:636
		{
			yyVAL.str = AST_EQ
		}
	case 113:
		//line sql.y:640
		{
			yyVAL.str = AST_LT
		}
	case 114:
		//line sql.y:644
		{
			yyVAL.str = AST_GT
		}
	case 115:
		//line sql.y:648
		{
			yyVAL.str = AST_LE
		}
	case 116:
		//line sql.y:652
		{
			yyVAL.str = AST_GE
		}
	case 117:
		//line sql.y:656
		{
			yyVAL.str = AST_NE
		}
	case 118:
		//line sql.y:660
		{
			yyVAL.str = AST_NSE
		}
	case 119:
		//line sql.y:666
		{
			yyVAL.colTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 120:
		//line sql.y:670
		{
			yyVAL.colTuple = yyS[yypt-0].subquery
		}
	case 121:
		//line sql.y:674
		{
			yyVAL.colTuple = ListArg(yyS[yypt-0].bytes)
		}
	case 122:
		//line sql.y:680
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 123:
		//line sql.y:686
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 124:
		//line sql.y:690
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 125:
		//line sql.y:696
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 126:
		//line sql.y:700
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 127:
		//line sql.y:704
		{
			yyVAL.valExpr = yyS[yypt-0].rowTuple
		}
	case 128:
		//line sql.y:708
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 129:
		//line sql.y:712
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:716
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 131:
		//line sql.y:720
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:724
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 133:
		//line sql.y:728
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 134:
		//line sql.y:732
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 135:
		//line sql.y:736
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 136:
		//line sql.y:740
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
	case 137:
		//line sql.y:755
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 138:
		//line sql.y:759
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 139:
		//line sql.y:763
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 140:
		//line sql.y:767
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 141:
		//line sql.y:771
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 142:
		//line sql.y:777
		{
			yyVAL.bytes = IF_BYTES
		}
	case 143:
		//line sql.y:781
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 144:
		//line sql.y:787
		{
			yyVAL.byt = AST_UPLUS
		}
	case 145:
		//line sql.y:791
		{
			yyVAL.byt = AST_UMINUS
		}
	case 146:
		//line sql.y:795
		{
			yyVAL.byt = AST_TILDA
		}
	case 147:
		//line sql.y:801
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 148:
		//line sql.y:806
		{
			yyVAL.valExpr = nil
		}
	case 149:
		//line sql.y:810
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 150:
		//line sql.y:816
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 151:
		//line sql.y:820
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 152:
		//line sql.y:826
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 153:
		//line sql.y:831
		{
			yyVAL.valExpr = nil
		}
	case 154:
		//line sql.y:835
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 155:
		//line sql.y:841
		{
			yyVAL.colName = &ColName{Name: bytes.ToLower(yyS[yypt-0].bytes)}
		}
	case 156:
		//line sql.y:845
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: bytes.ToLower(yyS[yypt-0].bytes)}
		}
	case 157:
		//line sql.y:851
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 158:
		//line sql.y:855
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 159:
		//line sql.y:859
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 160:
		//line sql.y:863
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 161:
		//line sql.y:868
		{
			yyVAL.valExprs = nil
		}
	case 162:
		//line sql.y:872
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 163:
		//line sql.y:877
		{
			yyVAL.boolExpr = nil
		}
	case 164:
		//line sql.y:881
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 165:
		//line sql.y:886
		{
			yyVAL.orderBy = nil
		}
	case 166:
		//line sql.y:890
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 167:
		//line sql.y:896
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 168:
		//line sql.y:900
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 169:
		//line sql.y:906
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 170:
		//line sql.y:911
		{
			yyVAL.str = AST_ASC
		}
	case 171:
		//line sql.y:915
		{
			yyVAL.str = AST_ASC
		}
	case 172:
		//line sql.y:919
		{
			yyVAL.str = AST_DESC
		}
	case 173:
		//line sql.y:924
		{
			yyVAL.limit = nil
		}
	case 174:
		//line sql.y:928
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 175:
		//line sql.y:932
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 176:
		//line sql.y:937
		{
			yyVAL.str = ""
		}
	case 177:
		//line sql.y:941
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 178:
		//line sql.y:945
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
	case 179:
		//line sql.y:958
		{
			yyVAL.columns = nil
		}
	case 180:
		//line sql.y:962
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 181:
		//line sql.y:968
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 182:
		//line sql.y:972
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 183:
		//line sql.y:977
		{
			yyVAL.updateExprs = nil
		}
	case 184:
		//line sql.y:981
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 185:
		//line sql.y:987
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 186:
		//line sql.y:991
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 187:
		//line sql.y:997
		{
			yyVAL.values = Values{yyS[yypt-0].rowTuple}
		}
	case 188:
		//line sql.y:1001
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].rowTuple)
		}
	case 189:
		//line sql.y:1007
		{
			yyVAL.rowTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 190:
		//line sql.y:1011
		{
			yyVAL.rowTuple = yyS[yypt-0].subquery
		}
	case 191:
		//line sql.y:1017
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 192:
		//line sql.y:1021
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 193:
		//line sql.y:1027
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 194:
		//line sql.y:1032
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1034
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1037
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1039
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1042
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1044
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1048
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1050
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1052
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1054
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1056
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1059
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1061
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1064
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		//line sql.y:1066
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		//line sql.y:1069
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		//line sql.y:1071
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		//line sql.y:1075
		{
			//comment by liuqi    $$ = bytes.ToLower($1)
			yyVAL.bytes = yyS[yypt-0].bytes //add by liuqi
		}
	case 212:
		//line sql.y:1081
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
