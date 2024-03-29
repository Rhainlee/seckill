package config

import (
	"regexp"
	"strings"
)

var (
	AuthPermitConfig AuthPermitAll
)

// Http配置
type AuthPermitAll struct {
	PermitAll []interface{}
}

// 正则表达式
func Match(str string) bool {
	if len(AuthPermitConfig.PermitAll) > 0 {
		targetValue := AuthPermitConfig.PermitAll
		for i := 0; i < len(targetValue); i++ {
			s := targetValue[i].(string)
			res, _ := regexp.MatchString(strings.ReplaceAll(s, "**", "(.*?)"), str)
			if res {
				return true
			}

		}
	}
	return false
}
