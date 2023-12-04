# miniblog

## 项目初始化

在 `cmd/miniblog`文件中新建 `main`

```golang
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

```

运行脚本

```bash
# 整理代码格式
go fmt -s -w ./
# 定好输出文件
go build -o _output/miniblog.exe ./cmd/miniblog/main.go
```

##### 热重载

使用 `air`实现代码热重载

###### **安装**

```golang
go install github.com/
```

###### 使用

在目录下新建配置文件 `.ari.toml`

```toml
cmd = "go build -o xxxx xxx.go"
bin = "xxxx" # 执行cmd执行后生成的文件
```
