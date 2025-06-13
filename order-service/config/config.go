package config

import (
	"order-service/common/util"
	"os"

	"github.com/sirupsen/logrus"
	_ "github.com/spf13/viper/remote"
)

var Config AppConfig //variabel config dari struct AppConfig


//struct AppConfig (mengambil dari file config.json)
type AppConfig struct {
	Port                  int      `json:"port"`
	AppName               string   `json:"appName"`
	AppEnv                string   `json:"appEnv"`
	SignatureKey          string   `json:"signatureKey"`
	Database              Database `json:"database"`
	RateLimiterMaxRequest float64  `json:"rateLimiterMaxRequest"`
	RateLimiterTimeSecond int      `json:"rateLimiterTimeSecond"`
	InternalService            InternalService `json:"internalService"`
	GCSType                    string          `json:"gcsType"`
	GCSProjectID               string          `json:"gcsProjectID"`
	GCSPrivateKeyID            string          `json:"gcsPrivateKeyID"`
	GCSPrivateKey              string          `json:"gcsPrivateKey"`
	GCSClientEmail             string          `json:"gcsClientEmail"`
	GCSClientID                string          `json:"gcsClientID"`
	GCSAuthURI                 string          `json:"gcsAuthURI"`
	GCSTokenURI                string          `json:"gcsTokenURI"`
	GCSAuthProviderX509CertURL string          `json:"gcsAuthProviderX509CertURL"`
	GCSClientX509CertURL       string          `json:"gcsClientX509CertURL"`
	GCSUniverseDomain          string          `json:"gcsUniverseDomain"`
	GCSBucketName              string          `json:"gcsBucketName"`
	Kafka	Kafka 	`json:"kafka"`
	Midtrans Midtrans `json:"midtrans"`
}

//stuct user-service
type InternalService struct {
	User User `json:"user"`
	Field User `json:"field"`
	Payment User `json:"payment"`
}

//struct dari user user-service
type User struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}
type Field struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}
type Payment struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}

//struct Database (mengambil dari file config.json)
type Database struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnections    int    `json:"maxOpenConnections"`
	MaxLifeTimeConnection int    `json:"maxLifeTimeConnection"`
	MaxIdleConnections    int    `json:"maxIdleConnections"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}
//config kafka
type Kafka struct {
	Brokers     []string `json:"brokers"`
	TimeoutInMS int      `json:"timeoutInMS"`
	MaxRetry    int      `json:"maxRetry"`
	Topic       string   `json:"topic"`
}

//config midtrans
type Midtrans struct {
	ServerKey    string `json:"serverKey"`
	ClientKey    string `json:"clientKey"`
	IsProduction bool   `json:"isProduction"`
}

//Fungsi untuk memngambil config dari config.json jika berada dilokal
//Namun akan menambil dair .env jika pada development atau production di server
func Init() {
	err := util.BindFromJSON(&Config, "config.json", ".")
	if err != nil {
		logrus.Infof("failed to bind config: %v", err)
		err = util.BindFromConsul(&Config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_PATH"))
		if err != nil {
			panic(err)
		}
	}
}