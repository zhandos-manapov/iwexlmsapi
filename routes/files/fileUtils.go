package files

import (
	"fmt"
	"iwexlmsapi/models"
	"os"
	"path"
	"strings"
)

func InitConstants() {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	CONTENT_ROOT_PATH = path.Join(rootPath, "Files")
	FILESYSTEM = os.DirFS(CONTENT_ROOT_PATH)
}

func fileStat(body *models.FileOperationsReqBody, filePath string) (*fileStruct, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	cwd := fileStruct{
		Name:         path.Base(filePath),
		Size:         getSize(fileInfo.Size()),
		IsFile:       !fileInfo.IsDir(),
		DateModified: fileInfo.ModTime(),
		Type:         path.Ext(filePath),
		FilterPath:   getRelativePath(CONTENT_ROOT_PATH, CONTENT_ROOT_PATH+body.Path),
		Permission:   nil,
		HasChild:     fileInfo.IsDir(),
	}
	return &cwd, nil
}

func getRelativePath(rootDirectory string, fullPath string) string {
	if strings.HasSuffix(rootDirectory, "/") {
		if strings.Contains(fullPath, rootDirectory) {
			return fullPath[len(rootDirectory)-1:]
		}
		return ""
	} else if strings.Contains(fullPath, rootDirectory+"/") {
		return "/" + fullPath[len(rootDirectory)+1:]
	} else {
		return ""
	}
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

func readDirectories(body *models.FileOperationsReqBody, files []string) ([]fileStruct, error) {
	filesCnt := len(files)
	filesInfo := make([]fileStruct, filesCnt)
	for i, file := range files {
		p := path.Join(CONTENT_ROOT_PATH, body.Path, file)
		fmt.Println(p)
		fileInfo, err := fileStat(body, p)
		if err != nil {
			return nil, err
		}
		filesInfo[i] = *fileInfo
	}
	return filesInfo, nil
}

func fileManagerDirectoryContent(body *models.FileOperationsReqBody, filepath string, searchFilterPath string) (*fileStruct, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	cwd := fileStruct{
		Name:         path.Base(filepath),
		Size:         getSize(fileInfo.Size()),
		IsFile:       !fileInfo.IsDir(),
		DateModified: fileInfo.ModTime(),
		Type:         path.Ext(filepath),
		FilterPath:   "",
		Permission:   nil,
		HasChild:     fileInfo.IsDir(),
	}
	if searchFilterPath != "" {
		cwd.FilterPath = searchFilterPath
	} else {
		if len(body.Data) > 0 {
			cwd.FilterPath = body.Path
		} else {
			cwd.FilterPath = ""
		}
	}
	return &cwd, nil
}
