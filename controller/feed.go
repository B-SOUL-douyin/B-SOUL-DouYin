package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"

	"github.com/RaymondCode/simple-demo/video/dal/db"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var FeedVideos []Video
	FeedVideos, err := db.GetFeedList()

	resp := Response{0, ""}
	if err != nil {
		resp.StatusCode++
		resp.StatusMsg = resp.StatusMsg + ";" + err.Error()
	}
	if len(FeedVideos) == 0 {
		resp.StatusCode++
		resp.StatusMsg = resp.StatusMsg + "没有新的视频了"
	}

	LatestTime, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if LatestTime <= 0 {
		LatestTime = time.Now().Unix()
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  resp,
		VideoList: FeedVideos,
		NextTime:  LatestTime,
	})
}
