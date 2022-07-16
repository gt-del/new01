package main

import (
	"fmt"
	"log"
	"taoguo/bolt"
)

func main() {
	db, err := bolt.Open("test.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic(err)
	}
	//对数据库进行读写操作
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建bucket(b1)失败")
			}
		}
		//写入数据
		bucket.Put([]byte("11111"), []byte("hello"))
		bucket.Put([]byte("22222"), []byte("world"))
		return nil
	})
	//读取数据
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			log.Panic("未找到抽屉")
		}
		v1 := bucket.Get([]byte("11111"))
		v2 := bucket.Get([]byte("22222"))
		fmt.Printf("v1 : %s\n", v1)
		fmt.Printf("v2 : %s\n", v2)
		return nil
	})
}
