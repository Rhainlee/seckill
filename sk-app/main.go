package main

import (
	"github.com/rhainlee/seckill/pkg/bootstrap"
	"github.com/rhainlee/seckill/sk-app/setup"
)

// 秒杀业务系统主要为前端/移动端提供秒杀活动查询和进行秒杀的HTTP接口，处理有关用户ID和IP
// 黑白名单 和进行流量限制的逻辑，并通过Redis将合法的秒杀请求发送给秒杀核心业务，
// 并将秒杀核心业务的处理结果返回给前端/移动端
// 秒杀业务系统和秒杀核心系统之间通过Redis的队列进行交互

// 从 Zookeeper 中加载秒杀活动数据到内存中，监听Zookeeper中的数据变化,
// 并实时更新数据到内存中.建立Redis连接，启动工作协程.
func main() {
	//mysql.InitMysql(conf.MysqlConf.Host, conf.MysqlConf.Port, conf.MysqlConf.User, conf.MysqlConf.Pwd, conf.MysqlConf.Db)
	setup.InitZk()
	setup.InitRedis()
	setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)
}
