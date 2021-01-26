package models

import (
	"strconv"
)

// GalQueueRegTags

const GScale = "g_scale"

const G01Id = "g01_id"
const G01Status = "g01_status"
const G01TempLilblu = "g01_temp_lilblu"
const G01TimeLilblu = "g01_time_lilblu"
const G01WtLilblu = "g01_wt_lilblu"
const G02Id = "g02_id"
const G02Status = "g02_status"
const G02TempLilblu = "g02_temp_lilblu"
const G02TimeLilblu = "g02_time_lilblu"
const G02WtLilblu = "g02_wt_lilblu"

var GIDTags = []string{G01Id, G02Id}
var GStatTags = []string{G01Status, G02Status}

// PreQueueRegTags

const WScale = "w_scale"

const W01Id = "w01_id"
const W01Status = "w01_status"
const W01WtBktEmpt = "w01_wt_bkt_empt"
const W01WtBktFull = "w01_wt_bkt_full"
const W01TimePicklingS1 = "w01_time_pickling_s1"
const W01TimePicklingS2 = "w01_time_pickling_s2"
const W01TimePicklingS3 = "w01_time_pickling_s3"
const W01TimeDegreaseS1 = "w01_time_degrease_s1"
const W01TimeDegreaseS2 = "w01_time_degrease_s2"
const W01TimeDegreaseS3 = "w01_time_degrease_s3"
const W01TimeHotWash1 = "w01_time_hot_wash1"
const W01TimeHotWash2 = "w01_time_hot_wash2"
const W01TimeHotWash3 = "w01_time_hot_wash3"
const W01TimeDegrease1 = "w01_time_degrease1"
const W01TimeDegrease2 = "w01_time_degrease2"
const W01TimeDegrease3 = "w01_time_degrease3"
const W01TimeDwash1 = "w01_time_dwash1"
const W01TimeDwash2 = "w01_time_dwash2"
const W01TimeDwash3 = "w01_time_dwash3"
const W01TimeFlux1 = "w01_time_flux1"
const W01TimeFlux2 = "w01_time_flux2"
const W01TimeFlux3 = "w01_time_flux3"
const W01TimeDry1 = "w01_time_dry1"
const W01TimeDry2 = "w01_time_dry2"
const W01TimeDry3 = "w01_time_dry3"
const W01TimePickling11 = "w01_time_pickling1_1"
const W01TimePickling12 = "w01_time_pickling1_2"
const W01TimePickling13 = "w01_time_pickling1_3"
const W01TimePickling21 = "w01_time_pickling2_1"
const W01TimePickling22 = "w01_time_pickling2_2"
const W01TimePickling23 = "w01_time_pickling2_3"
const W01TimePwash1 = "w01_time_pwash1"
const W01TimePwash2 = "w01_time_pwash2"
const W01TimePwash3 = "w01_time_pwash3"
const W01TimeSflux1 = "w01_time_sflux1"
const W01TimeSflux2 = "w01_time_sflux2"
const W01TimeSflux3 = "w01_time_sflux3"
const W01TimeBkt1 = "w01_time_bkt1"
const W01TimeBkt2 = "w01_time_bkt2"
const W01TimeBkt3 = "w01_time_bkt3"
const W02Id = "w02_id"
const W02Status = "w02_status"
const W02WtBktEmpt = "w02_wt_bkt_empt"
const W02WtBktFull = "w02_wt_bkt_full"
const W02TimePicklingS1 = "w02_time_pickling_s1"
const W02TimePicklingS2 = "w02_time_pickling_s2"
const W02TimePicklingS3 = "w02_time_pickling_s3"
const W02TimeDegreaseS1 = "w02_time_degrease_s1"
const W02TimeDegreaseS2 = "w02_time_degrease_s2"
const W02TimeDegreaseS3 = "w02_time_degrease_s3"
const W02TimeHotWash1 = "w02_time_hot_wash1"
const W02TimeHotWash2 = "w02_time_hot_wash2"
const W02TimeHotWash3 = "w02_time_hot_wash3"
const W02TimeDegrease1 = "w02_time_degrease1"
const W02TimeDegrease2 = "w02_time_degrease2"
const W02TimeDegrease3 = "w02_time_degrease3"
const W02TimeDwash1 = "w02_time_dwash1"
const W02TimeDwash2 = "w02_time_dwash2"
const W02TimeDwash3 = "w02_time_dwash3"
const W02TimeFlux1 = "w02_time_flux1"
const W02TimeFlux2 = "w02_time_flux2"
const W02TimeFlux3 = "w02_time_flux3"
const W02TimeDry1 = "w02_time_dry1"
const W02TimeDry2 = "w02_time_dry2"
const W02TimeDry3 = "w02_time_dry3"
const W02TimePickling11 = "w02_time_pickling1_1"
const W02TimePickling12 = "w02_time_pickling1_2"
const W02TimePickling13 = "w02_time_pickling1_3"
const W02TimePickling21 = "w02_time_pickling2_1"
const W02TimePickling22 = "w02_time_pickling2_2"
const W02TimePickling23 = "w02_time_pickling2_3"
const W02TimePwash1 = "w02_time_pwash1"
const W02TimePwash2 = "w02_time_pwash2"
const W02TimePwash3 = "w02_time_pwash3"
const W02TimeSflux1 = "w02_time_sflux1"
const W02TimeSflux2 = "w02_time_sflux2"
const W02TimeSflux3 = "w02_time_sflux3"
const W02TimeBkt1 = "w02_time_bkt1"
const W02TimeBkt2 = "w02_time_bkt2"
const W02TimeBkt3 = "w02_time_bkt3"
const W03Id = "w03_id"
const W03Status = "w03_status"
const W03WtBktEmpt = "w03_wt_bkt_empt"
const W03WtBktFull = "w03_wt_bkt_full"
const W03TimePicklingS1 = "w03_time_pickling_s1"
const W03TimePicklingS2 = "w03_time_pickling_s2"
const W03TimePicklingS3 = "w03_time_pickling_s3"
const W03TimeDegreaseS1 = "w03_time_degrease_s1"
const W03TimeDegreaseS2 = "w03_time_degrease_s2"
const W03TimeDegreaseS3 = "w03_time_degrease_s3"
const W03TimeHotWash1 = "w03_time_hot_wash1"
const W03TimeHotWash2 = "w03_time_hot_wash2"
const W03TimeHotWash3 = "w03_time_hot_wash3"
const W03TimeDegrease1 = "w03_time_degrease1"
const W03TimeDegrease2 = "w03_time_degrease2"
const W03TimeDegrease3 = "w03_time_degrease3"
const W03TimeDwash1 = "w03_time_dwash1"
const W03TimeDwash2 = "w03_time_dwash2"
const W03TimeDwash3 = "w03_time_dwash3"
const W03TimeFlux1 = "w03_time_flux1"
const W03TimeFlux2 = "w03_time_flux2"
const W03TimeFlux3 = "w03_time_flux3"
const W03TimeDry1 = "w03_time_dry1"
const W03TimeDry2 = "w03_time_dry2"
const W03TimeDry3 = "w03_time_dry3"
const W03TimePickling11 = "w03_time_pickling1_1"
const W03TimePickling12 = "w03_time_pickling1_2"
const W03TimePickling13 = "w03_time_pickling1_3"
const W03TimePickling21 = "w03_time_pickling2_1"
const W03TimePickling22 = "w03_time_pickling2_2"
const W03TimePickling23 = "w03_time_pickling2_3"
const W03TimePwash1 = "w03_time_pwash1"
const W03TimePwash2 = "w03_time_pwash2"
const W03TimePwash3 = "w03_time_pwash3"
const W03TimeSflux1 = "w03_time_sflux1"
const W03TimeSflux2 = "w03_time_sflux2"
const W03TimeSflux3 = "w03_time_sflux3"
const W03TimeBkt1 = "w03_time_bkt1"
const W03TimeBkt2 = "w03_time_bkt2"
const W03TimeBkt3 = "w03_time_bkt3"
const W04Id = "w04_id"
const W04Status = "w04_status"
const W04WtBktEmpt = "w04_wt_bkt_empt"
const W04WtBktFull = "w04_wt_bkt_full"
const W04TimePicklingS1 = "w04_time_pickling_s1"
const W04TimePicklingS2 = "w04_time_pickling_s2"
const W04TimePicklingS3 = "w04_time_pickling_s3"
const W04TimeDegreaseS1 = "w04_time_degrease_s1"
const W04TimeDegreaseS2 = "w04_time_degrease_s2"
const W04TimeDegreaseS3 = "w04_time_degrease_s3"
const W04TimeHotWash1 = "w04_time_hot_wash1"
const W04TimeHotWash2 = "w04_time_hot_wash2"
const W04TimeHotWash3 = "w04_time_hot_wash3"
const W04TimeDegrease1 = "w04_time_degrease1"
const W04TimeDegrease2 = "w04_time_degrease2"
const W04TimeDegrease3 = "w04_time_degrease3"
const W04TimeDwash1 = "w04_time_dwash1"
const W04TimeDwash2 = "w04_time_dwash2"
const W04TimeDwash3 = "w04_time_dwash3"
const W04TimeFlux1 = "w04_time_flux1"
const W04TimeFlux2 = "w04_time_flux2"
const W04TimeFlux3 = "w04_time_flux3"
const W04TimeDry1 = "w04_time_dry1"
const W04TimeDry2 = "w04_time_dry2"
const W04TimeDry3 = "w04_time_dry3"
const W04TimePickling11 = "w04_time_pickling1_1"
const W04TimePickling12 = "w04_time_pickling1_2"
const W04TimePickling13 = "w04_time_pickling1_3"
const W04TimePickling21 = "w04_time_pickling2_1"
const W04TimePickling22 = "w04_time_pickling2_2"
const W04TimePickling23 = "w04_time_pickling2_3"
const W04TimePwash1 = "w04_time_pwash1"
const W04TimePwash2 = "w04_time_pwash2"
const W04TimePwash3 = "w04_time_pwash3"
const W04TimeSflux1 = "w04_time_sflux1"
const W04TimeSflux2 = "w04_time_sflux2"
const W04TimeSflux3 = "w04_time_sflux3"
const W04TimeBkt1 = "w04_time_bkt1"
const W04TimeBkt2 = "w04_time_bkt2"
const W04TimeBkt3 = "w04_time_bkt3"
const W05Id = "w05_id"
const W05Status = "w05_status"
const W05WtBktEmpt = "w05_wt_bkt_empt"
const W05WtBktFull = "w05_wt_bkt_full"
const W05TimePicklingS1 = "w05_time_pickling_s1"
const W05TimePicklingS2 = "w05_time_pickling_s2"
const W05TimePicklingS3 = "w05_time_pickling_s3"
const W05TimeDegreaseS1 = "w05_time_degrease_s1"
const W05TimeDegreaseS2 = "w05_time_degrease_s2"
const W05TimeDegreaseS3 = "w05_time_degrease_s3"
const W05TimeHotWash1 = "w05_time_hot_wash1"
const W05TimeHotWash2 = "w05_time_hot_wash2"
const W05TimeHotWash3 = "w05_time_hot_wash3"
const W05TimeDegrease1 = "w05_time_degrease1"
const W05TimeDegrease2 = "w05_time_degrease2"
const W05TimeDegrease3 = "w05_time_degrease3"
const W05TimeDwash1 = "w05_time_dwash1"
const W05TimeDwash2 = "w05_time_dwash2"
const W05TimeDwash3 = "w05_time_dwash3"
const W05TimeFlux1 = "w05_time_flux1"
const W05TimeFlux2 = "w05_time_flux2"
const W05TimeFlux3 = "w05_time_flux3"
const W05TimeDry1 = "w05_time_dry1"
const W05TimeDry2 = "w05_time_dry2"
const W05TimeDry3 = "w05_time_dry3"
const W05TimePickling11 = "w05_time_pickling1_1"
const W05TimePickling12 = "w05_time_pickling1_2"
const W05TimePickling13 = "w05_time_pickling1_3"
const W05TimePickling21 = "w05_time_pickling2_1"
const W05TimePickling22 = "w05_time_pickling2_2"
const W05TimePickling23 = "w05_time_pickling2_3"
const W05TimePwash1 = "w05_time_pwash1"
const W05TimePwash2 = "w05_time_pwash2"
const W05TimePwash3 = "w05_time_pwash3"
const W05TimeSflux1 = "w05_time_sflux1"
const W05TimeSflux2 = "w05_time_sflux2"
const W05TimeSflux3 = "w05_time_sflux3"
const W05TimeBkt1 = "w05_time_bkt1"
const W05TimeBkt2 = "w05_time_bkt2"
const W05TimeBkt3 = "w05_time_bkt3"

// W01Id,
// W01Status,
// W01WtBktEmpt,
// W01WtBktFull,
// W01TimePicklingS1,
// W01TimePicklingS2,
// W01TimePicklingS3,
// W01TimeDegreaseS1,
// W01TimeDegreaseS2,
// W01TimeDegreaseS3,
// W01TimeHotWash1,
// W01TimeHotWash2,
// W01TimeHotWash3,
// W01TimeDegrease1,
// W01TimeDegrease2,
// W01TimeDegrease3,
// W01TimeDwash1,
// W01TimeDwash2,
// W01TimeDwash3,
// W01TimeFlux1,
// W01TimeFlux2,
// W01TimeFlux3,
// W01TimeDry1,
// W01TimeDry2,
// W01TimeDry3,
// W01TimePickling11,
// W01TimePickling12,
// W01TimePickling13,
// W01TimePickling21,
// W01TimePickling22,
// W01TimePickling23,
// W01TimePwash1,
// W01TimePwash2,
// W01TimePwash3,
// W01TimeSflux1,
// W01TimeSflux2,
// W01TimeSflux3,
// W01TimeBkt1,
// W01TimeBkt2,
// W01TimeBkt3

var WIDTags = []string{
	W01Id, W02Id, W03Id, W04Id, W05Id}

var WStatTags = []string{
	W01Status, W02Status, W03Status, W04Status, W05Status}

var W01Metrics = []string{
	W01TimePicklingS1, W01TimePicklingS2, W01TimePicklingS3,
	W01TimeDegreaseS1, W01TimeDegreaseS2, W01TimeDegreaseS3,
	W01TimeHotWash1, W01TimeHotWash2, W01TimeHotWash3,
	W01TimeDegrease1, W01TimeDegrease2, W01TimeDegrease3,
	W01TimeDwash1, W01TimeDwash2, W01TimeDwash3,
	W01TimeFlux1, W01TimeFlux2, W01TimeFlux3,
	W01TimeDry1, W01TimeDry2, W01TimeDry3,
	W01TimePickling11, W01TimePickling12, W01TimePickling13,
	W01TimePickling21, W01TimePickling22, W01TimePickling23,
	W01TimePwash1, W01TimePwash2, W01TimePwash3,
	W01TimeSflux1, W01TimeSflux2, W01TimeSflux3,
	W01TimeBkt1, W01TimeBkt2, W01TimeBkt3}

var W02Metrics = []string{
	// W02WtBktEmpt, W02WtBktFull, W02TimeDegrease1, W02TimeDegrease2, W02TimeDegrease3,
	// W02TimePickling1, W02TimePickling2, W02TimePickling3, W02TimeFlux,
	// Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	// W02WtGood
}

var W03Metrics = []string{
	// W03WtBktEmpt, W03WtBktFull, W03TimeDegrease1, W03TimeDegrease2, W03TimeDegrease3,
	// W03TimePickling1, W03TimePickling2, W03TimePickling3, W03TimeFlux,
	// Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	// W03WtGood
}

var W04Metrics = []string{
	// W04WtBktEmpt, W04WtBktFull, W04TimeDegrease1, W04TimeDegrease2, W04TimeDegrease3,
	// W04TimePickling1, W04TimePickling2, W04TimePickling3, W04TimeFlux,
	// Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	// W04WtGood
}

var W05Metrics = []string{
	// W05WtBktEmpt, W05WtBktFull, W05TimeDegrease1, W05TimeDegrease2, W05TimeDegrease3,
	// W05TimePickling1, W05TimePickling2, W05TimePickling3, W05TimeFlux,
	// Preproc01Temp, Preproc02Temp, Preproc03Temp, Preproc04Temp, Preproc05Temp,
	// W05WtGood
}

var jobMetrics = [][]string{
	W01Metrics, W02Metrics, W03Metrics, W04Metrics, W05Metrics}

// GalMeasureBufferTags

const Gal01Id = "gal01_id"
const Gal01TempSet = "gal01_temp_set"
const Gal01TempAct = "gal01_temp_act"

// PreMeasureBufferTags

const Preproc01Id = "preproc01_id"
const Preproc01Temp = "preproc01_temp"
const Preproc01Time = "preproc01_time"
const Preproc01Ph = "preproc01_ph"
const Preproc01Conduct = "preproc01_conduct"
const Preproc02Id = "preproc02_id"
const Preproc02Temp = "preproc02_temp"
const Preproc02Time = "preproc02_time"
const Preproc02Ph = "preproc02_ph"
const Preproc02Conduct = "preproc02_conduct"
const Preproc03Id = "preproc03_id"
const Preproc03Temp = "preproc03_temp"
const Preproc03Time = "preproc03_time"
const Preproc03Ph = "preproc03_ph"
const Preproc03Conduct = "preproc03_conduct"
const Preproc04Id = "preproc04_id"
const Preproc04Temp = "preproc04_temp"
const Preproc04Time = "preproc04_time"
const Preproc04Ph = "preproc04_ph"
const Preproc04Conduct = "preproc04_conduct"
const Preproc05Id = "preproc05_id"
const Preproc05Temp = "preproc05_temp"
const Preproc05Time = "preproc05_time"
const Preproc05Ph = "preproc05_ph"
const Preproc05Conduct = "preproc05_conduct"
const Preproc06Id = "preproc06_id"
const Preproc06Temp = "preproc06_temp"
const Preproc06Time = "preproc06_time"
const Preproc06Ph = "preproc06_ph"
const Preproc06Conduct = "preproc06_conduct"
const Preproc07Id = "preproc07_id"
const Preproc07Temp = "preproc07_temp"
const Preproc07Time = "preproc07_time"
const Preproc07Ph = "preproc07_ph"
const Preproc07Conduct = "preproc07_conduct"
const Preproc08Id = "preproc08_id"
const Preproc08Temp = "preproc08_temp"
const Preproc08Time = "preproc08_time"
const Preproc08Ph = "preproc08_ph"
const Preproc08Conduct = "preproc08_conduct"
const Preproc09Id = "preproc09_id"
const Preproc09Temp = "preproc09_temp"
const Preproc09Time = "preproc09_time"
const Preproc09Ph = "preproc09_ph"
const Preproc09Conduct = "preproc09_conduct"
const Preproc10Id = "preproc10_id"
const Preproc10Temp = "preproc10_temp"
const Preproc10Time = "preproc10_time"
const Preproc10Ph = "preproc10_ph"
const Preproc10Conduct = "preproc10_conduct"
const Preproc11Id = "preproc11_id"
const Preproc11Temp = "preproc11_temp"
const Preproc11Time = "preproc11_time"
const Preproc11Ph = "preproc11_ph"
const Preproc11Conduct = "preproc11_conduct"
const Preproc12Id = "preproc12_id"
const Preproc12Temp = "preproc12_temp"
const Preproc12Time = "preproc12_time"
const Preproc12Ph = "preproc12_ph"
const Preproc12Conduct = "preproc12_conduct"
const Preproc13Id = "preproc13_id"
const Preproc13Temp = "preproc13_temp"
const Preproc13Time = "preproc13_time"
const Preproc13Ph = "preproc13_ph"
const Preproc13Conduct = "preproc13_conduct"
const Preproc14Id = "preproc14_id"
const Preproc14Temp = "preproc14_temp"
const Preproc14Time = "preproc14_time"
const Preproc14Ph = "preproc14_ph"
const Preproc14Conduct = "preproc14_conduct"
const Preproc15Id = "preproc15_id"
const Preproc15Temp = "preproc15_temp"
const Preproc15Time = "preproc15_time"
const Preproc15Ph = "preproc15_ph"
const Preproc15Conduct = "preproc15_conduct"
const Preproc16Id = "preproc16_id"
const Preproc16Temp = "preproc16_temp"
const Preproc16Time = "preproc16_time"
const Preproc16Ph = "preproc16_ph"
const Preproc16Conduct = "preproc16_conduct"
const Preproc17Id = "preproc17_id"
const Preproc17Temp = "preproc17_temp"
const Preproc17Time = "preproc17_time"
const Preproc17Ph = "preproc17_ph"
const Preproc17Conduct = "preproc17_conduct"
const Preproc18Id = "preproc18_id"
const Preproc18Temp = "preproc18_temp"
const Preproc18Time = "preproc18_time"
const Preproc18Ph = "preproc18_ph"
const Preproc18Conduct = "preproc18_conduct"
const Preproc19Id = "preproc19_id"
const Preproc19Temp = "preproc19_temp"
const Preproc19Time = "preproc19_time"
const Preproc19Ph = "preproc19_ph"
const Preproc19Conduct = "preproc19_conduct"
const Preproc20Id = "preproc20_id"
const Preproc20Temp = "preproc20_temp"
const Preproc20Time = "preproc20_time"
const Preproc20Ph = "preproc20_ph"
const Preproc20Conduct = "preproc20_conduct"
const Preproc21Id = "preproc21_id"
const Preproc21Temp = "preproc21_temp"
const Preproc21Time = "preproc21_time"
const Preproc21Ph = "preproc21_ph"
const Preproc21Conduct = "preproc21_conduct"
const Preproc22Id = "preproc22_id"
const Preproc22Temp = "preproc22_temp"
const Preproc22Time = "preproc22_time"
const Preproc22Ph = "preproc22_ph"
const Preproc22Conduct = "preproc22_conduct"
const Preproc23Id = "preproc23_id"
const Preproc23Temp = "preproc23_temp"
const Preproc23Time = "preproc23_time"
const Preproc23Ph = "preproc23_ph"
const Preproc23Conduct = "preproc23_conduct"
const Preproc24Id = "preproc24_id"
const Preproc24Temp = "preproc24_temp"
const Preproc24Time = "preproc24_time"
const Preproc24Ph = "preproc24_ph"
const Preproc24Conduct = "preproc24_conduct"
const Preproc25Id = "preproc25_id"
const Preproc25Temp = "preproc25_temp"
const Preproc25Time = "preproc25_time"
const Preproc25Ph = "preproc25_ph"
const Preproc25Conduct = "preproc25_conduct"
const Preproc26Id = "preproc26_id"
const Preproc26Temp = "preproc26_temp"
const Preproc26Time = "preproc26_time"
const Preproc26Ph = "preproc26_ph"
const Preproc26Conduct = "preproc26_conduct"
const Preproc27Id = "preproc27_id"
const Preproc27Temp = "preproc27_temp"
const Preproc27Time = "preproc27_time"
const Preproc27Ph = "preproc27_ph"
const Preproc27Conduct = "preproc27_conduct"
const Preproc28Id = "preproc28_id"
const Preproc28Temp = "preproc28_temp"
const Preproc28Time = "preproc28_time"
const Preproc28Ph = "preproc28_ph"
const Preproc28Conduct = "preproc28_conduct"
const Preproc29Id = "preproc29_id"
const Preproc29Temp = "preproc29_temp"
const Preproc29Time = "preproc29_time"
const Preproc29Ph = "preproc29_ph"
const Preproc29Conduct = "preproc29_conduct"
const Preproc30Id = "preproc30_id"
const Preproc30Temp = "preproc30_temp"
const Preproc30Time = "preproc30_time"
const Preproc30Ph = "preproc30_ph"
const Preproc30Conduct = "preproc30_conduct"
const Preproc31Id = "preproc31_id"
const Preproc31Temp = "preproc31_temp"
const Preproc31Time = "preproc31_time"
const Preproc31Ph = "preproc31_ph"
const Preproc31Conduct = "preproc31_conduct"
const Preproc32Id = "preproc32_id"
const Preproc32Temp = "preproc32_temp"
const Preproc32Time = "preproc32_time"
const Preproc32Ph = "preproc32_ph"
const Preproc32Conduct = "preproc32_conduct"

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
