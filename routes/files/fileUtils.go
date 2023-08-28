package files

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

func InitConstants() {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	CONTENT_ROOT_PATH = path.Join(rootPath, "Files")
	FILESYSTEM = os.DirFS(CONTENT_ROOT_PATH)
}

func fileStat(bodyPath string) (*fileStruct, error) {
	filePath := path.Join(CONTENT_ROOT_PATH, bodyPath)
	fileSystem := os.DirFS(filePath)
	fileInfo, err := fs.Stat(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	cwd := fileStruct{
		Name:         path.Base(filePath),
		Size:         getSize(fileInfo.Size()),
		IsFile:       !fileInfo.IsDir(),
		DateModified: fileInfo.ModTime(),
		Type:         path.Ext(filePath),
		FilterPath:   bodyPath,
		Permission:   nil,
		HasChild:     fileInfo.IsDir(),
	}
	return &cwd, nil
}

func getSize(size int64) string {
	stringSize := ""
	if size < 1024 {
		stringSize = fmt.Sprintf("%.2f B", float64(size))
	} else if size < 1024*1024 {
		stringSize = fmt.Sprintf("%.2f KB", float64(size/1024.0))
	} else if size < 1024*1024*1024 {
		stringSize = fmt.Sprintf("%.2f MB", float64(size/1024.0/1024.0))
	} else {
		stringSize = fmt.Sprintf("%.2f GB", float64(size/1024.0/1024.0/1024.0))
	}
	return stringSize
}

func readDirectories(files []string) ([]fileStruct, error) {
	filesCnt := len(files)
	filesInfo := make([]fileStruct, filesCnt)
	for i, file := range files {
		fileInfo, err := fileStat(file)
		if err != nil {
			return nil, nil
		}
		filesInfo[i] = *fileInfo
	}
	return filesInfo, nil
}
