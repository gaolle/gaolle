---
title: git
date: 2020-10-27 01:15:50
tags: git
categories: 版本控制
---

# Git

切换至git仓库目录（最顶层文件目录）

##### 初始化仓库（默认进入主分支（master））

```shell
git init
```

##### 查看状态

```shell
git status
```

##### 添加文件到仓库（临时缓存区）

```shell
git add filename
```

##### 将文件提交到仓库

```shell
git commit -m "committext"
```

##### 添加文件并提交

```shell
git commit filename -m "committext"
```

##### 提交至远程仓库

```shell
git push origin/master
```

##### 查看日志

```shell
git log
```

##### 查看分支情况（*代表当前分支）

```shell
git branch
```

##### 增加分支

```shell
git branch branchname
```

##### 切换当前分支（当前尾部有提示所在行）

```shell
git checkout branchname
```

##### 创建并切换至新分支

```shell
git checkout -b branchname
```

##### 创建远程分支

```shell
git checkout -b branchname origin/branchname
```

##### 合并分支

无冲突

```shell
git merge branchname
```

存在冲突

```shell
解决冲突文件
上面为自己新代码，下面为合并的代码
git add mergefile
git commit
git push
```

##### 删除分支（主分支不能删除）

```shell
git branch -d branchname
```

##### SSH密钥和公钥生产（/c/Users/username/.ssh）id_rsa（密钥）id_rsa.pub（公钥） 公钥添加到Github

```shell
ssh-keygen.exe
```

##### 添加公匙之后测试

```shell
ssh -T git@github.com 
```

##### 关联远程仓库

```shell
git remote add origin githubaddress
```

##### 同步远程仓库

```shell
git pull
```

##### 强制拉取并覆盖本地代码

```shell
git fetch --all
git reset --hard origin/master
git pull
```

##### 放弃本次commit

```shell
git reset --soft HEAD^
```

##### 隐藏本地修改

```shell
git stash
```

##### 返回本地修改

```shell
git stash pop
```
