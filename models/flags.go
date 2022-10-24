package models

import "fmt"

type Flags struct {
	IP   string
	Port string
}

func (f *Flags) GetApplicationURL() (*string, *ErrorDetail) {
	f, err := GetFlags()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s:%s", f.IP, f.Port)
	return &url, nil
}

var flagsObj *Flags

func NewFlags(ip, port string) *Flags {
	if flagsObj == nil {
		flagsObj = &Flags{
			IP:   ip,
			Port: port,
		}
	}
	return flagsObj
}

func GetFlags() (*Flags, *ErrorDetail) {
	if flagsObj == nil {
		return nil, &ErrorDetail{
			ErrorType:    ErrorTypeFatal,
			ErrorMessage: "Flags not set",
		}
	}

	return flagsObj, nil
}
