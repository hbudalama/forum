package db

import (
	"learn.reboot01.com/git/hbudalam/forum/pkg/structs"
)

func GetUser(userId int) (structs.User, error) {
	panic("not implemented")
}

func GetUserPassword(username string) (string,error){
	panic("not implemented")
}

func GetPost(postId int)(structs.Post,error){
	panic("not implemented")
}

func GetComment(commentId int)(structs.Comment, error){
	panic("not implemented")
}

func GetSession(token string)(structs.Session , error){
	panic("not implemented")
}
