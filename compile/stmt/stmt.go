package stmt

type Stmt struct {

}

func NewStmt() *Stmt{
    return &Stmt{}
}

//编译stmt
func (this *Stmt) CompileStmt(cpt *compile.CompileCxt, stmt *ast.Stmt) {
    if nil != stmt {
        return
    }
    var stmtHdl Stmt
    switch stmt := stmt.(type) {
    case *ast.AssignStmt:
        stmtHdl = NewAssignStmt(stmt)
    case *ast.IncDecStmt:
        stmtHdl = NewIncDecStmt(stmt)
    case *ast.IfStmt:
        stmtHdl = NewIfStmt(stmt)
    case *ast.ReturnStmt:
        stmtHdl = NewReturnStmt(stmt)
    default:
        panic("syntax error: nonsupport stmt ", stmt.Pos(), stmt.End())
    }
    stmtHdl.Compile(cpt, stmt)
}
