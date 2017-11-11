This is a workload attestor plugin for docker. It works with the [SPIRE implementation of SPIFFE](https://github.com/spiffe/spire). Dependencies are managed by golang's dep tool.

As a prerequisite to working with this repository, please install [dep](https://github.com/golang/dep) and run:
```
dep ensure
```

To test, please run:
```
go test -cover .
```

To build an executable, please run:
```
go build -o dockerplugin
```
