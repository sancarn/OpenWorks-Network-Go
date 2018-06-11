//Could be useful:
//http://www.alangpierce.com/blog/2016/03/17/adventures-in-go-accessing-unexported-functions/


// You can edit this code!
// Click here and start typing.
package main

import (
  "encoding/csv"
	"fmt"
	"io"
	"os"
)
type Point struct {
  x float64
  y float64
}

type Node struct {
  _id       string                  //ID of link
  _type     string                  //Type of link
  _geo      Point                   //Geometry - Point
  data      *map[string]interface{} //Pointer to data
  ds_links  []*Link                 //Array of Link pointers
  us_links  []*Link                 //Array of Link pointers
}
type Link struct {
  _id       string                  //ID of link
  _type     string                  //Type of link
  _geo      []Point                 //Geometry - Array of Points
  data      *map[string]interface{} //Pointer to data
  us_node   *Node                   //Pointer to node
  ds_node   *Node                   //Pointer to node
}


func Node_new(ID string,TYPE string,GEO Point) Node {
  var n Node
  n._id = ID
  n._type = TYPE
  n._geo  = GEO
  return n
}

/*
#Alternatively
~~~~~~~~~~~~~~~~~~~~~~~
Node.go
~~~~~~~~~~~~~~~~~~~~~~~
package Node
import "Link"
type Node struct {
  //...
}

//capital letter is important!!
//Functions and Structs starting with a capital letter are exported from the package
func New(...) Node {
  return Node{...}
}

~~~~~~~~~~~~~~~~~~~~~~~
Link.go
~~~~~~~~~~~~~~~~~~~~~~~
package Link
import "Node"
type Link struct {
  //...
}

//capital letter is important!!
func New(...) Node {
  return Node{...}
}

~~~~~~~~~~~~~~~~~~~~~~~
Main.go
~~~~~~~~~~~~~~~~~~~~~~~
package Main

import (
  "Node"
  "Link"
)

func main(){
  n := Node.New(...)
  l := Link.New(...)
}
*/

type Header struct {
  data  interface{}
  m     map[string]int
}
func (h Header) get(s string) interface {
  return h.data[m[s]]
}

func getHeader(data interface) Header {
  //Header to return
  var h Header
  
  //Assign data
  h.data = data
  
  //Prepare map
  for i := 0; i < len(record); i++ {
    h.m[data(i)]=i
  }
  
  //Return header
  return h
}


/*
  # DYNAMIC Node array:
  //Define a new empty array
  a := []Node
  
  //Add a node to the array
  n := Node.new("abcdef")
  a=append(a,n)
*/


func getNodes() map[string]Node {
  file, err := os.Open("data/nodes.csv")
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