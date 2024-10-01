package main

import (
    "flag"
    "fmt"
    "math/rand"
    "time"

    "tools/cmd"
)

var defaultChar = `23456789ABCDEFGHJKMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz!@#$%^&*`

func main() {
    description := fmt.Sprintf(`This program generates random password. default characters:
%s`, defaultChar)
    use := `Usage: passwd --char [words] --len [number]

  -c, --char               Customize the character set used for generating passwords.
  -l, --len                Specify the length of the generated password.`
    appVersion := `1.0`
    argsCount := 0

    cmd.Set(use, description)

    var char string
    flag.StringVar(&char, "c", defaultChar, "character")
    flag.StringVar(&char, "char", defaultChar, "character")

    var length int
    flag.IntVar(&length, "l", 10, "length")
    flag.IntVar(&length, "len", 10, "length")

    cmd.SetArgs(argsCount, appVersion)

    count := len(char)

    var result string
    for len(result) < length {
        result += string(char[rand.Intn(count)])
    }
    fmt.Println(result)

}

func init() {
    rand.Seed(time.Now().UnixNano())
}
