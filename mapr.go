// Steve Phillips / elimisteve
// 2012.08.17

package mapr

import (
	"fmt"
	"strings"
)

// Printf takes a special format string -- one resembling "My name is
// {{myName}} and I'm {{age}} years old" -- and replaces `{{myName}}`
// with the value of m["myName"], `{{age}}` with the value of
// m["age"], etc, then prints the resulting string. m should be a map
// from strings to values of type string, int, int64, or interface{}.
func Printf(format string, m interface{}) {
	defer func() { fmt.Print(format) }()

	// You'd think a type switch could save the day... I don't think
	// so :-(
	if t, ok := m.(map[string]interface{}); ok {
		for k, _ := range t {
			format = replace(format, k, t[k])
		}
		return
	}
	if t, ok := m.(map[string]string); ok {
		for k, _ := range t {
			format = replace(format, k, t[k])
		}
		return
	}
	if t, ok := m.(map[string]int); ok {
		for k, _ := range t {
			format = replace(format, k, t[k])
		}
		return
	}
	if t, ok := m.(map[string]int64); ok {
		for k, _ := range t {
			format = replace(format, k, t[k])
		}
		return
	}
	// TODO: Use reflection to support structs in addition to maps

	return
}

func replace(format, key string, val interface{}) string {
	replacement := fmt.Sprintf("%v", val)
	format = strings.Replace(format, "{{"+key+"}}", replacement, -1)
	return format
}
