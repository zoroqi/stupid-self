# cron

[cron](https://github.com/robfig/cron) golang实现的一个定时模块 [文档](https://godoc.org/github.com/robfig/cron)

基本使用

```golang
c := cron.New() // 创建一个调度
stop := make(chan struct{})
count := 0
// 一分钟一次
c.AddFunc("*/1 * * * *", func() {
	fmt.Println(time.Now())
	count++
	if count > 2 {
		stop <- struct{}{}
	}
})
c.Start() // 开始执行调度
<-stop
```

整体结构

伪代码
```
for {
	case <-ticker.C: // 根据最小间隔计算
    	for _,j := range jobs {
        	if j.isZero() { // 判断是否需要执行
        		j.Next(now) // 计算下次执行时间
        		go j.Run()  // 执行
            }
        }
    case <-stop:
    	break
}
```

## 巧妙地设计

之前自己想实现一个解析cron表达式的工具, 自己实现费劲的要死. 我当初的结构大概是是. 通过时间计算. 
```java
public class TimeField {
	private int[] runs; // 也可能是用的Set<Integer>
    private Type type;
}

public enum Type {
	MIN, HOUR, DAY, MONTH, WEEK
}
```

作者使用的方案是`uint64`存储时间间隔. 加一个type, 感觉比我的简单很多, 而且可以做进一步的运算. `uint64`刚好可以囊括, 秒, 分, 时, 日, 月, 周的所有时间范围.

以分钟为例
```
00100000 00000000 00000000 00000000
就是3分钟的时候发起执行
```

最高位`==1`表示`*`

## 伟大的goto

验证方式,
```golang
WRAP:
	// 月份
	for 1<<t.Month() & s.Month == 0 {
    	t = t.AddDate(0, 1, 0)
        if t.Month() == time.January {
			goto WRAP
		}
    }
    // 日和周, 比较复杂但逻辑相似
    
    // 小时
    for 1<<t.Hour() & s.Hour == 0 {
    	t = t.Add(1 * time.Hour)
        if t.Day() == 1 {
			goto WRAP
		}
    }
    // 分钟
    // 秒

```

 每一个时间分割都有一个, 类似的判断, 将会goto到开始处. 
```golang
if t.Day() == 1 {
	goto WRAP
}
```
这里涉及到到当前时间发生了滚动就要, 就需要从高位在进行一次判断. 

`0 0 31 */2 *` 每个月31号执行, 实际执行日期是 1,3,5,7月, 其他月份无法执行.  9月满足月份逻辑, 但日期在执行会使下次日期执行进入10月1日, 这种情况需要从月份计算再一次时间.  

这里可以改成循环也需要使用到Label, 和goto差不多, 还是goto吧.


## 另一种方案

作者采用的是计算下次运行时间方案. 还有一种是验证方式.  这种方式不需要计算nextTime, 只需要验证就可以. 

验证方式,
```goalng
for {
	case <-ticker.C // 分钟
    	for _,j := range jobs {
        	if j.isRun() { // 现在时间是否需要执行
            	go j.Run()
            }
        }
}
```

两种方法区别, 计算型要从大时间向小时间算, 验证是从小时间向大时间验证.  验证型没法支持crontab意外的定时方式, spring框架下的 `fixedDelay` 上次执行完后间隔执行

当然看需求, 复杂的优先选择计算型, 要是简单的话选验证性