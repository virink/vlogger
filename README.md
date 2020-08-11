# vlogger

convenient log package

# 1. 使用说明
```go
    import  "github.com/virink/vlogger"

    // 设置完成后，即可在控制台和日志文件app.log中看到如下输出
	vlogger.Success("Success")
	vlogger.Failed("Failed")
	vlogger.Normal("Normal")
	// vlogger.Panic("Error")
	vlogger.Error("Error")
	vlogger.Debug("Debug")
```
输出结果：

# 2. 日志等级

等级由高到底，当配置为某个输出等级时，只有大于等于该等级的日志才会输出。不同的输出适配器支持不同的日志等级配置：

| 等级 | 配置    | 释义     | 控制台颜色 |
| ---- | ------- | -------- | ---------- |
| 0    | PANIC   | 系统错误 | 红底白色   |
| 1    | ERROR   | 用户错误 | 红色       |
| 3    | NORMAL  | 普通     | 白色       |
| 3    | SUCCESS | 成功     | 绿色       |
| 3    | FAILED  | 失败     | 紫色       |
| 4    | DEBUG   | 调试     | 黄色       |


