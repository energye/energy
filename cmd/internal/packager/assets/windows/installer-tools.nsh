; DO NOT EDIT - Generated automatically by `energy build`

!include "x64.nsh"
!include "WinVer.nsh"
!include "FileFunc.nsh"

!ifndef INFO_PROJECTNAME
    !define INFO_PROJECTNAME "{{.Name}}"
!endif
!ifndef INFO_COMPANYNAME
    !define INFO_COMPANYNAME "{{.Info.CompanyName}}"
!endif
!ifndef INFO_PRODUCTNAME
    !define INFO_PRODUCTNAME "{{.Info.ProductName}}"
!endif
!ifndef INFO_PRODUCTVERSION
    !define INFO_PRODUCTVERSION "{{.Info.ProductVersion}}"
!endif
!ifndef INFO_COPYRIGHT
    !define INFO_COPYRIGHT "{{.Info.Copyright}}"
!endif
!ifndef PRODUCT_EXECUTABLE
    !define PRODUCT_EXECUTABLE "${INFO_PROJECTNAME}.exe"
!endif
!ifndef UNINST_KEY_NAME
    !define UNINST_KEY_NAME "${INFO_COMPANYNAME}${INFO_PRODUCTNAME}"
!endif
!define UNINST_KEY "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}"

!ifndef REQUEST_EXECUTION_LEVEL
    !define REQUEST_EXECUTION_LEVEL "admin"
!endif

RequestExecutionLevel "${REQUEST_EXECUTION_LEVEL}"

!ifdef ARG_ENERGY_LANGUAGE
    !define ENERGY_LANGUAGE "${ARG_ENERGY_LANGUAGE}" # customer
!else
    !define ENERGY_LANGUAGE "English" # default
!endif

!macro energy.files

    File "/oname=${PRODUCT_EXECUTABLE}" "${ARG_ENERGY_BINARY}" ; energy app.exe path, ..\..\app.exe
    ;File /r "${ARG_ENERGY_CEF_FRAMEWORK}" ; cef framework path, ENERGY_HOME=/to/cef/path

!macroend

!macro energy.writeUninstaller
    WriteUninstaller "$INSTDIR\uninstall.exe"

    SetRegView 64
    WriteRegStr HKLM "${UNINST_KEY}" "Publisher" "${INFO_COMPANYNAME}"
    WriteRegStr HKLM "${UNINST_KEY}" "DisplayName" "${INFO_PRODUCTNAME}"
    WriteRegStr HKLM "${UNINST_KEY}" "DisplayVersion" "${INFO_PRODUCTVERSION}"
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
