package configs

var acessRoles = [][]string{
	{"A", ("person/get-by-id;person/get-all;person/insert;person/update;person/get-by-doc;person/reset-password;" +
		"user/get-by-id;user/get-all;user/insert;user/update;cbo/get-by-id;cbo/get-all;cbo/insert;" +
		"cbo/update;patient/get-by-id;patient/get-all;patient/insert;patient/update;doctor/get-by-id;doctor/get-all;" +
		"doctor/insert;doctor/update;services/get-by-id;services/get-all;services/insert;services/update;" +
		"quiz/get-by-id;quiz/get-all;quiz/insert;quiz/update;question/get-by-id;question/get-all;question/insert;question/update;question/update-bulk;" +
		"question/get-by-quiz;question/delete;anamnese/get-by-id;anamnese/get-all;anamnese/insert;anamnese/update;asking/get-by-id;asking/get-all;" +
		"asking/insert;asking/update;proximity/insert;proximity/get-by-id-person;proximity/get-by-id-patient;anamnesehasasking/get-by-anamnese-id;proximity/get-persons-by-id-patient/;" +
		"anamnesehasasking/get-by-asking-id;anamnesehasasking/get-all;anamnesehasasking/insert;proximityhasquiz/get-all;proximityhasquiz/get-by-id-quiz;" +
		"proximityhasquiz/get-by-id-patient;proximityhasquiz/get-by-id-person;proximityhasquiz/insert;proximityhasquiz/update;" +
		"patienthasdoctor/get-by-patient-id;patienthasdoctor/get-by-doctor-id;patienthasdoctor/get-all;patienthasdoctor/insert;anamnese/get-by-id-user-patient;" +
		"proximity/get-persons-by-id-patient;patient/reset-password")},

	{"U", ("person/get-by-id;person/get-all;person/insert;person/update;person/get-by-doc;person/reset-password;" +
		"user/get-by-id;user/get-all;user/insert;user/update;cbo/get-by-id;cbo/get-all;cbo/insert;" +
		"cbo/update;patient/get-by-id;patient/get-all;patient/insert;patient/update;doctor/get-by-id;doctor/get-all;" +
		"doctor/insert;doctor/update;services/get-by-id;services/get-all;services/insert;services/update;" +
		"quiz/get-by-id;quiz/get-all;quiz/insert;quiz/update;question/get-by-id;question/get-all;question/insert;question/update;question/update-bulk;" +
		"question/get-by-quiz;question/delete;anamnese/get-by-id;anamnese/get-all;anamnese/insert;anamnese/update;asking/get-by-id;asking/get-all;" +
		"asking/insert;asking/update;proximity/insert;proximity/get-by-id-person;proximity/get-by-id-patient;anamnesehasasking/get-by-anamnese-id;proximity/get-persons-by-id-patient/;" +
		"anamnesehasasking/get-by-asking-id;anamnesehasasking/get-all;anamnesehasasking/insert;proximityhasquiz/get-all;proximityhasquiz/get-by-id-quiz;" +
		"proximityhasquiz/get-by-id-patient;proximityhasquiz/get-by-id-person;proximityhasquiz/insert;proximityhasquiz/update;" +
		"patienthasdoctor/get-by-patient-id;patienthasdoctor/get-by-doctor-id;patienthasdoctor/get-all;patienthasdoctor/insert;anamnese/get-by-id-user-patient;" +
		"proximity/get-persons-by-id-patient")},

	{"P", ("user/get-by-id;user/get-all;" +
		"patient/get-by-id;patient/get-all;doctor/get-by-id;doctor/get-all;" +
		"quiz/get-by-id;quiz/get-all;quiz/update;question/get-by-id;question/get-all;question/update;question/update-bulk;" +
		"question/get-by-quiz;question/delete;anamnese/get-by-id;anamnese/get-all;asking/get-by-id;asking/get-all;" +
		"proximity/get-by-id-person;proximity/get-by-id-patient;anamnesehasasking/get-by-anamnese-id;" +
		"anamnesehasasking/get-by-asking-id;anamnesehasasking/get-all;proximityhasquiz/get-all;proximityhasquiz/get-by-id-quiz;" +
		"proximityhasquiz/get-by-id-patient;proximityhasquiz/get-by-id-person;" +
		"patienthasdoctor/get-by-patient-id;patienthasdoctor/get-all;anamnese/get-by-id-user-patient;proximity/get-persons-by-id-patient")},
}

func GetRoles() (roles [][]string) {
	return acessRoles
}

var _ = GetRoles()
