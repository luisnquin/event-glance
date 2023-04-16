package currency

import "regexp"

// ISO 4217 currency codes.
const (
	// United Arab Emirates Dirham
	AED = "AED"
	// Afghan Afghani
	AFN = "AFN"
	// Albanian Lek
	ALL = "ALL"
	// Armenian Dram
	AMD = "AMD"
	// Netherlands Antillean Guilder
	ANG = "ANG"
	// Angolan Kwanza
	AOA = "AOA"
	// Argentine Peso
	ARS = "ARS"
	// Australian Dollar
	AUD = "AUD"
	// Aruban Florin
	AWG = "AWG"
	// Azerbaijani Manat
	AZN = "AZN"
	// Bosnia-Herzegovina Convertible Mark
	BAM = "BAM"
	// Barbadian Dollar
	BBD = "BBD"
	// Bangladeshi Taka
	BDT = "BDT"
	// Bulgarian Lev
	BGN = "BGN"
	// Bahraini Dinar
	BHD = "BHD"
	// Burundian Franc
	BIF = "BIF"
	// Bermudan Dollar
	BMD = "BMD"
	// Brunei Dollar
	BND = "BND"
	// Bolivian Boliviano
	BOB = "BOB"
	// Brazilian Real
	BRL = "BRL"
	// Bahamian Dollar
	BSD = "BSD"
	// Bitcoin
	BTC = "BTC"
	// Bhutanese Ngultrum
	BTN = "BTN"
	// Botswanan Pula
	BWP = "BWP"
	// New Belarusian Ruble
	BYN = "BYN"
	// Belarusian Ruble
	BYR = "BYR"
	// Belize Dollar
	BZD = "BZD"
	// Canadian Dollar
	CAD = "CAD"
	// Congolese Franc
	CDF = "CDF"
	// Swiss Franc
	CHF = "CHF"
	// Chilean Unit of Account (UF)
	CLF = "CLF"
	// Chilean Peso
	CLP = "CLP"
	// Chinese Yuan
	CNY = "CNY"
	// Colombian Peso
	COP = "COP"
	// Costa Rican Colon
	CRC = "CRC"
	// Cuban Convertible Peso
	CUC = "CUC"
	// Cuban Peso
	CUP = "CUP"
	// Cape Verdean Escudo
	CVE = "CVE"
	// Czech Republic Koruna
	CZK = "CZK"
	// Djiboutian Franc
	DJF = "DJF"
	// Danish Krone
	DKK = "DKK"
	// Dominican Peso
	DOP = "DOP"
	// Algerian Dinar
	DZD = "DZD"
	// Egyptian Pound
	EGP = "EGP"
	// Eritrean Nakfa
	ERN = "ERN"
	// Ethiopian Birr
	ETB = "ETB"
	// Euro
	EUR = "EUR"
	// Fijian Dollar
	FJD = "FJD"
	// Falkland Islands Pound
	FKP = "FKP"
	// British Pound Sterling
	GBP = "GBP"
	// Georgian Lari
	GEL = "GEL"
	// Guernsey Pound
	GGP = "GGP"
	// Ghanaian Cedi
	GHS = "GHS"
	// Gibraltar Pound
	GIP = "GIP"
	// Gambian Dalasi
	GMD = "GMD"
	// Guinean Franc
	GNF = "GNF"
	// Guatemalan Quetzal
	GTQ = "GTQ"
	// Guyanaese Dollar
	GYD = "GYD"
	// Hong Kong Dollar
	HKD = "HKD"
	// Honduran Lempira
	HNL = "HNL"
	// Croatian Kuna
	HRK = "HRK"
	// Haitian Gourde
	HTG = "HTG"
	// Hungarian Forint
	HUF = "HUF"
	// Indonesian Rupiah
	IDR = "IDR"
	// Israeli New Sheqel
	ILS = "ILS"
	// Manx pound
	IMP = "IMP"
	// Indian Rupee
	INR = "INR"
	// Iraqi Dinar
	IQD = "IQD"
	// Iranian Rial
	IRR = "IRR"
	// Icelandic Kr\u00f3na
	ISK = "ISK"
	// Jersey Pound
	JEP = "JEP"
	// Jamaican Dollar
	JMD = "JMD"
	// Jordanian Dinar
	JOD = "JOD"
	// Japanese Yen
	JPY = "JPY"
	// Kenyan Shilling
	KES = "KES"
	// Kyrgystani Som
	KGS = "KGS"
	// Cambodian Riel
	KHR = "KHR"
	// Comorian Franc
	KMF = "KMF"
	// North Korean Won
	KPW = "KPW"
	// South Korean Won
	KRW = "KRW"
	// Kuwaiti Dinar
	KWD = "KWD"
	// Cayman Islands Dollar
	KYD = "KYD"
	// Kazakhstani Tenge
	KZT = "KZT"
	// Laotian Kip
	LAK = "LAK"
	// Lebanese Pound
	LBP = "LBP"
	// Sri Lankan Rupee
	LKR = "LKR"
	// Liberian Dollar
	LRD = "LRD"
	// Lesotho Loti
	LSL = "LSL"
	// Lithuanian Litas
	LTL = "LTL"
	// Latvian Lats
	LVL = "LVL"
	// Libyan Dinar
	LYD = "LYD"
	// Moroccan Dirham
	MAD = "MAD"
	// Moldovan Leu
	MDL = "MDL"
	// Malagasy Ariary
	MGA = "MGA"
	// Macedonian Denar
	MKD = "MKD"
	// Myanma Kyat
	MMK = "MMK"
	// Mongolian Tugrik
	MNT = "MNT"
	// Macanese Pataca
	MOP = "MOP"
	// Mauritanian Ouguiya
	MRO = "MRO"
	// Mauritian Rupee
	MUR = "MUR"
	// Maldivian Rufiyaa
	MVR = "MVR"
	// Malawian Kwacha
	MWK = "MWK"
	// Mexican Peso
	MXN = "MXN"
	// Malaysian Ringgit
	MYR = "MYR"
	// Mozambican Metical
	MZN = "MZN"
	// Namibian Dollar
	NAD = "NAD"
	// Nigerian Naira
	NGN = "NGN"
	// Nicaraguan Córdoba
	NIO = "NIO"
	// Norwegian Krone
	NOK = "NOK"
	// Nepalese Rupee
	NPR = "NPR"
	// New Zealand Dollar
	NZD = "NZD"
	// Omani Rial
	OMR = "OMR"
	// Panamanian Balboa
	PAB = "PAB"
	// Peruvian Nuevo Sol
	PEN = "PEN"
	// Papua New Guinean Kina
	PGK = "PGK"
	// Philippine Peso
	PHP = "PHP"
	// Pakistani Rupee
	PKR = "PKR"
	// Polish Zloty
	PLN = "PLN"
	// Paraguayan Guarani
	PYG = "PYG"
	// Qatari Rial
	QAR = "QAR"
	// Romanian Leu
	RON = "RON"
	// Serbian Dinar
	RSD = "RSD"
	// Russian Ruble
	RUB = "RUB"
	// Rwandan Franc
	RWF = "RWF"
	// Saudi Riyal
	SAR = "SAR"
	// Solomon Islands Dollar
	SBD = "SBD"
	// Seychellois Rupee
	SCR = "SCR"
	// Sudanese Pound
	SDG = "SDG"
	// Swedish Krona
	SEK = "SEK"
	// Singapore Dollar
	SGD = "SGD"
	// Saint Helena Pound
	SHP = "SHP"
	// Sierra Leonean Leone
	SLE = "SLE"
	// Sierra Leonean Leone
	SLL = "SLL"
	// Somali Shilling
	SOS = "SOS"
	// Surinamese Dollar
	SRD = "SRD"
	// Sao Tome and Principe Dobra
	STD = "STD"
	// Salvadoran Colon
	SVC = "SVC"
	// Syrian Pound
	SYP = "SYP"
	// Swazi Lilangeni
	SZL = "SZL"
	// Thai Baht
	THB = "THB"
	// Tajikistani Somoni
	TJS = "TJS"
	// Turkmenistani Manat
	TMT = "TMT"
	// Tunisian Dinar
	TND = "TND"
	// Tongan Pa'anga
	TOP = "TOP"
	// Turkish Lira
	TRY = "TRY"
	// Trinidad and Tobago Dollar
	TTD = "TTD"
	// New Taiwan Dollar
	TWD = "TWD"
	// Tanzanian Shilling
	TZS = "TZS"
	// Ukrainian Hryvnia
	UAH = "UAH"
	// Ugandan Shilling
	UGX = "UGX"
	// United States Dollar
	USD = "USD"
	// Uruguayan Peso
	UYU = "UYU"
	// Uzbekistan Som
	UZS = "UZS"
	// Venezuelan Bolivar Fuerte
	VEF = "VEF"
	// Sovereign Bolivar
	VES = "VES"
	// Vietnamese Dong
	VND = "VND"
	// Vanuatu Vatu
	VUV = "VUV"
	// Samoan Tala
	WST = "WST"
	// CFA Franc BEAC
	XAF = "XAF"
	// Silver (troy ounce)
	XAG = "XAG"
	// Gold (troy ounce)
	XAU = "XAU"
	// East Caribbean Dollar
	XCD = "XCD"
	// Special Drawing Rights
	XDR = "XDR"
	// CFA Franc BCEAO
	XOF = "XOF"
	// CFP Franc
	XPF = "XPF"
	// Yemeni Rial
	YER = "YER"
	// South African Rand
	ZAR = "ZAR"
	// Zambian Kwacha (pre-2013)
	ZMK = "ZMK"
	// Zambian Kwacha
	ZMW = "ZMW"
	// Zimbabwean Dollar
	ZWL = "ZWL"
)

var rxCurrencyCode = regexp.MustCompile("^[A-Z]{3}$")

// Returns a boolean indicating whether the code is like
// the ISO 4217 currency code via a regular expression.
func Could(code string) bool {
	return rxCurrencyCode.MatchString(code)
}

// Returns a boolean indicating whether the code is a ISO 4217 currency code.
func Is(code string) bool {
	for _, currency := range []string{
		AED, AFN, ALL, AMD, ANG, AOA, ARS, AUD, AWG, AZN, BAM, BBD, BDT, BGN, BHD, BIF, BMD,
		BND, BOB, BRL, BSD, BTC, BTN, BWP, BYN, BYR, BZD, CAD, CDF, CHF, CLF, CLP, CNY, COP,
		CRC, CUC, CUP, CVE, CZK, DJF, DKK, DOP, DZD, EGP, ERN, ETB, EUR, FJD, FKP, GBP, GEL,
		GGP, GHS, GIP, GMD, GNF, GTQ, GYD, HKD, HNL, HRK, HTG, HUF, IDR, ILS, IMP, INR, IQD,
		IRR, ISK, JEP, JMD, JOD, JPY, KES, KGS, KHR, KMF, KPW, KRW, KWD, KYD, KZT, LAK, LBP,
		LKR, LRD, LSL, LTL, LVL, LYD, MAD, MDL, MGA, MKD, MMK, MNT, MOP, MRO, MUR, MVR, MWK,
		MXN, MYR, MZN, NAD, NGN, NIO, NOK, NPR, NZD, OMR, PAB, PEN, PGK, PHP, PKR, PLN, PYG,
		QAR, RON, RSD, RUB, RWF, SAR, SBD, SCR, SDG, SEK, SGD, SHP, SLE, SLL, SOS, SRD, STD,
		SVC, SYP, SZL, THB, TJS, TMT, TND, TOP, TRY, TTD, TWD, TZS, UAH, UGX, USD, UYU, UZS,
		VEF, VES, VND, VUV, WST, XAF, XAG, XAU, XCD, XDR, XOF, XPF, YER, ZAR, ZMK, ZMW, ZWL,
	} {
		if currency == code {
			return true
		}
	}

	return false
}
