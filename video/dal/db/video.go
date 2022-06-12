package db

import (
	"github.com/RaymondCode/simple-demo/controller"
)

type Video = controller.Video

// CreateVideo create video
func CreateVideo(Videos []Video) error {
	if err := DB.Create(Videos).Error; err != nil {
		return err
	}
	return nil
}

// GetFeedList Get video feed list
func GetFeedList() ([]Video, error) {
	var Videos []Video

	err := DB.Order("ID desc").Limit(10).First(&Videos).Error
	// SELECT * FROM Video ORDER BY ID desc LIMIT 3;
	if err != nil {
		return Videos, err
	}
	return Videos, nil
}

// TODO:Update video
// TODO:Delete video
