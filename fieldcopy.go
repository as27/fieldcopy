package fieldcopy

import (
	"errors"
	"reflect"
	"strings"
)

const tagName = "fieldcopy"

func Copy(dst, src interface{}, tagvalue string) error {
	vdst := reflect.ValueOf(dst).Elem()
	vsrc := reflect.ValueOf(src)
	ta := vdst.Type()
	tb := vsrc.Type()
	if ta.Name() != tb.Name() {
		return errors.New("src and dst must have the same type")
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
