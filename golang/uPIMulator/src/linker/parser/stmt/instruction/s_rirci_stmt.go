package instruction

import (
	"errors"
	"uPIMulator/src/linker/lexer"
	"uPIMulator/src/linker/parser/expr"
)

type SRirciStmt struct {
	op_code   *expr.Expr
	suffix    *expr.Expr
	dc        *lexer.Token
	imm       *expr.Expr
	ra        *expr.Expr
	condition *expr.Expr
	pc        *expr.Expr
}

func (this *SRirciStmt) Init(
	op_code *expr.Expr,
	suffix *expr.Expr,
	dc *lexer.Token,
	imm *expr.Expr,
	ra *expr.Expr,
	condition *expr.Expr,
	pc *expr.Expr,
) {
	if op_code.ExprType() != expr.RRI_OP_CODE {
		err := errors.New("op code is not an RRI op code")
		panic(err)
	}

	if suffix.ExprType() != expr.SUFFIX {
		err := errors.New("suffix is not a suffix")
		panic(err)
	}

	if dc.TokenType() != lexer.PAIR_REG {
		err := errors.New("dc is not a pair reg")
		panic(err)
	}

	if imm.ExprType() != expr.PROGRAM_COUNTER {
		err := errors.New("imm is not a program counter")
		panic(err)
	}

	if ra.ExprType() != expr.SRC_REG {
		err := errors.New("ra is not a src reg")
		panic(err)
	}

	if condition.ExprType() != expr.CONDITION {
		err := errors.New("condition is not a condition")
		panic(err)
	}

	if pc.ExprType() != expr.PROGRAM_COUNTER {
		err := errors.New("pc is not a program counter")
		panic(err)
	}

	this.op_code = op_code
	this.suffix = suffix
	this.dc = dc
	this.imm = imm
	this.ra = ra
	this.condition = condition
	this.pc = pc
}

func (this *SRirciStmt) OpCode() *expr.Expr {
	return this.op_code
}

func (this *SRirciStmt) Suffix() *expr.Expr {
	return this.suffix
}

func (this *SRirciStmt) Dc() *lexer.Token {
	return this.dc
}

func (this *SRirciStmt) Imm() *expr.Expr {
	return this.imm
}

func (this *SRirciStmt) Ra() *expr.Expr {
	return this.ra
}

func (this *SRirciStmt) Condition() *expr.Expr {
	return this.condition
}

func (this *SRirciStmt) Pc() *expr.Expr {
	return this.pc
}
