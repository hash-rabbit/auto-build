# auto-build
golang 自动编译系统

## 使用
```shell
#normal
./auto-build ./config.toml
#nohup
nohup ./auto-build ./config.toml > nohup.log 2>&1 &
```

## 配置文件说明
```toml
port = 8000 # 监听端口
log_path = "./buildlog/" #编译/程序运行 log
go_env_path = "./goenv/" # go 环境安装目录
default_go_path = "./workspace/" # 针对 gomod 的 gopath 目录(缓存包)
dest_path = "./output/" # 输出文件目录
sql_file = "./dev.db" # sqlite 文件位置,会自动创建
web_path = "./dist/" # 前端目录,可以用下面的前端项目编译后的 dist 目录
```

## 前端
[链接](https://github.com/hash-rabbit/auto-build-web)