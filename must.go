// Must<Type> convert data to specified type but panics if convert fails.
package redis

import (
	"fmt"
)

func check(err error, name string, data interface{}) {
	if err != nil {
		panic(fmt.Sprintf("redis %s: [%#v]: %s", name, data, err.Error()))
	}
}

// convert data to string
func MustString(data interface{}, err error) string {
	str, err := String(data, err)
	check(err, "MustString", data)
	return str
}

// convert data to []string
func MustStrings(data interface{}, err error) []string {
	ss, err := Strings(data, err)
	check(err, "MustStrings", data)
	return ss
}

// convert data to int
func MustInt(data interface{}, err error) int {
	n, err := Int(data, err)
	check(err, "MustInt", data)
	return n
}

// convert data to int64
func MustInt64(data interface{}, err error) int64 {
	n, err := Int64(data, err)
	check(err, "MustInt64", data)
	return n
}
