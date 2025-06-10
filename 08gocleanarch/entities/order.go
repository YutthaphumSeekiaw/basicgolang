package entities

// Idea ของ Entities คือการประกาศโครงสร้างของ Data ไว้ ดังนั้น code ของ
// Entities จึงเป็นเพียงการประกาศ struct ของข้อมูลเพื่อเป็นการบอกว่าข้อมูลมีหน้าตาออกมาประมาณไหน
type Order struct {
	ID    uint
	Total float64
}
