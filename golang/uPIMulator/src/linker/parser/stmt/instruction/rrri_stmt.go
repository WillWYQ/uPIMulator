package instruction

import (
	"errors"
	"uPIMulator/src/linker/parser/expr"
)

type RrriStmt struct {
	op_code *expr.Expr
	rc      *expr.Expr
	ra      *expr.Expr
	rb      *expr.Expr
	imm     *expr.Expr
}

func (this *RrriStmt) Init(
	op_code *expr.Expr,
	rc *expr.Expr,
	ra *expr.Expr,
	rb *expr.Expr,
	imm *expr.Expr,
) {
	if op_code.ExprType() != expr.RRRI_OP_CODE {
		err := errors.New("op code is not an RRRI op code")
		panic(err)
	}

	if rc.ExprType() != expr.SRC_REG {
		err := errors.New("rc is not a src reg")
		panic(err)
	}

	if ra.ExprType() != expr.SRC_REG {
		err := errors.New("ra is not a src reg")
		panic(err)
	}

	if rb.ExprType() != expr.SRC_REG {
		err := errors.New("rb is not a src reg")
		panic(err)
	}

	if imm.ExprType() != expr.PROGRAM_COUNTER {
		err := errors.New("imm is not a progrcm counter")
		panic(err)
	}

	this.op_code = op_code
	this.rc = rc
	this.ra = ra
	this.rb = rb
	this.imm = imm
}

func (this *RrriStmt) OpCode() *expr.Expr {
	return this.op_code
}

func (this *RrriStmt) Rc() *expr.Expr {
	return this.rc
}

func (this *RrriStmt) Ra() *expr.Expr {
	return this.ra
}

func (this *RrriStmt) Rb() *expr.Expr {
	return this.rb
}

func (this *RrriStmt) Imm() *expr.Expr {
	return this.imm
}
