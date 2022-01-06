package main

import (
	"fmt"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func main() {
	repository, err := name.NewRepository("lchdzh/k8s-debug")
	if err != nil {
		panic(err)
	}

	tags, err := remote.List(repository)
	if err != nil {
		panic(err)
	}

	for _, tag := range tags {
		fmt.Println(tag)
	}
}
