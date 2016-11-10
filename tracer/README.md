## Middleware information

This folder contains a middleware which uses the opentracing library to send
timing statistics/tracing to Zipkin.


## Install

```sh
$ go get -u github.com/iris-contrib/middleware/tracer
```

## How to use

This is a Middleware

```go
New(config ...Config) iris.HandlerFunc
```

Simply add the middleware using `iris.Use()`:

```go
package main

import (
    "github.com/kataras/iris"
    "github.com/iris-contrib/middleware/tracer"
)

func main() {

    iris.Use(tracer.New(tracer.Config{
      // Host:Port where Zipkin can be contacted
      ZipkinHost: "zipkin.mycompany.com:9411",
    }))

    iris.Get("/", func(ctx *iris.Context) {
        ctx.Write("hello")
    })

    iris.Get("/1", func(ctx *iris.Context) {
        ctx.Write("hello")
    })

    iris.Get("/2", func(ctx *iris.Context) {
        ctx.Write("hello")
    })

    iris.Listen(":8080")

}
```
