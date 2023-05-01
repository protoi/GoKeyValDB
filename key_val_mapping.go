package main

type KeyValMapping struct {
	mapping map[string]string
}

func (kvmap *KeyValMapping) init() {
	kvmap.mapping = make(map[string]string)
}
func (kvmap *KeyValMapping) SetValue(k string, v string) bool {
	if kvmap.mapping == nil {
		return false
	}
	kvmap.mapping[k] = v
	return true
}
func (kvmap *KeyValMapping) GetValue(k string) (string, bool) {
	if kvmap.mapping == nil {
		return "", false
	}
	if v, ok := kvmap.mapping[k]; ok {
		return v, ok
	}
	return "", false
}
func (kvmap *KeyValMapping) DelValue(k string) bool {
	if kvmap.mapping == nil {
		return false
	}

	if _, ok := kvmap.mapping[k]; ok {
		delete(kvmap.mapping, k)
		return ok
	}
	return false
}
func (kvmap *KeyValMapping) FlushMap() bool {
	if kvmap.mapping == nil {
		return false
	}
	for key := range kvmap.mapping {
		delete(kvmap.mapping, key)
	}
	return true
}

/**
INIT
SET key value - sets the value of a key in Redis to a given string value. ✅
GET key - retrieves the value of a key in Redis. ✅
DEL key - deletes a key and its value from Redis. ✅
FLUSHDB - deletes all keys and their values from the current Redis database. ✅

*/
