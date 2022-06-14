module example.com/main

go 1.18

replace example.com/nodequeue => ../nodequeue

replace example.com/slicequeue => ../slicequeue

require example.com/nodequeue v0.0.0-00010101000000-000000000000 // indirect

require example.com/slicequeue v0.0.0-00010101000000-000000000000 // indirect
