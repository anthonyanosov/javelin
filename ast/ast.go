package ast

type Program []Stmt

type Stmt interface {
	stmtNode()
}

type DeclStmt struct {
	Ident   string
	Integer int64
}

func (v *DeclStmt) stmtNode() {}
