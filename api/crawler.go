package api

import (
	"FBcrawler/task"
	"FBcrawler/types"
)

func Crawler(name string) *[]*types.Record {
	var records *[]*types.Record
	records = new([]*types.Record)
	var fbc *task.FacebookCrawler
	fbc.Do(name, records)

	return records
}
