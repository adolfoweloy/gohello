package main

import (
	"fmt"
	"github.com/adolfoweloy/gohello/basics"
	"github.com/adolfoweloy/gohello/concurrency"
	"github.com/adolfoweloy/gohello/errors"
	"github.com/adolfoweloy/gohello/examples/calculator"
	"github.com/adolfoweloy/gohello/examples/fileio"
	"github.com/adolfoweloy/gohello/examples/httpclient"
	"github.com/adolfoweloy/gohello/examples/json"
	"github.com/adolfoweloy/gohello/oop"
)

func main() {
	fmt.Println("Go Programming Basics Examples")
	fmt.Println("=====================================")
	
	// Basic concepts
	basics.DemonstrateVariables()
	basics.DemonstrateArraysAndSlices()
	basics.DemonstrateMaps()
	basics.DemonstrateConstants()
	basics.DemonstrateIota()
	basics.DemonstrateControlStructures()
	basics.DemonstrateLoops()
	basics.DemonstrateSwitch()
	basics.DemonstrateFunctions()
	basics.DemonstratePointers()
	basics.DemonstrateStructPointers()
	
	// Object-oriented concepts
	oop.DemonstrateStructsAndMethods()
	oop.DemonstrateEmbedding()
	oop.DemonstrateNestedEmbedding()
	oop.DemonstrateInterfaces()
	oop.DemonstrateInterfaceComposition()
	oop.DemonstrateTypeAssertion()
	oop.DemonstrateEmptyInterface()
	
	// Error handling
	errors.DemonstrateBasicErrors()
	errors.DemonstrateCustomErrors()
	errors.DemonstratePanicRecover()
	errors.DemonstrateErrorWrapping()
	
	// Concurrency
	concurrency.DemonstrateGoroutines()
	concurrency.DemonstrateChannels()
	concurrency.DemonstrateSelect()
	concurrency.DemonstrateWaitGroups()
	concurrency.DemonstrateWorkerPool()
	
	// Practical examples
	calculator.DemonstrateCalculator()
	fileio.DemonstrateBasicFileOperations()
	fileio.DemonstrateLogging()
	httpclient.DemonstrateBasicHTTP()
	httpclient.DemonstrateAPIClient()
	jsonexample.DemonstrateBasicJSON()
	jsonexample.DemonstrateFileJSON()
	
	fmt.Println("\n=====================================")
	fmt.Println("All examples completed successfully!")
	fmt.Println("\nTo run individual examples:")
	fmt.Println("- Calculator CLI: go run examples/calculator/main.go")
	fmt.Println("- Or use the interactive features in each package")
}