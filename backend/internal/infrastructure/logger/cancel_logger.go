package logger

import (
    "context"
    "errors"
    "time"

    glogger "gorm.io/gorm/logger"
)

// CancelFilter は context.Canceled/DeadlineExceeded に伴うGORMのエラーログを抑制します。
type CancelFilter struct {
    glogger.Interface
}

func NewCancelFilter(base glogger.Interface) glogger.Interface {
    return &CancelFilter{Interface: base}
}

// Trace はクエリ結果のログ出力。キャンセル系エラーは出力しない。
func (l *CancelFilter) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
    if err != nil {
        if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
            return
        }
    }
    l.Interface.Trace(ctx, begin, fc, err)
}

