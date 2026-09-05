package composite

type FileSystemNode interface {
	Name() string
	Size() int
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int {
	return f.size
}

type Directory struct {
	name     string
	children []FileSystemNode
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:     name,
		children: []FileSystemNode{},
	}
}

func (d *Directory) Add(node FileSystemNode) {
	d.children = append(d.children, node)
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int {
	totalSize := 0
	for _, child := range d.children {
		totalSize += child.Size()
	}
	return totalSize
}
