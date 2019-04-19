package main

type group struct {
	id        string
	name      string
	readonly  []string
	readwrite []string
}

func (d database) addgroup(g group) bool {
	ok, _ := d.getgroupbyname(g.name)
	if ok {
		log.Errorln("cannot create group, group \"", g.name, "\", already exist")
		return false
	}
	ro, err := sliceToString(g.readonly)
	if err != nil {
		log.Errorln(err)
		return false
	}
	rw, err := sliceToString(g.readwrite)
	if err != nil {
		log.Errorln(err)
		return false
	}
	_, err = d.queries.Exec(d.db, "add-group", g.name, ro, rw)
	if err != nil {
		log.Errorln("cannot create group :\n", err)
		return false
	}
	return true
}

func (d database) delgroup(id int) bool {
	ok, _ := d.getgroupbyid(id)
	if !ok {
		log.Errorln("Cannot delete group, group do not exist")
		return false
	}
	_, err := d.queries.Exec(d.db, "del-group", id)
	if err != nil {
		log.Errorln("Error while deleting group : \n", err.Error())
		return false
	}
	return true
}

func (d database) getgroupbyid(id int) (bool, group) {
	var g group
	var strinfo, sliceinfo = make([]string, 2), make([][]string, 2)
	row, err := d.queries.QueryRow(d.db, "get-group-by-id", id)
	if err != nil {
		log.Errorln(errQuery, "\n", err)
		return false, group{}
	}
	err = row.Scan(&g.id, &g.name, &strinfo[0], &strinfo[1])
	if err != nil {
		log.Errorln("cannot scan rows :\n", err)
		return false, group{}
	}
	for k, v := range strinfo {
		sliceinfo[k], err = stringToSlice(v)
		if err != nil {
			log.Errorln(err)
			return false, group{}
		}
	}
	g.readonly, g.readwrite = sliceinfo[0], sliceinfo[1]
	return true, g
}

func (d database) getgroupbyname(name string) (bool, group) {
	var g group
	var strinfo, sliceinfo = make([]string, 2), make([][]string, 2)
	row, err := d.queries.QueryRow(d.db, "get-group-by-name", name)
	if err != nil {
		log.Errorln(errQuery, "\n", err)
		return false, group{}
	}
	err = row.Scan(&g.id, &g.name, &strinfo[0], &strinfo[1])
	if err != nil {
		log.Errorln("cannot scan rows :\n", err)
		return false, group{}
	}
	for k, v := range strinfo {
		sliceinfo[k], err = stringToSlice(v)
		if err != nil {
			log.Errorln(err)
			return false, group{}
		}
	}
	g.readonly, g.readwrite = sliceinfo[0], sliceinfo[1]
	return true, g
}

func (d database) grouplist() (bool, []group) {
	var g group
	var gslice, tmpstr, tmpslice = []group{}, make([]string, 2), make([][]string, 2)
	rows, err := d.queries.Query(d.db, "list-groups")
	if err != nil {
		log.Error(errQuery)
		return false, nil
	}
	for rows.Next() {
		err = rows.Scan(&g.id, &g.name, &tmpstr[0], &tmpstr[1])
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
		g.readonly, g.readwrite = tmpslice[0], tmpslice[1]
		gslice = append(gslice, g)
	}
	return true, gslice
}
