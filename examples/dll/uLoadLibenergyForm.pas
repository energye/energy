unit uForm;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, LCLType, StdCtrls, dynlibs;

type

  { TForm1 }

  TForm1 = class(TForm)
    Button1: TButton;
    procedure Button1Click(Sender: TObject);
    procedure FormCloseQuery(Sender: TObject; var CanClose: boolean);
    procedure FormCreate(Sender: TObject);
  private

  public
    procedure CEFCloseCallback();

  end;

var
  Form1: TForm1;

  DllHandle: HMODULE;
  initApplication: procedure(); stdcall;
  mainFormShow: procedure(); stdcall;
  mainFormClose: procedure(); stdcall;
  mainFormFree: procedure(); stdcall;
  initCEFApplication: procedure(); stdcall;
  cefFormShow: procedure(); stdcall;
  cefFormClose: procedure(); stdcall;
  cefFormFree: procedure(); stdcall;
  setWindowHwnd: procedure(ptr: HWND); stdcall;
  cefClose: procedure(); stdcall;

implementation

{$R *.lfm}

{ TForm1 }

procedure TForm1.FormCreate(Sender: TObject);
begin
  setWindowHwnd(self.Handle);
  initCEFApplication();
end;

procedure TForm1.Button1Click(Sender: TObject);
begin
  cefFormShow();
end;

procedure TForm1.CEFCloseCallback();
begin

end;

procedure TForm1.FormCloseQuery(Sender: TObject; var CanClose: boolean);
begin
  //WriteLn('FormCloseQuery');

  //cefClose();
  CanClose := True;
end;

procedure LoadLibEnergy();
begin
  DllHandle := LoadLibrary('libenergy.dll');
  if DllHandle = 0 then
  begin
    ShowMessage('Unable to load libenergy.dll');
    Exit;
  end;
  Pointer(initApplication) := GetProcAddress(DllHandle, 'initApplication');
  Pointer(mainFormShow) := GetProcAddress(DllHandle, 'mainFormShow');
  Pointer(mainFormClose) := GetProcAddress(DllHandle, 'mainFormClose');
  Pointer(mainFormFree) := GetProcAddress(DllHandle, 'mainFormFree');
  Pointer(initCEFApplication) := GetProcAddress(DllHandle, 'initCEFApplication');
  Pointer(cefFormShow) := GetProcAddress(DllHandle, 'cefFormShow');
  Pointer(cefFormClose) := GetProcAddress(DllHandle, 'cefFormClose');
  Pointer(cefFormFree) := GetProcAddress(DllHandle, 'cefFormFree');
  Pointer(setWindowHwnd) := GetProcAddress(DllHandle, 'setWindowHwnd');
  Pointer(cefClose) := GetProcAddress(DllHandle, 'cefClose');
end;

procedure FreeLibEnergy();
begin
  if DllHandle <> 0 then
  begin
    dynlibs.UnloadLibrary(DllHandle);
    DllHandle := 0;
  end;

end;

initialization
  begin
    //WriteLn('LoadLibEnergy');
    LoadLibEnergy;
  end;

finalization
  begin
    //WriteLn('FreeLibEnergy');
    FreeLibEnergy;
  end;

end.
