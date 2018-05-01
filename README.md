Go Dep Demo
===========

Use [dep](https://golang.github.io/dep) to manage dependency.

```
brew install dep
dep ensure
go run hello.go
```

Notice
------

For this project, I use command:

```
dep init
dep ensure -add github.com/pkg/errors
```