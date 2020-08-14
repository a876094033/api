package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type Members struct {
	Id           int    `json:"id" gorm:"type:int(11);primary_key"`      // id
	Name         string `json:"name" gorm:"type:varchar(100);"`          // 用户名
	Password     string `json:"password" gorm:"type:varchar(100);"`      // 密码
	Phone        string `json:"phone" gorm:"type:varchar(100);"`         // 手机号
	Email        string `json:"email" gorm:"type:varchar(100);"`         // 邮箱
	RealName     string `json:"realName" gorm:"type:varchar(100);"`      // 真实姓名
	FirstCardNum string `json:"firstCardNum" gorm:"type:varchar(100);"`  // 身份证首段
	LastCardNum  string `json:"lastCardNum" gorm:"type:varchar(100);"`   // 身份证尾号
	AmountAll    string `json:"amountAll" gorm:"type:decimal(18,2);"`    // 总金额
	AmountFrozen string `json:"amountFrozen" gorm:"type:decimal(10,0);"` // 冻结金额
	AmountUsed   string `json:"amountUsed" gorm:"type:decimal(10,0);"`   // 可用金额
	CreateBy     string `json:"createBy" gorm:"type:varchar(100);"`      // 注册人
	UpdateBy     string `json:"updateBy" gorm:"type:varchar(100);"`      // 修改人
	DataScope    string `json:"dataScope" gorm:"-"`
	Params       string `json:"params"  gorm:"-"`
	BaseModel
}

func (Members) TableName() string {
	return "members"
}

// 创建Members
func (e *Members) Create() (Members, error) {
	var doc Members
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取Members
func (e *Members) Get() (Members, error) {
	var doc Members
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if e.Password != "" {
		table = table.Where("password = ?", e.Password)
	}

	if e.Phone != "" {
		table = table.Where("phone = ?", e.Phone)
	}

	if e.Email != "" {
		table = table.Where("email = ?", e.Email)
	}

	if e.RealName != "" {
		table = table.Where("real_name = ?", e.RealName)
	}

	if e.FirstCardNum != "" {
		table = table.Where("first_card_num = ?", e.FirstCardNum)
	}

	if e.LastCardNum != "" {
		table = table.Where("last_card_num = ?", e.LastCardNum)
	}

	if e.AmountAll != "" {
		table = table.Where("amount_all = ?", e.AmountAll)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取Members带分页
func (e *Members) GetPage(pageSize int, pageIndex int) ([]Members, int, error) {
	var doc []Members

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if e.Password != "" {
		table = table.Where("password = ?", e.Password)
	}

	if e.Phone != "" {
		table = table.Where("phone = ?", e.Phone)
	}

	if e.Email != "" {
		table = table.Where("email = ?", e.Email)
	}

	if e.RealName != "" {
		table = table.Where("real_name = ?", e.RealName)
	}

	if e.FirstCardNum != "" {
		table = table.Where("first_card_num = ?", e.FirstCardNum)
	}

	if e.LastCardNum != "" {
		table = table.Where("last_card_num = ?", e.LastCardNum)
	}

	if e.AmountAll != "" {
		table = table.Where("amount_all = ?", e.AmountAll)
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

// 更新Members
func (e *Members) Update(id int) (update Members, err error) {
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

// 删除Members
func (e *Members) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&Members{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *Members) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&Members{}).Error; err != nil {
		return
	}
	Result = true
	return
}
