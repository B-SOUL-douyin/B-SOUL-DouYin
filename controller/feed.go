package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"

	"github.com/B-SOUL-douyin/B-SOUL-DouYin/pkg/constants"
	"github.com/B-SOUL-douyin/B-SOUL-DouYin/video/dal/db"
)

type FeedResponse struct {
	constants.Response
	VideoList []constants.Video `json:"video_list,omitempty"`
	NextTime  int64             `json:"next_time,omitempty"`
}

var StatusCode int32 = 0

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var FeedVideos []constants.Video
	VideoModelTmp, err := db.GetFeedList()
	resp := constants.Response{}
	if err != nil {
		resp.StatusMsg = resp.StatusMsg + ";" + err.Error()
	}
	if len(VideoModelTmp) == 0 {
		resp.StatusMsg = resp.StatusMsg + ";没有新的视频了"
	} else {
		for _, v := range VideoModelTmp {
			FeedVideos = append(FeedVideos, v.Video)
		}
	}

	LatestTime, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if LatestTime <= 0 {
		LatestTime = time.Now().Unix()
	}

	StatusCode++
	resp.StatusCode = StatusCode
	c.JSON(http.StatusOK, FeedResponse{
		Response:  resp,
		VideoList: FeedVideos,
		NextTime:  LatestTime,
	})
}
