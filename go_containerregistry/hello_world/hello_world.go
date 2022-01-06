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
	indentImageConfiguration, _ := json.MarshalIndent(imageConfiguration, "", "  ")
	os.WriteFile("./go_containerregistry/hello_world/configuration.json", indentImageConfiguration, 0666)
	// fmt.Println(string(indentImageConfiguration))
}

// 获取 Image Manifest 文件
func ImageManifest(img v1.Image) {
	imageManifest, _ := img.Manifest()
	indentImageManifest, _ := json.MarshalIndent(imageManifest, "", "  ")
	os.WriteFile("./go_containerregistry/hello_world/manifest.json", indentImageManifest, 0666)

	// fmt.Println(string(indentImageManifest))
}

// 获取 Image layer
func ImageLayers(img v1.Image) {
	layers, _ := img.Layers()
	for _, layer := range layers {
		// 解压层数据
		// layerIOReader, _ := layer.Uncompressed()

		// 层的
		layer.Digest()
		// 层的大小
		layerSize, _ := layer.Size()
		// 层的媒体类型
		layerMediaType, _ := layer.MediaType()
		fmt.Printf("媒体类型: %v,占用空间: %v\n",
			layerMediaType, layerSize)
	}
}

// 保存镜像
func SaveImage(img v1.Image) {
	// err := crane.SaveOCI(img, "./go_containerregistry/hello_world/k8s-debug")
	// if err != nil {
	// 	panic(err)
	// }

	// 保存
	err := crane.Save(img, "v1", "./go_containerregistry/hello_world/k8s-debug.tar")
	if err != nil {
		panic(err)
	}
}
func main() {
	img, err := crane.Pull("lchdzh/k8s-debug:v1")
	if err != nil {
		panic(err)
	}

	SaveImage(img)
	ImageConfiguration(img)
	ImageManifest(img)
	ImageLayers(img)
}
