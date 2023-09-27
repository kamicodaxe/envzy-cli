// Inspired by https://github.com/laurent22/go-sqlkv/blob/master/sqlkv.go
package kvstore

import (
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	PLACEHOLDER_QUESTION_MARK = 1
	PLACEHOLDER_DOLLAR        = 2
)

type SqlKv struct {
	db              *gorm.DB
	tableName       string
	placeholderType int
}

type SqlKvRow struct {
	Name  string
	Value string
}

func New(db *gorm.DB, tableName string) *SqlKv {
	output := new(SqlKv)
	output.db = db
	output.tableName = tableName
	output.placeholderType = PLACEHOLDER_QUESTION_MARK

	// Create the table if it doesn't exist
	err := db.AutoMigrate(&SqlKvRow{})
	if err != nil {
		println("Error on SQlRow migrate")
	}

	return output
}

func (this *SqlKv) placeholder(index int) string {
	if this.placeholderType == PLACEHOLDER_QUESTION_MARK {
		return "?"
	} else {
		return "$" + strconv.Itoa(index)
	}
}

func (this *SqlKv) rowByName(name string) (*SqlKvRow, error) {
	row := new(SqlKvRow)
	if err := this.db.Where("name = ?", name).First(row).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (this *SqlKv) All() []SqlKvRow {
	var rows []SqlKvRow
	if err := this.db.Find(&rows).Error; err != nil {
		return []SqlKvRow{}
	}
	return rows
}

func (this *SqlKv) String(name string) string {
	row, err := this.rowByName(name)
	if err == nil && row == nil {
		return ""
	}
	if err != nil {
		panic(err)
	}
	return row.Value
}

func (this *SqlKv) StringD(name string, defaultValue string) string {
	if !this.HasKey(name) {
		return defaultValue
	}
	return this.String(name)
}

func (this *SqlKv) SetString(name string, value string) error {
	row, err := this.rowByName(name)

	if row == nil && err == nil {
		newRow := SqlKvRow{
			Name:  name,
			Value: value,
		}
		if err := this.db.Create(&newRow).Error; err != nil {
			return err
		}
	} else if err == nil {
		if err := this.db.Model(&row).Where("name = ?", name).Update("Value", value).Error; err != nil {
			return err
		}
	} else {
		return err
	}
	return err
}

func (this *SqlKv) Int(name string) int {
	s := this.String(name)
	if s == "" {
		return 0
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func (this *SqlKv) IntD(name string, defaultValue int) int {
	if !this.HasKey(name) {
		return defaultValue
	}
	return this.Int(name)
}

func (this *SqlKv) SetInt(name string, value int) error {
	s := strconv.Itoa(value)
	return this.SetString(name, s)
}

func (this *SqlKv) SetUInt(name string, value uint) error {
	s := strconv.Itoa(int(value))
	return this.SetString(name, s)
}

func (this *SqlKv) Float(name string) float32 {
	s := this.String(name)
	if s == "" {
		return 0
	}

	o, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}
	return float32(o)
}

func (this *SqlKv) FloatD(name string, defaultValue float32) float32 {
	if !this.HasKey(name) {
		return defaultValue
	}
	return this.Float(name)
}

func (this *SqlKv) SetFloat(name string, value float32) {
	s := strconv.FormatFloat(float64(value), 'g', -1, 32)
	this.SetString(name, s)
}

func (this *SqlKv) Bool(name string) bool {
	s := this.String(name)
	return s == "1" || strings.ToLower(s) == "true"
}

func (this *SqlKv) BoolD(name string, defaultValue bool) bool {
	if !this.HasKey(name) {
		return defaultValue
	}
	return this.Bool(name)
}

func (this *SqlKv) SetBool(name string, value bool) {
	var s string
	if value {
		s = "1"
	} else {
		s = "0"
	}
	this.SetString(name, s)
}

func (this *SqlKv) Time(name string) time.Time {
	s := this.String(name)
	if s == "" {
		return time.Time{}
	}

	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		panic(err)
	}

	return t
}

func (this *SqlKv) TimeD(name string, defaultValue time.Time) time.Time {
	if !this.HasKey(name) {
		return defaultValue
	}
	return this.Time(name)
}

func (this *SqlKv) SetTime(name string, value time.Time) {
	this.SetString(name, value.Format(time.RFC3339Nano))
}

func (this *SqlKv) Del(name string) error {
	if err := this.db.Where("name = ?", name).Delete(&SqlKvRow{}).Error; err != nil {
		return err
	}
	return nil
}

func (this *SqlKv) Clear() {
	if err := this.db.Delete(&SqlKvRow{}).Error; err != nil {
		panic(err)
	}
}

func (this *SqlKv) HasKey(name string) bool {
	_, err := this.rowByName(name)
	if err == nil {
		return true
	}
	return false
}
