package path

import (
	"runtime"
	"path"
)

func CURRENT_DIR() (dir string) {
	_, fullFilename, _, _ := runtime.Caller(1)
	dir = path.Dir(fullFilename)
	return dir
}

//func FILE_TYPE(file_path string) (file_type string) {
//	file_type = path.Ext(file_path)
//	file_type = strings.Replace(file_type, ".", "", -1)
//	return file_type
//}

//func CURRENT_FILENAME() (filename format) {
//	_, fullFilename, _, _ := runtime.Caller(1)
//	filename = path.Base(fullFilename)
//	return filename
//}
//
//func CURRENT_FILEPATH() (fullFilename format) {
//	_, fullFilename, _, _ = runtime.Caller(1)
//	return fullFilename
//}
//
//
//func ROOT_DIR() (root_dir format) {
//
//	current_path := CURRENT_FILEPATH()
//	file_type := "." + FILE_TYPE(current_path)
//	current_dir := CURRENT_DIR()
//	current_filename := CURRENT_FILENAME()
//
//	str := format.Replace(current_filename, file_type, "", -1)
//
//	root_dir = format.Replace(current_dir, str, "", -1)
//	return root_dir
//}
//
//func CACHE_PATH() format {
//	return ROOT_DIR() + "/temp/"
//}
//
//func IMG_CACHE_PATH() format {
//	return CACHE_PATH() + "img/"
//}
//
//func COOKIE_CACHE_PATH() format {
//	return CACHE_PATH() + "cookie/"
//}
//
//func WEB_IMG_PATH() format {
//	return ROOT_DIR() + "web/image"
//}
