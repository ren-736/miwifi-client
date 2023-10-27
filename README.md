# 小米路由器客户端

参考[movsb][1]这位大佬的[gpon-client][2]项目，简单修改了下路由器的api借口。
这个程序用来在命令行下控制小米路由器的一些常用配置（已连接设备列表/端口映射/路由器网络信息），本人设备（AX6S）。

## 帮助

```bash
$ ./miwifi-client
./gpon-client 
A MiWIFI client used to modify router configurations

Usage:
  ./gpon-client [command]

Available Commands:
  devices     manage devices
  gwinfo      show gateway information
  help        Help about any command
  portmaps    manage port mappings

Flags:
  -h, --help   help for ./miwifi-client

Use "./miwifi-client [command] --help" for more information about a command.
```

## 运行环境初始化配置

请先在命令行下导出几个环境变量。

```bash
# 路由器IP地址，默认 192.168.31.1
$ export MI_IP=192.168.31.1

# 路由器用户名，默认 admin
$ export MI_USER=admin

# 路由器密码，无默认
$ export MI_PSD=
```

## 示例使用

### 设备列表

```bash
$ ./miwifi-client devices list
Name                IPv4              Type      MAC
----------------------------------------------------------------------------------------------------
60:BE:B4:08:71:33   192.168.31.10     0         60:BE:B4:08:71:33
AE:0B:38:C4:2F:28   192.168.31.157    2         AE:0B:38:C4:2F:28
```

### 端口转发配置

#### 枚举端口映射

```bash
$ ./miwifi-client portmaps list
ID   Name            Protocol    OuterPort   InnerIP             InnerPort
-----------------------------------------------------------------------------------
1    Ubuntu-Docker   1           12376       192.168.31.10       2376
2    Ubuntu-SSH      1           10022       192.168.31.10       22

```

#### 添加端口映射

```bash
$ ./miwifi-client portmaps create nginx 3 10443 192.168.31.2 443
```

#### 删除端口映射

```bash
$ ./miwifi-client portmaps delete 10443
```

### 查看网关信息

```bash
$ ./miwifi-client gwinfo
LAN IPV4: 192.168.31.1
LAN MAC:  D4:DA:21:74:15:1B
WAN IPV4: 192.168.1.2
WAN MAC:  D4:DA:21:1C:07:08
```

[1]: https://github.com/movsb      "movsb"
[2]: https://github.com/movsb/gpon-client  "gpon-client"
