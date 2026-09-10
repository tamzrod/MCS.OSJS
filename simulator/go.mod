module github.com/tamzrod/MCS.OSJS/simulator

go 1.25.0

require (
	github.com/tamzrod/MCS.OSJS/mma2composer v0.0.0
	gopkg.in/yaml.v3 v3.0.1
	mma2 v0.0.0
)

replace github.com/tamzrod/MCS.OSJS/mma2composer => ../mma2composer

replace mma2 => ../MMA2
