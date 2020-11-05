package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	//gorm本身做了大量的约定俗成
	//gorm会通过类名+ID的方式去找到这两个类之间的关联关系
	//Article有一个结构体成员是Tag，就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

// Preload就是一个预加载器，它会执行两条SQL，分别是
// SELECT * FROM blog_articles;和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);，
// 那么在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

//gorm本身做了大量的约定俗成
//Article有一个结构体成员是TagID，就是外键。gorm会通过类名+ID的方式去找到这两个类之间的关联关系
//Article有一个结构体成员是Tag，就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

// 注意硬删除要使用 Unscoped()，这是 GORM 的约定
func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}
