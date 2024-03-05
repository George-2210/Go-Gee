package gee

import (
	"fmt"
	"strings"
)

// 路由节点结构体
type node struct {
	pattern  string  // 路由规则
	part     string  // 路由节点的一部分
	children []*node // 子节点列表
	isWild   bool    // 是否为通配符节点
}

// String 返回节点的字符串表示
func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

// insert 向节点插入路由规则
func (n *node) insert(pattern string, parts []string, height int) {
	// 如果已经到达路由规则的末尾，则将当前节点的路由规则设置为传入的路由规则
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 获取当前层级的路由部分
	part := parts[height]

	// 查找当前节点的子节点，看是否存在与当前路由部分匹配的子节点
	child := n.matchChild(part)
	if child == nil {
		// 如果没有匹配的子节点，则创建一个新的子节点
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	// 递归向下插入路由规则
	child.insert(pattern, parts, height+1)
}

// search 根据路由规则查找节点
func (n *node) search(parts []string, height int) *node {
	// 如果已经到达路由规则的末尾，或者当前节点是一个 * 通配符节点，则返回当前节点
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	// 获取当前层级的路由部分
	part := parts[height]

	// 查找当前节点的子节点，看是否存在与当前路由部分匹配的子节点
	children := n.matchChildren(part)

	// 遍历子节点，继续向下查找
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

// travel 遍历节点，将节点添加到列表中
func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

// matchChild 查找与给定路由部分匹配的子节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 查找与给定路由部分匹配的所有子节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
