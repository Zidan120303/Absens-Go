package main

import (
	"APLIKASI_1/module"
	"APLIKASI_1/route"
)

func main() {
	module.InitialMigration()

	module.PresenIniMigration()

	route.InitializeRouter()
}
