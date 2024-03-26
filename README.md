# backend-pack
## 安装
```bash
git clone https://github.com/decert-me/backend-pack.git
cd backend-pack
```
## 编译
```bash
# 主程序
go build -o backend-pack
```
## 配置
```bash
cp ./config/config.demo.yaml ./config/config.yaml
vi ./config/config.yaml
```
克隆 tutorials 项目
```bash
git clone https://github.com/decert-me/tutorials.git
cd tutorials
yarn -i
```
配置 tutorials 项目路径
```shell
# 配置文件
pack:
  path: "/Users/mac/Code/tutorials" # tutorials 项目目录
```
## 运行
```bash
# 主程序
./backend-pack
```

## 配置说明

### 运行端口配置

配置项：
```yaml
# system configuration
system:
  env: develop
  addr: 9092
```
env：运行环境，可选值为 develop、test、production

addr：运行端口

### 打包配置

配置项：

```yaml
pack:
  path: "/Users/mac/Code/tutorials"
```

path: [tutorials项目本地路径](https://github.com/decert-me/tutorials)
