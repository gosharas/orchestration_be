package main

import (
	"github.com/gosharas/orchestration_be/impl"
	"log"
)

func main() {
	impl.Serve("test")
	log.Println("gateway started")
}
