package favorite

import (
	"github.com/jinzhu/copier"
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
	"github.com/okatu-loli/TikTokLite/internal/repository"
	"github.com/okatu-loli/TikTokLite/internal/request"
	"github.com/okatu-loli/TikTokLite/internal/response"
	"log"
)

type IFavoriteService interface {
	FavoriteAction(userId int64, favoriteRequest request.FavoriteActionParam) //点赞取消赞的操作
	FavoriteVideoList(r *request.FavoriteListParam) response.VideoList        // 喜欢列表
}

type FavoritesService struct {
	VideoRepository    repository.IVideoRepository
	FavoriteRepository repository.IFavoriteRepository
	UserRepository     repository.IUserRepository
}

func (f FavoritesService) FavoriteVideoList(r *request.FavoriteListParam) response.VideoList {

	userId := r.UserId
	user, err := f.UserRepository.GetUserById(userId)
	if err != nil {
		log.Printf("获取用户信息失败")
		return response.VideoList{}
	}
	var userResponse response.User
	err = copier.Copy(&userResponse, &user)
	if err != nil {
		return response.VideoList{}
	}

	likes, err := f.FavoriteRepository.QueryFavorVideoListByUserId(userId)
	if err != nil {
		return response.VideoList{}
	}
	//建立  在赋值
	videoIds := make([]int64, len(likes))
	for idx, like := range likes {
		videoIds[idx] = like.VideoId
	}
	videos, err := f.VideoRepository.GetVideoListByVideoIds(videoIds)
	if err != nil {
		return response.VideoList{}
	}
	videosResponse := make([]response.Video, len(videos))
	err = copier.Copy(&videosResponse, &videos)
	if err != nil {
		log.Printf("拷贝失败")
		return response.VideoList{}
	}
	for _, video := range videosResponse {
		video.User = userResponse
	}
	//for i := range  videos {
	//	user, err := f.UserRepository.GetUserById(userId)
	//	if err == nil { //若查询未出错则更新，否则不更新作者信息
	//		videos[i].UserId = user.ID
	//	}
	//	videos[i].IsFavorite = true //喜欢状态
	//}
	return response.VideoList{
		VideoList:      videosResponse,
		CommonResponse: response.CommonResponse{StatusMsg: "success", StatusCode: 0},
	}
}

// FavoriteAction 点赞和取消赞的操作
func (f FavoritesService) FavoriteAction(userId int64, favoriteRequest request.FavoriteActionParam) {
	actionType := favoriteRequest.ActionType
	videoId := favoriteRequest.VideoId
	begin := db.DB.Begin()
	if actionType == 1 {
		err := f.FavoriteRepository.PlusOneFavorByUserIdAndVideoId(videoId, userId)
		if err != nil {
			log.Printf("FavoriteAction|点赞失败|%v", err)
			begin.Rollback()
			return
		}

	} else if actionType == 2 {
		err := f.FavoriteRepository.MinusOneFavorByUserIdAndVideoId(userId, videoId)
		if err != nil {
			log.Printf("FavoriteAction|取消点赞失败|%v", err)
			begin.Rollback()
			return
		}
	} else {
		log.Printf("FavoriteAction|参数错误|actionType=%v", actionType)
	}
	begin.Commit()
}

func NewFavoriteService() IFavoriteService {
	favoriteService := FavoritesService{
		//点赞和增加点赞数量放到了一起  不在单独对video进行
		//VideoRepository:    repository.NewVideoRepository(),
		FavoriteRepository: repository.NewFavoriteRepository(),
		UserRepository:     repository.NewUserRepository(),
	}
	return favoriteService
}
