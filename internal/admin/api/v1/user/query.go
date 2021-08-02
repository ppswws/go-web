package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
)

//查询
func (u *UserHandler) GetByUsername(c *gin.Context) {

	user, err := u.srv.SysUser().GetByUsername(c.Param("name"))
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}
	util.WriteResponse(c, 0, nil, user)

}

//查询多条记录，参数为json格式
func (u *UserHandler) List(c *gin.Context) {
	var param model.SysUser
	// 此处不能传入空指针，否则绑定失败
	err := c.ShouldBindJSON(&param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	list, err := u.srv.SysUser().List(&param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	util.WriteResponse(c, 0, nil, list)
}

func (u *UserHandler) GetPage(c *gin.Context) {
	var param model.SysUserPage
	err := c.ShouldBindJSON(&param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	list, count, err := u.srv.SysUser().GetPage(&param)
	if err != nil {
		util.WriteResponse(c, http.StatusInternalServerError, err, nil)
		return
	}

	page := &model.Page{
		Records:  list,
		PageInfo: model.PageInfo{PageIndex: param.PageIndex, PageSize: param.PageSize},
	}
	page.SetPageNum(count)
	util.WriteResponse(c, 0, nil, page)
}

// 使用go-jwt授权
func (u *UserHandler) Login(c *gin.Context) (interface{}, error) {
	var param model.SysUser
	err := c.ShouldBindJSON(&param)
	if err != nil {
		return nil, err
	}

	user, err := u.srv.SysUser().Login(param.Username, param.Password)

	if err != nil || user == nil {
		return nil, err
	}

	return map[string]interface{}{
		"user": fmt.Sprintf("%d", user.Id),
	}, nil
}
