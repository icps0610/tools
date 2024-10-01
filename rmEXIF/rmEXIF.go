package rmEXIF

import (
    "fmt"
    "image"
    "image/jpeg"
    "image/png"
    "os"
    "path/filepath"
)

func Run(inputPath, outputPath string) {
    ext := filepath.Ext(inputPath)

    // 開啟檔案 另存檔名
    ioReader, _ := os.Open(inputPath)
    img, _, err := image.Decode(ioReader)
    printError(err)

    file, err := os.Create(outputPath)
    printError(err)
    defer file.Close()

    if ext == `.png` {
        err = png.Encode(file, img)
    } else if ext == `.jpeg` || ext == `.jpg` {
        err = jpeg.Encode(file, img, nil)
    }

    printError(err)
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
