package model

type Rolebinding struct {
	Id   int    `gorm:"auto_increment;primary_key"`
	UserId string `gorm:"not null"`
	DataCenter   string `gorm:"not null"`
	Cluster      string `gorm:"not null"`
	Workspace    string `gorm:""`
	Namespace     string `gorm:""`
	Role  string `gorm:"not null"`
	//PlanName     string `gorm:"not null"`
	//ServiceName  string `gorm:"not null"`
}

