package entity

import "gorm.io/gorm"

type Employees struct { 
	gorm.Model 
	Name   string   `valid:"stringlength(2|80)~name must be 2-80"`// ชื่อต้องยาว 2-80 ตัวอักษร 
	Salary    float64  `valid:"float,range(15000|200000)~salary must be 15000-200000"`// เงินเดือนต้องอยู่ในช่วง 15000-200000 
	EmployeeCode string   `valid:"matches(^[HR]-{4}$)"`// รหัสพนักงานต้องเป็นอักษรอังกฤษตัวใหญ่ 2 ตัว ตามด้วย "-" และตัวเลข 4 ตัว (เช่น "HR-1024") 
	} 