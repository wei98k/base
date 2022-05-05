# Overview

Package sql provides a generic interface around SQL (or SQL-like) databases.

The sql package must be used in conjunction with a database driver. See https://golang.org/s/sqldrivers for a list of drivers.

Drivers that do not support context cancellation will not return until after the query is completed.

For usage examples, see the wiki page at https://golang.org/s/sqlwiki.

包sql提供了一个围绕SQL（或类似SQL）数据库的通用接口。

sql包必须与数据库驱动一起使用。参见https://golang.org/s/sqldrivers，获取驱动程序的列表。

不支持上下文取消的驱动程序在查询完成后才会返回。

关于使用实例，请参见wiki页面https://golang.org/s/sqlwiki。

## Variables

`var ErrConnDone = errors.New("sql: connection is already closed")`

ErrConnDone is returned by any operation that is performed on a connection that has already been returned to the connection pool.

ErrConnDone由已经返回到连接池的连接上执行的任何操作返回。

`var ErrNoRows = errors.New("sql: no rows in result set")`

当QueryRow不返回一行时，Scan返回errnorlows。在这种情况下，QueryRow返回一个占位符*行值，将此错误推迟到Scan。

`var ErrTxDone = errors.New("sql: transaction has already been committed or rolled back")`

ErrTxDone由对已经提交或回滚的事务执行的任何操作返回。

## Functions

### func Drivers

`func Drivers() []string`

Drivers returns a sorted list of the names of the registered drivers.

驱动程序返回一个已注册的驱动程序名称的排序列表。

### func Register

`func Register(name string, driver driver.Driver)`

Register makes a database driver available by the provided name. If Register is called twice with the same name or if driver is nil, it panics.

Register通过提供的名称使一个数据库驱动可用。如果Register以相同的名字被调用两次，或者driver为nil，那么它就会陷入恐慌。

## Types

## type ColumnType

```
type ColumnType struct {
	// contains filtered or unexported fields
}
```

ColumnType contains the name and type of a column.

ColumnType包含一个列的名称和类型。


### func (*ColumnType) DatabaseTypeName

`func (ci *ColumnType) DatabaseTypeName() string`

DatabaseTypeName returns the database system name of the column type. If an empty string is returned, then the driver type name is not supported. Consult your driver documentation for a list of driver data types. Length specifiers are not included. Common type names include "VARCHAR", "TEXT", "NVARCHAR", "DECIMAL", "BOOL", "INT", and "BIGINT".

DatabaseTypeName返回列类型的数据库系统名称。如果返回的是一个空字符串，那么说明不支持该驱动类型的名称。关于驱动数据类型的列表，请查阅你的驱动文档。长度指定器不包括在内。常见的类型名称包括 "VARCHAR"，"TEXT"，"NVARCHAR"，"DECIMAL"，"BOOL"，"INT"，和 "BIGINT"。

### func (*ColumnType) DecimalSize

`func (ci *ColumnType) DecimalSize() (precision, scale int64, ok bool)`

DecimalSize returns the scale and precision of a decimal type. If not applicable or if not supported ok is false.

DecimalSize返回一个小数类型的比例和精度。如果不适用或不支持，ok为false。

### func (*ColumnType) Length

`func (ci *ColumnType) Length() (length int64, ok bool)`

Length returns the column type length for variable length column types such as text and binary field types. If the type length is unbounded the value will be math.MaxInt64 (any database limits will still apply). If the column type is not variable length, such as an int, or if not supported by the driver ok is false.

对于可变长度的列类型，如文本和二进制字段类型，Length返回列类型的长度。如果类型的长度是无界的，其值将是math.MaxInt64（任何数据库限制仍将适用）。如果列的类型不是可变长度的，比如int，或者驱动程序不支持，ok就是false。

### func (*ColumnType) Name

`func (ci *ColumnType) Name() string`

Name returns the name or alias of the column.

Name返回该列的名称或别名。

### func (*ColumnType) Nullable

`func (ci *ColumnType) Nullable() (nullable, ok bool)`

Nullable reports whether the column may be null. If a driver does not support this property ok will be false.

Nullable报告该列是否可以为空。如果一个驱动程序不支持这个属性，ok将是false。

### func (*ColumnType) ScanType

`func (ci *ColumnType) ScanType() reflect.Type`

ScanType returns a Go type suitable for scanning into using Rows.Scan. If a driver does not support this property ScanType will return the type of an empty interface.

ScanType返回一个适合使用Rows.Scan扫描到的Go类型。如果驱动程序不支持此属性，ScanType将返回一个空接口的类型。

## type Conn

```golang
type Conn struct {
	// contains filtered or unexported fields
}
```
Conn represents a single database connection rather than a pool of database connections. Prefer running queries from DB unless there is a specific need for a continuous single database connection.

A Conn must call Close to return the connection to the database pool and may do so concurrently with a running query.

After a call to Close, all operations on the connection fail with ErrConnDone.

Conn代表一个单一的数据库连接，而不是一个数据库连接池。除非对连续的单一数据库连接有特殊需要，否则更倾向于从DB运行查询。

Conn必须调用Close来将连接返回到数据库池中，并且可以与正在运行的查询同时进行。

在调用Close后，对连接的所有操作都以ErrConnDone失败。

### func (*Conn) BeginTx

`func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)`

BeginTx starts a transaction.

The provided context is used until the transaction is committed or rolled back. If the context is canceled, the sql package will roll back the transaction. Tx.Commit will return an error if the context provided to BeginTx is canceled.

The provided TxOptions is optional and may be nil if defaults should be used. If a non-default isolation level is used that the driver doesn't support, an error will be returned.

BeginTx开始一个事务。

所提供的上下文被使用，直到事务被提交或回滚。如果上下文被取消，sql包将回滚该事务。如果提供给BeginTx的上下文被取消，Tx.Commit将返回一个错误。

提供的TxOptions是可选的，如果应该使用默认值，可以是nil。如果使用了驱动程序不支持的非默认隔离级别，将返回一个错误。

### func (*Conn) Close

`func (c *Conn) Close() error`

Close returns the connection to the connection pool. All operations after a Close will return with ErrConnDone. Close is safe to call concurrently with other operations and will block until all other operations finish. It may be useful to first cancel any used context and then call close directly after.

Close将连接返回到连接池中。在Close之后的所有操作都将以ErrConnDone返回。Close与其他操作同时调用是安全的，它将阻塞直到所有其他操作完成。首先取消任何已使用的上下文，然后直接调用Close可能是有用的。

### func (*Conn) ExecContext

`func (c *Conn) ExecContext(ctx context.Context, query string, args ...any) (Result, error)`

ExecContext executes a query without returning any rows. The args are for any placeholder parameters in the query.

ExecContext执行一个查询，不返回任何行。args是用于查询中的任何占位参数。

### func (*Conn) PingContext

`func (c *Conn) PingContext(ctx context.Context) error`

PingContext verifies the connection to the database is still alive.

PingContext验证了与数据库的连接仍然有效。

### func (*Conn) PrepareContext

`func (c *Conn) PrepareContext(ctx context.Context, query string) (*Stmt, error)`

PrepareContext creates a prepared statement for later queries or executions. Multiple queries or executions may be run concurrently from the returned statement. The caller must call the statement's Close method when the statement is no longer needed.

The provided context is used for the preparation of the statement, not for the execution of the statement.

PrepareContext为以后的查询或执行创建一个准备好的语句。多个查询或执行可以在返回的语句中同时运行。当不再需要该语句时，调用者必须调用该语句的关闭方法。

所提供的上下文用于准备语句，而不是用于执行语句。

### func (*Conn) QueryContext

`func (c *Conn) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)`

QueryContext executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.

QueryContext执行一个返回行的查询，通常是一个SELECT。args是用于查询中的任何占位参数。

### func (*Conn) QueryRowContext

`func (c *Conn) QueryRowContext(ctx context.Context, query string, args ...any) *Row`

QueryRowContext executes a query that is expected to return at most one row. QueryRowContext always returns a non-nil value. Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.

QueryRowContext执行一个查询，预计最多返回一条记录。QueryRowContext总是返回一个非零的值。错误被推迟到Row的Scan方法被调用。如果查询没有选择任何行，*Row的扫描将返回ErrNoRows。否则，*Row's Scan会扫描第一条被选择的行，并丢弃其余的行。

### func (*Conn) Raw

`func (c *Conn) Raw(f func(driverConn any) error) (err error)`

Raw executes f exposing the underlying driver connection for the duration of f. The driverConn must not be used outside of f.

Once f returns and err is not driver.ErrBadConn, the Conn will continue to be usable until Conn.Close is called.

Raw执行f，在f的持续时间内暴露底层的驱动连接。

一旦f返回并且err不是driver.ErrBadConn，Conn将继续可用，直到Conn.Close被调用。

## type DB

```
type DB struct {
	// contains filtered or unexported fields
}
```

DB is a database handle representing a pool of zero or more underlying connections. It's safe for concurrent use by multiple goroutines.

The sql package creates and frees connections automatically; it also maintains a free pool of idle connections. If the database has a concept of per-connection state, such state can be reliably observed within a transaction (Tx) or connection (Conn). Once DB.Begin is called, the returned Tx is bound to a single connection. Once Commit or Rollback is called on the transaction, that transaction's connection is returned to DB's idle connection pool. The pool size can be controlled with SetMaxIdleConns.

DB是一个数据库句柄，代表一个由零个或多个底层连接组成的池。它对于多个goroutine的并发使用是安全的。

sql包自动创建和释放连接；它还维护一个空闲的连接池。如果数据库有一个每个连接状态的概念，这种状态可以在一个事务（Tx）或连接（Conn）中被可靠地观察到。一旦DB.Begin被调用，返回的Tx被绑定到一个单一的连接。一旦在事务上调用Commit或Rollback，该事务的连接就会返回到DB的闲置连接池。池的大小可以用SetMaxIdleConns来控制。

### func Open

`func Open(driverName, dataSourceName string) (*DB, error)`

Open opens a database specified by its database driver name and a driver-specific data source name, usually consisting of at least a database name and connection information.

Most users will open a database via a driver-specific connection helper function that returns a *DB. No database drivers are included in the Go standard library. See https://golang.org/s/sqldrivers for a list of third-party drivers.

Open may just validate its arguments without creating a connection to the database. To verify that the data source name is valid, call Ping.

The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections. Thus, the Open function should be called just once. It is rarely necessary to close a DB.

打开一个由其数据库驱动名称和特定驱动数据源名称指定的数据库，通常至少包括一个数据库名称和连接信息。

大多数用户会通过一个返回*DB的特定驱动程序连接辅助函数来打开数据库。Go标准库中没有包含数据库驱动。有关第三方驱动的列表，请参见https://golang.org/s/sqldrivers。

Open可能只是验证其参数，而不创建与数据库的连接。要验证数据源名称是否有效，请调用Ping。

返回的DB对于多个goroutine的并发使用是安全的，并且维护它自己的空闲连接池。因此，Open函数应该只被调用一次。很少有必要关闭一个数据库。

### func OpenDB

`func OpenDB(c driver.Connector) *DB`

OpenDB opens a database using a Connector, allowing drivers to bypass a string based data source name.

Most users will open a database via a driver-specific connection helper function that returns a *DB. No database drivers are included in the Go standard library. See https://golang.org/s/sqldrivers for a list of third-party drivers.

OpenDB may just validate its arguments without creating a connection to the database. To verify that the data source name is valid, call Ping.

The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections. Thus, the OpenDB function should be called just once. It is rarely necessary to close a DB.

OpenDB使用连接器打开数据库，允许驱动程序绕过基于字符串的数据源名称。

大多数用户会通过一个特定于驱动程序的连接辅助函数来打开数据库，该函数返回一个*DB。Go标准库中没有包含数据库驱动。有关第三方驱动的列表，请参见https://golang.org/s/sqldrivers。

OpenDB可能只是验证其参数，而不创建与数据库的连接。要验证数据源名称是否有效，请调用Ping。

返回的DB对于多个goroutine的并发使用是安全的，并且维护自己的空闲连接池。因此，OpenDB函数应该只被调用一次。很少有必要关闭一个数据库。

### func (*DB) Begin

`func (db *DB) Begin() (*Tx, error)`

Begin starts a transaction. The default isolation level is dependent on the driver.

Begin uses context.Background internally; to specify the context, use BeginTx.

Begin启动一个事务。默认的隔离级别取决于驱动程序。

Begin在内部使用context.background；要指定上下文，请使用BeginTx。

### func (*DB) BeginTx

`func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)`

BeginTx starts a transaction.

The provided context is used until the transaction is committed or rolled back. If the context is canceled, the sql package will roll back the transaction. Tx.Commit will return an error if the context provided to BeginTx is canceled.

The provided TxOptions is optional and may be nil if defaults should be used. If a non-default isolation level is used that the driver doesn't support, an error will be returned.

BeginTx开始一个事务。

所提供的上下文被使用，直到事务被提交或回滚。如果上下文被取消，sql包将回滚该事务。如果提供给BeginTx的上下文被取消，Tx.Commit将返回一个错误。

提供的TxOptions是可选的，如果应该使用默认值，可以是nil。如果使用了驱动程序不支持的非默认隔离级别，将返回一个错误。

### func (*DB) Close

`func (db *DB) Close() error`

Close closes the database and prevents new queries from starting. Close then waits for all queries that have started processing on the server to finish.

It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.

Close关闭数据库，防止新的查询开始。然后Close等待所有已经在服务器上开始处理的查询完成。

关闭一个数据库是很罕见的，因为数据库句柄是长期存在的，并且在许多goroutine之间共享。

### func (*DB) Conn

`func (db *DB) Conn(ctx context.Context) (*Conn, error)`

Conn returns a single connection by either opening a new connection or returning an existing connection from the connection pool. Conn will block until either a connection is returned or ctx is canceled. Queries run on the same Conn will be run in the same database session.

Every Conn must be returned to the database pool after use by calling Conn.Close.

Conn通过打开一个新的连接或从连接池返回一个现有的连接来返回一个单一的连接。Conn将阻塞，直到返回一个连接或取消ctx。在同一个Conn上运行的查询将在同一个数据库会话中运行。

每个Conn在使用后必须通过调用Conn.Close返回到数据库池中。


### func (*DB) Driver 

`func (db *DB) Driver() driver.Driver`

Driver returns the database's underlying driver.

驱动程序返回数据库的底层驱动程序。

### func (*DB) Exec

`func (db *DB) Exec(query string, args ...any) (Result, error)`

Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.

Exec uses context.Background internally; to specify the context, use ExecContext.

Exec执行一个查询，不返回任何记录。args是用于查询中的任何占位参数。

Exec内部使用context.Background；要指定上下文，请使用ExecContext。


### func (*DB) ExecContext 

`func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (Result, error)`

ExecContext executes a query without returning any rows. The args are for any placeholder parameters in the query.

ExecContext执行一个查询，不返回任何行。args是用于查询中的任何占位参数。

### func (*DB) Ping

`func (db *DB) Ping() error`

Ping verifies a connection to the database is still alive, establishing a connection if necessary.

Ping uses context.Background internally; to specify the context, use PingContext.

Ping验证与数据库的连接是否仍然存在，如果需要的话，建立一个连接。

Ping在内部使用context.Background；要指定上下文，请使用PingContext。

### func (*DB) PingContext 

`func (db *DB) PingContext(ctx context.Context) error`

PingContext verifies a connection to the database is still alive, establishing a connection if necessary.

PingContext验证与数据库的连接是否仍然有效，如果有必要，建立一个连接。

### func (*DB) Prepare

`func (db *DB) Prepare(query string) (*Stmt, error)`

Prepare creates a prepared statement for later queries or executions. Multiple queries or executions may be run concurrently from the returned statement. The caller must call the statement's Close method when the statement is no longer needed.

Prepare uses context.Background internally; to specify the context, use PrepareContext.


准备（Prepare）为以后的查询或执行创建一个准备好的语句。多个查询或执行可以在返回的语句中同时运行。当不再需要该语句时，调用者必须调用该语句的关闭方法。

Prepare在内部使用context.Background；要指定上下文，请使用PrepareContext。


### func (*DB) PrepareContext 

`func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error)`

PrepareContext creates a prepared statement for later queries or executions. Multiple queries or executions may be run concurrently from the returned statement. The caller must call the statement's Close method when the statement is no longer needed.

The provided context is used for the preparation of the statement, not for the execution of the statement.

PrepareContext为以后的查询或执行创建一个准备好的语句。多个查询或执行可以在返回的语句中同时运行。当不再需要该语句时，调用者必须调用该语句的关闭方法。

所提供的上下文用于准备语句，而不是用于执行语句。

### func (*DB) Query

`func (db *DB) Query(query string, args ...any) (*Rows, error)`

Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.

Query uses context.Background internally; to specify the context, use QueryContext.

查询执行一个返回行的查询，通常是一个SELECT。args是用于查询中的任何占位参数。

Query内部使用context.Background；要指定上下文，请使用QueryContext.Background。

### func (*DB) QueryContext

`func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)`

QueryContext executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.

QueryContext执行一个返回行的查询，通常是一个SELECT。args是用于查询中的任何占位参数。

### func (*DB) QueryRow

`func (db *DB) QueryRow(query string, args ...any) *Row`

QueryRow executes a query that is expected to return at most one row. QueryRow always returns a non-nil value. Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.

QueryRow uses context.Background internally; to specify the context, use QueryRowContext.

QueryRow执行一个查询，预计最多返回一条记录。QueryRow总是返回一个非零的值。错误被推迟到Row的Scan方法被调用。如果查询没有选择任何行，*Row的扫描将返回ErrNoRows。否则，*Row's Scan会扫描第一条被选择的行，并丢弃其余的。

QueryRow在内部使用context.Background；要指定上下文，请使用QueryRowContext。

### func (*DB) QueryRowContext 

`func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *Row`

QueryRowContext executes a query that is expected to return at most one row. QueryRowContext always returns a non-nil value. Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.

QueryRowContext执行一个查询，预计最多返回一条记录。QueryRowContext总是返回一个非零的值。错误被推迟到Row的Scan方法被调用。如果查询没有选择任何行，*Row的扫描将返回ErrNoRows。否则，*Row's Scan会扫描第一条被选择的行，并丢弃其余的行。

### func (*DB) SetConnMaxIdleTime

`func (db *DB) SetConnMaxIdleTime(d time.Duration)`

SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.

Expired connections may be closed lazily before reuse.

If d <= 0, connections are not closed due to a connection's idle time.

SetConnMaxIdleTime设置一个连接可能被闲置的最大时间。

过期的连接在重复使用前可能会被懒散地关闭。

如果d<=0，连接不会因为连接的空闲时间而被关闭。

### func (*DB) SetConnMaxLifetime 

`func (db *DB) SetConnMaxLifetime(d time.Duration)`

SetConnMaxLifetime sets the maximum amount of time a connection may be reused.

Expired connections may be closed lazily before reuse.

If d <= 0, connections are not closed due to a connection's age.

SetConnMaxLifetime设置一个连接可以被重复使用的最大时间。

过期的连接在重复使用前可能会被懒惰地关闭。

如果d<=0，连接不会因为连接的年龄而被关闭。

### func (*DB) SetMaxIdleConns

`func (db *DB) SetMaxIdleConns(n int)`

SetMaxIdleConns sets the maximum number of connections in the idle connection pool.

If MaxOpenConns is greater than 0 but less than the new MaxIdleConns, then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.

If n <= 0, no idle connections are retained.

The default max idle connections is currently 2. This may change in a future release.

SetMaxIdleConns设置空闲连接池的最大连接数。

如果MaxOpenConns大于0但小于新的MaxIdleConns，那么新的MaxIdleConns将被减少以符合MaxOpenConns的限制。

如果n<=0，则不保留空闲连接。

目前默认的最大空闲连接数是2。 这可能会在未来的版本中改变。

### func (*DB) SetMaxOpenConns

`func (db *DB) SetMaxOpenConns(n int)`

SetMaxOpenConns sets the maximum number of open connections to the database.

If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than MaxIdleConns, then MaxIdleConns will be reduced to match the new MaxOpenConns limit.

If n <= 0, then there is no limit on the number of open connections. The default is 0 (unlimited).


SetMaxOpenConns设置到数据库的最大开放连接数。

如果MaxIdleConns大于0，而新的MaxOpenConns小于MaxIdleConns，那么MaxIdleConns将被减少以匹配新的MaxOpenConns限制。

如果n<=0，那么对开放连接的数量没有限制。默认是0（无限）。


### func (*DB) Stats

`func (db *DB) Stats() DBStats`

Stats returns database statistics.

Stats返回数据库的统计数据。

### type DBStats 

```
type DBStats struct {
	MaxOpenConnections int // Maximum number of open connections to the database.

	// Pool Status
	OpenConnections int // The number of established connections both in use and idle.
	InUse           int // The number of connections currently in use.
	Idle            int // The number of idle connections.

	// Counters
	WaitCount         int64         // The total number of connections waited for.
	WaitDuration      time.Duration // The total time blocked waiting for a new connection.
	MaxIdleClosed     int64         // The total number of connections closed due to SetMaxIdleConns.
	MaxIdleTimeClosed int64         // The total number of connections closed due to SetConnMaxIdleTime.
	MaxLifetimeClosed int64         // The total number of connections closed due to SetConnMaxLifetime.
}
```

## type IsolationLevel

```
type IsolationLevel int
```
IsolationLevel is the transaction isolation level used in TxOptions.

IsolationLevel是TxOptions中使用的事务隔离级别

```
const (
	LevelDefault IsolationLevel = iota
	LevelReadUncommitted
	LevelReadCommitted
	LevelWriteCommitted
	LevelRepeatableRead
	LevelSnapshot
	LevelSerializable
	LevelLinearizable
)
```
Various isolation levels that drivers may support in BeginTx. If a driver does not support a given isolation level an error may be returned.

驱动程序在BeginTx中可能支持的各种隔离级别。如果一个驱动程序不支持给定的隔离级别，可能会返回一个错误。

### func (IsolationLevel) String

`func (i IsolationLevel) String() string`

String returns the name of the transaction isolation level.

字符串返回事务隔离级别的名称。

## type NamedArg

```
type NamedArg struct {

	// Name is the name of the parameter placeholder.
	//
	// If empty, the ordinal position in the argument list will be
	// used.
	//
	// Name must omit any symbol prefix.
	Name string

	// Value is the value of the parameter.
	// It may be assigned the same value types as the query
	// arguments.
	Value any
	// contains filtered or unexported fields
}
```

A NamedArg is a named argument. NamedArg values may be used as arguments to Query or Exec and bind to the corresponding named parameter in the SQL statement.

For a more concise way to create NamedArg values, see the Named function.

NamedArg是一个命名的参数。NamedArg值可以作为Query或Exec的参数，并与SQL语句中相应的命名参数绑定。

关于创建NamedArg值的更简洁的方法，请参见Named函数。

### func Named

`func Named(name string, value any) NamedArg`

Named provides a more concise way to create NamedArg values.

Named提供了一种更简洁的方式来创建NamedArg值。

## type NullBool

```
type NullBool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}
```

NullBool represents a bool that may be null. NullBool implements the Scanner interface so it can be used as a scan destination, similar to NullString.

NullBool表示一个可能为空的bool。NullBool实现了Scanner接口，所以它可以作为一个扫描目标，类似于NullString

### func (*NullBool) Scan

`func (n *NullBool) Scan(value any) error`

Scan implements the Scanner interface.

Scan实现了Scanner接口。

### func (NullBool) Value

`func (n NullBool) Value() (driver.Value, error)`

Value implements the driver Valuer interface.

Value实现了驱动Valuer接口。

## type NullByte

```
type NullByte struct {
	Byte  byte
	Valid bool // Valid is true if Byte is not NULL
}
```

NullByte represents a byte that may be null. NullByte implements the Scanner interface so it can be used as a scan destination, similar to NullString.

NullByte表示一个可能为空的字节。NullByte实现了Scanner接口，所以它可以作为一个扫描目标，类似于NullString。

### func (*NullByte) Scan

`func (n *NullByte) Scan(value any) error`

Scan implements the Scanner interface.

### func (NullByte) Value

`func (n NullByte) Value() (driver.Value, error)`

Value implements the driver Valuer interface.

## type Out

```
type Out struct {

	// Dest is a pointer to the value that will be set to the result of the
	// stored procedure's OUTPUT parameter.
	Dest any

	// In is whether the parameter is an INOUT parameter. If so, the input value to the stored
	// procedure is the dereferenced value of Dest's pointer, which is then replaced with
	// the output value.
	In bool
	// contains filtered or unexported fields
}
```

Out may be used to retrieve OUTPUT value parameters from stored procedures.

Not all drivers and databases support OUTPUT value parameters.

Out可以用来从存储过程中检索OUTPUT值参数。

不是所有的驱动程序和数据库都支持OUTPUT值参数。

## type RawBytes

```
type RawBytes []byte
```

RawBytes is a byte slice that holds a reference to memory owned by the database itself. After a Scan into a RawBytes, the slice is only valid until the next call to Next, Scan, or Close.

RawBytes是一个字节片，持有对数据库本身拥有的内存的引用。在对RawBytes进行扫描后，该分片只在下次调用Next、Scan或Close前有效。

## type Result

```
type Result interface {
	// LastInsertId returns the integer generated by the database
	// in response to a command. Typically this will be from an
	// "auto increment" column when inserting a new row. Not all
	// databases support this feature, and the syntax of such
	// statements varies.
	LastInsertId() (int64, error)

	// RowsAffected returns the number of rows affected by an
	// update, insert, or delete. Not every database or database
	// driver may support this.
	RowsAffected() (int64, error)
}
```
A Result summarizes an executed SQL command

一个结果总结了一个已执行的SQL命令

## type Row

```
type Row struct {
	// contains filtered or unexported fields
}
```

Row is the result of calling QueryRow to select a single row.

Row是调用QueryRow来选择一条记录的结果。

### func (*Row) Err

`func (r *Row) Err() error`

Err provides a way for wrapping packages to check for query errors without calling Scan. Err returns the error, if any, that was encountered while running the query. If this error is not nil, this error will also be returned from Scan.

Err为封装包提供了一种方法，可以在不调用Scan的情况下检查查询错误。Err 返回运行查询时遇到的错误（如果有的话）。如果这个错误不是nil，这个错误也将从Scan返回。

### func (*Row) Scan

`func (r *Row) Scan(dest ...any) error`

Scan copies the columns from the matched row into the values pointed at by dest. See the documentation on Rows.Scan for details. If more than one row matches the query, Scan uses the first row and discards the rest. If no row matches the query, Scan returns ErrNoRows.

扫描将匹配行中的列复制到dest所指向的值中。详情请参见Rows.Scan的文档。如果有多条记录与查询相匹配，Scan将使用第一条记录，并丢弃其余的记录。如果没有符合查询的行，Scan返回ErrNoRows。

## type Rows

```
type Rows struct {
	// contains filtered or unexported fields
}
```
Rows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance from row to row.

Rows是一个查询的结果。它的光标从结果集的第一行前开始。使用 "下一步 "来从一行推进到另一行。

### func (*Rows) Close

`func (rs *Rows) Close() error`

Close closes the Rows, preventing further enumeration. If Next is called and returns false and there are no further result sets, the Rows are closed automatically and it will suffice to check the result of Err. Close is idempotent and does not affect the result of Err.

Close关闭了Rows，防止进一步的枚举。如果调用Next并返回false，并且没有进一步的结果集，那么Rows将被自动关闭，检查Err的结果就足够了。Close是empotent的，不影响Err的结果。

### func (*Rows) ColumnTypes

`func (rs *Rows) ColumnTypes() ([]*ColumnType, error)`

ColumnTypes returns column information such as column type, length, and nullable. Some information may not be available from some drivers.

ColumnTypes返回列的信息，如列的类型、长度和nullable。有些信息可能无法从某些驱动程序中获得。

### func (*Rows) Columns

`func (rs *Rows) Columns() ([]string, error)`

Columns returns the column names. Columns returns an error if the rows are closed.

Columns返回列的名称。如果行被关闭，Columns返回一个错误。

### func (*Rows) Err

`func (rs *Rows) Err() error`

Err returns the error, if any, that was encountered during iteration. Err may be called after an explicit or implicit Close.

Err返回迭代过程中遇到的错误（如果有的话）。Err可以在显式或隐式关闭后调用。

### func (*Rows) Next

`func (rs *Rows) Next() bool`

Next prepares the next result row for reading with the Scan method. It returns true on success, or false if there is no next result row or an error happened while preparing it. Err should be consulted to distinguish between the two cases.

Every call to Scan, even the first one, must be preceded by a call to Next.

下一步是用扫描方法准备下一个结果行的读取。如果没有下一条结果行或者在准备过程中发生错误，则返回true，或者返回false。应该参考Err来区分这两种情况。

每一次对Scan的调用，甚至是第一次，都必须先对Next进行调用。

### func (*Rows) NextResultSet

`func (rs *Rows) NextResultSet() bool`

NextResultSet prepares the next result set for reading. It reports whether there is further result sets, or false if there is no further result set or if there is an error advancing to it. The Err method should be consulted to distinguish between the two cases.

After calling NextResultSet, the Next method should always be called before scanning. If there are further result sets they may not have rows in the result set.

NextResultSet为读取下一个结果集做准备。它报告是否有进一步的结果集，如果没有进一步的结果集或者有一个错误的推进，则报告为假。应该参考Err方法来区分这两种情况。

在调用NextResultSet之后，在扫描之前应该总是调用Next方法。如果有进一步的结果集，它们可能在结果集中没有行。

### func (*Rows) Scan 

`func (rs *Rows) Scan(dest ...any) error`

Scan copies the columns in the current row into the values pointed at by dest. The number of values in dest must be the same as the number of columns in Rows.

Scan converts columns read from the database into the following common Go types and special types provided by the sql package:

扫描将当前行中的列复制到dest所指向的值中。dest中的值的数量必须与Rows中的列的数量相同。

Scan将从数据库中读取的列转换为以下常见的Go类型和由sql包提供的特殊类型。

## type Scanner

```
type Scanner interface {
	// Scan assigns a value from a database driver.
	//
	// The src value will be of one of the following types:
	//
	//    int64
	//    float64
	//    bool
	//    []byte
	//    string
	//    time.Time
	//    nil - for NULL values
	//
	// An error should be returned if the value cannot be stored
	// without loss of information.
	//
	// Reference types such as []byte are only valid until the next call to Scan
	// and should not be retained. Their underlying memory is owned by the driver.
	// If retention is necessary, copy their values before the next call to Scan.
	Scan(src any) error
}    
```
Scanner is an interface used by Scan.

Scanner是Scan使用的一个接口。

## type Stmt 

```
type Stmt struct {
	// contains filtered or unexported fields
}
```

Stmt is a prepared statement. A Stmt is safe for concurrent use by multiple goroutines.

If a Stmt is prepared on a Tx or Conn, it will be bound to a single underlying connection forever. If the Tx or Conn closes, the Stmt will become unusable and all operations will return an error. If a Stmt is prepared on a DB, it will remain usable for the lifetime of the DB. When the Stmt needs to execute on a new underlying connection, it will prepare itself on the new connection automatically.

Stmt是一个准备好的语句。一个Stmt对于多个goroutine的并发使用是安全的。

如果一个Stmt是在Tx或Conn上准备的，它将永远被绑定到一个底层连接。如果Tx或Conn关闭，Stmt将变得不可用，所有操作将返回一个错误。如果一个Stmt是在一个DB上准备的，它将在该DB的生命周期内保持可用。当Stmt需要在一个新的底层连接上执行时，它将自动在新的连接上准备自己。

### func (*Stmt) Close

`func (s *Stmt) Close() error`

Close closes the statement.

### func (*Stmt) Exec

`func (s *Stmt) Exec(args ...any) (Result, error)`

Exec executes a prepared statement with the given arguments and returns a Result summarizing the effect of the statement.

Exec uses context.Background internally; to specify the context, use ExecContext.

Exec用给定的参数执行一个准备好的语句，并返回一个总结该语句效果的结果。

Exec内部使用context.Background；要指定上下文，请使用ExecContext.Background。

### func (*Stmt) ExecContext

`func (s *Stmt) ExecContext(ctx context.Context, args ...any) (Result, error)`

ExecContext executes a prepared statement with the given arguments and returns a Result summarizing the effect of the statement.

ExecContext用给定的参数执行一个准备好的语句，并返回一个总结该语句效果的结果。

### func (*Stmt) Query

`func (s *Stmt) Query(args ...any) (*Rows, error)`

Query executes a prepared query statement with the given arguments and returns the query results as a *Rows.

Query uses context.Background internally; to specify the context, use QueryContext.

Query用给定的参数执行一个准备好的查询语句，并以*Rows的形式返回查询结果。

Query内部使用context.Background；要指定上下文，请使用QueryContext.Background。

### func (*Stmt) QueryContext 

`func (s *Stmt) QueryContext(ctx context.Context, args ...any) (*Rows, error)`

QueryContext executes a prepared query statement with the given arguments and returns the query results as a *Rows.

QueryContext用给定的参数执行一个准备好的查询语句，并将查询结果作为一个*Rows返回。

### func (*Stmt) QueryRow

`func (s *Stmt) QueryRow(args ...any) *Row`

QueryRow executes a prepared query statement with the given arguments. If an error occurs during the execution of the statement, that error will be returned by a call to Scan on the returned *Row, which is always non-nil. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.

QueryRow用给定的参数执行一个准备好的查询语句。如果在语句的执行过程中发生错误，该错误将通过对返回的*Row的Scan的调用来返回，该错误总是非零。如果查询没有选择任何行，*Row的Scan将返回ErrNoRows。否则，*Row's Scan将扫描第一条被选择的行，并丢弃其余的行。


### func (*Stmt) QueryRowContext 

`func (s *Stmt) QueryRowContext(ctx context.Context, args ...any) *Row`

QueryRowContext executes a prepared query statement with the given arguments. If an error occurs during the execution of the statement, that error will be returned by a call to Scan on the returned *Row, which is always non-nil. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.

QueryRowContext用给定的参数执行一个准备好的查询语句。如果在语句的执行过程中发生错误，该错误将通过对返回的*Row的Scan的调用来返回，该错误总是非零。如果查询没有选择任何行，*Row的Scan将返回ErrNoRows。否则，*Row's Scan将扫描第一条被选择的行，并丢弃其余的行。

## type Tx

```
type Tx struct {
	// contains filtered or unexported fields
}
```

Tx is an in-progress database transaction.

A transaction must end with a call to Commit or Rollback.

After a call to Commit or Rollback, all operations on the transaction fail with ErrTxDone.

The statements prepared for a transaction by calling the transaction's Prepare or Stmt methods are closed by the call to Commit or Rollback.

Tx是一个正在进行的数据库事务。

一个事务必须以调用Commit或Rollback来结束。

在调用Commit或Rollback后，对事务的所有操作都以ErrTxDone失败。

通过调用事务的Prepare或Stmt方法为事务准备的语句会在调用Commit或Rollback时关闭。


### func (*Tx) Commit 

`func (tx *Tx) Commit() error`

Commit commits the transaction.

### func (*Tx) Exec 

`func (tx *Tx) Exec(query string, args ...any) (Result, error)`

Exec executes a query that doesn't return rows. For example: an INSERT and UPDATE.

Exec uses context.Background internally; to specify the context, use ExecContext.


Exec执行一个不返回记录的查询。例如：一个INSERT和UPDATE。

Exec内部使用context.Background；要指定上下文，请使用ExecContext.Background。

### func (*Tx) ExecContext 

`func (tx *Tx) ExecContext(ctx context.Context, query string, args ...any) (Result, error)`

ExecContext executes a query that doesn't return rows. For example: an INSERT and UPDATE.

ExecContext执行一个不返回记录的查询。例如：一个INSERT和UPDATE。

### func (*Tx) Prepare 

`func (tx *Tx) Prepare(query string) (*Stmt, error)`

Prepare creates a prepared statement for use within a transaction.

The returned statement operates within the transaction and will be closed when the transaction has been committed or rolled back.

To use an existing prepared statement on this transaction, see Tx.Stmt.

Prepare uses context.Background internally; to specify the context, use PrepareContext.

准备（Prepare）创建一个准备好的语句，在一个事务中使用。

返回的语句在事务中运行，当事务被提交或回滚时将被关闭。

要在这个事务中使用一个现有的准备好的语句，请看Tx.Stmt。

Prepare在内部使用context.Background；要指定上下文，请使用PrepareContext。

### func (*Tx) PrepareContext

`func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error)`

PrepareContext creates a prepared statement for use within a transaction.

The returned statement operates within the transaction and will be closed when the transaction has been committed or rolled back.

To use an existing prepared statement on this transaction, see Tx.Stmt.

The provided context will be used for the preparation of the context, not for the execution of the returned statement. The returned statement will run in the transaction context.


PrepareContext创建一个准备好的语句，以便在一个事务中使用。

返回的语句在事务中运行，当事务被提交或回滚时将被关闭。

要在这个事务中使用一个现有的准备好的语句，请参见Tx.Stmt。

提供的上下文将被用于准备上下文，而不是执行返回的语句。返回的语句将在事务上下文中运行。

### func (*Tx) Query

`func (tx *Tx) Query(query string, args ...any) (*Rows, error)`

Query executes a query that returns rows, typically a SELECT.

Query uses context.Background internally; to specify the context, use QueryContext.

Query执行一个返回行的查询，通常是一个SELECT。

Query在内部使用context.Background；要指定上下文，请使用QueryContext。

### func (*Tx) QueryContext 

`func (tx *Tx) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)`

QueryContext executes a query that returns rows, typically a SELECT.

### func (*Tx) QueryRow

`func (tx *Tx) QueryRow(query string, args ...any) *Row`

QueryRow executes a query that is expected to return at most one row. QueryRow always returns a non-nil value. Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.

QueryRow uses context.Background internally; to specify the context, use QueryRowContext.

QueryRow执行一个查询，预计最多返回一条记录。QueryRow总是返回一个非零的值。错误被推迟到Row的Scan方法被调用。如果查询没有选择任何行，*Row的扫描将返回ErrNoRows。否则，*Row's Scan会扫描第一条被选择的行，并丢弃其余的。

QueryRow在内部使用context.Background；要指定上下文，请使用QueryRowContext。

### func (*Tx) QueryRowContext

`func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...any) *Row`

QueryRowContext executes a query that is expected to return at most one row. QueryRowContext always returns a non-nil value. Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.

### func (*Tx) Rollback 

`func (tx *Tx) Rollback() error`

Rollback aborts the transaction.

### func (*Tx) Stmt

`func (tx *Tx) Stmt(stmt *Stmt) *Stmt`

Stmt returns a transaction-specific prepared statement from an existing statement.

Stmt从一个现有的语句中返回一个特定于事务的准备好的语句。

### func (*Tx) StmtContext

`func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt`

## type TxOptions

```
type TxOptions struct {
	// Isolation is the transaction isolation level.
	// If zero, the driver or database's default level is used.
	Isolation IsolationLevel
	ReadOnly  bool
}
```

TxOptions holds the transaction options to be used in DB.BeginTx.

TxOptions持有将在DB.BeginTx中使用的交易选项。