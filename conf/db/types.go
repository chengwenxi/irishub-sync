package db

import (
	"os"
	"github.com/irisnet/irishub-sync/util/constant"
	"github.com/irisnet/irishub-sync/module/logger"
)

var (
	Host     = "127.0.0.1"
	Port     = "27217"
	User     = "user"
	Passwd   = "passwd"
	Database = "sync_irishub"
)

// get value of env var
func init() {
	host, found := os.LookupEnv(constant.EnvNameDbHost)
	if found {
		Host = host
		logger.Info.Printf("The value of env var %v is %v\n",
			constant.EnvNameDbHost, host)
	}
	
	port, found := os.LookupEnv(constant.EnvNameDbPort)
	if found {
		Port = port
		logger.Info.Printf("The value of env var %v is %v\n",
			constant.EnvNameDbPort, port)
	}
	
	user, found := os.LookupEnv(constant.EnvNameDbUser)
	if found {
		User = user
		logger.Info.Printf("The value of env var %v is %v\n",
			constant.EnvNameDbUser, user)
	}
	
	passwd, found := os.LookupEnv(constant.EnvNameDbPassWd)
	if found {
		Passwd = passwd
		logger.Info.Printf("The value of env var %v is %v\n",
			constant.EnvNameDbPassWd, passwd)
	}
	
	database, found := os.LookupEnv(constant.EnvNameDbDataBase)
	if found {
		Database = database
		logger.Info.Printf("The value of env var %v is %v\n",
			constant.EnvNameDbDataBase, database)
	}
}