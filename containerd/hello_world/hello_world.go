package main

import (
	"context"
	"fmt"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

func helloWorld() error {
	// 通过 containerd 的 socket 实例化一个客户端以连接 containerd
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()

	// 使用 namespace 创建一个 context
	ctx := namespaces.WithNamespace(context.Background(), "default")

	// 通过 containerd 的客户端实例化一个镜像存储器
	imagesStore := client.ImageService()
	// 列出 ctx 环境中的所有镜像
	images, _ := imagesStore.List(ctx)
	for _, image := range images {
		fmt.Println(image.Name)
	}

	// 通过 containerd 的客户端实例化一个容器存储器
	containersStroe := client.ContainerService()
	// 列出 ctx 环境中的所有容器
	containers, _ := containersStroe.List(ctx)
	for _, container := range containers {
		fmt.Println(container.Spec)
	}
	return nil
}

func main() {
	if err := helloWorld(); err != nil {
		log.Fatal(err)
	}
}
