package main

import (
	_ "github.com/xstnet/starfire-cloud/internal/db"
	"github.com/xstnet/starfire-cloud/internal/routers"
)

func main() {

	r := routers.SetupRouters()
	r.Run("127.0.0.1:9999")

	// user := models.User{Username: "test1", Email: "aaaa@qq.com", Password: "1234455", UsedSpace: 1234567894125}

	// result := DB.Create(&user)

	// fmt.Printf("ID:%d, Err: %s, rows: %d", user.ID, result.Error, result.RowsAffected)

}
