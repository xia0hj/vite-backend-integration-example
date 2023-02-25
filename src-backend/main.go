package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"

	"github.com/labstack/echo/v4"
)

var cliMode = flag.String("mode", "prod", "launch in dev or prod")

//go:embed dist
var embededFiles embed.FS

func main() {
	e := echo.New()

	flag.Parse()

	if *cliMode == "dev" {
		log.Println("launch in development")
		e.Static("/", "dev-template.html")
	} else {
		log.Println("launch in production")
		fsys, err := fs.Sub(embededFiles, "dist")
		if err != nil {
			panic(err)
		}
		e.StaticFS("/", fsys)

	}
	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}
