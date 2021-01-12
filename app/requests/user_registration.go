package requests

import (
	"goblog/app/models/user"

	"github.com/thedevsaddam/govalidator"
)

func ValidateRegistrationForm(data user.User) map[string][]string {

	// 创建规则
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:3,20"},
		"email":            []string{"required", "email", "min:4", "max:30"},
		"password":         []string{"required", "min:6"},
		"password_comfirm": []string{"required"},
	}

	messages := govalidator.MapData{
		"name":             []string{"required:用户名不能为空", "alpha_num:用户名应为数字和字母", "between:用户名长度为3-20"},
		"email":            []string{"required:邮箱不能为空", "email:请输入正确的邮箱", "min:邮箱长度为4-30", "max:邮箱长度为4-30"},
		"password":         []string{"required:请输入密码", "min:密码长度必须大于6"},
		"password_comfirm": []string{"required:确认密码为必填项"},
	}

	// 配置信息
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	errs := govalidator.New(opts).ValidateStruct()

	if data.Password != data.PasswordComfirm {
		errs["password_comfirm"] = append(errs["password_comfirm"], "两次密码输入不一致")
	}

	return errs
}
