# gormt

gormt is a library designed for the [Go programming language](https://go.dev/). It allows to use custom structs as [gorm](https://gorm.io/index.html) JSON type columns through go generics. **Requires Go version 1.18 or higher**.

## How to use

Given the following Postgres table:

``` postgres
CREATE TABLE letters (
	id  UUID NOT NULL PRIMARY KEY,
	info JSON NOT NULL
);
```

where *info* will contains a JSON like:

``` json
{ 
  "sender": "Dvor",
  "recipient": "Magic",
  "content": "Hello World"
}
```

We implement the model as follows:

``` go
type Letter struct {
  ID   string `gorm:"primaryKey"`
  Info gormt.JSON[[]*Info]
}

type Info struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}
```

GORM requires the customized data type to implement the Scanner and Valuer interfaces. This way GORM knowns to how to receive/save it into the database. But using this library you get rid of any complication. You just need to add the label `gormt.JSON[T]` where T could be your custom type and you are good to go!
