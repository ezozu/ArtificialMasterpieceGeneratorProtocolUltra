// cmd/artificialmasterpiecegeneratorprotocolultra/main.go
package main

import (
"flag"
"log"
"os"

"artificialmasterpiecegeneratorprotocolultra/internal/artificialmasterpiecegeneratorprotocolultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialmasterpiecegeneratorprotocolultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
