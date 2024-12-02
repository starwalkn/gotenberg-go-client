package gotenberg

type httpHeader string

const (
	headerOutputFilename httpHeader = "Gotenberg-Output-Filename"
	headerTrace          httpHeader = "Gotenberg-Trace"
)

const (
	headerAuthorization httpHeader = "Authorization"
)

const (
	headerWebhookURL          httpHeader = "Gotenberg-Webhook-Url"
	headerWebhookErrorURL     httpHeader = "Gotenberg-Webhook-Error-Url"
	headerWebhookMethod       httpHeader = "Gotenberg-Webhook-Method"
	headerWebhookErrorMethod  httpHeader = "Gotenberg-Webhook-Error-Method"
	headerWebhookExtraHeaders httpHeader = "Gotenberg-Webhook-Extra-Http-Headers"
)
