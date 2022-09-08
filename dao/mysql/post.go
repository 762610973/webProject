package mysql

import (
	"strings"
	"webProject/models"

	"github.com/jmoiron/sqlx"
)

// CreatePost 将创建的帖子保存进数据库post表中
func CreatePost(p *models.Post) error {
	sqlStr := `INSERT INTO post (post_id,title,content,author_id,community_id)
				VALUES (?,?,?,?,?)`
	_, err := db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err != nil {
		return err
	}
	return nil
}

// GetPostById 通过id获取帖子
func GetPostById(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `SELECT post_id,title,content,author_id,community_id,create_time
				FROM post 
				WHERE post_id = ?`
	err = db.Get(post, sqlStr, id)
	return
}

// GetPostList 分页获取帖子
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `SELECT post_id,content,title,author_id,community_id,create_time
				FROM post ORDER BY create_time
				DESC
				LIMIT ?,?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// GetPostListByIDs 根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	// in的用法查看https://www.liwenzhou.com/posts/Go/sqlx
	sqlStr := `SELECT post_id,title,content,author_id,community_id,create_time
				FROM post
				WHERE post_id in (?)
				ORDER BY FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ""))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
