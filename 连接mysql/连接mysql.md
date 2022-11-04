<h3>连接</h3>
Go语言中的database/sql包提供了保证SQL或类SQL数据库的泛用接口，并不提供具体的数据库驱动。<br>
使用database/sql包时必须注入（至少）一个数据库驱动。<br>

我们常用的数据库基本上都有完整的第三方实现。例如：MySQL驱动<br>

<h4>下载依赖</h4>
go get -u github.com/go-sql-driver/mysql <br>

<h4>使用mysql驱动</h4>

db, err := sql.Open("mysql", dsn)<br>
Open打开一个dirverName指定的数据库，dataSourceName指定数据源，一般至少包括数据库文件名和其它连接必要的信息。<br>
Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。<br>

<h5>初始化连接</h5>
如果要检查数据源的名称是否真实有效，应该调用Ping方法。<br>

返回的DB对象可以安全地被多个goroutine并发使用，并且维护其自己的空闲连接池。<br>
因此，Open函数应该仅被调用一次，很少需要关闭这个DB对象。<br>

<h5>SetMaxOpenConns</h5>
func (db *DB) SetMaxOpenConns(n int)<br>
SetMaxOpenConns设置与数据库建立连接的最大数目。 如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。<br>
如果n<=0，不会限制最大开启连接数，默认为0（无限制）。<br>

<h5>SetMaxIdleConns</h5>
func (db *DB) SetMaxIdleConns(n int)<br>
SetMaxIdleConns设置连接池中的最大闲置连接数。 如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。 如果n<=0，不会保留闲置连接。<br>

<h4>mysql预处理</h4>
什么是预处理？<br>
