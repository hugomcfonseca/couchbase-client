package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"

	couchbase "github.com/couchbase/go-couchbase"
)

var (
	cbAction     = flag.String("action", "", "")
	cbURL        = flag.String("url", "http://couchbase:8091", "")
	cbBucket     = flag.String("bucket", "my_test", "")
	cbBucketAuth = flag.String("bucket-auth", "", "")
	cbBucketRAM  = flag.Int("bucket-ram", 200, "")
)

func main() {
	validArgs, errors := parseArgs()

	if !validArgs {
		// print all errors
		fmt.Printf("%+v", errors)
		return
	}

	c, err := couchbase.Connect(*cbURL)
	if err != nil {
		log.Fatalf("Error connecting:  %v", err)
	}

	_, err = c.GetPool("default")
	if err != nil {
		log.Fatalf("Error getting pool:  %v", err)
	}
}

// parseArgs parses all input arguments and validate each one to check if they are
// in the expected format/values
func parseArgs() (bool, []string) {
	var errors []string

	flag.Parse()

	if *cbAction == "" {
		errors = append(errors, "* No action defined.")
	}

	_, err := url.Parse(*cbURL)
	if *cbURL == "" || err != nil {
		errors = append(errors, "* Invalid Couchbase URL.")
	}

	if strings.Contains(*cbAction, "bucket") && *cbBucket == "" {
		if *cbBucket == "" {
			errors = append(errors, "* You should pass a valid bucket name.")
		}

		if *cbBucketRAM < 100 {
			msg := fmt.Sprintf("* Minimum RAM by bucket is 100 MB (passed %d).", *cbBucketRAM)
			errors = append(errors, msg)
		}
	}

	if len(errors) > 0 {
		return false, errors
	}

	return true, nil
}

// doServerSetup initializes and configures a new Couchbase Server node
// It should takes as an argument a struct with all configs
func doServerSetup() error {
	return nil
}

// changeServerSetup changes init configurations
// It should takes as an argument a struct with all configs
func changeServerSetup() error {
	return nil
}

// createBucket creates a new bucket
// It should takes as an argument a struct with all configs
func createBucket() error {
	return nil
}

// deleteBucket deletes a new bucket
// It should takes as an argument a struct with all configs
func deleteBucket() error {
	return nil
}
