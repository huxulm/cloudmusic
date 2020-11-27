package entities

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	CK_CSRF         = "__csrf"
	CK_MUSIC_U      = "MUSIC_U"
	CK_MUSIC_A      = "MUSIC_A"
	CK_OSVER        = "osver" // 系统版本
	CK_DEVICE_ID    = "deviceId"
	CK_APPVER       = "appver"
	CK_VERSION_CODE = "versioncode"
	CK_MOBILE_NAME  = "mobilename"
	CK_BUILD_VER    = "buildver"
	CK_RESOLUTION   = "resolution"
	CK_OS           = "os"
	CK_CHANNEL      = "channel"
	CK_REQUEST_ID   = "requestId"
)

type Cookies []*http.Cookie

func LoadFromJSON(d []byte) ([]*http.Cookie, error) {
	cokis := make([]*http.Cookie, 0)
	err := json.Unmarshal(d, &cokis)
	if err != nil {
		return nil, err
	}
	return cokis, nil
}

func LoadFromYaml(d []byte) ([]*http.Cookie, error) {
	cokis := make([]*http.Cookie, 0)
	err := yaml.Unmarshal(d, &cokis)
	if err != nil {
		return nil, err
	}
	return cokis, nil
}

func (cs *Cookies) String() string {
	cookie := make([]string, len(*cs))
	for i, c := range *cs {
		cookie[i] = c.String()
	}
	return strings.Join(cookie, "; ")
}

func (cs *Cookies) FindByKey(key string) *http.Cookie {
	for _, c := range *cs {
		if c.Name == key {
			return c
		}
	}
	return nil
}

type Result struct {
	Body    []byte
	Header  http.Header
	Cookies Cookies
	Status  int
}

func (r *Result) AsJSON() string {
	if d, err := json.Marshal(r); err == nil {
		return string(d)
	}
	return ""
}

func (r *Result) BodyAsString() string {
	if len(r.Body) > 0 {
		return string(r.Body)
	}
	return ""
}

func (r *Result) ToLoginCellRes() (*LoginCellRes, error) {
	if len(r.Body) > 0 {
		var lcr LoginCellRes
		if err := json.Unmarshal(r.Body, &lcr); err == nil {
			return &lcr, nil
		} else {
			return nil, err
		}
	} else {
		return nil, errors.New("body is empty")
	}
}
func (r *Result) ToLoginStatusRes() (*LoginCellRes, error) {
	var lcr LoginCellRes
	if err := r.toRes(&lcr); err == nil {
		return &lcr, nil
	} else {
		return nil, err
	}
}

func (r *Result) ToUserDetailRes() (*UserDetailRes, error) {
	var d UserDetailRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}

func (r *Result) ToSubcountRes() (*Subcount, error) {
	var d Subcount
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToPlaylistRes() (*PlaylistRes, error) {
	var d PlaylistRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToPlaylistDetailRes() (*PlaylistDetailRes, error) {
	var d PlaylistDetailRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToSongURLRes() (*SongURLRes, error) {
	var d SongURLRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToSongDetailRes() (*SongDetailRes, error) {
	var d SongDetailRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToSongLyricRes() (*SongLyricRes, error) {
	var d SongLyricRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToSearchRes() (*SearchRes, error) {
	var d SearchRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToHqTagsRes() (*HqTagsRes, error) {
	var d HqTagsRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToLikeRes() (*LikeRes, error) {
	var d LikeRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToLikeListRes() (*LikeListRes, error) {
	var d LikeListRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToRecSongsRes() (*RecSongsRes, error) {
	var d RecSongsRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToRecResourceRes() (*RecResourceRes, error) {
	var d RecResourceRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}
func (r *Result) ToPersonalFMRes() (*PersonalFMRes, error) {
	var d PersonalFMRes
	if err := r.toRes(&d); err == nil {
		return &d, nil
	} else {
		return nil, err
	}
}

func (r *Result) toRes(typo interface{}) error {
	if len(r.Body) > 0 {
		if err := json.Unmarshal(r.Body, typo); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return errors.New("body is empty")
	}
}

func (r *Result) Unmarshal(in interface{}) error {
	if err := r.toRes(in); err == nil {
		return nil
	} else {
		return err
	}
}

type Account struct {
	ID                 int64  `json:"id,omitempty"`
	UserName           string `json:"userName,omitempty"`
	Type               int    `json:"type,omitempty"`
	Status             int    `json:"status,omitempty"`
	WhitelistAuthority int    `json:"whitelistAuthority,omitempty"`
	CreateTime         int64  `json:"createTime,omitempty"`
	Salt               string `json:"salt,omitempty"`
	TokenVersion       int    `json:"tokenVersion,omitempty"`
	Ban                int    `json:"ban,omitempty"`
	BaoyueVersion      int    `json:"baoyueVersion,omitempty"`
	DonateVersion      int    `json:"donateVersion,omitempty"`
	VipType            int    `json:"vipType,omitempty"`
	ViptypeVersion     int64  `json:"viptypeVersion,omitempty"`
	AnonimousUser      bool   `json:"anonimousUser,omitempty"`
}
type Profile struct {
	Mutual                    bool        `json:"mutual,omitempty" yaml:"mutual"`                                       // false,
	RemarkName                string      `json:"remarkName,omitempty" yaml:"remarkName"`                               // null,
	ExpertTags                []string    `json:"expertTags,omitempty" yaml:"expertTags"`                               // null,
	AuthStatus                int         `json:"authStatus,omitempty" yaml:"authStatus"`                               // 0,
	Experts                   interface{} `json:"experts,omitempty" yaml:"experts"`                                     // {},
	Followed                  bool        `json:"followed,omitempty" yaml:"followed"`                                   // false,
	BackgroundUrl             string      `json:"backgroundUrl,omitempty" yaml:"backgroundUrl"`                         // "https://p3.music.126.net/XHZJQErGslqPfrcH7CJC9A==/109951163285349613.jpg",
	DetailDescription         string      `json:"detailDescription,omitempty" yaml:"detailDescription"`                 // "",
	AvatarImgIdStr            string      `json:"avatarImgIdStr,omitempty" yaml:"avatarImgIdStr"`                       // "2907108744406549",
	BackgroundImgIdStr        string      `json:"backgroundImgIdStr,omitempty" yaml:"backgroundImgIdStr"`               // "109951163285349613",
	VipType                   int         `json:"vipType,omitempty" yaml:"vipType"`                                     // 0,
	Gender                    int         `json:"gender,omitempty" yaml:"gender"`                                       // 1,
	AccountStatus             int         `json:"accountStatus,omitempty" yaml:"accountStatus"`                         // 0,
	AvatarImgID               int64       `json:"avatarImgId,omitempty" yaml:"avatarImgId"`                             // 2907108744406549,
	Nickname                  string      `json:"nickname,omitempty" yaml:"nickname"`                                   // "<your nickname>",
	Birthday                  int64       `json:"birthday,omitempty" yaml:"birthday"`                                   // 694195200000,
	City                      int64       `json:"city,omitempty" yaml:"city"`                                           // 500101,
	AvatarURL                 string      `json:"avatarUrl,omitempty" yaml:"avatarUrl"`                                 // "https://p4.music.126.net/uQv7-TsW88LaQYBwmkN5Kg==/2907108744406549.jpg",
	DefaultAvatar             bool        `json:"defaultAvatar,omitempty" yaml:"defaultAvatar"`                         // false,
	Province                  int64       `json:"province,omitempty" yaml:"province"`                                   // 500000,
	DjStatus                  int         `json:"djStatus,omitempty" yaml:"djStatus"`                                   // 0,
	BackgroundImgID           int64       `json:"backgroundImgId,omitempty" yaml:"backgroundImgId"`                     // 109951163285349613,
	Description               string      `json:"description,omitempty" yaml:"description"`                             // "",
	UserID                    int64       `json:"userId,omitempty" yaml:"userId"`                                       // 45964623,
	UserType                  int         `json:"userType,omitempty" yaml:"userType"`                                   // 0,
	Signature                 string      `json:"signature,omitempty" yaml:"signature"`                                 // "裸奔在无人的沙漠里",
	Authority                 int         `json:"authority,omitempty" yaml:"authority"`                                 // 0,
	Followeds                 int64       `json:"followeds,omitempty" yaml:"followeds"`                                 // 3,
	Follows                   int64       `json:"follows,omitempty" yaml:"follows"`                                     // 28,
	EventCount                int64       `json:"eventCount,omitempty" yaml:"eventCount"`                               // 7,
	AvatarDetail              interface{} `json:"avatarDetail,omitempty" yaml:"avatarDetail"`                           // null,
	PlaylistCount             int64       `json:"playlistCount,omitempty" yaml:"playlistCount"`                         // 21,
	PlaylistBeSubscribedCount int64       `json:"playlistBeSubscribedCount,omitempty" yaml:"playlistBeSubscribedCount"` // 0
}
type Binding struct {
	RefreshTime  int64  `json:"refreshTime,omitempty" yaml:"refreshTime"`
	BindingTime  int64  `json:"bindingTime,omitempty" yaml:"bindingTime"`
	ExpiresIn    int64  `json:"expiresIn,omitempty" yaml:"expiresIn"`
	TokenJsonStr string `json:"tokenJsonStr,omitempty" yaml:"tokenJsonStr"`
	UserID       int64  `json:"userId,omitempty" yaml:"userId"`
	URL          string `json:"url,omitempty" yaml:"url"`
	Expired      bool   `json:"expired,omitempty" yaml:"expired"`
	ID           int64  `json:"id,omitempty" yaml:"id"`
	Type         int    `json:"type,omitempty" yaml:"type"`
}
type LoginCellRes struct {
	LoginType  int        `json:"loginType,omitempty" yaml:"loginType"`
	ClientID   string     `json:"clientId,omitempty" yaml:"clientId"`
	EffectTime int64      `json:"effectTime,omitempty" yaml:"effectTime"`
	Code       int        `json:"code,omitempty" yaml:"code"`
	Account    *Account   `json:"account,omitempty" yaml:"account"`
	Token      string     `json:"token,omitempty" yaml:"token"`
	Profile    *Profile   `json:"profile,omitempty" yaml:"profile"`
	Bindings   []*Binding `json:"binding,omitempty" yaml:"binding"`
}

// UserDetailRes 用户详情
type UserDetailRes struct {
	Level       int   `json:"level,omitempty"`
	ListenSongs int64 `json:"listenSongs,omitempty"`
	MobileSign  bool  `json:"mobileSign,omitempty"`
	PcSign      bool  `json:"pcSign,omitempty"`
	UserPoint   struct {
		UserId       int64 `json:"userId,omitempty"`
		Balance      int64 `json:"balance,omitempty"`
		UpdateTime   int64 `json:"updateTime,omitempty"`
		Version      int64 `json:"version,omitempty"`
		Status       int64 `json:"status,omitempty"`
		BlockBalance int64 `json:"blockBalance,omitempty"`
	} `json:"userPoint,omitempty"`
	Profile                  Profile   `json:"profile,omitempty"`
	Bindings                 []Binding `json:"bindings,omitempty"`
	PeopleCanSeeMyPlayRecord bool      `json:"peopleCanSeeMyPlayRecord,omitempty"`
	AdValid                  bool      `json:"adValid,omitempty"`
	Code                     int       `json:"code,omitempty"`
	CreateTime               int64     `json:"createTime,omitempty"`
	CreateDays               int64     `json:"createDays,omitempty"`
}

// 歌单，收藏，mv，dj数量
type Subcount struct {
	ProgramCount         int64 `json:"programCount,omitempty"`
	DjRadioCount         int64 `json:"djRadioCount,omitempty"`
	MvCount              int64 `json:"mvCount,omitempty"`
	ArtistCount          int64 `json:"artistCount,omitempty"`
	NewProgramCount      int64 `json:"newProgramCount,omitempty"`
	CreateDjRadioCount   int64 `json:"createDjRadioCount,omitempty"`
	CreatedPlaylistCount int64 `json:"createdPlaylistCount,omitempty"`
	SubPlaylistCount     int64 `json:"subPlaylistCount,omitempty"`
	Code                 int64 `json:"code,omitempty"`
}

type PlaylistRes struct {
	Version    string `json:"version,omitempty"` //"1603414754372",
	More       bool   `json:"more,omitempty"`    //false,
	Subscribed bool   `json:"subscribed,omitempty"`
	Playlist   []struct {
		Creator               Profile     `json:"creator,omitempty"`
		Artists               interface{} `json:"artists,omitempty"`
		Tracks                interface{} `json:"tracks,omitempty"`
		UpdateFrequency       interface{} `json:"updateFrequency,omitempty"`
		BackgroundCoverId     int64       `json:"backgroundCoverId,omitempty"`
		BackgroundCoverUrl    string      `json:"backgroundCoverUrl,omitempty"`
		TitleImage            int64       `json:"titleImage,omitempty"`
		TitleImageUrl         string      `json:"titleImageUrl,omitempty"`
		EnglishTitle          string      `json:"englishTitle,omitempty"`
		OpRecommend           bool        `json:"opRecommend,omitempty"`
		RecommendInfo         interface{} `json:"recommendInfo,omitempty"`
		UserId                int64       `json:"userId,omitempty"`
		AdType                int         `json:"adType,omitempty"`
		TrackNumberUpdateTime int64       `json:"trackNumberUpdateTime,omitempty"`
		CreateTime            int64       `json:"createTime,omitempty"`
		HighQuality           bool        `json:"highQuality,omitempty"`
		CoverImgId            int64       `json:"coverImgId,omitempty"`
		NewImported           bool        `json:"newImported,omitempty"`
		Anonimous             bool        `json:"anonimous,omitempty"`
		UpdateTime            int64       `json:"updateTime,omitempty"`
		CoverImgUrl           string      `json:"coverImgUrl,omitempty"`
		SpecialType           int         `json:"specialType,omitempty"`
		TotalDuration         int64       `json:"totalDuration,omitempty"`
		TrackCount            int64       `json:"trackCount,omitempty"`
		CommentThreadId       string      `json:"commentThreadId,omitempty"`
		Privacy               int64       `json:"privacy,omitempty"`
		TrackUpdateTime       int64       `json:"trackUpdateTime,omitempty"`
		PlayCount             int64       `json:"playCount,omitempty"`
		SubscribedCount       int64       `json:"subscribedCount,omitempty"`
		CloudTrackCount       int64       `json:"cloudTrackCount,omitempty"`
		Description           string      `json:"description,omitempty"`
		Ordered               bool        `json:"ordered,omitempty"`
		Tags                  []string
		Status                int    `json:"status"`
		Name                  string `json:"name"`
		Id                    int64  `json:"id"`
	} `json:"playlist,omitempty"`
	Code int `json:"code,omitempty"`
}

type TrackItem struct {
	Name string `json:"name,omitempty"`
	Id   int64  `json:"id,omitempty"`
	Pst  int64  `json:"pst,omitempty"`
	T    int64  `json:"t,omitempty"`
	Ar   []struct {
		Id    int64         `json:"id,omitempty"`
		Name  string        `json:"name,omitempty"`
		Tns   []interface{} `json:"tns,omitempty"`
		Alias []interface{} `json:"alias,omitempty"`
	} `json:"ar,omitempty"`
	Alia []interface{} `json:"alia,omitempty"`
	Pop  float64       `json:"pop,omitempty"`
	St   int64         `json:"st,omitempty"`
	Rt   string        `json:"rt,omitempty"`
	Fee  int64         `json:"fee,omitempty"`
	V    int64         `json:"v,omitempty"`
	Crbt interface{}   `json:"crbt,omitempty"`
	Cf   string        `json:"cf,omitempty"`
	Al   struct {
		Id     int64         `json:"id,omitempty"`
		Name   string        `json:"name,omitempty"`
		PicUrl string        `json:"picUrl,omitempty"`
		Tns    []interface{} `json:"tns,omitempty"`
		Pic    int64         `json:"pic,omitempty"`
	} `json:"al,omitempty"`
	Dt int64 `json:"dt,omitempty"` // 280706,
	H  struct {
		Br   int64   `json:"br,omitempty"`
		Fid  int64   `json:"fid,omitempty"`
		Size int64   `json:"size,omitempty"`
		Vd   float64 `json:"vd,omitempty"`
	} `json:"h,omitempty"`
	M struct {
		Br   int64   `json:"br,omitempty"`
		Fid  int64   `json:"fid,omitempty"`
		Size int64   `json:"size,omitempty"`
		Vd   float64 `json:"vd,omitempty"`
	} `json:"m,omitempty"`
	L struct {
		Br   int64   `json:"br,omitempty"`
		Fid  int64   `json:"fid,omitempty"`
		Size int64   `json:"size,omitempty"`
		Vd   float64 `json:"vd,omitempty"`
	} `json:"l,omitempty"`
	A               interface{} `json:"a,omitempty"`
	Cd              string      `json:"cd,omitempty"`
	No              int64       `json:"no,omitempty"`
	RtUrl           string      `json:"rtUrl,omitempty"`
	Ftype           int         `json:"ftype,omitempty"`
	RtUrls          []string    `json:"rtUrls,omitempty"`
	DjId            int64       `json:"djId,omitempty"`
	Copyright       int64       `json:"copyright,omitempty"`
	S_id            int64       `json:"s_id,omitempty"`
	Mark            int64       `json:"mark,omitempty"`
	OriginCoverType int         `json:"originCoverType,omitempty"`
	NoCopyrightRcmd interface{} `json:"noCopyrightRcmd,omitempty"`
	Rtype           int         `json:"rtype,omitempty"`
	Rurl            string      `json:"rurl,omitempty"`
	Mst             int         `json:"mst,omitempty"`
	Cp              int64       `json:"cp,omitempty"`
	Mv              int64       `json:"mv,omitempty"`
	PublishTime     int64       `json:"publishTime,omitempty"`
}
type PlaylistDetailRes struct {
	Code          int         `json:"code,omitempty"`
	RelatedVideos interface{} `json:"relatedVideos,omitempty"`
	Playlist      struct {
		Creator     Profile       `json:"creator,omitempty"`
		Subscribers []interface{} `json:"subscribers,omitempty"`
		Subscribed  bool          `json:"subscribed,omitempty"`
		Tracks      []TrackItem   `json:"tracks,omitempty"`
	} `json:"playlist,omitempty"`
	URLs []string `json:"urls,omitempty"`
}

type SongItem struct {
	ID                 int64       `json:"id"`
	URL                string      `json:"url"`
	Br                 int64       `json:"br"`
	Size               int64       `json:"size"`
	Md5                string      `json:"md5"`
	Code               int         `json:"code"`
	Expi               int64       `json:"expi"`
	Type               string      `json:"type"`
	Gain               float64     `json:"gain"`
	Fee                float64     `json:"fee"`
	Uf                 interface{} `json:"uf"`
	Payed              int         `json:"payed"`
	Flag               int         `json:"flag"`
	CanExtend          bool        `json:"canExtend"`
	FreeTrialInfo      interface{} `json:"freeTrialInfo"`
	Level              string      `json:"level"`
	EncodeType         string      `json:"encodeType"`
	FreeTrialPrivilege struct {
		ResConsumable  bool `json:"resConsumable"`
		UserConsumable bool `json:"userConsumable"`
	}
	UrlSource int64 `json:"urlSource"` // 0
}
type SongURLRes struct {
	Data []*SongItem `json:"data,omitempty"`
	Code int         `json:"code,omitempty"`
}
type SongDetailRes struct {
	Songs []*Song `json:"songs,omitempty"`
	Code  int     `json:"code,omitempty"`
}

type Lyric struct {
	Version int    `json:"version,omitempty"`
	Lyric   string `json:"lyric,omitempty"`
}
type SongLyricRes struct {
	Sgc    bool `json:"sgc,omitempty"`
	Sfy    bool `json:"sfy,omitempty"`
	Qfy    bool `json:"qfy,omitempty"`
	Lyric  `json:"lrc,omitempty"`
	Klyric struct {
		Version int    `json:"version,omitempty"`
		Lyric   string `json:"lyric,omitempty"`
	} `json:"klyric"`
	Tlyric struct {
		Version int    `json:"version,omitemtpy"`
		Lyric   string `json:"lyric,omitemtpy"`
	} `json:"tlyric"`
	Code int `json:"code,omitempty"`
}

type Artist struct {
	ID        int64         `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	PicUrl    string        `json:"picUrl,omitempty"`
	Alias     []interface{} `json:"alias"`
	AlbumSize int64         `json:"albumSize,omitempty"`
	PicID     int64         `json:"picId,omitempty"`
	Img1v1Url string        `json:"img1v1Url,omitempty"`
	Img1v1    int64         `json:"img1v1,omitempty"`
	Trans     interface{}   `json:"trans,omitempty"`
}
type Album struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Artist      Artist `json:"artist,omitempty"`
	PublishTime int64  `json:"publishTime,omitempty"`
	Size        int    `json:"size,omitempty"`
	CopyrightId int64  `json:"copyrightId,omitempty"`
	Status      int    `json:"status,omitempty"`
	PicID       int64  `json:"picId,omitempty"`
	Mark        int64  `json:"mark,omitempty"`
}
type Song struct {
	ID          int64     `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Artists     []*Artist `json:"artists,omitempty"`
	Album       `json:"album,omitempty"`
	Duration    int64    `json:"duration,omitempty"`
	CopyrightId int64    `json:"copyrightId,omitempty"`
	Status      int      `json:"status,omitempty"`
	Alias       []string `json:"alias,omitempty"`
	Rtype       int      `json:"rtype,omitempty"`
	Ftype       int      `json:"ftype,omitempty"`
	Mvid        int64    `json:"mvid,omitempty"`
	Fee         int64    `json:"fee,omitempty"`
	RUrl        string   `json:"rUrl,omitempty"`
	Mark        int64    `json:"mark,omitempty"`
}
type SearchRes struct {
	Result struct {
		Songs []*Song `json:"songs,omitempty"`
	}
	Code int `json:"code,omitempty"`
}
type Tag struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     int    `json:"type,omitempty"`
	Category int    `json:"category,omitempty"`
	Hot      bool   `json:"hot,omitempty"`
}

type HqTagsRes struct {
	Tags []*Tag `json:"tags,omitempty"`
	Code int    `json:"code,omitempty"`
}

type LikeRes struct {
	Songs      []interface{} `json:"songs,omitempty"`
	PlaylistID int64         `json:"playlistId,omitempty"`
	Code       int           `json:"code,omitempty"`
}
type LikeListRes struct {
	IDs        []int64 `json:"ids,omitempty"`
	CheckPoint int64   `json:"checkPoint,omitempty"`
	Code       int     `json:"code,omitempty"`
}

// 推荐歌曲
type RecmReason struct {
	SongId int64  `json:"songId"`
	Reason string `json:"reason"`
}
type RecSongsRes struct {
	Data struct {
		DailySongs       []TrackItem  `json:"dailySongs,omitempty"`
		RecommendReasons []RecmReason `json:"recommendReasons,omitempty"`
	} `json:"data,omitempty"`
	Code int `json:"code,omitempty"`
}

// 推荐资源
type RecResourceRes struct {
	FeatureFirst  bool          `json:"featureFirst"`
	HaveRcmdSongs bool          `json:"haveRcmdSongs"`
	Recommend     []interface{} `json:"recommend,omitempty"`
	Code          int           `json:"code,omitempty"`
}

type MusicMeta struct {
	Name        string  `json:"name,omitempty"`        // null,
	Id          int64   `json:"id,omitempty"`          // 4936700402,
	Size        int64   `json:"size,omitempty"`        // 3996987,
	Extension   string  `json:"extension,omitempty"`   // "mp3",
	Sr          int64   `json:"sr,omitempty"`          // 44100,
	DfsId       int64   `json:"dfsId,omitempty"`       // 0,
	Bitrate     int64   `json:"bitrate,omitempty"`     // 128000,
	PlayTime    int64   `json:"playTime,omitempty"`    // 249782,
	VolumeDelta float64 `json:"volumeDelta,omitempty"` // -43295
}

// 私人FM
type FMSong struct {
	Song
	Position int64     `json:"position,omitempty"`
	No       int64     `json:"no,omitempty"`
	Disc     string    `json:"disc,omitempty"`
	Starred  bool      `json:"starred,omitempty"`
	MvID     int64     `json:"mvid,omitempty"`
	Mp3URL   string    `json:"mp3Url,omitempty"` // null
	BMusic   MusicMeta `json:"bMusic,omitempty"`
	HMusic   MusicMeta `json:"hMusic,omitempty"`
	MMusic   MusicMeta `json:"mMusic,omitempty"`
	LMusic   MusicMeta `json:"lMusic,omitempty"`
}
type PersonalFMRes struct {
	Code      int      `json:"code,omitempty"`
	PopAdjust bool     `json:"popAdjust,omitempty"`
	Data      []FMSong `json:"data,omitempty"`
}
