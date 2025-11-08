package ioc

import (
	"go.uber.org/dig"
)

type ContainerDI = dig.Container

func NewContainerDI() *ContainerDI {
	return dig.New()
}
