### vim 

**插入 **

i 光标前

I 行首

a 光标后

A 行末

o 当前行的下一行

O 当前行的上一行

cw 删除当前单词光标到末尾

**退出模式**

ESC建、ctrl+[

**分屏**

上下分屏

ctrl+w s

左右分屏

ctrl+w v

上下屏切换

ctrl+w h l

左右屏切换

ctrl+w j k

移动分屏

ctrl+w H J K L

关闭分屏

ctrl+w c

**打开保存文件**

a 所有 

q 退出 

！强制 

w 保存 

e 重新打开新文件

**设置行号**

set nu 显示行号

: n 某行

**关键字查找**

/?

n 向下查找

N 向上查找

**删除多行文本**

n, n d

**复制多行**

t

**粘贴**

p

**翻页操作**

ctrl+f 查看下一页

ctrl+b 查看上一页

**tar**

压缩

tar -zcvf tarname tarfile

解压

tar -zxvf tarname

**zip**

压缩

zip -r zipname zipfile

解压

unzip zipfile

**命令行操作**

CTRL+左右键 单词跳转

CTRL+a 行首

CTRL+e 行末

CTRL+u 删除光标前内容并复制

CTRL+k 删除光标后内容并复制

CTRL+y 粘贴

##### du

```
-b 以bytes为单位显示文件大小
-h 以K,M,为单位显示文件大小
-m 以1MB为单位显示文件大小
du option filename
```

##### df

```
显示文件系统磁盘使用情况统计
-h 刻度格式
df option
```

##### glances

```

```

##### cp

```
-f 覆盖
-r 复制目录文件
```

##### rm

```
-i 删除前询问
-rf 删除文件夹下全部
```

##### mv（移动）

```

```

shell

##### find

```
-amin n 过去几分钟被读取过
-cmin n 过去几分钟被修改过
-ctime n 过去几天被修改过
+n 表示超过那个时间 -n 表示在那个之间之内
-name 文件名字
-i 不区分大小写
find path option
```

##### awk

```

```

##### grep

```
-i 不区分大小写
-c 匹配的行数
-v 查找不匹配的行
-n 显示查找的行号和内容
grep "被查找的字符串" 文件名
```

##### sed

```
-e 以脚本处理文本文件
-f 以脚本文件处理文本文件
	a 新增
	c 取代
	d 删除
	i 插入
```

##### sort

```
-r 逆序
-o 输出排序结果到指定文件
sort filename
```

##### uniq

```
-c 行重复的次数
-d 显示重复的行
-u 显示不重复的行

uniq filename

sort filename | uniq options
```

##### cut

```

```

##### crontab

```

```

##### ping

```
-i 间隔秒数
-c 回应次数
```

##### telnet

```

```

##### scp

```
-C 允许压缩
-r 递归整个目录

scp t.rs devel@192.168.196.19:/home/devel/
scp devel@192.168.196.19:/home/devel/t.rs ~
```

##### ifconfig

```

```

##### ps

```
-u 某使用者的进程
-aux 显示所有包含其他使用者的进程
-ef 显示所有进程信息，连同命令行
```

##### lsof

```
lsof -i:端口号 查看端口占用情况
```

##### netstat

```
-t tcp相关
-u udp相关
-n 拒绝显示别名
-l 仅列出在监听的服务状态
-p 显示建立相关链接的程序名
netstat -tunlp | grep port
```

##### dlv 

```
dlv core corefile executable
```

##### kill

```
-1 重新加载进程
-9 强杀
-15 正常停止进程
kill 进程PID
```

##### killall

```
killall 进程名字
```

##### xargs 

```
标准输入转换为命令行参数
-I {} 将读取到的数据一行一行赋值给{}
-n num 命令在执行的时候一次使用argument的个数，默认所有
-d 分隔符 默认分隔符为空格
```

##### exec 

```

```



#### awk：格式化处理文本文件

```shell
cat text.txt
a b c d e f 
```

以空格分隔，打印第1、4项

```shell
awk '{print $1 $4}' text.txt
ad 
```

以空格分隔，打印第1、4项（按照指定格式）

```shell
awk '{printf "%-8s %-10s\n",$1,$4}' text.txt
a        d                
```

设置变量并打印出来

```shell
awk -va=1 '{print $2 $1+a}' text.txt
b1 
```

运行脚本

```shell
awk -f {awk脚本} {文件名}
```

```shell
cat text.txt
1 A                                                                                       2 B                                                                                       3 C                                                                                       4 D                                                                                       5 E                                                                                       6 F  
```

过滤第一列大于2并且第二列等于“E”的行

```shell
awk '$1>2 && $2=="E" {print $1,$2,$3}' text.txt
5 E    
```

输出第二列包含“E”，并打印第一列

```shell
awk '$2 ~ /E/ {print $1}' text.txt
5       
```

模式取反（输出第二列包含“E”，并打印第一列）

```shell
/text# awk '$2 !~ /E/ {print $1}' text.txt
1                                                                                         2                                                                                         3                                                                                         4                                                                                         6 
```

从文件中找出长度大于1的行

```shell
awk 'length>1' text.txt
1 A                                                                                       2 B                                                                                       3 C                                                                                       4 D                                                                                       5 E                                                                                       6 F                                                                         
```

#### sed利用脚本处理文件

```shell
cat sedtext                                                                               1 a                                                                                       2 b                                                                                       3 c                                                                                       4 d                          
```

在第4行之后添加5e

```shell
sed -e 4a\5e sedtext                                                                     1 a                                                                                       2 b                                                                                       3 c                                                                                       4 d                                                                                       5e                     
```

以行号显示文件，并删除第5行

```shell
nl sedtext | sed '5d'                                                                     1  1 a                                                                                   2  2 b                                                                                   3  3 c                                                                                   4  4 d                
```

以行号显示文件，并删除4到最后一行

```shell
nl sedtext | sed '4,$d'                                                                 
1  1 a                                                                         
2  2 b                                                                               
3  3 c                
```

将2-5行的内容替换为2 b

```shell
nl sedtext | sed '2,5c  2 b'                                                             1  1 a                                                                                   2 b                         
```

列出2-3行

```shell
nl sedtext | sed -n '2,3p'                                                               2  2 b                                                                                   3  3 c                  
```

搜索与b关键字有关的行

```shell
nl sedtext | sed -n '/b/p'                                                               2  2 b                
```

搜索与b关键字有关的行,并删除

```shell
nl sedtext | sed '/b/d'                                                                   1  1 a                                                                                   3  3 c                                                                                   4  4 d                                                                                   5  5e                  
```

直接修改文件中的内容

```shell
sed -i '2a 2 b' sedtext                                                                   cat sedtext                                                                               1 a                                                                                       2 b                                                                                       2 b                                                                                       3 c                                                                                       4 d                                                                                       5e                                    
```

#### ls：列出指定目录下的文件

* -a 显示所有文件
* -l 详细信息
* -t 时间顺序
* -R 层序文件输出

uniq：检查反复出现的行列

```shell
cat uniqtest                                                                             A                                                                                          A                                                                                         A                                                                                          B                                                                                         B                                                                                          C                            
```

删除重复的行

```shell
uniq uniqtest
A
B
C
```

删除重复的行,并统计次数

```shell
uniq -c uniqtext                                                                         3 A                                                                                       2 B                                                                                       1 C  
```

统计重复的行

```shell
uniq -d  uniqtext                                                                         A                                                                                         B                               
```

#### wc：用于计算字数

行数、字数、字节数

```shell
wc uniqtext                                                                             
6  6 12 uniqtext          
```

#### ip：用于显示或设置网络设备

* link：网络设备
* address：设备上的地址
* route：路由表条目
* rule：路由规则

#### vim快捷方式

##### Normal->Insert

* i键：当前光标插入
* a键：光标下一处插入

##### Normal

光标快速移动：

* hjkl 左下上右
* w 下一单词首位
* b 当前首位
* gg 文本首位
* shift + g 文本末尾
* number + shift + g 转至number行
* o 当前光标下方插入一行
* shift + o 当前光表插入一行
* x 删除光标字符
* shift + x 删除光标之前字符
* dd 删除当前行并将内容防止剪贴板
* de 删除当前光标之后并将内容防止剪贴板
* daw 删除光标所在的整个单词
* u 还原上一个操作
* shift + 4 行末
* shift + 6 行首

v 进入可视模式

* y 复制选定块
* yy 整行赋值
* p 粘贴

查找与替换

* f + 键 光标之后查找键
* F + 键 光标之前查找键
* r + 键  将光标值替换为键
