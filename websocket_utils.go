package slack

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type JSONTimeString string

// String converts the unix timestamp into a string
func (t JSONTimeString) String() string {
	if t == "" {
		return ""
	}
	floatN, err := strconv.ParseFloat(string(t), 64)
	if err != nil {
		log.Panicln(err)
		return ""
	}
	timeStr := int64(floatN)
	tm := time.Unix(int64(timeStr), 0)
	return fmt.Sprintf("\"%s\"", tm.Format("Mon Jan _2"))
}
