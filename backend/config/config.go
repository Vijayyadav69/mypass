package config

import (
	"log"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

type config struct {
	PGSQL_USERNAME       string
	PGSQL_PASSWORD       string
	PGSQL_HOSTNAME       string
	PGSQL_HOSTPORT       int
	PGSQL_DATABASE       string
	APPLICATION_HOSTNAME string
	APPLICATION_HOSTPORT int
}

var GLOBAL_CONFIG config

func LoadConfig() {
	GLOBAL_CONFIG = config{}
	err := gotenv.Load("./config.cfg")
	if err != nil {
		log.Fatalln("error in loading config.cfg ", err.Error())
	}

	GLOBAL_CONFIG.PGSQL_USERNAME = os.Getenv("PGSQL_USERNAME")
	GLOBAL_CONFIG.PGSQL_PASSWORD = os.Getenv("PGSQL_PASSWORD")
	GLOBAL_CONFIG.PGSQL_HOSTNAME = os.Getenv("PGSQL_HOSTNAME")
	GLOBAL_CONFIG.PGSQL_DATABASE = os.Getenv("PGSQL_DATABASE")
	GLOBAL_CONFIG.APPLICATION_HOSTNAME = os.Getenv("APPLICATION_HOSTNAME")

	if GLOBAL_CONFIG.PGSQL_HOSTPORT, err = strconv.Atoi(os.Getenv("PGSQL_HOSTPORT")); err != nil {
		log.Fatalln("error in loading PGSQL_HOSTPORT from config.cfg ", err.Error())
	}
	if GLOBAL_CONFIG.APPLICATION_HOSTPORT, err = strconv.Atoi(os.Getenv("APPLICATION_HOSTPORT")); err != nil {
		log.Fatalln("error in loading APPLICATION_HOSTPORT from config.cfg ", err.Error())
	}

}
