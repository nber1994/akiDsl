package akiDsl

import (
    "fmt"
    "go/parser"
    "go/token"
    "github.com/nber1994/akiDsl/dslCxt"
    "github.com/nber1994/akiDsl/compile"
)

type AkiDsl struct {
    FileName *string //dsl脚本地址
    DslCxt *dslCxt.DslCxt//dsl与上下文的交互
}

func New(fileName *string, Cxt *string) *AkiDsl{
    dslCxtNode := dslCxt.New(Cxt)
    return &AkiDsl{
        FileName: fileName,
        DslCxt: dslCxtNode,
    }
}

func (this *AkiDsl) Run() (interface{}, *dslCxt.DslCxt, error){
    //总体控制错误信息
    var retErr error
    defer func() {
        if r := recover(); r != nil {
            retErr = fmt.Errorf("internal error: %v", r)
        }
    }()

    fset := token.NewFileSet()
    //这块可以扩展不止传入文件名
    fAst, err := parser.ParseFile(fset, *this.FileName, nil, 0)
    if err != nil {
        panic(err)
    }

    pct := compile.New(fAst, fset, this.DslCxt)

    var ret interface{}

    decl := compile.NewDecl()
    d := pct.FAst.Decls[0]

    go func() {
        decl.CompileDecl(pct, d)
    }()

    ret = <-pct.ReturnCh
    close(pct.ReturnCh)
    return ret, pct.DslCxt, retErr
}
