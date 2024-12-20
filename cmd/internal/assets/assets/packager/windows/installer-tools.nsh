; DO NOT EDIT - Generated automatically by `energy build`

!include "x64.nsh"
!include "WinVer.nsh"
!include "FileFunc.nsh"

!define INFO_ProjectName "{{.Name}}"
!define INFO_InstallerFileName "{{.InstallerFileName}}"
!define INFO_CompanyName "{{.Info.CompanyName}}"
!define INFO_ProductName "{{.Info.ProductName}}"
!define INFO_ShortCutName "{{.NSIS.ShortCutName}}"
!define INFO_FileVersion "{{.Info.FileVersion}}"
!define INFO_ProductVersion "{{.Info.ProductVersion}}"
!define INFO_FileDescription "{{.Info.FileDescription}}"
!define INFO_Copyright "{{.Info.Copyright}}"
!define PRODUCT_EXECUTABLE "${INFO_ProjectName}.exe"
!define UNINST_KEY_NAME "${INFO_CompanyName}${INFO_ProductName}"
!define INFO_Icon "{{.NSIS.Icon}}"
!define INFO_UnIcon "{{.NSIS.UnIcon}}"
!define ENERGY_LANGUAGE "{{.NSIS.Language}}"

!define UNINST_KEY "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}"

{{if .NSIS.License}}
!define ENERGY_PAGE_LICENSE "{{.NSIS.License}}" ; license.txt path
{{end}}

{{if .NSIS.RequestExecutionLevel}}
!define REQUEST_EXECUTION_LEVEL "{{.NSIS.RequestExecutionLevel}}"
RequestExecutionLevel "${REQUEST_EXECUTION_LEVEL}" ; admin or ""
{{end}}


!macro energy.files

    File "/oname=${PRODUCT_EXECUTABLE}" "{{.ProjectPath}}\{{.ExeName}}.exe" ; app.exe path, ..\..\app.exe

{{if .NSIS.CompressFile}}
    File "{{.NSIS.CompressFile}}"
{{else if .FrameworkPath}}
    File /r "{{.FrameworkPath}}\*.*"
{{end}}

{{range $i,$path := .NSIS.Include }}
    File /r "{{$path}}"{{end}}
!macroend

!macro energy.compressNsis7z
{{if .NSIS.UseCompress}} ;CEF
    Nsis7z::ExtractWithCallback "$INSTDIR\{{.NSIS.CompressName}}"
    Delete "$OUTDIR\{{.NSIS.CompressName}}"
{{end}}
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
