# Sandbox

### A basic sandbox implementation built on Go and Docker

### Instruction

### Supported Language

For now, it only supports C, without any dependency. But, we have a plan to extend this as far as we can.

- [x] C 
- [ ] CPP 
- [ ] Go
- [ ] Python2.7 
- [ ] Python3 
- [ ] JavaScript 
- [ ] PHP
- [ ] Rust
- [ ] Bash
- [ ] Java


#### Build The Image

To build image, you need docker

    docker build -t sandbox ~/go/src/github.com/RafikFarhad/sandbox/image

**Note:** If you have other important image named `sandbox` that you can not change, you can tag this image with anything, 
but in that case you must have to provide `config.ImageName` property on `sandboxConfig`.


#### Basic Run

To try this you can follow the `drver\main.go` example file.

Just create a config and pass it to the `sandbox.Run(...)` with your docker client instance.

    	clientInstance := sandbox.GetDockerClient()
    
    	sandboxConfig := sandbox.GetDefaultSandboxConfig()
    
    	sandboxConfig.Timeout = time.Duration(10) * time.Second
    	sandboxConfig.CodeFile = "~/go/src/github.com/RafikFarhad/sandbox/container/solution.c"
    	sandboxConfig.OutputPath = "/tmp"
   
This repo provides a sample code file for testing at `container/solution.c`.

You can tweak different `config` values.

#### SandboxConfig
    
- `CodeFile`:
    
    The code file to run inside the container

- `AllowInternet`: default `false`
    
    Whether our code can access the internet or not.
    
- `Timeout`: default `time.Duration(10) * time.Second` // 10 seconds
    
    You can provide the time limit for execution by this value

- `AllowedMemory`: default `32` // 32 MegaByte
    
    The memory limit for the code execution. The unit is in `MegaByte`
    
- `OutputPath`: default `/tmp`
    
    In this path the `output.txt` file will write the code output.
    
- `ImageName`: default `sandbox`
    
    The image you preferred to run. If you want to run other image rather than this repo's `sandbox` image,
     the `Timeout`, `AllowedMemory` and `AllowInternet` will be used only. But if you want to compile and run a 
     specific code file only you can use this repo's Docker Image. For that default `sandbox` value will do the job.
     But in another case if you build the docker image with another tag, you can provide that here.
    
- `AutoRemove`: default `false`
    
    To remove the container after execution set this flag `true`

- `Verbose`: default `true`
    
    To remove the container after execution set this flag `true`

- `EnvVariables`: default `string[]{}` // Empty array
    
    If your code need any Env Variable you can pass it here as string, like `string[]{'KEY1=val','KEY2=empty'}`

- `CompileRequired`: default `true`
    
    For now this value is not necessary at all. In near future if we implement other language that need not a compilation, we can handle that by this.

- `DoNotRun`: default `false`
    
    If we just want to compile rather than running the code, we can use this.

### Issues

This is a repo all by myself, so that it may have issues or documents may not have been that much clear, but one thing I'm sure
that if you find any bug, issue, feature that can be introduced or any suggestion, I'm open to all of these.

