package stmt

import (
    "github.com/nber1994/akiDsl/runCxt"
    "github.com/nber1994/akiDsl/compile/stmt/expr"
    "go/token"
)

type IfStmt struct {
   Cond *ast.Expr //表达式 
   Body *ast.BlockStmt //{}内表达式
   Else *ast.Stmt
}

func NewIfStmt(stmt *ast.IfStmt) *IfStmt {
    return &IfStmt{
        Cond: stmt.Cond,
        Body: stmt.Body,
        Else: stmt.Else,
    }
}

func (this *IfStmt) Compile(cpt *compile.CompileCxt, stmt *ast.Stmt) {
    expr := NewExpr()
    condRet := expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, this.Cond)
    stmt := NewStmt()
    //如果条件成立
    if condRet {
        stmt.CompileStmt(cpt, this.Body)
    } else {
        if nil == this.Else {
            return
        }
        stmt.CompileStmt(cpt, this.Else)
    }
}
