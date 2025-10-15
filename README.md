# leek

<!-- [![Release Version](https://img.shields.io/github/v/release/shalldie/leek?display_name=tag&logo=github&style=flat-square)](https://github.com/shalldie/leek)
[![Docker Image Version](https://img.shields.io/docker/v/shalldie/leek/latest?style=flat-square&logo=docker)](https://hub.docker.com/r/shalldie/leek/tags)
[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/leek?label=go&logo=go&style=flat-square)](https://github.com/shalldie/leek)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/leek.svg)](https://pkg.go.dev/github.com/shalldie/leek)
[![Build Status](https://img.shields.io/github/actions/workflow/status/shalldie/leek/ci.yml?logo=github&style=flat-square)](https://github.com/shalldie/leek/actions)
[![License](https://img.shields.io/github/license/shalldie/leek?logo=github&style=flat-square)](https://github.com/shalldie/leek) -->

查询当前 `大A` 行情的 cli 工具。

## Installation

1. binary

[Download](https://github.com/shalldie/leek/releases)，下载后直接执行即可，加入 `PATH` 更佳。

| 文件                | 适用系统                 |
| :------------------ | :----------------------- |
| `leek.darwin-amd64` | `Mac amd64`、`Mac arm64` |
| `leek.linux-amd64`  | `Linux amd64`            |
| `leek.linux-arm64`  | `Linux arm64`            |

example:

```bash
# install
wget -O leek [url]
sudo chmod a+x leek
sudo mv leek /usr/local/bin/leek
# run
leek
```

2. go install

```bash
# install
go install github.com/shalldie/leek@latest
# run
leek
```

## LICENSE

MIT
