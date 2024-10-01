package script

import (
    "fmt"
    "net"
    "time"

    "golang.org/x/net/icmp"
    "golang.org/x/net/ipv4"
)

// Ping function to check if the target responds to ICMP Echo
func Ping(address string, wait int) bool {
    conn, err := net.Dial("ip4:icmp", address)
    if err != nil {
        fmt.Println("Error:", err)
        return false
    }
    defer conn.Close()

    // Construct ICMP Echo Request message
    icmpMessage := icmp.Message{
        Type: ipv4.ICMPTypeEcho,
        Code: 0,
        Body: &icmp.Echo{
            ID:   1,
            Seq:  1,
            Data: []byte("PING"),
        },
    }

    // Serialize the message
    messageBytes, err := icmpMessage.Marshal(nil)
    if err != nil {
        fmt.Println("Error:", err)
        return false
    }

    // Send ICMP Echo Request
    _, err = conn.Write(messageBytes)
    if err != nil {
        fmt.Println("Error:", err)
        return false
    }

    reply := make([]byte, 1500)
    conn.SetReadDeadline(time.Now().Add(time.Duration(wait) * time.Millisecond))

    _, err = conn.Read(reply)
    if err != nil {
        return false // Timeout or error, assume no response
    }

    return true // Received a response
}
