package main

import (
	"fmt"
	"time"

	"github.com/RafikFarhad/sandbox"
)

func main() {
	fmt.Println("... Hello World ...")
	clientInstance := sandbox.GetDockerClient()
	imageName := "sandbox"

	sandboxConfig := sandbox.GetDefaultSandboxConfig()

	sandboxConfig.ImageName = imageName
	sandboxConfig.Timeout = time.Duration(6) * time.Second
	sandboxConfig.HostPWD = "/home/farhad/go/src/github.com/RafikFarhad/sandbox"

	_ = sandbox.Run(clientInstance, sandboxConfig)

	//sandbox.GetContainerStats(clientInstance, containerId)
}
