package mysql

import (
	"fmt"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

// 如果需要使用事务相关，则在所有数据库的操作前使用StartTrans方法初始化ormer数据库操作对象。
type Model struct {
	ormer orm.Ormer `orm:"-" description:"beego操作数据库所需对象 用于处理事务相关操作"`
}

/* 根据配置文件初始化 */
func (m *Model) InitConf() (e error) {
	jdbc := beego.AppConfig.String("mysql_jdbc")                     // sso mysql 数据库连接
	maxIdleConns, err := beego.AppConfig.Int("mysql_max_idle_conns") // sso mysql 最大空闲连接
	if err != nil {
		maxIdleConns = 0
	}
	maxOpenConns, err := beego.AppConfig.Int("mysql_max_open_conns") // sso mysql 最大打开连接
	if err != nil {
		maxOpenConns = 30
	}

	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err!=nil {
		return err
	}
	
	for {	
		e = orm.RegisterDataBase("default", "mysql", jdbc, maxIdleConns, maxOpenConns);
		if e == nil {
			break
		}
		fmt.Println(err)
	}

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
	return nil
}

/**
	自定义初始化
 */
func (m *Model) CustomInit(jdbc string, dbName string) error {
	maxIdleConns, err := beego.AppConfig.Int("mysql_max_idle_conns") // sso mysql 最大空闲连接
	if err != nil {
		maxIdleConns = 0
	}
	maxOpenConns, err := beego.AppConfig.Int("mysql_max_open_conns") // sso mysql 最大打开连接
	if err != nil {
		maxOpenConns = 30
	}
	if dbName == "" {
		dbName = "default"
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(dbName, "mysql", jdbc, maxIdleConns, maxOpenConns)

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
	return nil
}

func (m *Model) QueryTable(obj interface{}, DbName string) (qs orm.QuerySeter) {
	return m.GetOrmer(DbName).QueryTable(obj)
}

//得到ormer对象，用于操作数据库
func (m *Model) GetOrmer(dbName string) orm.Ormer {
	if dbName == "" {
		dbName = "default"
	}

	// 不为空时 是需要使用事务操作的
	o := m.ormer

	// 为空是不需要使用事务操作的
	if o == nil {
		o = orm.NewOrm()
	}

	o.Using(dbName)
	return o
}

//开启事务
func (m *Model) StartTrans() error {
	m.ormer = orm.NewOrm()
	return m.ormer.Begin();
}

//事务回滚
func (m *Model) RollBack() (err error) {
	if err = m.ormer.Rollback(); err != nil {
		return err
	}
	m.ormer = nil
	return nil
}

//事务提交
func (m *Model) Commit() (err error) {
	err = m.Commit()
	m.ormer = nil
	return err
}
