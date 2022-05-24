package storage

type Book struct {
	ID    int32  `db:"id"`
	Title string `db:"title"`
}

type Author struct {
	ID   int32  `db:"id"`
	Name string `db:"name"`
}
