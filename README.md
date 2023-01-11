# CA Service Desk WebServices Go
`casdmws-go` is WSDL schema converted into Go code using [wsdl2go](https://github.com/fiorix/wsdl2go) with few additions:
- easy access login methods for both SOAP/REST;
- http client wrapped into SDM-specific REST client w/ support of attachments;
- mock interface;
- error struct and some tools like inline pointers to WSDL used data types;
- go mod (that's why it's here);
- has an example if you're new to Go/CASDM Schema.

## Why
Regarding to my experience, integrations with ITSM are mainly done by system-administrators and other kind of non-dev staff, who starts their trip by googling how-to examples.
My goal is to make this process easier and give the option to start writing code and call end-system methods without dealing with file-system, encryption, and other kind of preparation and system-related specifics.

## Usage
There are no advatages against your own converted WSDL schema but if you're too lazy to make your own, feel free to `go mod` this one.

## TODOs
- need to add and test REST Mock-related stuff;
- need to organize loggers, they're bit chaotic now.

## Licence
[MIT](https://opensource.org/licenses/MIT)