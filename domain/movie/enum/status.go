package movieenum

type EspisodeStatus string

const (
	MovieStatus_Upcoming EspisodeStatus = "UPCOMING"
	MovieStatus_OnGoing  EspisodeStatus = "ONGOING"
	MovieStatus_Ended    EspisodeStatus = "ENDED"
)
