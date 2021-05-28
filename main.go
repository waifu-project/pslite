package main

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/d1y/pslite/port"
	"github.com/wailsapp/wails"
)

//go:embed psd/**
var html embed.FS

var _PORT int

func init() {
	_PORT = port.GetPort()
	go func() {
		var filesystem = http.FS(html)
		var fs = http.FileServer(filesystem)
		http.Handle("/", fs)
		var port = fmt.Sprintf(":%d", _PORT)
		var err = http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func createHtml() string {
	return fmt.Sprintf(`<script type="text/javascript">window.location.href='http://127.0.0.1:%d/psd/psd.html'</script>`, _PORT)
}

func main() {

	var html = createHtml()

	app := wails.CreateApp(&wails.AppConfig{
		Width:            1024,
		Height:           768,
		HTML:             html,
		DisableInspector: false,
		Title:            "pslite",
		Resizable:        true,
	})
	app.Run()
}
