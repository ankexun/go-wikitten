package main

import (
	"log"
	"sort"

	"github.com/fsnotify/fsnotify"

	"path/filepath"

	"os"
)

type Watch struct {
	watch *fsnotify.Watcher
}

type FileNode struct {
	Name      string      `json:"name"`
	IsDir     bool        `json:"isDir"`
	Path      string      `json:"path"`
	FileNodes []*FileNode `json:"children"`
}

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

// 使用递归遍历
func walk(path string, info os.FileInfo, node *FileNode) {
	// 列出当前目录下的所有目录、文件
	files := listFiles(path)

	// 遍历这些文件
	for _, filename := range files {
		// 拼接全路径
		fpath := filepath.ToSlash(filepath.Clean(filepath.Join(path, filename)))

		// 构造文件结构
		fio, _ := os.Lstat(fpath)

		// 将当前文件作为子节点添加到目录下
		child := FileNode{filename, fio.IsDir(), fpath, []*FileNode{}}
		node.FileNodes = append(node.FileNodes, &child)

		// 如果遍历的当前文件是个目录，则进入该目录进行递归
		if fio.IsDir() {
			walk(fpath, fio, &child)
		}
	}

	return
}

// 相当于ls
func listFiles(dirname string) []string {
	f, _ := os.Open(dirname)

	names, _ := f.Readdirnames(-1)
	f.Close()

	sort.Strings(names)

	return names
}

// 递归查找后增加
func addFileNode(child []*FileNode, fileinfo bool, path, filename string) {
	for i := 0; i < len(child); i++ {
		if child[i].Path != path || len(child[i].FileNodes) != 0 {
			addFileNode(child[i].FileNodes, fileinfo, path, filename)
		} else {
			log.Println("add前: ", child[i].FileNodes)
			child[i].FileNodes = append(child[i].FileNodes, &FileNode{filename, fileinfo, path, []*FileNode{}})
			log.Println("add后: ", child[i].FileNodes)
		}
	}
}

// 递归查找后删除
func deleteFileNode(child []*FileNode, path string) {
	for i := 0; i < len(child); i++ {
		if child[i].Path != path || len(child[i].FileNodes) != 0 {
			deleteFileNode(child[i].FileNodes, path)
		} else {
			log.Println("delete: ", i, child[i])
			child = append(child[:i], child[i+1:]...)
		}
	}

	return
}

//监控目录
func (w *Watch) watchDir(dir string) *FileNode {
	// 判断目录是否存在,不存在则创建
	exist, err := PathExists(dir)
	if err != nil {
		log.Printf("get dir error![%v]\n", err)
		return nil
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

	rootpath := filepath.Clean(dir)
	root := FileNode{dir, true, rootpath, []*FileNode{}}
	fileInfo, _ := os.Lstat(rootpath)

	walk(rootpath, fileInfo, &root)

	// 通过Walk来遍历目录下的所有子目录
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		//这里判断是否为目录，只需监控目录即可

		//目录下的文件也在监控范围内，不需要我们一个一个加
		if info.IsDir() {
			path = filepath.Clean(path)
			err := w.watch.Add(path)

			if err != nil {
				return err
			}
			log.Printf("监控 : %s\n", path)
		} else {
			log.Printf("文件 : %s\n", path)
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

						if err == nil && fi.IsDir() {

							w.watch.Add(ev.Name)

							log.Println("添加监控 : ", ev.Name)
						}
						// filename := filepath.Base(ev.Name)
						// path := filepath.Dir(ev.Name)
						// addFileNode(root.FileNodes, fi.IsDir(), filepath.ToSlash(path), filename)
					}

					if ev.Op&fsnotify.Write == fsnotify.Write {

						log.Println("写入文件 : ", ev.Name)

					}

					if ev.Op&fsnotify.Remove == fsnotify.Remove {

						log.Println("删除文件 : ", ev.Name)

						//如果删除文件是目录，则移除监控
						fi, err := os.Stat(ev.Name)

						if err == nil && fi.IsDir() {

							w.watch.Remove(ev.Name)

							log.Println("删除监控 : ", ev.Name)

						}

						// deleteFileNode(root.FileNodes, filepath.ToSlash(ev.Name))
					}

					if ev.Op&fsnotify.Rename == fsnotify.Rename {

						log.Println("重命名文件 : ", ev.Name)

						//如果重命名文件是目录，则移除监控

						//注意这里无法使用os.Stat来判断是否是目录了

						//因为重命名后，go已经无法找到原文件来获取信息了

						//所以这里就简单粗爆的直接remove好了

						w.watch.Remove(ev.Name)

						// deleteFileNode(root.FileNodes, filepath.ToSlash(ev.Name))
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

	return &root
}
