package runCxt

import (
    "go/token"
)

type MemBlock struct {
    Var string //数据的值
    VarType token.Token //数据类型
}

var SupportTypeArr []token.Token = []token.Token{ 
    token.INT,
    token.FLOAT,
    token.STRING,
}

func NewMemBlock(v string, t token.Token) *MemBlock {
    //校验是否支持的类型
    inArr := false
    for _, SupportType := range SupportTypeArr {
        if SupportType == t {
            inArr = true
        }
    }

    if !inArr {
        panic("syntax error: nonsupport type ")
    }

    return &MemBlock{
        Var: v,
        VarType: t,
    }
}
