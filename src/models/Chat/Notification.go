package chat

type Notification struct {
	ID int `json:"id" gorm:"primary_key;not null;uniqueIndex;autoIncrement"`
}
