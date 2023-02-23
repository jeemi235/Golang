module mymodule

go 1.19

// it could be a file anywhere on your computer being run with go run. The first real benefit of Go modules is being able to add dependencies to your project in any directory
//and not just the GOPATH directory structure. You can also add packages to your module. In the next section, you will expand your module by creating an additional package within it.

require github.com/spf13/cobra v1.6.1

require (
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
