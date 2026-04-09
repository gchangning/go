package main

import (
	"star/internal/cmd"
	_ "star/internal/packed"

	//_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {

	cmd.Main.Run(gctx.GetInitCtx())

}
