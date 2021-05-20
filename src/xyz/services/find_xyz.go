package services

import (
	"context"
	"math"
	"q-change-assignment/src/xyz/dao"
)

func calculateValueFromIndex(i int64) int64 {
	// คิดไม่ออกว่าตัวเลขมันมาจากสมการอะไร เลยเอาตัวข้อมูลไปใส่ใน https://www.dcode.fr/function-equation-finder
	// ตอนใส่ข้อมูล กำหนดให้ x เป็นลำดับตัวเลข และ y เป็นค่าของตัวเลข (เอาเฉพาะที่ไม่ติดตัวแปร)
	// จะได้ออกมาเป็นคู่อันดับแบบนี้ (1,1), (3,8), (4,17), (7, 78), (8, 113)
	// และเลือก TARGET AND CALCULATION METHOD เป็น POWER (INCLUDING INVERSE AND NTH ROOT) USING CURVE FITTING)
	// จริงๆลองหลายๆตัวแล้ว อันนี้ดูจะพอดีที่สุด
	// หลังจากใส่ข้อมูลแล้วก็กด FIND A MATCHING EQUATION แล้วมันก็ออกมาเป็นสมการข้างล่าง
	return int64(math.Round(math.Pow(float64(i), 2.78833)*0.340517 + 0.698018))
}

func (sv *XYZService) FindXYZ(ctx context.Context) (dao.FindXYZResponse, error) {
	return dao.FindXYZResponse{
		X: calculateValueFromIndex(2),
		Y: calculateValueFromIndex(5),
		Z: calculateValueFromIndex(6),
	}, nil
}
