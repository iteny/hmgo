package file

import "os"

// IsExist checks whether a file or directory exists,If it exists it will return false.
// IsExist 检查文件或目录是否存在，如果文件或目录不存在将返回false。
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
