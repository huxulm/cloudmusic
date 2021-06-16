package api

import "fmt"

const (
	BASE_URL_1 = "https://music.163.com/weapi"

	BASE_URL_2 = "https://music.163.com"

	BASE_URL_3 = "https://music.163.com/weapi/v1"
	BASE_URL_4 = "https://music.163.com/api"
	BASE_URL_5 = "https://interface3.music.163.com/eapi"
)

func ConcatBase(base, segement string) string {
	return fmt.Sprintf("%s%s", base, segement)
}

var (
	// 手机号登录
	LOGIN_CELLPHONE = ConcatBase(BASE_URL_1, "/login/cellphone")
	// 邮箱登录
	LOGIN_EMAIL = ConcatBase(BASE_URL_1, "/login")
	// 刷新登录
	LOGIN_REFRESH = ConcatBase(BASE_URL_1, "/login/refresh")
	// 登录状态
	LOGIN_STATUS = ConcatBase(BASE_URL_2, "")
	// 发送验证码
	CAPTCHA_SENT = ConcatBase(BASE_URL_1, "/captcha/sent")
	// 验证验证码
	CAPTCHA_VERIFY = ConcatBase(BASE_URL_1, "/captcha/verify")
	// 退出登录
	LOGOUT = ConcatBase(BASE_URL_1, "/logout")
	// 用户详情
	USER_DETAIL = ConcatBase(BASE_URL_3, "/user/detail")
	// 用户收藏计数
	USER_SUBCOUNT = ConcatBase(BASE_URL_1, "/subcount")
	// 用户等级信息
	USER_LEVEL = ConcatBase(BASE_URL_1, "/user/level")
	// 账号信息
	ACCOUNT = ConcatBase(BASE_URL_4, "/nuser/account/get")
	// 用户歌单
	USER_PLAYLIST = ConcatBase(BASE_URL_4, "/user/playlist")
	// 歌单详情
	PLAYLIST_DETAIL = ConcatBase(BASE_URL_4, "/v6/playlist/detail")
	// 歌曲链接
	SONG_URL = ConcatBase(BASE_URL_5, "/song/enhance/player/url")
	// 歌曲详情
	SONG_DETAIL = ConcatBase(BASE_URL_1, "/v3/song/detail")
	// 歌词
	SONG_LYRIC = ConcatBase(BASE_URL_4, "/song/lyric")
	// 搜索
	SEARCH = ConcatBase(BASE_URL_1, "/search/get")
	// 精品歌单标签
	PLAYLIST_HIGHQUALITY_TAGS = ConcatBase(BASE_URL_4, "/playlist/highquality/tags")
	// 喜欢音乐列表
	LIKE_LIST = ConcatBase(BASE_URL_1, "/song/like/get")
	// 喜欢音乐
	LIKE = ConcatBase(BASE_URL_1, "/radio/like")
	// 每日推荐歌单
	RECOMMEND_RESOURCE = ConcatBase(BASE_URL_1, "/v1/discovery/recommend/resource")
	// 每日推荐歌曲
	RECOMMEND_SONGS = ConcatBase(BASE_URL_4, "/v3/discovery/recommend/songs")
	// 私人FM
	PERSONAL_FM = ConcatBase(BASE_URL_3, "/radio/get")
)
