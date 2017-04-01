package main

import "fmt"
import "flag"
import "time"
import "math/rand"
import "github.com/icrowley/fake"

// note, that variables are pointers
var strFlag = flag.String("word", "def_val", "Description")

//number of lines(nol)
var nol = flag.Int("nol", 1, "Number of lines")
var boolFlag = flag.Bool("bool", false, "Description of flag")

func init() {
	// example with short version for long flag
	flag.StringVar(strFlag, "s", "deff_val", "Description")

	generate(5)
}

func generate(nol int) {
	response := [...]string{"200", "404", "500", "301"}
	verb := [...]string{"GET", "POST", "DELETE", "PUT"}
	resources := [...]string{"/list", "/wp-content", "/wp-admin", "/explore", "/search/tag/list", "/app/main/posts", "/posts/posts/explore", "/apps/cart.jsp?appID="}
	ip := fake.IPv4()
	ua := fake.UserAgent()
	dt := time.Now().Format("02/Jan/2006:15:04:05:-0700")
	rand.Seed(time.Now().Unix())
	cresp := response[rand.Int()%len(response)]
	cverb := verb[rand.Int()%len(verb)]
	cresource := resources[rand.Int()%len(resources)]
	cbyte := rand.Int() % 10000
	fmt.Println(fake.UserAgent())
	// f.write('%s - - [%s %s] "%s %s HTTP/1.0" %s %s "%s" "%s"\n' % (ip,dt,tz,vrb,uri,resp,byt,referer,useragent))

	fmt.Printf("%s - - [%s] \"%s %s HTTP/1.0\" %s %s %d\n", ip, dt, cverb, cresource, cresp, ua, cbyte)
}

func main() {
	flag.Parse()
	// println("str flag: ", *strFlag)
	// println("bool flag: ", *boolFlag)
}
