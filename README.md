# calipers-go

Package calipers measures the dimensions of image files
quickly by not loading the entire image into memory.

# Usage

```go
package main

import (
  "github.com/mgartner/calipers-go"
)

func main() {
  result, err := calipers.Measure("path/to/file.png")

  fmt.Println(result) // e.g. Measurement{png, 100, 60}
}
```

# Supported File Types

* PNG
* ~JPEG~ (coming soon)
* ~GIF~ (coming soon)

# Contributing

Feel free to submit a PR!
