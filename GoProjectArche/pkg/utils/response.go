package utils

import (
	"encoding/json"
	"net/http"
)

/*
if we want to send JSON as response
we hv to convert struct data->JSON(ENCODE)
bydefault status code Ok(success)
*/
//                response             data may be any Type so interface{}
func JSONResponse(w http.ResponseWriter, data interface{}) {
	//setting response type is JSON
	w.Header().Set("Content-Type", "application/json")
	//response in JSON format
	json.NewEncoder(w).Encode(data)
}

// bydefault status code Ok(success) +msg
func SuccessResponse(w http.ResponseWriter, msg string) {
	JSONResponse(w, map[string]string{"message": msg})
}

// if failure am getting status code and sending that
func ErrorResponse(w http.ResponseWriter, msg string, code int) {
	//setting failure status code
	w.WriteHeader(code)
	//sending msg+failure status code
	JSONResponse(w, map[string]string{"error": msg})
}
