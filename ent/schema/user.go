package schema

import (
	"database/sql/driver"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Userスキーマを定義
type User struct {
	ent.Schema
}

// フィールドを定義
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.Time("created_at").
			Immutable().
			GoType(CustomTime{}).
			SchemaType(map[string]string{
				"mysql": "datetime",
			}),
		field.Time("updated_at").
			Immutable().
			GoType(CustomTime{}).
			SchemaType(map[string]string{
				"mysql": "datetime",
			}),
	}
}

type CustomTime struct {
	time.Time
}

func (t *CustomTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		t.Time = v
	case []byte:
		tt, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return fmt.Errorf("unable to parse date: %v", err)
		}
		t.Time = tt
	case string:
		tt, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return fmt.Errorf("unable to parse date: %v", err)
		}
		t.Time = tt
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into CustomTime", value)
	}
	return nil
}

func (t CustomTime) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}
