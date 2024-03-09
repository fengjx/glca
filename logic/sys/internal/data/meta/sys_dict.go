// Code generated by "daox.gen"; DO NOT EDIT.
package meta

import (
    "github.com/fengjx/daox/sqlbuilder"
    "github.com/fengjx/daox/sqlbuilder/ql"


    "time"

)



// SysDictM 系统字典表
type SysDictM struct {
    ID string
    Group string
    GroupName string
    Value string
    Label string
    Status string
    Remark string
    Utime string
    Ctime string
}

func (m SysDictM) TableName() string {
    return "sys_dict"
}

func (m SysDictM) PrimaryKey() string {
    return "id"
}

func (m SysDictM) Columns() []string {
	return []string{
        "id",
        "group",
        "group_name",
        "value",
        "label",
        "status",
        "remark",
        "utime",
        "ctime",
    }
}

var SysDictMeta = SysDictM{
    ID: "id",
    Group: "group",
    GroupName: "group_name",
    Value: "value",
    Label: "label",
    Status: "status",
    Remark: "remark",
    Utime: "utime",
    Ctime: "ctime",
}




func (m SysDictM) IdIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.ID).In(args...)
}

func (m SysDictM) IdNotIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.ID).NotIn(args...)
}

func (m SysDictM) IdEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).EQ(val)
}

func (m SysDictM) IdNotEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).NotEQ(val)
}

func (m SysDictM) IdLT(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).LT(val)
}

func (m SysDictM) IdLTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).LTEQ(val)
}

func (m SysDictM) IdGT(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).GT(val)
}

func (m SysDictM) IdGTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).GTEQ(val)
}

func (m SysDictM) IdLike(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).Like(val)
}

func (m SysDictM) IdNotLike(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).NotLike(val)
}

func (m SysDictM) IdDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.ID)
}

func (m SysDictM) IdAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.ID)
}



func (m SysDictM) GroupIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Group).In(args...)
}

func (m SysDictM) GroupNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Group).NotIn(args...)
}

func (m SysDictM) GroupEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Group).EQ(val)
}

func (m SysDictM) GroupNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Group).NotEQ(val)
}

func (m SysDictM) GroupLT(val string) sqlbuilder.Column {
	return ql.Col(m.Group).LT(val)
}

func (m SysDictM) GroupLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Group).LTEQ(val)
}

func (m SysDictM) GroupGT(val string) sqlbuilder.Column {
	return ql.Col(m.Group).GT(val)
}

func (m SysDictM) GroupGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Group).GTEQ(val)
}

func (m SysDictM) GroupLike(val string) sqlbuilder.Column {
	return ql.Col(m.Group).Like(val)
}

func (m SysDictM) GroupNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Group).NotLike(val)
}

func (m SysDictM) GroupDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Group)
}

func (m SysDictM) GroupAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Group)
}



func (m SysDictM) GroupNameIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.GroupName).In(args...)
}

func (m SysDictM) GroupNameNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.GroupName).NotIn(args...)
}

func (m SysDictM) GroupNameEQ(val string) sqlbuilder.Column {
	return ql.Col(m.GroupName).EQ(val)
}

func (m SysDictM) GroupNameNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.GroupName).NotEQ(val)
}

func (m SysDictM) GroupNameLT(val string) sqlbuilder.Column {
	return ql.Col(m.GroupName).LT(val)
}

func (m SysDictM) GroupNameLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.GroupName).LTEQ(val)
}

func (m SysDictM) GroupNameGT(val string) sqlbuilder.Column {
	return ql.Col(m.GroupName).GT(val)
}

func (m SysDictM) GroupNameGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.GroupName).GTEQ(val)
}

func (m SysDictM) GroupNameLike(val string) sqlbuilder.Column {
	return ql.Col(m.GroupName).Like(val)
}

func (m SysDictM) GroupNameNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.GroupName).NotLike(val)
}

func (m SysDictM) GroupNameDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.GroupName)
}

func (m SysDictM) GroupNameAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.GroupName)
}



func (m SysDictM) ValueIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Value).In(args...)
}

func (m SysDictM) ValueNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Value).NotIn(args...)
}

func (m SysDictM) ValueEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Value).EQ(val)
}

func (m SysDictM) ValueNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Value).NotEQ(val)
}

func (m SysDictM) ValueLT(val string) sqlbuilder.Column {
	return ql.Col(m.Value).LT(val)
}

func (m SysDictM) ValueLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Value).LTEQ(val)
}

func (m SysDictM) ValueGT(val string) sqlbuilder.Column {
	return ql.Col(m.Value).GT(val)
}

func (m SysDictM) ValueGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Value).GTEQ(val)
}

func (m SysDictM) ValueLike(val string) sqlbuilder.Column {
	return ql.Col(m.Value).Like(val)
}

func (m SysDictM) ValueNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Value).NotLike(val)
}

func (m SysDictM) ValueDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Value)
}

func (m SysDictM) ValueAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Value)
}



func (m SysDictM) LabelIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Label).In(args...)
}

func (m SysDictM) LabelNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Label).NotIn(args...)
}

func (m SysDictM) LabelEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Label).EQ(val)
}

func (m SysDictM) LabelNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Label).NotEQ(val)
}

func (m SysDictM) LabelLT(val string) sqlbuilder.Column {
	return ql.Col(m.Label).LT(val)
}

func (m SysDictM) LabelLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Label).LTEQ(val)
}

func (m SysDictM) LabelGT(val string) sqlbuilder.Column {
	return ql.Col(m.Label).GT(val)
}

func (m SysDictM) LabelGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Label).GTEQ(val)
}

func (m SysDictM) LabelLike(val string) sqlbuilder.Column {
	return ql.Col(m.Label).Like(val)
}

func (m SysDictM) LabelNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Label).NotLike(val)
}

func (m SysDictM) LabelDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Label)
}

func (m SysDictM) LabelAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Label)
}



func (m SysDictM) StatusIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Status).In(args...)
}

func (m SysDictM) StatusNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Status).NotIn(args...)
}

func (m SysDictM) StatusEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Status).EQ(val)
}

func (m SysDictM) StatusNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Status).NotEQ(val)
}

func (m SysDictM) StatusLT(val string) sqlbuilder.Column {
	return ql.Col(m.Status).LT(val)
}

func (m SysDictM) StatusLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Status).LTEQ(val)
}

func (m SysDictM) StatusGT(val string) sqlbuilder.Column {
	return ql.Col(m.Status).GT(val)
}

func (m SysDictM) StatusGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Status).GTEQ(val)
}

func (m SysDictM) StatusLike(val string) sqlbuilder.Column {
	return ql.Col(m.Status).Like(val)
}

func (m SysDictM) StatusNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Status).NotLike(val)
}

func (m SysDictM) StatusDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Status)
}

func (m SysDictM) StatusAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Status)
}



func (m SysDictM) RemarkIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Remark).In(args...)
}

func (m SysDictM) RemarkNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Remark).NotIn(args...)
}

func (m SysDictM) RemarkEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).EQ(val)
}

func (m SysDictM) RemarkNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).NotEQ(val)
}

func (m SysDictM) RemarkLT(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).LT(val)
}

func (m SysDictM) RemarkLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).LTEQ(val)
}

func (m SysDictM) RemarkGT(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).GT(val)
}

func (m SysDictM) RemarkGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).GTEQ(val)
}

func (m SysDictM) RemarkLike(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).Like(val)
}

func (m SysDictM) RemarkNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).NotLike(val)
}

func (m SysDictM) RemarkDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Remark)
}

func (m SysDictM) RemarkAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Remark)
}



func (m SysDictM) UtimeIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Utime).In(args...)
}

func (m SysDictM) UtimeNotIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Utime).NotIn(args...)
}

func (m SysDictM) UtimeEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).EQ(val)
}

func (m SysDictM) UtimeNotEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).NotEQ(val)
}

func (m SysDictM) UtimeLT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).LT(val)
}

func (m SysDictM) UtimeLTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).LTEQ(val)
}

func (m SysDictM) UtimeGT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).GT(val)
}

func (m SysDictM) UtimeGTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).GTEQ(val)
}

func (m SysDictM) UtimeLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).Like(val)
}

func (m SysDictM) UtimeNotLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).NotLike(val)
}

func (m SysDictM) UtimeDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Utime)
}

func (m SysDictM) UtimeAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Utime)
}



func (m SysDictM) CtimeIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Ctime).In(args...)
}

func (m SysDictM) CtimeNotIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Ctime).NotIn(args...)
}

func (m SysDictM) CtimeEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).EQ(val)
}

func (m SysDictM) CtimeNotEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).NotEQ(val)
}

func (m SysDictM) CtimeLT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).LT(val)
}

func (m SysDictM) CtimeLTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).LTEQ(val)
}

func (m SysDictM) CtimeGT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).GT(val)
}

func (m SysDictM) CtimeGTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).GTEQ(val)
}

func (m SysDictM) CtimeLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).Like(val)
}

func (m SysDictM) CtimeNotLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).NotLike(val)
}

func (m SysDictM) CtimeDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Ctime)
}

func (m SysDictM) CtimeAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Ctime)
}
