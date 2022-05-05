package hns

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"google.golang.org/grpc/examples/helloworld/greeter_server/db"
)

func StoryToMap(s SummedType) map[string][]map[string]interface{} {
	var resultFromDb []db.Item
	for _, v := range s.Items {
		resultFromDb = append(resultFromDb, db.Item{
			Title: v.Title,
			Score: v.Score,
		})
	}

	var topStoriesMap = make(map[string][]map[string]interface{})
	var sliceOfMaps = make([]map[string]interface{}, 0)

	for _, v := range resultFromDb {
		elem := reflect.ValueOf(&v).Elem()
		relType := elem.Type()

		var myMap = make(map[string]interface{})

		for i := 0; i < relType.NumField(); i++ {
			myMap[relType.Field(i).Name] = elem.Field(i).Interface()
		}
		delete(myMap, "DateStamp")
		delete(myMap, "ID")

		sliceOfMaps = append(sliceOfMaps, myMap)
		topStoriesMap["top_stories"] = sliceOfMaps
	}

	return topStoriesMap
}

func HandleUserJSONResponse(stories SummedType) http.Handler {
	filteredStories := StoryToMap(stories)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.MarshalIndent(filteredStories, "", "   ")
		if err != nil {
			log.Println(err)
		}

		w.Write([]byte(string(b)))
	})
}
