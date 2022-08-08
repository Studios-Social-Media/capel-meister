package logs

import (
    "fmt"
    "os"
)

func Info (msg string) {
    fmt.Println(">> " + msg);
}

func Error (err error) {
    if err == nil {
        return
    }

    fmt.Printf("\x1b[31;1m>> %s\x1b[0m\n", fmt.Sprintf("error: %s", err))
    os.Exit(1);
}
