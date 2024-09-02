package mongo

import (
	"context"
	"fmt"
	"testing"
)

func TestSaveAMap(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://koala:koala@localhost:27017/koala", "koala")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	insertedID, err := cli.SaveOne(ctx, "test", map[string]string{"name": "koala"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Scanf("Inserted ID %v", insertedID))
}

func TestSaveAStruct(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://koala:koala@localhost:27017/koala", "koala")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		Name string
	}
	insertedID, err := cli.SaveOne(ctx, "test", Test{"koala"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Scanf("Inserted ID %v", insertedID))
}

func TestSaveManyStruct(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://koala:koala@localhost:27017/koala", "koala")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		Name string
	}
	err = cli.SaveMany(ctx, "test", []interface{}{Test{"koala2"}, Test{"koala2"}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetStruct(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://koala:koala@localhost:27017/koala", "koala")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		Name string
	}
	var target Test
	err = cli.FindOne(ctx, "test", map[string]any{"name": "koala"}, &target)
	if err != nil {
		t.Fatal(err)
	}
	if target.Name != "koala" {
		t.Errorf("Expected koala, got %v", target.Name)
	}
}

func TestGetList(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://koala:koala@localhost:27017/koala", "koala")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		ID   string `bson:"_id"`
		Name string `bson:"name"`
	}
	var target []Test
	err = cli.FindMany(ctx, "test", map[string]any{"name": "koala"}, 2, 2, &target)
	if err != nil {
		t.Fatal(err)
	}
	if len(target) != 2 {
		t.Errorf("Expected 2, got %v", len(target))
	}
}

func TestUpdateOne(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://koala:koala@localhost:27017/koala", "koala")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		Name string
	}
	err = cli.UpdateOne(ctx, "test", map[string]any{"name": "koala"}, map[string]any{"$set": map[string]any{"name": "koala-update"}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateMany(t *testing.T) {
	ctx := context.TODO()
	cli, err := NewMongoClient(ctx, "mongodb://koala:koala@localhost:27017/koala", "koala")
	if err != nil {
		t.Errorf("Failed to create mongo client: %v", err)
	}
	type Test struct {
		Name string
	}
	err = cli.UpdateMany(ctx, "test", map[string]any{"name": "koala"}, map[string]any{"$set": map[string]any{"name": "koala-update"}})
	if err != nil {
		t.Fatal(err)
	}
}
