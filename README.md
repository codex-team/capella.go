# Go SDK for CodeX Capella

package contains methods to upload images to the Capella Server

## Installation

### Go Get

```golang
go get https://github.com/codex-team/capella.go
```

### Binary download

Binaries for Linux, macOS and Windows are available in the releases section of this repository

## Usage

```golang
import(
	"github.com/codex-team/capella.go"
)

func main() {

    // upload image from URL
    url := "imageUrl"
    response, err := capella.Upload(url)
    
    if err != nil {
       // handle or panic
    }
    
    ....
    
    // upload image from local path
    filepath := "~/images/file.png"
    response, err := capella.UploadFile(filepath)
    
    if err != nil {
       // handle or panic
    }
}
```

response implements capella.Response struct that has
`success`, `message`, `id`, `url`

`success` is `true` when CodeX capella saved the image
`url` - special allocated URL for uploaded image. If `success` is `false` this propery 
takes value of nil
`message` - in case of error you will get a message. 

## Docs

CodeX Capella [documentation](https://github.com/codex-team/capella#readme)

## Contribution

Feel free to aks a question or report a bug or fork and improve a package

