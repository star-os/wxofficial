package util

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

//进行sha1加密
func EncodingSha1(params ...string) string {
	sort.Strings(params)
	str := strings.Join(params, "")
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))

}