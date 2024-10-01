package main

import (
    "fmt"

    "tools/cmd"
    "tools/script"
)

func main() {
    description := `This program uses PowerShell to generate speech.
Add-Type -AssemblyName System.Speech; (New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak([text]).`
    use := `Usage: speech [text]`
    appVersion := `1.0`
    argsCount := 1

    cmd.Set(use, description)
    args := cmd.SetArgs(argsCount, appVersion)

    cmdStr := fmt.Sprintf(`Add-Type -AssemblyName System.Speech; (New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak('%s');`, args[0])

    script.RunPS(cmdStr)
}
