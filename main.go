package main

import (
	"fmt"

	"github.com/xstnet/starfire-cloud/configs"
	_ "github.com/xstnet/starfire-cloud/configs"
	_ "github.com/xstnet/starfire-cloud/internal/db"
	"github.com/xstnet/starfire-cloud/internal/routers"
)

func main() {
	r := routers.SetupRouters()
	r.Run(fmt.Sprintf("%s:%d", configs.Server.Host, configs.Server.Port))
}
