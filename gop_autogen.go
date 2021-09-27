package wep

import (
	fmt "fmt"
	reflect "reflect"
)

func Run() {
}
func Join(slices interface {
}, sep string) string {
//line /tmp/wep/wep.gop:9
	full := ""
//line /tmp/wep/wep.gop:10
	if reflect.TypeOf(slices).Kind() == reflect.Slice {
//line /tmp/wep/wep.gop:11
		s := reflect.ValueOf(slices)
//line /tmp/wep/wep.gop:12
		for
//line /tmp/wep/wep.gop:12
		i := 0; i < s.Len();
//line /tmp/wep/wep.gop:12
		i++ {
//line /tmp/wep/wep.gop:13
			full += fmt.Sprint(s.Index(i))
//line /tmp/wep/wep.gop:14
			if i < s.Len()-1 {
//line /tmp/wep/wep.gop:15
				full += sep
			}
		}
	}
//line /tmp/wep/wep.gop:19
	return full
}
