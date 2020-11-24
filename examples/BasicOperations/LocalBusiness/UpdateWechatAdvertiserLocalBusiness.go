/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
)

type WechatAdvertiserLocalBusinessUpdateExample struct {
	TAds                                    *ads.SDKClient
	AccessToken                             string
	HeadImage                               *os.File
	Name                                    string
	Description                             string
	ContactPerson                           string
	ContactPersonMobile                     string
	ContactPersonCardId                     string
	Corporation                             string
	CorporationLicence                      string
	IndustryId                              int64
	AccountId                               int64
	WechatAdvertiserLocalBusinessUpdateOpts *api.WechatAdvertiserLocalBusinessUpdateOpts
}

func (e *WechatAdvertiserLocalBusinessUpdateExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	file, err := os.Open("YOUR FILE PATH")
	if err != nil {
		e.HeadImage = file
	}
	e.Name = "name_example"
	e.Description = "description_example"
	e.ContactPerson = "contactPerson_example"
	e.ContactPersonMobile = "contactPersonMobile_example"
	e.ContactPersonCardId = "contactPersonCardId_example"
	e.Corporation = "corporation_example"
	e.CorporationLicence = "corporationLicence_example"
	e.IndustryId = 789
	e.AccountId = 789
	e.WechatAdvertiserLocalBusinessUpdateOpts = &api.WechatAdvertiserLocalBusinessUpdateOpts{}
}

func (e *WechatAdvertiserLocalBusinessUpdateExample) RunExample() (interface{}, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.WechatAdvertiserLocalBusiness().Update(ctx, e.HeadImage, e.Name, e.Description, e.ContactPerson, e.ContactPersonMobile, e.ContactPersonCardId, e.Corporation, e.CorporationLicence, e.IndustryId, e.AccountId, e.WechatAdvertiserLocalBusinessUpdateOpts)
}

func main() {
	e := &WechatAdvertiserLocalBusinessUpdateExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}
