package main

import (
	"fmt"
	cga "github.com/guaychou/corona-api"
	"strings"
)

func corona(message string)string{
	split:=strings.Split(message," ")
   if len(split)==1 {
		country:=message
		result:=cga.GetCorona(country)
		if result.Recovered.Value==-1 || result.Deaths.Value==-1 || result.Confirmed.Value==-1{
			return "Country not found."
		}else {
			deathrate:=fmt.Sprintf("%.2f",result.CaseFatalityRate)
			recoveryrate:=fmt.Sprintf("%.2f",result.CaseRecoveryRate)
			return "DR: "+deathrate+"% RR: "+recoveryrate+"%"
		}
	}
	return "Error: Something goes wrong"
}