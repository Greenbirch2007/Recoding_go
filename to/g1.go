package main

import "image"

type cfg struct {

}

var cfg * image.Config

func init()  {
	cfg = new(config)
}


//NewConfig提供获取实例的方法

func NewConfig() *config  {
	return cfg
}


type config struct {

}


//全局变量
var cfg *config = new(config)


//NewConfig提供获取实例的方法

func NewConfig() *config  {
	return cfg

}