package fieldcopy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCopy(t *testing.T) {
	type A struct {
		One   string `fieldcopy:"string,safe"`
		Two   string `fieldcopy:"string,other"`
		MyInt int    `fieldcopy:"int,safe"`
	}
	a := A{
		"A One",
		"A Two",
		1,
	}
	b := A{
		"B One",
		"B Two",
		5,
	}

	type args struct {
		dst      A
		src      interface{}
		tagvalue string
	}
	tests := []struct {
		name   string
		args   args
		expDst A
	}{
		{
			"Copy with tag 'string'",
			args{a, b, "string"},
			A{
				"B One",
				"B Two",
				1,
			},
		},
		{
			"Copy with tag 'safe' (string and int)",
			args{a, b, "safe"},
			A{
				"B One",
				"A Two",
				5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Copy(&tt.args.dst, tt.args.src, tt.args.tagvalue)
			if !reflect.DeepEqual(tt.args.dst, tt.expDst) {
				t.Errorf("Copy() = %#v, want %#v", tt.args.dst, tt.expDst)
			}
		})
	}
}

func Example_Copy() {
	type A struct {
		One   string `fieldcopy:"json" json:"one,omitempty"`
		Two   string `fieldcopy:"json" json:"two,omitempty"`
		MyInt int    `json:"my_int,omitempty"`
	}
	a := A{"one", "two", 42}
	b := A{}
	Copy(&b, a, "json")
	fmt.Printf("%v", b)
	// Output: {one two 0}
}
