package placeholder

import (
    "github.com/tsiparinda/platform/logging"
    "time"
    "fmt"
)

type DayHandler struct {
    logging.Logger
}

func (dh DayHandler) GetDay() string {
    return  fmt.Sprintf("Day: %v", time.Now().Day())
}
