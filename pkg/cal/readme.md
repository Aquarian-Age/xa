## 历法
- `data.go` map数组 阳历有效时间范围为1600~3500

## 数据库
MariaDB
- cal.sql 有效时间同`data.go`的map数组
- ccal.sql 有效时间-4000~8000
- 使用ccal.sql需要设置
- my.cnf
```text
[mysqld]
max_allowed_packet=100M #大于20M
```

## 重启服务
```bash
systemctl restart mysqld.service
```

## 导入数据库
```bash
mysql -u root -pxxxxx
create database ccal;
use ccal;
source /path/ccal.sql
```

## 参考

[农历编算法则](https://ytliu0.github.io/ChineseCalendar/rules_simp.html)

## 数据来源

[TDBtimes.txt](https://raw.githubusercontent.com/ytliu0/ChineseCalendar/master/src/TDBtimes.txt)