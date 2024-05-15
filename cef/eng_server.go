//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICEFServer Parent: ICefBaseRefCounted
//
//	Interface representing a server that supports HTTP and WebSocket requests. Server capacity is limited and is intended to handle only a small number of simultaneous connections (e.g. for communicating between applications on localhost). The functions of this interface are safe to call from any thread in the brower process unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h (cef_server_t))</a>
type ICEFServer interface {
	ICefBaseRefCounted
	// GetTaskRunner
	//  Returns the task runner for the dedicated server thread.
	GetTaskRunner() ICefTaskRunner // function
	// IsRunning
	//  Returns true (1) if the server is currently running and accepting incoming connections. See ICefServerHandler.OnServerCreated documentation for a description of server lifespan. This function must be called on the dedicated server thread.
	IsRunning() bool // function
	// GetAddress
	//  Returns the server address including the port number.
	GetAddress() string // function
	// HasConnection
	//  Returns true (1) if the server currently has a connection. This function must be called on the dedicated server thread.
	HasConnection() bool // function
	// IsValidConnection
	//  Returns true (1) if |connection_id| represents a valid connection. This function must be called on the dedicated server thread.
	IsValidConnection(connectionid int32) bool // function
	// Shutdown
	//  Stop the server and shut down the dedicated server thread. See ICefServerHandler.OnServerCreated documentation for a description of server lifespan.
	Shutdown() // procedure
	// SendHttp200response
	//  Send an HTTP 200 "OK" response to the connection identified by |connection_id|. |content_type| is the response content type (e.g. "text/html"), |data| is the response content, and |data_size| is the size of |data| in bytes. The contents of |data| will be copied. The connection will be closed automatically after the response is sent.
	SendHttp200response(connectionid int32, contenttype string, data uintptr, datasize NativeUInt) // procedure
	// SendHttp404response
	//  Send an HTTP 404 "Not Found" response to the connection identified by |connection_id|. The connection will be closed automatically after the response is sent.
	SendHttp404response(connectionid int32) // procedure
	// SendHttp500response
	//  Send an HTTP 500 "Internal Server Error" response to the connection identified by |connection_id|. |error_message| is the associated error message. The connection will be closed automatically after the response is sent.
	SendHttp500response(connectionid int32, errormessage string) // procedure
	// SendHttpResponse
	//  Send a custom HTTP response to the connection identified by |connection_id|. |response_code| is the HTTP response code sent in the status line (e.g. 200), |content_type| is the response content type sent as the "Content-Type" header (e.g. "text/html"), |content_length| is the expected content length, and |extra_headers| is the map of extra response headers. If |content_length| is >= 0 then the "Content-Length" header will be sent. If |content_length| is 0 then no content is expected and the connection will be closed automatically after the response is sent. If |content_length| is < 0 then no "Content-Length" header will be sent and the client will continue reading until the connection is closed. Use the SendRawData function to send the content, if applicable, and call CloseConnection after all content has been sent.
	SendHttpResponse(connectionid, responsecode int32, contenttype string, contentlength int64, extraheaders ICefStringMultimap) // procedure
	// SendRawData
	//  Send raw data directly to the connection identified by |connection_id|. |data| is the raw data and |data_size| is the size of |data| in bytes. The contents of |data| will be copied. No validation of |data| is performed internally so the client should be careful to send the amount indicated by the "Content-Length" header, if specified. See SendHttpResponse documentation for intended usage.
	SendRawData(connectionid int32, data uintptr, datasize NativeUInt) // procedure
	// CloseConnection
	//  Close the connection identified by |connection_id|. See SendHttpResponse documentation for intended usage.
	CloseConnection(connectionid int32) // procedure
	// SendWebSocketMessage
	//  Send a WebSocket message to the connection identified by |connection_id|. |data| is the response content and |data_size| is the size of |data| in bytes. The contents of |data| will be copied. See ICefServerHandler.OnWebSocketRequest documentation for intended usage.
	SendWebSocketMessage(connectionid int32, data uintptr, datasize NativeUInt) // procedure
}

// TCEFServer Parent: TCefBaseRefCounted
//
//	Interface representing a server that supports HTTP and WebSocket requests. Server capacity is limited and is intended to handle only a small number of simultaneous connections (e.g. for communicating between applications on localhost). The functions of this interface are safe to call from any thread in the brower process unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h (cef_server_t))</a>
type TCEFServer struct {
	TCefBaseRefCounted
}

// ServerRef -> ICEFServer
var ServerRef server

// server TCEFServer Ref
type server uintptr

func (m *server) UnWrap(data uintptr) ICEFServer {
	var resultCEFServer uintptr
	CEF().SysCallN(238, uintptr(data), uintptr(unsafePointer(&resultCEFServer)))
	return AsCEFServer(resultCEFServer)
}

func (m *TCEFServer) GetTaskRunner() ICefTaskRunner {
	var resultCefTaskRunner uintptr
	CEF().SysCallN(227, m.Instance(), uintptr(unsafePointer(&resultCefTaskRunner)))
	return AsCefTaskRunner(resultCefTaskRunner)
}

func (m *TCEFServer) IsRunning() bool {
	r1 := CEF().SysCallN(229, m.Instance())
	return GoBool(r1)
}

func (m *TCEFServer) GetAddress() string {
	r1 := CEF().SysCallN(226, m.Instance())
	return GoStr(r1)
}

func (m *TCEFServer) HasConnection() bool {
	r1 := CEF().SysCallN(228, m.Instance())
	return GoBool(r1)
}

func (m *TCEFServer) IsValidConnection(connectionid int32) bool {
	r1 := CEF().SysCallN(230, m.Instance(), uintptr(connectionid))
	return GoBool(r1)
}

func (m *TCEFServer) Shutdown() {
	CEF().SysCallN(237, m.Instance())
}

func (m *TCEFServer) SendHttp200response(connectionid int32, contenttype string, data uintptr, datasize NativeUInt) {
	CEF().SysCallN(231, m.Instance(), uintptr(connectionid), PascalStr(contenttype), uintptr(data), uintptr(datasize))
}

func (m *TCEFServer) SendHttp404response(connectionid int32) {
	CEF().SysCallN(232, m.Instance(), uintptr(connectionid))
}

func (m *TCEFServer) SendHttp500response(connectionid int32, errormessage string) {
	CEF().SysCallN(233, m.Instance(), uintptr(connectionid), PascalStr(errormessage))
}

func (m *TCEFServer) SendHttpResponse(connectionid, responsecode int32, contenttype string, contentlength int64, extraheaders ICefStringMultimap) {
	CEF().SysCallN(234, m.Instance(), uintptr(connectionid), uintptr(responsecode), PascalStr(contenttype), uintptr(unsafePointer(&contentlength)), GetObjectUintptr(extraheaders))
}

func (m *TCEFServer) SendRawData(connectionid int32, data uintptr, datasize NativeUInt) {
	CEF().SysCallN(235, m.Instance(), uintptr(connectionid), uintptr(data), uintptr(datasize))
}

func (m *TCEFServer) CloseConnection(connectionid int32) {
	CEF().SysCallN(225, m.Instance(), uintptr(connectionid))
}

func (m *TCEFServer) SendWebSocketMessage(connectionid int32, data uintptr, datasize NativeUInt) {
	CEF().SysCallN(236, m.Instance(), uintptr(connectionid), uintptr(data), uintptr(datasize))
}
