package main

import (
	"fmt"
	"os"

	"strconv"

	"math/rand"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tealeg/xlsx"
)

type StudentsInfo struct {
	gorm.Model
	NowClass                 int    `gorm:"type:integer"`
	Number                   string `gorm:"type:varchar(255)"`
	OldClass                 int    `gorm:"type:integer"`
	Name                     string `gorm:"type:varchar"`
	Gender                   string `gorm:"type:varchar"`
	MasterSubject            string `gorm:"type:varchar"`
	SeniorHighSchoolScore    string `gorm:"type:varchar"`
	SeniorHighSchoolRank     int    `gorm:"type:integer"`
	SeniorOneScoreBefore     string `gorm:"type:varchar"`
	SeniorOneScoreBeforeRank int    `gorm:"type:integer"`
	SeniorOneScoreLast       string `gorm:"type:varchar"`
	SeniorOneScoreLastRank   int    `gorm:"type:integer"`
	TotalRank                int    `gorm:"type:integer"`
	Chinese                  string `gorm:"type:varchar"`
	Math                     string `gorm:"type:varchar"`
	English                  string `gorm:"type:varchar"`
	Physics                  string `gorm:"type:varchar"`
	Chemistry                string `gorm:"type:varchar"`
	Microorganism            string `gorm:"type:varchar"`
	Politics                 string `gorm:"type:varchar"`
	History                  string `gorm:"type:varchar"`
	Geography                string `gorm:"type:varchar"`
}

type studentInfoBasic struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	ID     uint   `json:"id"`

	NowClass      int    `json:"now_class"`
	OldClass      int    `json:"old_class"`
	Gender        string `json:"gender"`
	MasterSubject string `json:"master_subject"`
}

//SeniorHighSchoolScore    string     `json:"senior_high_school_score"`
//SeniorHighSchoolRank     string     `json:"senior_high_school_rank"`
//SeniorOneScoreBefore     string     `json:"senior_one_score_before"`
//SeniorOneScoreBeforeRank string     `json:"senior_one_score_before_rank"`
//SeniorOneScoreLast       string     `json:"senior_one_score_last"`
//SeniorOneScoreLastRank   string     `json:"senior_one_score_last_rank"`
//TotalRank                string     `json:"total_rank"`
//Chinese                  string     `json:"chinese"`
//Math                     string     `json:"math"`
//English                  string     `json:"english"`
//Physics                  string     `json:"physics"`
//Chemistry                string     `json:"chemistry"`
//Microorganism            string     `json:"microorganism"`
//Politics                 string     `json:"politics"`
//History                  string     `json:"history"`
//Geography                string     `json:"geography"`

func (s *StudentsInfo) Serialzer() *studentInfoBasic {
	return &studentInfoBasic{
		ID:            s.ID,
		Name:          s.Name,
		NowClass:      s.NowClass,
		OldClass:      s.OldClass,
		MasterSubject: s.MasterSubject,
		Gender:        s.Gender,
		Number:        s.Number,
	}
}

var DB *gorm.DB

func init() {

	if _, err := os.Stat("students.db"); os.IsExist(err) != true {
		os.Create("students.db")

	} else {
		os.Remove("students.db")
	}
	db, err := gorm.Open("sqlite3", "./students.db")
	if err != nil {
		fmt.Println(err, "sqlite3")
		return
	}

	DB = db
	DB.LogMode(true)
	DB.AutoMigrate(&StudentsInfo{})
}

func StartInit() {

	path, _ := os.Getwd()
	fullPath := path + "/students.xlsx"
	excelFileName := fullPath
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	for _, sheet := range xlFile.Sheets {
		fmt.Println(sheet.Name)
		for index, row := range sheet.Rows {

			if len(row.Cells) != 22 {
				return
			}

			fmt.Println(index, "index")
			nowClass, _ := strconv.Atoi(row.Cells[0].Value)
			oldClass, _ := strconv.Atoi(row.Cells[1].Value)
			number := row.Cells[2].Value
			name := row.Cells[3].Value
			gender := row.Cells[4].Value
			master := row.Cells[5].Value
			sScore := row.Cells[6].Value
			sHigh := row.Cells[7].Value
			sLast := row.Cells[8].Value
			sRank, _ := strconv.Atoi(row.Cells[9].Value)
			gHighRank, _ := strconv.Atoi(row.Cells[10].Value)
			gLastRank, _ := strconv.Atoi(row.Cells[11].Value)
			totalRank, _ := strconv.Atoi(row.Cells[12].Value)
			c := row.Cells[13].Value
			m := row.Cells[14].Value
			e := row.Cells[15].Value
			p := row.Cells[16].Value
			chm := row.Cells[17].Value
			sw := row.Cells[18].Value
			zz := row.Cells[19].Value
			ls := row.Cells[20].Value
			dl := row.Cells[21].Value

			tempData := StudentsInfo{
				NowClass:                 nowClass,
				OldClass:                 oldClass,
				Number:                   number,
				Name:                     name,
				Gender:                   gender,
				MasterSubject:            master,
				SeniorHighSchoolScore:    sScore,
				SeniorOneScoreBefore:     sHigh,
				SeniorOneScoreLast:       sLast,
				SeniorHighSchoolRank:     sRank,
				SeniorOneScoreBeforeRank: gHighRank,
				SeniorOneScoreLastRank:   gLastRank,
				TotalRank:                totalRank,
				Chinese:                  c,
				Math:                     m,
				English:                  e,
				Physics:                  p,
				Chemistry:                chm,
				Microorganism:            sw,
				Politics:                 zz,
				History:                  ls,
				Geography:                dl,
			}
			if index == 0 {
				fmt.Println(tempData)
			} else {
				if dbError := DB.Save(&tempData).Error; dbError != nil {
					return
				}
			}

		}
	}
}

// Todo 输入班级，数字，随机蹦出一位学员名称
// 1. 将数据导入 sqlite
// 2. 页面仿照 搜狗搜索
// 3. 随机抽取出一位学员

func main() {
	StartInit()
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"Message": "幸运系统",
		})
	})
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("./templates/**/*")

	r.GET("/attachments/new", nil)
	r.GET("/attachments/create", nil)
	r.GET("/attachments", nil)

	r.GET("/reports/class", randomMax)
	r.GET("/reports/all_class/:type", statisticClass)
	r.GET("/reports/all_students", allStudents)
	r.GET("/reports/single", searchStudentByNameOrNumber)

	r.Run(":8080")
}

var IDs []int

type queryParam struct {
	PerPage int `form:"per_page default=1"`
	Offset  int `form:"offset default=10"`
	Offline int `form:"offline"`
	Number  int `form:"class_number"`
}

func randomMax(c *gin.Context) {

	var param queryParam

	if err := c.ShouldBindQuery(&param); err != nil {
		fmt.Println(err)
		return
	}

	class := param.Number
	randomNumber := param.Offline
	if (class <= 0 || class > 10) || (randomNumber <= 0) {
		errText := "请再次确定班级(1-10) 或 可能缺席人数(小于班级总人数)"
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"errText": errText,
		})
		return

	}

	var collections []StudentsInfo
	if dbErr := DB.Raw(fmt.Sprintf("select * from students_infos where now_class = %d", class)).Scan(&collections).Error; dbErr != nil {
		fmt.Println(dbErr)
		return
	}
	//randL := rand.New(rand.NewSource(time.Now().Unix()))
	randomIDC := rand.Perm(len(collections))
	fmt.Println(randomIDC, len(randomIDC), "randomIDC")
	randomIDs := []int{}
	for _, v := range randomIDC {
		randomIDs = append(randomIDs, int(collections[v].ID))
		if len(randomIDs) >= len(collections) || len(randomIDs) >= randomNumber {
			break
		}

	}

	var results []StudentsInfo
	if dbErr := DB.Where("id in (?)", randomIDs).Find(&results).Error; dbErr != nil {
		fmt.Println(dbErr.Error())
		return
	}
	resultData := make([]studentInfoBasic, len(randomIDs))
	for index, one := range results {
		resultData[index] = *one.Serialzer()
	}
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"attachs": resultData,
		"count":   false,
	})
}

type class struct {
	NowClass int `json:"now_class"`
	Count    int `json:"count"`
}

func statisticClass(c *gin.Context) {

	var classStatistics []class
	if dbError := DB.Raw(`select now_class as now_class, count(now_class) as count from students_infos group by now_class`).Scan(&classStatistics).Error; dbError != nil {
		return
	}

	var oldStatistics []class
	if dbError := DB.Raw(`select old_class as now_class, count(old_class) as count from students_infos group by old_class`).Scan(&oldStatistics).Error; dbError != nil {
		return
	}

	//c.JSON(http.StatusOK, gin.H{
	//
	//	"All_Now_Statistic": classStatistics,
	//	"All_Old_Statistic": oldStatistics,
	//})
	if c.Param("type") == "old" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"count": true,
			"list":  oldStatistics,
		})
	} else if c.Param("type") == "new" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"count": true,
			"list":  classStatistics,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"count":   true,
			"errText": "Choose old or new ",
		})
	}

}

func allStudents(c *gin.Context) {
	var collections []StudentsInfo
	if dbError := DB.Raw("select * from students_infos").Scan(&collections).Error; dbError != nil {
		return
	}

	result := make([]studentInfoBasic, len(collections))
	for index, val := range collections {
		result[index] = *val.Serialzer()
	}
	//c.JSON(http.StatusOK, result)
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"count":   false,
		"attachs": result,
	})

}

type search struct {
	Name   string `form:"name"`
	Number string `form:"number"`
}

func searchStudentByNameOrNumber(c *gin.Context) {

	var searchParam search
	if err := c.ShouldBindQuery(&searchParam); err != nil {
		return
	}
	var collection []StudentsInfo
	query := DB
	if searchParam.Name != "" {
		query = query.Where("name like ?", "%"+searchParam.Name+"%")

	}
	if searchParam.Number != "" {
		query = query.Where("number like ?", "%"+searchParam.Number+"%")
	}

	if dbError := query.Find(&collection).Error; dbError != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"Message": "Not Found"})
		return
	}
	result := make([]studentInfoBasic, len(collection))
	for index, val := range collection {
		result[index] = *val.Serialzer()
	}
	//c.JSON(http.StatusOK, result)
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"attachs": result,
		"count":   false,
	})
}
