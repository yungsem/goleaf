package inits

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// initDB 初始化 MySQL 数据库连接
func initDB() {
	Log.Debug("初始化MySQL数据库连接-开始")
	// 连接数据库
	// 返回的 db 是线程安全的，并且包含一个连接池
	// 在一个项目中，该方法大概率只需要调用一次，因为大多数情况下，一个项目只维护一个 db 连接池
	// db 也不需要手动关闭，没有这种需要。只要项目在运行中，这个 db 就是不能关闭的。但是项目终止运行，db 也自动被消失了
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goleaf")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		Conf.Mysql.Username, Conf.Mysql.Password, Conf.Mysql.Host, Conf.Mysql.Port, Conf.Mysql.DB))
	if err != nil {
		panic(err)
	}

	// 设置连接池的最大连接数
	// 默认值是 2
	db.SetMaxIdleConns(10)

	// 设置一个连接的最大存活时间
	// 如果设置为 <= 0 ，则表示永久存活
	db.SetConnMaxLifetime(0)

	DB = db

	Log.Debug("初始化MySQL数据库连接-成功")
	return
}
