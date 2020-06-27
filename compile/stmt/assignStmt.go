package stmt

import (
    "github.com/nber1994/akiDsl/runCxt"
    "github.com/nber1994/akiDsl/compile/stmt/expr"
    "github.com/nber1994/akiDsl/compile"
    "go/token"
)

type AssignStmt struct {
    Lhs []ast.Expr
    Tok token.Token //操作符
    Rhs []sat.Expr
}

func NewAssignStmt(stmt *ast.AssignStmt) *AssignStmt {
    return &AssignStmt{
        Lhs: stmt.Lhs,
        Tok: stmt.Tok,
        Rhs: stmt.Rhs,
    }
}

func (this *AssignStmt) Compile(cpt *compile.CompileCxt, stmt *ast.Stmt) {
    //只支持= :=
    if token.DEFINE != this.Tok || token.ASSIGN != this.Tok {
        panic("syntax error: nonsupport Tok ", this.Tok)
    }

    expr := NewExpr()

    //Lhs中的变量进行声明
    for idx, l := range this.Lhs {
        switch l := l.(type) {
        case *ast.Ident:
            r := this.Rhs[idx]
            cpt.RunCxt.Vars[l.Name] = expr.CompileExpr(cpt.RunCxt, r)
        default:
            panic("syntax error: assign type must be ident type")
        }
    }
}
