package dslCxt

import (
   "github.com/nber1994/akiDsl/nodejson"
   "fmt"
)

type DslCxt struct {
    OriginCxt *string //原始上下文
}


func New(originCxt *string) *DslCxt {
    return &DslCxt{
        OriginCxt: originCxt,
    }
}

//获取Cxt的值
func (this *DslCxt) Get(path string) interface{} {
    fmt.Println("....dsl orignCxt", *this.OriginCxt)
    node, _ := nodejson.UnmarshalToNode([]byte(*this.OriginCxt))
    value := node.Get(path)
    fmt.Println("....dsl ", value)
    fmt.Println("....dsl Get path ", path, " value ", value.Value())
    return value.Value()
}

func (this *DslCxt) Set(path string, value interface{}) interface{} {
    node, _ := nodejson.UnmarshalToNode([]byte(*this.OriginCxt))
    node.Set(path, value)
    fmt.Println("....dsl Set path ", path, " value ", value)
    return this.Get(path)
}
