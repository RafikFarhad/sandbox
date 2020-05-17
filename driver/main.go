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
	sandboxConfig.Timeout = time.Duration(10) * time.Second
	sandboxConfig.CodeFile = "/home/farhad/go/src/github.com/RafikFarhad/sandbox/container/solution.c"
	sandboxConfig.OutputPath = "/home/farhad/go/src/github.com/RafikFarhad/sandbox/container"

	_ = sandbox.Run(clientInstance, sandboxConfig)

}
