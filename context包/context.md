<h2>GO标准库Context</h2>

<h3>Context 接口</h3>

Context 包的核心就是Context接口<br>

type Context interface { <br>
&emsp;&emsp;&emsp;&emsp;//deadline返回ctx完成工作应该被取消的时间，当deadline没有被设置返回ok==false <br>
&emsp;&emsp;&emsp;&emsp;Deadline() (deadline time.Time, ok bool)<br>
&emsp;&emsp;&emsp;&emsp;//返回一个代表context完成的管道，若是context无法关闭，done返回nil<br>
&emsp;&emsp;&emsp;&emsp;//WithCancel 安排 Done 在调用 cancel 时关闭；<br>
&emsp;&emsp;&emsp;&emsp;//WithDeadline 安排 Done 在截止日期到期时关闭； WithTimeout 安排 Done 在超时过后关闭。<br>
&emsp;&emsp;&emsp;&emsp;Done() <-chan struct{}<br>
&emsp;&emsp;&emsp;&emsp;//若done没有关闭，err返回nil<br>
&emsp;&emsp;&emsp;&emsp;//若doneclosed。err返回非空解释，如canceled如果context被cancel，DeadlineExceeded如果context<br>
&emsp;&emsp;&emsp;&emsp;//的deadline超时<br>
&emsp;&emsp;&emsp;&emsp;Err() error<br>
&emsp;&emsp;&emsp;&emsp;//用于存储与此上下文关联的key值<br>
&emsp;&emsp;&emsp;&emsp;Value(key interface{}) interface{}<br>
}<br>

其中：<br>

Deadline 方法需要返回当前 Context 被取消的时间，也就是完成工作的截止时间（deadline）；<br>
Done 方法需要返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消之后关闭，多次调用 Done 方法会返回同一个Channel；<br>
Err 方法会返回当前 Context 结束的原因，它只会在 Done 返回的 Channel 被关闭时才会返回非空的值：<br>
如果当前 Context 被取消就会返回 Canceled 错误；<br>
如果当前 Context 超时就会返回 DeadlineExceeded 错误；<br>
Value 方法会从 Context 中返回键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法仅用于传递跨 API 和进程间跟请求域的数据。<br>
