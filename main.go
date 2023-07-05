package main

import "app/config"

func main() {
    r := config.SetupRoutes()
    r.Run()
}
