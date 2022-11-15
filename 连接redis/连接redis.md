<h4>安装</h4>
Go 社区中目前有很多成熟的 redis client 库，比如[https://github.com/gomodule/redigo 和https://github.com/go-redis/redis<br>
使用以下命令下安装 go-redis 库:<br>
go get github.com/go-redis/redis/v8<br>

<h4>普通连接</h4>
go-redis 库中使用 redis.NewClient 函数连接 Redis 服务器。<br>

rdb := redis.NewClient(&redis.Options{  <br>
&emsp;&emsp;&emsp;&emsp;Addr:     "localhost:6379",  <br>
&emsp;&emsp;&emsp;&emsp;Password: "", // 密码  <br>
&emsp;&emsp;&emsp;&emsp;DB:       0,  // 数据库<br>
&emsp;&emsp;&emsp;&emsp;PoolSize: 20, // 连接池大小<br>
})<br>

除此之外，还可以使用 redis.ParseURL 函数从表示数据源的字符串中解析得到 Redis 服务器的配置信息。<br>
opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")<br>
if err != nil {<br>
&emsp;&emsp;&emsp;&emsp;panic(err)<br>
}<br>

rdb := redis.NewClient(opt)<br>

<h4>TLS连接模式</h4>
如果使用的是 TLS 连接方式，则需要使用 tls.Config 配置。<br>

rdb := redis.NewClient(&redis.Options{<br>
&emsp;&emsp;&emsp;TLSConfig: &tls.Config{<br>
&emsp;&emsp;&emsp;&emsp;MinVersion: tls.VersionTLS12,<br>
// Certificates: []tls.Certificate{cert},<br>
// ServerName: "your.domain.com",<br>
},<br>
})<br>

<h4>Redis Sentinel模式</h4>
使用下面的命令连接到由 Redis Sentinel 管理的 Redis 服务器。<br>

rdb := redis.NewFailoverClient(&redis.FailoverOptions{<br>
&emsp;&emsp;&emsp;&emsp;MasterName:    "master-name",<br>
&emsp;&emsp;&emsp;&emsp;SentinelAddrs: []string{":9126", ":9127", ":9128"},<br>
})<br>

<h4>Redis Cluster模式</h4>
使用下面的命令连接到 Redis Cluster，go-redis 支持按延迟或随机路由命令。<br>

rdb := redis.NewClusterClient(&redis.ClusterOptions{<br>
&emsp;&emsp;&emsp;&emsp;Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},<br>
    // 若要根据延迟或随机路由命令，请启用以下命令之一<br>
    // RouteByLatency: true,<br>
    // RouteRandomly: true,<br>
})<br>

<h4>有序集合</h4>
go redis有序集合常用函数:<br>
&emsp;&emsp;&emsp;&emsp;ZAdd - 添加一个或者多个元素到集合，如果元素已经存在则更新分数<br>
&emsp;&emsp;&emsp;&emsp;ZCard - 返回集合元素个数<br>
&emsp;&emsp;&emsp;&emsp;ZCount - 统计某个分数范围内的元素个数<br>
&emsp;&emsp;&emsp;&emsp;ZIncrBy - 增加元素的分数<br>
&emsp;&emsp;&emsp;&emsp;ZRange,ZRevRange - 返回集合中某个索引范围的元素，根据分数从小到大排序<br>
&emsp;&emsp;&emsp;&emsp;ZRangeByScore,ZRevRangeByScore - 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。<br>
&emsp;&emsp;&emsp;&emsp;ZRem - 删除集合元素<br>
&emsp;&emsp;&emsp;&emsp;ZRemRangeByRank - 根据索引范围删除元素<br>
&emsp;&emsp;&emsp;&emsp;ZRemRangeByScore - 根据分数范围删除元素<br>
&emsp;&emsp;&emsp;&emsp;ZScore - 查询元素对应的分数<br>
&emsp;&emsp;&emsp;&emsp;ZRank, ZRevRank - 查询元素的排名<br>


<h4>遍历扫描所有key</h4>
Scan函数

<h4>pipeline单个client执行多条命令</h4>

<h4>处理redis事务</h4>
Redis 是单线程执行命令的，因此单个命令始终是原子的，但是来自不同客户端的两个给定命令可以依次执行，例如在它们之间交替执行。<br>
但是，Multi/exec能够确保在multi/exec两个语句之间的命令之间没有其他客户端正在执行命令。<br>
在这种场景我们需要使用 TxPipeline 或 TxPipelined 方法将 pipeline 命令使用 MULTI 和EXEC包裹起来。<br>

<h4>WATCH</h4>
从使用WATCH命令监视某个 key 开始，直到执行EXEC命令的这段时间里，如果有其他用户抢先对被监视的 key 进行了替换、更新、删除等操作，<br>
那么当用户尝试执行EXEC的时候，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。<br>

