package main

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Movie holds a movie data
type Movie struct {
	Name      string   ` bson:"name"`
	Year      string   ` bson:"year"`
	Directors []string ` bson:"directors"`
	Writers   []string ` bson:"writers"`
	BoxOffice ` bson:"boxOffice"`
}

//BoxOffice is nested in Movie
type BoxOffice struct {
	Budget uint64 ` bson:"budget"`
	Gross  uint64 ` bson:"gross"`
}

func main() {
	session, err := mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("appdb").C("movies")

	//Create a movie

	darkNight := &Movie{
		Name:      "The Dark Night 2",
		Year:      "2009",
		Directors: []string{"Diego Costa"},
		Writers:   []string{"Nolan, Christ"},
		BoxOffice: BoxOffice{
			Budget: 190000000,
			Gross:  689100001,
		},
	}

	// Insert into Mongo

	err = c.Insert(darkNight)
	if err != nil {
		log.Fatal(err)
	}

	// Noe query movie back

	result := Movie{}
	err = c.Find(bson.M{"boxOffice.budget": bson.M{"$gt": 150000000}}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie:", result.Name)

}
