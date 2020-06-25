package main

import (
	"UtilsApi/controller"
	_ "UtilsApi/docs"
)

// @title Utils API
// @version 1.0
// @description 希望这些API能帮到你
// @termsOfService https://imagician.net/

func main() {
	controller.Run()
}
