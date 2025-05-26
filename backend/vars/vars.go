package vars

import (
	"log"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func Get[T string | bool | float64 | int](name string) T {
	typeOfT := reflect.TypeOf(*new(T))

	var result T
	var valueAsStr string

	var envValueAsStr string
	var cliValueAsStr string
	var valuePresentInEnv bool
	var valuePresentInCli bool

	// normalize name
	var nameParts []string

	if strings.ContainsAny(name, "_") {
		nameParts = strings.Split(name, "_")
	} else {
		nameParts = strings.Split(name, "-")
	}

	// check env for variable
	namePartsInCaps := make([]string, len(nameParts))
	for i, np := range nameParts {
		namePartsInCaps[i] = strings.ToUpper(np)
	}
	nameInCaps := strings.Join(namePartsInCaps, "_")

	envValueAsStr, valuePresentInEnv = os.LookupEnv(nameInCaps)
	if valuePresentInEnv {
		valueAsStr = envValueAsStr
	}

	// check cli args for variable
	namePartsInLower := make([]string, len(nameParts))
	for i, np := range nameParts {
		namePartsInLower[i] = strings.ToLower(np)
	}
	nameInLower := strings.Join(namePartsInLower, "-")

	if typeOfT.Kind() == reflect.Bool {
		valueIsTrue := getCliFlag(nameInLower)
		if valueIsTrue {
			valueAsStr = "true"
			valuePresentInCli = true
		} else {
			valuePresentInCli = false
		}
	} else {
		cliValueAsStr, valuePresentInCli = getCliArgument(nameInLower)
		if valuePresentInCli {
			valueAsStr = cliValueAsStr
		}
	}

	if !valuePresentInCli && !valuePresentInEnv {
		log.Fatalf(
			"Could not find variable '%s' or '%s'\nPlease provide it as:\nAn environment variable in .env (%s=<value>)\nOr as a CLI argument (--%s <value>)",
			nameInCaps, nameInLower, nameInCaps, nameInLower,
		)
	}

	switch typeOfT.Kind() {
	case reflect.String:
		result = any(valueAsStr).(T)

	case reflect.Int:
		resultInt64, err := strconv.ParseInt(valueAsStr, 10, 64)
		if err != nil {
			log.Fatalf("Unable to coerce '%s' to type 'int'", valueAsStr)
		}

		result = any(int(resultInt64)).(T)

	case reflect.Float64:
		resultFloat64, err := strconv.ParseFloat(valueAsStr, 64)
		if err != nil {
			log.Fatalf("Unable to coerce '%s' to type 'float64'", valueAsStr)
		}

		result = any(resultFloat64).(T)

	case reflect.Bool:
		resultBool, err := strconv.ParseBool(valueAsStr)
		if err != nil {
			log.Fatalf("Unable to coerce '%s' to type 'bool'", valueAsStr)
		}

		result = any(resultBool).(T)
	default:
		log.Fatalf("Cannot read object of type '%T' from env/cli", *new(T))
	}

	return result
}

func getCliArgument(argument string) (string, bool) {
	args := os.Args[1:]

	argumentNameIndex := slices.Index(args, "--"+argument)
	if argumentNameIndex == -1 {
		// log.Fatalf("CLI argument not found, please provide '--%s <%s>'", argument, argument)
		return "", false
	}

	argumentIndex := argumentNameIndex + 1

	if len(args) < argumentIndex+1 || strings.HasPrefix(args[argumentIndex], "--") {
		log.Fatalf("No argument provided for flag '--%s'", argument)
	}

	return args[argumentIndex], true
}

func getCliFlag(flag string) bool {
	flagExists := slices.Contains(os.Args[1:], "--"+flag)
	return flagExists
}
