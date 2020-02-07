package model

import (
	"fmt"
	"strconv"
)

// 用来提供测试用例。之后应该用真实数据替代
func tryBedModel() {

	depart := &Depart{ID: 13}
	MYSQL.Create(depart)
	unit := &Unit{
		UID:         "13E3",
		Depart:      *depart,
		DepartRefer: depart.ID,
	}
	dorm := make([]Dorm, 40)
	bed := make([]Bed, 100)
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			for w := 0; w < 4; w++ {
				dorm[i*j+w] = Dorm{
					DormNumber: "13E" + "3" + strconv.Itoa(i+1) + strconv.Itoa(j+1) + strconv.Itoa(w+1),
					Floor:      uint8(i + 1),
					Unit:       *unit,
					UnitRefer:  unit.UID,
					Size:       uint8(w + 1),
					UnUseSize:  uint8((w + 1) - ((w + 2) / 2)),
					Gender:     1,
				}
				MYSQL.Create(&dorm[i*j+w])
				for d := 0; d < w+1; d++ {
					bed[i*j*w+d] = Bed{
						BedNumber: dorm[i*j+w].DormNumber + strconv.Itoa(d+1),
						Remain:    uint8(d % 2),
						Dorm:      dorm[i*j+w],
						DormRefer: dorm[i*j+w].DormNumber,
					}
					MYSQL.Create(&bed[i*j*w+d])
				}
			}
		}
	}

	depart2 := &Depart{ID: 14}
	MYSQL.Create(depart2)
	unit2 := &Unit{
		UID:         "14E2",
		Depart:      *depart2,
		DepartRefer: depart.ID,
	}
	dorm2 := make([]Dorm, 40)
	bed2 := make([]Bed, 100)
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			for w := 0; w < 4; w++ {
				dorm2[i*j+w] = Dorm{
					DormNumber: unit2.UID + strconv.Itoa(i+1) + strconv.Itoa(j+1) + strconv.Itoa(w+1),
					Floor:      uint8(i + 1),
					Unit:       *unit2,
					UnitRefer:  unit2.UID,
					Size:       uint8(w + 1),
					UnUseSize:  uint8((w + 1) - ((w + 2) / 2)),
					Gender:     0,
				}
				MYSQL.Create(&dorm2[i*j+w])
				for d := 0; d < w+1; d++ {
					bed2[i*j*w+d] = Bed{
						BedNumber: dorm2[i*j+w].DormNumber + strconv.Itoa(d+1),
						Remain:    uint8(d % 2),
						Dorm:      dorm2[i*j+w],
						DormRefer: dorm2[i*j+w].DormNumber,
					}
					MYSQL.Create(&bed2[i*j*w+d])
				}
			}
		}
	}
	i, _ := GetDepartBedNumber(13)
	fmt.Print(i)
}
