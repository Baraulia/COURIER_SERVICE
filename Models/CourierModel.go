package Models

import "github.com/Baraulia/COURIER_SERVICE/db"

func GetCouriers(Couriers []db.SmallInfo) []db.SmallInfo {
	Couriers = []db.SmallInfo{}
	db.GetCouriersFromDB(&Couriers)
	return Couriers
}

func GetOneCourier(id int) db.SmallInfo {
	Courier := db.SmallInfo{}
	db.GetOneCourierFromDB(&Courier, id)
	return Courier
}
