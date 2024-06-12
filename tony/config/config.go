package config

import "os"

func GetPort() string {
	port := os.Getenv("PORT")
	// if port == "" {
	// 	return ":3001"
	// }
	return port
}

func GetMongoUrI() string {
	mongoUri := os.Getenv("MONGOURI")
	if mongoUri == "" {
		return "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.5"
	}
	return mongoUri
}

func GetDBName() string {
	DBName := os.Getenv("DBNAME")
	if DBName == "" {
		return "tony"
	}
	return DBName
}

func GetUserColl() string {
	userColl := os.Getenv("USERCOLL")
	if userColl == "" {
		return "users"
	}
	return userColl
}
func GetJwtSecret() string {
	jwtSecret := os.Getenv("JWTSECRET")
	if jwtSecret == "" {
		return "little fan"
	}
	return jwtSecret
}
func GetMD5Secret() string {
	md5Secret := os.Getenv("MD5SECRET")
	if md5Secret == "" {
		return "users"
	}
	return md5Secret
}

func GetPostColl() string {
	postColl := os.Getenv("POSTCOLL")
	if postColl == "" {
		return "posts"
	}
	return postColl
}

func GetCommentsColl() string {
	commentsColl := os.Getenv("COMMENTSCOLL")
	if commentsColl == "" {
		return "comments"
	}
	return commentsColl
}

func GetTagsColl() string {
	tagsColl := os.Getenv("TAGSCOLL")
	if tagsColl == "" {
		return "tags"
	}
	return tagsColl
}
