package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/cto-new/imagesplit/imagesplit"
    testdata "github.com/cto-new/imagesplit/imagesplit/testdata"
)

func main() {
    outputDir := filepath.Join("output")
    if err := os.MkdirAll(outputDir, 0o755); err != nil {
        log.Fatalf("create output directory: %v", err)
    }

    inputImage := filepath.Join(outputDir, "sample_input.png")
    if err := testdata.WriteGradientPNG(inputImage); err != nil {
        log.Fatalf("prepare input image: %v", err)
    }

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

    batchInput := filepath.Join(outputDir, "batch_input")
    if err := os.MkdirAll(batchInput, 0o755); err != nil {
        log.Fatalf("create batch input directory: %v", err)
    }
    if err := testdata.WriteGradientPNG(filepath.Join(batchInput, "gradient.png")); err != nil {
        log.Fatalf("write gradient sample: %v", err)
    }
    if err := testdata.WriteBlocksJPEG(filepath.Join(batchInput, "blocks.jpg")); err != nil {
        log.Fatalf("write blocks sample: %v", err)
    }

    batchOutput := filepath.Join(outputDir, "batch")
    batchResults, err := imagesplit.SplitDirectory(batchInput, batchOutput, imagesplit.DirectorySplitConfig{
        Mode: imagesplit.DirectorySplitModeGrid,
        Rows: 2,
        Cols: 2,
    })
    if err != nil {
        log.Fatalf("batch split failed: %v", err)
    }

    fmt.Println("\nDirectory split results:")
    for src, generated := range batchResults {
        fmt.Println("Source:", src)
        for _, path := range generated {
            fmt.Println("  -", path)
        }
    }
}
