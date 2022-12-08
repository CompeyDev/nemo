package main

import (
	"github.com/CompeyDev/nemo/src/server"
	// "sync"
	// "github.com/CompeyDev/nemo/src/service/payload"
	// "github.com/CompeyDev/nemo/common/database"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(2)

	// go payload.Run(&wg)
	// go payload.ExecuteCommand(&wg, 2, "whoami", nil)

	// wg.Wait()
	server.InitializeConnection()
	// database.Test()
}
