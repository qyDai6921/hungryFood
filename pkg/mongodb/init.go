package mongodb

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
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
	Selector            bson.M // 通过bson.M(是一个map[string]interface{}类型)进行条件筛选，达到文档查询的目的
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
		PoolLimit: conf.MongoLimit, // Session.SetPoolLimit 管理连接数量， 最大连接数的 mongodb atlas
		Timeout:   5 * time.Second,
	}
	dbName = conf.MongoDB
	mgoSession, err = mgo.DialWithInfo(dialInfo) // //创建一个维护套接字池的session
	if err != nil {
		return
	}
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

// 获取session
func getSession() *mgo.Session {
	return mgoSession.Copy()
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

//添加单条记录
func Insert(collection string, records bson.M) error {
	session := getSession()
	defer session.Close() //如果每个 session都不调用 close()，会导致连接数会很快就达到 PoolLimit，并堵死其他请求
	err := session.DB(dbName).C(collection).Insert(records)
	return err
}
