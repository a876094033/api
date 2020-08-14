package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type BorrowDiya struct {
	DiyaId    int    `json:"diyaId" gorm:"type:int(11);primary_key"` // ID
	BorrowId  string `json:"borrowId" gorm:"type:int(11);"`          //
	Address   string `json:"address" gorm:"type:varchar(100);"`      // 地址
	Amount    string `json:"amount" gorm:"type:varchar(100);"`       // 市场价格
	Size      string `json:"size" gorm:"type:varchar(100);"`         // 大小
	Ltv       string `json:"ltv" gorm:"type:varchar(100);"`          // LTV
	CreateBy  string `json:"createBy" gorm:"type:varchar(100);"`     //
	UpdateBy  string `json:"updateBy" gorm:"type:varchar(100);"`     //
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (BorrowDiya) TableName() string {
	return "borrow_diya"
}

// 创建BorrowDiya
func (e *BorrowDiya) Create() (BorrowDiya, error) {
	var doc BorrowDiya
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取BorrowDiya
func (e *BorrowDiya) Get() (BorrowDiya, error) {
	var doc BorrowDiya
	table := orm.Eloquent.Table(e.TableName())

	if e.DiyaId != 0 {
		table = table.Where("diya_id = ?", e.DiyaId)
	}

	if e.BorrowId != "" {
		table = table.Where("borrow_id = ?", e.BorrowId)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取BorrowDiya带分页
func (e *BorrowDiya) GetPage(pageSize int, pageIndex int) ([]BorrowDiya, int, error) {
	var doc []BorrowDiya

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.DiyaId != 0 {
		table = table.Where("diya_id = ?", e.DiyaId)
	}

	if e.BorrowId != "" {
		table = table.Where("borrow_id = ?", e.BorrowId)
	}

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	dataPermission := new(DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

// 更新BorrowDiya
func (e *BorrowDiya) Update(id int) (update BorrowDiya, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("diya_id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除BorrowDiya
func (e *BorrowDiya) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("diya_id = ?", id).Delete(&BorrowDiya{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *BorrowDiya) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("diya_id in (?)", id).Delete(&BorrowDiya{}).Error; err != nil {
		return
	}
	Result = true
	return
}
