package util

var (
	// BaseSuccessResponse is the base of every successful API call + data
	BaseSuccessResponse = Message(true, "Success")
	// BaseFailedResponse is the base of every failed API call
	BaseFailedResponse = Message(false, "Something is wrong")
)
