package role

import (
	"errors"
	"strconv"

	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func (r *SysRoleHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	role, err := r.srv.SysRole().GetById(uint64(id))
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to get role"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, role)
}

func (r *SysRoleHandler) GetList(c *gin.Context) {
	var param model.SysRole
	err := c.ShouldBindJSON(&param)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
		return
	}
	whereOrders := createSysRoleQueryCondition(param)
	roles, err := r.srv.SysRole().GetList(whereOrders...)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to get roles"), nil)
		return
	}
	util.WriteResponse(c, 200, nil, roles)
}

func (r *SysRoleHandler) GetPage(c *gin.Context) {
	var param model.SysRolePage
	err := c.ShouldBindJSON(&param)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to bind param"), nil)
		return
	}
	whereOrders := createSysRoleQueryCondition(param.SysRole)
	roles, count, err := r.srv.SysRole().GetPage(param.PageIndex, param.PageSize, whereOrders...)
	if err != nil {
		util.WriteResponse(c, 500, errors.New("failed to get role page"), nil)
		return
	}

	page := &model.Page{
		Records:  roles,
		PageInfo: model.PageInfo{PageIndex: param.PageIndex, PageSize: param.PageSize},
	}
	page.SetPageNum(count)
	util.WriteResponse(c, 200, nil, page)
}

func createSysRoleQueryCondition(param model.SysRole) []model.WhereOrder {
	whereOrders := make([]model.WhereOrder, 0)

	if param.Name != "" {
		v := "%" + param.Name + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "name like ?", Value: []interface{}{v}})
	}
	if param.NameZh != "" {
		v := "%" + param.NameZh + "%"
		whereOrders = append(whereOrders, model.WhereOrder{Where: "name_zh like ?", Value: []interface{}{v}})
	}

	return whereOrders
}
