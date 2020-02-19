package tables

import (
	"os"
	"strings"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMeetingMeetsTable(ctx *context.Context) table.Table {

    meetingMeetsTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

    // Info
	info := meetingMeetsTable.GetInfo()
	info.HideExportButton()
	info.SetSortAsc()

	info.AddField("ID","id", db.Int)
	info.AddField("类型","type", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
	    if model.Value == "1" {
	        return "普通"
	    }
	    if model.Value == "2" {
	        return "定期"
	    }
	    return "未定义"
	})
	info.AddField("头图","banner", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			filePath, err := os.Getwd();
			if err != nil {
				panic(err)
			}
			value.Value = filePath + "/" + config.Get().Store.Prefix + "/" + value.Value
			return template.Default().Image().
				SetSrc(template.HTML(value.Value)).
				SetHeight("60").SetWidth("60").
				WithModal().GetContent()
		})
	info.AddField("标题","title", db.Varchar).
		FieldLimit(20).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
			FieldFilterProcess(func(s string) string {
		    return strings.TrimSpace(s)
		})
	info.AddField("开始时间","start_at", db.Timestamp)
	info.AddField("地点","place", db.Varchar).FieldLimit(20)
	info.AddField("费用","fee", db.Int)
	info.AddField("人数","person", db.Int)
	info.AddField("内容","content", db.Text).FieldLimit(20)
	info.AddField("状态","state", db.Tinyint).
		FieldSortable().
		FieldDisplay(func(model types.FieldModel) interface{} {
		    if model.Value == "1" {
		        return "有效"
		    }
		    if model.Value == "2" {
		        return "无效"
		    }
		    return "未定义"
		})
	info.SetTable("meeting_meets").SetTitle("活动").SetDescription("活动信息列表")

	// Detail
	detail := meetingMeetsTable.GetDetail()
	detail.AddField("活动ID","id", db.Int)
	detail.AddField("活动类型","type", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
	    if model.Value == "1" {
	        return "普通活动"
	    }
	    if model.Value == "2" {
	        return "定期活动"
	    }
	    return "未定义"
	})
	detail.AddField("头图","banner", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			filePath, err := os.Getwd();
			if err != nil {
				panic(err)
			}
			value.Value = filePath + "/" + config.Get().Store.Prefix + "/" + value.Value
			return template.Default().Image().
				SetSrc(template.HTML(value.Value)).
				SetHeight("150").SetWidth("150").
				WithModal().GetContent()
		})
	detail.AddField("标题","title", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
			FieldFilterProcess(func(s string) string {
		    return strings.TrimSpace(s)
		}).
		FieldLimit(20)
	detail.AddField("开始时间","start_at", db.Timestamp)
	detail.AddField("地点","place", db.Varchar)
	detail.AddField("费用","fee", db.Int)
	detail.AddField("人数","person", db.Int)
	detail.AddField("内容","content", db.Text)
	detail.AddField("状态","state", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
	    if model.Value == "1" {
	        return "有效"
	    }
	    if model.Value == "2" {
	        return "无效"
	    }
	    return "未定义"
	})
	detail.SetTable("meeting_meets").SetTitle("活动").SetDescription("活动信息详情")


	// Form
	formList := meetingMeetsTable.GetForm()
	formList.AddField("ID","id",db.Int,form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("类型","type",db.Tinyint,form.SelectSingle).
		FieldDefault("1").
        FieldOptions([]map[string]string{ 
            {"field": "普通活动","value": "1"},
            {"field": "长期活动","value": "2"},
        })
	formList.AddField("顶部图片","banner",db.Varchar,form.File).FieldHelpMsg("图片小于3M")
	formList.AddField("标题","title",db.Varchar,form.Text).FieldMust()
	formList.AddField("开始时间","start_at",db.Timestamp,form.Datetime).FieldMust()
	formList.AddField("地点","place",db.Varchar,form.Text)
	formList.AddField("费用","fee",db.Int,form.Currency)
	formList.AddField("人数","person",db.Int,form.Number)
	formList.AddField("活动介绍","content",db.Text,form.RichText)
	formList.AddField("状态","state",db.Tinyint,form.SelectSingle).
		FieldDefault("1").
        FieldOptions([]map[string]string{ 
            {"field": "有效","value": "1"},
            {"field": "无效","value": "2"},
        })
	formList.SetTable("meeting_meets").SetTitle("活动").SetDescription("活动信息编辑")

	return meetingMeetsTable
}