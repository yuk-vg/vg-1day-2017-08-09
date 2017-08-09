package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuk-vg/vg-1day-2017-08-09/yukimine/httputil"
	"github.com/yuk-vg/vg-1day-2017-08-09/yukimine/model"
)

// User is controller for requests to user
type User struct {
	DB     *sql.DB
	Stream chan *model.User
}

// All は全てのメッセージを取得してJSONで返します
func (u *User) All(c *gin.Context) {
	usrs, err := model.UserAll(u.DB)
	if err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(usrs) == 0 {
		c.JSON(http.StatusOK, make([]*model.User, 0))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": usrs,
		"error":  nil,
	})
}
