BEA API Client
=======
Client package that facilitates access to economic data via the [US Bureau of Economic Analysis (BEA)](http://www.bea.gov).  Be sure to obtain a free [registration key](http://www.bea.gov/API/signup/index.cfm) prior to accessing the API. 


### Usage
```go
package main

import (
	econ "github.com/openwonk/bea-api"
	"fmt"
)

func main() {
	s := econ.Series{
		"A1BC2D3E-4567-8F9G-0987-H65IJ4K53MN2", // UserId
		"GetData",      // Method
		"STATE",        // GeoFIPS
		"2009",         // Year
		"RegionalData", // DataSetName
		"PCPI_CI",      // KeyCode
	}

	fmt.Println(s.Request())

}
```
<br>
<br>

<hr>
<small>
<strong>OpenWonk &copy; 2015 </strong>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
</small>