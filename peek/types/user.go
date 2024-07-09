package types

type User struct {
	ID         int      `gorm:"id,pk" json:"id,omitempty"`
	Username   string   `bson:"username" json:"username"`
	Password   string   `bson:"password" json:"-"`
	Level      int      `bson:"level" json:"level"`
	Content    string   `bson:"content" json:"content"`
	Email      string   `bson:"email" json:"email"`
	Job        string   `bson:"job" json:"job"`
	Company    string   `bson:"company" json:"company"`
	Birthday   string   `bson:"birthday" json:"birthday"`
	Age        int      `bson:"age" json:"age"`
	Gender     string   `bson:"gender" json:"gender"`
	Avatar     string   `bson:"avatar" json:"avatar"`
	Tags       []string `bson:"tags" json:"tags"`
	Likes      string   `bson:"likes" json:"likes"`
	Describe   string   `bson:"describe" json:"describe"`
	Collection string   `bson:"collection" json:"collection"`
	IsAdmin    bool     `bson:"isAdmin" json:"isAdmin"`
}
