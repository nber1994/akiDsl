package compile

import (
    "fmt"
    "errors"
    "go/ast"
)

type Decl struct {
}

func NewDecl() *Decl {
    return &Decl{}
}

func (this *Decl) CompileDecl(pct *CompileCxt, d ast.Decl) {
    defer func() {
        if err := recover(); err != nil {
            retErr := errors.New(err.(string))
            pct.ErrCh <- retErr
        }
    }()

    switch d := d.(type) {
    case *ast.FuncDecl:
        //处理main函数
        //解析decls 目前只支持main函数
        if d.Name.String() == "main" {
            this.CompileFuncDecl(pct, d)
        } else {
            panic("syntax error: The entry point must be main function")
        }
    default:
        panic("syntax error: The entry point must be main function")
    }
    var ret interface{}
    pct.ReturnCh <- ret
}

func (this *Decl) CompileFuncDecl(pct *CompileCxt, d *ast.FuncDecl) {
    fmt.Println("---------in func decl")
    stmtHd := NewStmt()
    for _, stmt := range d.Body.List {
        stmtHd.CompileStmt(pct, stmt)
    }
}
