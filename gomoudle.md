<h3>go module介绍</h3>

go module是go官方自带的go依赖管理库,在1.13版本正式推荐使用<br>
go module可以将某个项目(文件夹)下的所有依赖整理成一个 go.mod 文件,里面写入了依赖的版本等 使用go module之后我们可不用将代码放置在src下<br>
使用 go module 管理依赖后会在项目根目录下生成两个文件 go.mod（会记录当前项目的所依赖）和go.sum（记录每个依赖库的版本和哈希值）<br>


<h3>Go mod使用方法</h3>
//初始化模块：<br>
Go mod init <项目模块名称><br>

//依赖关系处理，根据go.mod文件<br>
Go mod tidy<br>

//将依赖包复制到项目的vendor目录<br>
Go mod vendor<br>

//显示依赖关系<br>
Go list -m all<br>

//显示详细依赖关系<br>
Go list -m -json all<br>

//下载依赖<br>
Go mod download [path@version]<br>

<h3>proxy</h3>

安装 golang gin 依赖包的时候，发现长时间没有响应，无法下载，从返回的错误信息看应该是国内无法访问 golang.org。<br>

解决办法<br>
使用国内七牛云的 go module 镜像。<br>

参考 https://github.com/goproxy/goproxy.cn  <br>

golang 1.13 可以直接执行：<br>

go env -w GO111MODULE=on<br>
go env -w GOPROXY=https://goproxy.cn,direct <br>
然后再次使用 go get 下载 gin 依赖就可以了。为七牛云点个赞。<br>

阿里云 Go Module 国内镜像仓库服务<br>
除了七牛云，还可以使用阿里云的 golang 国内镜像。<br>

https://mirrors.aliyun.com/goproxy/ <br>

设置方法<br>
go env -w GO111MODULE=on<br>
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct <br>
