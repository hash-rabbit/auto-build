package model

type Project struct {
	Id        int64  `xorm:"pk"`
	Name      string `xorm:"varchar(20) not null"`
	LocalPath string `xorm:"varchar(50) not null"`
	Url       string `xorm:"varchar(50)"` // 去掉前面的 https://
	Token     string `xorm:"varchar(50)"`
	GoMod     bool   `xorm:"bool"`
	WorkSpace string `xorm:"varchar(50)"`  //only go mod used
	Env       string `xorm:"varchar(255)"` // 环境变量key1=value1;key2=value2
}

func (p *Project) Insert() error {
	_, err := engine.InsertOne(p)
	return err
}

func ListProject() ([]*Project, error) {
	ps := make([]*Project, 0)
	engine.Find(&ps)
	return ps, nil
}
