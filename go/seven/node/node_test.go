package node

import (
	"testing"
)

func testNodeData() *Node {
	root := &Node{
		Type: Directory,
		name: "",
	}
	aFile := &Node{
		Type: File,
		name: "a-file",
		Size: 200,
	}
	bDir := &Node{
		Type: Directory,
		name: "b-dir",
	}
	cFile := &Node{
		Type: File,
		name: "c-file",
		Size: 2022,
	}
	dFile := &Node{
		Type: File,
		name: "d-file",
		Size: 1111,
	}
	eDir := &Node{
		Type: Directory,
		name: "e-dir",
	}
	fFile := &Node{
		Type: File,
		name: "f-file",
		Size: 3333,
	}
	gFile := &Node{
		Type: File,
		name: "g-file",
		Size: 4444,
	}

	eDir.AddChildren(fFile, gFile)
	bDir.AddChildren(cFile, eDir, dFile)
	root.AddChildren(aFile, bDir)

	return root
}

func TestSizes(t *testing.T) {
	root := testNodeData()
	_ = root.Size
	expect := 200 + 2022 + 1111 + 3333 + 4444
	got := root.Size
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}

	expect = 7777
	got = root.children[1].children[1].Size
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestSizeFilter(t *testing.T) {
	root := testNodeData()
	expect := 10910
	got := root.SizeFilter(11000, 0)
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestPath(t *testing.T) {
	root := testNodeData()
	expect := "/b-dir/e-dir/g-file (4444 bytes)"
	got := root.Children()[1].Children()[1].Children()[1].Path()
	if got != expect {
		t.Errorf("Expected: %s, got: %s", expect, got)
	}
}

func TestCDRoot(t *testing.T) {
	root := testNodeData()
	expect := root
	got := root.Children()[1].Children()[1].Children()[1].Root()
	if got != expect {
		t.Errorf("Expected %#v, got %#v", expect, got)
	}
}

func TestFindOrCreateDirChild(t *testing.T) {
	root := testNodeData()
	expect := &Node{Type: Directory, name: "TestFindOrCreateDirChild"}
	got := root.FindOrCreateChild(Directory, "TestFindOrCreateDirChild", 0)
	if got.name != expect.name {
		t.Errorf("Expected: %s, got: %s", expect.name, got.name)
	}
	again := root.FindOrCreateChild(Directory, "TestFindOrCreateDirChild", 0)
	if again != got {
		t.Errorf("Expected: %#v, got: %#v", expect, got)
	}
}

func TestFindOrCreateFileChild(t *testing.T) {
	root := testNodeData()
	expect := &Node{Type: File, name: "TestFindOrCreateFileChild", Size: 12345}
	got := root.FindOrCreateChild(File, "TestFindOrCreateFileChild", 12345)
	if got.name != expect.name {
		t.Errorf("Expected: %s, got: %s", expect.name, got.name)
	}
	again := root.FindOrCreateChild(File, "TestFindOrCreateFileChild", 12345)
	if again != got {
		t.Errorf("Expected: %#v, got: %#v", expect, got)
	}
}
