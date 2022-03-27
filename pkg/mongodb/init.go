package mongodb

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
	"workspace/ginweb/conf"
)

var (
	mgoSession *mgo.Session
	dbName     string
)

type MgoLookUp struct {
	Collection          string
	JoinCollection      string
	LocalField          string
	ForeignField        string
	AsField             string
	Selector            bson.M
	Result              interface{}
	Sorter, Skip, Limit bson.M
}

type Err struct {
	Code int32  `bson:"id"`
	Msg  string `bson:"client_msg"`
}

type Mgo struct {
	Session *mgo.Session
	Name    *mgo.Collection
}

func (m *Mgo) NewMgo() {
	m.Session = mgoSession.Copy()
}
func (m *Mgo) ReleaseMgo() {
	m.Session.Close()
	*m = Mgo{}
}
func (m *Mgo) GetCollcetion(collection string) *mgo.Collection {
	return m.Session.DB(dbName).C(collection)
}

func Init() (err error) {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{conf.MongoAddr},
		Database:  conf.MongoDB,
		Username:  conf.MongoUser,
		Password:  conf.MongoPass,
		PoolLimit: conf.MongoLimit, // Session.SetPoolLimit
		Timeout:   5 * time.Second,
	}
	dbName = conf.MongoDB
	mgoSession, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		return
	}
	return nil
}

// 获取session
func getSession() *mgo.Session {
	return mgoSession.Copy()
}

//获取数量
func Count(collection string, selector bson.M) (int, error) {
	session := getSession()
	defer session.Close()
	return session.DB(dbName).C(collection).Find(selector).Count()
}

// 根据条件获取一条记录
func FindOne(collection string, selector bson.M, fields bson.M, result interface{}, sorter ...string) error {
	session := getSession()
	defer session.Close()
	query := session.DB(dbName).C(collection).Find(selector)
	if fields != nil && len(fields) > 0 {
		query = query.Select(fields)
	}
	if len(sorter) > 0 && sorter[0] != "" {
		query = query.Sort(sorter...)
	}
	err := query.One(result)
	if nil != err && "not found" == err.Error() {
		return nil
	}
	return err
}

// 获取多条记录
func FindAll(collection string, selector bson.M, fields bson.M, result interface{}, sorter ...string) error {
	session := getSession()
	defer session.Close()
	query := session.DB(dbName).C(collection).Find(selector)
	if fields != nil && len(fields) > 0 {
		query = query.Select(fields)
	}
	if len(sorter) > 0 && sorter[0] != "" {
		query = query.Sort(sorter...)
	}
	err := query.All(result)
	if nil != err && "not found" == err.Error() {
		return nil
	}
	return err
}

// 获取多条记录
func FindAllWithLimit(collection string, selector bson.M, fields bson.M, limit int, result interface{}, sorter ...string) error {
	session := getSession()
	defer session.Close()
	query := session.DB(dbName).C(collection).Find(selector)
	if fields != nil && len(fields) > 0 {
		query = query.Select(fields)
	}
	if len(sorter) > 0 && sorter[0] != "" {
		query = query.Sort(sorter...)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	err := query.All(result)
	if nil != err && "not found" == err.Error() {
		return nil
	}
	return err
}

// 获取多条分页记录（返回记录总数）
func FindAllWithPage(collection string, selector bson.M, fields bson.M, pageSize int, curPage int, result interface{}, sorter ...string) (int, error) {
	session := getSession()
	defer session.Close()
	// 获取记录总数
	nCount := 0
	nCount, err := session.DB(dbName).C(collection).Find(selector).Count()
	if err != nil {
		if "not found" == err.Error() {
			return nCount, nil
		}
		return nCount, err
	}
	//获取分页记录
	query := session.DB(dbName).C(collection).Find(selector)
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	if len(sorter) > 0 && sorter[0] != "" {
		query = query.Sort(sorter...)
	}
	err = query.Skip((curPage - 1) * pageSize).Limit(pageSize).All(result)
	if nil != err && "not found" == err.Error() {
		return nCount, nil
	}
	return nCount, err
}

//mongo apply
func GetMaxUserid(collection string, selector bson.M, changer mgo.Change) (maxid int32, err error) {
	var max = struct {
		Userid int32 `bson:"maxuserid"`
	}{}
	session := getSession()
	defer session.Close()
	_, err = session.DB(dbName).C(collection).Find(selector).Apply(changer, &max)
	return max.Userid, err
}

//mongo apply
func GetMaxBatchID(collection string, selector bson.M, changer mgo.Change) (maxid int32, err error) {
	var max = struct {
		BatchID int32 `bson:"maxbarchid"`
	}{}
	session := getSession()
	defer session.Close()
	_, err = session.DB(dbName).C(collection).Find(selector).Apply(changer, &max)
	return max.BatchID, err
}

//添加单条记录
func Insert(collection string, records bson.M) error {
	session := getSession()
	defer session.Close()
	err := session.DB(dbName).C(collection).Insert(records)
	return err
}

//批量添加记录
func InsertAll(collection string, records []interface{}) error {
	session := getSession()
	defer session.Close()
	err := session.DB(dbName).C(collection).Insert(records...)
	return err
}

//修改单条记录
func Update(collection string, selector bson.M, changer bson.M) error {
	session := getSession()
	defer session.Close()
	err := session.DB(dbName).C(collection).Update(selector, changer)
	/*	if collection == define.CUserinfo {
		if v, ok := changer["$set"]; ok {
			if _, _ok := (v.(bson.M))["baseinfo.dan"]; _ok {
				for i := 1; ; i++ {
					_, file, line, ok := runtime.Caller(i)
					if !ok {
						break
					}
					logs.Errorf("danuser%v,%s:%d,changer:%v", selector["userid"], file, line, changer)
				}
			}
		}
	}*/
	return err
}

//修改多条记录
func UpdateAll(collection string, selector bson.M, changer bson.M) (*mgo.ChangeInfo, error) {
	session := getSession()
	defer session.Close()
	changeInfo, err := session.DB(dbName).C(collection).UpdateAll(selector, changer)
	return changeInfo, err
}

//删除单条记录
func Remove(collection string, selector bson.M) error {
	session := getSession()
	defer session.Close()
	err := session.DB(dbName).C(collection).Remove(selector)
	return err
}

//删除多条记录
func RemoveAll(collection string, selector bson.M) (*mgo.ChangeInfo, error) {
	session := getSession()
	defer session.Close()
	changeInfo, err := session.DB(dbName).C(collection).RemoveAll(selector)
	return changeInfo, err
}

//获取mgo session
func GetSession() *mgo.Session {
	return mgoSession.Copy()
}

func DropCollection(collection string) error {
	session := getSession()
	defer session.Close()
	return session.DB(dbName).Run(bson.D{{Name: "drop", Value: collection}}, nil)
}

func LookUp(req *MgoLookUp) error {
	session := getSession()
	defer session.Close()
	if len(req.Sorter) == 0 {
		//不排序不分页
		return session.DB(dbName).C(req.Collection).Pipe([]bson.M{
			bson.M{"$match": req.Selector},
			bson.M{"$lookup": bson.M{
				"from":         req.JoinCollection,
				"localField":   req.LocalField,
				"foreignField": req.ForeignField,
				"as":           req.AsField,
			}},
		}).All(req.Result)
	} else {
		if len(req.Limit) == 0 {
			//排序不分页
			return session.DB(dbName).C(req.Collection).Pipe([]bson.M{
				bson.M{"$match": req.Selector},
				bson.M{"$lookup": bson.M{
					"from":         req.JoinCollection,
					"localField":   req.LocalField,
					"foreignField": req.ForeignField,
					"as":           req.AsField,
				}},
				bson.M{"$sort": req.Sorter},
			}).All(req.Result)
		} else {
			//排序分页
			return session.DB(dbName).C(req.Collection).Pipe([]bson.M{
				bson.M{"$match": req.Selector},
				bson.M{"$lookup": bson.M{
					"from":         req.JoinCollection,
					"localField":   req.LocalField,
					"foreignField": req.ForeignField,
					"as":           req.AsField,
				}},
				bson.M{"$sort": req.Sorter},
				bson.M{"$skip": req.Skip},
				bson.M{"$limit": req.Limit},
			}).All(req.Result)
		}
	}
}
