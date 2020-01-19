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

