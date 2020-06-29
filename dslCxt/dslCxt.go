package dslCxt

import (
   "github.com/nber1994/akiDsl/nodejson"
   "fmt"
)

type DslCxt struct {
    OriginCxt *string //原始上下文
	Node *nodejson.Node
}


func New(originCxt *string) (*DslCxt, error) {
    node, err := nodejson.UnmarshalToNode([]byte(*originCxt))
	if nil != err {
		return nil, err
	}
    return &DslCxt{
        OriginCxt: originCxt,
		Node: &node,
    }, nil
}

//获取Cxt的值
func (this *DslCxt) Get(path string) interface{} {
    value := this.Node.Get(path)
    fmt.Println("....dsl Get path ", path, " value ", value.Value())
    return value.Value()
}

func (this *DslCxt) Set(path string, value interface{}) interface{} {
    this.Node.Set(path, value)
    fmt.Println("....dsl Set path ", path, " value ", value)
    return this.Get(path)
}
