基本操作

```go
1.克隆代码到本地：git -clone git仓库地址 [指定文件夹]
2.拉取分支最新代码：git pull;
3.查看文件修改状态：git status;
4.添加修改文件到缓存区：
	git add <file>：添加某个修改文件
	git add -a：添加所有修改文件
5.提交代码到本地仓库：git commit -m "描述信息"
6.推送代码到远程分支：git push [-f 分支名（强制推送，会覆盖别人提交的修改）] 
7.删除某次commit的提交
	1）从缓存区删除某个文件并再次提交：
		git rm --cached <file>
		git commit -m "describe"
	2）删除某次commit
		git log 查询提交日志，获取要回滚的commit_id
		git reset --soft commit_id 本次提交的修改被退回到缓存区（git status可查看）
		git reset --hard commit_id 本次提交的修改不做任何保留（git status无法查看）
8.放弃某次提交
	git revert commit_id
	git revert commit_id -m 指定提交点
9.查看和切换分支
	git branch 当前分支
	git branch -a 所有分支
	git checkout 分支名 切换当前分支到‘分支名’
10.合并某个分支到当前分支
	git merge 分支名
```

