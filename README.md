# go-wasm

# Installation
```shell
go get github.com/jnsougata/go-wasm
```

# Example

```go
package main

import (
	"github.com/jnsougata/gowasm/bridge"
)

func main() {
	ch := make(chan struct{}, 0)
	document := bridge.New()
	document.StyleSheet("../static/style.css")
	document.CSSClasses(
		bridge.Class{
			Name: "button",
			Attributes: map[string]string{
				"background-color": "#fff",
				"color":            "#7c4dff",
				"padding":          "10px 20px",
				"border-radius":    "5px",
				"border":           "none",
				"cursor":           "pointer",
				"margin":           "10px",
			},
		})
	img := document.CreateElement("img")
	img.Set("src", "../static/wasm.png")
	img.Set("width", "200")
	img.Set("height", "200")
	button := document.CreateElement("button")
	button.Set("innerHTML", "Click Me!")
	button.OnClick(func(this bridge.Arg, args []bridge.Arg) interface{} {
		img.Set("src", "../static/img.png")
		return nil
	})
	div := document.CreateElement("div")
	div.Set("id", "container")
	div.Set("className", "abc")
	div.Call("appendChild", button)
	document.Body.AppendChild(img)
	document.Body.AppendChild(div)
	<-ch
}

```

# Sample Output
![img.png](example.png)
