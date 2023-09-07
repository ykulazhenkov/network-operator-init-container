# network-operator-init-container
Init container for NVIDIA Network Operator

The network-operator-init-container container has two required command line arguments:

 - `--config` path to the configuration file
 - `--node-name` name of the k8s node on which this app runs

The configuration file should be in JSON format:

```
{
  "safeDriverLoad": {
    "enable": true,
    "annotation": "some-annotation"
  }
}
```

- `safeDriverLoad` - contains settings related to safeDriverLoad feature
- `safeDriverLoad.enable` - enable safeDriveLoad feature
- `safeDriverLoad.annotation` - annotation to use for safeDriverLoad feature


If `safeDriverLoad` feature is enabled then the network-operator-init-container container will set annotation
provided in `safeDriverLoad.annotation` on the Kubernetes Node object identified by `--node-name`.
The container exits with code 0 when the annotation is removed from the Node object.

If `safeDriverLoad` feature is disabled then the container will immediately exit with code 0.

```
NVIDIA Network Operator init container

Usage:
  network-operator-init-container [flags]

Config flags:

      --config string                                                                                                                                                                                 
                path to the configuration file
      --node-name string                                                                                                                                                                              
                name of the k8s node on which this app runs

Logging flags:

      --log-flush-frequency duration                                                                                                                                                                  
                Maximum number of seconds between log flushes (default 5s)
      --log-json-info-buffer-size quantity                                                                                                                                                            
                [Alpha] In JSON format with split output streams, the info messages can be buffered for a while to increase performance. The default value of zero bytes disables buffering. The size can
                be specified as number of bytes (512), multiples of 1000 (1K), multiples of 1024 (2Ki), or powers of those (3M, 4G, 5Mi, 6Gi). Enable the LoggingAlphaOptions feature gate to use this.
      --log-json-split-stream                                                                                                                                                                         
                [Alpha] In JSON format, write error messages to stderr and info messages to stdout. The default is to write a single stream to stdout. Enable the LoggingAlphaOptions feature gate to use
                this.
      --logging-format string                                                                                                                                                                         
                Sets the log format. Permitted formats: "json" (gated by LoggingBetaOptions), "text". (default "text")
  -v, --v Level                                                                                                                                                                                       
                number for the log level verbosity
      --vmodule pattern=N,...                                                                                                                                                                         
                comma-separated list of pattern=N settings for file-filtered logging (only works for text log format)

General flags:

  -h, --help                                                                                                                                                                                          
                print help and exit
      --version                                                                                                                                                                                       
                print version and exit

Kubernetes flags:

      --kubeconfig string                                                                                                                                                                             
                Paths to a kubeconfig. Only required if out-of-cluster.

```