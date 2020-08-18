package models

import (
	"strconv"
)

// GalQueueRegTags

const G01Id = "g01_id"
const G01Status = "g01_status"
const G01TempLilblu = "g01_temp_lilblu"
const G01TimeLilblu = "g01_time_lilblu"
const G01WtBktEmpt = "g01_wt_bkt_empt"
const G01WtBktFull = "g01_wt_bkt_full"
const G01WtLilblu = "g01_wt_lilblu"
const G01WtGood = "g01_wt_good"
const G02Id = "g02_id"
const G02Status = "g02_status"
const G02TempLilblu = "g02_temp_lilblu"
const G02TimeLilblu = "g02_time_lilblu"
const G02WtBktEmpt = "g02_wt_bkt_empt"
const G02WtBktFull = "g02_wt_bkt_full"
const G02WtLilblu = "g02_wt_lilblu"
const G02WtGood = "g02_wt_good"

var GIDTags = []string{G01Id, G02Id}
var GStatTags = []string{G01Status, G02Status}

// PreQueueRegTags

const W01Id = "w01_id"
const W01Status = "w01_status"
const W01TimeDegrease1 = "w01_time_degrease1"
const W01TimeDegrease2 = "w01_time_degrease2"
const W01TimeDegrease3 = "w01_time_degrease3"
const W01TimeFlux = "w01_time_flux"
const W01TimePickling1 = "w01_time_pickling1"
const W01TimePickling2 = "w01_time_pickling2"
const W01TimePickling3 = "w01_time_pickling3"
const W01WtBktEmpt = "w01_wt_bkt_empt"
const W01WtBktFull = "w01_wt_bkt_full"
const W01WtGood = "w01_wt_good"
const W02Id = "w02_id"
const W02Status = "w02_status"
const W02TimeDegrease1 = "w02_time_degrease1"
const W02TimeDegrease2 = "w02_time_degrease2"
const W02TimeDegrease3 = "w02_time_degrease3"
const W02TimeFlux = "w02_time_flux"
const W02TimePickling1 = "w02_time_pickling1"
const W02TimePickling2 = "w02_time_pickling2"
const W02TimePickling3 = "w02_time_pickling3"
const W02WtBktEmpt = "w02_wt_bkt_empt"
const W02WtBktFull = "w02_wt_bkt_full"
const W02WtGood = "w02_wt_good"
const W03Id = "w03_id"
const W03Status = "w03_status"
const W03TimeDegrease1 = "w03_time_degrease1"
const W03TimeDegrease2 = "w03_time_degrease2"
const W03TimeDegrease3 = "w03_time_degrease3"
const W03TimeFlux = "w03_time_flux"
const W03TimePickling1 = "w03_time_pickling1"
const W03TimePickling2 = "w03_time_pickling2"
const W03TimePickling3 = "w03_time_pickling3"
const W03WtBktEmpt = "w03_wt_bkt_empt"
const W03WtBktFull = "w03_wt_bkt_full"
const W03WtGood = "w03_wt_good"
const W04Id = "w04_id"
const W04Status = "w04_status"
const W04TimeDegrease1 = "w04_time_degrease1"
const W04TimeDegrease2 = "w04_time_degrease2"
const W04TimeDegrease3 = "w04_time_degrease3"
const W04TimeFlux = "w04_time_flux"
const W04TimePickling1 = "w04_time_pickling1"
const W04TimePickling2 = "w04_time_pickling2"
const W04TimePickling3 = "w04_time_pickling3"
const W04WtBktEmpt = "w04_wt_bkt_empt"
const W04WtBktFull = "w04_wt_bkt_full"
const W04WtGood = "w04_wt_good"
const W05Id = "w05_id"
const W05Status = "w05_status"
const W05TimeDegrease1 = "w05_time_degrease1"
const W05TimeDegrease2 = "w05_time_degrease2"
const W05TimeDegrease3 = "w05_time_degrease3"
const W05TimeFlux = "w05_time_flux"
const W05TimePickling1 = "w05_time_pickling1"
const W05TimePickling2 = "w05_time_pickling2"
const W05TimePickling3 = "w05_time_pickling3"
const W05WtBktEmpt = "w05_wt_bkt_empt"
const W05WtBktFull = "w05_wt_bkt_full"
const W05WtGood = "w05_wt_good"
const W06Id = "w06_id"
const W06Status = "w06_status"
const W06TimeDegrease1 = "w06_time_degrease1"
const W06TimeDegrease2 = "w06_time_degrease2"
const W06TimeDegrease3 = "w06_time_degrease3"
const W06TimeFlux = "w06_time_flux"
const W06TimePickling1 = "w06_time_pickling1"
const W06TimePickling2 = "w06_time_pickling2"
const W06TimePickling3 = "w06_time_pickling3"
const W06WtBktEmpt = "w06_wt_bkt_empt"
const W06WtBktFull = "w06_wt_bkt_full"
const W06WtGood = "w06_wt_good"
const W07Id = "w07_id"
const W07Status = "w07_status"
const W07TimeDegrease1 = "w07_time_degrease1"
const W07TimeDegrease2 = "w07_time_degrease2"
const W07TimeDegrease3 = "w07_time_degrease3"
const W07TimeFlux = "w07_time_flux"
const W07TimePickling1 = "w07_time_pickling1"
const W07TimePickling2 = "w07_time_pickling2"
const W07TimePickling3 = "w07_time_pickling3"
const W07WtBktEmpt = "w07_wt_bkt_empt"
const W07WtBktFull = "w07_wt_bkt_full"
const W07WtGood = "w07_wt_good"
const W08Id = "w08_id"
const W08Status = "w08_status"
const W08TimeDegrease1 = "w08_time_degrease1"
const W08TimeDegrease2 = "w08_time_degrease2"
const W08TimeDegrease3 = "w08_time_degrease3"
const W08TimeFlux = "w08_time_flux"
const W08TimePickling1 = "w08_time_pickling1"
const W08TimePickling2 = "w08_time_pickling2"
const W08TimePickling3 = "w08_time_pickling3"
const W08WtBktEmpt = "w08_wt_bkt_empt"
const W08WtBktFull = "w08_wt_bkt_full"
const W08WtGood = "w08_wt_good"
const W09Id = "w09_id"
const W09Status = "w09_status"
const W09TimeDegrease1 = "w09_time_degrease1"
const W09TimeDegrease2 = "w09_time_degrease2"
const W09TimeDegrease3 = "w09_time_degrease3"
const W09TimeFlux = "w09_time_flux"
const W09TimePickling1 = "w09_time_pickling1"
const W09TimePickling2 = "w09_time_pickling2"
const W09TimePickling3 = "w09_time_pickling3"
const W09WtBktEmpt = "w09_wt_bkt_empt"
const W09WtBktFull = "w09_wt_bkt_full"
const W09WtGood = "w09_wt_good"
const W10Id = "w10_id"
const W10Status = "w10_status"
const W10TimeDegrease1 = "w10_time_degrease1"
const W10TimeDegrease2 = "w10_time_degrease2"
const W10TimeDegrease3 = "w10_time_degrease3"
const W10TimeFlux = "w10_time_flux"
const W10TimePickling1 = "w10_time_pickling1"
const W10TimePickling2 = "w10_time_pickling2"
const W10TimePickling3 = "w10_time_pickling3"
const W10WtBktEmpt = "w10_wt_bkt_empt"
const W10WtBktFull = "w10_wt_bkt_full"
const W10WtGood = "w10_wt_good"
const W11Id = "w11_id"
const W11Status = "w11_status"
const W11TimeDegrease1 = "w11_time_degrease1"
const W11TimeDegrease2 = "w11_time_degrease2"
const W11TimeDegrease3 = "w11_time_degrease3"
const W11TimeFlux = "w11_time_flux"
const W11TimePickling1 = "w11_time_pickling1"
const W11TimePickling2 = "w11_time_pickling2"
const W11TimePickling3 = "w11_time_pickling3"
const W11WtBktEmpt = "w11_wt_bkt_empt"
const W11WtBktFull = "w11_wt_bkt_full"
const W11WtGood = "w11_wt_good"
const W12Id = "w12_id"
const W12Status = "w12_status"
const W12TimeDegrease1 = "w12_time_degrease1"
const W12TimeDegrease2 = "w12_time_degrease2"
const W12TimeDegrease3 = "w12_time_degrease3"
const W12TimeFlux = "w12_time_flux"
const W12TimePickling1 = "w12_time_pickling1"
const W12TimePickling2 = "w12_time_pickling2"
const W12TimePickling3 = "w12_time_pickling3"
const W12WtBktEmpt = "w12_wt_bkt_empt"
const W12WtBktFull = "w12_wt_bkt_full"
const W12WtGood = "w12_wt_good"
const W13Id = "w13_id"
const W13Status = "w13_status"
const W13TimeDegrease1 = "w13_time_degrease1"
const W13TimeDegrease2 = "w13_time_degrease2"
const W13TimeDegrease3 = "w13_time_degrease3"
const W13TimeFlux = "w13_time_flux"
const W13TimePickling1 = "w13_time_pickling1"
const W13TimePickling2 = "w13_time_pickling2"
const W13TimePickling3 = "w13_time_pickling3"
const W13WtBktEmpt = "w13_wt_bkt_empt"
const W13WtBktFull = "w13_wt_bkt_full"
const W13WtGood = "w13_wt_good"
const W14Id = "w14_id"
const W14Status = "w14_status"
const W14TimeDegrease1 = "w14_time_degrease1"
const W14TimeDegrease2 = "w14_time_degrease2"
const W14TimeDegrease3 = "w14_time_degrease3"
const W14TimeFlux = "w14_time_flux"
const W14TimePickling1 = "w14_time_pickling1"
const W14TimePickling2 = "w14_time_pickling2"
const W14TimePickling3 = "w14_time_pickling3"
const W14WtBktEmpt = "w14_wt_bkt_empt"
const W14WtBktFull = "w14_wt_bkt_full"
const W14WtGood = "w14_wt_good"
const W15Id = "w15_id"
const W15Status = "w15_status"
const W15TimeDegrease1 = "w15_time_degrease1"
const W15TimeDegrease2 = "w15_time_degrease2"
const W15TimeDegrease3 = "w15_time_degrease3"
const W15TimeFlux = "w15_time_flux"
const W15TimePickling1 = "w15_time_pickling1"
const W15TimePickling2 = "w15_time_pickling2"
const W15TimePickling3 = "w15_time_pickling3"
const W15WtBktEmpt = "w15_wt_bkt_empt"
const W15WtBktFull = "w15_wt_bkt_full"
const W15WtGood = "w15_wt_good"
const W16Id = "w16_id"
const W16Status = "w16_status"
const W16TimeDegrease1 = "w16_time_degrease1"
const W16TimeDegrease2 = "w16_time_degrease2"
const W16TimeDegrease3 = "w16_time_degrease3"
const W16TimeFlux = "w16_time_flux"
const W16TimePickling1 = "w16_time_pickling1"
const W16TimePickling2 = "w16_time_pickling2"
const W16TimePickling3 = "w16_time_pickling3"
const W16WtBktEmpt = "w16_wt_bkt_empt"
const W16WtBktFull = "w16_wt_bkt_full"
const W16WtGood = "w16_wt_good"
const W17Id = "w17_id"
const W17Status = "w17_status"
const W17TimeDegrease1 = "w17_time_degrease1"
const W17TimeDegrease2 = "w17_time_degrease2"
const W17TimeDegrease3 = "w17_time_degrease3"
const W17TimeFlux = "w17_time_flux"
const W17TimePickling1 = "w17_time_pickling1"
const W17TimePickling2 = "w17_time_pickling2"
const W17TimePickling3 = "w17_time_pickling3"
const W17WtBktEmpt = "w17_wt_bkt_empt"
const W17WtBktFull = "w17_wt_bkt_full"
const W17WtGood = "w17_wt_good"
const W18Id = "w18_id"
const W18Status = "w18_status"
const W18TimeDegrease1 = "w18_time_degrease1"
const W18TimeDegrease2 = "w18_time_degrease2"
const W18TimeDegrease3 = "w18_time_degrease3"
const W18TimeFlux = "w18_time_flux"
const W18TimePickling1 = "w18_time_pickling1"
const W18TimePickling2 = "w18_time_pickling2"
const W18TimePickling3 = "w18_time_pickling3"
const W18WtBktEmpt = "w18_wt_bkt_empt"
const W18WtBktFull = "w18_wt_bkt_full"
const W18WtGood = "w18_wt_good"
const W19Id = "w19_id"
const W19Status = "w19_status"
const W19TimeDegrease1 = "w19_time_degrease1"
const W19TimeDegrease2 = "w19_time_degrease2"
const W19TimeDegrease3 = "w19_time_degrease3"
const W19TimeFlux = "w19_time_flux"
const W19TimePickling1 = "w19_time_pickling1"
const W19TimePickling2 = "w19_time_pickling2"
const W19TimePickling3 = "w19_time_pickling3"
const W19WtBktEmpt = "w19_wt_bkt_empt"
const W19WtBktFull = "w19_wt_bkt_full"
const W19WtGood = "w19_wt_good"
const W20Id = "w20_id"
const W20Status = "w20_status"
const W20TimeDegrease1 = "w20_time_degrease1"
const W20TimeDegrease2 = "w20_time_degrease2"
const W20TimeDegrease3 = "w20_time_degrease3"
const W20TimeFlux = "w20_time_flux"
const W20TimePickling1 = "w20_time_pickling1"
const W20TimePickling2 = "w20_time_pickling2"
const W20TimePickling3 = "w20_time_pickling3"
const W20WtBktEmpt = "w20_wt_bkt_empt"
const W20WtBktFull = "w20_wt_bkt_full"
const W20WtGood = "w20_wt_good"

var WIDTags = []string{
	W01Id, W02Id, W03Id, W04Id, W05Id,
	W06Id, W07Id, W08Id, W09Id, W10Id,
	W11Id, W12Id, W13Id, W14Id, W15Id,
	W16Id, W17Id, W18Id, W19Id, W20Id}

var WStatTags = []string{
	W01Status, W02Status, W03Status, W04Status, W05Status,
	W06Status, W07Status, W08Status, W09Status, W10Status,
	W11Status, W12Status, W13Status, W14Status, W15Status,
	W16Status, W17Status, W18Status, W19Status, W20Status}

var W01Metrics = []string{
	W01WtBktEmpt, W01WtBktFull, W01TimeDegrease1, W01TimeDegrease2, W01TimeDegrease3,
	W01TimePickling1, W01TimePickling2, W01TimePickling3, W01TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W01WtGood}

var W02Metrics = []string{
	W02WtBktEmpt, W02WtBktFull, W02TimeDegrease1, W02TimeDegrease2, W02TimeDegrease3,
	W02TimePickling1, W02TimePickling2, W02TimePickling3, W02TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W02WtGood}

var W03Metrics = []string{
	W03WtBktEmpt, W03WtBktFull, W03TimeDegrease1, W03TimeDegrease2, W03TimeDegrease3,
	W03TimePickling1, W03TimePickling2, W03TimePickling3, W03TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W03WtGood}

var W04Metrics = []string{
	W04WtBktEmpt, W04WtBktFull, W04TimeDegrease1, W04TimeDegrease2, W04TimeDegrease3,
	W04TimePickling1, W04TimePickling2, W04TimePickling3, W04TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W04WtGood}

var W05Metrics = []string{
	W05WtBktEmpt, W05WtBktFull, W05TimeDegrease1, W05TimeDegrease2, W05TimeDegrease3,
	W05TimePickling1, W05TimePickling2, W05TimePickling3, W05TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W05WtGood}

var W06Metrics = []string{
	W06WtBktEmpt, W06WtBktFull, W06TimeDegrease1, W06TimeDegrease2, W06TimeDegrease3,
	W06TimePickling1, W06TimePickling2, W06TimePickling3, W06TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W06WtGood}

var W07Metrics = []string{
	W07WtBktEmpt, W07WtBktFull, W07TimeDegrease1, W07TimeDegrease2, W07TimeDegrease3,
	W07TimePickling1, W07TimePickling2, W07TimePickling3, W07TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W07WtGood}

var W08Metrics = []string{
	W08WtBktEmpt, W08WtBktFull, W08TimeDegrease1, W08TimeDegrease2, W08TimeDegrease3,
	W08TimePickling1, W08TimePickling2, W08TimePickling3, W08TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W08WtGood}

var W09Metrics = []string{
	W09WtBktEmpt, W09WtBktFull, W09TimeDegrease1, W09TimeDegrease2, W09TimeDegrease3,
	W09TimePickling1, W09TimePickling2, W09TimePickling3, W09TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W09WtGood}

var W10Metrics = []string{
	W10WtBktEmpt, W10WtBktFull, W10TimeDegrease1, W10TimeDegrease2, W10TimeDegrease3,
	W10TimePickling1, W10TimePickling2, W10TimePickling3, W10TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W10WtGood}

var W11Metrics = []string{
	W11WtBktEmpt, W11WtBktFull, W11TimeDegrease1, W11TimeDegrease2, W11TimeDegrease3,
	W11TimePickling1, W11TimePickling2, W11TimePickling3, W11TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W11WtGood}

var W12Metrics = []string{
	W12WtBktEmpt, W12WtBktFull, W12TimeDegrease1, W12TimeDegrease2, W12TimeDegrease3,
	W12TimePickling1, W12TimePickling2, W12TimePickling3, W12TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W12WtGood}

var W13Metrics = []string{
	W13WtBktEmpt, W13WtBktFull, W13TimeDegrease1, W13TimeDegrease2, W13TimeDegrease3,
	W13TimePickling1, W13TimePickling2, W13TimePickling3, W13TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W13WtGood}

var W14Metrics = []string{
	W14WtBktEmpt, W14WtBktFull, W14TimeDegrease1, W14TimeDegrease2, W14TimeDegrease3,
	W14TimePickling1, W14TimePickling2, W14TimePickling3, W14TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W14WtGood}

var W15Metrics = []string{
	W15WtBktEmpt, W15WtBktFull, W15TimeDegrease1, W15TimeDegrease2, W15TimeDegrease3,
	W15TimePickling1, W15TimePickling2, W15TimePickling3, W15TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W15WtGood}

var W16Metrics = []string{
	W16WtBktEmpt, W16WtBktFull, W16TimeDegrease1, W16TimeDegrease2, W16TimeDegrease3,
	W16TimePickling1, W16TimePickling2, W16TimePickling3, W16TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W16WtGood}

var W17Metrics = []string{
	W17WtBktEmpt, W17WtBktFull, W17TimeDegrease1, W17TimeDegrease2, W17TimeDegrease3,
	W17TimePickling1, W17TimePickling2, W17TimePickling3, W17TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W17WtGood}

var W18Metrics = []string{
	W18WtBktEmpt, W18WtBktFull, W18TimeDegrease1, W18TimeDegrease2, W18TimeDegrease3,
	W18TimePickling1, W18TimePickling2, W18TimePickling3, W18TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W18WtGood}

var W19Metrics = []string{
	W19WtBktEmpt, W19WtBktFull, W19TimeDegrease1, W19TimeDegrease2, W19TimeDegrease3,
	W19TimePickling1, W19TimePickling2, W19TimePickling3, W19TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W19WtGood}

var W20Metrics = []string{
	W20WtBktEmpt, W20WtBktFull, W20TimeDegrease1, W20TimeDegrease2, W20TimeDegrease3,
	W20TimePickling1, W20TimePickling2, W20TimePickling3, W20TimeFlux,
	Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	W20WtGood}

var jobMetrics = [][]string{
	W01Metrics, W02Metrics, W03Metrics, W04Metrics, W05Metrics,
	W06Metrics, W07Metrics, W08Metrics, W09Metrics, W10Metrics,
	W11Metrics, W12Metrics, W13Metrics, W14Metrics, W15Metrics,
	W16Metrics, W17Metrics, W18Metrics, W19Metrics, W20Metrics,
}

// GalMeasureBufferTags

const Gal01Id = "gal01_id"
const Gal01TempSet = "gal01_temp_set"
const Gal01TempAct = "gal01_temp_act"

// PreMeasureBufferTags

const Preproc01Id = "preproc01_id"
const Preproc01Conduct = "preproc01_conduct"
const Preproc01Ph = "preproc01_ph"
const Preproc01Temp = "preproc01_temp"
const Preproc02Id = "preproc02_id"
const Preproc02Conduct = "preproc02_conduct"
const Preproc02Ph = "preproc02_ph"
const Preproc02Temp = "preproc02_temp"
const Preproc03Id = "preproc03_id"
const Preproc03Conduct = "preproc03_conduct"
const Preproc03Ph = "preproc03_ph"
const Preproc03Temp = "preproc03_temp"
const Preproc04Id = "preproc04_id"
const Preproc04Conduct = "preproc04_conduct"
const Preproc04Ph = "preproc04_ph"
const Preproc04Temp = "preproc04_temp"
const Preproc05Id = "preproc05_id"
const Preproc05Conduct = "preproc05_conduct"
const Preproc05Ph = "preproc05_ph"
const Preproc05Temp = "preproc05_temp"

// GetTagArray generates a number of array tag name strings specified by 'count'
func GetTagArray(baseTag string, count int) []string {
	var s = []string{}
	for i := 0; i < count; i++ {
		s = append(s, baseTag+"("+strconv.Itoa(i)+")")
	}
	return s
}

// GetJobMetrics ...
func GetJobMetrics(index int) []string {
	return jobMetrics[index]
}
