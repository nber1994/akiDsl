package compile

import (
    "go/ast"
    "go/token"
    "github.com/nber1994/akiDsl/runCxt"
    "github.com/nber1994/akiDsl/dslCxt"
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

func (this *CompileCxt) Rescue() {
}
