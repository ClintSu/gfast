// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package cms_news

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table cms_news.
type Entity struct {
	Id            uint64 `orm:"id,primary"     json:"id"`             //
	UserId        uint64 `orm:"user_id"        json:"user_id"`        // 发表者用户id
	NewsStatus    uint   `orm:"news_status"    json:"news_status"`    // 状态;1:已发布;0:未发布;
	IsTop         uint   `orm:"is_top"         json:"is_top"`         // 是否置顶;1:置顶;0:不置顶
	Recommended   uint   `orm:"recommended"    json:"recommended"`    // 是否推荐;1:推荐;0:不推荐
	IsSlide       uint   `orm:"is_slide"    json:"is_slide"`          // 是否幻灯;1:是;0:否
	NewsHits      uint64 `orm:"news_hits"      json:"news_hits"`      // 查看数
	NewsLike      uint64 `orm:"news_like"      json:"news_like"`      // 点赞数
	CreateTime    uint   `orm:"create_time"    json:"create_time"`    // 创建时间
	UpdateTime    uint   `orm:"update_time"    json:"update_time"`    // 更新时间
	PublishedTime uint   `orm:"published_time" json:"published_time"` // 发布时间
	DeleteTime    uint   `orm:"delete_time"    json:"delete_time"`    // 删除时间
	NewsTitle     string `orm:"news_title"     json:"news_title"`     // post标题
	NewsKeywords  string `orm:"news_keywords"  json:"news_keywords"`  // seo keywords
	NewsExcerpt   string `orm:"news_excerpt"   json:"news_excerpt"`   // post摘要
	NewsSource    string `orm:"news_source"    json:"news_source"`    // 转载文章的来源
	Thumbnail     string `orm:"thumbnail"      json:"thumbnail"`      // 缩略图
	IsJump        uint   `orm:"is_jump"        json:"is_jump"`        // 是否跳转地址
	JumpUrl       string `orm:"jump_url"       json:"jump_url"`       // 跳转地址
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}
