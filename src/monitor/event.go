package monitor

type event struct {
	Identity string
	Timestamp int64
	Etype int32
	WarnLevel string
	Message string
}