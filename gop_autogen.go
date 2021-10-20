package wep

import (
	fmt "fmt"
	reflect "reflect"
)

func Join(slices interface {
}, sep string) (full string) {
//line /Users/jessonchan/code/golang/wep/wep.gop:6
	if reflect.TypeOf(slices).Kind() == reflect.Slice {
//line /Users/jessonchan/code/golang/wep/wep.gop:7
		s := reflect.ValueOf(slices)
//line /Users/jessonchan/code/golang/wep/wep.gop:8
		for
//line /Users/jessonchan/code/golang/wep/wep.gop:8
		i := 0; i < s.Len();
//line /Users/jessonchan/code/golang/wep/wep.gop:8
		i++ {
//line /Users/jessonchan/code/golang/wep/wep.gop:9
			full += fmt.Sprint(s.Index(i))
//line /Users/jessonchan/code/golang/wep/wep.gop:10
			if i < s.Len()-1 {
//line /Users/jessonchan/code/golang/wep/wep.gop:11
				full += sep
			}
		}
	}
//line /Users/jessonchan/code/golang/wep/wep.gop:15
	return
}
func Each(slices interface {
}, sep string) (full string) {
//line /Users/jessonchan/code/golang/wep/wep.gop:18
	if reflect.TypeOf(slices).Kind() == reflect.Slice {
//line /Users/jessonchan/code/golang/wep/wep.gop:19
		s := reflect.ValueOf(slices)
//line /Users/jessonchan/code/golang/wep/wep.gop:20
		for
//line /Users/jessonchan/code/golang/wep/wep.gop:20
		i := 0; i < s.Len();
//line /Users/jessonchan/code/golang/wep/wep.gop:20
		i++ {
//line /Users/jessonchan/code/golang/wep/wep.gop:21
			full += fmt.Sprint(s.Index(i))
//line /Users/jessonchan/code/golang/wep/wep.gop:22
			if i < s.Len()-1 {
//line /Users/jessonchan/code/golang/wep/wep.gop:23
				full += sep
			}
		}
	}
//line /Users/jessonchan/code/golang/wep/wep.gop:27
	return
}
