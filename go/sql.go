package go

import (
	"context"
	"database/sql"
)

// go语言中使用原生sql语句
func SQL() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
        panic(err.Error())
    }
    defer db.Close()

	// 增删改
	_, err = db.Exec("  INSERT user(name, age) VALUES(?, ?)", "Alice", 18)
	_, err = db.ExecContext(context.Background(), "INSERT user(name, age) VALUES(?, ?)", "Alice", 18)

	// 查 Query
	_, err = db.QueryContext(context.Background(), "SELECT * FROM user WHERE id = ?", 1)

}




