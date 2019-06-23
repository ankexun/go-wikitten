package main

import (
	"log"

	"github.com/fsnotify/fsnotify"

	"path/filepath"

	"os"
)

type Watch struct {
	watch *fsnotify.Watcher
}

type Tree struct {
	Name  string
	IsDir bool
}

// 全局变量
var Data []Tree

// 判断目录是否存在
func PathExists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, nil
	}
	// err错误是否报告了一个文件或者目录不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//监控目录
func (w *Watch) watchDir(dir string) {
	// 判断目录是否存在,不存在则创建
	exist, err := PathExists(dir)
	if err != nil {
		log.Printf("get dir error![%v]\n", err)
		return
	}

	if exist {
		log.Printf("has dir![%v]\n", dir)
	} else {
		log.Printf("no dir![%v]\n", dir)
		// 创建目录
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Printf("mkdir failed![%v]\n", err)
		} else {
			log.Printf("mkdir success!\n")
		}
	}
	//通过Walk来遍历目录下的所有子目录
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		//这里判断是否为目录，只需监控目录即可

		//目录下的文件也在监控范围内，不需要我们一个一个加

		if info.IsDir() {

			// path, err := filepath.Abs(path)

			// if err != nil {

			// 	return err

			// }

			// err = w.watch.Add(path)
			path = filepath.Clean(path)
			err := w.watch.Add(path)

			if err != nil {

				return err

			}
			log.Printf("监控 : %s\n", path)
			//目录分隔统一转换成"/"
			path = filepath.ToSlash(path)

			Data = append(Data, Tree{path, true})

		} else {
			log.Printf("文件 : %s\n", path)
			//目录分隔统一转换成"/"
			path = filepath.ToSlash(path)

			Data = append(Data, Tree{path, false})
		}

		return nil

	})

	go func() {

		for {

			select {

			case ev := <-w.watch.Events:

				{

					if ev.Op&fsnotify.Create == fsnotify.Create {

						log.Println("创建文件 : ", ev.Name)

						//这里获取新创建文件的信息，如果是目录，则加入监控中

						fi, err := os.Stat(ev.Name)
						name := Tree{filepath.ToSlash(ev.Name), false}

						if err == nil && fi.IsDir() {

							w.watch.Add(ev.Name)

							log.Println("添加监控 : ", ev.Name)

							name.IsDir = true
						}
						Data = append(Data, name)
					}

					if ev.Op&fsnotify.Write == fsnotify.Write {

						log.Println("写入文件 : ", ev.Name)

					}

					if ev.Op&fsnotify.Remove == fsnotify.Remove {

						log.Println("删除文件 : ", ev.Name)

						for k, v := range Data {
							if v.Name == filepath.ToSlash(ev.Name) {
								Data = append(Data[:k], Data[k+1:]...)
							}
						}
						//如果删除文件是目录，则移除监控

						fi, err := os.Stat(ev.Name)

						if err == nil && fi.IsDir() {

							w.watch.Remove(ev.Name)

							log.Println("删除监控 : ", ev.Name)

						}

					}

					if ev.Op&fsnotify.Rename == fsnotify.Rename {

						log.Println("重命名文件 : ", ev.Name)

						for k, v := range Data {
							if v.Name == filepath.ToSlash(ev.Name) {
								Data = append(Data[:k], Data[k+1:]...)
							}
						}
						//如果重命名文件是目录，则移除监控

						//注意这里无法使用os.Stat来判断是否是目录了

						//因为重命名后，go已经无法找到原文件来获取信息了

						//所以这里就简单粗爆的直接remove好了

						w.watch.Remove(ev.Name)

					}

					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {

						log.Println("修改权限 : ", ev.Name)

					}

				}

			case err := <-w.watch.Errors:

				{

					log.Println("error : ", err)

					return

				}

			}

		}

	}()

}

//目录在前,文件在后排序
func sortData() []Tree {
	var sortedDir, sortedFile []Tree
	for k, v := range Data {
		if v.IsDir {
			sortedDir = append(sortedDir, Data[k])
		} else {
			sortedFile = append(sortedFile, Data[k])
		}
	}
	return append(sortedDir, sortedFile...)
}
