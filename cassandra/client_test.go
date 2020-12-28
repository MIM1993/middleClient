package middleClient

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestNewCassandraCluster(t *testing.T) {
	var address = []string{"127.0.0.1:9042"}
	cd := NewCassandraCluster(address,"example","any",3,2,3)
	log.Info("cassandra init ok")
	session,err:= cd.CreateSession()
	if err!=nil{
		panic(err)
	}

	session.Close()

	//session.Query()
}