package gotenberg

type httpHeader string

const (
	headerOutputFilename = "Gotenberg-Output-Filename"
)

const (
	headerAuthorization = "Authorization"
)

const (
	headerWebhookURL          httpHeader = "Gotenberg-Webhook-Url"
	headerWebhookErrorURL     httpHeader = "Gotenberg-Webhook-Error-Url"
	headerWebhookMethod       httpHeader = "Gotenberg-Webhook-Method"
	headerWebhookErrorMethod  httpHeader = "Gotenberg-Webhook-Error-Method"
	headerWebhookExtraHeaders httpHeader = "Gotenberg-Webhook-Extra-Http-Headers"
)
