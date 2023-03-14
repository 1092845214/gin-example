package global

import (
	"go.uber.org/zap"
	"os"
	"path"
	"time"
)

const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateOnlyFormat = "2006-01-02"
	TimeOnlyFormat = "15:04:05"
)

var (
	exePath, _  = os.Executable()
	ProjectPath = path.Dir(path.Dir(exePath))

	CstZone = time.FixedZone("CST", 8*3600)

	Logger *zap.SugaredLogger
	//DB     *gorm.DB
	//RDS    *redis.Client // 在 global 准备一个 cli, 由 config.redisCli 调用
)
