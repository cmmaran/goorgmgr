package main

import (
	"github.com/cmmaran/goorgmgr/internal/layer/business"
	"github.com/cmmaran/goorgmgr/internal/server"
)

func main() {
	layer := business.NewBusiness()
	server.NewServer(layer).Start()
}
