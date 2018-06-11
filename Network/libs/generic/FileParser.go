//Could be useful:
//http://www.alangpierce.com/blog/2016/03/17/adventures-in-go-accessing-unexported-functions/


// You can edit this code!
// Click here and start typing.
package CSV

import (
  "*/gabs/gabs.go"
	"io"
  "os"
)

//importFile

//    nodesDecoder = json.NewDecoder(os.Open("*/data/Nodes.json"))
//    jsonParsed, err = gabs.ParseJSONDecoder(nodesDecoder)
//    if err != nil {
//        log.Fatal("Failed to parse: %v", err)
//        return
//    }
//    children, _ := jsonParsed.S("").ChildrenMap()
//    
  


//linksDecoder = json.NewDecoder(os.Open("*/data/Links.json"))




/*

package main

import "fmt"
import "*/gabs/gabs"

func main(){
  jsonParsed, err := gabs.ParseJSON([]byte(`{
    "a":{"_id":"SK1"},
    "b":{"_id":"SK2"},
    "c":{"_id":"SK3"}
  }`))
  for key,child := range jsonParsed.ChildrenMap() {
    fmt.Println(child.Path("_id").Data())
  }
}
*/