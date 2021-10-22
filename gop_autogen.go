package wep

import (
	fmt "fmt"
	reflect "reflect"
)

func Join(slices interface {
}, sep string) (full string) {
//line /Users/jessonchan/code/golang/wep/wep.gop:7
	if reflect.TypeOf(slices).Kind() == reflect.Slice {
//line /Users/jessonchan/code/golang/wep/wep.gop:8
		s := reflect.ValueOf(slices)
//line /Users/jessonchan/code/golang/wep/wep.gop:9
		for
//line /Users/jessonchan/code/golang/wep/wep.gop:9
		i := 0; i < s.Len();
//line /Users/jessonchan/code/golang/wep/wep.gop:9
		i++ {
//line /Users/jessonchan/code/golang/wep/wep.gop:10
			full += fmt.Sprint(s.Index(i))
//line /Users/jessonchan/code/golang/wep/wep.gop:11
			if i < s.Len()-1 {
//line /Users/jessonchan/code/golang/wep/wep.gop:12
				full += sep
			}
		}
	}
//line /Users/jessonchan/code/golang/wep/wep.gop:16
	return
}
func Each(slices interface {
}, fn func(k int, v interface {
})) {
//line /Users/jessonchan/code/golang/wep/wep.gop:19
	if reflect.TypeOf(slices).Kind() == reflect.Slice {
//line /Users/jessonchan/code/golang/wep/wep.gop:20
		s := reflect.ValueOf(slices)
//line /Users/jessonchan/code/golang/wep/wep.gop:21
		for
//line /Users/jessonchan/code/golang/wep/wep.gop:21
		i := 0; i < s.Len();
//line /Users/jessonchan/code/golang/wep/wep.gop:21
		i++ {
//line /Users/jessonchan/code/golang/wep/wep.gop:22
			fn(i, s.Index(i).Interface())
		}
	}
}
