package compile

import (
    "fmt"
    "go/ast"
    "go/token"
	"github.com/spf13/cast"
)

type Stmt struct{
}

func NewStmt() *Stmt{
    return &Stmt{}
}

//编译stmt
func (this *Stmt) CompileStmt(cpt *CompileCxt, stmt ast.Stmt) {
    if nil == stmt {
        return
    }
    switch stmt := stmt.(type) {
    case *ast.AssignStmt:
        this.CompileAssignStmt(cpt, stmt)
    case *ast.IncDecStmt:
        this.CompileIncDecStmt(cpt, stmt)
    case *ast.IfStmt:
        this.CompileIfStmt(cpt, stmt)
    case *ast.ForStmt:
        this.CompileForStmt(cpt, stmt)
    case *ast.RangeStmt:
        this.CompileRangeStmt(cpt, stmt)
    case *ast.ReturnStmt:
        this.CompileReturnStmt(cpt, stmt)
    case *ast.BlockStmt:
        this.CompileBlockStmt(cpt, stmt)
    default:
        panic("syntax error: nonsupport stmt ")
    }
}

func (this *Stmt) CompileBlockStmt(cpt *CompileCxt, stmt *ast.BlockStmt) {
    fmt.Println("-----------------in block stmt")
    for _, b := range stmt.List {
        this.CompileStmt(cpt, b)
    }
}

func (this *Stmt) CompileAssignStmt(cpt *CompileCxt, stmt *ast.AssignStmt) {
    fmt.Println("-----------------in assign stmt")
    //只支持= :=
    if token.DEFINE != stmt.Tok && token.ASSIGN != stmt.Tok {
        panic("syntax error: nonsupport Tok ")
    }

    expr := NewExpr()

    //Lhs中的变量进行声明
    if len(stmt.Lhs) == len(stmt.Rhs) {
        for idx, l := range stmt.Lhs {
            switch l := l.(type) {
            case *ast.Ident:
                r := stmt.Rhs[idx]
                cpt.RunCxt.SetValue(l.Name, expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, r))
            default:
                panic("syntax error: assign type must be ident type")
            }
        }
    } else if len(stmt.Lhs) > len(stmt.Rhs) && 1 == len(stmt.Rhs){
        //声明语句不能嵌套，如果Rhs的元素是方法，则执行多返回值编译逻辑
        r := stmt.Rhs[0]
        switch r := r.(type) {
        case *ast.CallExpr:
            funcRet := expr.CompileCallMultiReturnExpr(cpt.DslCxt, cpt.RunCxt, r)
            if len(funcRet) != len(funcRet) {
                panic("syntax error: func return can not match")
            }
            for k, l := range stmt.Lhs {
                cpt.RunCxt.SetValue(l.(*ast.Ident).Name, funcRet[k])
            }
        case *ast.IndexExpr:
            if 2 == len(stmt.Lhs) && 1 == len(stmt.Rhs) {
                //处理v, exist := a[b]的情况
                target := expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, stmt.Rhs[0].(*ast.IndexExpr).X)
                switch target := target.(type) {
                case map[interface{}]interface{}:
                    idx := expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, stmt.Rhs[0].(*ast.IndexExpr).Index)
                    kName := stmt.Lhs[0].(*ast.Ident).Obj.Name
                    vName := stmt.Lhs[1].(*ast.Ident).Obj.Name
                    kVar, vExist := target[idx]
                    cpt.RunCxt.SetValue(kName, kVar)
                    cpt.RunCxt.SetValue(vName, vExist)
                default:
                    panic("syntax error: index exist assign stmt type error")
                }
            }
        default:
            panic("syntax error: assign nums not match")
        }

    }
}


func (this *Stmt) CompileForStmt(cpt *CompileCxt, stmt *ast.ForStmt) {
    fmt.Println("----------------in for stmt")
    stmtHd := NewStmt()
    expr := NewExpr()
    //编译初始条件
    stmtHd.CompileStmt(cpt, stmt.Init)
    for {
        if access := expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, stmt.Cond); !cast.ToBool(access) {
            break;
        }
        //执行body
        stmtHd.CompileStmt(cpt, stmt.Body)
        stmtHd.CompileStmt(cpt, stmt.Post)
    }
}

func (this *Stmt) CompileIfStmt(cpt *CompileCxt, stmt *ast.IfStmt) {
    fmt.Println("----------------in if stmt")
    expr := NewExpr()
    condRet := expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, stmt.Cond)
    stmtHd := NewStmt()
    fmt.Println("----------------in if init stmt")
    fmt.Println(stmt)
    stmtHd.CompileStmt(cpt, stmt.Init)
    //如果条件成立
    if cast.ToBool(condRet) {
        stmtHd.CompileStmt(cpt, stmt.Body)
    } else {
        stmtHd.CompileStmt(cpt, stmt.Else)
    }
}

//只支持变量
func (this *Stmt) CompileIncDecStmt(cpt *CompileCxt, stmt *ast.IncDecStmt) {
    fmt.Println("----------------in inc dec stmt")
    //只支持 ++ --
    if token.INC != stmt.Tok && token.DEC != stmt.Tok {
        panic("syntax error: nonsupport Tok ")
    }

    expr := NewExpr()
    varName := stmt.X.(*ast.Ident).Name
    switch stmt.Tok {
    case token.INC:
        cpt.RunCxt.SetValue(varName, expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, stmt.X))
        cpt.RunCxt.SetValue(varName, BInc(cpt.RunCxt.GetValue(varName)))
    case token.DEC:
        cpt.RunCxt.SetValue(varName, expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, stmt.X))
        cpt.RunCxt.SetValue(varName, BDec(cpt.RunCxt.GetValue(varName)))
    default:
        panic("syntax error: nonsupport Tok ")
    }
}

func (this *Stmt) CompileRangeStmt(cpt *CompileCxt, stmt *ast.RangeStmt) {
    fmt.Println("----------------in range stmt")
    expr := NewExpr()
    stmtHd := NewStmt()
    RangeTarget := expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, stmt.Key.(*ast.Ident).Obj.Decl.(*ast.AssignStmt).Rhs[0].(*ast.UnaryExpr).X)
    kName := stmt.Key.(*ast.Ident).Name
    vName := stmt.Key.(*ast.Ident).Obj.Decl.(*ast.AssignStmt).Lhs[1].(*ast.Ident).Name
    switch rt := RangeTarget.(type) {
    case []interface{}:
        for k, v := range rt {
            //设置kv的值
            cpt.RunCxt.SetValue(kName, k)
            cpt.RunCxt.SetValue(vName, v)
            //执行Body
            stmtHd.CompileStmt(cpt, stmt.Body)
        }
    case map[interface{}]interface{}:
        for k, v := range rt {
            //设置kv的值
            cpt.RunCxt.SetValue(kName, k)
            cpt.RunCxt.SetValue(vName, v)
            //执行Body
            stmtHd.CompileStmt(cpt, stmt.Body)
        }
    default:
        panic("syntax error: nonsupport range type")
    }
}

//支持返回只支持一个
func (this *Stmt) CompileReturnStmt(cpt *CompileCxt, stmt *ast.ReturnStmt) {
    fmt.Println("----------------in return stmt")
    var ret interface{}
    expr := NewExpr()
    e := stmt.Results[0]
    ret = expr.CompileExpr(cpt.DslCxt, cpt.RunCxt, e)
    fmt.Println("----------------return ", ret)
    cpt.ReturnCh <- ret
}
