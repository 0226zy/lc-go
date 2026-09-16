package designinmemoryfilesystem

import (
	"sort"
	"strings"
)

// node 文件系统节点，既可以是目录也可以是文件
type node struct {
	children map[string]*node // 目录：名称 -> 子节点
	isFile   bool
	content  strings.Builder // 文件内容（仅 isFile 为 true 时有效）
}

func newNode() *node {
	return &node{children: make(map[string]*node)}
}

// FileSystem 设计内存文件系统
// 用前缀树（Trie）模拟目录树，支持 ls / mkdir / addContentToFile / readContentFromFile。
// 时间复杂度: ls 为 O(d·L + k·log k)，其余操作为 O(d·L)
// 空间复杂度: O(路径分量总数 + 文件内容总长)
type FileSystem struct {
	root *node
}

// Constructor 初始化文件系统，只有一个根目录 "/"
func Constructor() FileSystem {
	return FileSystem{root: newNode()}
}

// split 把 "/a/b/c" 切分为 ["a","b","c"]，根路径 "/" 切分为空列表
func split(path string) []string {
	segs := make([]string, 0)
	for _, seg := range strings.Split(path, "/") {
		if seg != "" {
			segs = append(segs, seg)
		}
	}
	return segs
}

// walk 沿路径分量查找节点，返回最终节点；路径不存在时返回 nil
func (fs *FileSystem) walk(segs []string) *node {
	cur := fs.root
	for _, seg := range segs {
		next, ok := cur.children[seg]
		if !ok {
			return nil
		}
		cur = next
	}
	return cur
}

// walkOrCreate 沿路径分量查找节点，不存在的层自动创建
func (fs *FileSystem) walkOrCreate(segs []string) *node {
	cur := fs.root
	for _, seg := range segs {
		if _, ok := cur.children[seg]; !ok {
			cur.children[seg] = newNode()
		}
		cur = cur.children[seg]
	}
	return cur
}

// Ls 列出路径内容：文件返回自身名称，目录返回按字典序排列的所有条目
func (fs *FileSystem) Ls(path string) []string {
	segs := split(path)
	target := fs.walk(segs)
	if target.isFile {
		// path 是文件，返回文件名本身
		return []string{segs[len(segs)-1]}
	}
	names := make([]string, 0, len(target.children))
	for name := range target.children {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Mkdir 创建目录，中间目录不存在时一并创建；已存在则不做任何操作
func (fs *FileSystem) Mkdir(path string) {
	fs.walkOrCreate(split(path))
}

// AddContentToFile 向文件追加内容，文件不存在时创建
func (fs *FileSystem) AddContentToFile(filePath string, content string) {
	target := fs.walkOrCreate(split(filePath))
	target.isFile = true
	target.content.WriteString(content)
}

// ReadContentFromFile 返回文件的全部内容
func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	return fs.walk(split(filePath)).content.String()
}
