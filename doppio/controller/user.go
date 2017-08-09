package controller

import (
	"database/sql"
	//"errors"
	"github.com/gin-gonic/gin"
	"github.com/yuk-vg/vg-1day-2017-08-09/doppio/httputil"
	"github.com/yuk-vg/vg-1day-2017-08-09/doppio/model"
	"net/http"
)

type User struct {
	DB     *sql.DB
	Stream chan *model.User
}

// All は全てのメッセージを取得してJSONで返します
func (m *User) All(c *gin.Context) {
	usrs, err := model.UsersAll(m.DB)
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

// GetByID はパラメーターで受け取ったidのメッセージを取得してJSONで返します
func (m *User) GetByID(c *gin.Context) {
	usr, err := model.UserByID(m.DB, c.Param("id"))

	switch {
	case err == sql.ErrNoRows:
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusNotFound, resp)
		return
	case err != nil:
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": usr,
		"error":  nil,
	})
}
