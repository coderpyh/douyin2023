# douyin2023
第五届字节跳动青训营后端进阶班极简抖音项目2023队

项目运行方法：

1.创建数据库。MySQL执行create database douyin2023;。建议创建前首先将MySQL的默认字符集设为utf8mb4。

2.修改连接数据库的代码。repository文件夹init.go文件mysql.Open语句，账号、密码、端口号根据自己MySQL的设置进行修改。

3.设置服务器地址。端口号为8888，客户端中填写地址为http://服务器的IP地址:8888/。

4.运行项目。命令行在douyin2023文件夹下执行go run douyin2023。

5.服务端成功开始运行。
