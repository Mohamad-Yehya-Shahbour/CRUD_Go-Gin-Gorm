package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Posts(g *gin.Context) {
	//limit := g.DefaultQuery("limit","10")
	//offset := g.DefaultQuery("offset","0")
	var post []Post
	db.Find(&post)
	g.JSON(http.StatusOK, gin.H{
		"msg":  "all posts",
		"data": post,
	})
}

func Show(g *gin.Context) {

	post := GetById(g)
	if post.ID == 0 {
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"msg":  "",
		"data": post,
	})

}

func Store(g *gin.Context) {

	var post Post
	if err := g.ShouldBindJSON(&post); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	db.Create(&post)

	g.JSON(http.StatusCreated, gin.H{
		"msg":  "post has been added",
		"data": post,
	})

}

func Update(g *gin.Context) {

	oldpost := GetById(g)
	if oldpost.ID == 0 {
		return
	}

	var newpost Post
	if err := g.ShouldBindJSON(&newpost); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	oldpost.Post = newpost.Post
	oldpost.Desc = newpost.Desc
	db.Save(oldpost)

	g.JSON(http.StatusOK, gin.H{
		"msg":  "post has been updated",
		"data": oldpost,
	})
}

func Delete(g *gin.Context) {

	post := GetById(g)
	if post.ID == 0 {
		return
	}

	db.Delete(&post)

	g.JSON(http.StatusOK, gin.H{
		"msg":  "post has been deleted",
		"data": "",
	})

}

func GetById(g *gin.Context) Post {
	var post Post
	id := g.Param("id")
	db.First(&post, id)
	if post.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{
			"msg":  "",
			"data": "",
		})
	}
	return post
}
