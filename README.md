# go-starter

`go-starter` 是一个为 Golang 项目开发设计的基础模板，内置 VSCode Dev Container 支持、GitHub Actions CI/CD、统一代码规范校验、自动发布构建等常见开发工具链，开箱即用。

## ⚠️ 重要提示

**在使用此模板前，请将 `app/your-app` 目录重命名为您的实际应用名称，例如：**

- `app/my-service`
- `app/user-api`
- `app/order-service`

**同时需要更新以下文件中的路径引用：**

- `Makefile`
- `README.md`
- 配置文件中的相关路径

[![Unittest with Coverage](https://github.com/phenix3443/go-starter/actions/workflows/unittest.yaml/badge.svg)](https://github.com/phenix3443/go-starter/actions/workflows/unittest.yaml)
[![golangci-lint](https://github.com/phenix3443/go-starter/actions/workflows/golangci-lint.yaml/badge.svg)](https://github.com/phenix3443/go-starter/actions/workflows/golangci-lint.yaml)
[![Goreleaser](https://github.com/phenix3443/go-starter/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/phenix3443/go-starter/actions/workflows/goreleaser.yml)
[![Lint PR](https://github.com/phenix3443/go-starter/actions/workflows/lint-pr.yaml/badge.svg)](https://github.com/phenix3443/go-starter/actions/workflows/lint-pr.yaml)
[![Push Docker Image](https://github.com/phenix3443/go-starter/actions/workflows/docker.yml/badge.svg)](https://github.com/phenix3443/go-starter/actions/workflows/docker.yml)

---

## 🚀 特性概览

- ✅ **DevContainer 开发环境**（VSCode 原生支持）
- ✅ **完整 CI/CD 流水线**（基于 GitHub Actions）
- ✅ **统一代码规范检查**（支持 Go、Shell、Markdown、YAML 等）
- ✅ **多语言格式化与 Lint 检查**
- ✅ **一键构建发布（Docker + Goreleaser）**
- ✅ **灵活的配置管理系统**（支持多环境、环境变量、配置文件）
- ✅ **配置验证和类型安全**

---

## 🧱 技术栈与工具链

### 🧪 单元测试与覆盖率

- 测试框架：[testify](https://github.com/stretchr/testify)
- 覆盖率报告：[go-coverage-report](https://github.com/fgrosse/go-coverage-report)

### 🧹 代码风格与 Lint（通过 Husky + lint-staged）

- **Go 语言**
  - 格式化：[goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
  - Lint：[golangci-lint](https://github.com/golangci/golangci-lint)

- **Shell 脚本**
  - ✅ 格式化：[shfmt](https://github.com/patrickvane/shfmt)
  - 🚧 Lint（未启用）：[shellcheck](https://www.shellcheck.net/)

- **Markdown**
  - 格式化：[prettier](https://prettier.io/)
  - Lint：[markdownlint](https://github.com/DavidAnson/markdownlint)

- **YAML**
  - 格式化：[prettier](https://prettier.io/)

- **提交信息规范化**
  - 使用 [commitlint](https://github.com/conventional-changelog/commitlint) + [semantic-pull-request](https://github.com/marketplace/actions/semantic-pull-request)

---

## ⚙️ 配置管理

### 📁 配置文件结构

```
app/your-app/                # 应用目录（请修改为实际的应用名称）
├── config/                  # 配置文件目录
│   ├── config.yaml.example      # 配置文件模板
│   ├── local.yaml              # 本地测试配置
│   ├── development.yaml         # 开发环境配置
│   └── production.yaml          # 生产环境配置
├── env/                     # 环境变量文件目录
│   ├── local.env               # 本地测试环境变量
│   ├── development.env          # 开发环境变量
│   └── production.env           # 生产环境变量
├── cmd/                     # 命令定义
└── main.go                  # 主程序入口
```

### 🔧 配置加载方式

1. **配置文件方式**：指定配置文件路径

   ```bash
   go run app/your-app/main.go --config app/your-app/config/local.yaml api
   ```

2. **环境变量方式**：通过环境变量覆盖配置

   ```bash
   export GO_STARTER_APP_ENVIRONMENT=production
   export GO_STARTER_MYSQL_HOST=mysql.prod.com
   go run app/your-app/main.go api
   ```

3. **命令行参数方式**：通过命令行标志设置

   ```bash
   go run app/your-app/main.go --env production --debug=false api
   ```

### ✅ 配置验证

验证配置文件的正确性：

```bash
make config-validate
# 或者
go run cmd/config/validate.go -c config/development.yaml
```

### 🆕 创建配置文件

```bash
make config-create
```

### 🆕 创建环境变量文件

```bash
make env-create
```

### 📝 环境变量文件使用

环境变量文件提供了另一种配置方式，特别适合容器化部署：

```bash
# 使用本地环境变量
source app/your-app/env/local.env
go run app/your-app/main.go api

# 使用开发环境变量
source app/your-app/env/development.env
go run app/your-app/main.go api

# 使用生产环境变量
source app/your-app/env/production.env
go run app/your-app/main.go api
```

---

## ⚙️ GitHub Actions 流水线

| 名称                 | 功能                          |
| -------------------- | ----------------------------- |
| `unittest.yaml`      | 执行测试并生成覆盖率          |
| `golangci-lint.yaml` | 代码静态检查                  |
| `goreleaser.yml`     | 自动发布版本到 GitHub Release |
| `lint-pr.yaml`       | 检查 PR 标题是否符合约定规范  |
| `docker.yml`         | 构建并推送 Docker 镜像到 GHCR |

---

## 🐳 开发环境（DevContainer）

- 预配置 `.devcontainer.json` 和 Dockerfile
- 支持 VSCode Remote Container 一键启动
- 内置 Go + 常用开发工具（make, curl, git, zsh 等）

---

## 📦 版本发布

使用 [goreleaser](https://goreleaser.com/) 实现：

- 自动构建二进制
- 自动生成 release notes
- 可选 Docker 镜像构建和上传

---

## 🔧 本地开发推荐命令

```bash
# 基础命令
make lint           # 本地 lint 检查
make test           # 运行测试
make build          # 编译项目
make release        # 本地模拟 goreleaser 发布

# 配置管理
make config-create  # 创建配置文件
make config-validate # 验证配置文件
make env-create     # 创建环境变量文件

# 开发环境
make dev-setup      # 设置开发环境
make start-dc       # 启动 Docker Compose 服务
make run-api-local  # 运行本地测试 API
make run-api-dev    # 运行开发环境 API
make run-api-prod   # 运行生产环境 API
```

---

## 🌍 环境配置示例

### 本地测试

```bash
# 使用本地配置
go run app/your-app/main.go --config app/your-app/config/local.yaml api

# 或通过环境变量
export GO_STARTER_APP_ENVIRONMENT=development
export GO_STARTER_APP_DEBUG=true
go run app/your-app/main.go api
```

### 开发环境

```bash
# 使用开发配置
   go run app/your-app/main.go --config app/your-app/config/development.yaml api

# 或通过环境变量
export GO_STARTER_APP_ENVIRONMENT=development
export GO_STARTER_APP_DEBUG=true
go run app/your-app/main.go api
```

### 生产环境

```bash
# 使用生产配置
   go run app/your-app/main.go --config app/your-app/config/production.yaml api

# 或通过环境变量
export GO_STARTER_APP_ENVIRONMENT=production
export GO_STARTER_APP_DEBUG=false
export GO_STARTER_LOG_LEVEL=info
go run app/your-app/main.go api
```

---

## 📚 配置字段说明

#### 📁 类型文件结构

```
pkg/types/
├── types.go           # 包入口文件
├── app.go             # 主配置结构
├── app_info.go        # 应用信息配置
├── database.go        # 数据库配置
├── cache.go           # 缓存配置
├── message_queue.go   # 消息队列配置
├── api.go             # API配置
└── log.go             # 日志配置
```

### 应用配置 (app)

- `name`: 应用名称
- `version`: 应用版本
- `environment`: 运行环境 (development/staging/production)
- `debug`: 是否启用调试模式

### 数据库配置 (mysql/postgres)

- `host`: 数据库主机地址
- `port`: 数据库端口
- `username`: 数据库用户名
- `password`: 数据库密码
- `database`: 数据库名称（**重要：请根据实际情况修改为您的数据库名称**）
- `charset`: 字符集（MySQL）
- `parse_time`: 是否解析时间（MySQL）
- `loc`: 时区设置（MySQL，使用 "Local" 表示本地时间）
- `timezone`: 时区设置（PostgreSQL，使用 "Local" 表示本地时间）
- `max_open_conns`: 最大连接数
- `max_idle_conns`: 最大空闲连接数
- `conn_max_lifetime`: 连接最大生命周期（秒）

### Redis 配置 (redis)

- `host`: Redis 主机地址
- `port`: Redis 端口
- `password`: Redis 密码
- `database`: Redis 数据库编号
- `pool_size`: 连接池大小
- `min_idle_conns`: 最小空闲连接数

### API 配置 (api)

- `version`: API 版本
- `address`: 监听地址
- `read_timeout`: 读取超时（秒）
- `write_timeout`: 写入超时（秒）
- `idle_timeout`: 空闲超时（秒）
- `cors`: CORS 跨域配置
- `rate_limit`: 限流配置
