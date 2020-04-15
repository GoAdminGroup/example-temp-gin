package demo

import (
	"fmt"
	"github.com/GoAdminGroup/example-temp-gin/api/public"
	"github.com/GoAdminGroup/example-temp-gin/model/dbglobal"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/gin-gonic/gin"
	"github.com/sinlovgo/timezone"
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
				public.JSONErr(ctx, 0, bizPath, err)
				return
			}
			countRes = count
		} else {
			formDayNum, err := strconv.Atoi(fromDay)
			if err != nil {
				public.JSONErr(ctx, 0, bizPath, err)
				return
			}
			toDayNum, err := strconv.Atoi(toDay)
			if err != nil {
				public.JSONErr(ctx, 0, bizPath, err)
				return
			}

			fromUTCZero := timezone.CalcDayMicro(formDayNum)
			toUTCZero := timezone.CalcDayMicro(toDayNum)
			if fromUTCZero == "" || toUTCZero == "" {
				public.JSONErr(ctx, 0, bizPath, fmt.Errorf("time format error"))
				return
			}
			countByTimeRange, err := dbglobal.TableCountByTimeRange("demo_student", "created_at", fromUTCZero, toUTCZero)
			if err != nil {
				public.JSONErr(ctx, 0, bizPath, err)
				return
			}
			countRes = countByTimeRange
		}

		public.JSONSucc(ctx, struct {
			StudentCount int64 `json:"student_count"`
		}{
			StudentCount: countRes,
		})
	})
}
