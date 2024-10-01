package script

import (
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "strings"
    "time"
)

var IsWin bool

func init() {
    IsWin = runtime.GOOS == "windows"
}

func FileExist(file string) bool {
    if _, err := os.Stat(file); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func ShowFileExist(files ...string) bool {
    for _, filePath := range files {
        if !FileExist(filePath) {
            fmt.Println(fmt.Sprintf(`[ %s ] file not exist !!!`, filePath))
            return false
        }
    }
    return true
}

func GetFileName(localPath string) string {
    return filepath.Base(localPath)
}

func GetFullPath(dirPath, fileName string) string {
    filepath := filepath.Join(dirPath, fileName)
    return strings.Replace(filepath, `\`, `/`, -1)
}

func CreateBlankFile(filePath string) {
    err := ioutil.WriteFile(filePath, []byte(""), 0777)
    printError(err)
}

func RefreshFileTime(filePath string) {
    now := time.Now()
    err := os.Chtimes(filePath, now, now)
    printError(err)
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}

func RunCmd(cmd string) string {
    var bash, args = `bash`, `-c`
    if IsWin {
        bash, args = `cmd`, `/c`
    }
    output, _ := exec.Command(bash, args, cmd).CombinedOutput()
    return string(output)
}

func RunPS(cmd string) {
    _, err := exec.Command("pwsh", "-c", cmd).CombinedOutput()
    printError(err)
}

func PrintError(err error, msgs ...string) {
    if err != nil {
        for _, msg := range msgs {
            fmt.Println(msg)
        }
        os.Exit(1)
    }
}
