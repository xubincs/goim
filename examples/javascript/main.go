package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/", ".")
	e.Logger.Fatal(e.Start(":1999"))
}
