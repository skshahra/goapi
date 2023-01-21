package logs

import "log"



func ERROR(str string,err error){
	if err != nil {
		log.Println("❌🔥 Info :",str,", ERROR: ",err)
	}
}