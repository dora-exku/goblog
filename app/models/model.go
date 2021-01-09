package models

import "goblog/pkg/types"

type BaseModel struct {
	ID uint64
}

func (a BaseModel) GetStringId() string {
	return types.Unit64ToString(a.ID)
}
