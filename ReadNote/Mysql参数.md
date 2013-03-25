参考：http://lxneng.iteye.com/blog/451985

### 查看Mysql服务器配置项
`show variables;`

### 查看服务器运行状态项
`show global status;`

### 慢查询
`show variables like '%slow%';`

- low_slow_queries OFF 记录慢查询开关
- slow_launch_time 2 大于2秒的查询记录慢查询

`show global status like '%slow%';`

- Slow_launch_threads 0
- Slow_queries 270 一共有270条慢查询记录 

