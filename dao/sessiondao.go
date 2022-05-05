package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

//添加sesion
func AddSession(sess *model.Session) error {
	//写sql语句
	sqlStr := "insert into sessions values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, sess.SessionId, sess.UserName, sess.UserId)
	if err != nil {
		fmt.Println("预编译和执行出错:", err)
		return err
	}
	return nil
}

//根据sessionID从数据库中删除session
func DelSession(sessId string) error {
	//写sql语句
	sqlStr := "delete from sessions where session_id=?"
	_, err := utils.Db.Exec(sqlStr, sessId)
	if err != nil {
		return err
	}

	return nil
}

//根据sesiomId从数据库中查询sesion信息
func GetSession(sessId string) (*model.Session, error) {
	//写sql语句
	sqlStr := "select session_id,username,user_id from sessions where session_id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, sessId) //QueryRow执行一次返回一行结果
	sess := &model.Session{}
	row.Scan(&sess.SessionId, &sess.UserName, &sess.UserId)
	return sess, nil
}
