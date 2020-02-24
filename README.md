## for what

- this project used to goadmin with gin server
- [ ] rename `github.com/GoAdminGroup/example-temp-gin` to your api package name

## use

- `need dep to management golang dependenceis`, will change to go mod

```sh
$ make help
# check base dep
$ make init
# first run just use
$ make dep
# change conf/config.yaml
# run server as dev
$ make dev
```

# config

- config file is `config.yaml` demo see [conf/config.yaml](conf/config.yaml)

## log

+ `writers`: 输出位置，有2个可选项：file,stdout。选择file会将日志记录到`logger_file`指定的日志文件中，选择stdout会将日志输出到标准输出，当然也可以两者同时选择
+ `logger_level`: 日志级别，DEBUG, INFO, WARN, ERROR, FATAL
+ `logger_file`: 日志文件
+ `log_format_text`: 日志的输出格式，json或者plaintext，`false`会输出成json格式，`true`会输出成非json格式
+ `rollingPolicy`: rotate依据，可选的有：daily, size。如果选daily则根据天进行转存，如果是size则根据大小进行转存
+ `log_rotate_date`: rotate转存时间，配合`rollingPolicy: daily`使用
+ `log_rotate_size`: rotate转存大小，配合`rollingPolicy: size`使用
+ `log_backup_count`:当日志文件达到转存标准时，log系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数。

# dev

## evn

```bash
go version go1.13.4 darwin/amd64
gin version v1.5.0
go-admin v1.2.2
```
 
## fast-use

### db-init

- use docker-compose to init db entry [help db-docker-compose](db/README.md#db-docker-compose)
- and install db tools see [db/README.md](db/README.md)

```bash
# use script to init goAdmin base db
make dbPostgreImportAdmin
# init full biz db need db password
make dbPostgreAllBiz
```

### project-run

```bash
# check depends
make dep
# run dev
make dev
```

### adm

```bash
# adm install
make admInstall
# update adm
make admUpdate
# use see
make helpAdm
```

## folder-Def

```
.
├── Dockerfile
├── LIB.md
├── MakeAdm.mk
├── MakeDockerRun.mk
├── MakeGoMod.mk
├── Makefile     # make utils
├── README.md    # readme
├── api                         # api package for api
│   ├── api.go                  # api loader
│   └── demo
│       └── popup.go
├── build                       # build dir, not in git management
├── conf                        # service config folder makeFile use conf/test conf/release
│   └── config.yaml
├── config                      # load conf package
│   ├── baseConf.go
│   ├── config.go
│   ├── logConf.go
│   ├── probe.go
│   └── watchConf.go
├── db                          # db management
│   ├── README.md
│   └── default
│       ├── 1-demo.ini          # adm load config file
│       ├── admin.db
│       └── demo.sql
├── docker-compose.yml          # docker-compse use to run as docker 
├── go.mod                      # go mod package management
├── go.sum                      # not in git management
├── log                         # log folder, not in git management
│   └── server.log
├── main.go                     # program entry
├── model                       # model some db or model struct
│   └── dbglobal
│       └── db.go
├── pages                       # admin pages folder
│   └── index
│       ├── config.go
│       └── dashborad.go
├── pkg                         # some pkg for use in plural projects
│   ├── errdef
│   │   ├── errcode.go
│   │   ├── errdef.go
│   │   ├── errdef_test.go
│   │   └── noroute.go
│   └── tableutil
│       └── optionsYesOrNo.go
├── router                      # admin loader in router
│   ├── dbConn.go               # db connect
│   ├── display.go              # cli display
│   ├── interceptor             # rouder package for load tables
│   │   ├── admin.go
│   │   └── externalLink.go     # externalink page
│   ├── language.go
│   ├── middleware              # gin middleware
│   │   ├── header.go           # services header set
│   │   └── logger.go           # logger middlerware
│   ├── monitor.go              # monitor
│   ├── noRouteBiz.go           # no router biz
│   ├── plugin                  # admin plugin
│   │   ├── admin.go
│   │   └── example.go
│   ├── router.go               # router entry
│   └── static.go               # static file
├── tables                      # admin tables folder
│   ├── authors.go
│   ├── demo
│   │   ├── demo_class.go
│   │   ├── demo_student.go
│   │   ├── demo_student_class.go
│   │   ├── demo_student_score.go
│   │   └── tables.go
│   ├── posts.go
│   ├── tables.go
│   └── users.go
└── util                        # project util
    ├── folder
    │   └── path.go
    ├── security
    │   ├── base64.go
    │   ├── base64_test.go
    │   ├── stringplus.go
    │   ├── stringplus_test.go
    │   ├── unicode.go
    │   └── unicode_test.go
    ├── sys
    │   ├── network.go
    │   └── network_test.go
    └── timestamp
        └── now.go

```