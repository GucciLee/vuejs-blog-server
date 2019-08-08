package filepath

import (
	"os"
	"io"
)


// @Title 拷贝文件
// @Param string srcFileName 文件路径
// @Param string destFileName 目标文件路径
// @return (int64, error) 复制字节数, 错误信息
func CopyFile(srcFileName, destFileName string) (int64, error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return 0, err
	}
	destFile, err := os.OpenFile(destFileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()
	defer destFile.Close()
	return io.Copy(destFile, srcFile)
}

// @Title 新建文件
// @Param string src文件路径
// @return (int64, error) 复制字节数, 错误信息
func CreateFile(src string) bool {
	_, err := os.Create(src)
	if err != nil  {
		return false
	} else {
		return true
	}
}