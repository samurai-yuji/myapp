package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "os"
)

var (
    Db *gorm.DB
)

func InitDB() {

    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    dbase := os.Getenv("DB_DBASE")

    var dsn = user + ":" + pass+ "@" + "tcp(" + host + ")/" + dbase + "?parseTime=true"
    db, err := gorm.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }else{
        Db = db
    }
    Db.LogMode(true)
    autoMigrate()
}

func autoMigrate () {
    Db.AutoMigrate(&Table{})
}
