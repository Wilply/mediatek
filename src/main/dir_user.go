package main

type user struct {
	id        int
	name      string
	pass      string
	active    bool
	groups    []string //group id
	readonly  []string
	readwrite []string
}

func (d database) adduser(u user) bool {
	ok, _ := d.getuserbyname(u.name)
	if ok {
		log.Errorln("cannot create user, user \"", u.name, "\", already exist")
		return false
	}
	grp, err := sliceToString(u.groups)
	if err != nil {
		log.Errorln(err)
		return false
	}
	ro, err := sliceToString(u.readonly)
	if err != nil {
		log.Errorln(err)
		return false
	}
	rw, err := sliceToString(u.readwrite)
	if err != nil {
		log.Errorln(err)
		return false
	}
	_, err = d.queries.Exec(d.db, "add-user", u.name, u.pass, u.active, grp, ro, rw)
	if err != nil {
		log.Errorln("cannot create user :\n", err)
		return false
	}
	log.Info("succesfully added new user :\"", u.name, "\"")
	return true
}

func (d database) deluser(id int) bool {
	ok, _ := d.getuserbyid(id)
	if !ok {
		log.Errorln("Cannot delete user, user do not exist")
		return false
	}
	_, err := d.queries.Exec(d.db, "del-user", id)
	if err != nil {
		log.Errorln("Error while deleting user : \n", err.Error())
		return false
	}
	return true
}

//may be useless
func (d database) getuserbyid(id int) (bool, user) {
	var u user
	var strinfo, sliceinfo = make([]string, 3), make([][]string, 3)
	row, err := d.queries.QueryRow(d.db, "get-user-by-id", id)
	if err != nil {
		log.Errorln(errQuery, "\n", err)
		return false, user{}
	}
	err = row.Scan(&u.id, &u.name, &u.pass, &u.active, &strinfo[0], &strinfo[1], &strinfo[2])
	if err != nil {
		log.Errorln("cannot scan rows :\n", err)
		return false, user{}
	}
	for k, v := range strinfo {
		sliceinfo[k], err = stringToSlice(v)
		if err != nil {
			log.Errorln(err)
			return false, user{}
		}
	}
	u.groups, u.readonly, u.readwrite = sliceinfo[0], sliceinfo[1], sliceinfo[2]
	return true, u
}

func (d database) getuserbyname(name string) (bool, user) {
	var u user
	var strinfo, sliceinfo = make([]string, 3), make([][]string, 3)
	row, err := d.queries.QueryRow(d.db, "get-user-by-name", name)
	if err != nil {
		log.Errorln(errQuery, "\n", err)
		return false, user{}
	}
	err = row.Scan(&u.id, &u.name, &u.pass, &u.active, &strinfo[0], &strinfo[1], &strinfo[2])
	if err != nil {
		//log.Errorln("cannot scan rows :\n", err)
		return false, user{}
	}
	for k, v := range strinfo {
		sliceinfo[k], err = stringToSlice(v)
		if err != nil {
			log.Errorln(err)
			return false, user{}
		}
	}
	u.groups, u.readonly, u.readwrite = sliceinfo[0], sliceinfo[1], sliceinfo[2]
	return true, u
}

//dont return password
func (d database) userlist() (bool, []user) {
	var u = user{}
	tmpstr := make([]string, 3)
	tmpslice := make([][]string, 3)
	var userlist []user
	rows, err := d.queries.Query(d.db, "list-users-safe")
	if err != nil {
		log.Errorln(errQuery, "\n", err)
		return false, nil
	}
	for rows.Next() {
		err = rows.Scan(&u.id, &u.name, &u.active, &tmpstr[0], &tmpstr[1], &tmpstr[2])
		if err != nil {
			log.Errorln("cannot scan row :\n", err)
			return false, nil
		}
		for k, v := range tmpstr {
			tmpslice[k], err = stringToSlice(v)
			if err != nil {
				log.Errorln(err)
				return false, nil
			}
		}
		u.groups, u.readonly, u.readwrite = tmpslice[0], tmpslice[1], tmpslice[2]
		userlist = append(userlist, u)
	}
	return true, userlist
}
