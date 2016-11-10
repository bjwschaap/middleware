package tracer

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris"
	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

type tracerMiddleware struct {
	config Config
}

// Serve serves the middleware
func (tm *tracerMiddleware) Serve(ctx *iris.Context) {
	// Create a opentracing.Tracer that sends data to Zipkin
	collector, _ := zipkin.NewHTTPCollector(
		fmt.Sprintf("http://%s/api/v1/spans", tm.config.ZipkinHost))
	tracer, _ := zipkin.NewTracer(
		zipkin.NewRecorder(collector, false, "127.0.0.1:0", tm.config.ServiceName))

	parent := tracer.StartSpan("Request")

	child := tracer.StartSpan("Handler", opentracing.ChildOf(parent.Context()))
	// Call next handler on the stack
	ctx.Next()
	child.Finish()

	parent.LogEvent(fmt.Sprintf("%s %v %s %s %s",
		time.Now().Format("01/02 - 15:04:05"),
		strconv.Itoa(ctx.Response.StatusCode()),
		ctx.RemoteAddr(),
		ctx.MethodString(),
		ctx.PathString()))
	parent.Finish()
}

// New returns the logger middleware
// receives optional configs(logger.Config)
func New(cfg ...Config) iris.HandlerFunc {
	c := DefaultConfig().Merge(cfg)
	l := &tracerMiddleware{config: c}

	return l.Serve
}
