package movieenum

type ProgressionStatus string

const (
	ProgressionStatus_Upcoming ProgressionStatus = "UPCOMING"
	ProgressionStatus_OnGoing  ProgressionStatus = "ONGOING"
	ProgressionStatus_Ended    ProgressionStatus = "ENDED"
)

func (s ProgressionStatus) String() string {
	return string(s)
}
