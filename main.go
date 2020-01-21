package main

import (
	"github.com/daviresio/financeiro_api/controller"
)


func main() {

	controller.MapUrls()

	controller.GetRouter().Run(":8000")
}
