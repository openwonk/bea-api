package bea

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	// "strconv"
	// "strings"
	// "bytes"
	// "encoding/json"
)

// MetaData Methods
// =================
// - GetDataSetList
// - GetParameterList
// - GetParameterValues

// MetaData Method
// =================
// - GetData

// Request Struct
// =================
// UserID=B1AF3E5C-5710-4B5C-8675-B39EA3D54FC6
// method=GetData
// datasetname=RegionalData
// KeyCode=PCPI_CI
// GeoFIPS=STATE
// Year=2009
// ResultFormat=JSON

type Series struct {
	UserId, Method, GeoFIPS    string
	Year, DataSetName, KeyCode string
}

func (s *Series) Request() string {

	u, err := url.Parse("http://www.bea.gov/api/data")
	check(err)

	q := u.Query()
	q.Set("UserId", s.UserId)
	q.Set("Method", s.Method)
	q.Set("GeoFIPS", s.GeoFIPS)
	q.Set("Year", s.Year)
	q.Set("DataSetName", s.DataSetName)
	q.Set("KeyCode", s.KeyCode)
	q.Set("ResultForm", "JSON")
	u.RawQuery = q.Encode()
	fmt.Println("Query: ", u)

	resp, err := http.Get(u.String())
	check(err)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	output := string(body)
	fmt.Println("\n", output)
	return output
}

func check(err error) {
	if err != nil {
		panic(err)
	}

}
