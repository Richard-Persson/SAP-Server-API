package tools

import (
	"strings"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
)

func RemoveTZ(entries *[]models.TimeEntry)  {
	for i,obj := range *entries{
		before, _, _ := strings.Cut(obj.Date,"T")
		(*entries)[i].Date =  before
	}
}


func RemoveSingleTZ(entry *models.TimeEntry)  {
	before, _, _ := strings.Cut(entry.Date,"T")
	entry.Date = before
}
