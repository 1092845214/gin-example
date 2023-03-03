package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateOnlyFormat = "2006-01-02"
	TimeOnlyFormat = "15:04:05"
)

var (
	CstZone = time.FixedZone("CST", 8*3600)

	Logger *zap.SugaredLogger
	DB     *gorm.DB
)
