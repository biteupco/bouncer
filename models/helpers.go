package models

import (
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/mgo.v2"
)

type mongoConfiguration struct {
	Hosts    string
	Database string
	UserName string
	Password string
}

type mongoSession struct {
	dialInfo *mgo.DialInfo
	session  *mgo.Session
}

var singleton mongoSession

const (
	dbEnvPrefix = "benri"
)

func Start() error {
	if singleton.dialInfo != nil {
		return nil
	}
	var cfg mongoConfiguration
	// read from env var
	if err := envconfig.Process(dbEnvPrefix, &cfg); err != nil {
		return err
	}

	hosts := strings.Split(cfg.Hosts, ",")

	// setup with singleton
	singleton := mongoSession{
		dialInfo: &mgo.DialInfo{
			Addrs:    hosts,
			Timeout:  60 * time.Second,
			Database: cfg.Database,
			Username: cfg.UserName,
			Password: cfg.Password,
		},
	}

	mongoSession, err := mgo.DialWithInfo(singleton.dialInfo)
	if err != nil {
		return err
	}
	singleton.session = mongoSession
	singleton.session.SetSafe(&mgo.Safe{})
	return nil
}

func Close() {
	singleton.session.Close()
}
