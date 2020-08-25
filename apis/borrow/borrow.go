package borrow

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"go-admin/tools/app/msg"
)

func InsetBorrowImg(c *gin.Context) {
	file, _ := c.FormFile("file")
	guid := uuid.New().String()
	filePath := "static/uploadfile/" + guid + ".jpg"
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		app.OK(c, filePath, "上传失败")
		//自己完成信息提示
		return
	}
	app.OK(c, filePath, "上传成功")
}
func InsertBorrowProve(c *gin.Context) {
	file, _ := c.FormFile("file")
	guid := uuid.New().String()
	filePath := "static/uploadfile/" + guid + ".pdf"
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		app.OK(c, filePath, "上传失败")
		//自己完成信息提示
		return
	}
	app.OK(c, filePath, "上传成功")
}
func GetBorrowList(c *gin.Context) {
	var data models.Borrow
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.BorrowName = c.Request.FormValue("borrowName")
	data.InterestRate = c.Request.FormValue("interestRate")
	data.Term = c.Request.FormValue("term")
	data.TermType = c.Request.FormValue("termType")
	data.Amount = c.Request.FormValue("amount")
	data.RepayType = c.Request.FormValue("repayType")
	data.Purpose = c.Request.FormValue("purpose")
	data.RepayName = c.Request.FormValue("repayName")
	data.Diya = c.Request.FormValue("diya")
	data.AmountLimit = c.Request.FormValue("amountLimit")
	data.BorrowStatus = c.Request.FormValue("borrowStatus")

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetBorrow(c *gin.Context) {
	var data models.Borrow
	data.Id, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertBorrow(c *gin.Context) {

	var data models.Borrow
	err := c.ShouldBindJSON(&data)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateBorrow(c *gin.Context) {
	var data models.Borrow
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.Id)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteBorrow(c *gin.Context) {
	var data models.Borrow
	data.UpdateBy = tools.GetUserIdStr(c)

	IDS := tools.IdsStrToIdsIntGroup("id", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}
