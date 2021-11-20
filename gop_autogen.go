package wep

import (
	fmt "fmt"
	reflect "reflect"
	"strings"
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
func EachMap(slices interface {
}, fn func(k interface {
}, v interface {
})) {
	//line /Users/jessonchan/code/golang/wep/wep.gop:27
	if reflect.TypeOf(slices).Kind() == reflect.Map {
		//line /Users/jessonchan/code/golang/wep/wep.gop:28
		s := reflect.ValueOf(slices)
		//line /Users/jessonchan/code/golang/wep/wep.gop:29
		keys := s.MapKeys()
		//line /Users/jessonchan/code/golang/wep/wep.gop:30
		for
		//line /Users/jessonchan/code/golang/wep/wep.gop:30
		i := 0; i < len(keys);
		//line /Users/jessonchan/code/golang/wep/wep.gop:30
		i++ {
			//line /Users/jessonchan/code/golang/wep/wep.gop:31
			fn(keys[i].Interface(), s.MapIndex(keys[i]).Interface())
		}
	}
}
func SubString(str string, start int, end int) string {
	//line /Users/jessonchan/code/golang/wep/wep.gop:36
	if end < 0 {
		//line /Users/jessonchan/code/golang/wep/wep.gop:37
		end = len(str)
	}
	//line /Users/jessonchan/code/golang/wep/wep.gop:39
	if start >= end || str == "" || start >= len(str) {
		//line /Users/jessonchan/code/golang/wep/wep.gop:40
		return ""
	}
	//line /Users/jessonchan/code/golang/wep/wep.gop:42
	rs := []rune(str)
	//line /Users/jessonchan/code/golang/wep/wep.gop:43
	if end > len(rs) {
		//line /Users/jessonchan/code/golang/wep/wep.gop:44
		end = len(rs)
	}
	//line /Users/jessonchan/code/golang/wep/wep.gop:46
	return string(rs[start:end])
}
func Contains(str string, search string) bool {
	return strings.Contains(str, search)
}
