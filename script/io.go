package script

import (
    "io/ioutil"
    "strings"
)

func ReadFile(path string) string {
    bytes, err := ioutil.ReadFile(path)
    printError(err)
    return string(bytes)
}

func WriteFile(content, path string) {
    err := ioutil.WriteFile(path, []byte(content), 0777)
    printError(err)
}

func ReadLines(path string) []string {
    bytes, err := ioutil.ReadFile(path)
    printError(err)
    var contents []string
    for _, line := range strings.Split(string(bytes), "\n") {
        contents = append(contents, strings.TrimSpace(line))
    }
    return contents
}

func WriteLines(lines []string, path string) {
    bytes := []byte(strings.Join(lines, "\n"))
    ioutil.WriteFile(path, bytes, 0777)
}
