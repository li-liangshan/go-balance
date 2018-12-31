/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/31
 * @Last Modified by: liliangshan
 * *************************************************/

package web

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func UseMySQL() {
	db, err := sql.Open("mysql", "root:admin@/test?charset=utf8mb4")
	checkError(err)

	// insert into
	statement, err := db.Prepare("INSERT `user_info` SET `user_name` = ?, `depart_name`=?")
	checkError(err)

	res, err := statement.Exec("li-liangshan", "研发部门")
	checkError(err)

	id, err := res.LastInsertId()
	checkError(err)
	fmt.Printf("id: %d\n", id)

	// update
	statement, err = db.Prepare("UPDATE `user_info` SET `user_name` = ? WHERE `id` = ?")
	checkError(err)

	res, err = statement.Exec("li-liangshan-updated", id)
	checkError(err)

	rowsCount, err := res.RowsAffected()
	checkError(err)
	fmt.Printf("rows:%d\n", rowsCount)

	// select
	rows, err := db.Query("SELECT * FROM `user_info`")
	checkError(err)

	for rows.Next() {
		var id int
		var username string
		var department string
		var createdAt string
		var updatedAt string

		err = rows.Scan(&id, &username, &department, &createdAt, &updatedAt)
		checkError(err)
		fmt.Printf("id:%d\n", id)
		fmt.Printf("username: %s\n", username)
		fmt.Printf("department: %s\n", department)
		fmt.Printf("createdAt: %v\n", createdAt)
		fmt.Printf("updatedAt: %v\n", updatedAt)
	}

	// delete
	statement, err = db.Prepare("DELETE FROM `user_info` WHERE `id` = ?")
	checkError(err)

	res, err = statement.Exec(1)
	checkError(err)

	rowsCount, err = res.RowsAffected()
	checkError(err)
	fmt.Printf("rows:%d\n", rowsCount)

	db.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
