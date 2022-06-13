package db

import (
	"github.com/B-SOUL-douyin/B-SOUL-DouYin/pkg/constants"
	"gorm.io/gorm"
)

type VideoModel struct {
	gorm.Model
	Video constants.Video `gorm:"embedded"`
}

func (v *VideoModel) TableName() string {
	return "videos"
}

// TransVideoModel transform video type to video model
func TransVideoModel(OriginVideo []constants.Video) []VideoModel {
	var VideoModel []VideoModel
	for i, v := range OriginVideo {
		VideoModel[i].Video = v
	}
	return VideoModel
}

// CreateVideo create video
func CreateVideo(Videos []VideoModel) error {
	if err := DB.Create(Videos).Error; err != nil {
		return err
	}
	return nil
}

// GetFeedList Get video feed list
func GetFeedList() ([]VideoModel, error) {
	var Videos []VideoModel

	err := DB.Order("ID desc").Limit(10).Find(&Videos).Error
	// SELECT * FROM videos ORDER BY ID desc LIMIT 10;
	if err != nil {
		return Videos, err
	}
	// DEBUG
	//log.Println(len(Videos))
	return Videos, nil
}

// TODO:Update video
// TODO:Delete video
// TODO:视频去重
