# go-starter

golang starter 是一个适用于 golang 开发环境的配置模板。

## Features

- [x] Develop in [vscode dev container](https://code.visualstudio.com/docs/devcontainers/containers).
- [x] Enable Github Actions for:
  - [x] sync template with action.
  - [x] push docker image to ghcr
  - [x] lint PR by [semantic-pull-request](https://github.com/marketplace/actions/semantic-pull-request)
  - [x] lint code by [golangci-lint](https://github.com/golangci/golangci-lint)
  - [x] unittest and coverage report by [go-coverage-report](https://github.com/fgrosse/go-coverage-report)
  - [x] generate release by [goreleaser](https://goreleaser.com/)
- [x] Test with [testify](https://github.com/stretchr/testify).
- [x] Add more checks by [Husky](https://typicode.github.io/husky/).
  - [x] Lint commit by [commitlint](https://github.com/conventional-changelog/commitlint).
  - [x] Lint and format code by [Lint staged](https://github.com/lint-staged/lint-staged).
    - [x] shell script
      - [x] format by [shfmt](https://github.com/patrickvane/shfmt).
      - [ ] lint by [shellcheck](https://www.shellcheck.net/)
    - [x] markdown
      - [x] format by [prettier](https://prettier.io/).
      - [x] lint by [markdownlint](https://github.com/DavidAnson/markdownlint).
    - [x] yaml
      - [x] format by [prettier](https://prettier.io/).'
  - [x] golang
    - [x] format by [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports).
    - [x] lint by [golangci-lint](https://github.com/golangci/golangci-lint)
