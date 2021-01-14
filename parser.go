// GLUA - a Lua-like interpreter
//
// MIT License
//
// Copyright (c) 2021. Michael D Henderson
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"bytes"
)

//	chunk ::= block
type CHUNK struct {
	block *BLOCK
}

//	block ::= {stat} [retstat]
type BLOCK struct {
	stat    []*STAT
	retstat *RETSTAT
}

//	stat ::=  ‘;’ |
//		 varlist ‘=’ explist |
//		 functioncall |
//		 label |
//		 break |
//		 goto Name |
//		 do block end |
//		 while exp do block end |
//		 repeat block until exp |
//		 if exp then block {elseif exp then block} [else block] end |
//		 for Name ‘=’ exp ‘,’ exp [‘,’ exp] do block end |
//		 for namelist in explist do block end |
//		 function funcname funcbody |
//		 local function Name funcbody |
//		 local attnamelist [‘=’ explist]
type STAT struct {
	rule1  *STAT_RULE1
	rule2  *STAT_RULE2
	rule3  *STAT_RULE3
	rule4  *STAT_RULE4
	rule5  *STAT_RULE5
	rule6  *STAT_RULE6
	rule7  *STAT_RULE7
	rule8  *STAT_RULE8
	rule9  *STAT_RULE9
	rule10 *STAT_RULE10
	rule11 *STAT_RULE11
	rule12 *STAT_RULE12
	rule13 *STAT_RULE13
	rule14 *STAT_RULE14
	rule15 *STAT_RULE15
}
type STAT_RULE1 struct{}
type STAT_RULE2 struct {
	varlist *VARLIST
	explist *EXPLIST
}
type STAT_RULE3 struct {
	functioncall *FUNCTIONCALL
}
type STAT_RULE4 struct {
	label *LABEL
}
type STAT_RULE5 struct{}
type STAT_RULE6 struct {
	name *NAME
}
type STAT_RULE7 struct {
	block *BLOCK
}
type STAT_RULE8 struct {
	exp   *EXP
	block *BLOCK
}
type STAT_RULE9 struct {
	block *BLOCK
	exp   *EXP
}
type STAT_RULE10 struct {
	expblock  []*EXPBLOCK
	elseBlock *BLOCK
}
type EXPBLOCK struct {
	exp   *EXP
	block *BLOCK
}
type STAT_RULE11 struct {
	name             *NAME
	exp1, exp2, exp3 *EXP
	block            *BLOCK
}
type STAT_RULE12 struct {
	namelist *NAMELIST
	explist  *EXPLIST
	block    *BLOCK
}
type STAT_RULE13 struct {
	funcname *FUNCNAME
	funcbody *FUNCBODY
}
type STAT_RULE14 struct {
	name     *NAME
	funcbody *FUNCBODY
}
type STAT_RULE15 struct {
	attnamelist *ATTNAMELIST
	explist     *EXPLIST
}

//	attnamelist ::=  Name attrib {‘,’ Name attrib}
type ATTNAMELIST struct {
	attnamelist []*NAMEATTRIB
}
type NAMEATTRIB struct {
	name   *NAME
	attrib *ATTRIB
}

//	attrib ::= [‘<’ Name ‘>’]
type ATTRIB struct {
	name *NAME
}

//	retstat ::= return [explist] [‘;’]
type RETSTAT struct {
	explist *EXPLIST
}

//	label ::= ‘::’ Name ‘::’
type LABEL struct {
	name *NAME
}

//	funcname ::= Name {‘.’ Name} [‘:’ Name]
type FUNCNAME struct {
	name      *NAME
	dotName   []*NAME
	colonName *NAME
}

//	varlist ::= var {‘,’ var}
type VARLIST struct {
	variables []*VARIABLE
}

//	var ::=  Name | prefixexp ‘[’ exp ‘]’ | prefixexp ‘.’ Name
type VARIABLE struct{}

//	namelist ::= Name {‘,’ Name}
type NAMELIST struct {
	names []*NAME
}

//	explist ::= exp {‘,’ exp}
type EXPLIST struct {
	exps []*EXP
}

//	exp ::=  nil | false | true | Numeral | LiteralString | ‘...’ | functiondef | prefixexp | tableconstructor | exp binop exp | unop exp
type EXP struct {
	NIL, FALSE, TRUE, dotDotDot bool
	literal                     []byte
	literalString               *LITERALSTRING
	numeral                     *NUMERAL
	functiondef                 *FUNCTIONDEF
	prefixexp                   *PREFIXEXP
	tableconstructor            *TABLECONSTRUCTOR
	expBinopExp                 *EXP_BINOP_EXP
	unopExp                     *UNOP_EXP
}

type EXP_BINOP_EXP struct {
	exp1  *EXP
	binop *BINOP
	exp2  *EXP
}

type UNOP_EXP struct {
	unop *UNOP
	exp  *EXP
}

// prefixexp ::= var | functioncall | ‘(’ exp ‘)’
type PREFIXEXP struct {
	variable     *VARIABLE
	functioncall *FUNCTIONCALL
	exp          *EXP
}

//	functioncall ::=  prefixexp args | prefixexp ‘:’ Name args
type FUNCTIONCALL struct {
	prefixexp *PREFIXEXP
	name      *NAME
	args      *ARGS
}

//	args ::=  ‘(’ [explist] ‘)’ | tableconstructor | LiteralString
type ARGS struct {
	rule1 *ARGS_RULE1
	rule2 *ARGS_RULE2
	rule3 *ARGS_RULE3
}
type ARGS_RULE1 struct {
	explist *EXPLIST
}
type ARGS_RULE2 struct {
	tableconstructor *TABLECONSTRUCTOR
}
type ARGS_RULE3 struct {
	literalString *LITERALSTRING
}

//	functiondef ::= function funcbody
type FUNCTIONDEF struct {
	funcbody *FUNCBODY
}

//	funcbody ::= ‘(’ [parlist] ‘)’ block end
type FUNCBODY struct {
	parlist *PARLIST
	block   *BLOCK
}

//	parlist ::= namelist [‘,’ ‘...’] | ‘...’
type PARLIST struct {
	rule1 *PARLIST_RULE1
	rule2 *PARLIST_RULE2
}
type PARLIST_RULE1 struct {
	namelist  *NAMELIST
	comma     []byte
	dotDotDot []byte
}
type PARLIST_RULE2 struct {
	dotDotDot []byte
}

//	tableconstructor ::= ‘{’ [fieldlist] ‘}’
type TABLECONSTRUCTOR struct {
	fieldlist *FIELDLIST
}

//	fieldlist ::= field {fieldsep field} [fieldsep]
type FIELDLIST struct {
	fields []*FIELD
}

//	field ::= ‘[’ exp ‘]’ ‘=’ exp | Name ‘=’ exp | exp
type FIELD struct {
	rule1 *FIELD_RULE1
	rule2 *FIELD_RULE2
	rule3 *FIELD_RULE3
}
type FIELD_RULE1 struct {
	exp1 *EXP
	exp2 *EXP
}
type FIELD_RULE2 struct {
	name *NAME
	exp  *EXP
}
type FIELD_RULE3 struct {
	exp *EXP
}

// fieldsep ::= ‘,’ | ‘;’
type FIELDSEP struct {
	comma     bool
	semicolon bool
}

// binop ::=  ‘+’ | ‘-’ | ‘*’ | ‘/’ | ‘//’ | ‘^’ | ‘%’ |
//		 ‘&’ | ‘~’ | ‘|’ | ‘>>’ | ‘<<’ | ‘..’ |
//		 ‘<’ | ‘<=’ | ‘>’ | ‘>=’ | ‘==’ | ‘~=’ |
//		 and | or
type BINOP struct {
	plus                   bool
	hyphen                 bool
	asterisk               bool
	slash                  bool
	slashSlash             bool
	caret                  bool
	percent                bool
	ampersand              bool
	tilde                  bool
	pipe                   bool
	greaterThanGreaterThan bool
	lessThanlessThan       bool
	dotDot                 bool
	lessThan               bool
	lessThanEqual          bool
	greaterThan            bool
	greaterThanEqual       bool
	equalEqual             bool
	tildeEqual             bool
	and                    bool
	or                     bool
}

// unop ::= ‘-’ | not | ‘#’ | ‘~’
type UNOP struct {
	dash  bool
	not   bool
	hash  bool
	tilde bool
}

type parser struct {
	buf []byte
}

func eof(p parser) bool {
	return len(p.buf) == 0
}

func (p parser) skipch(n int) parser {
	if eof(p) {
		return p
	}
	if n > len(p.buf) {
		n = len(p.buf)
	}
	p.buf = p.buf[n:]
	return p
}

type node struct{}

type KEYWORD struct {
	val []byte
}

func bdup(src []byte) []byte {
	dst := make([]byte, len(src), len(src))
	copy(dst, src)
	return dst
}

func (p parser) accept_Keyword(b ...byte) (parser, *KEYWORD) {
	if eof(p) || !bytes.HasPrefix(p.buf, b) {
		return p, nil
	}
	pSaved, keyword := p, &KEYWORD{val: bdup(p.buf[len(b):])}
	// skip the keyword
	p.buf = p.buf[len(b):]
	// the next byte must be a keyword terminator
	if eof(p) {
		return pSaved, nil
	}
	switch p.buf[0] {
	case ' ', '\t', '\r', '\n',
		'(', ')', '{', '}', '<', '>', '[', ']',
		'.', ':', ';', '~', '=', '%', '/', '#', '*',
		'\'', '"':
		return p.skipch(1), keyword
	}
	return pSaved, nil
}

func (p parser) accept_Literal(b ...byte) (parser, []byte) {
	if eof(p) || !bytes.HasPrefix(p.buf, b) {
		return p, nil
	}
	p.buf = p.buf[len(b):]
	return p, b
}

type LITERALSTRING struct {
	level int
	val   []byte
}

// LiteralString is a terminal
func (p parser) accept_LiteralString() (parser, *LITERALSTRING, error) {
	panic("!implemented")
}

type NAME struct {
	val []byte
}

// Name is a terminal
func (p parser) accept_Name() (parser, *NAME, error) {
	if eof(p) {
		return p, nil, nil
	}
	panic("!implemented")
}

type NUMERAL struct {
	val []byte
}

// Numeral is a terminal
func (p parser) accept_Numeral() (parser, *NUMERAL, error) {
	panic("!implemented")
}
