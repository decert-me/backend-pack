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
git checkout feat-backend
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