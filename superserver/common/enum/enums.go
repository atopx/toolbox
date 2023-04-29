package enum

import (
	"superserver/domain/public_service"
)

var enumsCache map[Enum]map[int32]string

func InitEnums(enumsList []*public_service.Enums) {
	enumsCache = make(map[Enum]map[int32]string, len(enumsList))
	for _, enums := range enumsList {
		enumCache := make(map[int32]string, len(enums.Data))
		for _, enum := range enums.Data {
			enumCache[enum.Value] = enum.Desc
		}
		enumsCache[Enum(enums.Key)] = enumCache
	}
}

func Desc(key Enum, enum int32) string {
	return enumsCache[key][enum]
}
