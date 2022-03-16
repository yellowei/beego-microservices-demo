package models

//表的设计和创建
import (
	"github.com/astaxie/beego"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Name     string
	Pwd      string
	Articles []*Article `orm:"reverse(many)"`
}

//文章和类型是什么关系
//用户和文章是什么关系  多对多  单表操作  primary key
type Article struct {
	Id        int       `orm:"pk;auto"`
	Title     string    `orm:"unique;size(40)"`
	Content   string    `orm:"size(500)"`
	Img       string    `orm:"null"`
	Time      time.Time `orm:"type(datetime);auto_now_add"`
	ReadCount int       `orm:"default(0)"`

	ArticleType *ArticleType `orm:"rel(fk);null;on_delete(set_null)"`
	Users       []*User      `orm:"rel(m2m)"`
}
type ArticleType struct {
	Id       int
	TypeName string     `orm:"unique"`
	Article  []*Article `orm:"reverse(many)"`
}

func init() {
	//注册数据库
	beego.Info("注册数据库")
	//  mysql -h radondb-smileai-radondb-m-leader -uroot -p
	dbHost := beego.AppConfig.String("db_host")
	dbName := beego.AppConfig.String("db_dbname")
	dbPassword := beego.AppConfig.String("db_password")
	dbUser := beego.AppConfig.String("db_user")
	orm.RegisterDataBase("default", "mysql", dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/"+ dbName )
	//注册表
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	//跑起来
	orm.RunSyncdb("default", false, true)
}

