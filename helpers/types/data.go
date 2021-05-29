package types

type Map map[string]interface{}

type JsonApi struct {
	code int
	result string
	data Map
}

type StringApi struct {
	code int
	result string
	data string
}