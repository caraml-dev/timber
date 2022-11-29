package testutils

// Config captures generic information for testing purposes
type Config struct {
	Host   string `default:"localhost"`
	Port   int    `default:"3000"`
	Sentry Sentry
}

// Sentry captures config related to Sentry logging for testing purposes
type Sentry struct {
	URL    string `default:"https://xx.xx.xx"`
	Labels map[string]string
}
