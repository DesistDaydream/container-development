package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// 获取 Image Configuration 文件
func ImageConfiguration(img v1.Image) {
	imageConfiguration, _ := img.ConfigFile()
	indentJsonImageConfiguration, _ := json.MarshalIndent(imageConfiguration, "", "  ")
	err := os.WriteFile("./images/temp/configuration.json", indentJsonImageConfiguration, 0666)
	if err != nil {
		panic(err)
	}
}

// 获取 Image Manifest 文件
func ImageManifest(img v1.Image) {
	imageManifest, _ := img.Manifest()
	indentJsonImageManifest, _ := json.MarshalIndent(imageManifest, "", "  ")
	err := os.WriteFile("./images/temp/manifest.json", indentJsonImageManifest, 0666)
	if err != nil {
		panic(err)
	}
}

// 获取 Image layer
func ImageLayers(img v1.Image) {
	layers, _ := img.Layers()
	for _, layer := range layers {
		// 解压层数据
		// layerIOReader, _ := layer.Uncompressed()

		// 层的摘要
		layerDigest, _ := layer.Digest()
		// 层的大小
		layerSize, _ := layer.Size()
		// 层的媒体类型
		layerMediaType, _ := layer.MediaType()
		fmt.Printf("Digest: %v,\n  媒体类型: %v\n  占用空间: %v\n",
			layerDigest, layerMediaType, layerSize)
	}
}

// 保存镜像
func SaveImage(img v1.Image) {
	// 以 OCI 的 Image Layer 标准，将镜像保存到指定路径下.bolbs 目录、oci-layout 文件、index.json 文件
	err := crane.SaveOCI(img, "./images/k8s-debug")
	// 以 tarball 的方式保存镜像到本地。第二个参数会填写到 tarball 内 manifest.json 文件中的 RepoTags 字段。
	// err := crane.SaveLegacy(img, "lichenhao", "./images/k8s-debug.tar")
	if err != nil {
		panic(err)
	}
}

func main() {
	img, err := crane.Pull("lchdzh/k8s-debug:v1")
	if err != nil {
		panic(err)
	}

	ImageConfiguration(img)
	ImageManifest(img)
	ImageLayers(img)
	// SaveImage(img)
}
