package dto

type Post struct {
	Id int64 `db:"id"`
	Url string `db:"url"`
	Meta []PostMeta `db:"-"`
}

type PostBlogResult struct {
	Id int64 `db:id`
	Url string `db:url`
	Title string `db:title`
}

type PostMeta struct {
	PostId int64 `db:"post_id"`
	Key string `db:"key"`
	Value string `db:"value"`
}