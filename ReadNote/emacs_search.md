[参考链接](http://front.sjtu.edu.cn/~gwxie/search.html)

### 热键
* C-s 向前搜索/下一个
* C-r 向后搜索
* C-w 自动补全
* C-g 重回光标，搜索没成功C-g C-g
* C-s C-s 重新搜索前一次关键字
* C-s M-p 选择搜索过的关键字

### 配置
* 大小写匹配 默认设置的时候，你如果搜全是小写的字母的话，比如foobar，将会搜索所有大小写组成的foobar，如fooBar fOObAr都会成匹配，但如 果搜索带大写字母的单词如，Foobar，那么只和Foobar匹配，fooBar FoobaR 都不是匹配。 如果想更改这一特性的话，可以在.emacs中设置(setq-default case-fold-search nil)