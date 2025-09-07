package entity

import "time"

type User struct {
	DepartmentIds []string  // 部门
	HiredDate     time.Time // 入职时间
	CreatedAt     time.Time // 创建时间
	UpdatedAt     time.Time // 更新时间
	Name          string    // 姓名
	Nickname      string    // 别名
	Mobile        string    // 手机号
	Email         string    // 邮箱
	LeaderID      string    // 直属上级
	JobTitle      string    // 员工职位
	HashPassword  string    // 密码哈希
	JobNumber     int       // 员工工号
	UserID        int       // 用户ID
}
