# 后端

## 安装依赖

```sh
go mod tidy
```

## 新建环境变量

在后端目录里新建.env，里面应该有类似这样的东西

```env
DB_USERNAME=root
DB_PASSWORD=@Ws123456
DB_HOST=localhost
DB_PORT=3306
DB_NAME=jd
DB_LOG_LEVEL=silent
REDIS_HOST=localhost
REDIS_PORT=6379
```

## 启动后端

```sh
go run .
```