package main

import (
    "flag"
    "fmt"

    "tools/cmd"
    "tools/script"
)

func main() {
    description := `This program displays the current time or date. If no options are specified, the default output 
is the current month, day, hour, and minute in the format MMDDHHMM (e.g., 06100030)`
    use := `Usage: now [options]

  -f, --full         Display the full timestamp (e.g., 2024-01-01-00:00:00). Optional.
  -d, --day          Display the current date (e.g., 20240101). Optional.

  -h, --help         Show this help message and exit.
  -v, --version      Display version information and exit.

`
    appVersion := `1.0`
    argsCount := 0

    cmd.Set(use, description)

    var full bool
    flag.BoolVar(&full, "f", false, "Display the full timestamp")
    flag.BoolVar(&full, "full", false, "Display the full timestamp")

    var day bool
    flag.BoolVar(&day, "d", false, "Display the current date")
    flag.BoolVar(&day, "day", false, "Display the current date")

    cmd.SetArgs(argsCount, appVersion)

    now := script.TimeNow()

    var msg string
    if full {
        msg = fmt.Sprintf(`%v-%02v-%02v-%02v:%02v:%02v`, now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
    } else if day {
        msg = fmt.Sprintf(`%v%02v%02v`, now.Year(), int(now.Month()), now.Day())
    } else {
        msg = fmt.Sprintf(`%02v%02v%02v%02v`, int(now.Month()), now.Day(), now.Hour(), now.Minute())
    }

    fmt.Println(msg)
}
