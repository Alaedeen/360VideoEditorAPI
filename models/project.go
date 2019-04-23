package models

import (
	"github.com/jinzhu/gorm"
)

// Project Struct
type Project struct {
	gorm.Model
	UserID			int 			`json:"userId"`
	Title			string 			`json:"title"`
	Thumbnail 		string			`json:"thumbnail"`
	AFrame 			string			`json:"aFrame"`
	Video 			string			`json:"video"`
	Box 			int				`json:"box"`
	Sphere 			int				`json:"sphere"`
	Cone 			int				`json:"cone"`
	Cylinder 		int				`json:"cylinder"`
	Torus 			int				`json:"torus"`
	TorusKnot 		int				`json:"torus-knot"`
	Dodecahedron 	int				`json:"dodecahedron"`
	Tetrahedron 	int				`json:"tetrahedron"`
	Image 			int				`json:"image"`
	Video2D 		int				`json:"2dVideo"`
	Text 			int				`json:"text"`
	Tag 			int				`json:"tag"`
	ShapesList		[]AddedShapes 	`json:"shapesList"`
	TagsList		[]AddedTags		`json:"tagsList"`
}