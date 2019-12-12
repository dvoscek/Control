## coconut repository add

add a new git repository

### Synopsis

The repository add command adds a git repository
to the catalogue of repositories used for task and workflow configuration.

```
coconut repository add [flags]
```

### Options

```
  -h, --help   help for add
```

### Options inherited from parent commands

```
      --config string            optional configuration file for coconut (default $HOME/.config/coconut/settings.yaml)
      --config_endpoint string   configuration endpoint used by AliECS core as PROTO://HOST:PORT (default "consul://127.0.0.1:8500")
      --endpoint string          AliECS core endpoint as HOST:PORT (default "127.0.0.1:47102")
  -v, --verbose                  show verbose output for debug purposes
```

### SEE ALSO

* [coconut repository](coconut_repository.md)	 - manage git repositories for task and workflow configuration

###### Auto generated by spf13/cobra on 3-Oct-2019