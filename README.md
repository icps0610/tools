# tools

``` cmd
for /f "delims=" %%a in ('now.exe -d') do set "date=%%a"

echo %date%
```

| network  | Description                                                                       |
|----------|-----------------------------------------------------------------------------------|
| `host`   | This program shows the target hostname.                                           |
| `port`   | This program checks the responsiveness of a specified IP address and port number. |


|  system  | Description                                                                                           |
|----------|-------------------------------------------------------------------------------------------------------|
| `diff`   | This program compares two files for differences.                                                      |
| `now`    | This program displays the current time or date.                                                       |
| `passwd` | This program generates random password.                                                               |
| `sleep`  | This program pauses execution for a specified number of seconds.                                      |
| `speech` | This program uses PowerShell to generate speech.                                                      |
| `touch`  | This program updates the timestamp of a specified file or creates an empty file if it does not exist. |

|  image   | Description                                                        |
|----------|--------------------------------------------------------------------|
| `rmEXIF` | This program removes EXIF data from image by reconstructing them.  |

|   app    | Description                                                        |
|----------|--------------------------------------------------------------------|
| `lineMsg`| This program use LINE Notify to send message, or upload image.     |