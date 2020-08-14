package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type Invest struct {
	InvestId     int    `json:"investId" gorm:"type:int(11);primary_key"` // 投资id
	MamberId     string `json:"mamberId" gorm:"type:int(11);"`            // 用户id
	InvestAmount string `json:"investAmount" gorm:"type:decimal(18,2);"`  // 投资金额
	BorrowId     string `json:"borrowId" gorm:"type:int(11);"`            // 借款id
	CreateBy     string `json:"createBy" gorm:"type:varchar(100);"`       //
	UpdateBy     string `json:"updateBy" gorm:"type:varchar(100);"`       //
	DataScope    string `json:"dataScope" gorm:"-"`
	Params       string `json:"params"  gorm:"-"`
	BaseModel
}

func (Invest) TableName() string {
	return "invest"
}

// 创建Invest
func (e *Invest) Create() (Invest, error) {
	var doc Invest
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取Invest
func (e *Invest) Get() (Invest, error) {
	var doc Invest
	table := orm.Eloquent.Table(e.TableName())

	if e.InvestId != 0 {
		table = table.Where("invest_id = ?", e.InvestId)
	}

	if e.MamberId != "" {
		table = table.Where("mamber_id = ?", e.MamberId)
	}

	if e.InvestAmount != "" {
		table = table.Where("invest_amount = ?", e.InvestAmount)
	}

	if e.BorrowId != "" {
		table = table.Where("borrow_id = ?", e.BorrowId)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取Invest带分页
func (e *Invest) GetPage(pageSize int, pageIndex int) ([]Invest, int, error) {
	var doc []Invest
	
	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.InvestId != 0 {
		table = table.Where("invest_id = ?", e.InvestId)
	}

	if e.MamberId != "" {
		table = table.Where("mamber_id = ?", e.MamberId)
	}

	if e.InvestAmount != "" {
		table = table.Where("invest_amount = ?", e.InvestAmount)
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

// 更新Invest
func (e *Invest) Update(id int) (update Invest, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("invest_id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除Invest
func (e *Invest) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("invest_id = ?", id).Delete(&Invest{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *Invest) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("invest_id in (?)", id).Delete(&Invest{}).Error; err != nil {
		return
	}
	Result = true
	return
}
