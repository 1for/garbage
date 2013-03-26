## 设置合适的 `innodb_buffer_pool_size`

Setting buffer pool a bit larger than your database size will be enough,You need buffer pool a bit (say 10%) larger than your data (total size of Innodb TableSpaces) because it does not only contain data pages – it also contain adaptive hash indexes, insert buffer, locks which also take some time.

另外可以参考数据库运行数据 `show innodb status\G`

> 数据来自测试机81数据库

### 查看BUFFER POOL AND MEMORY一节

	Total memory allocated 624322114
	in additional pool 7251456
	Dictionary memory 797608
	Buffer pool size   32768	# 这个值表示页数 即Innodb_buffer_pool_pages_data， Innodb_buffer_pool_pages_data * Innodb_page_size = 使用的内存大小
	Free buffers       1
	Database pages     32585
	Modified db pages  0
	Pending reads 0
	Pending writes: LRU 0, flush list 0, single page 0
	Pages read 11800, created 866022, written 3340741
	0.00 reads/s, 0.00 creates/s, 0.22 writes/s
	Buffer pool hit rate 1000 / 1000

### 设置 `innodb_log_file_size`

If you change the parameter in the my.cnf file and restart the server.InnoDB will refuse to start because the existing log files don’t match the configured size.You need to shut the server down cleanly and normally, and move away (don’t delete) the log files, which are named ib_logfile0, ib_logfile1, and so on.
The basic point is that your log file needs to be big enough to let InnoDB optimize its I/O, but not so big that recovery takes a long time.

	mysql> show engine innodb status\G select sleep(60); show engine innodb status\G
	Log sequence number 5 725147216
	1 row in set (0.06 sec)
	 
	1 row in set (1 min 0.00 sec)
	 
	Log sequence number 5 738169434
	1 row in set (0.05 sec)

(5738169434-5725147216)/1024/1024 = 12.4M/Min  

an hour’s worth is more than enough so that it can reorder the writes to use sequential I/O during the flushing and checkpointing process.

12.4 * 60 = 744M 

and there's two log file, so we configure innodb_log_file_size 384M. 

PS. I think 10 minute is also enough, so I will set it to 64M.



