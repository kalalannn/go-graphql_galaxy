package main

import (
	"go-graphql_galaxy/internal/app"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	app.NewApp().Run(&wg)
	wg.Wait()
}
