migrateup:
	migrate -path "postgres/migrations" -database "postgres://postgres:1234@localhost:5432/meetup_dev?sslmode=disable" -verbose up

migratedown:
	migrate -path "postgres/migrations" -database "postgres://postgres:1234@localhost:5432/meetup_dev?sslmode=disable" -verbose down