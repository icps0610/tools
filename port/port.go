package port

import (
    "fmt"
    "net"
    "strings"
    "time"

    "tools/script"
)

func GetIP(str string) (string, int, int) {
    match := script.Scan(str, `(\d+\.\d+\.\d+)\.(\d+)-(\d+)`)
    var lan string
    var ipStart, ipEnd int
    if len(match) > 3 {
        lan = match[1]
        ipStart = script.To_i(match[2])
        ipEnd = script.To_i(match[3])
    } else {
        match = script.Scan(str, `(\d+\.\d+\.\d+)\.(\d+)`)
        if len(match) > 2 {
            lan = match[1]
            ipStart = script.To_i(match[2])
            ipEnd = script.To_i(match[2])

        }
    }

    return lan, ipStart, ipEnd
}

func GetPorts(str string) []string {
    match := script.Scan(str, `(\d+)-(\d+)`)

    var ports []string
    var portStart, portEnd int
    if len(match) > 2 {
        portStart = script.To_i(match[1])
        portEnd = script.To_i(match[2])

        for i := portStart; i <= portEnd; i++ {
            ports = append(ports, script.To_s(i))
        }
    } else if script.Match(str, ",") {
        for _, port := range strings.Split(str, ",") {
            ports = append(ports, strings.TrimSpace(port))
        }
    }
    // number
    match = script.Scan(str, `(\d+)`)
    if len(match) > 0 {
        ports = append(ports, str)
    }

    return ports
}

func GetRange(min, max int) []int {
    var list []int
    for idx := min; idx <= max; idx++ {
        list = append(list, idx)
    }
    return list
}

func CheckPort(ip, port string, wait int) error {
    address := fmt.Sprintf("%s:%s", ip, port)
    _, err := net.DialTimeout("tcp", address, time.Duration(wait)*time.Millisecond)
    return err
}
