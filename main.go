package main

import "gin-cli/base"

func main() {
	engine := base.RegisterEngine()
	base.RunEngine(engine)
}
