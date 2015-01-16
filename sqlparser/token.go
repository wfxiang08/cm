// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ngaut/arena"
	"github.com/wandoulabs/cm/hack"
	"github.com/wandoulabs/cm/sqltypes"
)

const EOFCHAR = 0x100

// Tokenizer is the struct used to generate SQL
// tokens for the parser.
type Tokenizer struct {
	InStream      *strings.Reader
	AllowComments bool
	ForceEOF      bool
	lastChar      uint16
	Position      int
	errorToken    []byte
	LastError     string
	posVarIndex   int
	ParseTree     Statement

	alloc arena.ArenaAllocator
}

// NewStringTokenizer creates a new Tokenizer for the
// sql string.
func NewStringTokenizer(sql string, alloc arena.ArenaAllocator) *Tokenizer {
	return &Tokenizer{
		InStream: strings.NewReader(sql),
		alloc:    alloc,
	}
}

var keywords = map[string]int{
	"by":            BY,
	"for":           FOR,
	"in":            IN,
	"to":            TO,
	"order":         ORDER,
	"using":         USING,
	"as":            AS,
	"distinct":      DISTINCT,
	"index":         INDEX,
	"inner":         INNER,
	"describe":      DESCRIBE,
	"limit":         LIMIT,
	"natural":       NATURAL,
	"ignore":        IGNORE,
	"rename":        RENAME,
	"values":        VALUES,
	"when":          WHEN,
	"else":          ELSE,
	"key":           KEY,
	"lock":          LOCK,
	"unique":        UNIQUE,
	"from":          FROM,
	"not":           NOT,
	"on":            ON,
	"all":           ALL,
	"desc":          DESC,
	"drop":          DROP,
	"intersect":     INTERSECT,
	"minus":         MINUS,
	"right":         RIGHT,
	"select":        SELECT,
	"null":          NULL,
	"show":          SHOW,
	"begin":         BEGIN,
	"rollback":      ROLLBACK,
	"commit":        COMMIT,
	"straight_join": STRAIGHT_JOIN,
	"then":          THEN,
	"duplicate":     DUPLICATE,
	"end":           END,
	"if":            IF,
	"left":          LEFT,
	"union":         UNION,
	"use":           USE,
	"table":         TABLE,
	"alter":         ALTER,
	"cross":         CROSS,
	"is":            IS,
	"join":          JOIN,
	"view":          VIEW,
	"where":         WHERE,
	"asc":           ASC,
	"delete":        DELETE,
	"force":         FORCE,
	"or":            OR,
	"case":          CASE,
	"create":        CREATE,
	"group":         GROUP,
	"having":        HAVING,
	"insert":        INSERT,
	"like":          LIKE,
	"and":           AND,
	"default":       DEFAULT,
	"except":        EXCEPT,
	"explain":       EXPLAIN,
	"between":       BETWEEN,
	"set":           SET,
	"outer":         OUTER,
	"update":        UPDATE,
	"analyze":       ANALYZE,
	"exists":        EXISTS,
	"into":          INTO,
	"BY":            BY,
	"FOR":           FOR,
	"IN":            IN,
	"TO":            TO,
	"AS":            AS,
	"DISTINCT":      DISTINCT,
	"INDEX":         INDEX,
	"INNER":         INNER,
	"ORDER":         ORDER,
	"USING":         USING,
	"DESCRIBE":      DESCRIBE,
	"LIMIT":         LIMIT,
	"NATURAL":       NATURAL,
	"IGNORE":        IGNORE,
	"RENAME":        RENAME,
	"VALUES":        VALUES,
	"WHEN":          WHEN,
	"ELSE":          ELSE,
	"KEY":           KEY,
	"LOCK":          LOCK,
	"UNIQUE":        UNIQUE,
	"FROM":          FROM,
	"ALL":           ALL,
	"DESC":          DESC,
	"DROP":          DROP,
	"INTERSECT":     INTERSECT,
	"NOT":           NOT,
	"ON":            ON,
	"MINUS":         MINUS,
	"RIGHT":         RIGHT,
	"SELECT":        SELECT,
	"DUPLICATE":     DUPLICATE,
	"END":           END,
	"IF":            IF,
	"LEFT":          LEFT,
	"NULL":          NULL,
	"SHOW":          SHOW,
	"STRAIGHT_JOIN": STRAIGHT_JOIN,
	"THEN":          THEN,
	"UNION":         UNION,
	"USE":           USE,
	"ALTER":         ALTER,
	"CROSS":         CROSS,
	"IS":            IS,
	"JOIN":          JOIN,
	"TABLE":         TABLE,
	"ASC":           ASC,
	"DELETE":        DELETE,
	"FORCE":         FORCE,
	"OR":            OR,

	"BEGIN":    BEGIN,
	"ROLLBACK": ROLLBACK,
	"COMMIT":   COMMIT,

	"VIEW":    VIEW,
	"WHERE":   WHERE,
	"CASE":    CASE,
	"CREATE":  CREATE,
	"AND":     AND,
	"DEFAULT": DEFAULT,
	"EXCEPT":  EXCEPT,
	"EXPLAIN": EXPLAIN,
	"GROUP":   GROUP,
	"HAVING":  HAVING,
	"INSERT":  INSERT,
	"LIKE":    LIKE,
	"BETWEEN": BETWEEN,
	"SET":     SET,
	"OUTER":   OUTER,
	"UPDATE":  UPDATE,
	"ANALYZE": ANALYZE,
	"EXISTS":  EXISTS,
	"INTO":    INTO,
}

// Lex returns the next token form the Tokenizer.
// This function is used by go yacc.
func (tkn *Tokenizer) Lex(lval *yySymType) int {
	typ, val := tkn.Scan()
	for typ == COMMENT {
		if tkn.AllowComments {
			break
		}
		typ, val = tkn.Scan()
	}
	switch typ {
	case ID, STRING, NUMBER, VALUE_ARG, LIST_ARG, COMMENT:
		lval.bytes = val
	}
	tkn.errorToken = val
	return typ
}

// Error is called by go yacc if there's a parsing error.
func (tkn *Tokenizer) Error(err string) {
	buf := bytes.NewBuffer(tkn.alloc.AllocBytes(32))
	if tkn.errorToken != nil {
		fmt.Fprintf(buf, "%s at position %v near %s", err, tkn.Position, tkn.errorToken)
	} else {
		fmt.Fprintf(buf, "%s at position %v", err, tkn.Position)
	}
	tkn.LastError = buf.String()
}

func (tkn *Tokenizer) scanHexValue(delim uint16, typ int) (int, []byte) {
	buffer := bytes.NewBuffer(tkn.alloc.AllocBytes(64))
	buffer.WriteString(`x'`)

	tkn.next()
	for {
		ch := tkn.lastChar
		tkn.next()
		if ch == delim {
			if tkn.lastChar == delim {
				tkn.next()
			} else {
				buffer.WriteByte(byte(ch))
				break
			}
		} else if ch == '\'' {
			if tkn.lastChar == EOFCHAR {
				return LEX_ERROR, buffer.Bytes()
			}
			if decodedChar := sqltypes.SqlDecodeMap[byte(tkn.lastChar)]; decodedChar == sqltypes.DONTESCAPE {
				ch = tkn.lastChar
			} else {
				ch = uint16(decodedChar)
			}
			tkn.next()
		}
		if ch == EOFCHAR {
			return LEX_ERROR, buffer.Bytes()
		}
		buffer.WriteByte(byte(ch))
	}
	return typ, buffer.Bytes()
}

// Scan scans the tokenizer for the next token and returns
// the token type and an optional value.
func (tkn *Tokenizer) Scan() (int, []byte) {
	if tkn.ForceEOF {
		return 0, nil
	}

	if tkn.lastChar == 0 {
		tkn.next()
	}
	tkn.skipBlank()
	switch ch := tkn.lastChar; {
	case isLetter(ch):
		if ch == 'x' {
			tkn.next()
			c := tkn.lastChar
			if c == '\'' {
				return tkn.scanHexValue('\'', STRING)
			}

			tkn.unReadByte()
		}
		return tkn.scanIdentifier()
	case isDigit(ch):
		return tkn.scanNumber(false)
	case ch == ':':
		return tkn.scanBindVar()
	default:
		tkn.next()
		switch ch {
		case EOFCHAR:
			return 0, nil
		case '=', ',', ';', '(', ')', '+', '*', '%', '&', '|', '^', '~':
			return int(ch), nil
		case '?':
			tkn.posVarIndex++
			buf := new(bytes.Buffer)
			fmt.Fprintf(buf, ":v%d", tkn.posVarIndex)
			return VALUE_ARG, buf.Bytes()
		case '.':
			if isDigit(tkn.lastChar) {
				return tkn.scanNumber(true)
			} else {
				return int(ch), nil
			}
		case '/':
			switch tkn.lastChar {
			case '/':
				tkn.next()
				return tkn.scanCommentType1("//")
			case '*':
				tkn.next()
				return tkn.scanCommentType2()
			default:
				return int(ch), nil
			}
		case '-':
			if tkn.lastChar == '-' {
				tkn.next()
				return tkn.scanCommentType1("--")
			} else {
				return int(ch), nil
			}
		case '<':
			switch tkn.lastChar {
			case '>':
				tkn.next()
				return NE, nil
			case '=':
				tkn.next()
				switch tkn.lastChar {
				case '>':
					tkn.next()
					return NULL_SAFE_EQUAL, nil
				default:
					return LE, nil
				}
			default:
				return int(ch), nil
			}
		case '>':
			if tkn.lastChar == '=' {
				tkn.next()
				return GE, nil
			} else {
				return int(ch), nil
			}
		case '!':
			if tkn.lastChar == '=' {
				tkn.next()
				return NE, nil
			} else {
				return LEX_ERROR, []byte("!")
			}
		case '\'', '"':
			return tkn.scanString(ch, STRING)
		case '`':
			return tkn.scanLiteralIdentifier()
		default:
			return LEX_ERROR, []byte{byte(ch)}
		}
	}
}

func (tkn *Tokenizer) skipBlank() {
	ch := tkn.lastChar
	for ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' {
		tkn.next()
		ch = tkn.lastChar
	}
}

func (tkn *Tokenizer) scanIdentifier() (int, []byte) {
	buffer := bytes.NewBuffer(tkn.alloc.AllocBytes(128))
	buffer.WriteByte(byte(tkn.lastChar))
	for tkn.next(); isLetter(tkn.lastChar) || isDigit(tkn.lastChar); tkn.next() {
		buffer.WriteByte(byte(tkn.lastChar))
	}

	if keywordId, found := keywords[hack.String(buffer.Bytes())]; found {
		return keywordId, buffer.Bytes()
	}

	return ID, buffer.Bytes()
}

func (tkn *Tokenizer) scanLiteralIdentifier() (int, []byte) {
	buffer := bytes.NewBuffer(tkn.alloc.AllocBytes(128))
	buffer.WriteByte(byte(tkn.lastChar))
	if !isLetter(tkn.lastChar) {
		return LEX_ERROR, buffer.Bytes()
	}
	for tkn.next(); isLetter(tkn.lastChar) || isDigit(tkn.lastChar); tkn.next() {
		buffer.WriteByte(byte(tkn.lastChar))
	}
	if tkn.lastChar != '`' {
		return LEX_ERROR, buffer.Bytes()
	}
	tkn.next()
	return ID, buffer.Bytes()
}

func (tkn *Tokenizer) scanBindVar() (int, []byte) {
	buffer := bytes.NewBuffer(tkn.alloc.AllocBytes(48))
	buffer.WriteByte(byte(tkn.lastChar))
	token := VALUE_ARG
	tkn.next()
	if tkn.lastChar == ':' {
		token = LIST_ARG
		buffer.WriteByte(byte(tkn.lastChar))
		tkn.next()
	}
	if !isLetter(tkn.lastChar) {
		return LEX_ERROR, buffer.Bytes()
	}
	for isLetter(tkn.lastChar) || isDigit(tkn.lastChar) || tkn.lastChar == '.' {
		buffer.WriteByte(byte(tkn.lastChar))
		tkn.next()
	}
	return token, buffer.Bytes()
}

func (tkn *Tokenizer) scanMantissa(base int, buffer *bytes.Buffer) {
	for digitVal(tkn.lastChar) < base {
		tkn.ConsumeNext(buffer)
	}
}

func (tkn *Tokenizer) scanNumber(seenDecimalPoint bool) (int, []byte) {
	buffer := bytes.NewBuffer(tkn.alloc.AllocBytes(64))
	if seenDecimalPoint {
		buffer.WriteByte('.')
		tkn.scanMantissa(10, buffer)
		goto exponent
	}

	if tkn.lastChar == '0' {
		// int or float
		tkn.ConsumeNext(buffer)
		if tkn.lastChar == 'x' || tkn.lastChar == 'X' {
			// hexadecimal int
			tkn.ConsumeNext(buffer)
			tkn.scanMantissa(16, buffer)
		} else {
			// octal int or float
			seenDecimalDigit := false
			tkn.scanMantissa(8, buffer)
			if tkn.lastChar == '8' || tkn.lastChar == '9' {
				// illegal octal int or float
				seenDecimalDigit = true
				tkn.scanMantissa(10, buffer)
			}
			if tkn.lastChar == '.' || tkn.lastChar == 'e' || tkn.lastChar == 'E' {
				goto fraction
			}
			// octal int
			if seenDecimalDigit {
				return LEX_ERROR, buffer.Bytes()
			}
		}
		goto exit
	}

	// decimal int or float
	tkn.scanMantissa(10, buffer)

fraction:
	if tkn.lastChar == '.' {
		tkn.ConsumeNext(buffer)
		tkn.scanMantissa(10, buffer)
	}

exponent:
	if tkn.lastChar == 'e' || tkn.lastChar == 'E' {
		tkn.ConsumeNext(buffer)
		if tkn.lastChar == '+' || tkn.lastChar == '-' {
			tkn.ConsumeNext(buffer)
		}
		tkn.scanMantissa(10, buffer)
	}

exit:
	return NUMBER, buffer.Bytes()
}

func (tkn *Tokenizer) scanString(delim uint16, typ int) (int, []byte) {
	buffer := bytes.NewBuffer(tkn.alloc.AllocBytes(128))
	for {
		ch := tkn.lastChar
		tkn.next()
		if ch == delim {
			if tkn.lastChar == delim {
				tkn.next()
			} else {
				break
			}
		} else if ch == '\\' {
			if tkn.lastChar == EOFCHAR {
				return LEX_ERROR, buffer.Bytes()
			}
			if decodedChar := sqltypes.SqlDecodeMap[byte(tkn.lastChar)]; decodedChar == sqltypes.DONTESCAPE {
				ch = tkn.lastChar
			} else {
				ch = uint16(decodedChar)
			}
			tkn.next()
		}
		if ch == EOFCHAR {
			return LEX_ERROR, buffer.Bytes()
		}
		buffer.WriteByte(byte(ch))
	}
	return typ, buffer.Bytes()
}

func (tkn *Tokenizer) scanCommentType1(prefix string) (int, []byte) {
	buffer := bytes.NewBuffer(tkn.alloc.AllocBytes(48))
	buffer.WriteString(prefix)
	for tkn.lastChar != EOFCHAR {
		if tkn.lastChar == '\n' {
			tkn.ConsumeNext(buffer)
			break
		}
		tkn.ConsumeNext(buffer)
	}
	return COMMENT, buffer.Bytes()
}

func (tkn *Tokenizer) scanCommentType2() (int, []byte) {
	buffer := bytes.NewBuffer(tkn.alloc.AllocBytes(48))
	buffer.WriteString("/*")
	for {
		if tkn.lastChar == '*' {
			tkn.ConsumeNext(buffer)
			if tkn.lastChar == '/' {
				tkn.ConsumeNext(buffer)
				break
			}
			continue
		}
		if tkn.lastChar == EOFCHAR {
			return LEX_ERROR, buffer.Bytes()
		}
		tkn.ConsumeNext(buffer)
	}
	return COMMENT, buffer.Bytes()
}

func (tkn *Tokenizer) ConsumeNext(buffer *bytes.Buffer) {
	if tkn.lastChar == EOFCHAR {
		// This should never happen.
		panic("unexpected EOF")
	}
	buffer.WriteByte(byte(tkn.lastChar))
	tkn.next()
}

func (tkn *Tokenizer) unReadByte() {
	tkn.InStream.UnreadByte()
}

func (tkn *Tokenizer) next() {
	if ch, err := tkn.InStream.ReadByte(); err != nil {
		// Only EOF is possible.
		tkn.lastChar = EOFCHAR
	} else {
		tkn.lastChar = uint16(ch)
	}
	tkn.Position++
}

func isLetter(ch uint16) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '@'
}

func digitVal(ch uint16) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch) - '0'
	case 'a' <= ch && ch <= 'f':
		return int(ch) - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return int(ch) - 'A' + 10
	}
	return 16 // larger than any legal digit val
}

func isDigit(ch uint16) bool {
	return '0' <= ch && ch <= '9'
}
