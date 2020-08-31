我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# Library-Api
Creating a REST API for books at Holberton School, using Go and Postgresql

## Usage
First Install postgres db

```
sudo apt-get install postgresql postgresql-contrib

```

+ Add the following environmental variables:
```

export DB_USERNAME="your_user_name"
export DB_PASSWD="your_passwd"
export DB_NAME="your_db_name"

```
+ to test your application inside the src file simply run:

```
go test -v
```
+ Then if you are running a Linux machine simply go to bin

```
$ ./src
```
Otherwise head over to the src directory and after setting up your GOPATH and your GOBIN run go build or go install.


+ Once the server is running head over to http://localhost:8080/books to see the
books inside your database.

+ If you have not added books yet, you can do manually inside psql or

```
curl -i -X POST -H 'Content-Type: application/json' -d '{"name":"book name", "description":"book description"}' http://localhost:8080/book

```

## TODO

+ add a front end to interact with api
