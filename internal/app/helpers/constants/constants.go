package constants

const ROOT_PATH_KEY = "ROOT_PATH"
const APPLICATION_PORT_KEY = "APPLICATION_PORT"
const APPLICATION_PORT_DEFAULT = ":8080"

type CorrelationIdKey struct{}

func (c CorrelationIdKey) String() string {
	return "correlation_id"
}
