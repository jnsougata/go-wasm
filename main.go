package main

import (
	"fmt"

	"github.com/jnsougata/gowasm/js"
)

func main() {
	ch := make(chan struct{}, 0)
	count := 0
	document := js.New()
	document.SetStyle("body {background-color:#7c4dff;text-align:center;align-items:center;padding:20px;color: #fff;font-family: Consolas,monaco,monospace;}")
	img := document.NewImg(
		&js.Image{
			Height: 200,
			Width:  200,
			Src:    "../static/wasm.png",
		},
	)
	htag := document.NewHTag("Hello, WebAssembly!", 3)
	htag.SetAttribute("style", "color: white;")
	button := document.NewButton("Click Me!")
	button.SetAttribute("style", "background-color: #fff;color: #7c4dff;padding: 10px 20px;border-radius: 5px;border: none;")
	button.SetAttribute("id", "btn")
	button.OnCLick(func(this js.Value, args []js.Value) interface{} {
		count++
		htag.SetAttribute("innerHTML", fmt.Sprintf("Clicked %d times!", count))
		return nil
	})
	document.Append(img.JSValue, htag.JSValue, button.JSValue)
	<-ch
}
