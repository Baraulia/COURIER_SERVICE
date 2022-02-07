package Models

import "github.com/Baraulia/COURIER_SERVICE/db"

func GetCouriers(Couriers []db.Courier) []db.Courier {
	Couriers = []db.Courier{}
	db.GetCouriersFromDB(&Couriers)
	return Couriers
}
