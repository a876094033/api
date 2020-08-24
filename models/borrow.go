package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type Borrow struct {
	Id           int    `json:"id" gorm:"type:int(11);primary_key"`      //
	BorrowName   string `json:"borrowName" gorm:"type:varchar(100);"`    // 借款名称
	InterestRate string `json:"interestRate" gorm:"type:decimal(10,2);"` // 利率
	Term         string `json:"term" gorm:"type:int(11);"`               // 借款期限
	TermType     string `json:"termType" gorm:"type:int(11);"`           // 日期类型
	Amount       string `json:"amount" gorm:"type:decimal(18,2);"`       // 借款金额
	RepayType    string `json:"repayType" gorm:"type:int(11);"`          // 还款方式
	Purpose      string `json:"purpose" gorm:"type:varchar(100);"`       // 借款目的
	RepayName    string `json:"repayName" gorm:"type:varchar(100);"`     // 还款人
	Diya         string `json:"diya" gorm:"type:varchar(100);"`          // 抵押
	AmountLimit  string `json:"amountLimit" gorm:"type:decimal(18,2);"`  // 投资限额
	BorrowStatus string `json:"borrowStatus" gorm:"type:int(11);"`       // 借款状态
	BorrowImg    string `json:"borrowImg" gorm:"_"`     // 借款状态
	CreateBy     string `json:"createBy" gorm:"type:varchar(100);"`      // 创建人
	UpdateBy     string `json:"updateBy" gorm:"type:varchar(100);"`      //
	DataScope    string `json:"dataScope" gorm:"-"`
	Params       string `json:"params"  gorm:"-"`
	BaseModel
}

func (Borrow) TableName() string {
	return "borrow"
}

// 创建Borrow
func (e *Borrow) Create() (Borrow, error) {
	var doc Borrow
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取Borrow
func (e *Borrow) Get() (Borrow, error) {
	var doc Borrow
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.BorrowName != "" {
		table = table.Where("borrow_name = ?", e.BorrowName)
	}

	if e.InterestRate != "" {
		table = table.Where("interest_rate = ?", e.InterestRate)
	}

	if e.Term != "" {
		table = table.Where("term = ?", e.Term)
	}

	if e.TermType != "" {
		table = table.Where("term_type = ?", e.TermType)
	}

	if e.Amount != "" {
		table = table.Where("amount = ?", e.Amount)
	}

	if e.RepayType != "" {
		table = table.Where("repay_type = ?", e.RepayType)
	}

	if e.Purpose != "" {
		table = table.Where("purpose = ?", e.Purpose)
	}

	if e.RepayName != "" {
		table = table.Where("repay_name = ?", e.RepayName)
	}

	if e.Diya != "" {
		table = table.Where("diya = ?", e.Diya)
	}

	if e.AmountLimit != "" {
		table = table.Where("amount_limit = ?", e.AmountLimit)
	}

	if e.BorrowStatus != "" {
		table = table.Where("borrow_status = ?", e.BorrowStatus)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取Borrow带分页
func (e *Borrow) GetPage(pageSize int, pageIndex int) ([]Borrow, int, error) {
	var doc []Borrow

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.BorrowName != "" {
		table = table.Where("borrow_name = ?", e.BorrowName)
	}

	if e.InterestRate != "" {
		table = table.Where("interest_rate = ?", e.InterestRate)
	}

	if e.Term != "" {
		table = table.Where("term = ?", e.Term)
	}

	if e.TermType != "" {
		table = table.Where("term_type = ?", e.TermType)
	}

	if e.Amount != "" {
		table = table.Where("amount = ?", e.Amount)
	}

	if e.RepayType != "" {
		table = table.Where("repay_type = ?", e.RepayType)
	}

	if e.Purpose != "" {
		table = table.Where("purpose = ?", e.Purpose)
	}

	if e.RepayName != "" {
		table = table.Where("repay_name = ?", e.RepayName)
	}

	if e.Diya != "" {
		table = table.Where("diya = ?", e.Diya)
	}

	if e.AmountLimit != "" {
		table = table.Where("amount_limit = ?", e.AmountLimit)
	}

	if e.BorrowStatus != "" {
		table = table.Where("borrow_status = ?", e.BorrowStatus)
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

// 更新Borrow
func (e *Borrow) Update(id int) (update Borrow, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除Borrow
func (e *Borrow) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&Borrow{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *Borrow) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&Borrow{}).Error; err != nil {
		return
	}
	Result = true
	return
}
