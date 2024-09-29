package commonenum

type AvailableStatus string

const (
	AvailableStatus_Draft    AvailableStatus = "DRAFT"
	AvailableStatus_Active   AvailableStatus = "ACTIVE"
	AvailableStatus_Inactive AvailableStatus = "INACTIVE"
)

func (s AvailableStatus) String() string {
	return string(s)
}
