package main

import (
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/spf13/pflag"
)

func main() {
	username := pflag.String("u", "", "用户名")
	password := pflag.String("p", "", "密码")
	pflag.Parse()

	// 实例化认证信息
	auth := authn.FromConfig(authn.AuthConfig{
		Username:      *username,
		Password:      *password,
		Auth:          "",
		IdentityToken: "",
		RegistryToken: "",
	})

	destOptions := remote.WithAuth(auth)

	// 获取源镜像
	srcRef, err := name.ParseReference("nginx")
	if err != nil {
		panic(err)
	}

	srcImg, err := remote.Image(srcRef)
	if err != nil {
		panic(err)
	}

	// 指定镜像推送的目标
	destRef, err := name.ParseReference("lchdzh/nginx:v1")
	if err != nil {
		panic(err)
	}

	// 使用 options 中的认证信息，将 img 推送到 ref 中
	err = remote.Write(destRef, srcImg, destOptions)
	if err != nil {
		panic(err)
	}
}
