// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// [START datastore_quickstart]
// Sample datastore_quickstart fetches an entity from Google Cloud Datastore.
package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"

	// Imports the Google Cloud Datastore client package
	"cloud.google.com/go/datastore"
)

type Task struct {
	Description string
}

func main() {
	ctx := context.Background()

	// Your Google Cloud Platform project ID
	projectID := "YOUR_PROJECT_ID"

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// The kind for the new entity
	kind := "Task"
	// The name/ID for the new entity
	//name := "sampletask1"
	// The Cloud Datastore key for the new entity
	key := datastore.IDKey(kind, 0, nil) 
	//alternatively datastore.NameKey(kind, name, nil)

	// Prepares the new entity
	task := new(Task)
	task.Description = "Buy Milk"

	// Saves the entity
	if _, err := client.Put(ctx, key, task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

	fmt.Printf("Saved %v: %v \n", key.String(), task.Description)
}

// [END datastore_quickstart]
