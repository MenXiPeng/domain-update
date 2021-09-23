package main

import (
	"log"
	"org.tubetrue01/domain-update/config"
	"org.tubetrue01/domain-update/notify"
	"org.tubetrue01/domain-update/util"
	"time"
)

var command *config.Config
var doNotify notify.Notify

func main() {
	task(command)
}

// task 执行定时任务，根据需要定期更新域名解析
func task(config *config.Config) {
	for {
		log.Println("定时任务开始...")
		pubIp := util.ObtainPubIp()
		if match, _ := util.IsValidIp(pubIp); match {
			if ip, exists := util.ObtainIpFromPool(); exists {
				log.Printf("ip 已存在，当前值为：%s, 当前本机公网 ip 为：%s\n", ip, pubIp)
				if ip != pubIp {
					log.Printf("ip 地址已经发生变化，准备进行推送...")
					doNotify.DoNotify(config, pubIp)
					util.UpdateIpPool(pubIp)
				}
				log.Println("开始睡眠 1 小时...")
				time.Sleep(time.Second * time.Duration(1*60*60))
			} else {
				log.Printf("ip 并不存在， 更新 ip 池...")
				util.UpdateIpPool(pubIp)
				doNotify.DoNotifyBefore(config, pubIp)
			}
		} else {
			log.Printf("获取到的公网 ip 地址：[%s] 非法，准备下次任务...\n", pubIp)
		}
	}
}

func init() {
	command = config.ObtainCommand()
	if command.IsEmail {
		doNotify = &notify.Mail{}
	} else {
		doNotify = &notify.Domain{}
	}

}
