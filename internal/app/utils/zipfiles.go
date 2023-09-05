package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//@function: ZipFiles
//@description: 压缩文件
//@param: filename string, files []string, oldForm, newForm string
//@return: error

func ZipFiles(filename string, files []string, oldForm, newForm string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = newZipFile.Close()
	}()

	zipWriter := zip.NewWriter(newZipFile)
	defer func() {
		_ = zipWriter.Close()
	}()

	// 把files添加到zip中
	for _, file := range files {

		err = func(file string) error {
			zipFile, err := os.Open(file)
			if err != nil {
				return err
			}
			defer zipFile.Close()
			// 获取file的基础信息
			info, err := zipFile.Stat()
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			// 使用上面的FileInforHeader() 就可以把文件保存的路径替换成我们自己想要的了，如下面
			header.Name = strings.Replace(file, oldForm, newForm, -1)

			// 优化压缩
			// 更多参考see http://golang.org/pkg/archive/zip/#pkg-constants
			header.Method = zip.Deflate

			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			if _, err = io.Copy(writer, zipFile); err != nil {
				return err
			}
			return nil
		}(file)
		if err != nil {
			return err
		}
	}
	return nil
}

// ZipDir 将指定的文件夹压缩为zip文件
func ZipDir(source, target string) error {
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// 获取文件夹的绝对路径
	absFolderPath, err := filepath.Abs(source)
	if err != nil {
		fmt.Println("无法获取文件夹路径:", err)
		return err
	}
	source = absFolderPath
	//fmt.Println("absFolderPath:", absFolderPath)
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//fmt.Println("run")
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			fmt.Println("relPath")
			return err
		}
		// 将文件路径转换为相对于文件夹的路径
		relPath, err := filepath.Rel(absFolderPath, path)
		if err != nil {
			return err
		}
		//fmt.Println(relPath)
		//fmt.Println(filepath.Join(filepath.Base(source), path))
		header.Name = relPath
		//header.Name = filepath.Join(filepath.Base(source), path)

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})

	return nil
}
