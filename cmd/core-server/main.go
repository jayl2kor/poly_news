package core_server

import (
	"os"
	"poly_news/cmd/core-server/flag"
	"poly_news/internal/config"
	"poly_news/internal/repository/gorm"

	"github.com/samber/lo"
)

func main() {
	coreServerFlag := lo.Must1(flag.ParseFlags(os.Args[1:]))
	coreServerConfig := lo.Must1(config.Parse(coreServerFlag))

	db := lo.Must1(gorm.New(coreServerConfig.Database))

}
