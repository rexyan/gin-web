### 介绍
1. JWT: 认证方案
2. Gin: web 框架
3. SQLX: 数据库查询
4. Swagger: 自动文档 
5. Zap + lumberjack: 日志
6. Viper: 配置管理，配置监听
7. Snowflake: 分布式 ID 生成器
8. Redis + MySQL: 数据库，缓存
9. Ratelimit: 限流策略


### 运行
1. 克隆代码
```shell
git clone git@github.com:rexyan/gin-web.git
```

2. 安装依赖
```shell
go mod dity
```

3. 编译 && 运行
```shell
make build && make run
```

### Swagger
1. 安装 [swag](https://github.com/swaggo/gin-swagger)
```shell
go get -u github.com/swaggo/swag/cmd/swag
```

2. 生成 swagger 文档
```shell
swag init
```

3. 访问 http://127.0.0.1:8081/swagger/index.html 即可

### 配置示例
```yaml
server:
  mode: "debug"
  name: "application"
  port: 8081
  secret: "123"

logger:
  level: "debug"
  filename: "./logs/web_app.log"
  maxSize: 200
  maxAge: 30
  maxBackups: 7

mysql:
  host: "127.0.0.1"
  port: 3306
  user: "root"
  password: ""
  dbname: "gin_web"
  maxOpenConns: 100
  maxIdleConns: 200

redis:
  db: 0
  host: "localhost"
  port: 6379
  password: ""
  poolSize: 100
```