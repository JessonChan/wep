package wep

import fmt "fmt"

func Run(args ...string) {
//line /tmp/wep/wep.gop:4
	addr := append(args, ":8080")[0]
//line /tmp/wep/wep.gop:5
	fmt.Println("go" + addr)
}
