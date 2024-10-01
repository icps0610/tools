package main

import (
    "tools/cmd"
    "tools/rmEXIF"
    "tools/script"
)

func main() {
    description := `This program removes EXIF data from image by reconstructing them.`
    use := `Usage: rmEXIF [inputPath] [outputPath]`
    appVersion := `1.0`
    argsCount := 2

    cmd.Set(use, description)
    args := cmd.SetArgs(argsCount, appVersion)

    inputPath, outputPath := args[0], args[1]
    // 檢查是否存在
    if !script.ShowFileExist(inputPath) {
        return
    }

    rmEXIF.Run(inputPath, outputPath)
}
