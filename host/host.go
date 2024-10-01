package host

import (
    "fmt"
    "net"
)

func GetName(ip string) []string {
    names, err := net.LookupAddr(ip)
    printError(err)
    return names
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
