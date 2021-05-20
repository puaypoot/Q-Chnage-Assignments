package services

import (
	"context"
	"q-change-assignment/src/xyz/dao"
)

type XYZService struct {
}

type IXYZService interface {
	FindXYZ(context.Context) (dao.FindXYZResponse, error)
}

func NewXYZService() IXYZService {
	return &XYZService{}
}
