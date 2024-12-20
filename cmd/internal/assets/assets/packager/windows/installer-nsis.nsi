Unicode true
!include "installer-tools.nsh"

; The version information for this two must consist of 4 parts
VIProductVersion "${INFO_ProductVersion}.0"
VIFileVersion    "${INFO_FileVersion}.0"

VIAddVersionKey "ProductName"     "${INFO_ProductName}"
VIAddVersionKey "CompanyName"     "${INFO_CompanyName}"
VIAddVersionKey "FileDescription" "${INFO_FileDescription}"
VIAddVersionKey "ProductVersion"  "${INFO_ProductVersion}"
VIAddVersionKey "FileVersion"     "${INFO_FileVersion}"
VIAddVersionKey "LegalCopyright"  "${INFO_Copyright}"

!include "MUI2.nsh"

!define MUI_ICON "${INFO_Icon}" ;"..\icon.ico"
!define MUI_UNICON "${INFO_UnIcon}" ;"..\icon.ico"

; !define MUI_WELCOMEFINISHPAGE_BITMAP "resources\leftimage.bmp" #Include this to add a bitmap on the left side of the Welcome Page. Must be a size of 164x314
!define MUI_FINISHPAGE_NOAUTOCLOSE # Wait on the INSTFILES page so the user can take a look into the details of the installation steps
!define MUI_ABORTWARNING # This will warn the user if they exit from the installer.

; !define MUI_WELCOMEPAGE_TITLE "Title" #
; !define MUI_WELCOMEPAGE_TEXT  "Text" #

!insertmacro MUI_PAGE_WELCOME # Welcome to the installer page.

; LICENSE Page
!ifdef ENERGY_PAGE_LICENSE
    !insertmacro MUI_PAGE_LICENSE "${ENERGY_PAGE_LICENSE}" # Add a LICENSE page to the installer
!endif

!insertmacro MUI_PAGE_DIRECTORY # In which folder install page.
!insertmacro MUI_PAGE_INSTFILES # Installing page.
!insertmacro MUI_PAGE_FINISH # Finished installation page.

!insertmacro MUI_UNPAGE_INSTFILES # Uinstalling page

!insertmacro MUI_LANGUAGE "${ENERGY_LANGUAGE}" # Set the Language of the installer

; The following two statements can be used to sign the installer and the uninstaller. The path to the binaries are provided in %1
;!uninstfinalize 'signtool --file "%1"'
;!finalize 'signtool --file "%1"'

Name "${INFO_ProductName}"
OutFile ".\${INFO_InstallerFileName}" # Name of the installer's file.
InstallDir "$PROGRAMFILES64\${INFO_CompanyName}\${INFO_ProductName}" # Default installing folder ($PROGRAMFILES is Program Files folder).
ShowInstDetails show # This will always show the installation details.

Function .onInit
FunctionEnd

Section
    !insertmacro energy.setShellContext

    SetOutPath $INSTDIR
    
    !insertmacro energy.files

    CreateShortcut "$SMPROGRAMS\${INFO_ShortCutName}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"
    CreateShortCut "$DESKTOP\${INFO_ShortCutName}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"

    !insertmacro energy.compressNsis7z

    !insertmacro energy.writeUninstaller
SectionEnd

Section "uninstall" 
    !insertmacro energy.setShellContext

    RMDir /r "$AppData\${PRODUCT_EXECUTABLE}"

    RMDir /r $INSTDIR

    Delete "$SMPROGRAMS\${INFO_ShortCutName}.lnk"
    Delete "$DESKTOP\${INFO_ShortCutName}.lnk"

    !insertmacro energy.deleteUninstaller
SectionEnd
