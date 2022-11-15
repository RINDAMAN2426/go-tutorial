module example.com/hello

go 1.19

require example.com/greetings v0.0.0-00010101000000-000000000000

require golang.org/x/example v0.0.0-20220412213650-2e68773dfca0 // indirect

replace example.com/greetings => ../greetings
