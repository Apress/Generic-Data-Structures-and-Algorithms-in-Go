module example.com/main

go 1.18

replace example.com/nodestack => ../nodestack

replace example.com/slicestack => ../slicestack

require example.com/nodestack v0.0.0-00010101000000-000000000000

require example.com/slicestack v0.0.0-00010101000000-000000000000
