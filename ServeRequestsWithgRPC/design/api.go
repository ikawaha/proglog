package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("proglog", func() {
	Title("Write a log")
	Description("Go言語による分散サービス:第Ⅱ部ネットワーク")
	Server("proglog", func() {
		Host("local", func() {
			URI("grpc://0.0.0.0:8080")
		})
	})
})
