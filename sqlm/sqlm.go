package sqlm

import (
	"database/sql"
	"reflect"
	"strings"
	"sync"

	"github.com/iteny/hmgo/sqlm/reflectm"
	"github.com/jmoiron/sqlx/reflectx"
)

// Although the NameMapper is convenient, in practice it should not
// be relied on except for application code.  If you are writing a library
// that uses sqlm, you should be aware that the name mappings you expect
// can be overridden by your user's application.
// 虽然名称映射器很方便，但他实际上不应该依赖应用程序代码。如果你写一个库使用sqlm,你应该知道你期望的名称映射器可以被用户的应用程序覆盖。
// NameMapper is used to map column names to struct field names.  By default,
// it uses strings.ToLower to lowercase struct field names.  It can be set
// to whatever you want, but it is encouraged to be set before sqlx is used
// as name-to-field mappings are cached after first use on a type.
// 名称映射器是用于映射字段名到结构体字段名。默认使用strings.ToLower来缩小结构体字段名。它能设置任何你想要的，但是鼓励在使用sqlm前进行设置作为名称到字段的映射在首次使用类型之后被缓存。
var NameMapper = strings.ToLower
var origMapper = reflect.ValueOf(NameMapper)

// DB is a wrapper around sql.DB which keeps track of the driverName upon Open,
// used mostly to automatically bind named queries using the right bindvars.
// DB结构体是一个围绕标准库sql.DB的包装器，它在Open上监控驱动名，主要用于使用正确的bindvars自动绑定命名查询。
type DB struct {
	*sql.DB
	driverName string
	unsafe     bool
	Mapper     *reflectm.Mapper
}

// Tx is an sqlx wrapper around sql.Tx with extra functionality
// Tm
type Tx struct {
	*sql.Tx
	driverName string
	unsafe     bool
	Mapper     *reflectm.Mapper
}

// Stmt is an sqlx wrapper around sql.Stmt with extra functionality
type Stmt struct {
	*sql.Stmt
	unsafe bool
	Mapper *reflectx.Mapper
}

// qStmt is an unexposed wrapper which lets you use a Stmt as a Queryer & Execer by
// implementing those interfaces and ignoring the `query` argument.
// qStmt是
type qStmt struct{ *Stmt }

// Rather than creating on init, this is created when necessary so that
// importers have time to customize the NameMapper.
// 在init上创建，这是必要时创建的，以便导入有空闲时自定义的名称映射器。
var mpr *reflectm.Mapper

// mprMu protects mpr.
// mprMu保护mpr。
var mprMu sync.Mutex

// Execer is an interface used by MustExec and LoadFile
// Execer 用于MustExec和LoadFile是一个接口
type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// mapper returns a valid mapper using the configured NameMapper func.
// mapper使用配置名称映射函数返回一个有效的映射器。
func mapper() *reflectm.Mapper {
	mprMu.Lock()
	defer mprMu.Unlock()

	if mpr == nil {
		mpr = reflectm.NewMapperFunc("db", NameMapper)
	} else if origMapper != reflect.ValueOf(NameMapper) {
		// if NameMapper has changed, create a new mapper
		mpr = reflectm.NewMapperFunc("db", NameMapper)
		origMapper = reflect.ValueOf(NameMapper)
	}
	return mpr
}

//NewEngine: Connect to a database and verify with a ping.
//NewEngine: 连接到一个数据库并且验证ping
func NewEngine(driverName, dataSourceName string) (*DB, error) {
	db, err := Open(driverName, dataSourceName)
	if err != nil {
		return db, err
	}
	// Ping verifies a connection to the database is still alive,
	// establishing a connection if necessary.
	// Ping验证链接到数据库是否还活着，必要时建立连接。
	err = db.Ping()
	return db, err
}

// Open is the same as sql.Open, but returns an *sqlx.DB instead.
//打开一个相同的sql.Open,但是返回一个*sqlm.DB代替
func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{DB: db, driverName: driverName, Mapper: mapper()}, err
}

// MustExec execs the query using e and panics if there was an error.
// Any placeholder parameters are replaced with supplied args.
// 如果出现错误，MustExec使用e和panic执行查询。任何占位符参数替换为提供的参数。
func MustExec(e Execer, query string, args ...interface{}) sql.Result {
	res, err := e.Exec(query, args...)
	if err != nil {
		panic(err)
	}
	return res
}

// MustBegin starts a transaction, and panics on error.  Returns an *sqlm.Tx instead
// of an *sql.Tx.
// MustBegin启动一个事物，并发生错误的panics。返回*sqlm.Tx取代*sql.Tx。
func (db *DB) MustBegin() *Tx {
	tx, err := db.Beginx()
	if err != nil {
		panic(err)
	}
	return tx
}

// Beginx begins a transaction and returns an *sqlx.Tx instead of an *sql.Tx.
// Beginx开始一个事物，并返回一个*sqlm.Tx取代*sql.Tx。
func (db *DB) Beginx() (*Tx, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{Tx: tx, driverName: db.driverName, unsafe: db.unsafe, Mapper: db.Mapper}, err
}

// MustExec (panic) runs MustExec using this database.
// Any placeholder parameters are replaced with supplied args.
// MustExec 运行MustExec使用此数据库。任何占位符参数都替换为提供的参数。
func (db *DB) MustExec(query string, args ...interface{}) sql.Result {
	return MustExec(db, query, args...)
}

// MustExec runs MustExec within a transaction.
// Any placeholder parameters are replaced with supplied args.
// MustExec在事物中运行MustExec。任何占位符参数都替换为提供的参数。
func (tx *Tx) MustExec(query string, args ...interface{}) sql.Result {
	return MustExec(tx, query, args...)
}
func (q *qStmt) Exec(query string, args ...interface{}) (sql.Result, error) {
	return q.Stmt.Exec(args...)
}
