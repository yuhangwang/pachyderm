package hashtree

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/pachyderm/pachyderm/src/client/pkg/require"
)

func TestGlobFile1(t *testing.T) {
	h := &HashTree{}
	fooNode := &Node{
		FileNode: &FileNode{
			Name: "foo",
		},
	}
	dirNode := &Node{
		DirNode: &DirectoryNode{
			Name: "dir",
		},
	}
	dirBarNode := &Node{
		FileNode: &FileNode{
			Name: "bar",
		},
	}
	dirBuzzNode := &Node{
		FileNode: &FileNode{
			Name: "buzz",
		},
	}
	h.Fs = make(map[string]*Node)
	h.Fs["/foo"] = fooNode
	h.Fs["/dir"] = dirNode
	h.Fs["/dir/bar"] = dirBarNode
	h.Fs["/dir/buzz"] = dirBuzzNode

	for _, pattern := range []string{"*", "/*"} {
		nodes, err := h.GlobFile(pattern)
		require.NoError(t, err)
		require.Equal(t, 2, len(nodes))
		for _, node := range nodes {
			require.EqualOneOf(t, []interface{}{fooNode, dirNode}, node)
		}
	}

	for _, pattern := range []string{"*/*", "/*/*"} {
		nodes, err := h.GlobFile(pattern)
		require.NoError(t, err)
		require.Equal(t, 2, len(nodes))
		for _, node := range nodes {
			require.EqualOneOf(t, []interface{}{dirBarNode, dirBuzzNode}, node)
		}
	}
}

func TestListDir(t *testing.T) {
	t := HashTree{}
	t.PutFile("/dir/a", nil)
	t.PutFile("/dir/b", nil)
	t.PutFile("/dir/c", nil)

	fmt.Println(proto.MarshalTextString(t))
	t.Fail()
}