package utility

import "reflect"

/*************************************************
 *	specification;
 *	name 			= StructToJsonTagMap
 *	Function 	= change struct to map[string]interface{}
 *	note			= map key is json string that tag json struct
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= data: struct(interface{})
 *  output    = jsontagmap: map[string]interface{}
 *  end of specification;
**************************************************/
func StructToJsonTagMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Tag.Get("json")
		value := elem.Field(i).Interface()
		result[field] = value
	}
	return result
}
