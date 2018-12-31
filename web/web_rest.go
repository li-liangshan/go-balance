/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/31
 * @Last Modified by: liliangshan
 * *************************************************/

package web

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are geting user %s", uid)
}

func modifyUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are modifing user %s", uid)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are deleting user %s", uid)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	fmt.Fprintf(w, "you are adding user %s", uid)
}

func TestRest() {
	mux := routes.New()
	mux.Get("/user/:uid", getUser)
	mux.Post("/user", addUser)
	mux.Del("/user/:uid", deleteUser)
	mux.Put("/user/:uid", modifyUser)
	http.Handle("/", mux)
	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
