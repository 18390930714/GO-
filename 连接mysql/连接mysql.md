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

普通SQL语句执行过程：<br>

客户端对SQL语句进行占位符替换得到完整的SQL语句。<br>
客户端发送完整SQL语句到MySQL服务端<br>
MySQL服务端执行完整的SQL语句并将结果返回给客户端。<br>
预处理执行过程：<br>

把SQL语句分成两部分，命令部分与数据部分。<br>
先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。<br>
然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。<br>
MySQL服务端执行完整的SQL语句并将结果返回给客户端。<br>

<h4>为什么要预处理？</h4>
优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。<br>
避免SQL注入问题。<br>

<h4>sql 注入问题,我们任何时候都不应该自己拼接SQL语句！</h4>

<h4>go语言处理mysql事务</h4>
事务：一个最小的不可再分的工作单元；通常一个事务对应一个完整的业务(例如银行账户转账业务，该业务就是一个最小的工作单元)，<br>
同时这个完整的业务需要执行多次的DML(insert、update、delete)语句共同联合完成。A转账给B，这里面就需要执行两次update操作。<br>

在MySQL中只有使用了Innodb数据库引擎的数据库或表才支持事务。事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，<br>
要么全部不执行。

<h4>事务的ACID</h4>
通常事务必须满足4个条件（ACID）：原子性（Atomicity，或称不可分割性）、一致性（Consistency）、隔离性（Isolation，又称独立性）、持久性（Durability）。<br>

原子性:   一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。<br>
一致性:	在事务开始之前和事务结束以后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精确度、串联性以及后续数据库可以自发性地完成预定的工作。<br>
隔离性:	数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable）。<br>
持久性:	事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。<br>

<h4>事务相关方法</h4>
Go语言中使用以下三个方法实现MySQL中的事务操作。 <br>
开始事务<br>
func (db *DB) Begin() (*Tx, error)<br>

提交事务<br>
func (tx *Tx) Commit() error <br>
回滚事务 <br>
func (tx *Tx) Rollback() error<br>