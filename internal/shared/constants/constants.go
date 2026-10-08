package constants

const ROOT_PATH_KEY = "ROOT_PATH"
const APPLICATION_PORT_KEY = "APPLICATION_PORT"
const APPLICATION_PORT_DEFAULT = ":8080"

const LOGGER_TYPE_KEY = "LOGGER_TYPE"
const LOGGER_TYPE_DEFAULT = "json"
const LOGGER_TYPE_JSON = "json"
const LOGGER_TYPE_TEXT = "text"
const LOGGER_TYPE_FILE = "file"

type CorrelationIdKey struct{}

func (c CorrelationIdKey) String() string {
	return "correlation_id"
}
