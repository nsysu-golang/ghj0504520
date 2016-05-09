package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

func main() {

	fmt.Println("Second, use network")

	/*network operation*/
	respon, err := http.Get("https://goo.gl/mSzFnk")
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	/*network to input stream*/
	defer respon.Body.Close()
	content, err := ioutil.ReadAll(respon.Body)
 	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	/*JSON decode*/
	var hashTable1 =map[string]interface{}{}
	json.Unmarshal(content,&hashTable1)

	for k, v := range hashTable1 {
		var hashTable2= v.(map[string]interface{})
		for i, j := range hashTable2 {
			if(i=="temp_c" || i=="icon_url") {
				fmt.Println(k,"->",i,":",j,"\n\n")
			}
		}
	}

}
