package handler

import(
	"log"
	"html/template"
	"net/http"
	"my_web/user"
	"strconv"
)

func RootHdl(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	right := user.UserSessionMgr.GetValue(r, "right")
	if right == nil {
		t := template.New("login.html")
		t, _ = t.ParseFiles("view/pages/examples/login.html")
		err := t.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
		return
	}

	t := template.New("index.html")
	t, _ = t.ParseFiles("view/index.html")
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func Index2Hdl(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	t := template.New("index2.html")
	t, _ = t.ParseFiles("view/index2.html")

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func IndexHdl(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	t := template.New("index1.html")
	t, _ = t.ParseFiles("view/index1.html")

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func LoginHdl(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	remember := r.FormValue("remember")
	name := r.FormValue("name")
	password := r.FormValue("password")

	sessionID, userInfo, err := user.UserLogin(name, password)
	if err != nil {
		w.Write([]byte(`{"errorCode":1, "str":"` + err.Error() + `"}`))
		return
	}
	if err = user.UserSessionMgr.SetValueByID(sessionID, "id", userInfo.ID); err != nil {
		w.Write([]byte(`{"errorCode":1, "str":"` + err.Error() + `"}`))
		return
	}
	if err = user.UserSessionMgr.SetValueByID(sessionID, "pfid", userInfo.PfID); err != nil {
		w.Write([]byte(`{"errorCode":1, "str":"` + err.Error() + `"}`))
		return
	}
	if err = user.UserSessionMgr.SetValueByID(sessionID, "right", userInfo.Right); err != nil {
		w.Write([]byte(`{"errorCode":1, "str":"` + err.Error() + `"}`))
		return
	}
	if err = user.UserSessionMgr.SetValueByID(sessionID, "src", userInfo.Address); err != nil {
		w.Write([]byte(`{"errorCode":1, "str":"` + err.Error() + `"}`))
		return
	}

	if remember == "true" {
		//cookie := http.Cookie{Name: "name", Value: name}
		//http.SetCookie(w, &cookie)
	}

	cookie := http.Cookie{Name: "usersession", Value: sessionID, Path: "/"}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "id", Value: strconv.Itoa(userInfo.ID), Path: "/"}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "pfid", Value: strconv.Itoa(userInfo.PfID), Path: "/"}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "name", Value: userInfo.Name, Path: "/"}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "right", Value: strconv.Itoa(userInfo.Right), Path: "/"}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "right_name", Value: userInfo.RightName, Path: "/"}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "src", Value: userInfo.Address, Path: "/"}
	http.SetCookie(w, &cookie)
	//rightValue, _ := user.GetRightValue(userInfo.Right)
	//cookie = http.Cookie{Name: "rightvalue_first", Value: strconv.FormatUint(rightValue&0xFFFFFFFF, 10), Path: "/"}
	//http.SetCookie(w, &cookie)
	//cookie = http.Cookie{Name: "rightvalue_second", Value: strconv.FormatUint(rightValue>>32, 10), Path: "/"}
	//http.SetCookie(w, &cookie)

	cookie = http.Cookie{Name: "name", Value: name}
	http.SetCookie(w, &cookie)
	w.Write([]byte(`{"errorCode":0}`))
}
func LogoutHdl(w http.ResponseWriter, r *http.Request) {
	id := user.UserSessionMgr.GetValue(r, "id")
	if id != nil {
		user.UserSessionMgr.DelSession(id.(int))
	}
	w.Write([]byte(`{"errorCode":0}`))
}
func UserManageHdl(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	t := template.New("data.html")
	t, _ = t.ParseFiles("view/pages/tables/data.html")

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
func UserListHdl(w http.ResponseWriter, r *http.Request) {
	right := user.UserSessionMgr.GetValue(r, "right")
	if right == nil {
		w.Write([]byte(`{"errorCode":2,"str":"please login"}`))
		return
	}
	var j *[]byte
	var err error

	j, err = user.GetAllUser()

	if err != nil {
		log.Println(err)
		w.Write([]byte(`{"errorCode":1,"str":"` + err.Error() + `"}`))
		return
	}
	w.Write([]byte(`{"errorCode":0,"users":` + string(*j) + `}`))
}
func UserAddHdl(w http.ResponseWriter, r *http.Request) {
	right := user.UserSessionMgr.GetValue(r, "right")
	if right == nil {
		w.Write([]byte(`{"errorCode":2,"str":"please login"}`))
		return
	}
	if err := r.ParseForm(); err != nil {
		w.Write([]byte(`{"errorCode":1,"str":"` + err.Error() + `"}`))
		return
	}
	pfID, _ := strconv.Atoi(r.FormValue("pfid"))
	name := r.FormValue("name")
	password := r.FormValue("password")
	userRight, _ := strconv.Atoi(r.FormValue("right"))

	id, err := user.AddUser(pfID, name, password, userRight)
	if err != nil {
		w.Write([]byte(`{"errorCode":1,"str":"` + err.Error() + `"}`))
		return
	}

	w.Write([]byte(`{"errorCode":0, "id":` + strconv.Itoa(id) + `}`))
}

func UserUpdateHdl(w http.ResponseWriter, r *http.Request) {
	right := user.UserSessionMgr.GetValue(r, "right")
	if right == nil {
		w.Write([]byte(`{"errorCode":2,"str":"please login"}`))
		return
	}
	if err := r.ParseForm(); err != nil {
		w.Write([]byte(`{"errorCode":1,"str":"` + err.Error() + `"}`))
		return
	}
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	userRight, _ := strconv.Atoi(r.FormValue("right"))
	pfID, _ := strconv.Atoi(r.FormValue("pfid"))

	var password *string = nil
	if r.FormValue("changepassword") == "true" {
		pwd := r.FormValue("password")
		password = &pwd
	}

	err := user.UpdateUser(id, name, password, userRight, pfID)
	if err != nil {
		w.Write([]byte(`{"errorCode":1,"str":"` + err.Error() + `"}`))
		return
	}

	w.Write([]byte(`{"errorCode":0}`))
}

func UserDelHdl(w http.ResponseWriter, r *http.Request) {
	right := user.UserSessionMgr.GetValue(r, "right")
	if right == nil {
		w.Write([]byte(`{"errorCode":2,"str":"please login"}`))
		return
	}
	if err := r.ParseForm(); err != nil {
		w.Write([]byte(`{"errorCode":1,"str":"` + err.Error() + `"}`))
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	pfID, _ := strconv.Atoi(r.FormValue("pfid"))

	err := user.DelUser(id, pfID)

	if err != nil {
		w.Write([]byte(`{"errorCode":1,"str":"` + err.Error() + `"}`))
		return
	}
	w.Write([]byte(`{"errorCode":0}`))
}

func RightListHdl(w http.ResponseWriter, r *http.Request) {
	right := user.UserSessionMgr.GetValue(r, "right")
	if right == nil {
		w.Write([]byte(`{"errorCode":2,"str":"please login"}`))
		return
	}
	/*
	rightVal, errVal := user.GetRightValue(right.(int))
	if errVal != nil {
		w.Write([]byte(`{"errorCode":1,"str":"` + errVal.Error() + `"}`))
		return
	}
	if (rightVal & user.UserRightViewRight) == 0 {
		w.Write([]byte(`{"errorCode":1,"str":"permission denied"}`))
		return
	}
	*/
	j, err := user.GetAllRight()

	if err != nil {
		log.Println(err)
		w.Write([]byte(`{"errorCode":1,"str":"` + err.Error() + `"}`))
		return
	}
	w.Write([]byte(`{"errorCode":0,"rights":` + string(*j) + `}`))
}

func PlatformListHdl(w http.ResponseWriter, r *http.Request) {
	right := user.UserSessionMgr.GetValue(r, "right")
	if right == nil {
		w.Write([]byte(`{"errorCode":2,"str":"please login"}`))
		return
	}
	/*
	rightVal, errVal := user.GetRightValue(right.(int))
	if errVal != nil {
		w.Write([]byte(`{"errorCode":1,"str":"` + errVal.Error() + `"}`))
		return
	}
	if (rightVal & user.UserRightViewRight) == 0 {
		w.Write([]byte(`{"errorCode":1,"str":"permission denied"}`))
		return
	}
	*/
	j, err := user.GetAllPlatform()
	if err != nil {
		log.Println(err)
		w.Write([]byte(`{"errorCode":1,"str":"` + err.Error() + `"}`))
		return
	}
	w.Write([]byte(`{"errorCode":0,"platforms":` + string(*j) + `}`))
}