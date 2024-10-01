package main

import (
    "flag"
    "fmt"

    "tools/cmd"
    "tools/lineMsg"
    "tools/script"
)

func main() {
    description := `This program use LINE Notify to send message, or upload image.`
    use := `Usage: lineMsg --token [token] --msg [message] --img [imagePath]

  -t, --token              LINE Notify API token (required).
  -m, --msg                Message to send (required).
  -i, --img                Path to the image file (optional).`

    appVersion := `1.0`
    argsCount := 0

    cmd.Set(use, description)

    var token string
    flag.StringVar(&token, "t", "", "token")
    flag.StringVar(&token, "token", "", "token")

    var message string
    flag.StringVar(&message, "m", "", "message")
    flag.StringVar(&message, "msg", "", "message")

    var imagePath string
    flag.StringVar(&imagePath, "i", "", "imagePath")
    flag.StringVar(&imagePath, "img", "", "imagePath")

    cmd.SetArgs(argsCount, appVersion)

    // required token msg
    if token == "" || message == "" {
        cmd.ShowHelp()
    }

    if imagePath != "" {
        script.ShowFileExist(imagePath)
    }

    err := lineMsg.Send(token, message, imagePath)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("ok")
    }
}

var _ = fmt.Println
