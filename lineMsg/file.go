package lineMsg

import (
    "bytes"
    "fmt"
    "image/jpeg"
    "image/png"
    "io"
    "os"
    "path/filepath"
)

func checkImageFormat(imgPath string) (io.Reader, string) {
    f, err := os.Open(imgPath)
    printError(err)
    defer f.Close()

    ext := filepath.Ext(imgPath)
    var buffer bytes.Buffer
    switch ext {
    case ".jpeg", ".jpg", ".JPEG", ".JPG":
        ext = "jpeg"
        img, err := jpeg.Decode(f)
        printError(err)
        err = jpeg.Encode(&buffer, img, &jpeg.Options{Quality: 100})
        printError(err)
    case ".png", ".PNG":
        ext = "png"
        img, err := png.Decode(f)
        printError(err)
        err = png.Encode(&buffer, img)
        printError(err)
    default:
        fmt.Println("imagePath not image !")
        os.Exit(1)
    }

    return &buffer, ext
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
