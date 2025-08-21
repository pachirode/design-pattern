package composite

type INode interface {
	Add(node INode)
	Tree(space int)
}

type Node struct {
	name string
}

func NewNode(name string) *Node {
	return &Node{name: name}
}

func (n *Node) Add(node INode) {
	panic("implement me")
}

func (n *Node) Tree(space int) {
	for i := 0; i < space; i++ {
		print(" ")
	}
	println(n.name)
}

type Folder struct {
	Node
	nodes []INode
}

func NewFolder(name string) *Folder {
	return &Folder{Node: Node{
		name: name,
	}}
}

func (f *Folder) Add(node INode) {
	if len(f.nodes) == 0 {
		f.nodes = make([]INode, 0)
	}

	f.nodes = append(f.nodes, node)
}

func (f *Folder) Tree(space int) {
	f.Node.Tree(space)

	for _, node := range f.nodes {
		node.Tree(space + 2)
	}
}

type File struct {
	Node
}

func NewFile(name string) *File {
	return &File{Node: Node{
		name: name,
	}}
}

func (f *File) Add(node INode) {
	panic("can not add node to file")
}

func (f *File) Tree(space int) {
	f.Node.Tree(space)
}
