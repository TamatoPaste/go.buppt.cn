package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	//找类并加载class文件
	//参数是相对路径，用正斜杠分割，文件名带.class后缀
	//返回值是最终读取的字节码，最终定位到的class文件的Entry和错误信息
	String() string
	// 类似java中的toString
}

func newEntry(path string) Entry {
	// 根据参数不同，创建不同的Entry实例
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
