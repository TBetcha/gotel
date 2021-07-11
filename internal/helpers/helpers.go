package helpers

import (
	"fmt"
	"github.com/tbetcha/gotel/internal/config"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

// NewHelpers sets up app config for herlpers
func NewHelpers(a *config.AppConfig){
	app = a
}


func ClientError(w http.ResponseWriter, status int){
	app.InfoLog.Println("Client error with status of ", status)
	http.Error(w,http.StatusText(status),status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(),debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)

}
//
//func ReturnFormattedDate(str string, w http.ResponseWriter) time.Time {
//	layout := "2006-01-02"
//	parsedDate, err := time.Parse(layout, str)
//		if err != nil{
//			ServerError(w, err)
//		}
//	return parsedDate
//	}
