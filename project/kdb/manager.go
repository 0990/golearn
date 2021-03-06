package kdb

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const defaultGroupName = "mysql"

var m = newManager()

type manager struct {
	dbs map[string]map[string][]*sql.DB
}

func newManager() *manager {
	m := new(manager)
	m.dbs = make(map[string]map[string][]*sql.DB)
	return m
}

func (p *manager) addDB(groupName string, isMaster bool, db *sql.DB) {
	dc := "master"

	if !isMaster {
		dc = "slave"
	}

	group, ok := m.dbs[groupName]
	if !ok {
		group = make(map[string][]*sql.DB)
	}
	if _, ok := group[dc]; ok {
		group[dc] = append(group[dc], db)
	} else {
		group[dc] = []*sql.DB{db}
	}

	m.dbs[groupName] = group
}

func (p *manager) getDB(names ...string) (*sql.DB, error) {
	groupName := defaultGroupName
	dc := "master"

	if len(names) > 0 {
		name := names[0]
		segment := strings.Split(name, "::")
		groupName = segment[0]
		if len(segment) > 1 {
			dc = segment[1]
		}
	}

	if dbs, ok := m.dbs[groupName][dc]; ok {
		max := len(dbs)
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(max)
		return dbs[i], nil
	}
	return nil, fmt.Errorf("DataBase")
}

func (p *manager) getReadDB(names ...string) (*sql.DB, error) {
	groupName := defaultGroupName
	if len(names) > 0 {
		groupName = names[0]
	}
	name := fmt.Sprintf("%s::%s", groupName, "slave")
	return m.getDB(name)
}
