package lib

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Data Data
}

type Data struct {
	Status    string
	Track     int
	Timestamp int64
	Face      Face
	Person    Person
	Quality   float64
}

type Face struct {
	Image string
}

type Person struct {
	FeatureID  int `json:"feature_id"`
	Confidence float64
	Tag        string
	ID         string `json:"id"`
}

type Tag struct {
	Description   string
	Remark        string
	SubjectType   int
	Name          string
	Avatar        string
	OriginPhotoID string
	Department    string
}

func ProcessBytesToFile(bytes []byte) {
	result := &Result{}
	errJSON := json.Unmarshal(bytes, result)
	if errJSON != nil {
		panic(errJSON)
	}
	fmt.Println(result.Data.Status)
	if result.Data.Status == "recognized" {
		fmt.Println("----- Tag:", result.Data.Person.Tag, result.Data.Person.DecodedTag())
		SaveBase64Image(result.Data.Face.Image, result.Data.Person.DecodedTag().Name)
	}
}

func (person *Person) DecodedTag() *Tag {
	tag := &Tag{}
	err := json.Unmarshal([]byte(person.Tag), tag)
	CheckError(err)
	return tag
}
