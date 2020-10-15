package util

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type JSONB struct {
	json.RawMessage
}

func (that *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("转换json出错")
	}
	result := JSONB{}
	err := json.Unmarshal(bytes, &result)
	*that = result
	return err
}

func (that JSONB) Value() (driver.Value, error) {
	if len(that.RawMessage) == 0 {
		return nil, nil
	}
	return json.RawMessage(that.RawMessage).MarshalJSON()
}

func (JSONB) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "JSONB"
}
