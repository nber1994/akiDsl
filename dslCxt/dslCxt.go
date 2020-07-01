package dslCxt

import (
   "github.com/nber1994/akiDsl/nodejson"
   //"fmt"
)

type DslCxt struct {
    OriginCxt *string //原始上下文
	Node *nodejson.Node
}

var (
    SupFuncList = map[string]string{
        "GetInt":"GetInt",
        "GetFloat":"GetFloat",
        "GetBool":"GetBool",
        "GetString":"GetString",
        "Set":"Set",
        "Exist":"Exist",
    }
)


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

func (this *DslCxt) ToJsonString() string {
    ret, _ := this.Node.ToJsonString()
    return ret
}

//获取Cxt的值
func (this *DslCxt) GetInt(path string) int64 {
    value := this.Node.Get(path)
    //fmt.Println("....dsl Get path ", path, " value ", value.Value())
    return value.Int64()
}

//获取Cxt的值
func (this *DslCxt) GetFloat(path string) float64 {
    value := this.Node.Get(path)
    //fmt.Println("....dsl Get path ", path, " value ", value.Value())
    return value.Float64()
}

//获取Cxt的值
func (this *DslCxt) GetBool(path string) bool {
    value := this.Node.Get(path)
    //fmt.Println("....dsl Get path ", path, " value ", value.Value())
    return value.Bool()
}

//获取Cxt的值
func (this *DslCxt) GetString(path string) string {
    value := this.Node.Get(path)
    //fmt.Println("....dsl Get path ", path, " value ", value.Value())
    return value.String()
}

func (this *DslCxt) Set(path string, value interface{}) interface{} {
    this.Node.Set(path, value)
    //fmt.Println("....dsl Set path ", path, " value ", value)
    return this.GetString(path)
}

func (this *DslCxt) Exist(path string) bool {
    return this.Node.Exist(path)
}
