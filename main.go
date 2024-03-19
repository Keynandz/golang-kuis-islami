package main

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Question struct {
	ID       int
	Question string
	Options  []string
	Answer   string
}

var questions = []Question{
	{1, "Siapakah nama penemu angka 0 (nol) yang memberikan kontribusi besar dalam perkembangan matematika selama masa kejayaan Islam?", []string{"Al-Khwarizmi", "Ibnu Sina", "Al-Jazari", "Al-Battani"}, "Al-Khwarizmi"},
	{2, "Apa nama bangunan megah yang dibangun pada masa kejayaan Islam di Spanyol yang menjadi pusat ilmu pengetahuan dan budaya pada saat itu?", []string{"Alhambra", "Hagia Sophia", "Al-Aqsa", "Kubah Batu"}, "Alhambra"},
	{3, "Siapakah ilmuwan Muslim yang dikenal sebagai 'Bapak Optik' yang melakukan banyak penelitian dalam bidang optik pada masa kejayaan Islam?", []string{"Ibnu al-Haitham", "Al-Kindi", "Al-Razi", "Ibnu Khaldun"}, "Ibnu al-Haitham"},
	{4, "Apakah nama yang dikenal sebagai 'Madrasah Baitul Hikmah', sebuah pusat ilmu pengetahuan yang terkenal di Baghdad pada masa kejayaan Islam?", []string{"House of Wisdom", "Al-Azhar", "University of Qarawiyyin", "Nalanda"}, "House of Wisdom"},
	{5, "Siapakah yang merupakan pelopor dalam pengembangan sistem angka desimal dan pengenalan konsep nol (nol) pada masa kejayaan Islam?", []string{"Al-Khwarizmi", "Ibnu Sina", "Al-Jazari", "Al-Battani"}, "Al-Khwarizmi"},
	{6, "Kitab manakah yang ditulis oleh Ibnu Sina yang menjadi salah satu karya terpenting dalam sejarah filsafat kedokteran?", []string{"Al-Qanun fi al-Tibb", "Tafsir Ibnu Katsir", "Al-Tashrif", "Zubdat al-Tawarikh"}, "Al-Qanun fi al-Tibb"},
	{7, "Siapakah penulis Kitab Al-Jabr wal-Muqabalah yang merupakan salah satu karya penting dalam perkembangan matematika?", []string{"Al-Khwarizmi", "Ibnu Sina", "Al-Jazari", "Al-Battani"}, "Al-Khwarizmi"},
	{8, "Apakah nama kota yang menjadi pusat perdagangan dan ilmu pengetahuan pada masa kejayaan Islam di Jazirah Arab?", []string{"Baghdad", "Cairo", "Damaskus", "Kufah"}, "Baghdad"},
	{9, "Apa nama sistem pendidikan yang berkembang pesat pada masa kejayaan Islam yang mencakup studi agama, ilmu pengetahuan, dan filsafat?", []string{"Madrasah", "Maktab", "Hawza", "Madrassa"}, "Madrasah"},
	{10, "Kitab manakah yang ditulis oleh Ibnu Khaldun yang membahas sejarah dunia dan metodologi ilmiah dalam riset sejarah?", []string{"Al-Muqaddimah", "Al-Qanun fi al-Tibb", "Tafsir Ibnu Katsir", "Zubdat al-Tawarikh"}, "Al-Muqaddimah"},
	{11, "Apakah nama masjid yang menjadi simbol kejayaan arsitektur Islam dan dianggap sebagai salah satu bangunan paling indah di dunia?", []string{"Masjid Katedral Cordoba", "Masjid Hagia Sophia", "Masjid Al-Aqsa", "Masjid Al-Haram"}, "Masjid Katedral Cordoba"},
	{12, "Siapakah ilmuwan Muslim yang dikenal sebagai 'Bapak Kimia' yang melakukan banyak penelitian dalam bidang kimia pada masa kejayaan Islam?", []string{"Al-Razi", "Ibnu al-Haitham", "Al-Kindi", "Ibnu Khaldun"}, "Al-Razi"},
	{13, "Apakah nama sistem irigasi yang diperkenalkan oleh ilmuwan Muslim pada masa kejayaan Islam yang memiliki dampak besar dalam pertanian?", []string{"Qanat", "Aqueduct", "Siphon", "Reservoir"}, "Qanat"},
	{14, "Siapakah penulis Kitab Al-Jazari yang merupakan salah satu karya penting dalam bidang teknik dan ilmu pengetahuan pada masa kejayaan Islam?", []string{"Al-Jazari", "Al-Khwarizmi", "Ibnu Sina", "Al-Battani"}, "Al-Jazari"},
	{15, "Apa nama institusi pendidikan tertua yang didirikan di Fes, Maroko, dan masih berfungsi hingga sekarang?", []string{"University of Qarawiyyin", "Al-Azhar University", "House of Wisdom", "Nalanda University"}, "University of Qarawiyyin"},
	{16, "Apakah nama kota yang menjadi pusat ilmu pengetahuan dan kebudayaan pada masa kejayaan Islam di Spanyol?", []string{"Cordoba", "Granada", "Seville", "Toledo"}, "Cordoba"},
	{17, "Siapakah ilmuwan Muslim yang menemukan prinsip hukum bacaan pada cahaya yang membentuk dasar bagi banyak peralatan optik modern?", []string{"Ibnu al-Haitham", "Al-Razi", "Ibnu Khaldun", "Al-Kindi"}, "Ibnu al-Haitham"},
	{18, "Apakah nama bangunan megah yang dibangun pada masa kejayaan Islam di India yang menjadi simbol kebesaran arsitektur Islam?", []string{"Taj Mahal", "Qutub Minar", "Hawa Mahal", "Red Fort"}, "Taj Mahal"},
	{19, "Siapakah ilmuwan Muslim yang menulis Kitab Al-Tashrif, yang membahas bidang kedokteran dan bedah pada masa kejayaan Islam?", []string{"Al-Zahrawi", "Ibnu Sina", "Al-Biruni", "Al-Kindi"}, "Al-Zahrawi"},
	{20, "Siapakah penulis karya 'Al-Tashrih al-Badan' yang merupakan salah satu karya penting dalam bidang anatomi pada masa kejayaan Islam?", []string{"Al-Zahrawi", "Ibnu Sina", "Al-Biruni", "Al-Kindi"}, "Al-Zahrawi"},
}


func mainPage(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	return tmpl.Execute(c.Response().Writer, questions)
}

func checkAnswer(c echo.Context) error {
	type Answer struct {
		QuestionID int    `json:"question_id"`
		Answer     string `json:"answer"`
	}

	var ans Answer
	if err := c.Bind(&ans); err != nil {
		return err
	}

	var correct bool
	for _, q := range questions {
		if q.ID == ans.QuestionID && q.Answer == ans.Answer {
			correct = true
			break
		}
	}

	return c.JSON(http.StatusOK, map[string]bool{"correct": correct})
}

func main() {
	e := echo.New()

	e.Static("/static", "static")

	e.GET("/", mainPage)
	e.POST("/check", checkAnswer)

	e.Start(":8080")
}
