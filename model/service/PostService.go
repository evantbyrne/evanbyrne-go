package service

import (
	"github.com/evantbyrne/evanbyrne-go/model/dto"
	"github.com/evantbyrne/evanbyrne-go/util"
	"log"
)

func GetPostById(db *util.Database, id string) (dto.Post, bool) {
	var res dto.Post;
	if err := db.Gorp().SelectOne(&res, "select * from post where id = $1 limit 1", id); err == nil {
		res.Meta = GetPostMetaByPostId(db, res.Id)
		return res, true
	}

	return dto.Post{}, false
}

func GetPostByUrl(db *util.Database, url string) (dto.Post, bool) {
	var res dto.Post;
	if err := db.Gorp().SelectOne(&res, "select * from post where url = $1 limit 1", url); err == nil {
		res.Meta = GetPostMetaByPostId(db, res.Id)
		return res, true
	}

	return dto.Post{}, false
}

func GetPostListing(db *util.Database) []dto.Post {
	var res []dto.Post;
	if _, err := db.Gorp().Select(&res, "select * from post order by url asc"); err != nil {
		log.Fatal(err)
	}

	return res
}

func GetPostMetaByPostId(db *util.Database, id int64) []dto.PostMeta {
	var res []dto.PostMeta;
	if _, err := db.Gorp().Select(&res, "select * from post_meta where post_id = $1", id); err != nil {
		log.Fatal(err)
	}

	return res
}

func CreatePost(db *util.Database, post dto.Post) error {
	dbmap := db.Gorp()
	err := dbmap.Insert(&post);
	if err != nil {
		return err
	}

	for _, meta := range post.Meta {
		meta.PostId = post.Id
		if err = dbmap.Insert(&meta); err != nil {
			return err
		}
	}

	return nil
}

func EditPost(db *util.Database, post dto.Post) error {
	if _, err := db.Connection.Exec("delete from post_meta where post_id = $1", post.Id); err != nil {
		return err
	}

	dbmap := db.Gorp()
	if _, err := dbmap.Update(&post); err != nil {
		return err
	}

	for _, meta := range post.Meta {
		meta.PostId = post.Id
		if err := dbmap.Insert(&meta); err != nil {
			return err
		}
	}

	return nil
}