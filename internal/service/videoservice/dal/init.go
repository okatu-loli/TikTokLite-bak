package dal

import (
	"TikTokLiteV2/internal/service/videoservice/dal/db"
	"TikTokLiteV2/internal/service/videoservice/dal/rdb"
)

func Init() {
	db.Init()
	rdb.Init()
}
