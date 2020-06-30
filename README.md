# akiDsl
一个golang语法dsl的脚本解释器

## Install
````
go get github.com/nber1994/akiDsl
````

# Example
```
package main

import (
    "fmt"
    "flag"
    "github.com/nber1994/akiDsl"
)

var (
    mFile = flag.String("f", "", "file mode")
    mCxt = flag.String("c", "{}", "cxt str")
)

func main() {
    flag.Parse()
    if *mFile == "" {
        fmt.Println("no file input")
    }

    akiDslNode, newErr := akiDsl.New(mFile, mCxt)
    if nil != newErr {
        fmt.Println("new err: ", newErr)
    }
    dslRet, cxt, err := akiDslNode.Run()
    if nil != err {
        fmt.Println(">> the run error is: ", err)
        return
    }
    fmt.Println(">> the dsl run result is: ", dslRet)
    fmt.Println(">> the cxt run result is: ", cxt.ToJsonString())
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
  return a
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
    aaaa, aaaaa := 1
    aaaa, aaaaa = 1
    aaaaaaaaaaa := map[string]string{}

    //集合类型
    b := []int{1,2,3}
    bb := []string{"i", "have", "some", "special"}
    bbb := []map[string]string{{"nber":"1994","1994":"nber"},{"abc":"666","666":"abc"}}

    //复合集合
    c := map[string]string{"jty":"1994-11-10", "ldh":"1987-12-22","wyz":"1988-12-01"}
    cc := map[string][]string{"jty":[]string{"so", "a", "big", "awesome", "boy"}, "ldh":[]string{"as", "good", "as", "jty"}}

    //取值
    d := b[1]
    dd := b[a]
    ddd := bbb[1]["nber"]
    dddd := c["jty"]
    ddddd := cc[aaa][1]
    dddddd, exist := c["jty"]
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
    return a
}
```

## If
```
package main

func main() {
    a := 1
    a++
    a--
}
12:41:59 test cat testIf
package main

func main() {
    a := 1
    b := ""
    if a > 2 {
        b = "bigger than"
    } else {
        b = "smaller than"
    }

    c := 1
    if b == "smaller than" {
        c = 2
    } else if b == "bigger than" {
        c = 3
    } else {
        c = 4
    }

    x := map[string]string{"a":"aa","b":"bb"}
    if _, exist := x["a"]; exist {
        x["a"] = "nber1994"
    }
}
```

## For
```
package main

func main() {
    a := 1
    for i:=0;i<10;i++ {
        a = a + 1
    }

    for i,j := 0, 0; i+j < 10;i++ {
        a = a + 1
    }
}
```

## Range
```
package main

func main() {
    a := []int{1,2,3}
    b := 0
    for _, v := range a {
        b = v
    }

    c := map[string]string{"123":"456", "nber1994":"jty"}
    d := ""
    for k, v := range c {
        d = k
        d = v
    }
}
```

## Call Func
```
package main

func main() {
    c, x := 1
    a := Get("dt.year")
    Set("dt.year", 666)
    a = Get("dt.year")


    arr := []int{}
    arr = append(arr, 1)
    return arr
}
```
目前支持的函数列表 appending...
```
Get(path string)
Set(path string value interface{})
append(a []interface{}, item interface{})
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
    a = 3
    b := 3
}
```

项目疯狂迭代中，如有问题我们issue见

