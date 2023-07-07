package main

import (
    app "app/bootstrap"
)

func main() {
    app.Server().Run(":8080")
}
