// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package snippets contains snippets for the Google BigQuery Go package.
package snippets

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

func createDataset(client *bigquery.Client, datasetID string) error {
	ctx := context.Background()
	// [START bigquery_create_dataset]
	if err := client.Dataset(datasetID).Create(ctx); err != nil {
		return err
	}
	// [END bigquery_create_dataset]
	return nil
}

func listDatasets(client *bigquery.Client) error {
	ctx := context.Background()
	// [START bigquery_list_datasets]
	it := client.Datasets(ctx)
	for {
		dataset, err := it.Next()
		if err == iterator.Done {
			break
		}
		fmt.Println(dataset.DatasetID)
	}
	// [END bigquery_list_datasets]
	return nil
}

// [START bigquery_create_table]

// Item represents a row item.
type Item struct {
	Name  string
	Count int
}

// Save implements the ValueSaver interface.
func (i *Item) Save() (map[string]bigquery.Value, string, error) {
	return map[string]bigquery.Value{
		"Name":  i.Name,
		"Count": i.Count,
	}, "", nil
}

// [END bigquery_create_table]

func createTable(client *bigquery.Client, datasetID, tableID string) error {
	ctx := context.Background()
	// [START bigquery_create_table]
	schema, err := bigquery.InferSchema(Item{})
	if err != nil {
		return err
	}
	table := client.Dataset(datasetID).Table(tableID)
	if err := table.Create(ctx, schema); err != nil {
		return err
	}
	// [END bigquery_create_table]
	return nil
}

func insertRows(client *bigquery.Client, datasetID, tableID string) error {
	ctx := context.Background()
	// [START bigquery_insert_stream]
	u := client.Dataset(datasetID).Table(tableID).Uploader()
	items := []*Item{
		// Item implements the ValueSaver interface.
		{Name: "n1", Count: 7},
		{Name: "n2", Count: 2},
		{Name: "n3", Count: 1},
	}
	if err := u.Put(ctx, items); err != nil {
		return err
	}
	// [ENDSTART bigquery_insert_stream]
	return nil
}

func copyTable(client *bigquery.Client, datasetID, srcID, dstID string) error {
	ctx := context.Background()
	// [START bigquery_copy_table]
	dataset := client.Dataset(datasetID)
	copier := dataset.Table(dstID).CopierFrom(dataset.Table(srcID))
	copier.WriteDisposition = bigquery.WriteTruncate
	job, err := copier.Run(ctx)
	if err != nil {
		return err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}
	if err := status.Err(); err != nil {
		return err
	}
	// [END bigquery_copy_table]
	return nil
}

func deleteTable(client *bigquery.Client, datasetID, tableID string) error {
	ctx := context.Background()
	// [START bigquery_delete_table]
	table := client.Dataset(datasetID).Table(tableID)
	if err := table.Delete(ctx); err != nil {
		return err
	}
	// [END bigquery_delete_table]
	return nil
}
