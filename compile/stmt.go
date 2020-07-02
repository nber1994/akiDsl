package compile

import (
    "fmt"
    "reflect"
    "go/ast"
    "go/token"
	"github.com/spf13/cast"
    "github.com/nber1994/akiDsl/runCxt"
)

type Stmt struct{
    Rct *runCxt.RunCxt //变量作用空间
    Type int
    Father *Stmt //子节点可以访问到父节点的内存空间
    Fset *ast.File
}

func NewStmt() *Stmt {
    rct := runCxt.NewRunCxt()
    return &Stmt{
        Rct: rct,
    }
}

func (this *Stmt) NewChild() *Stmt {
    stmt := NewStmt()
    stmt.Father = this
    return stmt
}

//编译stmt
func (this *Stmt) CompileStmt(cpt *CompileCxt, stmt ast.Stmt) {
    if nil == stmt {
        return
    }
    cStmt := this.NewChild()
    switch stmt := stmt.(type) {
    case *ast.AssignStmt:
        //赋值在本节点的内存中
        this.CompileAssignStmt(cpt, stmt)
    case *ast.IncDecStmt:
        this.CompileIncDecStmt(cpt, stmt)
    case *ast.IfStmt:
        cStmt.CompileIfStmt(cpt, stmt)
    case *ast.ForStmt:
        cStmt.CompileForStmt(cpt, stmt)
    case *ast.RangeStmt:
        cStmt.CompileRangeStmt(cpt, stmt)
    case *ast.ReturnStmt:
        cStmt.CompileReturnStmt(cpt, stmt)
    case *ast.BlockStmt:
        cStmt.CompileBlockStmt(cpt, stmt)
    case *ast.ExprStmt:
        cStmt.CompileExprStmt(cpt, stmt)
    default:
        panic(fmt.Sprintf("syntax error: Bad Stmt Type %v", cpt.Fset.Position(stmt.Pos())))
    }
}

////声明stmt
//func (this *Stmt) CompileDeclStmt(cpt *CompileCxt, stmt *ast.DeclStmt) {
//	expr := NewExpr()
//	for _, spec := range stmt.Decl.Specs {
//		switch spec := spec.(type) {
//		case *ast.ValueSpec:
//			if nil != spec.Values {
//				if len(spec.Values) == len(spec.Names) {
//					for i, n := range spec.Names {
//						v := spec.Values[i]
//						this.SetValue(n.(*ast.Ident).Name, expr.CompileExpr(cpt, this, v), true)
//					}
//				} else if len(spec.Names) > len(spec.Values) && 1 == len(spec.Values) {
//					for i, n := range spec.Names {
//						v := spec.Values[0]
//						this.SetValue(n.(*ast.Ident).Name, expr.CompileExpr(cpt, this, v), true)
//					}
//
//				} else {
//					panic("syntax error: nonsupport spec num not match")
//				}
//			} else {
//				for _, n := range spec.Names {
//					switch spec.Type.(*ast.Ident).Name {
//					case "int":
//						var v int
//					case "int32":
//						var v int32
//					case "int64":
//						var v int64
//					case "string":
//						var v string
//					case "float":
//						var v float
//					case "float32":
//						var v float32
//					case "float64":
//						var v float64
//					case "bool":
//						var v bool
//					}
//					this.SetValue(n.(*ast.Ident).Name, expr.CompileExpr(cpt, this, v), true)
//				}
//			}
//		default:
//			panic("syntax error: nonsupport spec ")
//		}
//	}
//}

//表达式stmt，目前只支持callExpr
func (this *Stmt) CompileExprStmt(cpt *CompileCxt, stmt *ast.ExprStmt) {
	expr := NewExpr()
	switch X := stmt.X.(type) {
	case *ast.CallExpr:
		expr.CompileExpr(cpt, this, X)
	default:
        panic(fmt.Sprintf("syntax error: Bad ExprStmt Type %v", cpt.Fset.Position(stmt.Pos())))

	}
}

func (this *Stmt) CompileBlockStmt(cpt *CompileCxt, stmt *ast.BlockStmt) {
    //fmt.Println("-----------------in block stmt")
    for _, b := range stmt.List {
        this.CompileStmt(cpt, b)
    }
}

//获取值 回溯所有父节点获取值
func (this *Stmt) GetValue(name string) interface{} {
    stmt := this
    for nil != stmt {
        //fmt.Println("now stmt rct is ", stmt.Rct.ToString())
        if _, exist := stmt.Rct.Vars[name]; exist {
            return stmt.Rct.Vars[name]
        }
        stmt = stmt.Father
    }
    panic(fmt.Sprintf("syntax error: Item Can Not Reach %v", name))
}

func (this *Stmt) ValueExist(name string) (bool, *Stmt) {
    stmt := this
    for nil != stmt {
        //fmt.Println("now stmt rct is ", stmt.Rct.ToString())
        if _, exist := stmt.Rct.Vars[name]; exist {
            return true, stmt
        }
        stmt = stmt.Father
    }
    return false, nil
}

func (this *Stmt) SetValue(name string, value interface{}, create bool) {
    if create {
        //只在本节点内存中做校验
        if exist := this.Rct.ValueExist(name); exist {
            panic(fmt.Sprintf("syntax error: Redeclare Value %v", name))
        } else {
            this.Rct.SetValue(name, value)
        }
    } else {
        //只在本节点内存中做校验
        if exist, node := this.ValueExist(name); !exist {
            panic(fmt.Sprintf("syntax error: Undeclare Value %v", name))
        } else {
            nowValue := node.Rct.GetValue(name)
            //做类型校验
            if reflect.TypeOf(value).Kind() != reflect.TypeOf(nowValue).Kind() {
                panic(fmt.Sprintf("syntax error: Value Type Not Match %v have %v want %v", name, reflect.TypeOf(nowValue), reflect.TypeOf(value)))
            }
            node.Rct.SetValue(name, value)
        }
    }
}

func (this *Stmt) CompileAssignStmt(cpt *CompileCxt, stmt *ast.AssignStmt) {
    //fmt.Println("-----------------in assign stmt")
    //只支持= :=
    if token.DEFINE != stmt.Tok && token.ASSIGN != stmt.Tok {
        panic(fmt.Sprintf("syntax error: Bad Tok %v", cpt.Fset.Position(stmt.Pos())))
    }

    expr := NewExpr()
    //Lhs中的变量进行声明
    if len(stmt.Lhs) == len(stmt.Rhs) {
        for idx, l := range stmt.Lhs {
            switch l := l.(type) {
            case *ast.Ident:
                r := stmt.Rhs[idx]
                this.SetValue(l.Name, expr.CompileExpr(cpt, this, r), token.DEFINE == stmt.Tok)
            case *ast.IndexExpr:
                r := stmt.Rhs[idx]
                target := expr.CompileExpr(cpt, this, l.X)
                idx := expr.CompileExpr(cpt, this, l.Index)
                switch target := target.(type) {
                case map[interface{}]interface{}:
                    target[idx] = expr.CompileExpr(cpt, this, r)
                    this.SetValue(l.X.(*ast.Ident).Name, target, false)
                case []interface{}:
                    switch idx := idx.(type) {
                    case int:
                        target[idx] = expr.CompileExpr(cpt, this, r)
                    default:
                        panic(fmt.Sprintf("syntax error: Bad Index Type %v", cpt.Fset.Position(stmt.Pos())))
                    }
                    this.SetValue(l.X.(*ast.Ident).Name, target, false)
                default:
                    panic(fmt.Sprintf("syntax error: Bad Assign Type %v", cpt.Fset.Position(stmt.Pos())))
                }
            default:
                panic(fmt.Sprintf("syntax error: Bad Assign Type %v", cpt.Fset.Position(stmt.Pos())))
            }
        }
    } else if len(stmt.Lhs) > len(stmt.Rhs) && 1 == len(stmt.Rhs) {
        //声明语句不能嵌套，如果Rhs的元素是方法，则执行多返回值编译逻辑
        r := stmt.Rhs[0]
        switch r := r.(type) {
        case *ast.CallExpr:
            funcRet := expr.CompileCallMultiReturnExpr(cpt, this, r)
            if len(funcRet) != len(funcRet) {
                panic(fmt.Sprintf("syntax error: Func Return Nums Not Match %v", cpt.Fset.Position(stmt.Pos())))
            }
            for k, l := range stmt.Lhs {
                this.SetValue(l.(*ast.Ident).Name, funcRet[k], token.DEFINE == stmt.Tok)
            }
        case *ast.IndexExpr:
            if 2 == len(stmt.Lhs) && 1 == len(stmt.Rhs) {
                //处理v, exist := a[b]的情况
                target := expr.CompileExpr(cpt, this, stmt.Rhs[0].(*ast.IndexExpr).X)
                switch target := target.(type) {
                case map[interface{}]interface{}:
                    idx := expr.CompileExpr(cpt, this, stmt.Rhs[0].(*ast.IndexExpr).Index)
                    kName := stmt.Lhs[0].(*ast.Ident).Obj.Name
                    vName := stmt.Lhs[1].(*ast.Ident).Obj.Name
                    kVar, vExist := target[idx]
                    this.SetValue(kName, kVar, token.DEFINE == stmt.Tok)
                    this.SetValue(vName, vExist, token.DEFINE == stmt.Tok)
                default:
                    panic(fmt.Sprintf("syntax error: Bad ExistStmt Type %v", cpt.Fset.Position(stmt.Pos())))
                }
            }
		default:
			for _, l := range stmt.Lhs {
				switch l := l.(type) {
				case *ast.Ident:
					this.SetValue(l.Name, expr.CompileExpr(cpt, this, r), token.DEFINE == stmt.Tok)
				default:
                    panic(fmt.Sprintf("syntax error: Bad Index Type ExistStmt %v", cpt.Fset.Position(stmt.Pos())))
				}
			}
        }
    } else {
        panic(fmt.Sprintf("syntax error: Bad AssignStmt Nums %v", cpt.Fset.Position(stmt.Pos())))
    }
}


func (this *Stmt) CompileForStmt(cpt *CompileCxt, stmt *ast.ForStmt) {
    //fmt.Println("----------------in for stmt")
    stmtHd := this.NewChild()
    expr := NewExpr()
    //初始条件
    this.CompileStmt(cpt, stmt.Init)
    for {
        if access := expr.CompileExpr(cpt, this, stmt.Cond); !cast.ToBool(access) {
            break;
        }
        //执行body
        stmtHd.CompileStmt(cpt, stmt.Body)
        this.CompileStmt(cpt, stmt.Post)
    }
}

func (this *Stmt) CompileIfStmt(cpt *CompileCxt, stmt *ast.IfStmt) {
    //fmt.Println("----------------in if stmt")
    stmtHd := this.NewChild()
    expr := NewExpr()
    //赋值操作,在本节点赋值
    this.CompileStmt(cpt, stmt.Init)
    condRet := expr.CompileExpr(cpt, this, stmt.Cond)
    //如果条件成立
    if cast.ToBool(condRet) {
        stmtHd.CompileStmt(cpt, stmt.Body)
    } else {
        stmtHd.CompileStmt(cpt, stmt.Else)
    }
}

//只支持变量
func (this *Stmt) CompileIncDecStmt(cpt *CompileCxt, stmt *ast.IncDecStmt) {
    //fmt.Println("----------------in inc dec stmt")
    //只支持 ++ --
    if token.INC != stmt.Tok && token.DEC != stmt.Tok {
        panic(fmt.Sprintf("syntax error: Bad Tok %v", cpt.Fset.Position(stmt.Pos())))
    }

    varName := stmt.X.(*ast.Ident).Name
    switch stmt.Tok {
    case token.INC:
        //this.SetValue(varName, expr.CompileExpr(cpt, this, stmt.X))
        this.SetValue(varName, BInc(this.GetValue(varName)), false)
    case token.DEC:
        //this.SetValue(varName, expr.CompileExpr(cpt, this, stmt.X))
        this.SetValue(varName, BDec(this.GetValue(varName)), false)
    default:
        panic(fmt.Sprintf("syntax error: Bad Tok %v", cpt.Fset.Position(stmt.Pos())))
    }
}

func (this *Stmt) CompileRangeStmt(cpt *CompileCxt, stmt *ast.RangeStmt) {
    //fmt.Println("----------------in range stmt")
    expr := NewExpr()
    stmtHd := this.NewChild()
    RangeTarget := expr.CompileExpr(cpt, this, stmt.Key.(*ast.Ident).Obj.Decl.(*ast.AssignStmt).Rhs[0].(*ast.UnaryExpr).X)
    kName := stmt.Key.(*ast.Ident).Name
    vName := stmt.Key.(*ast.Ident).Obj.Decl.(*ast.AssignStmt).Lhs[1].(*ast.Ident).Name
    switch rt := RangeTarget.(type) {
    case []interface{}:
        for k, v := range rt {
            //设置kv的值
            isCreate := true
            if k > 0 {
                isCreate = false
            }
            stmtHd.SetValue(kName, k, isCreate)
            stmtHd.SetValue(vName, v, isCreate)
            //执行Body
            stmtHd.CompileStmt(cpt, stmt.Body)
        }
    case map[interface{}]interface{}:
        i := 0
        for k, v := range rt {
            //设置kv的值
            isCreate := true
            if i > 0 {
                isCreate = false
            }
            stmtHd.SetValue(kName, k, isCreate)
            stmtHd.SetValue(vName, v, isCreate)
            //执行Body
            stmtHd.CompileStmt(cpt, stmt.Body)
            i++
        }
    default:
        panic(fmt.Sprintf("syntax error: Bad RangeStmt Type %v", cpt.Fset.Position(stmt.Pos())))
    }
}

//支持返回只支持一个
func (this *Stmt) CompileReturnStmt(cpt *CompileCxt, stmt *ast.ReturnStmt) {
    //fmt.Println("----------------in return stmt")
    var ret interface{}
    expr := NewExpr()
    e := stmt.Results[0]
    ret = expr.CompileExpr(cpt, this, e)
    //fmt.Println("----------------return ", ret)
    cpt.ReturnCh <- ret
}
