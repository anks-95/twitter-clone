//Defines a Post model with Gorm tags

package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
}
