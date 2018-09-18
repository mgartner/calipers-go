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

  fmt.Printf("%#v\n", result)
  // calipers.Measurement{Type:"png", Width: 100, Height:60}
}
```

# Supported File Types

* GIF
* PNG
* ~JPEG~ (coming soon)

# Contributing

Feel free to submit a PR!
