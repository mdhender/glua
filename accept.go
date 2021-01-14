// GLUA - a Lua-like interpreter
//
// MIT License
//
// Copyright (c) 2021 Michael D Henderson
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
	"fmt"
)

func (p parser) accept() (parser, *node, error) {
	if eof(p) {
		return p, nil, nil
	}
	return p, nil, nil
}

// chunk ::= block
func (p parser) accept_chunk() (parser, *CHUNK, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, chunk := p, &CHUNK{}
	var err error

	p, chunk.block, err = p.accept_block()
	if err != nil {
		return pSaved, nil, err
	}
	if chunk.block == nil {
		return pSaved, nil, nil
	}
	return p, chunk, nil
}

// block ::= {stat} [retstat]
func (p parser) accept_block() (parser, *BLOCK, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, block := p, &BLOCK{}

	// accept {stat}
	for {
		pp, stat, err := p.accept_stat()
		if err != nil {
			return pSaved, nil, err
		} else if stat == nil {
			break
		}
		block.stat = append(block.stat, stat)
		p = pp
	}

	if pp, retstat, err := p.accept_retstat(); err != nil {
		return pSaved, nil, err
	} else if retstat != nil {
		block.retstat = retstat
		p = pp
	}

	return p, block, nil
}

// stat ::= ';'
//        | varlist '=' explist
//        | functioncall
//        | label
//        | 'break'
//        | 'goto' Name
//        | 'do' block 'end'
//        | 'while' exp 'do' block 'end'
//        | 'repeat' block 'until' exp
//        | 'if' exp 'then' block {'elseif' exp 'then' block} ['else' block] 'end'
//        | 'for' Name '=' exp ',' exp [',' exp] 'do' block 'end'
//        | 'for' namelist 'in' explist 'do' block 'end'
//        | 'function' funcname funcbody
//        | 'local' 'function' Name funcbody
//        | 'local' attnamelist ['=' explist]
func (p parser) accept_stat() (parser, *STAT, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved := p

	// accept stat.rule1
	if pp, rule, err := p.accept_stat_rule1(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule1: rule}, nil
	}

	// accept stat.rule2
	if pp, rule, err := p.accept_stat_rule2(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule2: rule}, nil
	}

	// accept stat.rule3
	if pp, rule, err := p.accept_stat_rule3(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule3: rule}, nil
	}

	// accept stat.rule4
	if pp, rule, err := p.accept_stat_rule4(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule4: rule}, nil
	}

	// accept stat.rule5
	if pp, rule, err := p.accept_stat_rule5(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule5: rule}, nil
	}

	// accept stat.rule6
	if pp, rule, err := p.accept_stat_rule6(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule6: rule}, nil
	}

	// accept stat.rule7
	if pp, rule, err := p.accept_stat_rule7(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule7: rule}, nil
	}

	// accept stat.rule8
	if pp, rule, err := p.accept_stat_rule8(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule8: rule}, nil
	}

	// accept stat.rule9
	if pp, rule, err := p.accept_stat_rule9(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule9: rule}, nil
	}

	// accept stat.rule10
	if pp, rule, err := p.accept_stat_rule10(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule10: rule}, nil
	}

	// accept stat.rule11
	if pp, rule, err := p.accept_stat_rule11(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule11: rule}, nil
	}

	// accept stat.rule12
	if pp, rule, err := p.accept_stat_rule12(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule12: rule}, nil
	}

	// accept stat.rule13
	if pp, rule, err := p.accept_stat_rule13(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule13: rule}, nil
	}

	// accept stat.rule13
	if pp, rule, err := p.accept_stat_rule13(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule13: rule}, nil
	}

	// accept stat.rule14
	if pp, rule, err := p.accept_stat_rule14(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule14: rule}, nil
	}

	// accept stat.rule15
	if pp, rule, err := p.accept_stat_rule15(); err != nil {
		return pSaved, nil, err
	} else if rule != nil {
		return pp, &STAT{rule15: rule}, nil
	}

	return pSaved, nil, nil
}

// stat.rule1 ::= ';'
func (p parser) accept_stat_rule1() (parser, *STAT_RULE1, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved := p

	// accept ';'
	var semiColon []byte
	if p, semiColon = p.accept_Literal(';'); semiColon == nil {
		return pSaved, nil, nil
	}

	return pSaved, &STAT_RULE1{}, nil
}

// stat.rule2 ::= varlist '=' explist
func (p parser) accept_stat_rule2() (parser, *STAT_RULE2, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, rule := p, &STAT_RULE2{}
	var err error

	// accept varlist
	if p, rule.varlist, err = p.accept_varlist(); err != nil {
		return pSaved, nil, err
	} else if rule.varlist == nil {
		return pSaved, nil, nil
	}
	// expect '='
	var equals []byte
	if p, equals = p.accept_Literal('='); equals == nil {
		return pSaved, nil, fmt.Errorf("expected '='")
	}
	// accept explist
	if p, rule.explist, err = p.accept_explist(); err != nil {
		return pSaved, nil, err
	} else if rule.explist == nil {
		return pSaved, nil, fmt.Errorf("expected explist")
	}

	return p, rule, nil
}

// stat.rule3 ::= functioncall
func (p parser) accept_stat_rule3() (parser, *STAT_RULE3, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, rule := p, &STAT_RULE3{}
	var err error

	// accept functioncall
	if p, rule.functioncall, err = p.accept_functioncall(); err != nil {
		return pSaved, nil, err
	} else if rule.functioncall == nil {
		return pSaved, nil, nil
	}

	return p, rule, nil
}

// stat.rule4 ::= label
func (p parser) accept_stat_rule4() (parser, *STAT_RULE4, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, rule := p, &STAT_RULE4{}
	var err error

	// accept label
	if p, rule.label, err = p.accept_label(); err != nil {
		return pSaved, nil, err
	} else if rule.label == nil {
		return pSaved, nil, nil
	}

	return p, rule, nil
}

// stat.rule5 ::= 'break'
func (p parser) accept_stat_rule5() (parser, *STAT_RULE5, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved := p

	// accept 'break'
	var kwBreak *KEYWORD
	if p, kwBreak = p.accept_Keyword('b', 'r', 'e', 'a', 'k'); kwBreak == nil {
		return pSaved, nil, nil
	}

	return p, &STAT_RULE5{}, nil
}

// stat.rule6 ::= 'goto' Name
func (p parser) accept_stat_rule6() (parser, *STAT_RULE6, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, rule := p, &STAT_RULE6{}
	var err error

	// accept 'goto'
	var kwGoto *KEYWORD
	if p, kwGoto = p.accept_Keyword('g', 'o', 't', 'o'); kwGoto == nil {
		return pSaved, nil, nil
	}
	// expect Name
	if p, rule.name, err = p.accept_Name(); err != nil {
		return pSaved, nil, err
	} else if rule.name == nil {
		return pSaved, nil, fmt.Errorf("expected name")
	}

	return p, rule, nil
}

// stat.rule7 ::= 'do' block 'end'
func (p parser) accept_stat_rule7() (parser, *STAT_RULE7, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, rule := p, &STAT_RULE7{}
	var err error

	// accept 'do'
	var kwDo *KEYWORD
	if p, kwDo = p.accept_Keyword('d', 'o'); kwDo == nil {
		return pSaved, nil, nil
	}
	// expect block
	if p, rule.block, err = p.accept_block(); err != nil {
		return pSaved, nil, err
	} else if rule.block == nil {
		return pSaved, nil, fmt.Errorf("expected block")
	}
	// expect 'end'
	var kwEnd *KEYWORD
	if p, kwEnd = p.accept_Keyword('e', 'n', 'd'); kwEnd == nil {
		return pSaved, nil, nil
	}

	return p, rule, nil
}

// stat.rule8 ::= 'while' exp 'do' block 'end'
func (p parser) accept_stat_rule8() (parser, *STAT_RULE8, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, rule := p, &STAT_RULE8{}
	var err error

	// accept 'while'
	var kwWhile *KEYWORD
	if p, kwWhile = p.accept_Keyword('w', 'h', 'i', 'l', 'e'); kwWhile == nil {
		return pSaved, nil, nil
	}
	// expect exp
	if p, rule.exp, err = p.accept_exp(); err != nil {
		return pSaved, nil, err
	} else if rule.exp == nil {
		return pSaved, nil, fmt.Errorf("expected exp")
	}
	// expect 'do'
	var kwDo *KEYWORD
	if p, kwDo = p.accept_Keyword('d', 'o'); kwDo == nil {
		return pSaved, nil, fmt.Errorf("expected 'do'")
	}
	// expect block
	if p, rule.block, err = p.accept_block(); err != nil {
		return pSaved, nil, err
	} else if rule.block == nil {
		return pSaved, nil, fmt.Errorf("expected block")
	}
	// expect 'end'
	var kwEnd *KEYWORD
	if p, kwEnd = p.accept_Keyword('e', 'n', 'd'); kwEnd == nil {
		return pSaved, nil, nil
	}

	return p, rule, nil
}

// stat.rule9 ::= 'repeat' block 'until' exp
func (p parser) accept_stat_rule9() (parser, *STAT_RULE9, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, rule := p, &STAT_RULE9{}
	var err error

	// accept 'repeat'
	var kwRepeat *KEYWORD
	if p, kwRepeat = p.accept_Keyword('r', 'e', 'p', 'e', 'a', 't'); kwRepeat == nil {
		return pSaved, nil, nil
	}
	// expect block
	if p, rule.block, err = p.accept_block(); err != nil {
		return pSaved, nil, err
	} else if rule.block == nil {
		return pSaved, nil, fmt.Errorf("expected block")
	}
	// expect 'until'
	var kwUntil *KEYWORD
	if p, kwUntil = p.accept_Keyword('u', 'n', 't', 'i', 'l'); kwUntil == nil {
		return pSaved, nil, fmt.Errorf("expected 'until'")
	}
	// expect exp
	if p, rule.exp, err = p.accept_exp(); err != nil {
		return pSaved, nil, err
	} else if rule.exp == nil {
		return pSaved, nil, fmt.Errorf("expected exp")
	}

	return p, rule, nil
}

// stat.rule10 ::= 'if' exp 'then' block {'elseif' exp 'then' block} ['else' block] 'end'
func (p parser) accept_stat_rule10() (parser, *STAT_RULE10, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, rule := p, &STAT_RULE10{}
	var err error

	// accept 'if'
	var kwIf *KEYWORD
	if p, kwIf = p.accept_Keyword('i', 'f'); kwIf == nil {
		return pSaved, nil, nil
	}
	// expect exp
	var exp *EXP
	if p, exp, err = p.accept_exp(); err != nil {
		return pSaved, nil, err
	} else if exp == nil {
		return pSaved, nil, fmt.Errorf("expected exp")
	}
	// expect 'then'
	var kwThen *KEYWORD
	if p, kwThen = p.accept_Keyword('t', 'h', 'e', 'n'); kwThen == nil {
		return pSaved, nil, fmt.Errorf("expected 'then'")
	}
	// expect block
	var block *BLOCK
	if p, block, err = p.accept_block(); err != nil {
		return pSaved, nil, err
	} else if block == nil {
		return pSaved, nil, fmt.Errorf("expected block")
	}
	rule.expblock = append(rule.expblock, &EXPBLOCK{exp: exp, block: block})
	// accept {'elseif' exp 'then' block}
	for {
		// accept 'elseif' exp 'then' block
		pp, kwElseIf := p.accept_Keyword('e', 'l', 's', 'e', 'i', 'f')
		if kwElseIf == nil {
			break
		}
		// expect exp
		if p, block, err = pp.accept_block(); err != nil {
			return pSaved, nil, err
		} else if block == nil {
			return pSaved, nil, fmt.Errorf("expected block")
		}
		// expect 'then'
		if p, kwThen = p.accept_Keyword('t', 'h', 'e', 'n'); kwThen == nil {
			return pSaved, nil, fmt.Errorf("expected 'then'")
		}
		// expect block
		if p, block, err = pp.accept_block(); err != nil {
			return pSaved, nil, err
		} else if block == nil {
			return pSaved, nil, fmt.Errorf("expected block")
		}
		rule.expblock = append(rule.expblock, &EXPBLOCK{exp: exp, block: block})
	}
	// accept ['else' block]
	if pp, kwElse := p.accept_Keyword('e', 'l', 's', 'e'); kwElse != nil {
		// expect block
		if p, rule.elseBlock, err = pp.accept_block(); err != nil {
			return pSaved, nil, err
		} else if rule.elseBlock == nil {
			return pSaved, nil, fmt.Errorf("expected block")
		}
	}
	// expect 'end'
	var kwEnd *KEYWORD
	if p, kwEnd = p.accept_Keyword('e', 'n', 'd'); kwEnd == nil {
		return pSaved, nil, fmt.Errorf("expected 'end'")
	}

	return p, rule, nil
}

// stat.rule11 ::= 'for' Name '=' exp ',' exp [',' exp] 'do' block 'end'
func (p parser) accept_stat_rule11() (parser, *STAT_RULE11, error) {
	if eof(p) {
		return p, nil, nil
	}
	panic("!")
}

// stat.rule12 ::= 'for' namelist 'in' explist 'do' block 'end'
func (p parser) accept_stat_rule12() (parser, *STAT_RULE12, error) {
	if eof(p) {
		return p, nil, nil
	}
	panic("!")
}

// stat.rule13 ::= 'function' funcname funcbody
func (p parser) accept_stat_rule13() (parser, *STAT_RULE13, error) {
	if eof(p) {
		return p, nil, nil
	}
	panic("!")
}

// stat.rule14 ::= 'local' 'function' Name funcbody
func (p parser) accept_stat_rule14() (parser, *STAT_RULE14, error) {
	if eof(p) {
		return p, nil, nil
	}
	panic("!")
}

// stat.rule15 ::= 'local' attnamelist ['=' explist]
func (p parser) accept_stat_rule15() (parser, *STAT_RULE15, error) {
	if eof(p) {
		return p, nil, nil
	}
	panic("!")
}

// attnamelist ::= Name attrib {',' Name attrib}
func (p parser) accept_attnamelist() (parser, *ATTNAMELIST, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, attnamelist := p, &ATTNAMELIST{}
	var err error

	// accept Name
	var name *NAME
	if p, name, err = p.accept_Name(); err != nil {
		return pSaved, nil, err
	} else if name == nil {
		return pSaved, nil, nil
	}
	// accept attrib
	var attrib *ATTRIB
	if p, attrib, err = p.accept_attrib(); err != nil {
		return pSaved, nil, err
	} else if attrib == nil {
		return pSaved, nil, nil
	}
	attnamelist.attnamelist = append(attnamelist.attnamelist, &NAMEATTRIB{name: name, attrib: attrib})
	// accept {',' Name attrib}
	for {
		// accept ','
		pp, comma := p.accept_Literal(',')
		if comma == nil {
			break
		}
		// expect Name
		if p, name, err = pp.accept_Name(); err != nil {
			return pSaved, nil, err
		} else if name == nil {
			return pSaved, nil, fmt.Errorf("expected name")
		}
		// expect attrib
		if p, attrib, err = p.accept_attrib(); err != nil {
			return pSaved, nil, err
		} else if attrib == nil {
			return pSaved, nil, fmt.Errorf("expected attrib")
		}
		attnamelist.attnamelist = append(attnamelist.attnamelist, &NAMEATTRIB{name: name, attrib: attrib})
	}

	return p, attnamelist, nil
}

// attrib ::= ['<' Name '>']
func (p parser) accept_attrib() (parser, *ATTRIB, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, attrib := p, &ATTRIB{}
	var err error

	// accept '<'
	var oBroket []byte
	if p, oBroket = p.accept_Literal('<'); oBroket == nil {
		return pSaved, nil, nil
	}
	// expect Name
	if p, attrib.name, err = p.accept_Name(); err != nil {
		return pSaved, nil, err
	} else if attrib.name == nil {
		return pSaved, nil, fmt.Errorf("expected name")
	}
	// expect '>'
	var cBroket []byte
	if p, cBroket = p.accept_Literal(); cBroket == nil {
		return pSaved, nil, fmt.Errorf("expected '<'")
	}

	return p, attrib, nil
}

// retstat ::= 'return' [explist] [';']
func (p parser) accept_retstat() (parser, *RETSTAT, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, retstat := p, &RETSTAT{}
	var err error

	// accept 'return'
	var kwReturn *KEYWORD
	if p, kwReturn = p.accept_Keyword('r', 'e', 't', 'u', 'r', 'n'); kwReturn == nil {
		return pSaved, nil, nil
	}

	// accept [explist]
	if p, retstat.explist, err = p.accept_explist(); err != nil {
		return pSaved, nil, err
	}

	// accept [';']
	p, _ = p.accept_Literal(';')

	return p, retstat, nil
}

// label ::= '::' Name '::'
func (p parser) accept_label() (parser, *LABEL, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, label := p, &LABEL{}
	var err error

	// accept '::
	var colonColon []byte
	if p, colonColon = p.accept_Literal(':', ':'); colonColon == nil {
		return pSaved, nil, nil
	}
	// expect Name
	if p, label.name, err = p.accept_Name(); err != nil {
		return pSaved, nil, err
	} else if label.name == nil {
		return pSaved, nil, fmt.Errorf("expected name")
	}
	// expect '::
	if p, colonColon = p.accept_Literal(':', ':'); colonColon == nil {
		return pSaved, nil, fmt.Errorf("expected '::'")
	}

	return p, label, nil
}

// funcname ::= Name {'.' Name} [':' Name]
func (p parser) accept_funcname() (parser, *FUNCNAME, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, funcname := p, &FUNCNAME{}
	var err error

	// accept Name
	if p, funcname.name, err = p.accept_Name(); err != nil {
		return pSaved, nil, err
	}

	// accept {'.' Name}
	for {
		// accept '.'
		pp, dot := p.accept_Literal('.')
		if dot == nil {
			break
		}
		// expect Name
		var name *NAME
		if p, name, err = pp.accept_Name(); err != nil {
			return pSaved, nil, err
		} else if name == nil {
			return pSaved, nil, fmt.Errorf("expected name")
		}
		funcname.dotName = append(funcname.dotName, name)
	}

	// accept [':' Name]
	pp, colon := p.accept_Literal(':')
	if colon == nil {
		return p, funcname, nil
	}
	// expect Name
	if p, funcname.colonName, err = pp.accept_Name(); err != nil {
		return pSaved, nil, err
	} else if funcname.colonName == nil {
		return pSaved, nil, fmt.Errorf("expected name")
	}

	return p, funcname, nil
}

// varlist ::= var {',' var}
func (p parser) accept_varlist() (parser, *VARLIST, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, varlist := p, &VARLIST{}
	var err error

	// accept var
	var variable *VARIABLE
	if p, variable, err = p.accept_var(); err != nil {
		return pSaved, nil, err
	} else if variable == nil {
		return pSaved, nil, nil
	}
	varlist.variables = append(varlist.variables, variable)

	// accept {',' var}
	for {
		pp, comma := p.accept_Literal(',')
		if comma == nil {
			break
		}
		if pp, variable, err = pp.accept_var(); err != nil {
			return pSaved, nil, err
		} else if variable == nil {
			return pSaved, nil, fmt.Errorf("expected var")
		}
		p, varlist.variables = pp, append(varlist.variables, variable)
	}

	return p, varlist, nil
}

// var ::= Name
//       | prefixexp '[' exp ']'
//       | prefixexp '.' Name
func (p parser) accept_var() (parser, *VARIABLE, error) {
	if eof(p) {
		return p, nil, nil
	}
	panic("!implemented")
}

// namelist ::= Name {',' Name}
func (p parser) accept_namelist() (parser, *NAMELIST, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, namelist := p, &NAMELIST{}
	var err error

	// accept Name
	var name *NAME
	if p, name, err = p.accept_Name(); err != nil {
		return pSaved, nil, err
	} else if name == nil {
		return pSaved, nil, nil
	}
	namelist.names = append(namelist.names, name)

	// accept {',' Name}
	for {
		pp, comma := p.accept_Literal(',')
		if comma == nil {
			break
		}
		if pp, name, err = pp.accept_Name(); err != nil {
			return pSaved, nil, err
		} else if name == nil {
			return pSaved, nil, fmt.Errorf("expected Name")
		}
		p, namelist.names = pp, append(namelist.names, name)
	}

	return p, namelist, nil
}

// explist ::= exp {',' exp}
func (p parser) accept_explist() (parser, *EXPLIST, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, explist := p, &EXPLIST{}
	var err error

	// accept exp
	var exp *EXP
	if p, exp, err = p.accept_exp(); err != nil {
		return pSaved, nil, err
	} else if exp == nil {
		return pSaved, nil, nil
	}
	explist.exps = append(explist.exps, exp)

	// accept {',' exp}
	for {
		pp, comma := p.accept_Literal(',')
		if comma == nil {
			break
		}
		if pp, exp, err = pp.accept_exp(); err != nil {
			return pSaved, nil, err
		} else if exp == nil {
			return pSaved, nil, fmt.Errorf("expected exp")
		}
		p, explist.exps = pp, append(explist.exps, exp)
	}

	return p, explist, nil
}

// exp ::= 'nil'
//       | 'false'
//       | 'true'
//       | Numeral
//       | LiteralString
//       | '...'
//       | functiondef
//       | prefixexp
//       | tableconstructor
//       | exp binop exp
//       | unop exp
func (p parser) accept_exp() (parser, *EXP, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, exp := p, &EXP{}
	var err error

	var literal []byte
	if p, literal = p.accept_Literal('n', 'i', 'l'); literal != nil {
		return p, &EXP{NIL: true}, nil
	}
	if p, literal = p.accept_Literal('f', 'a', 'l', 's', 'e'); literal != nil {
		return p, &EXP{FALSE: true}, nil
	}
	if p, literal = p.accept_Literal('t', 'r', 'u', 'e'); literal != nil {
		return p, &EXP{TRUE: true}, nil
	}
	if p, exp.numeral, err = p.accept_Numeral(); err != nil {
		return pSaved, nil, err
	} else if exp.numeral != nil {
		return p, exp, nil
	}
	if p, exp.literalString, err = p.accept_LiteralString(); err != nil {
		return pSaved, nil, err
	} else if exp.literalString != nil {
		return p, exp, nil
	}
	if p, literal = p.accept_Literal('.', '.', '.'); literal != nil {
		return p, &EXP{dotDotDot: true}, nil
	}
	if p, exp.functiondef, err = p.accept_functiondef(); err != nil {
		return pSaved, nil, err
	} else if exp.functiondef != nil {
		return p, exp, nil
	}
	if p, exp.prefixexp, err = p.accept_prefixexp(); err != nil {
		return pSaved, nil, err
	} else if exp.prefixexp != nil {
		return p, exp, nil
	}
	if p, exp.tableconstructor, err = p.accept_tableconstructor(); err != nil {
		return pSaved, nil, err
	} else if exp.tableconstructor != nil {
		return p, exp, nil
	}

	// accept exp binop exp
	if pp, exp1, err := p.accept_exp(); err != nil {
		return pSaved, nil, err
	} else if exp1 != nil {
		exp.expBinopExp = &EXP_BINOP_EXP{exp1: exp1}
		// expect binop
		if pp, exp.expBinopExp.binop, err = p.accept_binop(); err != nil {
			return pSaved, nil, err
		} else if exp.expBinopExp.binop == nil {
			return pSaved, nil, fmt.Errorf("expected binop")
		}
		// expect exp
		if pp, exp.expBinopExp.exp2, err = p.accept_exp(); err != nil {
			return pSaved, nil, err
		} else if exp.expBinopExp.exp2 == nil {
			return pSaved, nil, fmt.Errorf("expected exp")
		}
		return pp, exp, nil
	}

	// accept unop exp
	if pp, unop, err := p.accept_unop(); err != nil {
		return pSaved, nil, err
	} else if unop != nil {
		exp.unopExp = &UNOP_EXP{unop: unop}
		// expect exp
		if pp, exp.unopExp.exp, err = p.accept_exp(); err != nil {
			return pSaved, nil, err
		} else if exp.unopExp.exp == nil {
			return pSaved, nil, fmt.Errorf("expected exp")
		}
		return pp, exp, nil
	}

	return pSaved, nil, nil
}

// prefixexp ::= var
//             | functioncall
//             | '(' exp ')'
func (p parser) accept_prefixexp() (parser, *PREFIXEXP, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, prefixexp := p, &PREFIXEXP{}
	var err error

	// accept var
	p, prefixexp.variable, err = p.accept_var()
	if err != nil {
		return pSaved, nil, err
	}
	if prefixexp.variable != nil {
		return p, prefixexp, nil
	}

	// accept functioncall
	p, prefixexp.functioncall, err = p.accept_functioncall()
	if err != nil {
		return pSaved, nil, err
	}
	if prefixexp.functioncall != nil {
		return p, prefixexp, nil
	}

	// accept '(' exp ')'
	var oParen []byte
	p, oParen = p.accept_Literal('(')
	if oParen == nil {
		return pSaved, nil, nil
	}
	// expect exp
	p, prefixexp.exp, err = p.accept_exp()
	if prefixexp.exp == nil {
		return pSaved, nil, fmt.Errorf("expected exp")
	}
	// expect ')'
	var cParen []byte
	p, cParen = p.accept_Literal('(')
	if cParen == nil {
		return pSaved, nil, fmt.Errorf("expected ')'")
	}

	return p, prefixexp, nil
}

// functioncall ::= prefixexp args
//                | prefixexp ':' Name args
func (p parser) accept_functioncall() (parser, *FUNCTIONCALL, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, functioncall := p, &FUNCTIONCALL{}
	var err error

	// accept prefixexp
	p, functioncall.prefixexp, err = p.accept_prefixexp()
	if err != nil {
		return pSaved, nil, err
	}
	if functioncall.prefixexp == nil {
		return pSaved, nil, nil
	}

	// accept [':' Name]
	if pp, colon := p.accept_Literal(':'); colon != nil {
		// expect Name
		p, functioncall.name, err = pp.accept_Name()
		if err != nil {
			return pSaved, nil, err
		}
		if functioncall.name == nil {
			return pSaved, nil, fmt.Errorf("expected Name")
		}
	}

	// expect args
	p, functioncall.args, err = p.accept_args()
	if err != nil {
		return pSaved, nil, err
	}
	if functioncall.args == nil {
		return pSaved, nil, fmt.Errorf("expected args")
	}

	return p, functioncall, nil
}

// args ::= '(' [explist] ')'
//        | tableconstructor
//        | LiteralString
func (p parser) accept_args() (parser, *ARGS, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved := p

	// accept args.rule1
	pp, rule1, err := p.accept_args_rule1()
	if err != nil {
		return pSaved, nil, err
	}
	if rule1 != nil {
		return pp, &ARGS{rule1: rule1}, nil
	}

	// accept args.rule2
	pp, rule2, err := p.accept_args_rule2()
	if err != nil {
		return pSaved, nil, err
	}
	if rule2 != nil {
		return pp, &ARGS{rule2: rule2}, nil
	}

	// accept args.rule3
	pp, rule3, err := p.accept_args_rule3()
	if err != nil {
		return pSaved, nil, err
	}
	if rule3 != nil {
		return pp, &ARGS{rule3: rule3}, nil
	}

	return pSaved, nil, nil
}

// args.rule1 ::= '(' [explist] ')'
func (p parser) accept_args_rule1() (parser, *ARGS_RULE1, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, args := p, &ARGS_RULE1{}
	var err error

	// accept '('
	var oParen []byte
	p, oParen = p.accept_Literal('(')
	if oParen == nil {
		return pSaved, nil, nil
	}

	// accept [explist]
	p, args.explist, err = p.accept_explist()
	if err != nil {
		return pSaved, nil, err
	}

	// expect ')'
	var cParen []byte
	p, cParen = p.accept_Literal(')')
	if cParen == nil {
		return pSaved, nil, fmt.Errorf("expected ')'")
	}

	return p, args, nil
}

// args.rule2 ::= tableconstructor
func (p parser) accept_args_rule2() (parser, *ARGS_RULE2, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, args := p, &ARGS_RULE2{}
	var err error

	p, args.tableconstructor, err = p.accept_tableconstructor()
	if err != nil {
		return pSaved, nil, err
	}

	return p, args, nil
}

// args.rule3 ::= LiteralString
func (p parser) accept_args_rule3() (parser, *ARGS_RULE3, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, args := p, &ARGS_RULE3{}
	var err error

	p, args.literalString, err = p.accept_LiteralString()
	if err != nil {
		return pSaved, nil, err
	}

	return p, args, nil
}

// functiondef ::= 'function' funcbody
func (p parser) accept_functiondef() (parser, *FUNCTIONDEF, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, functiondef := p, &FUNCTIONDEF{}
	var err error

	// accept 'function'
	var function []byte
	p, function = p.accept_Literal('f', 'u', 'n', 'c', 't', 'i', 'o', 'n')
	if function == nil {
		return pSaved, nil, nil
	}

	// expect funcbody
	p, functiondef.funcbody, err = p.accept_funcbody()
	if err != nil {
		return pSaved, nil, err
	}
	if functiondef.funcbody == nil {
		return pSaved, nil, fmt.Errorf("expected funcbody")
	}

	return p, functiondef, nil
}

// funcbody ::= '(' [parlist] ')' block 'end'
func (p parser) accept_funcbody() (parser, *FUNCBODY, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, funcbody := p, &FUNCBODY{}
	var err error

	// accept '('
	var oParen []byte
	p, oParen = p.accept_Literal('(')
	if oParen == nil {
		return pSaved, nil, nil
	}

	// accept [parlist]
	p, funcbody.parlist, err = p.accept_parlist()
	if err != nil {
		return pSaved, nil, err
	}

	// expect ')'
	var cParen []byte
	p, cParen = p.accept_Literal(')')
	if cParen == nil {
		return pSaved, nil, fmt.Errorf("expected ')'")
	}

	// expect block
	p, funcbody.block, err = p.accept_block()
	if err != nil {
		return pSaved, nil, err
	}
	if funcbody.block == nil {
		return pSaved, nil, fmt.Errorf("expected block")
	}

	// expect 'end'
	var end []byte
	p, end = p.accept_Literal('e', 'n', 'd')
	if end == nil {
		return pSaved, nil, fmt.Errorf("expected 'end'")
	}
	panic("!implemented")
}

// parlist ::= namelist [',' '...'] | '...'
func (p parser) accept_parlist() (parser, *PARLIST, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved := p

	// accept parlist.rule1
	pp, rule1, err := p.accept_parlist_rule1()
	if err != nil {
		return pSaved, nil, err
	}
	if rule1 != nil {
		return pp, &PARLIST{rule1: rule1}, nil
	}

	// accept parlist.rule2
	pp, rule2, err := p.accept_parlist_rule2()
	if err != nil {
		return pSaved, nil, err
	}
	if rule2 != nil {
		return pp, &PARLIST{rule2: rule2}, nil
	}

	return pSaved, nil, nil
}

// parlist.rule1 ::= namelist [',' '...']
func (p parser) accept_parlist_rule1() (parser, *PARLIST_RULE1, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, parlist := p, &PARLIST_RULE1{}
	var err error

	// accept namelist [',' '...']
	p, parlist.namelist, err = p.accept_namelist()
	if err != nil {
		return pSaved, nil, err
	}
	if parlist.namelist == nil {
		return pSaved, nil, nil
	}

	p, parlist.comma = p.accept_Literal(',')
	if parlist.comma == nil {
		return p, parlist, nil
	}

	p, parlist.dotDotDot = p.accept_Literal('.', '.', '.')
	if parlist.dotDotDot == nil {
		return pSaved, nil, fmt.Errorf("expected '...'")
	}

	return p, parlist, nil
}

// parlist.rule2 ::= '...'
func (p parser) accept_parlist_rule2() (parser, *PARLIST_RULE2, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, parlist := p, &PARLIST_RULE2{}

	// accept '...'
	p, parlist.dotDotDot = p.accept_Literal('.', '.', '.')
	if parlist.dotDotDot == nil {
		return pSaved, nil, nil
	}

	return p, parlist, nil
}

// tableconstructor ::= '{' [fieldlist] '}'
func (p parser) accept_tableconstructor() (parser, *TABLECONSTRUCTOR, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, tableconstructor := p, &TABLECONSTRUCTOR{}
	var err error

	// accept '{'
	var oBrace []byte
	p, oBrace = p.accept_Literal('{')
	if oBrace == nil {
		return pSaved, nil, nil
	}

	// accept [fieldlist]
	p, tableconstructor.fieldlist, err = p.accept_fieldlist()
	if err != nil {
		return pSaved, nil, err
	}

	// expect '}'
	var cBrace []byte
	p, cBrace = p.accept_Literal('}')
	if cBrace == nil {
		return pSaved, nil, fmt.Errorf("expected '}'")
	}

	return p, tableconstructor, nil
}

// fieldlist ::= field {fieldsep field} [fieldsep]
func (p parser) accept_fieldlist() (parser, *FIELDLIST, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, fieldlist := p, &FIELDLIST{}
	var err error

	// accept field
	var field *FIELD
	p, field, err = p.accept_field()
	if err != nil {
		return pSaved, nil, err
	}
	if field == nil {
		return pSaved, nil, nil
	}
	fieldlist.fields = append(fieldlist.fields, field)

	// accept {fieldsep field}
	for {
		pp, fieldsep, _ := p.accept_fieldsep()
		if fieldsep == nil {
			break
		}
		pp, field, _ = pp.accept_field()
		if field == nil {
			break
		}
		fieldlist.fields = append(fieldlist.fields, field)
		p = pp
	}

	// accept [fieldsep]
	if pp, fieldsep, _ := p.accept_fieldsep(); fieldsep != nil {
		p = pp
	}

	return p, fieldlist, nil
}

// field ::= '[' exp ']' '=' exp
//         | Name '=' exp
//         | exp
func (p parser) accept_field() (parser, *FIELD, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, field := p, &FIELD{}
	var err error

	// accept field.rule1
	p, field.rule1, err = p.accept_field_rule1()
	if err != nil {
		return pSaved, nil, err
	}
	if field.rule1 != nil {
		return p, field, nil
	}

	// accept field.rule2
	p, field.rule2, err = p.accept_field_rule2()
	if err != nil {
		return pSaved, nil, err
	}
	if field.rule2 != nil {
		return p, field, nil
	}

	// accept field.rule3
	p, field.rule3, err = p.accept_field_rule3()
	if err != nil {
		return pSaved, nil, err
	}
	if field.rule3 != nil {
		return pSaved, field, nil
	}

	return pSaved, nil, nil
}

// field.rule1 ::= '[' exp ']' '=' exp
func (p parser) accept_field_rule1() (parser, *FIELD_RULE1, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, field := p, &FIELD_RULE1{}

	var err error

	var oBracket []byte
	p, oBracket = p.accept_Literal('[')
	if oBracket == nil {
		return pSaved, nil, nil
	}
	p, field.exp1, err = p.accept_exp()
	if err != nil {
		return pSaved, nil, err
	}
	if field.exp1 == nil {
		return pSaved, nil, fmt.Errorf("expected expression")
	}

	var cBracket []byte
	p, cBracket = p.accept_Literal('=')
	if cBracket == nil {
		return pSaved, nil, fmt.Errorf("expected ']'")
	}

	var equals []byte
	p, equals = p.accept_Literal('=')
	if equals == nil {
		return pSaved, nil, fmt.Errorf("expected '='")
	}

	p, field.exp2, err = p.accept_exp()
	if err != nil {
		return pSaved, nil, err
	}
	if field.exp2 == nil {
		return pSaved, nil, fmt.Errorf("expected expression")
	}

	return p, field, nil
}

// field.rule2 ::= Name '=' exp
func (p parser) accept_field_rule2() (parser, *FIELD_RULE2, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, field := p, &FIELD_RULE2{}

	var err error

	// accept Name
	p, field.name, _ = p.accept_Name()
	if field.name == nil {
		return pSaved, nil, nil
	}

	// expect '='
	var equals []byte
	p, equals = p.accept_Literal('=')
	if equals == nil {
		return pSaved, nil, fmt.Errorf("expected '='")
	}

	// expect exp
	p, field.exp, err = p.accept_exp()
	if err != nil {
		return pSaved, nil, err
	}
	if field.exp == nil {
		return pSaved, nil, fmt.Errorf("expected expression")
	}
	return p, field, nil
}

// field.rule3 ::= exp
func (p parser) accept_field_rule3() (parser, *FIELD_RULE3, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSaved, field := p, &FIELD_RULE3{}
	var err error

	// accept exp
	p, field.exp, err = p.accept_exp()
	if err != nil {
		return pSaved, nil, err
	}
	if field.exp == nil {
		return pSaved, nil, err
	}

	return p, field, nil
}

// fieldsep ::= ',' | ';'
func (p parser) accept_fieldsep() (parser, *FIELDSEP, error) {
	if eof(p) {
		return p, nil, nil
	}
	switch p.buf[0] {
	case ',':
		return p.skipch(1), &FIELDSEP{comma: true}, nil
	case ';':
		return p.skipch(1), &FIELDSEP{semicolon: true}, nil
	}
	return p, nil, nil
}

// binop ::= '+' | '-'  | '*' | '/'  | '//' | '^' | '%'
//         | '&' | '~'  | '|' | '>>' | '<<' | '..'
//         | '<' | '<=' | '>' | '>=' | '==' | '~='
//         | and | or
func (p parser) accept_binop() (parser, *BINOP, error) {
	if eof(p) {
		return p, nil, nil
	}
	switch p.buf[0] {
	case '+':
		return p.skipch(1), &BINOP{plus: true}, nil
	case '-':
		return p.skipch(1), &BINOP{hyphen: true}, nil
	case '*':
		return p.skipch(1), &BINOP{asterisk: true}, nil
	case '/': // '/' or '//'
		if match := []byte{'/', '/'}; bytes.HasPrefix(p.buf, match) {
			return p.skipch(len(match)), &BINOP{slash: true}, nil
		}
		return p.skipch(1), &BINOP{slashSlash: true}, nil
	case '^':
		return p.skipch(1), &BINOP{caret: true}, nil
	case '%':
		return p.skipch(1), &BINOP{percent: true}, nil
	case '&':
		return p.skipch(1), &BINOP{ampersand: true}, nil
	case '~': // '~' | '~='
		if match := []byte{'~', '='}; bytes.HasPrefix(p.buf, match) {
			return p.skipch(len(match)), &BINOP{tildeEqual: true}, nil
		}
		return p.skipch(1), &BINOP{tilde: true}, nil
	case '|':
		return p.skipch(1), &BINOP{pipe: true}, nil
	case '.': // '..'
		if match := []byte{'.', '.'}; bytes.HasPrefix(p.buf, match) {
			return p.skipch(len(match)), &BINOP{dotDot: true}, nil
		}
	case '=': // '=='
		if match := []byte{'=', '='}; bytes.HasPrefix(p.buf, match) {
			return p.skipch(len(match)), &BINOP{equalEqual: true}, nil
		}
	case '<': // '<' | '<<' | '<='
		if match := []byte{'<', '<'}; bytes.HasPrefix(p.buf, match) {
			return p.skipch(len(match)), &BINOP{lessThanlessThan: true}, nil
		}
		if match := []byte{'<', '='}; bytes.HasPrefix(p.buf, match) {
			return p.skipch(len(match)), &BINOP{lessThanEqual: true}, nil
		}
		return p.skipch(1), &BINOP{lessThan: true}, nil
	case '>': // '>' | '>>' | '>='
		if match := []byte{'>', '>'}; bytes.HasPrefix(p.buf, match) {
			return p.skipch(len(match)), &BINOP{greaterThanGreaterThan: true}, nil
		}
		if match := []byte{'>', '='}; bytes.HasPrefix(p.buf, match) {
			return p.skipch(len(match)), &BINOP{greaterThanEqual: true}, nil
		}
		return p.skipch(1), &BINOP{greaterThan: true}, nil
	case 'a': // 'and'
		if match := []byte{'a', 'n', 'd'}; bytes.HasPrefix(p.buf, match) {
			return p.skipch(len(match)), &BINOP{and: true}, nil
		}
	case 'o': // 'or'
		if match := []byte{'o', 'r'}; bytes.HasPrefix(p.buf, match) {
			return p.skipch(len(match)), &BINOP{or: true}, nil
		}
	}
	return p, nil, nil
}

// unop ::= '-' | not | '#' | '~'
func (p parser) accept_unop() (parser, *UNOP, error) {
	if eof(p) {
		return p, nil, nil
	}
	pSave := p
	switch p.buf[0] {
	case '-':
		return p.skipch(1), &UNOP{dash: true}, nil
	case '#':
		return p.skipch(1), &UNOP{hash: true}, nil
	case 'n': // 'not'
		if match := []byte{'n', 'o', 't'}; bytes.HasPrefix(p.buf, match) {
			return p.skipch(len(match)), &UNOP{not: true}, nil
		}
	case '~':
		return p.skipch(1), &UNOP{tilde: true}, nil
	}
	return pSave, nil, nil
}
