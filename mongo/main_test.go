package mongo_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	mongo_db "github.com/lysofts/database-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	UserCollectionName = "test_users"
)

//TestData .. is the object used to test data objects CRUD
type TestData struct {
	ID            string  `json:"id,omitempty" bson:"_id"`
	Name          string  `json:"name,omitempty"`
	Price         float32 `json:"price,omitempty"`
	PostalAddress string  `json:"postalAddress,omitempty" bson:"postalAddress"`
}

//createTestUser create test data
func createTestUser(ctx context.Context, t *testing.T, uid string) error {
	db := mongo_db.New()
	data := TestData{
		ID:            uid,
		Name:          "Test",
		Price:         200.50,
		PostalAddress: "Home, Test Address",
	}
	_, err := db.Create(ctx, UserCollectionName, data)
	if err != nil {
		t.Errorf("error, unable to create test user, %v", err)
	}

	return err
}

func TestDatabaseImpl_GetOne(t *testing.T) {
	ctx := context.Background()
	db := mongo_db.New()

	UID := uuid.NewString()

	err := createTestUser(ctx, t, UID)

	if err != nil {
		t.Errorf(err.Error())
		return

	}

	type args struct {
		ctx            context.Context
		collectionName string
		data           interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "happy found one",
			args: args{
				ctx:            ctx,
				collectionName: UserCollectionName,
				data:           bson.M{"id": UID},
			},
			want:    UID,
			wantErr: false,
		},

		{
			name: "sad unabale to find one",
			args: args{
				ctx:            ctx,
				collectionName: UserCollectionName,
				data:           bson.M{"id": "test"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.ReadOne(tt.args.ctx, tt.args.collectionName, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DatabaseImpl.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != got["id"] {
				t.Errorf("DatabaseImpl.GetOne() = %v, want %v", got, tt.want)
			}
		})
	}

	_, err = db.Delete(ctx, UserCollectionName, bson.M{"id": UID})
	if err != nil {
		t.Errorf("error, unable to delete test user, %v", err)
		return
	}
}
func TestGet(t *testing.T) {
	ctx := context.Background()
	db := mongo_db.New()

	UID := uuid.NewString()

	err := createTestUser(ctx, t, UID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	type args struct {
		ctx            context.Context
		collectionName string
		filter         bson.M
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "happy got test user by name",
			args: args{
				ctx: ctx, collectionName: UserCollectionName, filter: bson.M{
					"name": "Test",
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "sad: query does not match any",
			args: args{
				ctx: ctx, collectionName: UserCollectionName, filter: bson.M{
					"name": "123",
				},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.Read(tt.args.ctx, tt.args.collectionName, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != len(got) {
				t.Errorf("Update() got = %v, wantNil %v", got, tt.want)
				return
			}
		})
	}

	_, err = db.Delete(ctx, UserCollectionName, bson.M{"_id": UID})
	if err != nil {
		t.Errorf("error, unable to delete test user, %v", err)
		return
	}
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	db := mongo_db.New()

	type args struct {
		ctx            context.Context
		collectionName string
		data           interface{}
	}

	UID := uuid.NewString()

	data := TestData{
		ID:            UID,
		Name:          "Test",
		Price:         200.50,
		PostalAddress: "Home, Test Address",
	}

	tests := []struct {
		name    string
		args    args
		want    *mongo.InsertOneResult
		wantErr bool
	}{
		{
			name: "happy created user",
			args: args{
				ctx:            ctx,
				collectionName: UserCollectionName,
				data:           data,
			},
			wantErr: false,
			want: &mongo.InsertOneResult{
				InsertedID: UID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.Create(tt.args.ctx, tt.args.collectionName, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
	_, err := db.Delete(ctx, UserCollectionName, bson.M{"_id": UID})
	if err != nil {
		t.Errorf("error, unable to delete test user, %v", err)
		return
	}
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	db := mongo_db.New()

	UID := uuid.NewString()
	_ = createTestUser(ctx, t, UID)

	type args struct {
		ctx            context.Context
		collectionName string
		filter         bson.M
		data           bson.M
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
		wantErr bool
	}{

		{
			name: "happy updated data",
			args: args{
				ctx:            ctx,
				collectionName: UserCollectionName,
				filter:         bson.M{"_id": UID},
				data:           map[string]interface{}{"firstName": "Test3"},
			},
			wantNil: false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.Update(tt.args.ctx, tt.args.collectionName, tt.args.filter, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantNil && got != nil {
				t.Errorf("Update() got = %v, wantNil %v", got, tt.wantNil)
				return
			}
		})
	}

	_, err := db.Delete(ctx, UserCollectionName, bson.M{"_id": UID})
	if err != nil {
		t.Errorf("error, unable to delete test user, %v", err)
		return
	}
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	db := mongo_db.New()

	UID := uuid.NewString()
	_ = createTestUser(ctx, t, UID)

	type args struct {
		ctx            context.Context
		collectionName string
		filer          bson.M
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "happy deleted data",
			args: args{
				ctx:            ctx,
				collectionName: UserCollectionName,
				filer:          bson.M{"_id": UID},
			},
			wantErr: false,
			want:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.Delete(tt.args.ctx, tt.args.collectionName, tt.args.filer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
