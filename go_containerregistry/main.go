package main

import (
	"fmt"

	"github.com/google/go-containerregistry/pkg/crane"
)

func main() {
	img, err := crane.Pull("nginx")
	if err != nil {
		panic(err)
	}

	// 获取 Image Configuration 文件
	// imageConfiguration, _ := img.ConfigFile()
	// indentImageConfiguration, _ := json.MarshalIndent(imageConfiguration, "", "  ")
	// fmt.Println(string(indentImageConfiguration))

	// 获取 Image Manifest 文件
	// imageManifest, _ := img.Manifest()
	// indentImageManifest, _ := json.MarshalIndent(imageManifest, "", "  ")
	// fmt.Println(string(indentImageManifest))

	// 获取 Image layer
	layers, _ := img.Layers()
	for _, layer := range layers {
		// 解压层数据
		// layerIOReader, _ := layer.Uncompressed()

		// 层的大小
		layerSize, _ := layer.Size()
		fmt.Println(layerSize)
		// 层的媒体类型
		layerMediaType, _ := layer.MediaType()
		fmt.Println(layerMediaType)

	}

}
