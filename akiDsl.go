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

func New(fileName *string, cxt *dslCxt.DslCxt) *AkiDsl {
    return &AkiDsl{
        FileName: fileName,
        DslCxt: cxt,
    }
}

func NewCxt(cxt string) (*dslCxt.DslCxt, error) {
    return dslCxt.New(&cxt)
}

func (this *AkiDsl) Run() (interface{}, *dslCxt.DslCxt, error){
    fset := token.NewFileSet()
    //这块可以扩展不止传入文件名
    fAst, err := parser.ParseFile(fset, *this.FileName, nil, 0)
    if err != nil {
        panic(err)
    }

    pct := compile.New(fAst, fset, this.DslCxt)
    decl := compile.NewDecl()
    d := pct.FAst.Decls[0]
    decl.CompileDecl(pct, d)

    //定义一个空方法，不然gc会在CompileDecl方法结束以后，把pct结构体回收 pct.Rescue()
    return pct.Return, pct.DslCxt, pct.Err
}
