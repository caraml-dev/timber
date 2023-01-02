package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// DeploymentConfig captures the config related to the deployment of Timber Services
type DeploymentConfig struct {
	// EnvironmentType describes the environment Timber services are deployed in
	EnvironmentType string `default:"local"`
	// ProjectName describes the CaraML project Timber services are deployed for
	ProjectName string `default:""`
	// ServiceName describes the CaraML Timber service name
	ServiceName string `default:""`
	// LogLevel captures the selected supported logging level
	LogLevel LogLevel `split_words:"false" default:"INFO"`
	// Maximum no. of go-routines that is allowed
	MaxGoRoutines int `default:"1000"`
}

// LogLevel type is used to capture the supported logging levels
type LogLevel string

const (
	// DebugLevel is used for verbose logs at debug level
	DebugLevel LogLevel = "DEBUG"
	// InfoLevel is used for logs that are info level and higher
	InfoLevel LogLevel = "INFO"
	// WarnLevel is used for logs that are warning level and higher
	WarnLevel LogLevel = "WARN"
	// ErrorLevel is used for logs that are error level and higher
	ErrorLevel LogLevel = "ERROR"
)

// ParseConfig parses configurations from various files into a Viper config
func ParseConfig(spec interface{}, filepaths []string) error {
	v := viper.New()

	err := reflectViperConfig("", spec, v)
	if err != nil {
		return fmt.Errorf("failed to read default config via reflection: %s", err)
	}

	// Load config values from the provided config files
	for _, f := range filepaths {
		v.SetConfigFile(f)
		err := v.MergeInConfig()
		if err != nil {
			return fmt.Errorf("failed to read config from file '%s': %s", f, err)
		}
	}

	// Load config values from environment variables.
	// Nested keys in the config is represented by variable name separated by '::'.
	// For example, DbConfig.Host can be set from environment variable DBCONFIG::HOST.
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "::"))
	v.AutomaticEnv()

	// Unmarshal config values into the config object.
	err = v.Unmarshal(spec)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config values: %s", err)
	}

	return nil
}

// reflectViperConfig recursively traverse over a map and convert nested values into
// environment variables delimited with a full-stop (.)
func reflectViperConfig(prefix string, spec interface{}, v *viper.Viper) error {
	s := reflect.ValueOf(spec)
	s = s.Elem()
	typeOfSpec := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		ftype := typeOfSpec.Field(i)

		viperKey := ftype.Name
		// Nested struct tags
		if prefix != "" {
			viperKey = fmt.Sprintf("%s.%s", prefix, ftype.Name)
		}
		value := ftype.Tag.Get("default")
		v.SetDefault(viperKey, value)
		// Create dynamic map using reflection
		if ftype.Type.Kind() == reflect.Map {
			mapValue := reflect.MakeMapWithSize(ftype.Type, 0)
			v.SetDefault(viperKey, mapValue)
		}

		for f.Kind() == reflect.Ptr {
			if f.IsNil() {
				if f.Type().Elem().Kind() != reflect.Struct {
					// nil pointer to a non-struct: leave it alone
					break
				}
				// nil pointer to struct: create a zero instance
				f.Set(reflect.New(f.Type().Elem()))
			}
			f = f.Elem()
		}

		if f.Kind() == reflect.Struct {
			// Capture information about the config parent prefix
			parentPrefix := prefix
			if !ftype.Anonymous {
				parentPrefix = viperKey
			}

			// Use recursion to resolve nested config
			nestedPtr := f.Addr().Interface()
			err := reflectViperConfig(parentPrefix, nestedPtr, v)
			if err != nil {
				return err
			}
			continue
		}
	}

	return nil
}
