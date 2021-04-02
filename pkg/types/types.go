package types

//Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.).
type Money int64

//Category представляет собой категорию, в который был совершён платёж (авто, аптеки, рестораны и т.д.).
type Category string

//Status представляет собой статус платежа.
type Status string

//Предопределённые статусы платежой.
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

//Payment представляет информацию о платежа.
type Payment struct{
	ID        int
	Amount    Money
	Category  Category
	Status    Status
}