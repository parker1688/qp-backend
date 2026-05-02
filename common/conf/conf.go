package conf

import (
	"bootpkg/common/tool"

	"github.com/spf13/viper"
)

/*
 CONFIG配置
*/

func NewOnceConfig(path string) *Config {
	v := viper.New()
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var config *Config
	if err = v.Unmarshal(&config); err != nil {
		panic(err)
	}
	if config.SHA256Salt == "" {
		config.SHA256Salt = config.General.SHA256Salt
	}
	if config.Mysql.EncryptionDSN {
		config.Mysql.DSN = tool.DecryptAESPrefixRandKeySalt(config.Mysql.DSN, "DSN")
	}
	if config.MysqlSharding.EncryptionDSN {
		config.MysqlSharding.DSN = tool.DecryptAESPrefixRandKeySalt(config.MysqlSharding.DSN, "DSN")
	}
	if config.Redis.EncryptionPassword {
		config.Redis.Password = tool.DecryptAESPrefixRandKeySalt(config.Redis.Password, "Password")
	}

	return config
}

type Config struct {
	General       general        `yaml:"General"`
	Mysql         mysql          `yaml:"Mysql"`
	MysqlSharding mysql          `yaml:"MysqlSharding"`
	Redis         redis          `yaml:"Redis"`
	SHA256Salt    string         `yaml:"SHA256Salt"`
	Mq            MQSetting      `yaml:"MQ"`
	ES            ESSetting      `yaml:"ES"`
	Log           Log            `yaml:"Log"`
	Venue         Venue          `yaml:"Venue"`
	Chatroom      Chatroom       `yaml:"Chatroom"`
	Payment       PaymentSetting `yaml:"Payment"`
	Email         struct {
		Aws struct {
			Region          string `yaml:"Region"`
			Sender          string `yaml:"Sender"`
			AccessKeyID     string `yaml:"AccessKeyID"`
			SecretAccessKey string `yaml:"SecretAccessKey"`
		} `yaml:"Aws"`
	} `yaml:"Email"`
	NATS struct {
		IsInit bool   `yaml:"IsInit"`
		URL    string `yaml:"URL"`
	} `yaml:"NATS"`
	Fingerprint Fingerprint `yaml:"Fingerprint"`
	Sms         Sms         `yaml:"Sms"`
}

type Fingerprint struct {
	IsOpen     bool   `yaml:"IsOpen" json:"IsOpen"`
	SecretKeys string `yaml:"SecretKeys" json:"SecretKeys"`
	ApiKey     string `yaml:"ApiKey" json:"ApiKey"`
}

type general struct {
	ENV                 string `yaml:"ENV"`
	AesEcbKey           string `yaml:"AesEcbKey"`
	AesUserKey          string `yaml:"AesUserKey"`
	HttpProxy           string `yaml:"HttpProxy"`
	Issuer              string `yaml:"Issuer"`
	UploadFilePath      string `yaml:"UploadFilePath"`
	UploadSecondPass    bool   `yaml:"UploadSecondPass"`
	ApiSHA256Salt       string `yaml:"ApiSHA256Salt"`
	CryptoAuthToken     string `yaml:"CryptoAuthToken"`
	RsaPublicKey        string `yaml:"RsaPublicKey"`
	RsaPrivateKey       string `yaml:"RsaPrivateKey"`
	DefaultCurrency     string `yaml:"DefaultCurrency"`
	ImageDomain         string `yaml:"ImageDomain"`
	TestEnvLaunchApiUrl string `yaml:"TestEnvLaunchApiUrl"`
	GoogleAuthFlag      bool   `yaml:"GoogleAuthFlag"`
	Ipdb                string `yaml:"Ipdb"`
	SHA256Salt          string `yaml:"SHA256Salt"`
	BetRecordUrl        string `yaml:"BetRecordUrl"`
}

type mysql struct {
	IsInit                    bool   `yaml:"IsInit"`
	DSN                       string `yaml:"DSN"`
	MaxOpenConn               int    `yaml:"MaxOpenConn"`
	MaxIdleConn               int    `yaml:"MaxIdleConn"`
	EncryptionDSN             bool   `yaml:"EncryptionDSN"`
	SkipInitializeWithVersion bool   `yaml:"SkipInitializeWithVersion"`
}

type redis struct {
	IsInit             bool   `yaml:"IsInit"`
	Addr               string `yaml:"Addr"`
	Password           string `yaml:"Password"`
	Db                 int    `yaml:"Db"`
	PoolSize           int    `yaml:"PoolSize"`
	EncryptionPassword bool   `yaml:"EncryptionPassword"`
	MinIdleConn        int    `yaml:"MinIdleConn"`
}

type MQSetting struct {
	IsInit bool         `yaml:"IsInit"`
	Kafka  KafkaSetting `yaml:"Kafka"`
}

type ESSetting struct {
	IsInit bool   `yaml:"IsInit"`
	Addr   string `yaml:"Addr"`
}

type KafkaSetting struct {
	Addr    string `yaml:"Addr"`
	Version string `yaml:"Version"`
	IsAsync bool   `yaml:"IsAsync"`
}

type Log struct {
	Prefix  string `yaml:"prefix"`
	LogFile bool   `yaml:"log-file"`
	Stdout  string `yaml:"stdout"`
	File    string `yaml:"file"`
	Path    string `yaml:"path"`
}

type Sms struct {
	AccessKeyId     string `yaml:"AccessKeyId"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	Endpoint        string `yaml:"Endpoint"`
	SignName        string `yaml:"SignName"`
	TemplateCode    string `yaml:"templateCode"`
}

type Venue struct {
	IsInit bool  `yaml:"IsInit"`
	Fbty   Fbty  `yaml:"Fbty"`
	HGTY   HGTY  `yaml:"HGTY"`
	PPDZ   PPDZ  `yaml:"PPDZ"`
	PTDZ   PTDZ  `yaml:"PTDZ"`
	PGDZ   PGDZ  `yaml:"PGDZ"`
	WUGDZ  WUGDZ `yaml:"WUGDZ"`
	FGDZ   FGDZ  `yaml:"FGDZ"`
	CQ9    CQ9   `yaml:"CQ9"`
	JDB    JDB   `yaml:"JDB"`
	AGZR   AGZR  `yaml:"AGZR"`
	BGZR   BGZR  `yaml:"BGZR"`
	BBIN   BBIN  `yaml:"BBIN"`
	DBDJ   DBDJ  `yaml:"DBDJ"`
	DBZR   DBZR  `yaml:"DBZR"`
	KYQP   KYQP  `yaml:"KYQP"`
	TYQP   TYQP  `yaml:"TYQP"`
	LYQP   LYQP  `yaml:"LYQP"`
	VGQP   VGQP  `yaml:"VGQP"`
	WALI   WALI  `yaml:"WALI"`
	MTQP   MTQP  `yaml:"MTQP"`
	MGDZ   MGDZ  `yaml:"MGDZ"`
	VRCP   VRCP  `yaml:"VRCP"`
	JLDZ   JLDZ  `yaml:"JLDZ"`
	KXDZ   KXDZ  `yaml:"KXDZ"`
	LGDDZ  LGDDZ `yaml:"LGDDZ"`
	MWDZ   MWDZ  `yaml:"MWDZ"`
	SGDZ   SGDZ  `yaml:"SGDZ"`
	TTQP   TTQP  `yaml:"TTQP"`
	IMTY   IMTY  `yaml:"IMTY"`
	SBTY   SBTY  `yaml:"SBTY"`
	PDTY   PDTY  `yaml:"PDTY"`
}

type HGTY struct {
	Url      string `yaml:"url"`
	Aeskey   string `yaml:"aeskey"`
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
	AgId     string `yaml:"agid"`
}

type LYQP struct {
	Url    string `yaml:"url"`
	Agent  string `yaml:"agent"`
	Md5key string `yaml:"md5key"`
	AesKey string `yaml:"aesKey"`
}

type VGQP struct {
	Url      string `yaml:"url"`
	Channel  string `yaml:"channel"`
	Agent    string `yaml:"agent"`
	Password string `yaml:"password"`
}

type FGDZ struct {
	Url          string `yaml:"url"`
	Merchantname string `yaml:"merchantname"`
	Merchantcode string `yaml:"merchantcode"`
}

type DBDJ struct {
	Url      string `yaml:"url"`
	Merchant string `yaml:"merchant"`
	Md5key   string `yaml:"md5key"`
}

type Fbty struct {
	Url          string `yaml:"url"`
	MerchantCode string `yaml:"merchantCode"`
	MerchantId   string `yaml:"merchantId"`
	Secret       string `yaml:"secret"`
}

type PPDZ struct {
	Url             string `yaml:"url"`
	SecureLogin     string `yaml:"secureLogin"`
	SecretKey       string `yaml:"secretKey"`
	DemoGamesDomain string `yaml:"demoGamesDomain"`
}

type WUGDZ struct {
	Url          string `yaml:"url"`
	HostId       string `yaml:"hostId"`
	HostToken    string `yaml:"hostToken"`
	ReturnUrl    string `yaml:"returnUrl"`
	ReturnTarget string `yaml:"returnTarget"`
}

type PGDZ struct {
	Url       string `yaml:"url"`
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

type VRCP struct {
	Url    string `yaml:"url"`
	AppID  string `yaml:"appId"`
	AesKey string `yaml:"aesKey"`
}

type DBZR struct {
	Url       string `yaml:"url"`
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

type WALI struct {
	Url       string `yaml:"url"`
	Acount    string `yaml:"acount"`
	AesKey    string `yaml:"aesKey"`
	SignKey   string `yaml:"signKey"`
	AgentName string `yaml:"agentName"`
}

type PTDZ struct {
	Url       string `yaml:"url"`
	AdminName string `yaml:"adminName"`
	AppSecret string `yaml:"appSecret"`
}

type CQ9 struct {
	Url      string `yaml:"url"`
	TOKEN    string `yaml:"token"`
	Gamehall string `yaml:"gamehall"`
}

type Chatroom struct {
	WebsocketAddr string `yaml:"WebsocketAddr"`
}

type JDB struct {
	Url     string `yaml:"url"`
	DCName  string `yaml:"DCName"`
	AgentId string `yaml:"agentId"`
	IV      string `yaml:"iv"`
	Key     string `yaml:"key"`
}

type AGZR struct {
	Url     string `yaml:"url"`
	GameUrl string `yaml:"gameurl"`
	Agent   string `yaml:"agent"`
	Md5Key  string `yaml:"md5Key"`
	DesKey  string `yaml:"desKey"`
}

type BGZR struct {
	Url       string `yaml:"url"`
	Sn        string `yaml:"sn"`
	LoginId   string `yaml:"loginId"`
	Password  string `yaml:"password"`
	SecretKey string `yaml:"secretKey"`
}

type BBIN struct {
	Url     string `yaml:"url"`
	LoginId string `yaml:"loginId"`
	Md5Key  string `yaml:"md5Key"`
	prefix  string `yaml:"prefix"`
}

type KYQP struct {
	Url     string `yaml:"url" json:"url"`
	AgentId string `yaml:"AgentId" json:"AgentId"`
	Md5Key  string `yaml:"Md5Key" json:"Md5Key"`
	DesKey  string `yaml:"DesKey" json:"DesKey"`
}

type TYQP struct {
	Url       string `yaml:"url" json:"url"`
	Userid    string `yaml:"userid" json:"userid"`
	Md5Key    string `yaml:"md5Key" json:"md5key"`
	DesKey    string `yaml:"deskey" json:"deskey"`
	Apisuffix string `yaml:"apisuffix" json:"apisuffix"`
}

type MTQP struct {
	Url          string `yaml:"url"`
	MerchantId   string `yaml:"merchantId"`
	MerchantName string `yaml:"merchantName"`
	Secret       string `yaml:"secret"`
}

type MGDZ struct {
	Url       string `yaml:"url"`
	TokenUrl  string `yaml:"tokenUrl"`
	AgentCode string `yaml:"agentCode"`
	Secret    string `yaml:"secret"`
}

type IMTY struct {
	Url       string `yaml:"url"`
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

type JLDZ struct {
	Url       string `yaml:"url"`
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

type KXDZ struct {
	Url    string `yaml:"url"`
	Agent  string `yaml:"agent"`
	Md5Key string `yaml:"md5Key"`
	DesKey string `yaml:"desKey"`
}

type LGDDZ struct {
	Url       string `yaml:"url"`
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

type MWDZ struct {
	Url       string `yaml:"url"`
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

type PDTY struct {
	Url       string `yaml:"url"`
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

type SBTY struct {
	Url        string `yaml:"url"`
	VendorId   string `yaml:"vendorId"`
	OperatorId string `yaml:"operatorId"`
	Currency   string `yaml:"currency"`
}

type SGDZ struct {
	Url       string `yaml:"url"`
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

type TTQP struct {
	Url       string `yaml:"url"`
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}
