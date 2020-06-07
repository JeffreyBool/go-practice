package main

import (
	"math/rand"
	"time"

	prometheusMiddleware "github.com/JeffreyBool/go-practice/pkg/prometheus"
	"github.com/kataras/iris"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	app := iris.New()
	m := prometheusMiddleware.New("serviceName", 0.3, 1.2, 5.0)

	app.Use(m.ServeHTTP)

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		// error code handlers are not sharing the same middleware as other routes, so we have
		// to call them inside their body.
		m.ServeHTTP(ctx)

		ctx.Writef("Not Found")
	})

	app.Get("/", func(ctx iris.Context) {
		sleep := rand.Intn(4999) + 1
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ctx.Writef("Slept for %d milliseconds", sleep)
	})

	app.Get("/metrics", iris.FromStd(promhttp.Handler()))

	// http://localhost:8080/
	// http://localhost:8080/anotfound

	// http://localhost:8080/metrics
	app.Run(iris.Addr("127.0.0.1:8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithCharset("UTF-8"))
}
