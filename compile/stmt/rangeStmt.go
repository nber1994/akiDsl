package stmt

import (
    "github.com/nber1994/akiDsl/runCxt"
    "github.com/nber1994/akiDsl/compile/stmt/expr"
    "github.com/nber1994/akiDsl/compile"
    "go/token"
)

type RangeStmt struct {
    Key *ast.Ident
    Body *ast.BlockStmt
}

func NewRangeStmt(stmt *ast.RangeStmt) *RangeStmt {
    return &RangeStmt {
        Key: stmt.Key,
        Body: stmt.Body,
    }
}

func (this *RangeStmt) Compile(cpt *compile.CompileCxt, stmt *ast.Stmt) {
    expr := expr.NewExpr()
    stmtHd := stmt.NewStmt()
    RangeTarget := expr.CompileExpr(cpt.dct, cpt.rct, stmt.Obj.Decl[0].Rhs[0].X)
    kName := this.Key.Name
    vName := this.Key.Obj.Decl[1].Name
    for k, v := range RangeTarget {
        //设置kv的值
        cpt.rct.SetValue(kName, k)
        cpt.rct.SetValue(vName, v)
        //执行Body
        stmtHd.CompileStmt(this.Body)
    }
}

