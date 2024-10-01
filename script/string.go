package script

import (
    "fmt"
    "regexp"
    "strconv"
    "strings"
)

func To_i(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}

func To_s(i int) string {
    return fmt.Sprintf(`%v`, i)
}

func ToLower(s string) string {
    return strings.ToLower(s)
}

func Scan(str, keyword string) []string {
    re := regexp.MustCompile(keyword)
    return re.FindStringSubmatch(str)
}

func Match(str, keyword string) bool {
    match, _ := regexp.MatchString(keyword, str)
    return match
}

func FmtPrintln[T any](lists []T) {
    for _, list := range lists {
        fmt.Println(list)
    }
}
