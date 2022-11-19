package database

import (
	"bytes"
	"ferry/global/orm"
	"ferry/pkg/logger"
	"ferry/tools/config"
	"strconv"

	"github.com/spf13/viper"

	"gorm.io/driver/mysql" //加载mysql
	"gorm.io/gorm"
)

var (
	DbType   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
)

func (e *Mysql) Setup() {

	var err error
	var db Database

	db = new(Mysql)
	orm.MysqlConn = db.GetConnect()
	orm.Eloquent, err = db.Open(orm.MysqlConn)

	if err != nil {
		logger.Fatalf("%s connect error %v", orm.MysqlConn, err)
	} else {
		logger.Infof("%s connect success!", orm.MysqlConn)
	}

	if orm.Eloquent.Error != nil {
		logger.Fatalf("database error %v", orm.Eloquent.Error)
	}

	// 是否开启详细日志记录
	// orm.Eloquent.LogMode(viper.GetBool("settings.gorm.logMode"))
	sqlDB, err := orm.Eloquent.DB()
	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(viper.GetInt("settings.gorm.maxOpenConn"))

	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用
	sqlDB.SetMaxIdleConns(viper.GetInt("settings.gorm.maxIdleConn"))
}

type Mysql struct {
}

func (e *Mysql) Open(conn string) (db *gorm.DB, err error) {
	return gorm.Open(mysql.Open(conn), &gorm.Config{})
}

func (e *Mysql) GetConnect() string {

	DbType = config.DatabaseConfig.Dbtype
	Host = config.DatabaseConfig.Host
	Port = config.DatabaseConfig.Port
	Name = config.DatabaseConfig.Name
	Username = config.DatabaseConfig.Username
	Password = config.DatabaseConfig.Password

	var conn bytes.Buffer
	conn.WriteString(Username)
	conn.WriteString(":")
	conn.WriteString(Password)
	conn.WriteString("@tcp(")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(Name)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=10000ms")
	return conn.String()
}
