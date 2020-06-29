package dslCxt

import (
   "github.com/nber1994/akiDsl/nodejson"
   "fmt"
)

type DslCxt struct {
    OriginCxt string //原始上下文
}


func New(originCxt string) *DslCxt {
    fmt.Println("....originCxt ", originCxt)
    return &DslCxt{
        OriginCxt: originCxt,
    }
}

//获取Cxt的值
func (this *DslCxt) Get(path string) interface{} {
    fmt.Println("....dsl ", this.OriginCxt)
    Node, e := nodejson.UnmarshalToNode([]byte(this.OriginCxt))
    fmt.Println(e)
    value := Node.Get(path)
    fmt.Println("....dsl ", Node)
    fmt.Println("....dsl Get path ", path, " value ", value.Value())
    return value.Value()
}

func (this *DslCxt) Set(path string, value interface{}) interface{} {
    Node, _ := nodejson.UnmarshalToNode([]byte(this.OriginCxt))
    Node.Set(path, value)
    fmt.Println("....dsl Set path ", path, " value ", value)
    return this.Get(path)
}
