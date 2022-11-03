package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("unauthorized to access this resource. user_type dosn't match")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "User" && uid != userId {
		err = errors.New("unauthorized to access this resource match user. user_type dosn't match to uid")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
