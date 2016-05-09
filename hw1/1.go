package main

import "fmt"
import "os"
import "io/ioutil"
import "encoding/json"

func main() {

	fmt.Println("First, use file")

	/*file operation*/
	file, err := os.Open("NSYSU.json")
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	/*file to input stream*/
	content, err := ioutil.ReadAll(file)
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
