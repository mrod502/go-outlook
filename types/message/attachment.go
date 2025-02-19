package message

import "time"

type Attachment struct {
	Id                   string `json:"id"`
	IsInline             bool   `json:"isInline"`
	LastModifiedDateTime string `json:"lastModifiedDateTime"`
	Name                 string `json:"name"`
	Size                 int32  `json:"size"`
	ContentType          string `json:"contentType"`
}

type FileAttachment struct {
	Attachment
	ContentBytes    string `json:"contentBytes"`
	ContentId       string `json:"contentId"`
	ContentLocation string `json:"contentLocation"`
}

type ItemAttachment struct {
	Attachment
	ContentBytes []byte      `json:"contentBytes"`
	Item         OutlookItem `json:"item"`
}

type OutlookItem struct {
	Categories           []string  `json:"categories"`
	ChangeKey            string    `json:"changeKey"`
	CreatedDateTime      time.Time `json:"createdDateTime"`
	Id                   string    `json:"id"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
}

type ReferenceAttachment struct {
	Attachment
}
