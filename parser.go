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
	"unicode/utf8"
)

//	chunk ::= block
//
//	block ::= {stat} [retstat]
//
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
//
//	attnamelist ::=  Name attrib {‘,’ Name attrib}
//
//	attrib ::= [‘<’ Name ‘>’]
//
//	retstat ::= return [explist] [‘;’]
//
//	label ::= ‘::’ Name ‘::’
//
//	funcname ::= Name {‘.’ Name} [‘:’ Name]
//
//	varlist ::= var {‘,’ var}
//
//	var ::=  Name | prefixexp ‘[’ exp ‘]’ | prefixexp ‘.’ Name
//
//	namelist ::= Name {‘,’ Name}
//
//	explist ::= exp {‘,’ exp}
//
//	exp ::=  nil | false | true | Numeral | LiteralString | ‘...’ | functiondef | prefixexp | tableconstructor | exp binop exp | unop exp
//
// prefixexp ::= var | functioncall | ‘(’ exp ‘)’
//
//	functioncall ::=  prefixexp args | prefixexp ‘:’ Name args
//
//	args ::=  ‘(’ [explist] ‘)’ | tableconstructor | LiteralString
//
//	functiondef ::= function funcbody
//
//	funcbody ::= ‘(’ [parlist] ‘)’ block end
//
//	parlist ::= namelist [‘,’ ‘...’] | ‘...’
//
//	tableconstructor ::= ‘{’ [fieldlist] ‘}’
//
//	fieldlist ::= field {fieldsep field} [fieldsep]
//
//	field ::= ‘[’ exp ‘]’ ‘=’ exp | Name ‘=’ exp | exp
//
// fieldsep ::= ‘,’ | ‘;’
//
// binop ::=  ‘+’ | ‘-’ | ‘*’ | ‘/’ | ‘//’ | ‘^’ | ‘%’ |
//		 ‘&’ | ‘~’ | ‘|’ | ‘>>’ | ‘<<’ | ‘..’ |
//		 ‘<’ | ‘<=’ | ‘>’ | ‘>=’ | ‘==’ | ‘~=’ |
//		 and | or
//
// unop ::= ‘-’ | not | ‘#’ | ‘~’

type parser struct {
	buf []byte
}

func eof(p parser) bool {
	return len(p.buf) == 0
}

func skipch(p parser) parser {
	if eof(p) {
		return p
	}
	_, w := utf8.DecodeRune(p.buf)
	p.buf = p.buf[w:]
	return p
}

func skipchars(p parser, n int) parser {
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

func (p parser) accept_Keyword(b ...byte) (parser, *KEYWORD) {
	if eof(p) || !bytes.HasPrefix(p.buf, b) {
		return p, nil
	}
	p.buf = p.buf[len(b):]
	return p, &KEYWORD{
		val: b,
	}
}

func (p parser) accept_Literal(b ...byte) (parser, []byte) {
	if eof(p) || !bytes.HasPrefix(p.buf, b) {
		return p, nil
	}
	p.buf = p.buf[len(b):]
	return p, b
}

// LiteralString is a terminal
func (p parser) accept_LiteralString() (parser, []byte) {
	panic("!implemented")
}

// Name is a terminal
func (p parser) accept_Name() (parser, *node) {
	panic("!implemented")
}

// Numeral is a terminal
func (p parser) accept_Numeral() (parser, []byte) {
	panic("!implemented")
}
