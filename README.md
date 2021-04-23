# goleaf

## 简介
[美团 Leaf](https://github.com/Meituan-Dianping/Leaf) 的 Go 语言版本。供学习使用。

## 快速开始
1. 执行 `srcript/segments.sql` 脚本。初始化数据库。

2. 修改 pkg/db/db.go 中的数据库连接信息：
```go
db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goleaf")
```

3. 运行系统

4. 交叉编译
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

## 参考
https://tech.meituan.com/2019/03/07/open-source-project-leaf.html

https://tech.meituan.com/2017/04/21/mt-leaf.html