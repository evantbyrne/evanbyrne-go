package service

import (
	"github.com/evantbyrne/evanbyrne-go/model/dto"
	"github.com/evantbyrne/evanbyrne-go/util"
	"log"
)

func GetPostById(id string) (dto.Post, bool) {
	var db = new(util.Database).Gorp()
	var res dto.Post;
	if err := db.SelectOne(&res, "select * from post where id = $1 limit 1", id); err == nil {
		res.Meta = GetPostMetaByPostId(res.Id)
		return res, true
	}

	return dto.Post{}, false
}

func GetPostByUrl(url string) (dto.Post, bool) {
	var db = new(util.Database).Gorp()
	var res dto.Post;
	if err := db.SelectOne(&res, "select * from post where url = $1 limit 1", url); err == nil {
		res.Meta = GetPostMetaByPostId(res.Id)
		return res, true
	}

	return dto.Post{}, false
}

func GetPostListing() []dto.Post {
	var db = new(util.Database).Gorp()
	var res []dto.Post;
	if _, err := db.Select(&res, "select * from post order by url asc"); err != nil {
		log.Fatal(err)
	}

	return res
}

func GetPostMetaByPostId(id int64) []dto.PostMeta {
	var db = new(util.Database).Gorp()
	var res []dto.PostMeta;
	if _, err := db.Select(&res, "select * from post_meta where post_id = $1", id); err != nil {
		log.Fatal(err)
	}

	return res
}

func CreatePost(post dto.Post) error {
	dbmap := new(util.Database).Gorp()
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