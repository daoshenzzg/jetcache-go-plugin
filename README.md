<p>
<a href="https://github.com/mgtv-tech/jetcache-go-plugin/actions"><img src="https://github.com/mgtv-tech/jetcache-go-plugin/workflows/Go/badge.svg" alt="Build Status"></a>
<a href="https://codecov.io/gh/mgtv-tech/jetcache-go-plugin"><img src="https://codecov.io/gh/mgtv-tech/jetcache-go-plugin/master/graph/badge.svg?v=1" alt="codeCov"></a>
<a href="https://goreportcard.com/badge/github.com/mgtv-tech/jetcache-go-plugin"><img src="https://goreportcard.com/badge/github.com/mgtv-tech/jetcache-go-plugin" alt="Go Repport Card"></a>
<a href="https://github.com/mgtv-tech/jetcache-go-plugin/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-green" alt="License"></a>
</p>

# Overview
[jetcache-go-plugin](https://github.com/mgtv-tech/jetcache-go-plugin) 是 [jetcache-go](https://github.com/mgtv-tech/jetcache-go) 维护的插件项目。

# Getting started

## Remote

### goRedisV8Adapter - [redis/go-redis v8](https://github.com/go-redis/redis/v8)
```go
import (
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

TODO

## Encoding

TODO

