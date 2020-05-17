package sandbox

import "time"

type Config struct {
	ImageName string

	ContainerName string

	AutoRemove bool

	Verbose bool

	Timeout time.Duration

	EnvVariables []string

	CompileRequired bool

	DoNotRun bool

	AllowInternet bool

	CodeFile string

	OutputPath string

	AllowedMemory int64
}

func GetDefaultSandboxConfig() Config {
	defaultConfig := Config{
		ImageName:       "sandbox",
		ContainerName:   "",
		AutoRemove:      false,
		Verbose:         true,
		Timeout:         time.Duration(10) * time.Second,
		EnvVariables:    []string{},
		CompileRequired: true,
		DoNotRun:        false,
		AllowInternet:   false,
		AllowedMemory:   32,
		OutputPath:      "/tmp",
	}
	return defaultConfig
}
