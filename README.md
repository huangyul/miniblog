# miniblog

## 项目初始化

### 新建main文件

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

### 热重载

使用 `air`实现代码热重载

#### **安装**

```golang
go install github.com/
```

#### 使用

在目录下新建配置文件 `.ari.toml`

```toml
cmd = "go build -o xxxx xxx.go"
bin = "xxxx" # 执行cmd执行后生成的文件
```

### API文档

使用swagger

1. 先在[swagger editor](https://link.juejin.cn/?target=https%3A%2F%2Feditor-next.swagger.io%2F)
2. 选择一个模板，导出成 `YAML`文件，放到项目根目录中
3. 安装 `go-swagger`工具：**go install github**.**com**/**go**-**swagger**/**go**-**swagger**/**cmd**/**swagger@latest**
4. 启动工具：swagger serve -F=swagger --no-open --port **65534**./**api**/**openapi**/**openapi**.yaml
