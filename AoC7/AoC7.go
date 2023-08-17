package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int64
}

type Dir struct {
	Size         int64
	Name         string
	ParentDir    *Dir
	ContentFiles []*File
	ContentDirs  []*Dir
}

type FileSystem struct {
	RootDir *Dir
	CurDir  *Dir
}

func main() {
	lines := ReadInput("./input.txt")
	fs := ParseInput(lines)
	fs.RootDir.CalcDirSize()
	fmt.Println(Sol1(fs))
	fmt.Println(Sol2(fs))
}

func ReadInput(inputFilename string) []string {
	result := make([]string, 0)

	file, err := os.Open(inputFilename)
	if err != nil {
		return result
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func ParseInput(input []string) FileSystem {
	res := FileSystem{
		RootDir: &Dir{
			Name: "/",
		},
	}
	for _, line := range input {
		isCommand := strings.HasPrefix(line, "$ ")
		if isCommand {
			command, _ := strings.CutPrefix(line, "$ ")
			res.EvalCommand(command)
		} else {
			res.HandleCommandOutput(line)
		}
	}

	res.CurDir = res.RootDir
	return res
}

func (fs *FileSystem) EvalCommand(command string) {
	splitedCommand := strings.Split(command, " ")
	switch splitedCommand[0] {
	case "cd":
		if fs.CurDir == nil && splitedCommand[1] == "/" {
			fs.CurDir = fs.RootDir
			break
		}
		if splitedCommand[1] == ".." {
			fs.CurDir = fs.CurDir.ParentDir
			break
		}
		for _, contentDir := range fs.CurDir.ContentDirs {
			if contentDir.Name == splitedCommand[1] {
				fs.CurDir = contentDir
				break
			}
		}
	}
}

func (fs *FileSystem) HandleCommandOutput(output string) {
	isDir := strings.HasPrefix(output, "dir")
	splitedOutput := strings.Split(output, " ")
	objectName := splitedOutput[1]
	if isDir {
		newDir := Dir{
			Name:      objectName,
			ParentDir: fs.CurDir,
		}
		fs.CurDir.ContentDirs = append(fs.CurDir.ContentDirs, &newDir)
	} else {
		fileSize, _ := strconv.ParseInt(splitedOutput[0], 0, 64)
		newFile := File{
			Name: objectName,
			Size: fileSize,
		}
		fs.CurDir.ContentFiles = append(fs.CurDir.ContentFiles, &newFile)
	}
}

func (d *Dir) CalcDirSize() {
	var size int64 = 0

	for _, f := range d.ContentFiles {
		size += f.Size
	}

	for _, cd := range d.ContentDirs {
		cd.CalcDirSize()
		size += cd.Size
	}
	d.Size = size
}

func Sol1(fs FileSystem) int64 {
	return SumOfAcceptedDirs(fs.RootDir)
}

func SumOfAcceptedDirs(curDir *Dir) int64 {
	var res int64 = 0
	if curDir.Size <= 100000 {
		res += curDir.Size
	}
	for _, d := range curDir.ContentDirs {
		res += SumOfAcceptedDirs(d)
	}

	return res
}

func Sol2(fs FileSystem) int64 {
	var freeSpace = 70000000 - fs.RootDir.Size
	var bytesRequired = 30000000 - freeSpace
	return DirMinSizeToDelete(fs.RootDir, bytesRequired, fs.RootDir.Size)
}

func DirMinSizeToDelete(dir *Dir, bytesRequired int64, foundMin int64) int64 {
	var res int64 = foundMin
	if dir.Size > bytesRequired && foundMin > dir.Size {
		res = dir.Size
	}
	for _, d := range dir.ContentDirs {
		minInBranch := DirMinSizeToDelete(d, bytesRequired, res)
		if minInBranch < res {
			res = minInBranch
		}
	}
	return res
}
