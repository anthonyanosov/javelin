package ast

type Program []Stmt

type Stmt interface {
	stmtNode()
}

type VarStmt struct {
	Ident   string
	Integer int64
}

func (v *VarStmt) stmtNode() {}
