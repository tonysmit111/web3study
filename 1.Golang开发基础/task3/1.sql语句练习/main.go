package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// question1()
	question2()
}

type Student struct {
	gorm.Model
	Name  string
	Age   uint32
	Grade string
}

type Account struct {
	gorm.Model
	Balance float32
}

type Transaction struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount float32
}

var db *gorm.DB

func init() {
	fmt.Println("init ...")
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	// db.Delete(&Account{}, "1=1")
	as := []Account{}
	db.Find(&as)
	if len(as) == 0 {
		as = []Account{
			{Balance: 100},
			{Balance: 100},
		}
		db.Create(&as)
	} else {
		db.Model(&as).Updates(Account{Balance: float32(100)})
	}

}

/*
题目1：基本CRUD操作

假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

func question1() {
	// insert(&Student{
	// 	Name:"张三",
	// 	Age:10,
	// 	Grade:"三年级",
	// })
	// db.Create(&Student{
	// 	Name : "李四",
	// 	Age: 11,
	// 	Grade: "三年级",
	// })
	var stus []Student
	db.Find(&stus, "age>=?", 10)
	fmt.Println(stus)

	var stus2 []Student
	db.Find(&stus2, "name=?", "张三")
	db.Model(&stus2).Update("Grade", "四年级")

	// db.Where("age<?",16).Delete(&Student{})
	db.Delete(Student{}, "age<?", 16)
}

func insert(stu *Student) {
	tx := db.Exec("insert into students(name, age, grade) values(?,?,?)", stu.Name, stu.Age, stu.Grade)
	// tx := db.Exec("insert into students(name, age, grade) values(?,?,?)",stu)
	r := tx.RowsAffected
	fmt.Println(r)
}

/*
题目2：事务语句

假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
func question2() {
	// trans(3, 4, 20, false)
	// err := trans(3, 4, 100, true)
	err := trans(3, 4, 200, false)
	if err != nil {
		fmt.Println(err)
	}
}

func trans(id1, id2 uint, m float32, hasError bool) error {
	var a1, a2 Account
	db.First(&a1, id1)
	db.First(&a2, id2)

	return db.Transaction(func(tx *gorm.DB) error {
		if a1.Balance < m{
			return errors.New("余额不足")
		}
		if err := tx.Model(&a1).Update("balance", a1.Balance-m).Error; err != nil {
			return err
		}
		if hasError {
			return errors.New("转账出现异常")
		}
		if err := tx.Model(&a2).Update("balance", a2.Balance+m).Error; err != nil {
			return err
		}
		trans := Transaction{
			FromAccountId: a1.ID,
			ToAccountId: a2.ID,
			Amount: m,
		}
		if err := tx.Create(&trans).Error;err!=nil {
			return err
		}
		return nil
	})
}
