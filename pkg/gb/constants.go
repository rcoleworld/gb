package gb

const (
    RequestsUsage = "the number of requests to send"
    ConcurrencyUsage = "the number of concurrent requests"
    RequestMethodUsage = "the http method to use. allowed: GET, POST" 
    RequestBodyUsage = "the http request body" 
)

const (
    RequestsFlag = "requests"
    ConcurrencyFlag = "concurrency"
    RequestMethodFlag = "request-method"
    RequestBodyFlag = "body"
)

const (
    ConcurrencyExceedsRequestsWarning = "warning, number of concurrent requests exceeds total number of requests. truncating number of concurrent requests to the total number of requests"
)

