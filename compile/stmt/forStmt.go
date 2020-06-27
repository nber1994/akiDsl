package stmt

import (
    "akiDsl/compile/stmt/expr"
    "akiDsl/compile"
)

type ForStmt struct {
    Init *ast.Stmt
    Cond *ast.Expr
    Post *ast.Stmt
    Body *ast.stmt
}

func NewForStmt(stmt *ast.ForStmt) *ForStmt {
    return &ForStmt{
        Init: stmt.Init,
        Cond: stmt.Cond,
        Post: stmt.Post,
        Body: stmt.Body,
    }
}

func (this *AssignStmt) Compile(cpt *compile.CompileCxt, stmt *ast.Stmt) {
    stmtHd := NewStmt()
    expr := NewExpr()
    //编译初始条件
    stmt.CompileStmt(cpt, this.Init)
    for {
        if access := expr.CompileExpr(cpt.dct, cpt.rct, this.Cond); !access {
            break;
        }
        //执行body
        stmt.CompileStmt(this.Body)
        stmt.CompileStmt(this.Post)
    }
}
