package order

const (
	STATUS_PLACED        = "placed"
	STATUS_COMPLETED      = "completed"
	STATUS_CANCELLED = "cancelled"
)

var allowedStatus []string = []string{
	STATUS_PLACED, STATUS_COMPLETED, STATUS_CANCELLED,
}

var AllowedTransitions map[string]map[string]bool = map[string]map[string]bool{
	STATUS_PLACED: {
		STATUS_COMPLETED: true,
		STATUS_CANCELLED:true,
		STATUS_PLACED:true,
	},
	STATUS_COMPLETED: {
		STATUS_CANCELLED: true,
		STATUS_COMPLETED:true,
		STATUS_PLACED:false,
	},
	STATUS_CANCELLED: {
		STATUS_CANCELLED:true,
		STATUS_COMPLETED:false,
		STATUS_PLACED:false,
	},
}
