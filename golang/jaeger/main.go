package main

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"time"
)

func main() {
	closer := initJaeger("in-process")
	defer closer.Close()
	// 获取jaeger tracer
	tracer := opentracing.GlobalTracer()
	// 创建root span
	span := tracer.StartSpan("in-process-service")
	// main执行完结束这个span
	defer span.Finish()
	// 将span传递给Foo
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	Foo(ctx)
}

func Foo(ctx context.Context) {
	// 开始一个span, 设置span的operation_name为Foo
	span, ctx := opentracing.StartSpanFromContext(ctx, "Foo")
	defer span.Finish()
	// 将context传递给Bar
	Bar(ctx)
	// 模拟执行耗时
	time.Sleep(1 * time.Second)
}

func Bar(ctx context.Context) {
	// 开始一个span，设置span的operation_name为Bar
	span, ctx := opentracing.StartSpanFromContext(ctx, "Bar")
	defer span.Finish()

	// 模拟执行耗时
	time.Sleep(2 * time.Second)

	// 假设Bar发生了某些错误
	err := errors.New("something wrong")
	span.LogFields(
		log.String("event", "error"),
		log.String("message", err.Error()),
	)
	span.SetTag("error", true)
}
