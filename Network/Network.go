//Could be useful:
//http://www.alangpierce.com/blog/2016/03/17/adventures-in-go-accessing-unexported-functions/


// You can edit this code!
// Click here and start typing.
package main

import (
  "*/libs/generic/CSV"
  "*/libs/libs/Geo"
  "*/libs/Node"
  "*/libs/Link"
)

var NodeTable map[string]Node
var LinkTable map[string]Link

func getNodes() map[string]Node {
  csv,err := CSV.new()
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		fmt.Println("Error:", err)
		return
	}
	// automatically call Close() at the end of current method
	defer file.Close()
	// 
	reader := csv.NewReader(file)
	// options are available at:
	// http://golang.org/src/pkg/encoding/csv/reader.go?s=3213:3671#L94
	reader.Comma = ';'
	lineCount := 0
	for {
    // read just one record, but we could ReadAll() as well
		record, err := reader.Read()
    
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
    
    if lineCount == 0 {
      Header header = getHeader(record)
    }
    
    lineCount += 1
  }
}

func main(){
  
  
  
  var node Node
  node.id = 1

  var child1 Link
  child1.us_node = &node
  child1.id = 91

  var child2 Link
  child2.us_node = &node
  child2.id = 92

  node.ds_links[0]=&child1
  node.ds_links[1]=&child2


  fmt.Println("Value: %d",(*node.ds_links[0]).id)
  fmt.Println("Value: %d",(*node.ds_links[1]).id)
}