package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SessionID string      `json:"sessionID" binding:"required"`
	Content   interface{} `json:"content" binding:"required" gorm:"json"`
}

// type ContentText string

// func (*ContentText) Value() ([]byte, error) {
// 	return nil, nil
// }

// func (*ContentText) Scan(bytes []byte, value interface{}) error {
// 	return nil
// }

// type ContentIcon uint64

// func (*ContentIcon) Value() ([]byte, error) {
// 	return nil, nil
// }

// func (*ContentIcon) Scan(bytes []byte, value interface{}) error {
// 	return nil
// }

// type ContentTest struct {
// 	Name string
// 	Age  int
// }

// func (c *ContentTest) Value() (driver.Value, error) {
// 	return json.Marshal(c)
// }

// func (c *ContentTest) Scan(value interface{}) error {
// 	buf, err := json.Marshal(c)
// 	if err != nil {
// 		return err
// 	}
// 	return json.Unmarshal(buf, c)
// }
