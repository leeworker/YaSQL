/*
@Time    :   2022/06/29 15:30:31
@Author  :   zongfei.fu
*/

package controllers

import (
	"goInsight/internal/pkg/kv"
	"goInsight/internal/pkg/utils"
)

type RuleHint struct {
	Summary        []string `json:"summary"` // 规则摘要
	AffectedRows   int      `json:"affected_rows"`
	IsSkipNextStep bool     // 是否跳过接下来的检查步骤
	DB             *utils.DB
	KV             *kv.KVCache
	Query          string // 原始SQL
	MergeAlter     string
	// AuditConfig    *config.AuditConfiguration
}
