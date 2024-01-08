# miniblog

项目介绍：一个博客后端系统，主要实现了用户的注册、增删改查和博客的增删改查

## 项目初始化

### 程序实时加载、构建、启动

使用`air`可以做到程序热加载

#### 1. 安装`air`工具
```bash
go install github.com/cosmtrek/air@latest 
```
#### 2. 配置`air`工具

在项目根目录下新建.air.toml
```toml
# 这里编写shell命令，可以使用make
cmd = "go build"
# 这里指定执行的文件名
bin = "_output/miniblog"
```
