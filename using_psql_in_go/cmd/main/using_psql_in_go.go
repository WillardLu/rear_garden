// @copyright Copyright 2024 Willard Lu
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package main

import (
	"database/sql"
	"fmt"
	"using_psql_in_go/internal/mypsql"
)

// 读取task_state表内容
func readTaskState(db *sql.DB) {
	sql := "SELECT code, name FROM task_state where code>=6 order by code"
	rows := mypsql.PsqlSelect(db, sql)
	if rows != nil {
		for rows.Next() {
			var code string
			var name string
			err := rows.Scan(&code, &name)
			if err != nil {
				fmt.Println("查询数据时发生错误")
				return
			}
			fmt.Printf("%s, %s\n", code, name)
		}
		rows.Close()
	}
	return
}

// 插入数据到task_state
func insertTaskState(db *sql.DB, code string, name string) {
	sql := "INSERT INTO task_state (code, name) VALUES (" +
		code + ", " + name + ");"
	if mypsql.PsqlExec(db, sql) == 0 {
		fmt.Println("\n插入数据成功！")
	}
	return
}

// 修改task_state中的数据
func updateTaskState(db *sql.DB, code string, name string) {
	sql := "UPDATE task_state SET name = '" + name + "' WHERE code = " +
		code
	if mypsql.PsqlExec(db, sql) == 0 {
		fmt.Println("\n更新数据成功！")
	}
	return
}

// 删除task_state中的数据
func deleteTaskState(db *sql.DB, code string) {
	sql := "DELETE FROM task_state WHERE code = " + code
	if mypsql.PsqlExec(db, sql) == 0 {
		fmt.Println("\n删除数据成功！")
	}
	return
}

func main() {
	// 获取数据库连接参数
	config, err := mypsql.GetPsqlConfig("pg_config.toml")
	if err != 0 {
		fmt.Println("获取PostgreSQL数据库连接参数时发生错误")
		return
	}
	// 连接数据库
	db, msg := mypsql.ConnectPsql(config)
	fmt.Println(msg)
	if db == nil {
		return
	}
	// 查询数据
	fmt.Println("\n-------- 查询数据 --------")
	readTaskState(db)
	// 插入数据
	insertTaskState(db, "10", "'测试'")
	// 再次查询数据
	fmt.Println("\n-------- 再次查询数据 --------")
	readTaskState(db)
	// 修改数据
	updateTaskState(db, "10", "被修改的数据")
	// 再次查询数据
	fmt.Println("\n-------- 再次查询数据 --------")
	readTaskState(db)
	// 删除数据
	deleteTaskState(db, "10")
	// 再次查询数据
	fmt.Println("\n-------- 再次查询数据 --------")
	readTaskState(db)
	// 关闭数据库连接
	msg, err = mypsql.ClosePsql(db)
	fmt.Println(msg)

	return
}
