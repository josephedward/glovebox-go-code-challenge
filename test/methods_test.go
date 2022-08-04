package test

import (
	"fmt"
	"goscraper/acloud"
	"goscraper/local"
	"goscraper/proxy"
	"testing"

)



func TestMethods(t *testing.T) {
	//load env credentials from .env file
	login, err := local.LoadEnv()
	local.PanicIfErr(err)
	fmt.Println("login : ", login)
	t.Log("login : ", login)
	
	//connect to website
	connect, err := proxy.Login(login)
	local.PanicIfErr(err)
	fmt.Println("connect : ", connect)
	t.Log("connect : ", connect)

	//scrape credentials
	vals, err := acloud.Sandbox(connect, login.Url)
	local.PanicIfErr(err)
	fmt.Println("vals : ", vals)
	t.Log("vals : ", vals)

	//copy credentials to clipboard
	creds, err := acloud.Copy(vals)
	local.PanicIfErr(err)
	fmt.Println("creds : ", creds.User)
	t.Log("creds : ", creds.User)

keys, KeyVals := acloud.KeyVals(creds)

	//create policies with map
	policies, err := proxy.Policies(keys, keyVals)
	local.PanicIfErr(err)
	fmt.Println("policies : ", policies)
	t.Log("policies : ", policies)

	//download text file of policies
	err = proxy.DocumentDownload("creds", policies)
	local.PanicIfErr(err)
	fmt.Println("Document Downloaded")
	t.Log("Document Downloaded")

	//create LocalCreds from creds
	localCreds, err := local.CreateLocalCreds(creds.User, creds.KeyID, creds.AccessKey)
	local.PanicIfErr(err)
	fmt.Println("localCreds : ", localCreds)
	t.Log("localCreds : ", localCreds)

	//append aws creds to .aws/credentials file
	err = local.AppendAwsCredentials(localCreds)
	local.PanicIfErr(err)
	fmt.Println("aws credentials appended")
	t.Log("aws credentials appended")


}
