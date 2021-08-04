# CA Service Desk WebServices Go
`casdmws-go` is WSDL schema converted into Go code using [wsdl2go](https://github.com/fiorix/wsdl2go) with few additions:
- added mock interface;
- added error struct, methods to make inline pointers to WSDL used data types;
- go mod (that's why it's here);
- has an example if you're new to Go/CASDM Schema.

## Usage
There are no advatages against your own converted WSDL schema but if you're too lazy to make your own, feel free to `go mod` this one.