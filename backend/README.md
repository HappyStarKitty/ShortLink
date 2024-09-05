# 项目文档
## 项目主要依赖
- Gin: 用于构建高性能Web应用的HTTP框架。
- GORM: 用于数据库操作的ORM库，提供了简单、灵活的数据库交互方式。
- Go Config: 用于加载和解析配置文件的库，帮助从外部文件加载环境配置。
- CORS: 用于处理跨源资源共享的中间件，支持跨域请求。
-Log: 用于日志记录，帮助调试和维护项目。
## 项目目录结构
```bash
.
├── api/                    # API 层，处理请求和响应
│   ├── dto/                # 数据传输对象 (Data Transfer Objects)
│   │   ├── general.go      # 通用 DTO
│   │   ├── link.go         # 短链接相关 DTO
│   │   ├── login.go        # 登录相关 DTO
│   │   └── user.go         # 用户相关 DTO
│   ├── route/              # 路由层，定义 API 路由
│   │   └── route.go        # 路由配置文件
│   └── init.go             # API 初始化文件
├── internal/               # 内部实现逻辑
│   ├── controller/         # 控制器层，处理 HTTP 请求
│   │   ├── link.go         # 短链接相关控制器
│   │   └── user.go         # 用户相关控制器
│   ├── dao/                # 数据访问层，负责与数据库交互
│   │   ├── model/          # 数据库模型
│   │   │   ├── link.go     # 短链接模型
│   │   │   └── user.go     # 用户模型
│   ├── service/            # 服务层，业务逻辑
│   │   ├── link_service.go # 短链接业务逻辑
│   │   └── user_service.go # 用户业务逻辑
│   ├── complex_crud.go     # 复杂的增删改查逻辑
│   └── dao.go              # DAO 层的初始化和配置
├── utils/                  # 工具库
│   ├── config.go           # 配置加载工具
│   ├── errors.go           # 错误处理工具
│   ├── log.go              # 日志工具
│   └── security.go         # 安全工具
├── config/                 # 配置文件目录
│   └── config.yaml         # 应用程序配置文件
├── main.go                 # 应用程序入口
├── Dockerfile              # Docker 部署配置文件
└── go.mod                  # Go 模块配置文件

```
## 如何运行

1. 下载依赖

在项目根目录下运行以下命令来安装依赖：
```bash
go mod tidy
```
2. 运行服务器

使用以下命令启动服务器：
```bash
go run main.go
```
3. 访问应用

服务器将运行在 http://localhost:8080。

## 如何测试
为了保证代码质量并测试controller层的功能，项目使用 testify 和 gin 提供的模拟工具。

编写测试

测试文件位于 internal/controller/ 目录下，例如：

```bash

internal/controller/link_controller_test.go
internal/controller/user_controller_test.go
```

运行测试

使用以下命令运行测试：

```bash
go test ./internal/controller -v
```

## Docker 部署
构建Docker镜像

运行以下命令构建镜像：

```bash
docker build -t shortlink-backend .
```

运行Docker容器

使用以下命令运行容器：

```bash
docker run -d -p 8080:8080 shortlink-backend
```

理论上是这样，不过我失败了qaq

## 感想建议
1. 模块化设计：

项目按层次划分（controller, service, dao），结构清晰，逻辑分明，便于维护和扩展。
错误处理和日志记录：

2. 错误处理：
通过完善的错误处理机制和日志记录，项目的可维护性大大增强，
