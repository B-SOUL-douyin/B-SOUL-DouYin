package controller

import (
	"github.com/B-SOUL-douyin/B-SOUL-DouYin/pkg/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	constants.Response
	UserList []constants.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, constants.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, constants.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: constants.Response{
			StatusCode: 0,
		},
		UserList: []constants.User{constants.DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: constants.Response{
			StatusCode: 0,
		},
		UserList: []constants.User{constants.DemoUser},
	})
}
