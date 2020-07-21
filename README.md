
项目思路：[实现一个简单的DSL解释器](http://nber1994.com/2020/07/02/go-%E5%AE%9E%E7%8E%B0%E4%B8%80%E4%B8%AA%E7%AE%80%E5%8D%95%E7%9A%84DSL%E8%A7%A3%E9%87%8A%E5%99%A8.html)

# akiDsl
这个项目简单来说就是一个golang的脚本解释器

## Install
````
go get github.com/nber1994/akiDsl
````

# Example
```
package main

import (
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "flag"
    "github.com/nber1994/akiDsl"
)

var (
    mFile = flag.String("f", "", "file mode")
    mCxt = flag.String("c", "", "cxt str")
)

func main() {
    flag.Parse()
    if *mFile == "" {
        fmt.Println("no file input")
    }

    fset := token.NewFileSet()
    fAst, err := parser.ParseFile(fset, *mFile, nil, 0)
    if err != nil {
        fmt.Println(err)
        return
    }

    ast.Print(fset, fAst)
    fmt.Println("==========================================")

    dslCxt, newErr := akiDsl.NewCxt(*mCxt)
    if nil != newErr {
        fmt.Println("new dsl cxt err: ", newErr)
    }
    akiDslNode := akiDsl.New(mFile, dslCxt)
    dslRet, cxt, err := akiDslNode.Run()
    fmt.Println("==========================================")
    fmt.Println(">> err: ", err)
    fmt.Println(">> ret ", dslRet)
    fmt.Println(">> cxt ", cxt.ToJsonString())

}

```
# Context
dsl一般应用于特定领域，且通过上下文与嵌入的系统进行通信，akiDsl的Context为字典形式json，且提供内置方法对Context进行读写

# Syntax
dsl提供了与golang相同的语法（not all），减少dsl的语法学习成本
以下为目前支持的语法：(觉得不够？我们issue见)

## main
所有的代码，必须包在main()函数中，并且允许在main中return任意格式的数据用于返回dsl执行结果
```
package main

func main() {
  a := 1
  return a （我知道这个不符合语法，但是作为一个dsl需要一个返回值，算是半个语法糖吧）
}
```

## Assign
```
package main

func main() {
    //基本数据
    a := 1
    aa := 1.2
    aaa := "jty"
    aaaa := true
    aaaaa := []int{1,2,3,4,5}
    aaaaaa := []string{"a", "b", "c"}
    aaaaaaa := []float{1.2, 1.3, 1.33}
    aaaaaaaa := map[int]int{1:1, 2:2, 3:3}
    aaaaaaaaa := map[int]float{1:1.1, 2:2.11, 3:22.1}
    aaaaaaaaaa := map[int]string{1:"a", 2:"b", 3:"c"}
    aaaaaaaaaaa := map[int]bool{1:true, 2:false, 3:true}
    aaaaaaaaaaaa := map[string]string{"a":"a", "b":"b", "c":"c"}
    aaaaaaaaaaaaa := map[float]float{1.1:1.2, 1.22:1.3, 1.11:2.22}
    aaaaaaaaaaaaaa := map[float]bool{1.2:true, 2.11:false}

    //复合集合
    b := map[string][]string{"jty":[]string{"so", "a", "big", "awesome", "boy"}, "ldh":[]string{"as", "good", "as", "jty"}}
    bb := []map[string]string{map[string]string{"jty":"1994-11-10", "ldh":"1987-12-22","wyz":"1988-12-01"}, map[string]string{"jty":"1994-11-10", "ldh":"1987-12-22","wyz":"1988-12-01"}}

    //其他操作
    c, cc := 1
    ccc, cccc := "a", "b"


    println(a)
    println(aa)
    println(aaa)
    println(aaaa)
    println(aaaaa)
    println(aaaaaa)
    println(aaaaaaa)
    println(aaaaaaaa)
    println(aaaaaaaaa)
    println(aaaaaaaaaa)
    println(aaaaaaaaaaa)
    println(aaaaaaaaaaaa)
    println(aaaaaaaaaaaaa)
    println(b)
    println(bb)
    println(c)
    println(cc)
    println(ccc)
    println(cccc)
}
```
声明（declare）语法暂不支持（感觉也没啥用）

## IncDec
```
package main

func main() {
    a := 1
    a++
    a--
    println(a)
}
```

## If
```
package main

func main() {
    a := 1
    b := ""
    if a > 2 {
        b = "bigger than"
    } else {
        b = "smaller than"
    }
    println(b)

    c := 1
    if b == "smaller than" {
        c = 2
    } else if b == "bigger than" {
        c = 3
    } else {
        c = 4
    }
    println(c)

    x := map[string]string{"a":"aa","b":"bb"}
    if _, exist := x["a"]; exist {
        x["a"] = "nber1994"
    }
    println(x)
}
```

## For
```
package main

func main() {
    a := 1
    for i:=0;i<10;i++ {
        a = a + 1
        println(a)
    }

    for i,j := 0, 0; i+j < 10;i++ {
        a = a + i
        println(a)
    }
}
```

## Range
```
package main

func main() {
    a := []int{1,2,3}
    b := 0
    for k, v := range a {
        println("key: ", k, " value: ", v)
    }

    c := map[string]string{"123":"456", "nber1994":"jty"}
    for k, v := range c {
        println("key: ", k, " value: ", v)
    }
}
```

## Call Func
```
package main

func main() {
    if Exist("dt.year") {
        println(GetInt("dt.year") + 1)
    }
    Set("dt.year", 666)

    arr := []int{}
    arr = append(arr, 1)
    println(arr)
    c := "nber1994"
    println(sprintf("sprintf %v", c))
}
```
目前支持的函数列表 appending...
```
Dsl Context操作相关函数
GetInt(path string) Int
GetFloat(path string) Float64
GetString(path string) string
GetBool(path string) bool
Set(path string value interface{})
Exist(path string) bool

常用函数
append(a []interface{}, item interface{})
println(v ...interface{})
sprintf(v ...interface{})
```
不够？我们issue见

## Local Variable
```
package main

func main() {
    a := 2

    for i:=0;i<10;i++ {
        a++
        b := 2
    }
    println(b)
}
```

项目疯狂迭代中，如有问题我们issue见

# change log
* v0.0.1 
    * init
* v0.0.2 
    * 代码优化
    * 修改了入参，将上下文从string类型改为内部结构体传入。这样在每个dsl执行节点，都会省去unmarshal步骤
