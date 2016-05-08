/*------------------------
name        mongo
describe    wmntec mongo library
author      ailn(z.ailn@wmntec.com)
create      2016-05-07
version     1.0
------------------------*/
package mongo

import (
	//third party package
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	Url     string
	Session *mgo.Session
}

//create new mongo
func NewMongo(url string) (*Mongo, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return &Mongo{}, err
	}
	session.SetMode(mgo.Monotonic, true)
	return &Mongo{url, session}, nil
}

//insert some document
func (m *Mongo) Insert(dbName, collName string, docs ...interface{}) error {
	return m.Session.DB(dbName).C(collName).Insert(docs...)
}

//find one document
func (m *Mongo) FindOne(dbName, collName string, query map[string]interface{}, result interface{}) error {
	return m.Session.DB(dbName).C(collName).Find(bson.M(query)).One(result)
}

//find all document
func (m *Mongo) FindAll(dbName, collName string, query map[string]interface{}, result interface{}) error {
	return m.Session.DB(dbName).C(collName).Find(bson.M(query)).All(result)
}

//count of document
func (m *Mongo) Count(dbName, collName string, query map[string]interface{}) int {
	count, err := m.Session.DB(dbName).C(collName).Find(bson.M(query)).Count()
	if err != nil {
		return 0
	}
	return count
}

//remove documents
func (m *Mongo) Remove(dbName, collName string, query map[string]interface{}) error {
	return m.Session.DB(dbName).C(collName).Remove(bson.M(query))
}

//close mongo
func (m *Mongo) Close() {
	m.Session.Close()
}
