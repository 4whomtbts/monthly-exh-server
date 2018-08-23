package core

import (
	"goServer/store"
	"goServer/store/mongoStore"
	"goServer/models"
	"goServer/mlog"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"os"
)

type Core struct {

	Srv *Server // 서버관련 설정 및 메소드
	Log *mlog.Logger
	Configuration *models.Config
	DBSetting     *models.MongoSetting
	Logger  *mlog.Logger
	AwsSession *session.Session
	S3 *s3.S3

	logListener string
	configListenerId string


}

func NewCore() *Core {

	core := &Core{
		Srv:       InitServer(),
		DBSetting: DBSetting(),
	}
	core.Srv.Store = store.NewLayeredStore(mongoStore.NewMongoStore(core.DBSetting))

	core.Configuration = &models.Config{
		AwsSettings : models.AwsSettings{},
		LogSettings: models.LogSettings{},
	}

	core.Configuration.SetDefaults()
	core.Logger = mlog.NewLogger(*core.Configuration.LogSettings.LogConfig)


	sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState : session.SharedConfigEnable,
	}))
	svc := s3.New(sess, &aws.Config{
		Region : aws.String(endpoints.ApNortheast2RegionID),
	})
	result , err := svc.ListBuckets(nil)
	if err != nil {
		fmt.Println("unable to list bucket : ",err)
	}


	core.AwsSession = sess
	core.S3 = svc

	for _,  b := range result.Buckets {
		fmt.Println("* %s created on %s \n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
	pwd , _ := os.Getwd()
	file , err := os.Open(pwd+"/robot.jpg")
	if err != nil {
		fmt.Println("file load error  : ",err)
	}
	defer file.Close()

	fmt.Println(file)
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()

	buffer := make([]byte, size)
	file.Read(buffer)



	mlog.InitGlobalLogger(core.Logger)
	mlog.Info("서버가동")

	return core
}

	func StartCore() *Core {

	Core := NewCore()
	Core.Start()
	return Core
}

func DBSetting() *models.MongoSetting {
	m0 := &models.MongoSetting{}
	m0.SetDefaults()
	return m0
}
