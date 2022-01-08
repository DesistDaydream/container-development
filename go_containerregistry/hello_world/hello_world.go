package main

import (
	"fmt"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func main() {
	// 通过镜像的 tag 或者 digest 实例化一个镜像的 Reference。
	// 这里之所以不在 remote.Image() 的第一个参数中只写填写镜像的 tag 或 digest，是为了可以在实例化 Image{} 之前对镜像名称进行一些操作，
	// 比如通过 Reference{},我们可以获取镜像的 注册中心、名称 等等信息，
	// 假如 ParseReference() 的参数是通过外部变量传递进来的，那么在实例化 Image{} 之前，我们可以先分析一下镜像的名称，对其进行过滤。
	ref, _ := name.ParseReference("nginx")

	// 通过镜像的引用实例化 Image{}
	img, _ := remote.Image(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain))

	// 通过实例化的镜像控制镜像，这里是获取镜像所占容量的大小
	imageSize, _ := img.Size()
	fmt.Println(imageSize)
}
