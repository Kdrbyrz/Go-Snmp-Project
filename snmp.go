package main
import (
	"fmt"
	"net/http"
	"log"
	"strconv"
	"strings"
	g "github.com/soniah/gosnmp"
)

func main() {
	g.Default.Target = "127.0.0.1"
	err := g.Default.Connect()
	if err != nil {
	    log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	oids := []string{"1.3.6.1.2.1.1.4.0", "1.3.6.1.2.1.1.7.0"}
	result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
	    log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
	    fmt.Printf("%d: oid: %s ", i, variable.Name)

	    // the Value of each variable returned by Get() implements
	    // interface{}. You could do a type switch...
	    switch variable.Type {
	    case g.OctetString:
	        bytes := variable.Value.([]byte)
	        fmt.Printf("string: %s\n", string(bytes))
	        resp, err := http.NewRequest("POST","URL",  strings.NewReader(string(bytes)) )
	        if err != nil {
				log.Fatalf("Post() err: %v", err)
			}else{
				fmt.Printf("%s sent successful\n",resp)
				defer resp.Body.Close()
			}

	    default:
	        // Get int data , convert to string and post 
		    typeInt := variable.Value.(int)
		    intToStr := strconv.Itoa(typeInt)
	        fmt.Printf("number: %s\n",intToStr)
	        resp, err := http.NewRequest("POST","URL",  strings.NewReader(intToStr))
	        if err != nil {
				log.Fatalf("Post() err: %v", err)
			}else{
				fmt.Printf("%s sent successful\n",resp)
				defer resp.Body.Close()
			}
	    }
	}
}
	
