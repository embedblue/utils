/// safemap.go
/// Created by PanDaZhong on 2015/05/11.
/// Copyright (c) 2015年 PanDaZhong. All rights reserved.
///

package utils
import (
	"sync"
)

type SafeMap struct {
	lock *sync.RWMutex
	bm   map[interface{}]interface{}
	len int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
		len: 0,
	}
}

// Get the num of the keys
func(m *SafeMap) Len()int{
	return m.len
}

// Get from maps return the k's value
func (m *SafeMap) Get(k interface{}) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.bm[k]; ok {
		return val
	}
	return nil
}

// Maps the given key and value. Returns false
// if the key is already in the map and changes nothing.
func (m *SafeMap) Set(k interface{}, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.bm[k]; !ok {
		m.bm[k] = v
		m.len++
	} else if val != v {
		m.bm[k] = v
	} else {
		return false
	}
	return true
}

// Returns true if k is exist in the map.
func (m *SafeMap) HasKey(k interface{}) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.bm[k]; !ok {
		return false
	}
	return true
}

// Delete the given key and value.
func (m *SafeMap) Delete(k interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.bm, k)
	if 0 < m.len {
		m.len--
	}
}

// Items returns all items in safemap.
func (m *SafeMap) Items() map[interface{}]interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	r := make(map[interface{}]interface{})
	for k, v := range m.bm {
		r[k] = v
	}
	return r
}
