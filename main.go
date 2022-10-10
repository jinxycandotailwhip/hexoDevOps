package main

import (
	"hexoDevOps/server"
)

func main() {
	hexoServer := server.Server{}
	hexoServer.StartServe()
}
