package main

import "fmt"

type BuilderPattern struct {
}

func (p BuilderPattern) Show() {
	builder := SimpleSQLBuilder{}

	result := builder.Select().Where("age > 1").Limit(50).Build()
	fmt.Println(result)
	result = builder.Select().Where("name = 'Yevhenii'").Limit(2).Build()
	fmt.Println(result)
}

func (p BuilderPattern) Name() string {
	return "Builder"
}

type SimpleSQLBuilder struct {
	queryType string
	where     string
	limit     int
}

type SQLBuilder interface {
	Select() SQLBuilder
	Where(string) SQLBuilder
	Limit(int) SQLBuilder
	Build() string
}

func (s *SimpleSQLBuilder) Select() SQLBuilder {
	s.queryType = "SELECT"
	return s
}

func (s *SimpleSQLBuilder) Where(where string) SQLBuilder {
	s.where = where
	return s
}

func (s *SimpleSQLBuilder) Limit(limit int) SQLBuilder {
	s.limit = limit
	return s
}

func (s *SimpleSQLBuilder) Build() string {
	return fmt.Sprintf("Your SQL query... kind of SQL query: '%s * FROM some_table WHERE %s LIMIT %d'",
		s.queryType, s.where, s.limit)
}
