package goi

import "strings"

//使用前缀树实现动态路由解析
type node struct {
	pattern  string  //待匹配的路由
	part     string  //路由中的一部分
	children []*node //子节点
	isWild   bool    //是否精确匹配
}

//根据待匹配串匹配某个节点中的子节点，返回第一个匹配中的子节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//根据待匹配串匹配某个节点中的子节点，返回所有匹配中的节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//把一条url加入到前缀树中,直到匹配到一个url的最后一部分才把整个url设置进去
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

//根据url到前缀树中查找，找到最后一个节点
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
