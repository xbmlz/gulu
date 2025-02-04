// Gulu - Golang common utilities for everyone.
// Copyright (c) 2019-present, b3log.org
//
// Gulu is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package gulu

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestIsValidFilename(t *testing.T) {
	if !File.IsValidFilename("hello.go") {
		t.Errorf("[hello.go] should be a valid filename")
	}
	if File.IsValidFilename("hello?.go") {
		t.Errorf("[hello?.go] should not be a valid filename")
	}
}

func TestWriteFileSaferByReader(t *testing.T) {
	writePath := "testdata/filewrite.go"
	defer os.RemoveAll(writePath)
	if err := File.WriteFileSaferByReader(writePath, strings.NewReader("test"), 0644); nil != err {
		t.Errorf("write file [%s] failed: %s", writePath, err)
	}
}

func TestWriteFileSafer(t *testing.T) {
	writePath := "testdata/filewrite.go"
	defer os.RemoveAll(writePath)
	if err := File.WriteFileSafer(writePath, []byte("test"), 0644); nil != err {
		t.Errorf("write file [%s] failed: %s", writePath, err)
	}
}

func TestIsHidden(t *testing.T) {
	filename := "./file.go"
	isHidden := File.IsHidden(filename)
	if isHidden {
		t.Error("file [" + filename + "] is not hidden")
	}
}

func TestGetFileSize(t *testing.T) {
	filename := "./file.go"
	size := File.GetFileSize(filename)
	if 0 > size {
		t.Error("file [" + filename + "] size is [" + strconv.FormatInt(size, 10) + "]")
	}
}

func TestIsExist(t *testing.T) {
	if !File.IsExist(".") {
		t.Error(". must exist")
		return
	}
}

func TestIdBinary(t *testing.T) {
	if File.IsBinary("not binary content") {
		t.Error("The content should not be binary")
		return
	}
}

func TestIsImg(t *testing.T) {
	if !File.IsImg(".jpg") {
		t.Error(".jpg should be a valid extension of a image file")
		return
	}
}

func TestIsDir(t *testing.T) {
	if !File.IsDir(".") {
		t.Error(". should be a directory")
		return
	}
}

func TestCopyDir(t *testing.T) {
	source := "testcopydir"
	os.Mkdir(source, 0644)
	dest := filepath.Join(testdataDir, source)
	defer os.Remove(dest)

	err := File.CopyDir(source, dest)
	if nil != err {
		t.Error("Copy dir failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() != destStat.ModTime() {
		t.Error("mod time is not equal")
		return
	}
}

func TestCopyFile(t *testing.T) {
	source := "./file.go"
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := File.CopyFile(source, dest)
	if nil != err {
		t.Error("Copy file failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() != destStat.ModTime() {
		t.Error("mod time is not equal")
		return
	}
}

func TestCopy(t *testing.T) {
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := File.Copy("./file.go", dest)
	if nil != err {
		t.Error("Copy failed: ", err)
		return
	}

	sourceStat, _ := os.Stat("./file.go")
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() != destStat.ModTime() {
		t.Error("mod time is not equal")
		return
	}
}

func TestCopyDirNewtimes(t *testing.T) {
	source := "testcopydir"
	os.Mkdir(source, 0644)
	dest := filepath.Join(testdataDir, source)
	defer os.Remove(dest)

	time.Sleep(100 * time.Millisecond) // CI

	err := File.CopyDirNewtimes(source, dest)
	if nil != err {
		t.Error("Copy dir failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() == destStat.ModTime() {
		t.Error("mod time is equal")
		return
	}
}

func TestCopyFileNewtimes(t *testing.T) {
	source := "./file.go"
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := File.CopyFileNewtimes(source, dest)
	if nil != err {
		t.Error("Copy file failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() == destStat.ModTime() {
		t.Error("mod time is equal")
		return
	}
}

func TestCopyNewtimes(t *testing.T) {
	source := "./file.go"
	dest := filepath.Join(testdataDir, "file.go")
	defer os.Remove(dest)
	err := File.CopyNewtimes(source, dest)
	if nil != err {
		t.Error("Copy failed: ", err)
		return
	}

	sourceStat, _ := os.Stat(source)
	destStat, _ := os.Stat(dest)
	if sourceStat.ModTime() == destStat.ModTime() {
		t.Error("mod time is equal")
		return
	}
}
