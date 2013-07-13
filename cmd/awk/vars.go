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
