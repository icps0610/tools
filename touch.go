package main

import (
    "tools/cmd"
    "tools/script"
)

func main() {
    description := `This program updates the timestamp of a specified file or creates an empty file if it does not exist.`
    use := `Usage: touch [filePath]`
    appVersion := `1.0`
    argsCount := 1

    cmd.Set(use, description)
    args := cmd.SetArgs(argsCount, appVersion)

    filePath := args[0]
    if script.FileExist(filePath) {
        script.RefreshFileTime(filePath)
    } else {
        script.CreateBlankFile(filePath)
    }

}
