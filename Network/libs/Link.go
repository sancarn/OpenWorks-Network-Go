//Could be useful:
//http://www.alangpierce.com/blog/2016/03/17/adventures-in-go-accessing-unexported-functions/


// You can edit this code!
// Click here and start typing.
package Link

import (
  "*/libs/Geo"
)

type Link struct {
  _id       string                  //ID of link
  _type     string                  //Type of link
  _geo      []Point                 //Geometry - Array of Points
  data      *map[string]interface{} //Pointer to data
  us_node   *Node                   //Pointer to node
  ds_node   *Node                   //Pointer to node
}

func New(ID string,TYPE string) Link {
  var l Link
  l._id = ID
  l._type = TYPE
  return l
}