namespace go tiktok

struct FeedRequest {
  1: optional i32 latest_time, // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
  2: optional string token // 可选参数，登录用户设置
}


struct FeedResponse {
  1: required i32 status_code, // 状态码，0-成功，其他值-失败
  2: optional string status_msg, // 返回状态描述
  3: list<Video> video_list, // 视频列表
  4: optional i64 next_time // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct PublishListRequest {
    1: i64 user_id,
    2: string token,
}

struct PublishListResponse {
    1: i32 status_code,
    2: optional string status_msg,
    3: list<Video> video_list,
}

struct Video {
  1: required i32 id, // 视频唯一标识
  2: required User author, // 视频作者信息
  3: required string play_url, // 视频播放地址
  4: required string cover_url, // 视频封面地址
  5: required i32 favorite_count, // 视频的点赞总数
  6: required i32 comment_count, // 视频的评论总数
  7: required bool is_favorite, // true-已点赞，false-未点赞
  8: required string title // 视频标题
}

struct User {
  1: required i32 id, // 用户id
  2: required string name, // 用户名称
  3: optional i32 follow_count, // 关注总数
  4: optional i32 follower_count, // 粉丝总数
  5: required bool is_follow, // true-已关注，false-未关注
  6: optional string avatar, // 用户头像
  7: optional string background_image, // 用户个人页顶部大图
  8: optional string signature, // 个人简介
  9: optional i32 total_favorited, // 获赞数量
  10: optional i32 work_count, // 作品数量
  11: optional i32 favorite_count // 点赞数量
}

struct PublishActionRequest {
  1: required string token, // 用户鉴权token
  2: FileData data, // 视频数据
  3: required string title // 视频标题
}

struct PublishActionResponse {
  1: required i32 status_code, // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
}

struct FileData {
    1: string name,     // 文件名
    2: string type,     // 文件类型
    3: binary content,  // 文件内容
}

service VideoService {
  FeedResponse Feed(1:FeedRequest request)
  PublishActionResponse PublishAction(1:PublishActionRequest request)
  PublishListResponse PublishList(1:PublishListRequest request)
}