package composite

import "testing"

func TestComposite(t *testing.T) {
	driveD := NewFolder("DriveD")
	driveD.Add(NewFile("file1"))
	driveD.Add(NewFile("file2"))
	driveD.Add(NewFile("file3"))
	driveD.Add(NewFolder("folder1"))
	driveD.Add(NewFolder("folder2"))
	folder3 := NewFolder("folder3")
	folder3.Add(NewFile("file4"))
	driveD.Add(folder3)
	driveD.Tree(0)
}
