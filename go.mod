module github.com/zdunecki/k8s-publishing-bot-example-source

go 1.19

replace github.com/zdunecki/k8s-publishing-bot-example-source-client => ./pkg/client

require github.com/zdunecki/k8s-publishing-bot-example-source-client v0.0.0-00010101000000-000000000000 // indirect
