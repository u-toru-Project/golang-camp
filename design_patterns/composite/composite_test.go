package composite

import "testing"

func TestCompositeSize(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() FileSystemNode
		expected int
	}{
		{
			name: "単一のファイルのサイズ",
			setup: func() FileSystemNode {
				return NewFile("test.txt", 100)
			},
			expected: 100,
		},
		{
			name: "空のディレクトリのサイズ",
			setup: func() FileSystemNode {
				return NewDirectory("empty_dir")
			},
			expected: 0,
		},
		{
			name: "ファイルが複数入ったディレクトリのサイズ",
			setup: func() FileSystemNode {
				dir := NewDirectory("docs")
				dir.Add(NewFile("a.txt", 10))
				dir.Add(NewFile("b.txt", 20))
				return dir
			},
			expected: 30,
		},
		{
			name: "入れ子になったディレクトリ",
			setup: func() FileSystemNode {
				rootDir := NewDirectory("root")
				subDir := NewDirectory("sub")

				subDir.Add(NewFile("sub_file.txt", 50))
				rootDir.Add(subDir)
				rootDir.Add(NewFile("root_file.txt", 100))

				return rootDir
			},
			expected: 150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := tt.setup()
			if got := node.Size(); got != tt.expected {
				t.Errorf("Size() = %d, want %d", got, tt.expected)
			}
		})
	}
}
