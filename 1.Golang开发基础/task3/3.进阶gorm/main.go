package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 题目1：模型定义

// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	gorm.Model
	Name    string
	PostCnt uint
	Posts   []Post
}

type Post struct {
	gorm.Model
	UserId        uint
	Title         string
	CommentCnt    uint
	CommentStatus string
	Comments      []Comment
}

type Comment struct {
	gorm.Model
	PostId      uint
	Description string
}

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
}

func question1() {
	db.Unscoped().Where("1=1").Delete(&User{})
	users := []User{
		{
			Name: "tom",
		},
		{

			Name: "jack",
		},
	}
	db.Create(&users)

	db.Unscoped().Where("1=1").Delete(&Post{})
	posts := []Post{
		{
			UserId: users[0].ID,
			Title:  "tom-文章一",
		},
		{
			UserId: users[0].ID,
			Title:  "tom-文章二",
		},
		{
			UserId: users[1].ID,
			Title:  "jack-文章一",
		},
	}
	db.Create(&posts)

	db.Unscoped().Where("1=1").Delete(&Comment{})
	comments := []Comment{
		{PostId: posts[0].ID, Description: "tom-文章一评论一"},
		{PostId: posts[0].ID, Description: "tom-文章一评论二"},
		{PostId: posts[1].ID, Description: "tom-文章二评论一"},
		{PostId: posts[2].ID, Description: "jack-文章一评论一"},
	}
	db.Create(&comments)

}

// 题目2：关联查询

// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

func question2() {
	users := []User{}
	// err = db.Joins("JOIN posts ON users.id = posts.user_id ").Joins("JOIN comments ON posts.id = comments.post_id").Find(&users).Error
	err = db.Preload("Posts.Comments").Where("name = ?", "tom").Find(&users).Error
	// err = db.Joins("Posts.Comments").Where("name = ?", "tom").Find(&users).Error
	if err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	fmt.Println("---------------------")
	ps := []Post{}
	db.Model(&Post{}).
		Select("posts.*", "count(1) as cnt").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Limit(1).
		Order("cnt desc").
		Preload("Comments").
		Find(&ps)

	pb, err := json.Marshal(&ps)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pb))

}

// 题目3：钩子函数

// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("post hook after create method is excuting")
	return tx.Model(&User{}).Where("id=?", p.UserId).Update("post_cnt", gorm.Expr("IFNULL(post_cnt,0) + ?", 1)).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("comment hook after delete method is excuting")
	var cnt int64
	p := &Post{}
	tx.Model(&Comment{}).Where("post_id=?", c.PostId).Count(&cnt)
	if cnt == 0 {
		return tx.Model(&p).Where("id = ?", c.PostId).Updates(&Post{CommentCnt: 0, CommentStatus: "无评论"}).Error
	} else {
		return tx.Model(&p).Where("id = ?", c.PostId).Update("comment_cnt", cnt).Error
	}

}

func question3() {
	u := User{}
	db.First(&u)
	posts := []Post{
		{
			UserId: u.ID,
			Title:  "tom-文章一",
		},
		{
			UserId: u.ID,
			Title:  "tom-文章二",
		},
	}
	db.Create(&posts)

	c := Comment{}
	db.Model(&c).Where("id=?", 22).First(&c)
	db.Unscoped().Where("id=?", 22).Delete(&c)
}

func main() {
	question3()
}
