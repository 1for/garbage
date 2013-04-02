## 2 求值实践

### 2.1 缓冲区名
`(buffer-name)` 显示缓冲区名
`(buffer-file-name)` 显示文件名，包含绝对路径

### 2.2 获得缓冲区
`(current-buffer)` 获得当前buffer,buffer-name仅仅获得缓冲区名
`(other-buffer)` 最近使用过的buffer

### 2.3 切换缓冲区
`(switch-to-buffer (other-buffer))` 切换到之前的buffer
`(set-buffer)` 设置活动buffer

### 2.4 缓冲区大小和位点定位
`(buffer-size)` 缓冲区大小
`(point)` 当前位点的字符个数

