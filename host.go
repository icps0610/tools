package main

import (
    "fmt"

    "tools/cmd"
    "tools/host"
)

func main() {
    description := `This program shows the target hostname.`
    use := `Usage: host [ip]`
    appVersion := `1.0`
    argsCount := 1

    cmd.Set(use, description)
    args := cmd.SetArgs(argsCount, appVersion)

    for _, name := range host.GetName(args[0]) {
        fmt.Println(name)
    }
}
