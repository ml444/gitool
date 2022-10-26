# gitool
gitlab clone all repo or pull all...

## Install 
```shell
$ go install github.com/ml444/gitool@lastest
```

## Setting
```shell
export GITLAB_LOCAL_BASE_DIR=/Users/your/path
export GITLAB_ACCESS_TOKEN=YourAccessToken
export GITLAB_BASE_URL=https://git.your.cn/api/v4
export GITLAB_USERNAME=UserName
```
## Commands

```shell
# clone all repo:
$ gitool clone -all 

# clone all repo by setting concurrent count
$ gitool clone -all -C 5

# clone one repo by search 
$ gitool clone -s repo_name

# >>>>pull<<<< 
# pull all repo:
$ gitool pull -all 

# pull all repo by setting concurrent count
$ gitool pull -all -C 5

# pull one repo by search 
$ gitool pull -s repo_name

```
