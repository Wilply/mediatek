package main

import (
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

/*func testhandler(w http.ResponseWriter, r *http.Request) {
	buf, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr == nil {
		fmt.Println(bytes.NewBuffer(buf))
	}
	fmt.Println(r.Form)
	fmt.Println(r.PostForm)
	fmt.Println(r.FormValue("x_test1"))
	fmt.Println(r.PostFormValue("x_test1"))
}*/

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("POST methode only"))
		w.WriteHeader(405)
		return
	}
	name, pass := r.PostFormValue("username"), r.PostFormValue("password")
	if name == "" || pass == "" {
		w.WriteHeader(400)
		return
	}
	exist, _ := db.getuserbyname(name)
	if exist {
		w.WriteHeader(400)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	ok := db.adduser(user{
		name:   name,
		pass:   string(hash),
		active: true,
	})
	if !ok {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
	return
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("POST methode only"))
		w.WriteHeader(405)
		return
	}
	name, pass := r.PostFormValue("username"), r.PostFormValue("password")
	if name == "" || pass == "" {
		w.WriteHeader(400)
		return
	}
	exist, u := db.getuserbyname(name)
	if !exist {
		w.WriteHeader(401)
		return
	}
	if !comparePassword(u.pass, pass) {
		w.WriteHeader(401)
		return
	}
	w.WriteHeader(200)
	return
}

func userlistHandler(w http.ResponseWriter, r *http.Request) {
	_, ulist := db.userlist()
	for _, v := range ulist {
		grp, _ := sliceToString(v.groups)
		ro, _ := sliceToString(v.readonly)
		rw, _ := sliceToString(v.readwrite)
		//TODO format with sprintf
		w.Write([]byte(strings.Join([]string{v.name, grp, ro, rw, "\n"}, " | ")))
	}
	return
}

func comparePassword(hash, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return false
	}
	return true
}
