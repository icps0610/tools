package lineMsg

import (
    "bytes"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "net/textproto"
    "os"
)

func Send(token, msg, imgPath string) error {
    lineUrl := "https://notify-api.line.me/api/notify"

    buffer, contentType := append(msg, imgPath)

    req, err := http.NewRequest("POST", lineUrl, buffer)
    printError(err)

    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", contentType)

    resp, err := http.DefaultClient.Do(req)
    if resp.StatusCode != 200 {
        fmt.Println(resp.Status)
        os.Exit(1)
    }
    return err
}

func append(msg, imgPath string) (io.Reader, string) {
    var buffer bytes.Buffer
    writer := multipart.NewWriter(&buffer)

    fw, err := writer.CreateFormField("message")
    printError(err)
    // text
    _, err = fw.Write([]byte(msg))
    printError(err)

    //image
    if imgPath != "" {
        part := make(textproto.MIMEHeader)
        part.Set("Content-Disposition", `form-data; name="imageFile"; filename=`+imgPath)

        img, format := checkImageFormat(imgPath)
        switch format {
        case "jpeg":
            part.Set("Content-Type", "image/jpeg")
        case "png":
            part.Set("Content-Type", "image/png")
        }

        fw, err = writer.CreatePart(part)
        printError(err)

        io.Copy(fw, img)
    }

    writer.Close()
    contentType := writer.FormDataContentType()

    return &buffer, contentType
}
