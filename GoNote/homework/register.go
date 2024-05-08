package homework

import "fmt"

// 註冊流程
func Register() {
	for {
		var username, pwd, cpwd string
		fmt.Println("歡迎註冊")
		fmt.Println("請輸入用戶名:")
		fmt.Scanf("%s\n", &username)
		fmt.Println("請輸入密碼:")
		fmt.Scanf("%s\n", &pwd)
		fmt.Println("請再次輸入密碼:")
		fmt.Scanf("%s\n", &cpwd)
		if username == "" || pwd == "" || cpwd == "" {
			fmt.Println("用戶名或密碼不能為空")
			continue
		}
		if pwd != cpwd {
			fmt.Println("兩次密碼不一致")
			continue
		}
		fmt.Println("註冊成功")
		break
	}
}
