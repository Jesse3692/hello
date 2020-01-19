# hello

## 前置操作：

1. 在github创建一个仓库`hello`

2. 在`$GOPATH/src/github.com/Jesse3692`路径下面进行远程仓库克隆`git clone https://github.com/Jesse3692/hello.git`

3. 将代码提交到添加到本地缓存区`git add .`

4. 提交代码到本地仓库`git commit -m "first commit"`

5. 将代码推送到本地仓库`git push -u origin master`

## 第一次提交的hello world

```go
package main

import "fmt"

func main()  {
	fmt.Println("Hello, World")
}
```

代码解析：

`main package`不同于其他`library package`，它定义了一个可执行程序。其中的`main`函数即是执行文件的入口函数。包是一种将相关的Go代码组合到一起的方式。

func关键字通过函数名和函数体来定义函数。

通过`import "fmt"`导入一个包含`Println`函数的包，我们用它来打印输出。

## 对代码进行测试

由于之前的代码是没有返回值的（fmt.Println会打印内容到标准输出），所以是不能进行测试的。

现在进行如下的修改：

```go
// hello.go
package main

import "fmt"

func Hello() string  {
	return "Hello, World"
}

func main()  {
	fmt.Println(Hello())
}
```

我们在代码中新建了一个函数`Hello`，然后定义了另一个关键字`string`，这意味着函数将返回一个字符串。

下面为`Hello`函数编写测试，创建一个`hello_test.go`文件

```go
// hello_test.go
package main

import "testing"

func TestHello(t *testing.T)  {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got '%q' want '%q'", got ,want)
	}
}
```

此时运行`go test`对`hello.go`文件进行测试。一般的情况下是可以正常测试的，但是由于开启了`GO111MODULE`，它是会爆出一个错误的：

```
go: cannot find main module, but found .git/config in /home/jesse/goProject/src/github.com/Jesse3692/hello
        to create a module there, run:
        go mod init
```

根据给出的错误提示，我们运行以下的命令来创建一个模块`go mod init hello`，然后重新进行测试`go test`

测试通过：

```
PASS
ok      hello   0.002s
```

编写测试的规则：

- 程序需要在一个名为`xxx_test.go`的文件中编写

- 测试函数的命名必须以单词`Test`开始

- 测试函数只接收一个参数`t *testing.T`

测试代码解析：

类型为`*testing.T`的变量`t`是测试框架中的钩子（hook），当测试失败后可以执行`t.Fail()`之类的操作。

`want := "Hello, World"`是声明变量的语法（`varName := value`），是`var want string = "Hello, World"`的简写

`t.Errorf`是调用`t`的`Errorf`方法，当测试失败后会打印一条消息。`Errorf`中的`f`表示格式化，它允许我们构建一个字符串，并将值插入到占位符值`%q`中。

## 使用Go文档

由于 `go doc`和`godoc`在Golang1.13中被移除了，所以需要自行安装

```shell
$ go get golang.org/x/tools/cmd/godoc
```

在本地启动文档`godoc -http :8888`


## 对函数使用参数

```go
// hello.go
func Hello(name string) string  {
	return "Hello, " + name
}
```

为函数添加一个类型是`string`的变量`name`

## 使用常量

```go
const englishHelloPrefix = "Hello, "

// Hello func return "Hello, World"
func Hello(name string) string  {
	return englishHelloPrefix + name
}
```