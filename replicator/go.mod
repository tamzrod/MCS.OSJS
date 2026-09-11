module github.com/tamzrod/MCS.OSJS/replicator

go 1.25.0

require (
	github.com/tamzrod/MCS.OSJS/mma2composer v0.0.0
	github.com/tamzrod/MCS.OSJS/mma2raw v0.0.0
	gopkg.in/yaml.v3 v3.0.1
	mma2 v0.0.0 // indirect
)

replace github.com/tamzrod/MCS.OSJS/mma2composer => ../mma2composer

replace github.com/tamzrod/MCS.OSJS/mma2raw => ../mma2raw

replace mma2 => ../MMA2
