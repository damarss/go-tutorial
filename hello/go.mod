module com.example/hello

go 1.20

require (
	com.example/greetings v0.0.0-00010101000000-000000000000
	golang.org/x/example v0.0.0-20230515183114-5bec75697667
)

replace com.example/greetings => ../greetings
