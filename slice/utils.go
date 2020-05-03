package slice

import (
	"fmt"
)

func FindStr(slice []byte, key interface{}) (flag bool) {
	fmt.Println(slice)
	fmt.Println(string(slice))
	return false
}
