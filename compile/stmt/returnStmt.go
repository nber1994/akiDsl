package stmt

import (
    "github.com/nber1994/akiDsl/runCxt"
    "github.com/nber1994/akiDsl/compile/stmt/expr"
    "github.com/nber1994/akiDsl/compile"
    "go/token"
)

type ReturnStmt struct {
    Results []ast.Expr
}

func NewReturnStmt(stmt *ast.ReturnStmt) *ReturnStmt {
    return &ReturnStmt{
        Results: stmt.Results,
    }
}

func (this *ReturnStmt) Compile(cpt *compile.CompileCxt, stmt *ast.Stmt) {
    var ret []interface{}
    expr := expr.NewExpr()
    for _, e := range this.Results {
        ret = append(ret, expr.CompileExpr(cpt.dct, cpt.rct, e))
    }
    cpt.ReturnCh <- ret
}
