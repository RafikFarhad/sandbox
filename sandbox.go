package sandbox

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func GetDockerClient() *client.Client {

	clientInstance, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return clientInstance
}

func appendNecessaryFlag(sandboxConfig *Config) {
	if sandboxConfig.CompileRequired {
		sandboxConfig.EnvVariables = append(sandboxConfig.EnvVariables, "COMPILE=YES")
	}
	if !sandboxConfig.DoNotRun {
		sandboxConfig.EnvVariables = append(sandboxConfig.EnvVariables, "RUN=YES")
	}
	sandboxConfig.EnvVariables = append(sandboxConfig.EnvVariables, "ALLOWED_MEMORY="+string(sandboxConfig.AllowedMemory*1e+6))
}

func Run(clientInstance *client.Client, sandboxConfig Config) string {
	appendNecessaryFlag(&sandboxConfig)
	ctx := context.Background()
	containerConfig := container.Config{
		Image: sandboxConfig.ImageName,
		Env:   sandboxConfig.EnvVariables,
	}
	hostConfig := container.HostConfig{
		Mounts: []mount.Mount{
			mount.Mount{
				Type:   mount.TypeBind,
				Source: sandboxConfig.OutputPath,
				Target: "/home/sandbox/raw",
			},
			mount.Mount{
				Type:   mount.TypeBind,
				Source: sandboxConfig.CodeFile,
				Target: "/home/sandbox/raw/code_file.c",
			},
		},
		Resources: container.Resources{
			Memory: sandboxConfig.AllowedMemory * 1e+6,
			CpusetCpus: "0",
		},
		AutoRemove:  sandboxConfig.AutoRemove,
		NetworkMode: container.NetworkMode(map[bool]string{true: "bridge", false: "none"}[sandboxConfig.AllowInternet]),
	}
	nc := network.NetworkingConfig{}
	limitCtx, cancel := context.WithTimeout(context.Background(), sandboxConfig.Timeout)
	defer cancel() // required if the assigned job is finished before the timeout

	containerBody, err := clientInstance.ContainerCreate(
		ctx,
		&containerConfig,
		&hostConfig,
		&nc,
		"",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(containerBody.ID)

	if err := clientInstance.ContainerStart(ctx, containerBody.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	//GetContainerStats(clientInstance, containerBody.ID)
	var statusCode int64 = -1
	statusChannel, errorChannel := clientInstance.ContainerWait(limitCtx, containerBody.ID, container.WaitConditionNextExit)
	select {
	case err := <-errorChannel:
		{
			fmt.Println("Error occurred")
			if err != nil {
				fmt.Println(err)
			}
		}
	case output := <-statusChannel:
		statusCode = output.StatusCode
		fmt.Println("Status: ", output.StatusCode)
	}

	switch statusCode {
	case -1:
		fmt.Println("Timed out")
		stopTimeout := time.Second * 5 // 5 second is timeout for stopping the container
		err := clientInstance.ContainerStop(ctx, containerBody.ID, &stopTimeout)
		if err != nil {
			fmt.Println("Container not stopped")
		}
		break
	case 139:
		fmt.Println("Memory limit exceeded")
		break
	case 254:
		fmt.Println("Compile failed")
		break
	default:
		if sandboxConfig.Verbose {
			reader, err := clientInstance.ContainerLogs(ctx, containerBody.ID, types.ContainerLogsOptions{
				ShowStderr: true,
				ShowStdout: true,
			})
			if err != nil {
				fmt.Println(err)
			}
			defer reader.Close()
			content, _ := ioutil.ReadAll(reader)
			fmt.Println(string(content))
		}
		fmt.Println("Running succeeded")
	}
	return containerBody.ID
}

func GetContainerStats(clientInstance *client.Client, containerId string) {
	fmt.Println("Stats for: ", containerId)
	stats, err := clientInstance.ContainerStatsOneShot(context.Background(), containerId)
	if err != nil {
		fmt.Println(err)
	}
	defer stats.Body.Close()
	content, _ := ioutil.ReadAll(stats.Body)
	fmt.Println(string(content))

}
