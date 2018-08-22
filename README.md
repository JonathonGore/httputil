# httputil

This package defines useful HTTP Utilities for implementing http handler functions.

## Examples

**main.go**
```
package main

import (
        "net/http"

        "github.com/JonathonGore/httputil"
)

func main() {
        http.HandleFunc("/success", func (w http.ResponseWriter, r *http.Request) {
                httputil.Success(w)
        })

        http.HandleFunc("/error", func (w http.ResponseWriter, r *http.Request) {
                httputil.Error(w, "unable to complete operation", 500)
        })

        http.ListenAndServe(":3000", nil)
}
```

Now if you curl these commands they will automatically be formatted into JSON for you.

```
$ curl http://localhost:3000/success

{"message":"success","code":200}
```

```
$ curl http://localhost:3000/error

{"message":"unable to complete operation","code":500}
```
