<h3>Gin框架介绍</h3>
Gin是一个用Go语言编写的web框架。它是一个类似于martini但拥有更好性能的API框架, 由于使用了httprouter，速度提高了近40倍。<br>
Go世界里最流行的Web框架，Github上有32K+star。 基于httprouter开发的Web框架。 中文文档齐全，简单易用的轻量级框架。

<h3>Restful API</h3>


<h3>HTML渲染</h3>
Gin框架中使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染。

<h3>自定义模板函数</h3>



<h3>处理静态文件</h3>

func main() {<br>
&emsp;&emsp;&emsp;&emsp;r := gin.Default()<br>
&emsp;&emsp;&emsp;&emsp;r.Static("/static", "./static")<br>
&emsp;&emsp;&emsp;&emsp;r.LoadHTMLGlob("templates/**/*")<br>
&emsp;&emsp;&emsp;&emsp;// ...<br>
&emsp;&emsp;&emsp;&emsp;r.Run(":8080")<br>
}<br>

<h3>模板继承</h3>

<h3>补充文件路径处理</h3>
func getCurrentPath() string {<br>
&emsp;&emsp;&emsp;&emsp;if ex, err := os.Executable(); err == nil {<br>
&emsp;&emsp;&emsp;&emsp;return filepath.Dir(ex)<br>
&emsp;&emsp;&emsp;&emsp;}<br>
&emsp;&emsp;&emsp;&emsp;return "./"<br>
}<br>

<h3>json渲染</h3>

<h3>XML渲染</h3>

<h3>YAML渲染</h3>

<h3>protobuf渲染</h3>

<h2>获取参数</h2>

<h3>获取querystring参数</h3>
c.DefaultQuery()
c.Query()

<h3>获取form表单</h3>
c.DefaultPostForm()
c.PostForm()

<h3>获取json参数</h3>

<h3>获取path参数</h3>

<h3>参数绑定</h3>
提取请求中 QueryString、form表单、json、xml等参数到结构体中<br>
shoudbind方法:<br>

type Login struct {<br>
&emsp;&emsp;&emsp;&emsp;    User  string  `json: "name" uri: "name" binding: "required"`<br>
&emsp;&emsp;&emsp;&emsp;    Password  string  `json: "password" uri: "password" binding: "required"`<br>
}<br>

var login Login<br>

err := c.ShouldBindJSON(&login)<br>
err := c.ShouldBindUri(&login)<br>
err := c.ShouldBindQuery(&login)<br>

binding: "required" 参数：<br>该标签表示要求绑定不能为空的值<br>

自定义验证器：<br>

<h3>单个文件上传</h3>

<h3>多个文件上传</h3>

<h3>HTTP重定向</h3>

<h3>路由重定向</h3>

<h3>Gin中间件</h3>

<h3>运行多个服务</h3>

<h3>gin框架路由拆分与注册</h3>
路由拆分成多个文件：<br>
路由拆分到不同的APP：<br>

<h3>在gin框架中使用JWT认证</h3>
<h4>什么是JWT？</h4>
JWT全称JSON WEB TOKEN 是一种跨域认证解决方案，属于一个开放标准；<br>
如何使用JWT：<br>
1.创建一个JWT
2.解析一个JWT