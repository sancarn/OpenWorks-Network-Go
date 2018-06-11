//Could be useful:
//http://www.alangpierce.com/blog/2016/03/17/adventures-in-go-accessing-unexported-functions/


// You can edit this code!
// Click here and start typing.
package Node

import (
  "*/libs/Geo"
)

type Node struct {
  _id       string                  //ID of link
  _type     string                  //Type of link
  _geo      Geo.Point               //Geometry - Point
  data      *map[string]interface{} //Pointer to data
  ds_links  []*Link                 //Array of Link pointers
  us_links  []*Link                 //Array of Link pointers
}

func New(ID string,TYPE string) Node {
  var n Node
  n._id = ID
  n._type = TYPE
  return n
}