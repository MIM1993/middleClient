package middleClient

import (
	"github.com/gocql/gocql"
	"time"
)

type Cassandra struct {
	Session *gocql.Session
}

func NewCassandraCluster(address []string, keySpace string, Consistency string, Timeout, NumConns, ProtoVer int) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(address...)
	cluster.Keyspace = keySpace
	switch Consistency {
	case "any":
		cluster.Consistency = gocql.Any
	case "one":
		cluster.Consistency = gocql.One
	case "two":
		cluster.Consistency = gocql.Two
	case "three":
		cluster.Consistency = gocql.Three
	case "quorum":
		cluster.Consistency = gocql.Quorum
	case "all":
		cluster.Consistency = gocql.All
	case "local_quorum":
		cluster.Consistency = gocql.LocalQuorum
	case "each_quorum":
		cluster.Consistency = gocql.EachQuorum
	case "local_one":
		cluster.Consistency = gocql.LocalOne
	}
	cluster.Timeout = time.Duration(Timeout) * time.Second
	//每台主机的连接数(默认值:2)
	cluster.NumConns = NumConns
	cluster.ProtoVersion = ProtoVer
	cluster.PageSize = 0
	cluster.SocketKeepalive = 10 * time.Minute
	return cluster
}
