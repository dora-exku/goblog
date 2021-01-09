package types

import (
	"goblog/pkg/logger"
	"strconv"
)

func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}

func Unit64ToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}