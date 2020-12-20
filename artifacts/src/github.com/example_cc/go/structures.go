

package main

import (
	//"bytes"
	// "encoding/json"
	"time"
)



type Document struct {

ObjectType          string   		`json:"Document"`
Fieldname		    string   		`json:"fieldname"`
Originalname		string   		`json:"originalname"`
Encoding			string   		`json:"encoding"`
Mimetype			string   		`json:"mimetype"`
Destination			string   		`json:"destination"`
Filename			string   		`json:"filename"`
Path				string   		`json:"path"`
Size				string   		`json:"size"`
Txid				string   		`json:"txid"`
FileHash			string			`json:"fileHash"`
DocumentID			string    		`json:documentId`
UploadedTime        time.Time  		`json:"uploadedTime"`

}