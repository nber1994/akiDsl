package stmt

import (
    "akiDsl/runCxt"
    "akiDsl/compile/stmt/expr"
    "go/token"
)

type IncDecStmt struct {
    X *ast.Ident
    Tok token.Token
}

func NewIncDecStmt(stmt *ast.IncDecStmt) *IncDecStmt {
    return &IncDecStmt{
        X: stmt.X,
        Tok: stmt.Tok,
    }
}

func (this *IncDecStmt) Compile(cpt *compile.CompileCxt, stmt *ast.Stmt) {
    //只支持 ++ --
    if token.INC != this.Tok || token.DEC != this.Tok {
        panic("syntax error: nonsupport Tok ", this.Tok)
    }

    expr := NewExpr()
    switch this.Tok {
    case token.INC:
        cpt.RunCxt.Vars[this.X.Name] = expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, this.X)
        cpt.RunCxt.Vars[this.X.Name]++
    case token.DEC:
        cpt.RunCxt.Vars[this.X.Name] = expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, this.X)
        cpt.RunCxt.Vars[this.X.Name]--
    default:
        panic("syntax error: nonsupport Tok ", this.Tok)
    }
}
