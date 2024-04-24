module multiversion

go 1.22.0

replace github.com/zmaxic1978/FooModule/100 => github.com/zmaxic1978/FooModule v1.0.0

replace github.com/zmaxic1978/FooModule/110 => github.com/zmaxic1978/FooModule v1.1.0

require (
	github.com/zmaxic1978/FooModule/100 v0.0.0-00010101000000-000000000000
	github.com/zmaxic1978/FooModule/110 v0.0.0-00010101000000-000000000000
)
