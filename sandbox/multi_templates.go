// DBDeployer - The MySQL Sandbox
// Copyright © 2006-2020 Giuseppe Maxia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sandbox

// Templates for multiple sandboxes

var (
	startMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
echo "# executing 'start' on $SBDIR"
{{ range .Nodes }}
echo 'executing "start" on {{.NodeLabel}} {{.Node}}'
$SBDIR/{{.NodeLabel}}{{.Node}}/start "$@"
{{end}}
`

	restartMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
$SBDIR/stop_all
$SBDIR/start_all "$@"
`
	useMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
if [ "$1" = "" ]
then
  echo "syntax: $0 command"
  exit 1
fi

USE_SCRIPT=use
if [ -n "$ADMIN_USE" ]
then
    USE_SCRIPT=use_admin
fi
{{range .Nodes}}
echo "# server: {{.Node}} "
echo "$@" | $SBDIR/{{.NodeLabel}}{{.Node}}/$USE_SCRIPT $MYCLIENT_OPTIONS
{{end}}
`

	execMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
if [ "$1" = "" ]
then
  echo "syntax: $0 command"
  exit 1
fi

{{range .Nodes}}
echo "# server: {{.Node}} "
  cd $SBDIR/{{.NodeLabel}}{{.Node}}
  $@
  cd - > /dev/null
{{end}}
`

	metadataMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
if [ "$1" = "" ]
then
  echo "syntax: $0 keyword"
  exit 1
fi

{{range .Nodes}}
echo "# server: {{.Node}} "
$SBDIR/{{.NodeLabel}}{{.Node}}/metadata $@
{{end}}
`
	useMultiAdminTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}

export ADMIN_USE=1
$SBDIR/use_all $@
`
	stopMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
echo "# executing 'stop' on $SBDIR"
for node in {{.StopNodeList}}
do
    echo "executing 'stop' on {{.NodeLabel}}$node"
    $SBDIR/{{.NodeLabel}}$node/stop "$@"
done
`
	sendKillMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
echo "# executing 'send_kill' on $SBDIR"
{{ range .Nodes }}
echo 'executing "send_kill" on {{.NodeLabel}} {{.Node}}'
$SBDIR/{{.NodeLabel}}{{.Node}}/send_kill "$@"
{{end}}
`
	clearMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
echo "# executing 'clear' on $SBDIR"
{{range .Nodes}}
echo 'executing "clear" on {{.NodeLabel}} {{.Node}}'
$SBDIR/{{.NodeLabel}}{{.Node}}/clear "$@"
{{end}}
`
	statusMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
echo "MULTIPLE  $SBDIR"
{{ range .Nodes }}
nstatus=$($SBDIR/{{.NodeLabel}}{{.Node}}/status )
if [ -f $SBDIR/{{.NodeLabel}}{{.Node}}/data/mysql_sandbox{{.NodePort}}.pid ]
then
	nport=$($SBDIR/{{.NodeLabel}}{{.Node}}/use -BN -e "show variables like 'port'")
fi
echo "{{.NodeLabel}}{{.Node}} : $nstatus  -  $nport ({{.NodePort}})"
{{end}}
`
	testSbMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
echo "# executing 'test_sb' on $SBDIR"
{{ range .Nodes }}
echo 'executing "test_sb" on {{.NodeLabel}} {{.Node}}'
$SBDIR/{{.NodeLabel}}{{.Node}}/test_sb "$@"
exit_code=$?
if [ "$exit_code" != "0" ] ; then exit $exit_code ; fi
{{end}}
`

	nodeTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
$SBDIR/{{.NodeLabel}}{{.Node}}/use "$@"
`

	nodeAdminTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}
$SBDIR/{{.NodeLabel}}{{.Node}}/use_admin "$@"
`
	replicateFromMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}

replicate_from=$SBDIR/{{.NodeLabel}}1/replicate_from

if [ ! -x $replicate_from ]
then
    echo "$replicate_from not found"
    exit 1
fi

$replicate_from $@
`
	sysbenchMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}

sysbench=$SBDIR/{{.NodeLabel}}1/sysbench

if [ ! -x $sysbench ]
then
    echo "$sysbench not found"
    exit 1
fi

$sysbench $@
`

	sysbenchReadyMultiTemplate string = `#!{{.ShellPath}}
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
SBDIR={{.SandboxDir}}

sysbench_ready=$SBDIR/{{.NodeLabel}}1/sysbench_ready

if [ ! -x $sysbench_ready ]
then
    echo "$sysbench_ready not found"
    exit 1
fi

$sysbench_ready $@
`

	MultipleTemplates = TemplateCollection{
		"start_multi_template": TemplateDesc{
			Description: "Starts all nodes (with optional mysqld arguments)",
			Notes:       "",
			Contents:    startMultiTemplate,
		},
		"restart_multi_template": TemplateDesc{
			Description: "Restarts all nodes (with optional mysqld arguments)",
			Notes:       "",
			Contents:    restartMultiTemplate,
		},
		"use_multi_template": TemplateDesc{
			Description: "Runs the same SQL query in all nodes",
			Notes:       "",
			Contents:    useMultiTemplate,
		},
		"exec_multi_template": TemplateDesc{
			Description: "Runs the same command in all nodes",
			Notes:       "",
			Contents:    execMultiTemplate,
		},
		"metadata_multi_template": TemplateDesc{
			Description: "Runs a metadata query in all nodes",
			Notes:       "",
			Contents:    metadataMultiTemplate,
		},
		"use_multi_admin_template": TemplateDesc{
			Description: "Runs the same SQL query (as admin user) in all nodes",
			Notes:       "",
			Contents:    useMultiAdminTemplate,
		},
		"stop_multi_template": TemplateDesc{
			Description: "Stops all nodes",
			Notes:       "",
			Contents:    stopMultiTemplate,
		},
		"send_kill_multi_template": TemplateDesc{
			Description: "Sends kill signal to all nodes",
			Notes:       "",
			Contents:    sendKillMultiTemplate,
		},
		"clear_multi_template": TemplateDesc{
			Description: "Removes data from all nodes",
			Notes:       "",
			Contents:    clearMultiTemplate,
		},
		"status_multi_template": TemplateDesc{
			Description: "Shows status for all nodes",
			Notes:       "",
			Contents:    statusMultiTemplate,
		},
		"test_sb_multi_template": TemplateDesc{
			Description: "Run sb test on all nodes",
			Notes:       "",
			Contents:    testSbMultiTemplate,
		},
		"node_template": TemplateDesc{
			Description: "Runs the MySQL client for a given node",
			Notes:       "",
			Contents:    nodeTemplate,
		},
		"node_admin_template": TemplateDesc{
			Description: "Runs the MySQL client for a given node as admin user",
			Notes:       "",
			Contents:    nodeAdminTemplate,
		},
		"replicate_from_multi_template": TemplateDesc{
			Description: "calls script replicate_from from node #1",
			Notes:       "",
			Contents:    replicateFromMultiTemplate,
		},
		"sysbench_multi_template": TemplateDesc{
			Description: "calls script sysbench from node #1",
			Notes:       "",
			Contents:    sysbenchMultiTemplate,
		},
		"sysbench_ready_multi_template": TemplateDesc{
			Description: "calls script sysbench_ready from node #1",
			Notes:       "",
			Contents:    sysbenchReadyMultiTemplate,
		},
	}
)
