# 说明
这个文件夹下包含文件(目录)遍历的pkg示例
相比于标准库filepath.Walk 这些包的核心价值在于提供了"并行"能力,让你可以更快的扫描磁盘上的文件

# 异步化
这里先区分一下"扫描过程的异步化"和"对扫描结果的处理异步化"是两个不同的概念
前者的逻辑是将将一个或者多个RootPath拆分成互不交叉的组(subRootPath) 然后分别异步的对他们进行扫描,
这在多磁盘环境下是比较有用的:
    比如 rootpath = /mnt 下挂在两个磁盘 diskA 和 diskB
根目录/mnt可以拆解成:/mnt/diskA/ 和 /mnt/diskB/两个互不交叉的子目录,使用两个"线程"分别对他们进行扫描
后者则关注的是对扫描出来结果处理的异步化

# 杀毒软件误报
由于这些pkg的行为特性(扫描用户磁盘)
以及他们作为开源通用pkg会被其他程序引用(可能是有一些恶意软使用了这些包)
会导致你编译出来的结果被杀毒软件误报为病毒.
比如我尝试编译skywalker的示例 就让Windows Defender无情的杀死了.
Trojan:Win64/Wingo.psyA!MTB
其他包也会存在被其他杀毒软件误报的风险.
如果存在这类问题.建议使用其他pkg代替,或者fork一份代码 尝试修改混淆一下
当然最好的方式还是参考他们自己实现逻辑.

# 示例编译Size(无任何优化参数)
不同功能,不同逻辑,不同代码.这样比较其实没有意义,只是作为一个参考
walker       2160k
godirwalk    2171k
skywalker    2444k

## skywalker提供了更多的功能:
1.设置"工作协程"数量,
需要注意的是skywalker似乎只是对filepath.Walk的简单封装,换句话说它的扫描过程依然是同步的,只是对扫描结果的处理可以异步的.这并不符合我的预期.
    err := filepath.Walk(sw.Root, sw.walker(workerChan))
你可以设置NumWorkers来控制对扫描出来结果处理的并发
    ```golang
    &Skywalker{
		Root:       root,
		NumWorkers: 20,
		QueueSize:  100,
		Worker:     worker,
		FilesOnly:  true,
	}
    ```
2.支持glob,黑名单、白名单、后缀过滤等,代价是引入github.com/gobwas/glob 这个包,
这使得编译结果会明显比其他两个大上一点(当然,我知道这么比较是非常不合理的,而且也很可能不是glob这个pkg的锅)
或许你并不需要glob,你可能希望使用正则表达式或者更简单github.com/tidwall/match来做这个事情

将特定的过滤规则硬编码进去其实不是一个很好的设计,作为一个通用的框架,更好的方式是提供一个filter接口,然后由用户来实现具体筛选过滤出哪些文件
(是根据后缀,正则或者glob,文件大小,文件内容,属性等)完全可以由filter来操作,这样还能避免框架本身依赖特定第三方pkg

3.他扫描获得一个文件列表
这是另一个问题:它的回调方法只包含一个路径 path string 因此你的处理逻辑重新打开文件来获取相关信息.
```
//Worker is anything that knows what to do with a path.
type Worker interface {
	Work(path string)
}
```
4.杀毒软件误报
包含该包的软件会被win安全中心误报为木马程序
Windows Defender
Trojan:Win64/Wingo.psyA!MTB

### 

# ...
另外一些跟文件扫描管理有关的工具 比如 gdu/syncthing等也有自己的目录扫码实现 如果需要自己实现目录并行扫描,其实可以参考他们