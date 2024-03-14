// -------------------------------------------------------------------
// Generated by 365admin-publish 
// Service wrapper  v1
// -------------------------------------------------------------------
/*
---
title: Web deploy to production
---
*/
package cmds
import (
"context"
"encoding/json"
"os"
"path"
"github.com/nats-io/nats.go"
"github.com/nats-io/nats.go/micro"
"github.com/spf13/cobra"
"github.com/365admin/magicbox-superset/schemas"
"github.com/365admin/magicbox-superset/execution"
"github.com/365admin/magicbox-superset/utils"
)
func ProvisionWebdeployproductionPost(ctx context.Context, args  []string)  ( *string, error) {

_, pwsherr := execution.ExecutePowerShell("john","*","magicbox-superset","60-provision","10-web.ps1","" )
if (pwsherr != nil) {
	return nil,pwsherr
}
return nil, nil
	
// end result mapping

func init(){
	
}
}