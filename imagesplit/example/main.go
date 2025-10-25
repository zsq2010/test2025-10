package main

import (
    "fmt"
    "log"
    "path/filepath"

    "github.com/cto-new/imagesplit/imagesplit"
)

func main() {
    inputImage := filepath.Join("..", "testdata", "gradient.png")
    outputDir := filepath.Join("output")

    gridFiles, err := imagesplit.GridSplit(inputImage, 2, 3, imagesplit.SplitOptions{
        OutputDir:  outputDir,
        FilePrefix: "sample_grid",
    })
    if err != nil {
        log.Fatalf("grid split failed: %v", err)
    }
    fmt.Println("Grid split results:")
    for _, path := range gridFiles {
        fmt.Println(" -", path)
    }

    tileFiles, err := imagesplit.TileSplit(inputImage, 4, 4, imagesplit.SplitOptions{
        OutputDir:  outputDir,
        FilePrefix: "sample_tile",
        Format:     "jpeg",
        Quality:    80,
    })
    if err != nil {
        log.Fatalf("tile split failed: %v", err)
    }
    fmt.Println("\nTile split results:")
    for _, path := range tileFiles {
        fmt.Println(" -", path)
    }
}
