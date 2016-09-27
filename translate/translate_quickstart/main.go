// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// [START translate_quickstart]
// Sample translate_quickstart translates "Hello, world!" into Russian.
package main

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/text/language"
	"log"

	// Imports the Google Cloud Translate client package
	"cloud.google.com/go/translate"
)

func main() {
	ctx := context.Background()

	// Creates a client
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// The text to translate
	text := "Hello, world!"
	// The target language
	target, err := language.Parse("ru")
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates some text into Russian
	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	fmt.Printf("Translation: %v", translations[0].Text)
}

// [END translate_quickstart]
