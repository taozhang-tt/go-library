package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)


//访问谷歌退货查询接口的证书
var CA_PRIVATE_PEM = `
-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC2A5N1aLixCqfy
xieYWT98LKmpIqa/3sqguBLRMtnNnp/dgcxof3+2MvlUqX2xs8ArMnnvkO+NLZky
0qmklNs93JYix+LorMYVp+Jx8YNzayyPXitR9DRfJKQrPsRdoOMd1tORrmWsk2Sf
n97hmw4rqs42BvSDOXBN6OpaP5gGN4XipFEN/StXXWeojgBJ7drk0JSuFn2EJ4kU
4vKtIWF1PCpX/Ul86231R6MSQ0yk5U/Lu/CKhDkpAJdMfEDDi5cGn8nMGVzdn7ZC
9BC2nw9SuXrtfYQUrTa1fmb91wUUPZamHqI4bGjzUNn0AfdOrXf4kSfBeyqXq0eo
8I0WCtlXAgMBAAECggEBAJNAOV/XJPQ2BsGmBgp+yZmQGII5AmZfu9ZilPaiCQsM
tZDinZg87fldK9GOfZ3yKhnIcFWcD/FLXpF7VLtNokFd4CirRauve2GxvMUp+oq5
vgcOzEU1J9mBLz+9O+fzbUqKrgdB7Ae+Br4M/KfQZZp2SPX0koRHR9AACviQUlFT
4skyT7RR+qfOpdvs7GfT63rx/TJx9xaLwffXslCCk9EZn3JqVG9pqlje6XRE/xbO
ru8h95SgChLuBHgTXVTZWcINAfexWSd5vzI3et7aiJ18AhhHgrt1Bp++ytiNc/pN
rficsj45pDEKoo393LwZuYz84LWQGOczUsCQoD4oqcECgYEA7+PSeAjDkB/Zjpr9
0iO7xrSThFp0PNpFW0wbp4QdseS15vU8wB/NAAcveeYNd3ODRjWPhn/ghX7PwPjU
vfZhG7Y+5ddGqphKeg8MsZEbnRPsSkKXgRpwmWwsvIGz0y0GiPeT0HbusavrzB4r
r22vT+BYsGwd9Qr6spzsjvPBtI8CgYEAwjzAsWpD0/1k/we3lAE6YxHyqujaiz/v
dsqHKZeoLVDrE65p5uP0kZIKUje/drFsNQ1uBz2cCmgUkLEOGOeX111m1sMJRarP
q23esdXlhcpbjTxSNd+ay88O1C+T1BlOYf5LQ6cNLxsiShIrmDGxIGUOn/7utHY1
MsfSn6OVwrkCgYEA0ae94TRfno2HeU/SN58NVKurJYhcgHaPGSyaGrynqzrlE5YP
YyYn5zdgcdvytwwYlfrnljgaxPFApzmRzPgQhMSxkfSkiSISLT0L7gCRLqYj4/7x
JF80O5JEQvfkbtKyHHCsGgxnrpY/vaQ8r9Rf3KKJQJ8tS2UuF354bLNy/tcCgYBJ
24AJ9jZEbZ3xu6V6id0BvtlfU5uGR5FuwiQTFK/GS3aXzUJHoXZlw/pYuQn8wAo9
QhYMesjSzDDFtA2AOs/p/IIWN6NW8lR1Axoi1QqpjQCy+7Tm91eNF/aziEPV6ql7
aY6E8wh4WQD42V/qRzrq7oJZJD7wL+KapJDaUwGnQQKBgGvQPD/2TGshvJnr8WPa
c2K9pTgNpHQeTCXsWXj2mvYZAUJ98kcaYzcJ17ZpnkTS1f783FrtZAktd9DXMM4l
EjDff2qAK5JL9ZMefEJIv6OBa1wpj4wvl4mu6w23IjzfwzmFBUyKCDg3N6rmmZk4
rUcVar2klot2ac0xU4s0tgEi
-----END PRIVATE KEY-----
	`
func main() {
    GetVoidedOrderFromGoogle("com.joyyou.sdk.facebook.test", "lb_0.09", "lgkihllcjdpindgjopbmdhei.AO-J1OwftUsQ2ZoX_7Sk9M-JuWtr_WaTeHylLjlOooQyNtojojrFpiXvf6vH0X-ZhWClMwC1k9tqmnRirsvgXCgbd-8P5WhAQi-aF2yEHzFg-IDwBmsjrb4")
}

//通过谷歌 api 获取已退货的订单
func GetVoidedOrderFromGoogle(packageName, productId, token string) (err error) {
	var (
		googleApi             = "https://androidpublisher.googleapis.com/androidpublisher/v3/applications/"+packageName+"/purchases/products/"+productId+"/tokens/"+token+":acknowledge"
		resp                  *http.Response
		body                  []byte
	)
	conf := &jwt.Config{
		Email:      "test-581@api-6696778595688237216-170767.iam.gserviceaccount.com",
		PrivateKey: []byte(CA_PRIVATE_PEM),
		Scopes: []string{
			"https://www.googleapis.com/auth/androidpublisher",
		},
		TokenURL: google.JWTTokenURL,
	}

	client := conf.Client(oauth2.NoContext)
	req, _ := http.NewRequest(http.MethodPost, googleApi, nil)
	if resp, err = client.Do(req); err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(resp.StatusCode)
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println("err: ", err)
		return
	}
    fmt.Println(string(body))
	return
}
