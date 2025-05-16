# go-starter

`go-starter` 是一个为 Golang 项目开发设计的基础模板，内置 VSCode Dev Container 支持、GitHub Actions CI/CD、统一代码规范校验、自动发布构建等常见开发工具链，开箱即用。

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
make lint       # 本地 lint 检查
make test       # 运行测试
make build      # 编译项目
make release    # 本地模拟 goreleaser 发布
```
