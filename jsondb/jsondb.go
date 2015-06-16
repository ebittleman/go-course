package jsondb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

const (
	tableDir = "./data"
)

var locksLock *sync.Mutex
var locks map[string]*sync.Mutex

func init() {
	locksLock = &sync.Mutex{}
	locks = make(map[string]*sync.Mutex)
}

func LoadTable(name string, table interface{}) error {
	lock := getLock(name)
	lock.Lock()
	defer lock.Unlock()

	data, err := ioutil.ReadFile(getTableFile(name))

	if err != nil {
		return err
	}

	return json.Unmarshal(data, table)
}

func WriteTable(name string, table interface{}) error {
	lock := getLock(name)
	lock.Lock()
	defer lock.Unlock()

	data, err := json.Marshal(table)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(getTableFile(name), data, 0644)
}

func getTableFile(name string) string {
	return fmt.Sprintf("%s/%s.json", tableDir, name)
}

func getLock(name string) *sync.Mutex {
	locksLock.Lock()
	defer locksLock.Unlock()

	lock, ok := locks[name]

	if !ok {
		lock = &sync.Mutex{}
		locks[name] = lock
	}

	return lock
}
