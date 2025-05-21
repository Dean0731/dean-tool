package utils

import (
	_const "github.com/dean0731/dean-tool/const"
)

const (
	defaultNumPerPage = 20
	defaultPageNum    = 1
)

func GetResourceFromPageNum[T any](pageNum, numPerPage int, data []T) []T {
	if pageNum <= _const.NUM_ZERO {
		pageNum = defaultPageNum
	}
	if numPerPage <= _const.NUM_ZERO {
		numPerPage = defaultNumPerPage
	}

	count := len(data)
	start := (pageNum - _const.NUM_ONE) * numPerPage
	end := start + numPerPage

	if start >= count {
		data = []T{}
	} else {
		if end > count {
			end = count
		}
		data = data[start:end]
	}
	return data
}

// Paginate 分页函数：将任意类型切片分页并转为 map[int][]T
func Paginate[T any](data []T, pageSize int) map[int][]T {
	result := make(map[int][]T)
	if pageSize <= 0 || len(data) == 0 {
		return result
	}

	pageNum := 1
	for i := 0; i < len(data); i += pageSize {
		end := i + pageSize
		if end > len(data) {
			end = len(data)
		}
		result[pageNum] = data[i:end]
		pageNum++
	}

	return result
}
