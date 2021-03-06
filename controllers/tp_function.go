package controllers

import (
	gvalid "ThingsPanel-Go/initialize/validate"
	"ThingsPanel-Go/models"
	"ThingsPanel-Go/services"
	response "ThingsPanel-Go/utils"
	valid "ThingsPanel-Go/validate"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
)

type TpFunctionController struct {
	beego.Controller
}

// 列表
func (TpFunctionController *TpFunctionController) List() {
	var TpFunctionService services.TpFunctionService
	isSuccess, d := TpFunctionService.GetFunctionList()
	if !isSuccess {
		response.SuccessWithMessage(1000, "查询失败", (*context2.Context)(TpFunctionController.Ctx))
	}
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(TpFunctionController.Ctx))
}

// 编辑
func (TpFunctionController *TpFunctionController) Edit() {
	TpFunctionValidate := valid.TpFunctionValidate{}
	err := json.Unmarshal(TpFunctionController.Ctx.Input.RequestBody, &TpFunctionValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(TpFunctionValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(TpFunctionValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(TpFunctionController.Ctx))
			break
		}
		return
	}
	if TpFunctionValidate.Id == "" {
		response.SuccessWithMessage(1000, "id不能为空", (*context2.Context)(TpFunctionController.Ctx))
	}
	var TpFunctionService services.TpFunctionService
	TpFunction := models.TpFunction{
		Id:           TpFunctionValidate.Id,
		FunctionName: TpFunctionValidate.FunctionName,
	}
	isSucess := TpFunctionService.EditFunction(TpFunction)
	if isSucess {
		response.SuccessWithDetailed(200, "success", TpFunction, map[string]string{}, (*context2.Context)(TpFunctionController.Ctx))
	} else {
		response.SuccessWithMessage(400, "编辑失败", (*context2.Context)(TpFunctionController.Ctx))
	}
}

// 新增
func (TpFunctionController *TpFunctionController) Add() {
	TpFunctionValidate := valid.TpFunctionValidate{}
	err := json.Unmarshal(TpFunctionController.Ctx.Input.RequestBody, &TpFunctionValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(TpFunctionValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(TpFunctionValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(TpFunctionController.Ctx))
			break
		}
		return
	}
	var TpFunctionService services.TpFunctionService
	TpFunction := models.TpFunction{
		Id:           TpFunctionValidate.Id,
		FunctionName: TpFunctionValidate.FunctionName,
	}
	isSucess, d := TpFunctionService.AddFunction(TpFunction)
	if isSucess {
		response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(TpFunctionController.Ctx))
	} else {
		response.SuccessWithMessage(400, "新增失败", (*context2.Context)(TpFunctionController.Ctx))
	}
}

// 删除
func (TpFunctionController *TpFunctionController) Delete() {
	TpFunctionValidate := valid.TpFunctionValidate{}
	err := json.Unmarshal(TpFunctionController.Ctx.Input.RequestBody, &TpFunctionValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(TpFunctionValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(TpFunctionValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(TpFunctionController.Ctx))
			break
		}
		return
	}
	if TpFunctionValidate.Id == "" {
		response.SuccessWithMessage(1000, "id不能为空", (*context2.Context)(TpFunctionController.Ctx))
	}
	var TpFunctionService services.TpFunctionService
	TpFunction := models.TpFunction{
		Id:           TpFunctionValidate.Id,
		FunctionName: TpFunctionValidate.FunctionName,
	}
	isSucess := TpFunctionService.DeleteFunction(TpFunction)
	if isSucess {
		response.SuccessWithMessage(200, "success", (*context2.Context)(TpFunctionController.Ctx))
	} else {
		response.SuccessWithMessage(400, "编辑失败", (*context2.Context)(TpFunctionController.Ctx))
	}
}
