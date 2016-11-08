package Figo

import (
	"github.com/quexer/utee"
	"log"
	"testing"
	"github.com/bmizerany/assert"
)

func TestFileOpen(t *testing.T) {
	_, err := FileOpen("./test.txt")
	utee.Chk(err)
	fp := FilePath("./test.txt")
	fullPath, err := fp.FullPath()
	utee.Chk(err)
	log.Println(fullPath)
	fp = FilePath(fullPath)
	log.Println(fp.WindowsPath())
	log.Println(fp.UnixPath())
	fp = FilePath(fp.UnixPath())
	log.Println(fp.WindowsPath())
}

func TestFileExist(t *testing.T){
	assert.Equal(t,true,FileExist("./file.go"))
	assert.Equal(t,false,FileExist("./fileNotExist"))
	log.Println("test pass")
}