# slurp
Universal API data gathering to the storage solution of your choice (among those implementd :) )

# Usage

## As a module in another project

Download the module with: 
```bash
go get github.com/xgourmandin/slurp
```

Then you must use it like this :
```go
context, err := slurp.NewContextFactory().CreateContextFromConfig(&apiConfiguration)
if err != nil {
  return err
}
engine := slurp.NewSlurpEngine()
engine.SlurpAPI(*context)
```

Where `apiConfiguration` is a `slurp.ApiConfiguration` struct
