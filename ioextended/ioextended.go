package ioextended


import(
	"net/http"
	"fmt"
	"time"

	"appengine"
	"appengine/datastore"
)

type Dato struct {
	Dato string
	Created time.Time
}

func init() {
    http.HandleFunc("/", rootWS)
    http.HandleFunc("/setdata", setdata)
    http.HandleFunc("/getdata", getdata)
    }

func rootWS(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"hola mundo en la io extended")
}

func setdata(w http.ResponseWriter, r *http.Request) {
	var data string
	data = r.FormValue("data")

	if data == ""{
		fmt.Fprintf(w, "Error")
	}else{
		fmt.Fprintf(w,"para poner datos: '%s'",data)
		dataobj := new(Dato)
		dataobj.Created = time.Now()
		dataobj.Dato = data

		c := appengine.NewContext(r)
		k := datastore.NewIncompleteKey(c,"Dato",nil)
		datastore.Put(c,k,dataobj)
	}

	

	
}

func getdata(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Dato").Order("-Created").Limit(1)
	var datos []Dato
	q.GetAll(c,&datos)

	fmt.Fprintf(w,datos[0].Dato)
}