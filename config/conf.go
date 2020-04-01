package config

var global Conf

func Global() Conf {
	return global
}

type Conf struct {
	RunMode    string  `json:"run_mode" mapstructure:"run_mode"`
	Name       string  `json:"name" mapstructure:"name"`
	Addr       string  `json:"addr" mapstructure:"addr"`
	ApiUrl     string  `json:"api_url" mapstructure:"api_url"`
	ApiVersion string  `json:"api_version" mapstructure:"api_version"`
	SslEnable  bool    `json:"ssl_enable" mapstructure:"ssl_enable"`
	Zap        Zap     `json:"zap" mapstructure:"zap"`
	Monitor    Monitor `json:"monitor" mapstructure:"monitor"`
	GoAdmin    GoAdmin `json:"go_admin" mapstructure:"go_admin"`
	Redis      Redis   `json:"Redis" mapstructure:"Redis"`
}

type Monitor struct {
	Debug          bool           `json:"debug" mapstructure:"debug"`
	Hardware       bool           `json:"hardware" mapstructure:"hardware"`
	Health         string         `json:"health" mapstructure:"health"`
	PProf          bool           `json:"pprof" mapstructure:"pprof"`
	RetryCount     int            `json:"retryCount" mapstructure:"retryCount"`
	Security       bool           `json:"security" mapstructure:"security"`
	SecurityUser   SecurityUser   `json:"securityUser" mapstructure:"securityUser"`
	Status         bool           `json:"status" mapstructure:"status"`
	StatusHardware StatusHardware `json:"status_hardware" mapstructure:"status_hardware"`
}

type Swagger struct {
	Description string `json:"description" mapstructure:"description"`
	Host        string `json:"host" mapstructure:"host"`
	Index       string `json:"index" mapstructure:"index"`
	Root        string `json:"root" mapstructure:"root"`
	Security    bool   `json:"security" mapstructure:"security"`
	Title       string `json:"title" mapstructure:"title"`
	User        User   `json:"user" mapstructure:"user"`
	Version     int    `json:"version" mapstructure:"version"`
}

type Zap struct {
	AtomicLevel   int           `json:"AtomicLevel" mapstructure:"AtomicLevel"`
	Development   bool          `json:"Development" mapstructure:"Development"`
	EncoderConfig EncoderConfig `json:"EncoderConfig" mapstructure:"EncoderConfig"`
	Encoding      string        `json:"Encoding" mapstructure:"Encoding"`
	Fields        Fields        `json:"Fields" mapstructure:"Fields"`
	FieldsAuto    bool          `json:"FieldsAuto" mapstructure:"FieldsAuto"`
	Rotate        Rotate        `json:"rotate"  mapstructure:"rotate"`
}

type SecurityUser struct {
	Admin string `json:"admin" mapstructure:"admin"`
}

type StatusHardware struct {
	CPU  string `json:"cpu" mapstructure:"cpu"`
	Disk string `json:"disk" mapstructure:"disk"`
	RAM  string `json:"ram" mapstructure:"ram"`
}

type User struct {
	Admin string `json:"admin" mapstructure:"admin"`
	User  string `json:"user" mapstructure:"user"`
}

type EncoderConfig struct {
	CallerKey      string `json:"CallerKey" mapstructure:"CallerKey"`
	LevelKey       string `json:"LevelKey" mapstructure:"LevelKey"`
	MessageKey     string `json:"MessageKey" mapstructure:"MessageKey"`
	NameKey        string `json:"NameKey" mapstructure:"NameKey"`
	StacktraceKey  string `json:"StacktraceKey"  mapstructure:"StacktraceKey"`
	TimeKey        string `json:"TimeKey" mapstructure:"TimeKey"`
	TimeEncoder    string `json:"TimeEncoder" mapstructure:"TimeEncoder"`
	EncodeDuration string `json:"EncodeDuration" mapstructure:"EncodeDuration"`
	EncodeLevel    string `json:"EncodeLevel" mapstructure:"EncodeLevel"`
	EncodeCaller   string `json:"EncodeCaller" mapstructure:"EncodeCaller"`
}

type Fields struct {
	Key string `json:"Key" mapstructure:"Key"`
	Val string `json:"Val" mapstructure:"Val"`
}

type Rotate struct {
	Compress   bool   `json:"Compress" mapstructure:"Compress"`
	Filename   string `json:"Filename" mapstructure:"Filename"`
	MaxAge     int    `json:"MaxAge" mapstructure:"MaxBackups"`
	MaxBackups int    `json:"MaxBackups" mapstructure:"MaxAge"`
	MaxSize    int    `json:"MaxSize" mapstructure:"Compress"`
}

type GoAdmin struct {
	Title       string    `json:"title" mapstructure:"title"`
	Theme       string    `json:"theme" mapstructure:"theme"`
	ColorScheme string    `json:"color_scheme" mapstructure:"color_scheme"`
	Captcha     bool      `json:"captcha" mapstructure:"captcha"`
	DashBoard   DashBoard `json:"dash_board" mapstructure:"dash_board"`
	DataBases   DataBases `json:"data_bases" mapstructure:"data_bases"`
	IndexURL    string    `json:"index_url" mapstructure:"index_url"`
	Language    string    `json:"language" mapstructure:"language"`
	Store       Store     `json:"store" mapstructure:"store"`
	URLPrefix   string    `json:"url_prefix" mapstructure:"url_prefix"`
	UseCustom   bool      `json:"use_custom" mapstructure:"use_custom"`
	Custom      Custom    `json:"custom" mapstructure:"custom"`
}

type Custom struct {
	LogoUrl string `json:"logo_url" mapstructure:"logo_url"`
}

type DashBoard struct {
	Title string `json:"title" mapstructure:"title"`
}

type DataBases struct {
	Default `json:"default" mapstructure:""`
}

type Store struct {
	Path   string `json:"path" mapstructure:"path"`
	Prefix string `json:"prefix" mapstructure:"prefix"`
}

type Default struct {
	Driver     string `json:"driver" mapstructure:"driver"`
	Host       string `json:"host" mapstructure:"host"`
	MaxIdleCon int    `json:"max_idle_con" mapstructure:"max_idle_con"`
	MaxOpenCon int    `json:"max_open_con" mapstructure:"max_open_con"`
	Name       string `json:"name" mapstructure:"name"`
	Port       string `json:"port" mapstructure:"port"`
	Pwd        string `json:"pwd" mapstructure:"pwd"`
	User       string `json:"user" mapstructure:"user"`
}

type Redis struct {
	Addr         string `json:"addr" mapstructure:"addr"`
	Password     string `json:"password" mapstructure:"password"`
	DB           int    `json:"db" mapstructure:"db"`
	MaxRetries   int    `json:"max_retries" mapstructure:"max_retries"`
	DialTimeout  int    `json:"dial_timeout" mapstructure:"dial_timeout"`
	ReadTimeout  int    `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout int    `json:"write_timeout" mapstructure:"write_timeout"`
}
