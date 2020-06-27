package compile

import (
    "go/ast"
    "go/token"
    "github.com/nber1994/akiDsl/runCxt"
    "github.com/nber1994/akiDsl/dslCxt"
    "github.com/nber1994/akiDsl/compile/stmt"
)

type CompileCxt struct {
    RunCxt *runCxt.RunCxt //运行时内存空间
    FAst *ast.File
    Fset *token.FileSet
    DslCxt *dslCxt.DslCxt
    ReturnCh chan interface{}
}

func New(fAst *ast.File, fset *token.FileSet, dslCxtNode *dslCxt.DslCxt) *CompileCxt {
    rct := runCxt.NewRunCxt()
    retChan := make(chan interface{})
    return &CompileCxt {
        RunCxt: rct,
        FAst: fAst,
        Fset: fset,
        DslCxt: dslCxtNode,
        ReturnCh: retChan,
    }
}

func (this *CompileCxt) Run() {
    //解析decls 目前只支持main函数
    d := this.FAst.Decls[0]
    switch d := d.(type) {
    case *ast.FuncDecl:
        //处理main函数
        if d.Name.String() == "main" {
            this.CompolieMainFuncDecl(d)
        } else {
            panic("syntax error: The entry point must be main function")
        }
    default:
        panic("syntax error: The entry point must be main function")
    }
    var ret interface{}
    this.ReturnCh <- ret
}

//解释执行函数声明 目前只支持main函数，所以默认执行编译
func (this *CompileCxt) CompolieMainFuncDecl(d *ast.FuncDecl) {
    for _, stmt := range d.Body.List {
        this.CompileStmt(&stmt)
    }
}

//编译各个stmt
func (this *CompileCxt) CompileStmt(stmt *ast.Stmt) {
    var stmtHdl stmt.Stmt
    switch stmt := stmt.(type) {
    case *ast.AssignStmt:
        stmtHdl = NewAssignStmt(stmt)
    case *ast.IncDecStmt:
        stmtHdl = NewIncDecStmt(stmt)
    case *ast.IfStmt:
        stmtHdl = NewIfStmt(stmt)
    default:
        panic("syntax error: nonsupport stmt ", stmt.Pos(), stmt.End())
    }
    stmtHdl.Compile(this.DslCxt, this.RunCxt)
}


