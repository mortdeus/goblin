package main

import (
	"sync"
)

type cell struct {
	//ctype byte
	//csub  byte
	name string
	sval string
	nval float64
	tval int
	next *cell
	lock sync.Mutex
}

type varTable struct {
	lock sync.Mutex
	tab  map[string]*cell
}

var vTab = &varTable{
	tab: map[string]*cell{
		"CONVFMT":  &cell{name: "CONVFMT", sval: `%.6g`},
		"FS":       &cell{name: "FS", sval: " "},
		"NF":       &cell{name: "NF"},
		"NR":       &cell{name: "NR"},
		"FNR":      &cell{name: "FNR"},
		"FILENAME": &cell{name: "FILENAME"},
		"RS":       &cell{name: "RS", sval: `\n`},
		"OFS":      &cell{name: "OFS"},
		"ORS":      &cell{name: "ORS", sval: `\n`},
		"OFMT":     &cell{name: "OFMT", sval: `%.6g`},
		"SUBSEP":   &cell{name: "SUBSEP", sval: "034", nval: 034},
		"ARGC":     &cell{name: "ARGC"},
		"ARGV":     &cell{name: "ARGV"},
		"ENVIRON":  &cell{name: "ENVIRON"},
	}}

func (vt *varTable) set(id, val string) error {
	vt.lock.Lock()
	defer vt.lock.Unlock()

	if v, ok := vt.tab[id]; ok {
		v.lock.Lock()
		v.sval = val
		v.lock.Unlock()
	} else {
		return unknownVarErr(0)
	}
	return nil
}

func (vt *varTable) add(id, val string) error {
	vt.lock.Lock()
	defer vt.lock.Unlock()
	if _, ok := vt.tab[id]; ok {
		return varExistsErr(0)
	} else {
		vt.tab[id] = &cell{name: id, sval: val}
	}
	return nil
}

func (vt *varTable) remove(id string) error {
	vt.lock.Lock()
	defer vt.lock.Unlock()
	if v, ok := vt.tab[id]; ok {
		v.lock.Lock()
		defer release(&v)
	} else {
		return unknownVarErr(0)
	}
	delete(vt.tab, id)
	return nil
}

func (vt *varTable) fetch(id string) (*cell, error) {
	vt.lock.Lock()
	defer vt.lock.Unlock()
	if c, ok := vt.tab[id]; ok {
		c.lock.Lock()
		return c, nil
	} else {
		return nil, unknownVarErr(0)
	}
}

func release(cp **cell) {
	(*cp).lock.Unlock()
	*cp = nil
}
