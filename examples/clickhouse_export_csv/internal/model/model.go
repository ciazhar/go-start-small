package model

type Event struct {
	AmpEnabled            bool
	BounceClass           int
	CampaignID            string
	ClickTracking         bool
	CustomerID            string
	DelvMethod            string
	DeviceToken           string
	ErrorCode             string
	EventID               string
	FriendlyFrom          string
	InitialPixel          bool
	InjectionTime         int64
	IPAddress             string
	IPpool                string
	MailboxProvider       string
	MailboxProviderRegion string
	MessageID             string
	MsgFrom               string
	MsgSize               int
	NumRetries            int
	OpenTracking          bool
	RcptMeta              map[string]string
	RcptTags              []string
	RcptTo                string
	RcptHash              string
	RawRcptTo             string
	RcptType              string
	RawReason             string
	Reason                string
	RecipientDomain       string
	RecvMethod            string
	RoutingDomain         string
	ScheduledTime         int64
	SendingDomain         string
	SendingIP             string
	SmsCoding             string
	SmsDst                string
	SmsDstNpi             string
	SmsDstTon             string
	SmsSrc                string
	SmsSrcNpi             string
	SmsSrcTon             string
	SubaccountID          string
	Subject               string
	TemplateID            string
	TemplateVersion       string
	Timestamp             int64
	Transactional         bool
	TransmissionID        string
	Type                  string
}