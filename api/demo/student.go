package demo

import (
	"fmt"
	pubilc "github.com/GoAdminGroup/example-temp-gin/api/public"
	"github.com/GoAdminGroup/example-temp-gin/model/dbglobal"
	"github.com/GoAdminGroup/example-temp-gin/util/timestamp"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/gin-gonic/gin"
	"strconv"
)

func StudentCount(g *gin.Engine, eng *engine.Engine) gin.IRoutes {
	bizPath := "/demo/studentCount"
	return g.GET(bizPath, func(ctx *gin.Context) {

		fromDay := ctx.Query("created_at_day_from")
		toDay := ctx.Query("created_at_day_to")

		var countRes int64
		if fromDay == "" || toDay == "" {
			count, err := dbglobal.TableCount("demo_student")
			if err != nil {
				pubilc.JSONErr(ctx, 0, bizPath, err)
				return
			}
			countRes = count
		} else {
			formDayNum, err := strconv.Atoi(fromDay)
			if err != nil {
				pubilc.JSONErr(ctx, 0, bizPath, err)
				return
			}
			toDayNum, err := strconv.Atoi(toDay)
			if err != nil {
				pubilc.JSONErr(ctx, 0, bizPath, err)
				return
			}

			fromUTCZero := timestamp.CalcDayMicroUTCZero(formDayNum)
			toUTCZero := timestamp.CalcDayMicroUTCZero(toDayNum)
			if fromUTCZero == "" || toUTCZero == "" {
				pubilc.JSONErr(ctx, 0, bizPath, fmt.Errorf("time format error"))
				return
			}
			countByTimeRange, err := dbglobal.TableCountByTimeRange("demo_student", "created_at", fromUTCZero, toUTCZero)
			if err != nil {
				pubilc.JSONErr(ctx, 0, bizPath, err)
				return
			}
			countRes = countByTimeRange
		}

		pubilc.JSONSucc(ctx, struct {
			StudentCount int64 `json:"student_count"`
		}{
			StudentCount: countRes,
		})
	})
}
