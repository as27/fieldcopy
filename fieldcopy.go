// Package fieldcopy allows you to copy fields, which are defined via
// a tag. The standard usecase is a big struct which is part of an API.
// If you don't want that all the values are visible you can use fieldcopy
// just to copy that fields you want.
//
// For example:
// You read values from a DB and there are fields which just admins are
// allowed to see. You can define that property direct on the struct.
//
//   type Node struct {
//     Username string `fieldcopy:"user,admin"`
//     Secret   string `fieldcopy:"admin"`
//   }
//
// Then you copy the values from your internal variable to the one which
// is visble.
//
//   internalNode := readDB()
//   exernal := Node{}
//   fieldcopy.Copy(&extern,internalNode,"user")
package fieldcopy

import (
	"errors"
	"reflect"
	"strings"
)

const tagName = "fieldcopy"

// ErrDstNeedPtr is used, when the input for the dst is not a pointer
var ErrDstNeedPtr = errors.New("need a pointer as dst input")

var ErrNeedSameType = errors.New("src and dst must have the same type")

// Copy copies the values from src to dst. It is important that both
// variables are the same type, otherwise just an error is returned.
// It is also important that dst is a pointer, that the values can
// be copied into that variable.
//
// Tagvalue defines the values, which are copied.
func Copy(dst, src interface{}, tagvalue string) error {
	var vdst reflect.Value
	if reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return ErrDstNeedPtr
	}
	vdst = reflect.ValueOf(dst).Elem()
	vsrc := reflect.ValueOf(src)
	ta := vdst.Type()
	tb := vsrc.Type()
	if ta.Name() != tb.Name() {
		return ErrNeedSameType
	}
	for i := 0; i < vdst.NumField(); i++ {
		t := ta.Field(i)
		if valueInTag(t.Tag.Get(tagName), tagvalue) {
			dstfield := vdst.Field(i)
			dstfield.Set(vsrc.Field(i))
		}
	}
	return nil
}

func valueInTag(tagval, value string) bool {
	return stringInSlice(value, parseTagVal(tagval))
}
func parseTagVal(s string) []string {
	return strings.Split(s, ",")
}

func stringInSlice(search string, sl []string) bool {
	for _, s := range sl {
		if search == s {
			return true
		}
	}
	return false
}
