package dslCxt

import (
   "akiDsl/nodejson"
)

type DslCxt struct {
    OriginCxt *string //原始上下文
    Node nodejson.Node //动态json节点
}


func New(originCxt *string) *DslCxt {
    node, err := nodejson.UnmarshalToNode([]byte(*originCxt))
    if err != nil {
        panic("dslCxt UnmarshalToNode err ")
    }

    return &DslCxt{
        OriginCxt: originCxt,
        Node: node,
    }
}

//获取Cxt的值
func (this *DslCxt) Get(path string) interface{} {
    value := this.Node.Get(path)
    return value
}

func (this *DslCxt) Set(path string, value interface{}) interface{} {
    this.Node.Set(path, value)
    return this.Get(path)
}
