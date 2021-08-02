### 背景

写自己的脚手架，二开开源脚手架，让程序写程序，解放生产力

### 知识点

* go元编程
    * 编译期生成代码
    * 运行时改变代码行为
      * 反射
      * 调试器
      * 外挂
* go-generate
* go-template
* ast

组合知识点，实现各种各样的功能

### 代码生成的本质

字符串 -> 文件 

goformat

需要思考的就是->如何花式去构建字符串

* 手动组装
* 模板生成
* 语法解析...
  （通过网络请求获取字符串）


### go-gen 

* 扫描 Go 语言源文件，查找待执行的 //go:generate 预编译指令；
* 执行预编译指令，再次扫描源文件并根据源文件中的代码生成代码；

### 例子


#### 玩具
step_test

#### go-gen
cmd/go/internal/generate/generate.go:163

#### strings
cmd/stringer/stringer.go:106

#### gen-grpc

cmd/protoc-gen-go-grpc/main.go:45

```
 protoc --go-grpc_out=. user.proto
```

#### go-zero

tools/goctl/goctl.go:447

tools/goctl/api/parser/parser.go:19


### 总结

https://www.processon.com/diagraming/60cc72a11efad410510dc849

### 参考

[元编程]
https://zh.wikipedia.org/wiki/%E5%85%83%E7%BC%96%E7%A8%8B

[go代码生成]
https://draveness.me/golang/docs/part4-advanced/ch08-metaprogramming/golang-code-gen/

[go-ast]
https://juejin.cn/post/6844903982683389960

[go-ast]
https://segmentfault.com/a/1190000039215176

[go-zero]
https://github.com/tal-tech/go-zero

[构建编译器]
https://craftinginterpreters.com/

[Golang template 小抄]
https://colobu.com/2019/11/05/Golang-Templates-Cheatsheet/

[pbgo: 基于Protobuf的框架]
https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-07-pbgo.html
