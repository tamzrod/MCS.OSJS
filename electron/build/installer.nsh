!include "nsDialogs.nsh"
!include "LogicLib.nsh"

!macro MCS_REMOVE_SERVICE service
  IfFileExists "$INSTDIR\resources\bin\nssm.exe" 0 +5
  nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" stop "${service}"'
  Pop $0
  nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" remove "${service}" confirm'
  Pop $0
!macroend

!ifndef BUILD_UNINSTALLER
Var MCSServicePage
Var MCSInstallMMA2
Var MCSInstallSimulator
Var MCSInstallReplicator
Var MCSStartServices
Var MCSInstallMMA2State
Var MCSInstallSimulatorState
Var MCSInstallReplicatorState
Var MCSStartServicesState

!macro customWelcomePage
  !insertmacro MUI_PAGE_WELCOME
!macroend

!macro customPageAfterChangeDir
  Page custom MCSServicePageCreate MCSServicePageLeave
!macroend

Function MCSServicePageCreate
  nsDialogs::Create 1018
  Pop $MCSServicePage
  ${If} $MCSServicePage == error
    Abort
  ${EndIf}

  ${NSD_CreateLabel} 0 0 100% 14u "Service Configuration"
  Pop $0
  CreateFont $1 "$(^Font)" "10" "700"
  SendMessage $0 ${WM_SETFONT} $1 1

  ${NSD_CreateLabel} 0 18u 100% 22u "Choose which MCS backend runtimes Windows will run as services. Electron remains a normal desktop application."
  Pop $0

  ${NSD_CreateCheckbox} 0 48u 100% 12u "MMA2 - shared Modbus memory appliance"
  Pop $MCSInstallMMA2
  ${NSD_Check} $MCSInstallMMA2

  ${NSD_CreateCheckbox} 0 68u 100% 12u "Simulator runtime"
  Pop $MCSInstallSimulator
  ${NSD_Check} $MCSInstallSimulator

  ${NSD_CreateCheckbox} 0 88u 100% 12u "Replicator runtime"
  Pop $MCSInstallReplicator
  ${NSD_Check} $MCSInstallReplicator

  ${NSD_CreateCheckbox} 0 116u 100% 12u "Start selected services after installation"
  Pop $MCSStartServices
  ${NSD_Check} $MCSStartServices

  ${NSD_CreateLabel} 0 142u 100% 30u "Service names: MCS-MMA2, MCS-Simulator, MCS-Replicator. Startup type: Automatic."
  Pop $0

  nsDialogs::Show
FunctionEnd

Function MCSServicePageLeave
  ${NSD_GetState} $MCSInstallMMA2 $MCSInstallMMA2State
  ${NSD_GetState} $MCSInstallSimulator $MCSInstallSimulatorState
  ${NSD_GetState} $MCSInstallReplicator $MCSInstallReplicatorState
  ${NSD_GetState} $MCSStartServices $MCSStartServicesState
FunctionEnd

!macro MCS_REQUIRE_SUCCESS action
  ${If} $0 != 0
    MessageBox MB_ICONSTOP|MB_OK "${action} failed with exit code $0. Installation cannot continue."
    Abort
  ${EndIf}
!macroend

!macro MCS_NSSM service executable
  nsExec::ExecToStack 'sc.exe query "${service}"'
  Pop $0
  Pop $1

  ${If} $0 == 0
    nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" stop "${service}"'
    Pop $0
    nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" set "${service}" Application "$INSTDIR\resources\bin\${executable}"'
    Pop $0
    !insertmacro MCS_REQUIRE_SUCCESS "Updating ${service} executable"
  ${Else}
    nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" install "${service}" "$INSTDIR\resources\bin\${executable}"'
    Pop $0
    !insertmacro MCS_REQUIRE_SUCCESS "Installing ${service} service"
  ${EndIf}

  ; All packaged runtimes are zero-argument service entry points. Clearing stale
  ; AppParameters also repairs installs created by older toolkit installers.
  nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" reset "${service}" AppParameters'
  Pop $0
  !insertmacro MCS_REQUIRE_SUCCESS "Resetting ${service} arguments"

  nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" set "${service}" AppDirectory "$LOCALAPPDATA\MCS Modbus Toolkit\runtime"'
  Pop $0
  !insertmacro MCS_REQUIRE_SUCCESS "Configuring ${service} AppDirectory"

  nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" set "${service}" AppEnvironmentExtra "OSJS_DATA_DIR=$LOCALAPPDATA\MCS Modbus Toolkit\runtime" "MCS_DATA_ROOT=$LOCALAPPDATA\MCS Modbus Toolkit\runtime"'
  Pop $0
  !insertmacro MCS_REQUIRE_SUCCESS "Configuring ${service} environment"

  nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" set "${service}" Start SERVICE_AUTO_START'
  Pop $0
  !insertmacro MCS_REQUIRE_SUCCESS "Configuring ${service} startup"

  nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" set "${service}" AppExit Default Restart'
  Pop $0
  !insertmacro MCS_REQUIRE_SUCCESS "Configuring ${service} restart policy"

  nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" set "${service}" AppRestartDelay 3000'
  Pop $0
  !insertmacro MCS_REQUIRE_SUCCESS "Configuring ${service} restart delay"
!macroend

!macro customInstall
  SetShellVarContext all
  CreateDirectory "$LOCALAPPDATA\MCS Modbus Toolkit\runtime"
  CreateDirectory "$LOCALAPPDATA\MCS Modbus Toolkit\runtime\config"
  CreateDirectory "$LOCALAPPDATA\MCS Modbus Toolkit\runtime\config\mma2"

  IfFileExists "$LOCALAPPDATA\MCS Modbus Toolkit\runtime\config\mma2\config.yaml" mma2_config_ready 0
  FileOpen $1 "$LOCALAPPDATA\MCS Modbus Toolkit\runtime\config\mma2\config.yaml" w
  FileWrite $1 "{}$\r$\n"
  FileClose $1
mma2_config_ready:

  IfFileExists "$INSTDIR\resources\bin\nssm.exe" +3 0
  MessageBox MB_ICONSTOP|MB_OK "NSSM was not packaged. Place nssm.exe in electron\bin and rebuild the installer."
  Abort

  ${If} $MCSInstallMMA2State == ${BST_CHECKED}
    IfFileExists "$INSTDIR\resources\bin\mma2-supervisor.exe" +3 0
    MessageBox MB_ICONSTOP|MB_OK "mma2-supervisor.exe was not packaged. Rebuild the installer."
    Abort

    !insertmacro MCS_NSSM "MCS-MMA2" "mma2-supervisor.exe"
  ${Else}
    !insertmacro MCS_REMOVE_SERVICE "MCS-MMA2"
  ${EndIf}

  ${If} $MCSInstallSimulatorState == ${BST_CHECKED}
    !insertmacro MCS_NSSM "MCS-Simulator" "modbus-simulator-runtime.exe"
    ${If} $MCSInstallMMA2State == ${BST_CHECKED}
      nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" set "MCS-Simulator" DependOnService "MCS-MMA2"'
      Pop $0
      !insertmacro MCS_REQUIRE_SUCCESS "Configuring MCS-Simulator dependency"
    ${Else}
      nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" reset "MCS-Simulator" DependOnService'
      Pop $0
    ${EndIf}
  ${Else}
    !insertmacro MCS_REMOVE_SERVICE "MCS-Simulator"
  ${EndIf}

  ${If} $MCSInstallReplicatorState == ${BST_CHECKED}
    !insertmacro MCS_NSSM "MCS-Replicator" "modbus-replicator-runtime.exe"
    ${If} $MCSInstallMMA2State == ${BST_CHECKED}
      nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" set "MCS-Replicator" DependOnService "MCS-MMA2"'
      Pop $0
      !insertmacro MCS_REQUIRE_SUCCESS "Configuring MCS-Replicator dependency"
    ${Else}
      nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" reset "MCS-Replicator" DependOnService'
      Pop $0
    ${EndIf}
  ${Else}
    !insertmacro MCS_REMOVE_SERVICE "MCS-Replicator"
  ${EndIf}

  ${If} $MCSStartServicesState == ${BST_CHECKED}
    ${If} $MCSInstallMMA2State == ${BST_CHECKED}
      nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" start "MCS-MMA2"'
      Pop $0
    ${EndIf}
    ${If} $MCSInstallSimulatorState == ${BST_CHECKED}
      nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" start "MCS-Simulator"'
      Pop $0
    ${EndIf}
    ${If} $MCSInstallReplicatorState == ${BST_CHECKED}
      nsExec::ExecToLog '"$INSTDIR\resources\bin\nssm.exe" start "MCS-Replicator"'
      Pop $0
    ${EndIf}
  ${EndIf}
!macroend
!endif

!macro customUnWelcomePage
  !insertmacro MUI_UNPAGE_WELCOME
!macroend

!macro customUnInstall
  !insertmacro MCS_REMOVE_SERVICE "MCS-Replicator"
  !insertmacro MCS_REMOVE_SERVICE "MCS-Simulator"
  !insertmacro MCS_REMOVE_SERVICE "MCS-MMA2"
!macroend
