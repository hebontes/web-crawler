package main
import (
  "testing"
  "reflect"
)



func TestGetURLsFromHTML(t *testing.T){

  tests := []struct{
    name string
    inputURL string
    inputBody string
    expected []string
  }{
    {
      name:     "absolute and relative URLs",
      inputURL: "https://blog.boot.dev",
      inputBody: `
      <html>
      <body>
      <a href="/path/one">
      <span>Boot.dev</span>
      </a>
      <a href="https://other.com/path/one">
      <span>Boot.dev</span>
      </a>
      </body>
      </html>
      `,
      expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
    },
    {
			name:     "absolute URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="https://blog.boot.dev">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name:     "relative URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
  } 


  for i, tc := range tests{
    t.Run(tc.name, func (t *testing.T){
      actual, err :=getURLsFromHTML(tc.inputBody,tc.inputURL)  
      if err != nil{
        t.Errorf("err != nill" )
        return
      }
      
      if !reflect.DeepEqual(tc.expected, actual) {
        t.Errorf("Test %v - %s Fail: Expected: %v, Actual: %v", i, tc.name, tc.expected, actual)
        return
      }



    })
  }
}
