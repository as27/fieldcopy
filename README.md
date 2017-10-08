# fieldcopy

fieldcopy copies the values of two structs. Via the fieldcopy tag not all values are copied. 

This is usefull, when you have an external API (for example in JSON) and you don't want all exported fields to be visible.

A simple example looks like this:

```go
    type A struct {
		One   string `fieldcopy:"json" json:"one,omitempty"`
		Two   string `fieldcopy:"json" json:"two,omitempty"`
		MyInt int    `json:"my_int,omitempty"`
	}
	a := A{"one", "two", 42}
	b := A{}
	fieldcopy.Copy(&b, a, "json")
	fmt.Printf("%v", b)
    // Output: {one two 0}

```    
