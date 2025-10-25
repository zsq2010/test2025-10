package imagesplit

import (
    "errors"
    "image"
    "image/color"
    "image/jpeg"
    "image/png"
    "os"
    "path/filepath"
    "strings"
    "testing"
)

func ensureTestImages(t *testing.T) (string, string) {
    t.Helper()
    pngPath := filepath.Join("testdata", "gradient.png")
    if _, err := os.Stat(pngPath); err != nil {
        t.Fatalf("expected test image %s: %v", pngPath, err)
    }

    jpegPath := filepath.Join("testdata", "blocks.jpg")
    if _, err := os.Stat(jpegPath); errors.Is(err, os.ErrNotExist) {
        createBlocksJPEG(t, jpegPath)
    } else if err != nil {
        t.Fatalf("stat jpeg: %v", err)
    }

    return pngPath, jpegPath
}

func createBlocksJPEG(t *testing.T, path string) {
    t.Helper()
    if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
        t.Fatalf("mkdir testdata: %v", err)
    }

    img := image.NewRGBA(image.Rect(0, 0, 12, 8))
    for y := 0; y < 8; y++ {
        for x := 0; x < 12; x++ {
            r := uint8(float64(x) / 11.0 * 255)
            g := uint8(float64(y) / 7.0 * 255)
            img.Set(x, y, color.RGBA{R: r, G: g, B: 180, A: 255})
        }
    }

    file, err := os.Create(path)
    if err != nil {
        t.Fatalf("create jpeg: %v", err)
    }
    defer file.Close()

    if err := jpeg.Encode(file, img, &jpeg.Options{Quality: 85}); err != nil {
        t.Fatalf("encode jpeg: %v", err)
    }

    t.Cleanup(func() {
        _ = os.Remove(path)
    })
}

func TestGridSplitPNG(t *testing.T) {
    pngPath, _ := ensureTestImages(t)
    outDir := t.TempDir()

    files, err := GridSplit(pngPath, 3, 4, SplitOptions{OutputDir: outDir})
    if err != nil {
        t.Fatalf("GridSplit returned error: %v", err)
    }
    if len(files) != 12 {
        t.Fatalf("expected 12 tiles, got %d", len(files))
    }

    expectedHeights := []int{4, 3, 3}
    expectedWidths := []int{3, 3, 2, 2}

    for i, path := range files {
        if _, err := os.Stat(path); err != nil {
            t.Fatalf("expected file %s: %v", path, err)
        }

        f, err := os.Open(path)
        if err != nil {
            t.Fatalf("open tile: %v", err)
        }
        img, err := png.Decode(f)
        f.Close()
        if err != nil {
            t.Fatalf("decode tile png: %v", err)
        }
        bounds := img.Bounds()
        row := i / 4
        col := i % 4
        if bounds.Dx() != expectedWidths[col] || bounds.Dy() != expectedHeights[row] {
            t.Errorf("unexpected tile size for row %d col %d: got %dx%d", row, col, bounds.Dx(), bounds.Dy())
        }
    }
}

func TestGridSplitJPEGInput(t *testing.T) {
    _, jpegPath := ensureTestImages(t)
    outDir := t.TempDir()

    files, err := GridSplit(jpegPath, 2, 3, SplitOptions{OutputDir: outDir})
    if err != nil {
        t.Fatalf("GridSplit returned error: %v", err)
    }
    if len(files) != 6 {
        t.Fatalf("expected 6 tiles, got %d", len(files))
    }

    for _, path := range files {
        if filepath.Ext(path) != ".jpg" {
            t.Errorf("expected jpg extension, got %s", filepath.Ext(path))
        }
        f, err := os.Open(path)
        if err != nil {
            t.Fatalf("open tile: %v", err)
        }
        if _, err := jpeg.Decode(f); err != nil {
            t.Fatalf("decode jpeg tile: %v", err)
        }
        f.Close()
    }
}

func TestTileSplitRemainders(t *testing.T) {
    pngPath, _ := ensureTestImages(t)
    outDir := t.TempDir()

    files, err := TileSplit(pngPath, 4, 3, SplitOptions{OutputDir: outDir, FilePrefix: "tiles"})
    if err != nil {
        t.Fatalf("TileSplit returned error: %v", err)
    }
    if len(files) != 12 {
        t.Fatalf("expected 12 tiles, got %d", len(files))
    }

    last := files[len(files)-1]
    f, err := os.Open(last)
    if err != nil {
        t.Fatalf("open last tile: %v", err)
    }
    img, err := png.Decode(f)
    f.Close()
    if err != nil {
        t.Fatalf("decode last tile: %v", err)
    }
    bounds := img.Bounds()
    if bounds.Dx() != 2 || bounds.Dy() != 1 {
        t.Errorf("expected last tile size 2x1, got %dx%d", bounds.Dx(), bounds.Dy())
    }

    for i, path := range files {
        if !strings.Contains(filepath.Base(path), "tiles_tile_") {
            t.Errorf("unexpected filename for tile %d: %s", i, path)
        }
    }
}

func TestTileSplitJPEGOutput(t *testing.T) {
    pngPath, _ := ensureTestImages(t)
    outDir := t.TempDir()

    files, err := TileSplit(pngPath, 6, 5, SplitOptions{OutputDir: outDir, Format: "jpeg", Quality: 75})
    if err != nil {
        t.Fatalf("TileSplit returned error: %v", err)
    }
    if len(files) != 4 {
        t.Fatalf("expected 4 tiles, got %d", len(files))
    }

    for _, path := range files {
        if filepath.Ext(path) != ".jpg" {
            t.Errorf("expected jpg extension, got %s", filepath.Ext(path))
        }
        f, err := os.Open(path)
        if err != nil {
            t.Fatalf("open tile: %v", err)
        }
        if _, err := jpeg.Decode(f); err != nil {
            t.Fatalf("decode jpeg tile: %v", err)
        }
        f.Close()
    }
}

func TestGridSplitInvalidParameters(t *testing.T) {
    pngPath, _ := ensureTestImages(t)
    if _, err := GridSplit(pngPath, 0, 2, SplitOptions{}); err == nil {
        t.Fatalf("expected error for zero rows")
    }
    if _, err := GridSplit(pngPath, 2, -1, SplitOptions{}); err == nil {
        t.Fatalf("expected error for negative cols")
    }
}

func TestTileSplitInvalidParameters(t *testing.T) {
    pngPath, _ := ensureTestImages(t)
    if _, err := TileSplit(pngPath, 0, 5, SplitOptions{}); err == nil {
        t.Fatalf("expected error for zero width")
    }
    if _, err := TileSplit(pngPath, 5, 0, SplitOptions{}); err == nil {
        t.Fatalf("expected error for zero height")
    }
}

func TestUnsupportedOutputFormat(t *testing.T) {
    pngPath, _ := ensureTestImages(t)
    _, err := GridSplit(pngPath, 2, 2, SplitOptions{Format: "gif"})
    if err == nil {
        t.Fatalf("expected error for unsupported output format")
    }
}

func TestUnsupportedInputFormat(t *testing.T) {
    tmp := filepath.Join(t.TempDir(), "not_image.txt")
    if err := os.WriteFile(tmp, []byte("not an image"), 0o644); err != nil {
        t.Fatalf("write temp file: %v", err)
    }

    if _, err := GridSplit(tmp, 2, 2, SplitOptions{}); err == nil {
        t.Fatalf("expected error for invalid input format")
    }
}

func TestOutputDirectoryCreated(t *testing.T) {
    pngPath, _ := ensureTestImages(t)
    base := t.TempDir()
    outDir := filepath.Join(base, "nested", "dir")

    files, err := TileSplit(pngPath, 5, 5, SplitOptions{OutputDir: outDir})
    if err != nil {
        t.Fatalf("TileSplit returned error: %v", err)
    }
    if len(files) == 0 {
        t.Fatalf("expected at least one output file")
    }

    if _, err := os.Stat(outDir); err != nil {
        t.Fatalf("expected output directory to be created: %v", err)
    }
}
