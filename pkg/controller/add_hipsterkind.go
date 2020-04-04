package controller

import (
	"github.com/alyarctiq/hipster-op/pkg/controller/hipsterkind"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, hipsterkind.Add)
}
