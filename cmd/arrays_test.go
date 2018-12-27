/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/27
 * @Last Modified by: liliangshan
 * *************************************************/

package cmd

import (
	"fmt"
	"log"
	"testing"
)

type User struct {
	Id int
	Name string
}

func TestUsers(t *testing.T) {
	users := []User{
		{ 0, "User0"},
		{ 1, "User1"},
	}

	fmt.Println(users, len(users))
	log.Println(users, len(users))
}
