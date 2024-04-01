package service


func GetUserDetails(name string, id int32)(map[string]interface{}){
     details := make(map[string]interface{})
	 details["name"] = name
	 details["id"] = id
	 return details

}