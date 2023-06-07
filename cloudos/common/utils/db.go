package utils

import (
	"cloudos/common/interface/common_iface"
	"fmt"
)

func NewLikeValue(value string) string {
	return fmt.Sprintf("%%%s%%", value)
}

func NewLeftLikeValue(value string) string {
	return fmt.Sprintf("%s%%", value)
}

func NewRangeValue(rang *common_iface.RangeI64) (left string, right string) {
	switch rang.Scope {
	case common_iface.RangeScope_RANGE_ALL:
		left = ">="
		right = "<="
	case common_iface.RangeScope_RANGE_LEFT:
		left = ">="
		right = "<"
	case common_iface.RangeScope_RANGE_RIGHT:
		left = ">"
		right = "<="
	}
	return left, right
}
