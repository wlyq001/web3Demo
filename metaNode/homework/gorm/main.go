package main

import (
	//"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id         uint `gorm:"primaryKey"`
	Name       string
	Age        uint8
	Grade      string
	PostAmount int64
}

type Post struct {
	Id         uint `gorm:"primaryKey"`
	Title      string
	Body       string
	UserId     uint
	HasComment bool
	User       User
}

func (post *Post) AfterSave(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Post{}).Where(&Post{UserId: post.UserId}).Count(&count)
	tx.Model(&User{}).Where(&User{Id: post.UserId}).Update("PostAmount", count)
	return nil
}

func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Comment{}).Where(&Comment{PostId: comment.PostId}).Count(&count)
	if count > 0 {
		return nil
	}
	tx.Model(&Post{}).Where(&Post{Id: comment.PostId}).Update("HasComment", false)
	return nil
}

type Comment struct {
	Id      uint `gorm:"primaryKey"`
	Comment string
	PostId  uint
	Post    Post
}

type Account struct {
	Id      uint `gorm:"primaryKey"`
	Balance uint
}

type Transaction struct {
	Id            uint `gorm:"primaryKey"`
	FromAccountId uint
	ToAccountId   uint
	Amount        uint
	Account1      Account `gorm:"foreignKey:FromAccountId"`
	Account2      Account `gorm:"foreignKey:ToAccountId"`
}

func main() {
	dsn := "root:wl980805@tcp(127.0.0.1:3306)/web3_learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db.Debug().AutoMigrate(&User{})

	//userZS := User{
	//	Name:  "李四",
	//	Age:   13,
	//	Grade: "一年级",
	//}
	//
	//db.Create(&userZS)

	//result := []User{}
	//db.Model(&User{}).Where("age > ?", 20).Find(&User{}).Scan(&result)
	//fmt.Println(result)

	//db.Model(&User{}).Where("name = ?", "张三").Update("grade", "四年级")

	//db.Model(&User{}).Where("age < ?", 15).Delete(&User{})

	//db.AutoMigrate(&Account{}, &Transaction{})

	//db.Transaction(func(tx *gorm.DB) error {
	//	var accountA Account
	//	tx.Model(&Account{}).Where("id = 3").Find(&accountA)
	//	if accountA.Balance < 100 {
	//		return errors.New("balance is less than 100")
	//	}
	//	var accountB Account
	//	tx.Model(&Account{}).Where("id = 2").Find(&accountB)
	//	accountA.Balance -= 100
	//	accountB.Balance += 100
	//	updatesA := tx.Model(&Account{}).Updates(accountA)
	//	updatesB := tx.Model(&accountB).Updates(accountB)
	//
	//	if updatesA.Error != nil || updatesB.Error != nil {
	//		return updatesA.Error
	//	}
	//
	//	transaction := Transaction{
	//		FromAccountId: accountA.Id,
	//		ToAccountId:   accountB.Id,
	//		Amount:        100,
	//	}
	//	createT := tx.Create(&transaction)
	//	if createT.Error != nil {
	//		return createT.Error
	//	}
	//	return nil
	//})

	//db.AutoMigrate(&User{}, &Post{}, &Comment{})
	//
	//user := User{
	//	Name:  "wlyq",
	//	Age:   27,
	//	Grade: "graduated",
	//}
	//db.Create(&user)
	//
	//post := Post{
	//	Title:  "无敌",
	//	Body:   "alsdjaoijdoaw",
	//	UserId: user.Id,
	//}
	//db.Create(&post)
	//
	//comment := Comment{
	//	Comment: "aaaa",
	//	PostId:  post.Id,
	//}
	//db.Create(&comment)

	//var comments []Comment
	//db.Debug().Joins("Post").Where(&Comment{Post: Post{UserId: 1}}).Find(&comments)
	//fmt.Println(comments)

	//var posts []Post
	//db.Where(&Post{UserId: 1}).Find(&posts)
	//infos := []PostInfo{}
	//for _, post := range posts {
	//	var comments []Comment
	//	db.Where(&Comment{PostId: post.Id}).Find(&comments)
	//	info := PostInfo{post: post, comments: comments}
	//	infos = append(infos, info)
	//}
	//fmt.Println(infos)
	//var commentCounts []map[string]interface{}
	//db.Model(&Comment{}).Select("post_id ,count(*) as count").Group("post_id").Find(&commentCounts)
	//fmt.Println(commentCounts)
	//var biggest int64 = 0
	//var finalPostId uint
	//for _, comment := range commentCounts {
	//	if comment["count"].(int64) > biggest {
	//		biggest = comment["count"].(int64)
	//		finalPostId = comment["post_id"].(uint)
	//	}
	//}
	//var finalPost Post
	//db.Model(&Post{}).Where(&Post{Id: finalPostId}).Find(&finalPost)
	//fmt.Println(finalPost)

	//post := &Post{Title: "woshishen", Body: "tiantian", UserId: 1}
	//db.Create(post)

	db.Where(&Comment{PostId: 1}).Delete(&Comment{PostId: 1})

}

type PostInfo struct {
	post     Post
	comments []Comment
}
