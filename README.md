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

```bash
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

### 添加LICENSE

安装工具：`go install github.com/nishanths/license/v5@latest`

为项目添加 `license`

```bash
license -n 'huangyul 1103221645@qq.com' -o LICENSE mit

```

### 为源文件添加声明

安装 `addlicense`

`go install github.com/marmotedu/addlicense@latest`

执行命令为所有文件添加声明

`addlicense -v -f ./scripts/boilerplate.txt --skip-dirs=third_party,vendor,_output .`

## 构建应用

使用`pflag`,`viper`,`cobra`来构建一个应用程序（应用程序包括应用配置，业务配置，启动框架）
- `pflag`: 读取命令行参数。[教程](https://link.juejin.cn/?target=https%3A%2F%2Fgithub.com%2Fmarmotedu%2Fgeekbang-go%2Fblob%2Fmaster%2F%25E5%25A6%2582%25E4%25BD%2595%25E4%25BD%25BF%25E7%2594%25A8Pflag%25E7%25BB%2599%25E5%25BA%2594%25E7%2594%25A8%25E6%25B7%25BB%25E5%258A%25A0%25E5%2591%25BD%25E4%25BB%25A4%25E8%25A1%258C%25E6%25A0%2587%25E8%25AF%2586.md)
- `viper`: 读取配置文件（yaml、json等）。[教程](https://link.juejin.cn/?target=https%3A%2F%2Fgithub.com%2Fmarmotedu%2Fgeekbang-go%2Fblob%2Fmaster%2F%25E9%2585%258D%25E7%25BD%25AE%25E8%25A7%25A3%25E6%259E%2590%25E7%25A5%259E%25E5%2599%25A8-Viper%25E5%2585%25A8%25E8%25A7%25A3.md)
- `cobra`: 命令行工具。[教程](https://link.juejin.cn/?target=https%3A%2F%2Fgithub.com%2Fmarmotedu%2Fgeekbang-go%2Fblob%2Fmaster%2F%25E7%258E%25B0%25E4%25BB%25A3%25E5%258C%2596%25E7%259A%2584%25E5%2591%25BD%25E4%25BB%25A4%25E8%25A1%258C%25E6%25A1%2586%25E6%259E%25B6-Cobra%25E5%2585%25A8%25E8%25A7%25A3.md)