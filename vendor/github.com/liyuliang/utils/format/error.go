package format

func ErrorToStr(err interface{}) string {
	if err == nil {
		return ""
	} else {
		return err.(error).Error()
	}
}
