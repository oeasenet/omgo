package options

import "go.mongodb.org/mongo-driver/mongo/options"

type DatabaseOptions struct {
	*options.DatabaseOptions
}

type ListCollectionsOptions struct {
	*options.ListCollectionsOptions
}
