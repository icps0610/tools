package main

import (
    "flag"
    "fmt"
    "os"

    "tools/cmd"
    "tools/port"
    "tools/script"
)

func main() {
    description := `This program checks the responsiveness of a specified IP address and port number.`
    use := `Usage: port --ip [ip] --port [port] --wait [millisecond]

Example: 
    port --ip 192.168.0.1-5
    port --ip 192.168.0.1     --port 8080
    port --ip 192.168.0.1     --port 8000-8080
    port --ip 192.168.0.1-100 --port 80,8100,8200
    port --ip 192.168.0.1-100 --port common

  --ip                    Accepts one of the following formats:
                            1. A single IP. (e.g., 192.168.0.1).
                            2. A range of consecutive IPs (e.g., 192.168.0.1-100).
  --port                  Accepts one of the following formats:
                            1. If --port is omitted, only ping results will be displayed.
                            2. A single port number (e.g., 8080).
                            3. A range of consecutive ports (e.g., 8000-8080).
                            4. Multiple ports separated by commas (e.g., 80, 8100, 8200).
                            5. Common ports can be specified using the keyword common, which includes:
                               21, 22, 23, 25, 53, 67, 68, 80, 110, 143, 443, 445, 3389.
  --wait                  Specifies the response wait time. Default: 300 milliseconds.`
    appVersion := `1.0`
    argsCount := 0

    cmd.Set(use, description)

    var ip, portStr string
    flag.StringVar(&ip, "ip", "", "ip")
    flag.StringVar(&portStr, "port", "", "port")

    var wait int
    flag.IntVar(&wait, "wait", 300, "wait")

    cmd.SetArgs(argsCount, appVersion)

    // 整理 ip
    lan, ipStart, ipEnd := port.GetIP(ip)
    if lan == "" {
        flag.Usage()
        os.Exit(0)
    }

    // 整理 ports
    var ports []string
    if portStr == "common" {
        ports = []string{"21", "22", "23", "25", "53", "67", "68", "80", "110", "143", "443", "445", "3389"}
    } else {
        nports := port.GetPorts(portStr)
        if len(nports) > 0 {
            ports = nports
        }
    }

    var flag bool
    // main

    for _, ip := range port.GetRange(ipStart, ipEnd) {
        lanIp := fmt.Sprintf(`%s.%v`, lan, ip)
        response := script.Ping(lanIp, wait)

        if response {
            fmt.Println(lanIp, "O")
            flag = true
            for _, portNumber := range ports {
                fmt.Printf(fmt.Sprintf(`%s %s`, lanIp, portNumber))

                err := port.CheckPort(lanIp, portNumber, wait)

                if err == nil {
                    fmt.Printf(" O\n")
                    flag = true
                } else {
                    fmt.Printf(" X\n")
                    flag = false
                }
            }
        } else {
            fmt.Println(lanIp, "X")
            flag = false
        }
    }
    if flag {
        os.Exit(0)
    } else {
        os.Exit(1)
    }
}
