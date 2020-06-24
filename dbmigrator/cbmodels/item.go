package cbmodels

type Item struct {
	ItemID uint   `gorm:"primary_key;column:idx"`
	Table  string `gorm:"type:varchar(50);column:bo_table"`
	Column string `gorm:"type:varchar(50);column:tcolumn"`
	Item   string `gorm:"type:varchar(50);column:titem"`
	Desc   int    `gorm:"column:tdesc"`
}

func (Item) TableName() string {
	return "bo_item"
}

type SubItem struct {
	SubItemID uint   `gorm:"primary_key;column:sidx"`
	Item      Item   `gorm:"foreignkey:ItemID"`
	ItemID    uint   `gorm:"column:idx"`
	SubItem   string `gorm:"type:varchar(200);column:stitem"`
	Desc      int    `gorm:"column:tdesc"`
}

func (SubItem) TableName() string {
	return "bo_item_sub"
}
