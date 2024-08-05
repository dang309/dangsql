package backend

type Ast struct {
	Statements []*Statement
}

type AstKind uint

const (
	SelectKind AstKind = iota
	CreateTableKind
	InsertKind
)

type Statement struct {
	SelectStatement      *SelectStatement
	CreateTableStatement *CreateTableStatement
	InsertStatement      *InsertStatement
	Kind                 AstKind
}

type InsertStatement struct {
	table  Token
	values *[]*expression
}

type expressionKind uint

const (
	literalKind expressionKind = iota
)

type expression struct {
	kind    expressionKind
	literal *Token
}

type columnDefinition struct {
	name     Token
	datatype Token
}

type CreateTableStatement struct {
	name Token
	cols *[]*columnDefinition
}

type SelectStatement struct {
	item []*expression
	from Token
}
