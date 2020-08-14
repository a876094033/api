package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type Area struct {
	AreaId    int    `json:"areaId" gorm:"type:int(11);primary_key"` //
	AreaName  string `json:"areaName" gorm:"type:varchar(100);"`     // 区域名称
	ParentId  string `json:"parentId" gorm:"type:varchar(100);"`     // 上级区域
	LevelId   string `json:"levelId" gorm:"type:varchar(100);"`      // 区域级别
	Postcode  string `json:"postcode" gorm:"type:varchar(100);"`     // 邮编
	CreateBy  string `json:"createBy" gorm:"type:varchar(100);"`     //
	UpdateBy  string `json:"updateBy" gorm:"type:varchar(100);"`     //
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (Area) TableName() string {
	return "area"
}

// 创建Area
func (e *Area) Create() (Area, error) {
	var doc Area
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取Area
func (e *Area) Get() (Area, error) {
	var doc Area
	table := orm.Eloquent.Table(e.TableName())

	if e.AreaId != 0 {
		table = table.Where("area_id = ?", e.AreaId)
	}

	if e.AreaName != "" {
		table = table.Where("area_name like ?", "%"+e.AreaName+"%")
	}

	if e.ParentId != "" {
		table = table.Where("parent_id = ?", e.ParentId)
	}

	if e.LevelId != "" {
		table = table.Where("level_id = ?", e.LevelId)
	}

	if e.Postcode != "" {
		table = table.Where("postcode = ?", e.Postcode)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取Area带分页
func (e *Area) GetPage(pageSize int, pageIndex int) ([]Area, int, error) {
	var doc []Area

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.AreaId != 0 {
		table = table.Where("area_id = ?", e.AreaId)
	}

	if e.AreaName != "" {
		table = table.Where("area_name like ?", "%"+e.AreaName+"%")
	}

	if e.ParentId != "" {
		table = table.Where("parent_id = ?", e.ParentId)
	}

	if e.LevelId != "" {
		table = table.Where("level_id = ?", e.LevelId)
	}

	if e.Postcode != "" {
		table = table.Where("postcode = ?", e.Postcode)
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

// 更新Area
func (e *Area) Update(id int) (update Area, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("area_id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除Area
func (e *Area) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("area_id = ?", id).Delete(&Area{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *Area) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("area_id in (?)", id).Delete(&Area{}).Error; err != nil {
		return
	}
	Result = true
	return
}
