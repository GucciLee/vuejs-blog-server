package learnku_filepath

import "os"

// 判断所给路径文件/文件夹是否存在
func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}