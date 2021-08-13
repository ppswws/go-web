package menu

import (
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/response"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

/*
	DELETE: /v1/menu/delete
*/
func (m *SysMenuHandler) BatchDelete(c *gin.Context) {

	var param model.IdParam
	err := c.ShouldBind(&param)
	if err != nil {
		response.FailWithCode(response.ParameterBindingError)
		return
	}
	ids := util.Str2Uint64Array(param.Ids)
	err = m.srv.SysMenu().BatchDelete(ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	response.Success()
}