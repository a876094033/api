package borrowdiya

import (
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"go-admin/tools/app/msg"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetBorrowDiyaList(c *gin.Context) {
	var data models.BorrowDiya
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.DiyaId, _ = tools.StringToInt(c.Request.FormValue("diyaId"))
	data.BorrowId = c.Request.FormValue("borrowId")

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetBorrowDiya(c *gin.Context) {
	var data models.BorrowDiya
	data.DiyaId, _ = tools.StringToInt(c.Param("diyaId"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertBorrowDiya(c *gin.Context) {
	var data models.BorrowDiya
	err := c.ShouldBindJSON(&data)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateBorrowDiya(c *gin.Context) {
	var data models.BorrowDiya
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.DiyaId)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteBorrowDiya(c *gin.Context) {
	var data models.BorrowDiya
	data.UpdateBy = tools.GetUserIdStr(c)

	IDS := tools.IdsStrToIdsIntGroup("diyaId", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}
