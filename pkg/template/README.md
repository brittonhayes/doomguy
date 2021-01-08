<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# template

```go
import "github.com/brittonhayes/doomguy/pkg/template"
```

## Index

- [func Usage(command string) string](<#func-usage>)
- [type Templates](<#type-templates>)
  - [func NewTemplates(box *packr.Box) *Templates](<#func-newtemplates>)
  - [func (t *Templates) GameDetails(d *rawg.GameDetailed) string](<#func-templates-gamedetails>)
  - [func (t *Templates) ProcessFile(fileName string, vars interface{}) string](<#func-templates-processfile>)
  - [func (t *Templates) ProcessString(str string, vars interface{}) string](<#func-templates-processstring>)


## func Usage

```go
func Usage(command string) string
```

## type Templates

```go
type Templates struct {
    Box *packr.Box
}
```

### func NewTemplates

```go
func NewTemplates(box *packr.Box) *Templates
```

NewTemplates generates a new instance of the Templates struct with a packr box

### func \(\*Templates\) GameDetails

```go
func (t *Templates) GameDetails(d *rawg.GameDetailed) string
```

### func \(\*Templates\) ProcessFile

```go
func (t *Templates) ProcessFile(fileName string, vars interface{}) string
```

### func \(\*Templates\) ProcessString

```go
func (t *Templates) ProcessString(str string, vars interface{}) string
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)