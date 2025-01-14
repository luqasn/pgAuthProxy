package proxy

import (
	"github.com/luqasn/pgAuthProxy/auth"
	"github.com/luqasn/pgAuthProxy/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net"
)

// Start PostgreSQL authentication proxy server
func Start() {
	//log.SetReportCaller(true)
	log.Info("Starting auth pgAuthProxy...")
	log.SetLevel(log.DebugLevel)

	var bindAddr = viper.GetString(utils.ConfigListenFlag)
	server, _ := net.Listen("tcp", bindAddr)
	log.WithField("address", bindAddr).Info("Started listening")
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Debug("Connection initialization error: " + err.Error())
		} else {
			go func() {
				defer conn.Close()
				front := NewProxyFront(conn, auth.Exec)
				defer front.Close()
				front.Run()
			}()
		}
	}
}
