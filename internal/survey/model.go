package survey

import "time"

type Survey struct {
	ID            string    `json:"id" db:"id"`
	Description   string    `json:"description" db:"description"`
	FromPhNum     string    `json:"fromPhNum" db:"from_ph_num"`
	ToPhNum       string    `json:"toPhNum" db:"to_ph_num`
	TwilioFlowSID string    `json:"twilioFlowSid" db:"twilio_flow_sid"`
	FlowParams    string    `json:"flowParams" db:"flow_params"`
	CreatedTime   time.Time `json:"createdTime" db:"created_time"`
	UpdatedTime   time.Time `json:"updatedTime" db:"updated_time"`
	DeletedTime   time.Time `json:"deletedTime" db:"deleted_time"`
}
