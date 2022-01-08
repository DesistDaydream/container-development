package main

import (
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/sirupsen/logrus"
)

func main() {
	ref, _ := name.ParseReference("lchdzh/k8s-debug:v1")

	logrus.Info(ref.Identifier())
	logrus.Info(ref.Context().Name())
	logrus.Info(ref.Context().String()) // String() 就是直接返回 r.Name()
	logrus.Info(ref.Context().Scheme())
	logrus.Info(ref.Context().RegistryStr())
	logrus.Info(ref.Context().RepositoryStr())
	logrus.Info(ref.Context().Digest(ref.Identifier()))
	logrus.Info(ref.Context().Tag(ref.Identifier()))
	logrus.Info("###################################")
	logrus.Info(ref.Name())
	logrus.Info(ref.Scope("1"))
	logrus.Info(ref.String())
}
