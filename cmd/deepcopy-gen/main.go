package main

import (
	"k8s.io/gengo/args"
	"path/filepath"
	"k8s.io/gengo/examples/deepcopy-gen/generators"
	"github.com/spf13/pflag"
	"github.com/golang/glog"
)

func main() {
	arguments := args.Default()

	// Override defaults.
	arguments.OutputFileBaseName = "deepcopy_generated"
	arguments.GoHeaderFilePath = filepath.Join(args.DefaultSourceTree(), "C:/Users/MyPC/go/src/github.com/xieydd/kubenetes-crd/hack/boilerplate/boilerplate.go.txt")

	// Custom args.
	customArgs := &generators.CustomArgs{}
	pflag.CommandLine.StringSliceVar(&customArgs.BoundingDirs, "bounding-dirs", customArgs.BoundingDirs,
		"Comma-separated list of import paths which bound the types for which deep-copies will be generated.")
	arguments.CustomArgs = customArgs

	// Run it.
	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		glog.Fatalf("Error: %v", err)
	}
	glog.V(2).Info("Completed successfully.")
}
