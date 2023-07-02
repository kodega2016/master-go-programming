package main

import "fmt"

func main() {
	rd := RemoteDataSource{}
	ld := LocalDataSource{}

	logData(rd)
	logData(ld)
}

type DataSource interface {
	getData() string
}

type RemoteDataSource struct {
}

// getData implements DataSource
func (RemoteDataSource) getData() string {
	return "data from the remote server"
}

type LocalDataSource struct{}

// getData implements DataSource
func (LocalDataSource) getData() string {
	return "data from the local data source"
}

func logData(source DataSource) {
	fmt.Println(source.getData())
}
