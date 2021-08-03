### gin-web
gin + sqlx + zap + viper + jwt + swagger + ...

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

3. 访问地址
http://127.0.0.1:8081/swagger/index.html

