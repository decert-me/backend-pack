package service

import (
	"backend-pack/internal/app/global"
	"backend-pack/internal/app/model"
	"backend-pack/internal/app/model/response"
	"backend-pack/internal/app/utils"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/tidwall/gjson"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

var l sync.Mutex

// PackRequest 处理打包请求
func PackRequest(tutorial model.Tutorial) (res response.PackResponse, err error) {
	l.Lock()
	defer l.Unlock()
	packPath := "assest/pack"
	resourcePath := "assest/resource"
	// 清空打包目录
	err = os.RemoveAll(packPath)
	if err != nil {
		res.Message = "清空打包目录失败"
		return res, err
	}
	// 创建打包目录
	err = os.MkdirAll(packPath, os.ModePerm)
	if err != nil {
		res.Message = "创建资源目录失败"
		return res, err
	}
	// 如果目录不存在则创建
	if exist, _ := utils.PathExists(resourcePath); !exist {
		err = os.MkdirAll(resourcePath, os.ModePerm)
		if err != nil {
			res.Message = "创建资源目录失败"
			return res, err
		}
	}
	data := []model.Tutorial{tutorial}
	// 生成JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		res.Message = "JSON生成失败"
		return res, err
	}
	// 将JSON数据写入文件
	file, err := os.Create(path.Join(global.CONFIG.Pack.Path, "tutorials.json"))
	if err != nil {
		fmt.Println("Error:", err)
		res.Message = "JSON写入失败"
		return res, err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		res.Message = "JSON写入失败"
		return res, err
	}
	// npm run build -- blockchain-basic
	args := []string{"run", "build"}
	dir := global.CONFIG.Pack.Path
	stdoutRes, stdoutErr, err := execCommand(global.CONFIG.Pack.Path, "npm", args...)
	if err != nil {
		fmt.Println(err)
		res.Message = "命令执行失败"
		return res, err
	}
	var packLog strings.Builder
	for _, v := range stdoutErr {
		packLog.WriteString(v + "<br />")
	}
	var success bool
	var startPage string
	for _, v := range stdoutRes {
		v = strings.Replace(v, "\n", "", -1)
		// 判断打包是否成功
		if !success && v == "Build completed successfully" {
			success = true
		}
		if gjson.Valid(v) {
			startPage = gjson.Get(v, "startPage").String()
		}
	}
	var status uint8
	if success {
		status = 2
	} else {
		status = 3
	}
	// 写入打包日志
	res.PackLog = model.PackLog{
		TutorialID: tutorial.ID,
		Status:     status,
	}
	if status == 2 {
		packLog.Reset()
		packLog.WriteString("打包成功  <br />")
		packLog.WriteString(fmt.Sprintf("打包时间：%s <br />", time.Now().Format("2006-01-02 15:04:05")))
		if tutorial.CommitHash != nil && *tutorial.CommitHash != "" {
			packLog.WriteString(fmt.Sprintf("版本：%s <br />", *tutorial.CommitHash))
		}
	}
	if tutorial.PackStatus == 2 && status == 3 {
		// 写入日志
		res.Tutorial = model.Tutorial{PackLog: packLog.String()}
		return res, nil
	}
	// 将结果写入
	res.Tutorial = model.Tutorial{StartPage: startPage, PackStatus: status, PackLog: packLog.String()}
	// 复制文件到发布项目路径
	utils.CopyContents(path.Join(dir, "build"), packPath)
	// 读取目录下文件夹
	entries, err := os.ReadDir(packPath)
	if err != nil {
		res.Message = "读取目录失败"
		return res, err
	}
	var dirName string
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println(entry.Name(), "is a directory")
			dirName = entry.Name()
			break
		}
	}
	_ = dirName
	//
	UUID := uuid.NewV4()
	// 压缩文件夹
	fileName := UUID.String() + ".zip"
	fmt.Println("packPath", packPath)
	fmt.Println("res", path.Join(resourcePath, fileName))
	err = utils.ZipDir(packPath, path.Join(resourcePath, fileName))
	if err != nil {
		res.Message = "压缩文件失败"
		return res, err
	}
	res.FileName = fileName
	// Build completed successfully
	// Error running build command:
	// 同步到远程服务器
	//if global.CONFIG.Sync.Enable && global.CONFIG.Sync.Type == 0 && global.CONFIG.Sync.Server != "" {
	//
	//}
	return res, nil
}
