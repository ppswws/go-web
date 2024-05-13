<h1 align="center" >Go Web</h1>

<div align="center">
这是学习分支
</div>
<p align="center">
<img src="https://img.shields.io/badge/Go-v1.16-blue" alt="Go version"/>
<img src="https://img.shields.io/badge/Gin-v1.7.2-brightgreen" alt="Gin version"/>
<img src="https://img.shields.io/badge/Gorm-v1.21.11-brightgreen" alt="Gorm version"/>
<img src="https://img.shields.io/github/license/cyj19/go-web" alt="License"/>
</p>


## 特性
- `RESTful API` 设计风格
- `MySQL`       数据库存储
- `gin`         golang web 的微框架
- `gorm`        数据库的ORM管理框架
- `gin-jwt`     gin封装的jwt中间件，用户认证
- `casbin`      轻量级开源访问控制框架，RBAC
- `go-redis`    redis客户端开发工具
- `viper`       轻便的golang配置管理工具
- `zap`         高性能日志库，提供多种级别的日志打印
- `lumberjack`  日志文件切割归档工具

## 项目结构
```
├── cmd
│    └── admin # admin项目主程序入口
├── configs # 配置目录
├── internal # 内部目录，不对外公开
│    ├── admin # admin项目目录
│    │     ├── api # api目录
│    │     │    └── v1 # v1版本接口目录(类似于Java中的controller), 如果有新版本可以继续添加v2/v3
│    │     ├── global # 全局公用模型目录
│    │     ├── router # 路由目录
│    │     ├── service # 业务逻辑目录
│    │     │    └── v1 # v1版本业务目录, 如果有新版本可以继续添加v2/v3
│    │     ├── store # 数据操作目录
│    │     ├── data.go # 初始化数据
│    │     └── router.go # 定义路由规则
│    ├── pkg # 内部公共模块目录
│    │     ├── cache # redis操作目录
│    │     ├── cofnig # 配置实体目录
│    │     ├── db # 数据库目录
│    │     ├── initialize # 工具初始化目录
│    │     ├── logger # 日志目录
│    │     ├── middleware # 中间件目录
│    │     ├── model # 传输模型目录
│    │     ├── response # 响应模型目录
│    │     └── util # 工具包目录
├── logs # 日志文件目录
├── pkg # 外部公共模块目录
├── .gitignore # git忽略
├── go.mod # go依赖列表
├── go.sum # go依赖下载历史
├── LICENSE # 开源证书
├── README.md # 说明文档
