package main

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)




var db *sql.DB
var dbx *sqlx.DB
var err error

func init() {
	db, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Print(err)
		panic("db init failed")
	}
	dbx = sqlx.NewDb(db, "sqlite3")
}

// 题目1：使用SQL扩展库进行查询

// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。


// type Employee struct {
// 	ID         uint	`db:"id"`
// 	Name       string `db:"name"`
// 	Department string `db:"department"`
// 	Salary     float64 `db:"salary"`
// }

type Employee struct {
	ID         uint	
	Name       string 
	Department string 
	Salary     float64 
}

func question1() {

/*	
create table employees(
id integer primary key,
name varchar(64),
department varchar(64),
salary double
)
	insert into employees 
values
(1, "张三", "技术部", 500),
(2, "李四", "客服部", 200),
(3, "王五", "技术部", 800),
(4, "赵六", "客服部", 200);
*/

	emps := []Employee{}
	err = dbx.Select(&emps, "select * from employees where department = ?", "技术部")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(emps)

	emp := Employee{}
	err = dbx.QueryRow("select * from employees order by salary desc limit 1").Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(emp)

	emp2 := Employee{}
	err = dbx.Get(&emp2, "select * from employees order by salary desc limit 1") 
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(emp2)
}



// 题目2：实现类型安全映射

// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type Book struct {
	ID uint
	Title string
	Author string
	Price float64
}


func question2() {
	/*
	create table books(
		id integer primary key,
		title varchar(64),
		author varchar(64),
		price double
	)

	insert into books 
	values
	(1, "农妇", "张三", 500),
	(2, "山泉", "李四",  200),
	(3, "有点田", "王五", 800),
	(4, "和闲", "赵六",  200);
	*/
	books := []Book{}
	err = dbx.Select(&books, "select * from books where price > ?", 50)
	if err != nil {
		panic(err)
	}
	fmt.Println(books)
}

func main() {
	question1()
	question2()
}

