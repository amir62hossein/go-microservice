package helper


func ErrorHandler(err error) {
	panic(err.Error())
}