package transform

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"reflect"
	"sort"
	"strings"
	// "fmt"

	"github.com/nsf/jsondiff"
)

const (
	changePattern = "changed"
)

var topLevelPropertyIndicators = []string{"    \"", "    {\"changed\":[\""}

// Diff compares 2 JSON strings and
// returns the changed/added/removed
// top level json objects
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
		if yes, str := isNewTopLevelProp(stringRead); yes {
			currentProperty = str
		}
		if strings.Contains(stringRead, changePattern) {
			diffMap[currentProperty] = true
		}
	}
	return mapKeysToSlice(reflect.ValueOf(diffMap).MapKeys())
}

func isNewTopLevelProp(str string) (bool, string) {
	for _, indicator := range topLevelPropertyIndicators {
		if strings.HasPrefix(str, indicator) {
			return true, (strings.Split(strings.TrimPrefix(str, indicator), "\""))[0]
		}
	}
	return false, ""
}

func mapKeysToSlice(keys []reflect.Value) []string {
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	sort.Strings(strkeys)
	return strkeys
}

// ObjectDiff compares specified object in 2 json strings
// returns the top level differences between the found objects
func ObjectDiff(a, b, objectAPath, objectBPath string) []string {
	objTree := strings.Split(objectAPath, ".")
	objA := getJSONObject(a, objTree)

	objTree = strings.Split(objectBPath, ".")
	objB := getJSONObject(b, objTree)

	return Diff(objA, objB)
}

func getJSONObject(a string, objTree []string) string {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(a), &result); err != nil {
		log.Printf("error getting specified json object %v", err)
		return ""
	}
	for k, v := range result {
		if k == objTree[0] {
			obj, _ := json.Marshal(v)
			if len(objTree) == 1 {
				return string(obj)
			}
			return getJSONObject(string(obj), objTree[1:])
		}
	}
	return ""
}
