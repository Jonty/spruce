## simpleyaml

a Go package to interact with arbitrary YAML, similar as [go-simplejson](https://github.com/bitly/go-simplejson).

[![GoDoc](https://godoc.org/github.com/smallfish/simpleyaml?status.svg)](http://godoc.org/github.com/smallfish/simpleyaml)

#### INSTALL
    
```bash
go get -u gopkg.in/yaml.v2 # required
go get -u github.com/smallfish/simpleyaml
```

#### EXAMPLE

```go
var data = []byte(`
name: smallfish
age: 99
bool: true
emails:
   - xxx@xx.com
   - yyy@yy.com
bb:
    cc:
        dd:
            - 111
            - 222
            - 333
        ee: aaa
`)

y, err := NewYaml(data)
if err != nil {
        // ERROR
}

name, err := y.Get("name").String()
if err != nil {
        // ERROR
}
fmt.Println("name:", name)

// y.Get("age").Int()
// y.Get("bool").Bool()
// y.Get("bb").Get("cc").Get("ee").String()
// y.Get("bb").Get("cc").Get("ee").GetIndex(1).Int()
// y.GetPath("bb", "cc", "ee").String()
```

__END__
