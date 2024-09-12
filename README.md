<p>
<a href="https://github.com/mgtv-tech/jetcache-go-plugin/actions"><img src="https://github.com/mgtv-tech/jetcache-go-plugin/workflows/Go/badge.svg" alt="Build Status"></a>
<a href="https://codecov.io/gh/mgtv-tech/jetcache-go-plugin"><img src="https://codecov.io/gh/mgtv-tech/jetcache-go-plugin/master/graph/badge.svg?v=1" alt="codeCov"></a>
<a href="https://goreportcard.com/badge/github.com/mgtv-tech/jetcache-go-plugin"><img src="https://goreportcard.com/badge/github.com/mgtv-tech/jetcache-go-plugin?v=1" alt="Go Repport Card"></a>
<a href="https://github.com/mgtv-tech/jetcache-go-plugin/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-green" alt="License"></a>
</p>

# Overview
[jetcache-go-plugin](https://github.com/mgtv-tech/jetcache-go-plugin) 是 [jetcache-go](https://github.com/mgtv-tech/jetcache-go) 维护的插件项目。

# Getting started

## Remote Adapter

### [redis/go-redis v8](https://github.com/go-redis/redis/v8) 
```go
import (
    "github.com/mgtv-tech/jetcache-go"
    "github.com/mgtv-tech/jetcache-go-plugin/remote"
)

mycache := cache.New(cache.WithName("any"),
    cache.WithRemote(remote.NewGoRedisV8Adapter(ring)),
    cache.WithLocal(local.NewFreeCache(256*local.MB, time.Minute)),
    // ...
)
```

## Local

TODO

## Stats

### [prometheus](https://prometheus.io/)
```go
import (
    "github.com/mgtv-tech/jetcache-go"
    "github.com/mgtv-tech/jetcache-go-plugin/remote"
    pstats "github.com/mgtv-tech/jetcache-go-plugin/stats"
    "github.com/mgtv-tech/jetcache-go/stats"
)

cacheName := "demo"
jetcache := cache.New(cache.WithRemote(remote.NewGoRedisV8Adapter(ring)),
    cache.WithStatsHandler(
        stats.NewHandles(false,
            stats.NewStatsLogger(cacheName), 
            pstats.NewPrometheus(cacheName))))
```
> 同时集成日志统计和Prometheus统计。

## Encoding

TODO

