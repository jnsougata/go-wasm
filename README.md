# go-wasm

# Installation
```shell
go get github.com/jnsougata/go-wasm
```

# Example

```go
package main

import (
	"fmt"
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
				"margin":           "20px",
			},
		})
	img := document.CreateElement("img")
	img.Set("src", "../static/wasm.png")
	img.Set("width", "200")
	img.Set("height", "200")
	button := document.CreateElement("button")
	button.Set("innerHTML", "Click Me!")
	val := 0
	button.OnClick(func(this bridge.Arg, args []bridge.Arg) interface{} {
		val++
		tag := document.GetElementById("count")
		tag.Set("innerHTML", fmt.Sprintf("Clicked x%d", val))
		return nil
	})
	h3 := document.CreateElement("h3")
	h3.Set("innerHTML", "Clicked x0")
	h3.Set("id", "count")
	div := document.CreateElement("div")
	div.Set("id", "container")
	div.Call("appendChild", button)
	document.Body.AppendChild(img)
	document.Body.AppendChild(h3)
	document.Body.AppendChild(div)
	<-ch
}

```

# Sample Output
![img.png](example.png)
