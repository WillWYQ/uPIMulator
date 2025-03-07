package instruction

import (
	"errors"
	"uPIMulator/src/linker/lexer"
	"uPIMulator/src/linker/parser/expr"
)

type SRircStmt struct {
	op_code   *expr.Expr
	suffix    *expr.Expr
	dc        *lexer.Token
	imm       *expr.Expr
	ra        *expr.Expr
	condition *expr.Expr
}

func (this *SRircStmt) Init(
	op_code *expr.Expr,
	suffix *expr.Expr,
	dc *lexer.Token,
	imm *expr.Expr,
	ra *expr.Expr,
	condition *expr.Expr,
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

	this.op_code = op_code
	this.suffix = suffix
	this.dc = dc
	this.imm = imm
	this.ra = ra
	this.condition = condition
}

func (this *SRircStmt) OpCode() *expr.Expr {
	return this.op_code
}

func (this *SRircStmt) Suffix() *expr.Expr {
	return this.suffix
}

func (this *SRircStmt) Dc() *lexer.Token {
	return this.dc
}

func (this *SRircStmt) Imm() *expr.Expr {
	return this.imm
}

func (this *SRircStmt) Ra() *expr.Expr {
	return this.ra
}

func (this *SRircStmt) Condition() *expr.Expr {
	return this.condition
}
