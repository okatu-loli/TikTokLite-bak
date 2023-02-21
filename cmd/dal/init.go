package dal

import (
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
	"github.com/okatu-loli/TikTokLite/cmd/dal/rdb"
)

func Init() {
	db.Init()
	rdb.Init()
}
