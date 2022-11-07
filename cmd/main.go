package main

import (
	"goodGopher/MyMenu/core"
)

func main() {

	g := core.MyServer()
	g.Run(":8081")
}
