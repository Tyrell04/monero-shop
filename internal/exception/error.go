package exception

// PanicLogging panics if error is not nil
func PanicLogging(err interface{}) {
	if err != nil {
		panic(err)
	}
}
