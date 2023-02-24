package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

//go:embed dist
var embededFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("dist"))
	}

	log.Print("using embed mode")
	fsys, err := fs.Sub(embededFiles, "dist")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func main() {
	e := echo.New()
	useOS := len(os.Args) > 1 && os.Args[1] == "live"

	if useOS {
		e.Static("/", "dev-template.html")
	} else {
		fsys, err := fs.Sub(embededFiles, "dist")
		if err != nil {
			panic(err)
		}
		assetHandler := http.FileServer(http.FS(fsys))
		e.GET("/", echo.WrapHandler(assetHandler))
		e.GET("/assets/*", echo.WrapHandler(assetHandler))
	}

	// assetHandler := http.FileServer(getFileSystem(useOS))
	// e.GET("/", echo.WrapHandler(assetHandler))
	// e.GET("/assets/*", echo.WrapHandler(assetHandler))
	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}
