package akiDsl

import (
    "go/parser"
    "go/token"
    "github.com/nber1994/akiDsl/dslCxt"
    "github.com/nber1994/akiDsl/compile"
)

type AkiDsl struct {
    FileName *string //dsl脚本地址
    DslCxt *dslCxt.DslCxt//dsl与上下文的交互
}

func New(fileName *string, Cxt *string) (*AkiDsl, error) {
    dslCxtNode, err := dslCxt.New(*Cxt)
    if nil != err {
        return nil, err
    }
    return &AkiDsl{
        FileName: fileName,
        DslCxt: dslCxtNode,
    }, nil
}

func (this *AkiDsl) Run() (interface{}, *dslCxt.DslCxt, error){
    //总体控制错误信息
    var retErr error
    var ret interface{}

    fset := token.NewFileSet()
    //这块可以扩展不止传入文件名
    fAst, err := parser.ParseFile(fset, *this.FileName, nil, 0)
    if err != nil {
        panic(err)
    }

    pct := compile.New(fAst, fset, this.DslCxt)
    decl := compile.NewDecl()
    d := pct.FAst.Decls[0]

    go func() {
        decl.CompileDecl(pct, d)
    }()

    select {
    case ret = <-pct.ReturnCh:
    case retErr = <-pct.ErrCh:
    }
    //定义一个空方法，不然gc会在CompileDecl方法结束以后，把pct结构体回收 pct.Rescue()
    return ret, pct.DslCxt, retErr
}
