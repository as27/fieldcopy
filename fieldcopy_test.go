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
	t.Run("Copy nested struct", func(t *testing.T) {
		type N struct {
			Name     string `fieldcopy:"export"`
			Children []A    `fieldcopy:"export"`
		}
		n1 := N{
			"n1",
			[]A{
				a,
				b,
			},
		}
		n2 := N{}
		Copy(&n2, n1, "export")
		fmt.Println("--------", n2)
	})
}

func TestCopyNested(t *testing.T) {
	type A struct {
		One   string `fieldcopy:"string,safe"`
		Two   string `fieldcopy:"string"`
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

	t.Run("Copy nested struct", func(t *testing.T) {
		type N struct {
			Name     string `fieldcopy:"admin"`
			Children []A    `fieldcopy:"export"`
		}
		n1 := N{
			"n1",
			[]A{
				a,
				b,
			},
		}
		n2 := N{}
		Copy(&n2, n1, "export")
		exp := N{
			"",
			[]A{
				a,
				b,
			},
		}
		if !reflect.DeepEqual(exp, n2) {
			t.Errorf("\nExpect: %v\nGot: %v", exp, n2)
		}

	})
}

func TestCopyError(t *testing.T) {
	t.Run("Error when different types", func(t *testing.T) {
		type A struct {
			One   string `fieldcopy:"string,safe"`
			Two   string `fieldcopy:"string,other"`
			MyInt int    `fieldcopy:"int,safe"`
		}
		type B struct {
			One   string `fieldcopy:"string,safe"`
			Two   string `fieldcopy:"string,other"`
			MyInt int    `fieldcopy:"int,safe"`
		}
		a := A{
			"A One",
			"A Two",
			1,
		}
		b := B{
			"B One",
			"B Two",
			5,
		}
		err := Copy(&a, b, "string")
		if err == nil {
			t.Error("Expect an error, because of didfferent types")
		}
	})
	t.Run("Error when dst is not a pointer", func(t *testing.T) {
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
			4,
		}
		err := Copy(a, b, "string")
		if err == nil {
			t.Error("Expect an error, when input for dst is not a pointer")
		}
		if err != ErrDstNeedPtr {
			t.Errorf("Wrong error\nExp:%v\nGot:%v\n", ErrDstNeedPtr, err)
		}
	})
}

func ExampleCopy() {
	type A struct {
		One   string `fieldcopy:"jsonexport" json:"one,omitempty"`
		Two   string `fieldcopy:"jsonexport" json:"two,omitempty"`
		MyInt int    `json:"my_int,omitempty"`
	}
	a := A{"one", "two", 42}
	b := A{}
	Copy(&b, a, "jsonexport")
	fmt.Printf("%v", b)
	// Output: {one two 0}
}
