package files

import (
	"fmt"
	"io"
	"io/ioutil"
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

func fileStat(body *models.FileOperationsReqBody, filePath string) (*models.FileStruct, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	cwd := models.FileStruct{
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

func readDirectories(body *models.FileOperationsReqBody, files []string) ([]models.FileStruct, error) {
	filesCnt := len(files)
	filesInfo := make([]models.FileStruct, filesCnt)
	for i, file := range files {
		fileInfo, err := fileStat(body, path.Join(CONTENT_ROOT_PATH, body.Path, file))
		if err != nil {
			return nil, err
		}
		filesInfo[i] = *fileInfo
	}
	return filesInfo, nil
}

func fileManagerDirectoryContent(body *models.FileOperationsReqBody, filepath string, searchFilterPath string) (*models.FileStruct, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	cwd := models.FileStruct{
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

func checkForDuplicates(directory string, name string, isFile bool) (bool, error) {
	filenames, err := ioutil.ReadDir(directory)
	if err != nil {
		return false, err
	}
	for _, file := range filenames {
		if file.Name() == name {
			if !isFile && file.IsDir() {
				return true, nil
			} else if isFile && !file.IsDir() {
				return true, nil
			} else {
				return false, nil
			}
		}
	}
	return false, nil
}

func fileDetails(filepath string, filterpath string) (*models.FileStruct, error) {
	if info, err := os.Stat(filepath); err != nil {
		return nil, err
	} else {
		cwd := models.FileStruct{
			Name:         path.Base(filepath),
			Size:         getSize(info.Size()),
			IsFile:       !info.IsDir(),
			DateModified: info.ModTime(),
			Type:         path.Ext(filepath),
			Location:     filterpath,
		}
		return &cwd, nil
	}
}

func getFolderSize(directory string) (int, error) {
	size := 0
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		return 0, err
	}

	for _, file := range files {
		if file.IsDir() {
			currSize, err := getFolderSize(path.Join(directory, file.Name()))
			if err != nil {
				return 0, err
			}
			size += currSize
		} else {
			size += int(file.Size())
		}
	}
	return size, nil
}

func checkForFileUpdate(body *models.FileOperationsReqBody, fromPath string, toPath string, item *models.FileDetails, copyName *string) (bool, error) {
	count := 1
	name := item.Name
	if fromPath == toPath {
		if duplicate, err := checkForDuplicates(path.Join(CONTENT_ROOT_PATH, body.TargetPath), name, item.IsFile); err != nil {
			return false, err
		} else {
			if duplicate {
				updateCopyName(path.Join(CONTENT_ROOT_PATH, body.TargetPath), name, count, item.IsFile, copyName)
			}
		}
	} else {
		if len(body.RenameFiles) > 0 && stringInSlice(item.Name, body.RenameFiles) {
			updateCopyName(path.Join(CONTENT_ROOT_PATH, body.TargetPath), name, count, item.IsFile, copyName)
		} else {
			if duplicate, err := checkForDuplicates(path.Join(CONTENT_ROOT_PATH, body.TargetPath), name, item.IsFile); err != nil {
				return false, err
			} else {
				if duplicate {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

func updateCopyName(path string, name string, count int, isFile bool, copyName *string) error {
	subName := ""
	extension := ""
	if isFile {
		extension = name[strings.LastIndex(name, "."):]
		subName = name[:strings.LastIndex(name, ".")]
	}
	if !isFile {
		*copyName = fmt.Sprintf("%s(%d)", name, count)
	} else {
		*copyName = fmt.Sprintf("%s(%d)%s", subName, count, extension)
	}
	if duplicate, err := checkForDuplicates(path, *copyName, isFile); err != nil {
		return err
	} else {
		if duplicate {
			count += 1
			updateCopyName(path, name, count, isFile, copyName)
		}
	}
	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func copyFolder(source string, dest string) error {
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		if err := os.Mkdir(dest, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if files, err := ioutil.ReadDir(source); err != nil {
		return err
	} else {
		for _, file := range files {
			currSource := path.Join(source, file.Name())
			targetPath := path.Join(dest, file.Name())
			if details, err := os.Stat(currSource); err != nil {
				return err
			} else {
				if details.IsDir() {
					if err := copyFolder(currSource, targetPath); err != nil {
						return err
					}
				} else {
					originalFile, err := os.Open(currSource)
					defer originalFile.Close()
					if err != nil {
						return err
					}

					newFile, err := os.Open(targetPath)
					defer newFile.Close()
					if err != nil {
						return err
					}

					if _, err := io.Copy(newFile, originalFile); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func moveFolder(source string, dest string) error {
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		if err := os.Mkdir(dest, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if files, err := ioutil.ReadDir(source); err != nil {
		return err
	} else {
		for _, file := range files {
			currSource := path.Join(source, file.Name())
			targetPath := path.Join(dest, file.Name())
			if details, err := os.Stat(currSource); err != nil {
				return err
			} else {
				if details.IsDir() {
					if err := moveFolder(currSource, targetPath); err != nil {
						return err
					}
				} else {
					if err := os.Rename(currSource, targetPath); err != nil {
						return err
					}
				}
			}

		}
	}
	return nil
}
