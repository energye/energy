//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package notice

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"golang.org/x/sys/execabs"
)

const notificationTemplate = `$title = "%s"
$content = "%s"

[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] > $null
$template = [Windows.UI.Notifications.ToastNotificationManager]::GetTemplateContent([Windows.UI.Notifications.ToastTemplateType]::ToastText02)
$toastXml = [xml] $template.GetXml()
$toastXml.GetElementsByTagName("text")[0].AppendChild($toastXml.CreateTextNode($title)) > $null
$toastXml.GetElementsByTagName("text")[1].AppendChild($toastXml.CreateTextNode($content)) > $null

$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
$xml.LoadXml($toastXml.OuterXml)
$toast = [Windows.UI.Notifications.ToastNotification]::new($xml)
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier("%s").Show($toast);`

var scriptNum = 0

func SendNotification(n *Notification) {
	title := escapeNotificationString(n.Title)
	content := escapeNotificationString(n.Content)
	appID := UniqueID()
	script := fmt.Sprintf(notificationTemplate, title, content, appID)
	go runScript("notify", script)
}

func escapeNotificationString(in string) string {
	noSlash := strings.Replace(in, "`", "``", -1)
	return strings.Replace(noSlash, "\"", "`\"", -1)
}

func runScript(name, script string) {
	scriptNum++
	appID := UniqueID()
	fileName := fmt.Sprintf("notice-%s-%s-%d.ps1", appID, name, scriptNum)

	tmpFilePath := filepath.Join(os.TempDir(), fileName)
	err := ioutil.WriteFile(tmpFilePath, []byte(script), 0600)
	if err != nil {
		return
	}
	defer os.Remove(tmpFilePath)

	launch := "(Get-Content -Encoding UTF8 -Path " + tmpFilePath + " -Raw) | Invoke-Expression"
	cmd := execabs.Command("PowerShell", "-ExecutionPolicy", "Bypass", launch)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = cmd.Run()
}
