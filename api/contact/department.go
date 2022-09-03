package contact

type DepartmentIdType string

const (
	DepartmentIdTypeDepartmentId     DepartmentIdType = "department_id"
	DepartmentIdTypeOpenDepartmentId                  = "open_department_id"
)

type DepartmentI18nName struct {
	ZhCn string `json:"zh_cn"`
	EnUs string `json:"en_us"`
	JaJp string `json:"ja_jp"`
}
