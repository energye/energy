; DO NOT EDIT - Generated automatically by `energy build`

!include "x64.nsh"
!include "WinVer.nsh"
!include "FileFunc.nsh"

!define INFO_ProjectName "{{.Name}}"
!define INFO_CompanyName "{{.Info.CompanyName}}"
!define INFO_ProductName "{{.Info.ProductName}}"
!define INFO_FileVersion "{{.Info.FileVersion}}"
!define INFO_ProductVersion "{{.Info.ProductVersion}}"
!define INFO_FileDescription "{{.Info.FileDescription}}"
!define INFO_Copyright "{{.Info.Copyright}}"
!define PRODUCT_EXECUTABLE "${INFO_ProjectName}.exe"
!define UNINST_KEY_NAME "${INFO_CompanyName}${INFO_ProductName}"
!define INFO_Icon "{{.Info.InstallPack.Icon}}"
!define INFO_UnIcon "{{.Info.InstallPack.UnIcon}}"
!define ENERGY_LANGUAGE "{{.Info.InstallPack.Language}}"

!define UNINST_KEY "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}"

{{if .Info.InstallPack.License}}
!define ENERGY_PAGE_LICENSE "{{.Info.InstallPack.License}}" ; license.txt path
{{end}}

{{if .Info.InstallPack.RequestExecutionLevel}}
!define REQUEST_EXECUTION_LEVEL "{{.Info.InstallPack.RequestExecutionLevel}}"
RequestExecutionLevel "${REQUEST_EXECUTION_LEVEL}" ; admin or ""
{{end}}


!macro energy.files

    File "/oname=${PRODUCT_EXECUTABLE}" "{{.ProjectPath}}\{{.Name}}.exe" ; app.exe path, ..\..\app.exe
    ;File /r "{{.FrameworkPath}}" ; cef framework path, ENERGY_HOME=/to/cef/path

!macroend

!macro energy.writeUninstaller
    WriteUninstaller "$INSTDIR\uninstall.exe"

    SetRegView 64
    WriteRegStr HKLM "${UNINST_KEY}" "Publisher" "${INFO_CompanyName}"
    WriteRegStr HKLM "${UNINST_KEY}" "DisplayName" "${INFO_ProductName}"
    WriteRegStr HKLM "${UNINST_KEY}" "DisplayVersion" "${INFO_ProductVersion}"
    WriteRegStr HKLM "${UNINST_KEY}" "DisplayIcon" "$INSTDIR\${PRODUCT_EXECUTABLE}"
    WriteRegStr HKLM "${UNINST_KEY}" "UninstallString" "$\"$INSTDIR\uninstall.exe$\""
    WriteRegStr HKLM "${UNINST_KEY}" "QuietUninstallString" "$\"$INSTDIR\uninstall.exe$\" /S"

    ${GetSize} "$INSTDIR" "/S=0K" $0 $1 $2
    IntFmt $0 "0x%08X" $0
    WriteRegDWORD HKLM "${UNINST_KEY}" "EstimatedSize" "$0"
!macroend

!macro energy.deleteUninstaller
    Delete "$INSTDIR\uninstall.exe"

    SetRegView 64
    DeleteRegKey HKLM "${UNINST_KEY}"
!macroend

!macro energy.setShellContext
    ${If} ${REQUEST_EXECUTION_LEVEL} == "admin"
        SetShellVarContext all
    ${else}
        SetShellVarContext current
    ${EndIf}
!macroend
