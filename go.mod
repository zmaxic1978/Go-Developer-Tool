module multiversion

go 1.22.0

replace github.com/zmaxic1978/FooModule/100 => github.com/zmaxic1978/FooModule v1.0.0

replace github.com/zmaxic1978/FooModule/110 => github.com/zmaxic1978/FooModule v1.1.0

replace github.com/zmaxic1978/trans_mod/100 => github.com/zmaxic1978/trans_mod v1.0.0

replace github.com/zmaxic1978/trans_mod/110 => github.com/zmaxic1978/trans_mod v1.1.0

replace github.com/zmaxic1978/trans_mod/v2/200 => github.com/zmaxic1978/trans_mod/v2 v2.0.0

replace github.com/zmaxic1978/trans_mod/v2/210 => github.com/zmaxic1978/trans_mod/v2 v2.1.0

require (
	github.com/zmaxic1978/FooModule/100 v0.0.0-00010101000000-000000000000
	github.com/zmaxic1978/FooModule/110 v0.0.0-00010101000000-000000000000
    github.com/zmaxic1978/trans_mod/100 v0.0.0-00010101000000-000000000000
	github.com/zmaxic1978/trans_mod/110 v0.0.0-00010101000000-000000000000
    github.com/zmaxic1978/trans_mod/v2/200 v0.0.0-00010101000000-000000000000
    github.com/zmaxic1978/trans_mod/v2/210 v0.0.0-00010101000000-000000000000
)