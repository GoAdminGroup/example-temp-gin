## for what

- this project used to goadmin with gin server
- [ ] rename `github.com/GoAdminGroup/example-temp-gin` to your api package name

## use

- `need go mod to management golang dependenceis`

- [ ] change API Port for dev replace `39000` to your want port
- [ ] change mysql Port for dev replace `39005` to your want port
- [ ] change redis Port for dev replace `39006` to your want port

```sh
$ make help
# check base dep
$ make init

# beacuse projec use mysql and reids so must run
$ make dockerDependDevFileInit dockerDependDevUp
# then see log to init database

# first run just use dep to get dep
$ make dep
# change conf/config.yaml
# run server as dev
$ make dev

# to remove dev docker and data can use
$ make dockerDependDevStop dockerDependDevContainRemove && rm -rf ./docker/example-temp-gin-dev

# run as docker contant
# Before run this project in docker must or can not find docker image
$ make dockerLocalFileRest dockerLocalImageRebuild
# if use linux
$ make dockerRunLinux
# if use macOS
$ make dockerRunDarwin
# stop or remove docker
$ make dockerStop
$ make dockerContainRemove
# more fast use of dockerStop dockerContainRemove dockerLocalImageRemove
$ make dockerBuildRemove

# can build less file run
# build less
$ make clean dockerLessBuild

# rest to dev mod
$ make dockerBuildRemove dockerLocalFileRest
```

# config

- config file is `config.yaml` demo see [conf/config.yaml](conf/config.yaml)

## log

```yaml
zap:
  AtomicLevel: -1 # DebugLevel:-1 InfoLevel:0 WarnLevel:1 ErrorLevel:2
  FieldsAuto: false # is use auto Fields key set
  Fields:
    Key: key
    Val: val
  Development: true # is open Open file and line number
  Encoding: console # output format, only use console or json, default is console
  rotate:
    Filename: log/temp-gin-api-swag.log # Log file path
    MaxSize: 16 # Maximum size of each zlog file, Unit: M
    MaxBackups: 10 # How many backups are saved in the zlog file
    MaxAge: 7 # How many days can the file be keep, Unit: day
    Compress: true # need compress
  EncoderConfig:
    TimeKey: time
    LevelKey: level
    NameKey: logger
    CallerKey: caller
    MessageKey: msg
    StacktraceKey: stacktrace
    TimeEncoder: ISO8601TimeEncoder # ISO8601TimeEncoder EpochMillisTimeEncoder EpochNanosTimeEncoder EpochTimeEncoder default is ISO8601TimeEncoder
    EncodeDuration: SecondsDurationEncoder # NanosDurationEncoder SecondsDurationEncoder StringDurationEncoder default is SecondsDurationEncoder
    EncodeLevel: CapitalColorLevelEncoder # CapitalLevelEncoder CapitalColorLevelEncoder LowercaseColorLevelEncoder LowercaseLevelEncoder default is CapitalLevelEncoder
    EncodeCaller: ShortCallerEncoder # ShortCallerEncoder FullCallerEncoder default is FullCallerEncoder
```

# dev

## evn

```bash
go version go1.13.4 darwin/amd64
gin version v1.5.0
go-admin v1.1.7
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

工程文件定义
