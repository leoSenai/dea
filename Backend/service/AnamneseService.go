package service

import (
	"api/models"
	"api/repository"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/signintech/gopdf"
)

func GetAnamneseById(id int64) (anamnese models.Anamnese, err error) {
	anamnese, err = repository.GetAnamneseById(int64(id))
	return anamnese, err
}

func GetAnamneseByIdUserPatient(idUser int64, idPatient int64) (anamnese models.Anamnese, err error) {
	anamnese, err = repository.GetAnamneseByIdUserPatient(idUser, idPatient)
	return anamnese, err
}

func GetAllAnamnese() (anamnese []models.Anamnese, err error) {
	anamneses, err := repository.GetAllAnamnese()
	return anamneses, err
}

func PostAnamnese(anamnesePost models.Anamnese) (anamneseBack models.Anamnese, err error) {
	anamneseBack, err = repository.PostAnamnese(anamnesePost)
	return anamneseBack, err
}

func PutAnamnese(anamnesePut models.Anamnese) (anamneseBack models.Anamnese, err error) {
	anamneseBack, err = repository.PutAnamnese(anamnesePut)
	return anamneseBack, err
}

func GetMonthPTBR(month_int int) (month string) {
	switch month_int {
	case 1:
		return "janeiro"
	case 2:
		return "fevereiro"
	case 3:
		return "março"
	case 4:
		return "abril"
	case 5:
		return "maio"
	case 6:
		return "junho"
	case 7:
		return "julho"
	case 8:
		return "agosto"
	case 9:
		return "setembro"
	case 10:
		return "outubro"
	case 11:
		return "novembro"
	case 12:
		return "dezembro"
	}
	return
}

func GetReport(anamnese models.Anamnese, grau int) (anamnese_file http.File, err error) {

	patient, _ := GetPatientById(int64(anamnese.IdPatient))
	user, _ := GetUserById(int64(anamnese.IdUser))
	user_cbo, _ := GetCboById(int64(user.IdCbo))
	//montar pdf
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	//configura fontes

	err = pdf.AddTTFFont("ArialUni", "fonts/arial.ttf")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	err = pdf.AddTTFFont("ArialBold", "fonts/arialbd.ttf")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	err = pdf.SetFont("ArialBold", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	// ---------------------- INICIO CABEÇALHO -------------------------

	//Imagem cabecalho
	pdf.Image("../Frontend/public/logo.png", 10, 10, &gopdf.Rect{
		W: 70,
		H: 70,
	})

	//Cabeçalho Titulo clinica
	pdf.SetX(85)
	pdf.SetY(pdf.GetY() + 2)
	pdf.Cell(&gopdf.Rect{
		W: 60,
		H: 30,
	}, "Clínica Motivar")

	//Cabeçalho paciente
	pdf.SetX(85)
	patient_name_y := pdf.GetY() + 15
	pdf.SetY(patient_name_y)
	pdf.Cell(&gopdf.Rect{
		W: 60,
		H: 30,
	}, "Paciente: ")

	//Cabeçalho cns
	pdf.SetX(85)
	patient_cns_y := pdf.GetY() + 14
	pdf.SetY(patient_cns_y)
	pdf.Cell(&gopdf.Rect{
		W: 60,
		H: 30,
	}, "CNS: ")

	//Cabeçalho idade
	pdf.SetX(85)
	patient_age_y := pdf.GetY() + 14
	pdf.SetY(patient_age_y)
	pdf.Cell(&gopdf.Rect{
		W: 60,
		H: 30,
	}, "Idade: ")

	//Cabeçalho emissão
	pdf.SetX(85)
	emmit_datetime_y := pdf.GetY() + 15
	pdf.SetY(emmit_datetime_y)
	pdf.Cell(&gopdf.Rect{
		W: 60,
		H: 30,
	}, "Emitido em: ")
	err = pdf.SetFont("ArialUni", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	//Set patient name
	pdf.SetX(160)
	pdf.SetY(patient_name_y)
	pdf.Cell(&gopdf.Rect{
		W: 60,
		H: 30,
	}, patient.Name)

	//Set patient CNS
	cns_formated := patient.Cns
	cns_formated = (cns_formated[:4] + "-" + cns_formated[4:])
	cns_formated = (cns_formated[:9] + "-" + cns_formated[9:])
	cns_formated = (cns_formated[:14] + "-" + cns_formated[14:])
	pdf.SetX(160)
	pdf.SetY(patient_cns_y)
	pdf.Cell(&gopdf.Rect{
		W: 60,
		H: 30,
	}, cns_formated)

	//Set patient Age
	patient_age_date, _ := time.Parse("2006-01-02", patient.BornDate)
	fmt.Print(patient_age_date)
	patient_age_years := (time.Now().Year() - patient_age_date.Year())
	patient_age_months := (time.Now().Month() - patient_age_date.Month())
	pdf.SetX(160)
	pdf.SetY(patient_age_y)
	pdf.Cell(&gopdf.Rect{
		W: 60,
		H: 30,
	}, strconv.Itoa(patient_age_years)+" anos e "+strconv.Itoa(int(patient_age_months))+" meses")

	//Set date time emmited
	current_time := time.Now()
	year, month, day := current_time.Date()
	emmit_date := strconv.Itoa(day) + "/" + strconv.Itoa(int(month)) + "/" + strconv.Itoa(year) + " " + strconv.Itoa(current_time.Hour()) + ":" + strconv.Itoa(current_time.Minute())
	pdf.SetX(160)
	pdf.SetY(emmit_datetime_y)
	pdf.Cell(&gopdf.Rect{
		W: 60,
		H: 30,
	}, emmit_date)

	//Linha separação cabeçalho
	pdf.SetLineWidth(1)
	pdf.Line(5, 86, 590, 86)
	pdf.SetLineWidth(1)
	pdf.Line(5, 5, 590, 5)
	pdf.SetLineWidth(1)
	pdf.Line(5, 5, 5, 86)
	pdf.SetLineWidth(1)
	pdf.Line(590, 5, 590, 86)

	// ----------------------- FIM CABEÇALHO --------------------------

	// ------------------------ COMEÇO CORPO  -------------------------

	//Corpo declaração
	pdf.SetX(45)
	declaration_corpo_y := 120.0
	pdf.SetY(declaration_corpo_y)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, "A clínica Motivar, portadora do CNPJ")

	err = pdf.SetFont("ArialBold", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	pdf.SetX(246)
	pdf.SetY(declaration_corpo_y)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, "02.679.048/0001-01")

	err = pdf.SetFont("ArialUni", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	pdf.SetX(356)
	pdf.SetY(declaration_corpo_y)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, "atesta hoje, dia ")

	err = pdf.SetFont("ArialBold", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	pdf.SetX(441)
	pdf.SetY(declaration_corpo_y)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, strconv.Itoa(current_time.Day())+" de "+GetMonthPTBR(int(current_time.Month()))+" de "+strconv.Itoa(current_time.Year()))

	err = pdf.SetFont("ArialUni", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	pdf.SetX(26)
	declaration_corpo_y = 140.0
	pdf.SetY(declaration_corpo_y)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, "que pelo presente processo de diagnóstico do DEA (Diagnóstico do Espectro Autista), atuado pelo  pro-")

	pdf.SetX(26)
	declaration_corpo_y = 160.0
	pdf.SetY(declaration_corpo_y)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, "fissional ")

	err = pdf.SetFont("ArialBold", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	pdf.SetX(74)
	pdf.SetY(declaration_corpo_y)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, user.Name)

	err = pdf.SetFont("ArialUni", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	calc_after_cbo_x := 75.0 + (6.0 * float64(len(user.Name)))
	pdf.SetX(calc_after_cbo_x)
	pdf.SetY(declaration_corpo_y)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, " exercente da ocupação de")

	err = pdf.SetFont("ArialBold", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	user_cbo_desc_y := calc_after_cbo_x + 147.0
	pdf.SetX(user_cbo_desc_y)
	pdf.SetY(declaration_corpo_y)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, strings.ToLower(user_cbo.Desc))

	err = pdf.SetFont("ArialUni", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	pdf.SetX(user_cbo_desc_y + (6.0 * float64(len(user_cbo.Desc))) - 2.0)
	pdf.SetY(declaration_corpo_y)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, ", concluí-se que:")

	pdf.SetX(26)
	pdf.SetY(200)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, "O paciente ")

	err = pdf.SetFont("ArialBold", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	pdf.SetX(88)
	pdf.SetY(200)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, patient.Name)

	err = pdf.SetFont("ArialUni", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	calc_after_patient_name_x := 89 + (7.0 * float64(len(patient.Name)))
	pdf.SetX(calc_after_patient_name_x)
	pdf.SetY(200)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, "portador do CPF")

	err = pdf.SetFont("ArialBold", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	calc_after_patient_cpf := calc_after_patient_name_x + 91
	pdf.SetX(calc_after_patient_cpf)
	pdf.SetY(200)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, patient.Cpf)

	indicativo_str := ""
	if anamnese.Indicative == 0 {
		indicativo_str = "não dispõe:"
		pdf.SetFillColor(226, 44, 44)
	} else {
		indicativo_str = "dispõe:"
		pdf.SetFillColor(55, 178, 63)
	}

	pdf.SetX(calc_after_patient_cpf + 89)
	pdf.SetY(200)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, indicativo_str)

	err = pdf.SetFont("ArialUni", "", 12)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	pdf.SetLineWidth(1)
	pdf.Line(200, 740, 415, 740)

	pdf.SetFillColor(0, 0, 0)

	pdf.SetX(250)
	pdf.SetY(760)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, "Assinatura do "+user.Name)

	indicativo_desc := ""

	if anamnese.Indicative == 1 {
		indicativo_desc = "no grau " + GetGrau(grau)
	}

	pdf.SetX(26)
	pdf.SetY(230)
	pdf.Cell(&gopdf.Rect{
		W: 180,
		H: 60,
	}, "- Do Transtorno do Espectro Autista (TEA) "+indicativo_desc)

	// -------------------------- FIM LAUDO ---------------------------

	id_random := uuid.New().String()

	report_path := "reports/report" + id_random + ".pdf"
	pdf.WritePdf(report_path)

	return anamnese_file, err
}

func GetGrau(grau int) (graustr string) {
	if grau == 1 {
		graustr = "leve."
	} else if grau == 2 {
		graustr = "moderado."
	} else if grau == 3 {
		graustr = "severo."
	}
	return graustr
}
