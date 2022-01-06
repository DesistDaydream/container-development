package main

import (
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ImagesList map[string]string `yaml:"images"`
}

type ImagesTag struct {
	source      string
	destination string
}

func handleSrcImages(srcImages string) (v1.Image, error) {
	// 获取源镜像
	srcRef, err := name.ParseReference(srcImages)
	if err != nil {
		return nil, err
	}

	srcImg, err := remote.Image(srcRef)
	if err != nil {
		return nil, err
	}

	return srcImg, nil
}

func handleDestImages(destImages string) (name.Reference, error) {
	// 指定镜像推送的目标
	destRef, err := name.ParseReference(destImages)
	if err != nil {
		return nil, err
	}
	return destRef, nil
}

func main() {
	username := pflag.String("u", "", "用户名")
	password := pflag.String("p", "", "密码")
	pflag.Parse()

	var config Config
	var imagestag ImagesTag

	imagesListFile, err := os.ReadFile("./go_containerregistry/images_sync/images.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(imagesListFile, &config.ImagesList)
	if err != nil {
		panic(err)
	}

	logrus.Info(config.ImagesList)

	for source, destination := range config.ImagesList {
		imagestag.source = source
		imagestag.destination = destination
		logrus.Info(imagestag.source, imagestag.destination)
		// 处理源镜像与目的镜像
		srcImg, err := handleSrcImages(imagestag.source)
		if err != nil {
			panic(err)
		}
		logrus.Info("源:", srcImg)
		destRef, err := handleDestImages(imagestag.destination)
		if err != nil {
			panic(err)
		}
		logrus.Info("目:", destRef)
	}

	// 实例化认证信息
	auth := authn.FromConfig(authn.AuthConfig{
		Username:      *username,
		Password:      *password,
		Auth:          "",
		IdentityToken: "",
		RegistryToken: "",
	})
	logrus.Debug(auth)

	// destOptions := remote.WithAuth(auth)

	// // 使用 options 中的认证信息，将 img 推送到 ref 中
	// err = remote.Write(destRef, srcImg, destOptions)
	// if err != nil {
	// 	panic(err)
	// }
}
