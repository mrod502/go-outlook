package message

import "time"

type Message struct {
	BccRecipients                 []interface{} `json:"BccRecipients"` // [{"@odata.type": "microsoft.graph.recipient"}],
	Body                          interface{}   `json:"Body"`          //{"@odata.type": "microsoft.graph.itemBody"},
	BodyPreview                   string        `json:"BodyPreview"`
	Categories                    []string      `json:"categories"`
	CcRecipients                  []interface{} `json:"CcRecipients"` //[{"@odata.type": "microsoft.graph.recipient"}],
	ChangeKey                     string        `json:"ChangeKey"`
	ConversationId                string        `json:"ConversationId"`
	ConversationIndex             []byte        `json:"ConversationIndex"`
	CreatedDateTime               time.Time     `json:"CreatedDateTime"`
	Flag                          interface{}   `json:"Flag"` //{"@odata.type": "microsoft.graph.followupFlag"},
	From                          interface{}   `json:"From"` // {"@odata.type": "microsoft.graph.recipient"},
	HasAttachments                bool          `json:"HasAttachments"`
	Id                            string        `json:"Id"`
	Importance                    string        `json:"Importance"`
	InferenceClassification       string        `json:"InferenceClassification"`
	InternetMessageHeaders        []interface{} `json:"InternetMessageHeaders"` // [{"@odata.type": "microsoft.graph.internetMessageHeader"}],
	InternetMessageId             string        `json:"InternetMessageId"`
	IsDeliveryReceiptRequested    bool          `json:"IsDeliveryReceiptRequested"`
	IsDraft                       bool          `json:"IsDraft"`
	IsRead                        bool          `json:"IsRead"`
	IsReadReceiptRequested        bool          `json:"IsReadReceiptRequested"`
	LastModifiedDateTime          time.Time     `json:"LastModifiedDateTime"`
	ParentFolderId                string        `json:"ParentFolderId"`
	ReceivedDateTime              time.Time     `json:"ReceivedDateTime"`
	ReplyTo                       []interface{} `json:"ReplyTo"` // [{"@odata.type": "microsoft.graph.recipient"}],
	Sender                        interface{}   `json:"Sender"`  //{"@odata.type": "microsoft.graph.recipient"},
	SentDateTime                  time.Time     `json:"SentDateTime"`
	Subject                       string        `json:"Subject"`
	ToRecipients                  []interface{} `json:"ToRecipients"` //[{"@odata.type": "microsoft.graph.recipient"}],
	UniqueBody                    interface{}   `json:"UniqueBody"`   // {"@odata.type": "microsoft.graph.itemBody"},
	WebLink                       string        `json:"WebLink"`
	Attachments                   []Attachment  `json:"Attachments"`                   // [{"@odata.type": "microsoft.graph.attachment"}],
	Extensions                    []interface{} `json:"Extensions"`                    // [{"@odata.type": "microsoft.graph.extension"}],
	MultiValueExtendedProperties  []interface{} `json:"MultiValueExtendedProperties"`  // [{"@odata.type": "microsoft.graph.multiValueLegacyExtendedProperty"}],
	SingleValueExtendedProperties []interface{} `json:"SingleValueExtendedProperties"` //[{"@odata.type": "microsoft.graph.singleValueLegacyExtendedProperty"}]
}
