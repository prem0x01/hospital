package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func StrPtr(s string) *string {
	return &s
}

func Int32Ptr(i int32) *int32 {
	return &i
}


func TimeToTimestamp(t time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{
		Time:  t,
		Valid: true,
	}
}

func TimestampToTime(ts pgtype.Timestamp) time.Time {
	if !ts.Valid {
		return time.Time{}
	}
	return ts.Time
}

func IintPtrToInt32Ptr(i *int) *int32 {
	if i == nil {
		return nil
	}
	val := int32(*i)
	return &val
}
