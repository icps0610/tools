package main

import (
    "fmt"

    "tools/cmd"
    "tools/script"
)

func main() {
    description := `This program pauses execution for a specified number of seconds.`
    use := `Usage: sleep [number]`
    appVersion := `1.0`
    argsCount := 1

    cmd.Set(use, description)
    args := cmd.SetArgs(argsCount, appVersion)

    seconds := script.To_i(args[0]) * 1000
    fmt.Println(seconds, `sec`)

    script.Sleep(seconds)
}
