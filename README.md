日志组件默认读取项目根目录下面的config.ini配置文件

文件格式

[log]<br/>
level = info<br/>
filename = xxx.log.%Y-%m-%d

xxx为日志文件名前缀

日志分割的时间为一天分割一次

最大保存30天
