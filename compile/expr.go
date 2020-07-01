package compile

import (
    "github.com/nber1994/akiDsl/dslCxt"
    "go/token"
    "go/ast"
    "github.com/spf13/cast"
    "reflect"
    "fmt"
	"strconv"
)

type Expr struct {

}

var (
    SupFuncList = map[string]string{
        "append":"Append",
        "len":"Len",
        "timeParse":"TimeParse",
        "println":"Println",
        "sprintf":"Sprintf",
    }
)

func NewExpr() *Expr {
    return &Expr{}
}

func (this *Expr) CompileExpr(cpt *CompileCxt, rct *Stmt, r ast.Expr) interface{} {
    var ret interface{}
	if nil == r {
		return ret
	}
    switch r := r.(type) {
    case *ast.BasicLit: //基本类型
        ret = this.CompileBasicLitExpr(cpt, rct, r)
    case *ast.BinaryExpr: //二元表达式
        ret = this.CompileBinaryExpr(cpt, rct, r)
    case *ast.CompositeLit: //集合类型
        switch  r.Type.(type) {
        case *ast.ArrayType: //数组
            ret = this.CompileArrayExpr(cpt, rct, r)
        case *ast.MapType: //map
            ret = this.CompileMapExpr(cpt, rct, r)
        default:
            panic(fmt.Sprintf("syntax error: Bad CompositeList Type %v", cpt.Fset.Position(r.Pos())))
        }
    case *ast.CallExpr:
        ret = this.CompileCallExpr(cpt, rct, r)
    case *ast.Ident:
        ret = this.CompileIdentExpr(cpt, rct, r)
    case *ast.IndexExpr:
        ret = this.CompileIndexExpr(cpt, rct, r)
    case *ast.SliceExpr:
        ret = this.CompileSliceExpr(cpt, rct, r)
    default:
        panic(fmt.Sprintf("syntax error: Bad Expr Type %v", cpt.Fset.Position(r.Pos())))
    }
    return ret
}

func (this *Expr) CompileSliceExpr(cpt *CompileCxt, rct *Stmt, r *ast.SliceExpr) interface{} {
    //fmt.Println("------------------------in Slice expr")
    var ret interface{}
    x := this.CompileExpr(cpt, rct, r.X)
    low := this.CompileExpr(cpt, rct, r.Low)
    high := this.CompileExpr(cpt, rct, r.High)
    switch x := x.(type) {
    case []interface{}:
        if nil != low && nil != high {
            ret = x[cast.ToInt(low):cast.ToInt(high)]
        } else if nil == low && nil != high {
            ret = x[:cast.ToInt(high)]
        } else if nil != low && nil == high {
            ret = x[cast.ToInt(low):]
        } else if nil == low && nil == high {
            ret = x[:]
        }
    default:
        panic(fmt.Sprintf("syntax error: Bad SliceExpr Type %v", cpt.Fset.Position(r.Pos())))
    }
    return ret
}

//index操作
func (this *Expr) CompileIndexExpr(cpt *CompileCxt, rct *Stmt, r *ast.IndexExpr) interface{} {
    //fmt.Println("------------------------in Index expr")
    var ret interface{}
    target := this.CompileExpr(cpt, rct, r.X)
    index := this.CompileExpr(cpt, rct, r.Index)
    switch target := target.(type) {
    case []interface{}:
        idx := cast.ToInt(index)
        if idx >= len(target) {
            panic(fmt.Sprintf("syntax error: Index Out Of Range %v", cpt.Fset.Position(r.Pos())))
        }
        ret = target[idx]
    case map[interface{}]interface{}:
        if _, exist := target[index]; !exist {
            panic(fmt.Sprintf("syntax error: Key Not Exist %v", cpt.Fset.Position(r.Pos())))
        }
        ret = target[index]
    default:
        panic(fmt.Sprintf("syntax error: Bad IndexExpr Type %v", cpt.Fset.Position(r.Pos())))
    }
    return ret
}

//内置函数 MethodByName会panic
func (this *Expr) CompileCallExpr(cpt *CompileCxt, rct *Stmt, r *ast.CallExpr) interface{} {
    //fmt.Println("------------------------in Call expr")
    var ret interface{}
    //校验内置函数
    var funcArgs []reflect.Value
    funcName := r.Fun.(*ast.Ident).Name
    //fmt.Println("------------------------in Call expr ", funcName)
    //初始化入参
    for _, arg := range r.Args {
        funcArgs = append(funcArgs, reflect.ValueOf(this.CompileExpr(cpt, rct, arg)))
    }
    //fmt.Println("------------------------in Call expr args", funcArgs)
    var res []reflect.Value
    if RealFuncName, exist := SupFuncList[funcName]; exist {
        flib := NewFuncLib()
        res = reflect.ValueOf(flib).MethodByName(RealFuncName).Call(funcArgs)
    } else if CxtFuncName, cxtExist := dslCxt.SupFuncList[funcName]; cxtExist {
        res = reflect.ValueOf(cpt.DslCxt).MethodByName(CxtFuncName).Call(funcArgs)
    } else {
        panic(fmt.Sprintf("syntax error: Bad Func Name %v", cpt.Fset.Position(r.Pos())))
    }
    if nil == res {
        return ret
    }
    return res[0].Interface()
}

//处理多返回值函数
func (this *Expr) CompileCallMultiReturnExpr(cpt *CompileCxt, rct *Stmt, r *ast.CallExpr) []interface{} {
    //fmt.Println("------------------------in Call multi expr")
    funcLib := NewFuncLib()
    var ret []interface{}
    //校验内置函数
    var funcArgs []reflect.Value
    funcName := r.Fun.(*ast.Ident).Name
    //初始化入参
    for _, arg := range r.Args {
        funcArgs = append(funcArgs, reflect.ValueOf(this.CompileExpr(cpt, rct, arg)))
    }
    res := reflect.ValueOf(funcLib).MethodByName(funcName).Call(funcArgs)
    for _, v := range res {
        ret = append(ret, v.Interface())
    }
    return ret
}

func (this *Expr) CompileBasicLitExpr(cpt *CompileCxt, rct *Stmt, r *ast.BasicLit) interface{} {
    //fmt.Println("------------------------in basiclit expr")
    var ret interface{}
    switch r.Kind {
    case token.INT:
        ret = cast.ToInt64(r.Value)
    case token.FLOAT:
        ret = cast.ToFloat64(r.Value)
    case token.STRING:
        retStr := cast.ToString(r.Value)
		var err error
		//去掉转义的双引号 这个真tm天坑
		ret, err = strconv.Unquote(retStr)
		if nil != err {
            panic(fmt.Sprintf("syntax error: Bad String %v", cpt.Fset.Position(r.Pos())))
		}
    default:
        panic(fmt.Sprintf("syntax error: Bad BasicList Type %v", cpt.Fset.Position(r.Pos())))
    }
    //fmt.Println("------------------------expr res ", ret)
    return ret
}

func (this *Expr) CompileArrayExpr(cpt *CompileCxt, rct *Stmt, r *ast.CompositeLit) interface{} {
    //fmt.Println("------------------------in array expr")
    var ret []interface{}
    for _, e := range r.Elts {
        switch e := e.(type) {
        case *ast.BasicLit:
            ret = append(ret, this.CompileExpr(cpt, rct, e))
        case *ast.CompositeLit:
            //拼接结构体
            compLit := &ast.CompositeLit{
                Type: r.Type.(*ast.ArrayType).Elt,
                Elts: e.Elts,
            }
            ret = append(ret, this.CompileExpr(cpt, rct, compLit))
        default:
            panic(fmt.Sprintf("syntax error: Bad Array Item Type %v", cpt.Fset.Position(r.Pos())))
        }
    }
    return ret
}

func (this *Expr) CompileMapExpr(cpt *CompileCxt, rct *Stmt, r *ast.CompositeLit) interface{} {
    //fmt.Println("------------------------in map expr")
    ret := make(map[interface{}]interface{})
    var key interface{}
    var value interface{}
    for _, e := range r.Elts {
        key = this.CompileExpr(cpt, rct, e.(*ast.KeyValueExpr).Key)
        value = this.CompileExpr(cpt, rct, e.(*ast.KeyValueExpr).Value)
        ret[key] = value
    }
    return ret
}


func (this *Expr) CompileIdentExpr(cpt *CompileCxt, rct *Stmt, r *ast.Ident) interface{} {
    //fmt.Println("------------------------in ident expr")
    var ret interface{}
    ret = rct.GetValue(r.Name)
    return ret
}

func (this *Expr) CompileBinaryExpr(cpt *CompileCxt, rct *Stmt, r *ast.BinaryExpr) interface{} {
    //fmt.Println("------------------------in binary expr")
    var ret interface{}
    switch r.Op {
        //+ - * / %
    case token.ADD:
        ret = BAdd(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.SUB:
        ret = BSub(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.MUL:
        ret = BMul(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.QUO:
        ret = BQuo(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.REM:
        ret = BRem(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
        // &&, ||, &, |, >, <, >=, <=, ==, !=
    case token.AND:
        ret = BAnd(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.OR:
        ret = BOr(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.LAND:
        ret = BLand(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.LOR:
        ret = BLor(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.GTR:
        ret = BGtr(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.LSS:
        ret = BLss(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.GEQ:
        ret = BGeq(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.LEQ:
        ret = BLeq(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.NEQ:
        ret = BNeq(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    case token.EQL:
        ret = BEql(this.CompileExpr(cpt, rct, r.X), this.CompileExpr(cpt, rct, r.Y))
    default:
        panic(fmt.Sprintf("syntax error: Bad BinaryExpr Type %v", cpt.Fset.Position(r.Pos())))
    }
    return ret
}

