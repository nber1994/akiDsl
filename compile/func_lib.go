package compile

type FuncLib struct {

}

func NewFuncLib() *FuncLib {
    return &FuncLib{}
}

func (this *FuncLib) Append(target []interface{}, item interface{}) interface{} {
    return append(target, item)
}

func (this *FuncLib) Len(target []interface{}) int {
    return len(target)
}
