package main

import (
	"fmt"

	_ "github.com/xstnet/starfire-cloud/boot"
	"github.com/xstnet/starfire-cloud/configs"
	"github.com/xstnet/starfire-cloud/internal/routers"
)

func main() {
	r := routers.SetupRouters()
	r.Run(fmt.Sprintf("%s:%d", configs.Server.Host, configs.Server.Port))
}
