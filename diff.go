package main

import (
    "flag"
    "fmt"

    "tools/cmd"
    "tools/diff"
    "tools/script"
)

func main() {
    description := `This program compares two files for differences.You can specify the comparison 
mode using the --mode parameter. If you do not specify the --mode parameter, the system will
automatically use the default value diff. Available mode options include:

  diff:   Shows the differences between the two files (default mode).
  same:   Returns results only when the files are identical.
  all:    Displays all content, including differences and similarities.`
    use := `Usage: diff --mode [*diff/same/all] --output [outputFullPath] [filePath1] [filePath2]

  -m, --mode               diff/same/all default: diff
  -o, --output             Generate an HTML report of the comparison results.`
    appVersion := `1.0`
    argsCount := 2

    cmd.Set(use, description)
    var mode string
    flag.StringVar(&mode, "m", "diff", "mode")
    flag.StringVar(&mode, "mode", "diff", "mode")

    var outputPath string
    flag.StringVar(&outputPath, "o", "", "Generate an HTML report of the comparison results.")
    flag.StringVar(&outputPath, "output", "", "Generate an HTML report of the comparison results.")

    args := cmd.SetArgs(argsCount, appVersion)

    filePath1, filePath2 := args[0], args[1]
    // 檢查是否存在
    if !script.ShowFileExist(filePath1, filePath2) {
        return
    }

    content1, content2 := script.ReadLines(filePath1), script.ReadLines(filePath2)
    difference, same, all := diff.Run(content1, content2)

    fmt.Printf(`Mode: `)

    switch mode {
    case "same":
        fmt.Println(mode)
        diff.Print(same)
    case "all":
        fmt.Println(mode)
        diff.Print(all)
    default:
        fmt.Println(`diff`)
        diff.Print(difference)
    }

    if outputPath != "" {
        fmt.Println()
        fmt.Println("Generate an HTML report", outputPath)

        content := diff.OutputHtml(filePath1, filePath2, difference, same, all)
        script.WriteFile(content, outputPath)
    }
}
