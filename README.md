# imagesplit

`imagesplit` 是一个用于 Go 语言的图片分割库，支持按网格（行×列）和固定尺寸两种方式对 PNG / JPEG 图片进行分割。库提供简洁的 API，可以作为独立依赖被其他项目调用。

## 功能特性

- ✅ 支持 PNG (`.png`) 和 JPEG (`.jpg`, `.jpeg`) 格式
- ✅ 网格分割：按照指定的行列数自动生成小图块
- ✅ 固定尺寸分割：按照固定的宽高切割，自动处理边缘剩余区域
- ✅ 灵活的输出配置：输出目录、文件前缀、图片格式、JPEG 质量
- ✅ 完善的错误处理：格式不支持、参数错误、输出目录创建失败等

## 安装

```bash
go get github.com/cto-new/imagesplit
```

## 快速上手

```go
package main

import (
    "fmt"
    "log"

    "github.com/cto-new/imagesplit/imagesplit"
)

func main() {
    files, err := imagesplit.GridSplit("input.png", 3, 4, imagesplit.SplitOptions{
        OutputDir:  "output",
        FilePrefix: "sample",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(files)
}
```

## API 文档

```go
func GridSplit(inputPath string, rows, cols int, opts imagesplit.SplitOptions) ([]string, error)
```
- `rows` / `cols`: 网格行列数，必须大于 0。
- `SplitOptions`
  - `OutputDir`: 输出目录（为空时使用原图所在目录）。
  - `FilePrefix`: 输出文件前缀（为空时使用原图文件名）。
  - `Format`: 输出格式（`"png"`、`"jpeg"`，为空使用原图格式）。
  - `Quality`: JPEG 质量，范围 1-100（默认 90）。
- 返回值为生成的文件路径列表。

```go
func TileSplit(inputPath string, tileWidth, tileHeight int, opts imagesplit.SplitOptions) ([]string, error)
```
- `tileWidth` / `tileHeight`: 图块宽高，必须大于 0。
- 其余参数与 `GridSplit` 一致。

## 命名规则

- 网格分割：`{prefix}_row{i}_col{j}.{ext}` → 例如：`image_row0_col2.png`
- 固定尺寸：`{prefix}_tile_{index}.{ext}` → 例如：`image_tile_5.jpg`

## 示例

仓库包含一个完整的示例程序，位于 `imagesplit/example/main.go`：

```bash
cd imagesplit/example
go run .
```

示例将按照网格和固定尺寸两种方式分割 `imagesplit/testdata/gradient.png` 图片，并把结果输出到 `imagesplit/example/output` 目录。

## 测试

项目提供了覆盖主要功能和边界情况的单元测试，包括：

- 网格分割及无法整除时的边缘尺寸验证
- 固定尺寸分割及剩余像素处理
- PNG / JPEG 输入输出
- 错误处理与参数校验
- 输出目录自动创建

运行测试：

```bash
cd imagesplit
go test ./...
```

## 许可证

该项目基于 MIT License 发布，欢迎自由使用与贡献。
