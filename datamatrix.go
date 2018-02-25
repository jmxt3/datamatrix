package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

//RUN ON THE TERMINAL
//go build datamatrix.go && ./datamatrix

func main() {

	skus := [231]string{
		"ACUC-A-AMARELO-SID-10X1KG",
		"ACUC-B-DCG-SID-10X1KG",
		"ACUC-ICING-SID-1X10KG",
		"ACUC-BRANCO-SID-1X25KG",
		"ACUC-G-SACOS-SID-1X25KG",
		"ACUC-B-1200-DOS-SID-1X96KG",
		"AZEI-SA-3X5LT",
		"AZEI-SA-1X12LT",
		"OLEO-A-GIRASSOL-ABR-1X5LT",
		"GORD-V-FRIT-ABR-1X10LT",
		"OLEO-ALIMENTA-FUL-15X1LT",
		"AZEI-B-DE-AZEIT-COO-3X5LT",
		"OLEO-A-TEMPERO-GDO-3X5LT",
		"VINA-UNI-20X075LT",
		"VINA-D-MESA-SA-12X025LT",
		"AZEI-D-MESA-SA-12X025LT",
		"BOLA-TOSTADA-CUE-13X800GR",
		"BOLA-MARIA-UNI-10X800GR",
		"CAFE-D-MMP-DEL-20X250GR",
		"CEVA-SOLUVEL-NES-12X200GR",
		"CHA-D-CIDREIRA-LIP-6X25UN",
		"CHA-D-CAMOMILA-LIP-6X25UN",
		"CHA-D-TILIA-LIP-6X25UN",
		"CHA-D-LUCIA-LI-LIP-6X25UN",
		"CHA-G-TEA-ORIE-LIP-6X25UN",
		"CHA-G-TEA-CLAS-LIP-6X25UN",
		"CHA-P-YELLOW-L-LIP-6X25UN",
		"CHA-F-VERMELHO-LIP-6X25UN",
		"CHA-M-E-CANELA-LIP-6X25UN",
		"CHA-TILIA-VIV-12X25UN",
		"CHA-D-CIDREIRA-VIV-12X25UN",
		"CHA-D-CAMOMILA-VIV-12X25UN",
		"CHA-D-LUCIA-LI-VIV-12X25UN",
		"CHOC-C-44%-CACA-NES-20X200GR",
		"CHOC-B-MESTRE-P-MAS-1X3KG",
		"CREM-D-CMP-MAS-1X5KG",
		"LATA-D-PESSEGO-UNI-3X265KG",
		"LATA-D-PESSEGO-UNI-12X840GR",
		"LATA-D-ANANAS-CIM-6X3KG",
		"LATA-D-ANANAS-UNI-12X822GR",
		"LATA-D-COGUMELO-UNI-3X25KG",
		"LATA-D-COGUMELO-UNI-12X780GR",
		"LATA-D-COGUMELO-UNI-12X780GR-1",
		"LATA-D-FEIJAO-M-PIC-3X25KG",
		"LATA-D-FMC-UNI-12X780GR",
		"LATA-D-FEIJAO-B-PIC-3X25KG",
		"LATA-D-PIMENTOS-CIM-12X780GR",
		"LATA-D-FBC-UNI-12X780GR",
		"LATA-D-GRAO-UNI-3X25KG",
		"LATA-D-GRAO-UNI-12X780GR",
		"LATA-D-FFC-UNI-12X780GR",
		"LATA-D-FVC-UNI-12X780GR",
		"SACO-D-FBS-PL-1X5KG",
		"SACO-D-FCS-PL-1X5KG",
		"SACO-D-GRAO-SEC-PL-1X5KG",
		"SACO-D-FFS-PL-1X5KG",
		"LATA-D-AEP-UNI-6X173KG",
		"LATA-D-AEP-UNI-25X110GR",
		"LATA-D-SET-UNI-25X110GR",
		"LATA-D-SEO-JUP-25X110GR",
		"LATA-D-FDC-JUP-25X110GR",
		"LATA-D-CAVALINH-JUP-25X110GR",
		"LATA-D-CARAPAUS-JUP-25X110GR",
		"LATA-D-TOMATE-UNI-12X780GR",
		"LATA-D-TOMATE-UNI-6X255KG",
		"LATA-D-TOMATE-P-KNO-6X25KG",
		"LATA-D-TOMATE-C-CIR-6X255KG",
		"LATA-D-PTT-KNO-6X255KG",
		"LATA-D-CUBOS-TO-KNO-6X255KG",
		"CREM-D-TOMATE-KNO-6X810GR",
		"FRAS-D-FMC-UNI-12X570GR",
		"FRAS-D-FEIJAO-B-UNI-12X570GR",
		"FRAS-D-FEIJAO-F-UNI-12X570GR",
		"FRAS-D-FEIJAO-V-UNI-12X570GR",
		"FRAS-D-FBC-UNI-12X570GR",
		"BALD-D-DOCE-CHI-AM-1X15KG",
		"TABU-D-MARMELAD-UNI-1X55KG",
		"FRAS-D-MEL-FRU-1X1KG",
		"FRAS-D-DDM-ALT-12X280GR",
		"FRAS-D-DDM-ALT-12X280GR-1",
		"FRAS-D-DDP-ALT-12X280GR",
		"SACO-D-OREGAOS-MEL-1X500GR",
		"FRAS-A-INDIAS-CON-1X570GR",
		"FRAS-P-MOIDO-CON-1X345GR",
		"FRAS-C-MOIDOS-CON-1X345GR",
		"FRAS-P-BRANCA-M-CON-1X555GR",
		"FRAS-D-PBE-CON-1X655GR",
		"FRAS-D-PIMENTAO-CHA-1X550GR",
		"FRAS-P-VAGEM-CON-1X270GR",
		"FRAS-D-CEP-CON-1X250GR",
		"FRAS-D-CARIL-CON-1X495GR",
		"FRAS-D-EDE-CON-1X430GR",
		"FRAS-D-EDM-CON-1X405GR",
		"SACO-D-CANELA-M-CON-1X1KG",
		"SACO-D-CEP-CON-1X500GR",
		"CANE-E-PAU-INDI-CON-1X500UN",
		"FRAS-D-TAE-CHA-1X300GR",
		"SACO-D-GLUTAMAT-CON-1X100GR",
		"SACO-D-PIRIPIRI-CON-1X5KG",
		"BALD-D-MDA-INC-1X5KG",
		"BALD-D-MDA-INC-1X1KG",
		"BALD-D-MDP-INC-1X5KG",
		"BISN-D-SDM-UNI-20X250GR",
		"PACO-D-SAL-GROS-UNI-1X20KG",
		"SACO-D-SAL-GROS-SAL-1X25KG",
		"SACO-D-SAL-FINO-SAL-1X25KG",
		"SACO-D-SAL-REGE-SAL-1X25KG",
		"PAST-D-CDG-KNO-6X1KG",
		"PAST-D-CDG-KNO-1X5KG",
		"CUBO-DCD-KNO-6X960GR",
		"PAST-D-CDC-KNO-6X1KG",
		"CUBO-DCD-KNO-6X960GR-1",
		"CALD-D-PEIXE-KNO-6X700GR",
		"CALD-D-PDM-KNO-6X700GR",
		"CREM-D-DE-MARIS-KNO-6X683GR",
		"CREM-D-DE-MARIS-KNO-1X25KG",
		"MOLH-D-ASSADO-KNO-6X800GR",
		"MOLH-D-ASSADO-1-KNO-1X3KG",
		"MOLH-D-DG1-KNO-6X750GR",
		"MOLH-D-DEMI-GLA-KNO-6X900GR",
		"MOLH-D-DG1-KNO-1X3KG",
		"MOLH-CERVEJEI-KNO-6X700GR",
		"PURE-D-BATATA-KNO-1X4KG",
		"PURE-D-BATATA-KNO-6X850GR",
		"FLOC-D-BATATA-KNO-1X10KG",
		"PRIM-T-PASTA-MA-KNO-2X340GR",
		"PRIM-T-PED-KNO-2X340GR",
		"FARI-A-DE-MILHO-MAI-4X25KG",
		"NATA-C-CUISINE-VAQ-12X1LT",
		"CATY-F-EM-PO-MAS-1X5KG",
		"CATY-BISCUIT-MAS-1X15KG",
		"FARI-ARROZ-CER-1X25KG",
		"FARI-M-TIPO-175-CER-1X25KG",
		"FARI-T-E-ALFARR-CER-1X10KG",
		"FARI-I-TIPO-150-CER-1X25KG",
		"FARI-C-TIPO-130-CER-1X25KG",
		"FARI-C-TIPO-85-CER-1X25KG",
		"FARI-T-00-PIZZA-CER-1X10KG",
		"FARI-T-65-CER-1X25KG",
		"FARI-T-TIPO-65-GRA-1X25KG",
		"FARI-T-55-GRA-1X25KG",
		"FARI-T-55-CER-1X25KG",
		"FARI-S-FT5-UNI-10X1KG",
		"FARI-C-FERMENTO-UNI-10X1KG",
		"ESPA-UNI-20X500GR",
		"MASS-D-COTOVELO-UNI-20X500GR",
		"MASS-COTOVELI-UNI-20X500GR",
		"MASS-PEVIDE-UNI-40X250GR",
		"MASS-CUSCUS-UNI-40X250GR",
		"MASS-ESPIRAL-ALT-20X500GR",
		"MASS-MACARRON-UNI-20X500GR",
		"ARRO-D-TAMBORIL-SA-12X1KG",
		"ARRO-D-TAMBORIL-SA-12X1KG-1",
		"ARRO-D-TAMBORIL-SA-12X1KG-2",
		"ARRO-D-TAMBORIL-SA-12X1KG-3",
		"REBU-DR-35KG",
		"CARA-D-FRUTA-PEN-15KG",
		"CARA-D-NATA-PEN-15KG",
		"REBU-D-MENTOL-PEN-15KG",
		"REBU-D-MEL-PEN-15KG",
		"LEIT-1-MEIO-GOR-SER-6X1LT",
		"LEIT-CHOCOLAT-UCA-1X24UN",
		"LEIT-CONDENSA-NES-12X370GR",
		"LEIT-CONDENSA-FF-12X397GR",
		"LEIT-NESQUIK-NES-1X30UN",
		"NATA-C-&-CREME-VAQ-12X1LT",
		"TEMP-D-AROMAT-KNO-6X500GR",
		"TEMP-D-MVD-KNO-6X810GR",
		"MANT-RP-6X1KG",
		"MANT-INS-6X1KG",
		"MANT-D-COM-ALHO-MIM-100X10GR",
		"MANT-DOSES-MIM-100X10GR",
		"LICO-BEIRAO-LB-1X1LT",
		"BRAN-MAC-1X1LT",
		"BRAN-CRO-1X1LT",
		"AGUA-F-DE-SECUL-1X1LT",
		"AGUA-ANTIQUA-CA-1X1LT",
		"AGUA-PDM-1X1LT",
		"AGUA-RESERVA-CR&-1X1LT",
		"VINH-D-PORTO-VE-VEL-1X1LT",
		"VINH-D-PORTO-VD-1X1LT",
		"MOSC-JP-1X1LT",
		"MOSC-ALH-1X1LT",
		"MOSC-PEG-1X1LT",
		"AMEN-AMARGA-AMA-1X1LT",
		"BAGA-C-ROCHEDO-1X1LT",
		"MART-TINTOS-MAR-1X50UN",
		"MART-TINTO-MAR-1X1LT",
		"BRAN-1920-BRA-1X1LT",
		"WHIS-J-5-ANOS-J&B-1X1LT",
		"WHIS-CS-1X075LT",
		"WHIS-G-5-ANOS-GRA-1X075LT",
		"WHIS-R-LABEL-JW-1X075LT",
		"WHIS-B-LABEL-JW-1X075LT",
		"WHIS-FG-1X075LT",
		"WHIS-L-12-ANOS-LOG-1X07LT",
		"WHIS-C-1X07-LT-CHI-1X07LT",
		"WHIS-J-B10-J&B-1X07LT",
		"GINJ-ALB-1X07LT",
		"GUAR-GP-40X40UN",
		"GUAR-2-FOLHAS-GC-40X40UN",
		"TOAL-P-DE-MARIS-GP-1X2000UN",
		"SAQU-P-BRANCAS-AP-1X1000UN",
		"SAQU-P-BRANCAS-AP-1X1000UN-1",
		"SAQU-P-KRAFT-F3-AP-1X1000UN",
		"SAQU-P-KRAFT-F4-AP-1X1000UN",
		"SACO-A-BRANCAS-S88-1X1000UN",
		"SACO-A-BI4-S88-1X1000UN",
		"SACO-A-CI3-S88-1X1000UN",
		"SACO-A-CI4-S88-1X1000UN",
		"CAIX-B-CIT-EN-1X500UN",
		"CAIX-B-CIT-EN-1X500UN-1",
		"CAIX-B-CIT-EN-1X500UN-2",
		"CAIX-B-CIT-EN-1X500UN-3",
		"CAIX-B-CIT-EN-1X100UN",
		"CAIX-B-CIT-EN-1X100UN-1",
		"CAIX-B-CIT-EN-1X100UN-2",
		"CAIX-B-CIT-EN-1X100UN-3",
		"CAIX-P-MCI-END-1X100UN",
		"CAIX-P-MCI-END-1X100UN-1",
		"CAIX-P-MCI-END-1X100UN-2",
		"CAIX-P-MCI-END-1X100UN-3",
		"TOAL-M-IMPRESSA-VAL-1X1000UN",
		"TOAL-M-IMPRESSA-VAL-1X1000UN-1",
		"TOAL-M-IMPRESSA-VAL-1X1000UN-2",
		"TOAL-M-IMPRESSA-VAL-1X1000UN-3",
		"TOAL-M-IMPRESSA-VAL-1X1000UN-4",
		"TOAL-M-IMPRESSA-VAL-1X1000UN-5",
		"RESM-P-EKI-AP-1X250UN",
		"ROLO-P-KI3-AP-1X30UN",
		"ROLO-P-KI3-AP-1X30UN-1",
	}

	os.RemoveAll("images/")
	os.MkdirAll("images/", os.ModePerm)

	for index, element := range skus {

		fmt.Printf("%d Data Matrix Created: %s.png\n", index, element)
		// Create the barcode
		qrCode, _ := qr.Encode(element, qr.M, qr.Auto)
		// Scale the barcode to 800x800 pixels
		qrCode, _ = barcode.Scale(qrCode, 800, 800)
		// create the output file
		file, _ := os.Create("images/" + element + ".png")
		defer file.Close()
		// encode the barcode as png
		png.Encode(file, qrCode)
	}
}
