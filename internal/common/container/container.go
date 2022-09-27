package container

import (
	"github.com/eddieowens/axon"
	"log"
)

func RegisterHandler[T any](instance T) T {
	err := axon.Inject(instance)
	if err != nil {
		log.Fatal(err)
	}
	return instance
}

func Register[T any](reference T) {
	axon.Add(axon.NewTypeKey[T](reference))
}
