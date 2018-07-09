package faceAPI

import (
	"FaceServer/lib"
	"fmt"
)

func GetSubject(ID int) {
	result, err := lib.Get(lib.Host+fmt.Sprintf("/subject/%d", ID), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}

func CreateSubject() {

}
