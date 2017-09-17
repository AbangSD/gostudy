# docker mysql

启一个 MySQL

```docker
// windows
docker run -d -p 127.0.0.1:3306:3306 --name shopapi -v C:\Users\AbangSD\docker\shop:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 mysql:latest
// Mac
docker run -d -p 127.0.0.1:3306:3306 --name shopapi -v /Users/abang/docker/shop:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 mysql:latest

docker run -d -p 3307:27017 --name shopmgo mongo

// -d 后台运行
// -p 本地 3306 端口对应 docker 3306
// --name 容器名字
// -v 创建数据卷 C:\Users\AbangSD\docker\database 对应 /var/lib/mysql
// -e MYSQL_ROOT_PASSWORD=123456 MySQL 密码
// mysql:latest 使用的镜像
```

启动后，进到 MySQL 内部

```docker
// Enter password:
docker exec -it shopapi mysql -uroot -p
```

```docker
// -uroot -p"password" 这是进到 MySQL 后的命令
docker exec -it shopapi mysql -uroot -p"password"
```

MySQL 基本操作

``注意 MySQL 保留``

``MySQL 命令``

```sql
// hostname 本机使用 127.0.0.1 或 localhost
mysql -h hostname -u username -p

// 修改 root 密码
/usr/bin/mysqladmin -u root password '123456'

// 查看当前所有存在的数据库
show databases;

// 创建数据库
create database database_name;

// 查看创建好的数据库
show create database database_name\G

// 删除数据库
drop database database_name;

// 查看数据库的存储引擎
show engines\G

// 选择当前数据库
use database_name

// 创建数据表
create table <表名>
(
字段名1, 数据类型 [列级别约束条件][默认值],
字段名2, 数据类型 [列级别约束条件][默认值],
...
[表级别约束条件]
);

// 查看数据表
show tables;

// 查看具体的表
describe table_name;
```

更多 MySQL操作 请查看

``
http://www.cnblogs.com/tuhooo/p/5441897.html
``

``
http://blog.csdn.net/cleanness/article/details/42967661
``