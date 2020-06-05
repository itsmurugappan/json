package pkg

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/nsf/jsondiff"
)

const (
	NEW_PROPERTY_INDICATOR = "    \""
	CHANGE_PATTERN         = "changed"
)

func Diff(a, b string) []string {
	opts := jsondiff.Options{
		Added:            jsondiff.Tag{Begin: "{\"changed\":[", End: "]}"},
		Removed:          jsondiff.Tag{Begin: "{\"changed\":[", End: "]}"},
		Changed:          jsondiff.Tag{Begin: "{\"changed\":[", End: "]}"},
		ChangedSeparator: ", ",
		Indent:           "    ",
	}

	result, comparedStr := jsondiff.Compare([]byte(a), []byte(b), &opts)

	if !(result == jsondiff.NoMatch || result == jsondiff.SupersetMatch) {
		return nil
	}

	reader := bufio.NewReader(bytes.NewReader([]byte(comparedStr)))
	diffMap := make(map[string]bool)
	var currentProperty string
	for {
		stringRead, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error finding difference in json strings %v", err)
		}
		if strings.HasPrefix(stringRead, NEW_PROPERTY_INDICATOR) {
			tmp := strings.Split(strings.TrimPrefix(stringRead, NEW_PROPERTY_INDICATOR), "\"")
			currentProperty = tmp[0]
		}
		if strings.Contains(stringRead, CHANGE_PATTERN) {
			diffMap[currentProperty] = true
		}
	}
	return mapToStringArray(reflect.ValueOf(diffMap).MapKeys())
}

func mapToStringArray(keys []reflect.Value) []string {
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	sort.Strings(strkeys)
	return strkeys
}
