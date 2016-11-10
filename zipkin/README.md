## Middleware information

This folder contains a middleware which uses the opentracing library to send
timing statistics/tracing to Zipkin.


## Install

```sh
$ go get -u github.com/iris-contrib/middleware/zipkin
```

## How to use
This is a Middleware

```go
New(config ...Config) iris.HandlerFunc
```

Add it using `iris.Use()`:

```go
func main() {

    iris.Use(tracer.New(tracer.Config{
        ZipkinEndpoint: "zipkin.mycompany.com:9411",
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
