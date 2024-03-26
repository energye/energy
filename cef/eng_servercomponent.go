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

// ICEFServerComponent Parent: IComponent
//
//	The TCEFServerComponent class puts together all CEF server procedures, functions, properties and events in one place.
type ICEFServerComponent interface {
	IComponent
	// Initialized
	//  Returns true when the server and the handler are initialized.
	Initialized() bool // property
	// IsRunning
	//  Returns true(1) if the server is currently running and accepting incoming
	//  connections. See ICefServerHandler.OnServerCreated documentation for a
	//  description of server lifespan. This function must be called on the
	//  dedicated server thread.
	IsRunning() bool // property
	// Address
	//  Returns the server address including the port number.
	Address() string // property
	// HasConnection
	//  Returns true(1) if the server currently has a connection. This function
	//  must be called on the dedicated server thread.
	HasConnection() bool // property
	// IsValidConnection
	//  Returns true(1) if |connection_id| represents a valid connection. This
	//  function must be called on the dedicated server thread.
	IsValidConnection(connectionid int32) bool // function
	// CreateServer
	//  Create a new server that binds to |address| and |port|. |address| must be a
	//  valid IPv4 or IPv6 address(e.g. 127.0.0.1 or::1) and |port| must be a port
	//  number outside of the reserved range(e.g. between 1025 and 65535 on most
	//  platforms). |backlog| is the maximum number of pending connections. A new
	//  thread will be created for each CreateServer call(the "dedicated server
	//  thread"). It is therefore recommended to use a different
	//  ICefServerHandler instance for each CreateServer call to avoid thread
	//  safety issues in the ICefServerHandler implementation. The
	//  ICefServerHandler.OnServerCreated function will be called on the
	//  dedicated server thread to report success or failure. See
	//  ICefServerHandler.OnServerCreated documentation for a description of
	//  server lifespan.
	CreateServer(address string, port uint16, backlog int32) // procedure
	// Shutdown
	//  Stop the server and shut down the dedicated server thread. See
	//  ICefServerHandler.OnServerCreated documentation for a description of
	//  server lifespan.
	Shutdown() // procedure
	// SendHttp200response
	//  Send an HTTP 200 "OK" response to the connection identified by
	//  |connection_id|. |content_type| is the response content type(e.g.
	//  "text/html"), |data| is the response content, and |data_size| is the size
	//  of |data| in bytes. The contents of |data| will be copied. The connection
	//  will be closed automatically after the response is sent.
	SendHttp200response(connectionid int32, contenttype string, data uintptr, datasize NativeUInt) // procedure
	// SendHttp404response
	//  Send an HTTP 404 "Not Found" response to the connection identified by
	//  |connection_id|. The connection will be closed automatically after the
	//  response is sent.
	SendHttp404response(connectionid int32) // procedure
	// SendHttp500response
	//  Send an HTTP 500 "Internal Server Error" response to the connection
	//  identified by |connection_id|. |error_message| is the associated error
	//  message. The connection will be closed automatically after the response is
	//  sent.
	SendHttp500response(connectionid int32, errormessage string) // procedure
	// SendHttpResponse
	//  Send a custom HTTP response to the connection identified by
	//  |connection_id|. |response_code| is the HTTP response code sent in the
	//  status line(e.g. 200), |content_type| is the response content type sent
	//  as the "Content-Type" header(e.g. "text/html"), |content_length| is the
	//  expected content length, and |extra_headers| is the map of extra response
	//  headers. If |content_length| is >= 0 then the "Content-Length" header will
	//  be sent. If |content_length| is 0 then no content is expected and the
	//  connection will be closed automatically after the response is sent. If
	//  |content_length| is < 0 then no "Content-Length" header will be sent and
	//  the client will continue reading until the connection is closed. Use the
	//  SendRawData function to send the content, if applicable, and call
	//  CloseConnection after all content has been sent.
	SendHttpResponse(connectionid, responsecode int32, contenttype string, contentlength int64, extraheaders ICefStringMultimap) // procedure
	// SendRawData
	//  Send raw data directly to the connection identified by |connection_id|.
	//  |data| is the raw data and |data_size| is the size of |data| in bytes. The
	//  contents of |data| will be copied. No validation of |data| is performed
	//  internally so the client should be careful to send the amount indicated by
	//  the "Content-Length" header, if specified. See SendHttpResponse
	//  documentation for intended usage.
	SendRawData(connectionid int32, data uintptr, datasize NativeUInt) // procedure
	// CloseConnection
	//  Close the connection identified by |connection_id|. See SendHttpResponse
	//  documentation for intended usage.
	CloseConnection(connectionid int32) // procedure
	// SendWebSocketMessage
	//  Send a WebSocket message to the connection identified by |connection_id|.
	//  |data| is the response content and |data_size| is the size of |data| in
	//  bytes. The contents of |data| will be copied. See
	//  ICefServerHandler.OnWebSocketRequest documentation for intended usage.
	SendWebSocketMessage(connectionid int32, data uintptr, datasize NativeUInt) // procedure
	// SetOnServerCreated
	//  Called when |server| is created. If the server was started successfully
	//  then ICefServer.IsRunning will return true(1). The server will
	//  continue running until ICefServerShutdown is called, after which time
	//  OnServerDestroyed will be called. If the server failed to start then
	//  OnServerDestroyed will be called immediately after this function returns.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnServerCreated(fn TOnServerCreated) // property event
	// SetOnServerDestroyed
	//  Called when |server| is destroyed. The server thread will be stopped after
	//  this function returns. The client should release any references to
	//  |server| when this function is called. See OnServerCreated documentation
	//  for a description of server lifespan.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnServerDestroyed(fn TOnServerDestroyed) // property event
	// SetOnClientConnected
	//  Called when a client connects to |server|. |connection_id| uniquely
	//  identifies the connection. Each call to this function will have a matching
	//  call to OnClientDisconnected.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnClientConnected(fn TOnClientConnected) // property event
	// SetOnClientDisconnected
	//  Called when a client disconnects from |server|. |connection_id| uniquely
	//  identifies the connection. The client should release any data associated
	//  with |connection_id| when this function is called and |connection_id|
	//  should no longer be passed to ICefServer functions. Disconnects can
	//  originate from either the client or the server. For example, the server
	//  will disconnect automatically after a ICefServer.SendHttpXXXResponse
	//  function is called.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnClientDisconnected(fn TOnClientDisconnected) // property event
	// SetOnHttpRequest
	//  Called when |server| receives an HTTP request. |connection_id| uniquely
	//  identifies the connection, |client_address| is the requesting IPv4 or IPv6
	//  client address including port number, and |request| contains the request
	//  contents(URL, function, headers and optional POST data). Call
	//  ICefServer functions either synchronously or asynchronusly to send a
	//  response.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnHttpRequest(fn TOnHttpRequest) // property event
	// SetOnWebSocketRequest
	//  Called when |server| receives a WebSocket request. |connection_id|
	//  uniquely identifies the connection, |client_address| is the requesting
	//  IPv4 or IPv6 client address including port number, and |request| contains
	//  the request contents(URL, function, headers and optional POST data).
	//  Execute |callback| either synchronously or asynchronously to accept or
	//  decline the WebSocket connection. If the request is accepted then
	//  OnWebSocketConnected will be called after the WebSocket has connected and
	//  incoming messages will be delivered to the OnWebSocketMessage callback. If
	//  the request is declined then the client will be disconnected and
	//  OnClientDisconnected will be called. Call the
	//  ICefServer.SendWebSocketMessage function after receiving the
	//  OnWebSocketConnected callback to respond with WebSocket messages.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnWebSocketRequest(fn TOnWebSocketRequest) // property event
	// SetOnWebSocketConnected
	//  Called after the client has accepted the WebSocket connection for |server|
	//  and |connection_id| via the OnWebSocketRequest callback. See
	//  OnWebSocketRequest documentation for intended usage.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnWebSocketConnected(fn TOnWebSocketConnected) // property event
	// SetOnWebSocketMessage
	//  Called when |server| receives an WebSocket message. |connection_id|
	//  uniquely identifies the connection, |data| is the message content and
	//  |data_size| is the size of |data| in bytes. Do not keep a reference to
	//  |data| outside of this function. See OnWebSocketRequest documentation for
	//  intended usage.
	//  This event will be called on the CEF server thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h(cef_server_handler_t)</a>
	SetOnWebSocketMessage(fn TOnWebSocketMessage) // property event
}

// TCEFServerComponent Parent: TComponent
//
//	The TCEFServerComponent class puts together all CEF server procedures, functions, properties and events in one place.
type TCEFServerComponent struct {
	TComponent
	serverCreatedPtr      uintptr
	serverDestroyedPtr    uintptr
	clientConnectedPtr    uintptr
	clientDisconnectedPtr uintptr
	httpRequestPtr        uintptr
	webSocketRequestPtr   uintptr
	webSocketConnectedPtr uintptr
	webSocketMessagePtr   uintptr
}

func NewCEFServerComponent(aOwner IComponent) ICEFServerComponent {
	r1 := CEF().SysCallN(204, GetObjectUintptr(aOwner))
	return AsCEFServerComponent(r1)
}

func (m *TCEFServerComponent) Initialized() bool {
	r1 := CEF().SysCallN(207, m.Instance())
	return GoBool(r1)
}

func (m *TCEFServerComponent) IsRunning() bool {
	r1 := CEF().SysCallN(208, m.Instance())
	return GoBool(r1)
}

func (m *TCEFServerComponent) Address() string {
	r1 := CEF().SysCallN(201, m.Instance())
	return GoStr(r1)
}

func (m *TCEFServerComponent) HasConnection() bool {
	r1 := CEF().SysCallN(206, m.Instance())
	return GoBool(r1)
}

func (m *TCEFServerComponent) IsValidConnection(connectionid int32) bool {
	r1 := CEF().SysCallN(209, m.Instance(), uintptr(connectionid))
	return GoBool(r1)
}

func CEFServerComponentClass() TClass {
	ret := CEF().SysCallN(202)
	return TClass(ret)
}

func (m *TCEFServerComponent) CreateServer(address string, port uint16, backlog int32) {
	CEF().SysCallN(205, m.Instance(), PascalStr(address), uintptr(port), uintptr(backlog))
}

func (m *TCEFServerComponent) Shutdown() {
	CEF().SysCallN(224, m.Instance())
}

func (m *TCEFServerComponent) SendHttp200response(connectionid int32, contenttype string, data uintptr, datasize NativeUInt) {
	CEF().SysCallN(210, m.Instance(), uintptr(connectionid), PascalStr(contenttype), uintptr(data), uintptr(datasize))
}

func (m *TCEFServerComponent) SendHttp404response(connectionid int32) {
	CEF().SysCallN(211, m.Instance(), uintptr(connectionid))
}

func (m *TCEFServerComponent) SendHttp500response(connectionid int32, errormessage string) {
	CEF().SysCallN(212, m.Instance(), uintptr(connectionid), PascalStr(errormessage))
}

func (m *TCEFServerComponent) SendHttpResponse(connectionid, responsecode int32, contenttype string, contentlength int64, extraheaders ICefStringMultimap) {
	CEF().SysCallN(213, m.Instance(), uintptr(connectionid), uintptr(responsecode), PascalStr(contenttype), uintptr(unsafePointer(&contentlength)), GetObjectUintptr(extraheaders))
}

func (m *TCEFServerComponent) SendRawData(connectionid int32, data uintptr, datasize NativeUInt) {
	CEF().SysCallN(214, m.Instance(), uintptr(connectionid), uintptr(data), uintptr(datasize))
}

func (m *TCEFServerComponent) CloseConnection(connectionid int32) {
	CEF().SysCallN(203, m.Instance(), uintptr(connectionid))
}

func (m *TCEFServerComponent) SendWebSocketMessage(connectionid int32, data uintptr, datasize NativeUInt) {
	CEF().SysCallN(215, m.Instance(), uintptr(connectionid), uintptr(data), uintptr(datasize))
}

func (m *TCEFServerComponent) SetOnServerCreated(fn TOnServerCreated) {
	if m.serverCreatedPtr != 0 {
		RemoveEventElement(m.serverCreatedPtr)
	}
	m.serverCreatedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(219, m.Instance(), m.serverCreatedPtr)
}

func (m *TCEFServerComponent) SetOnServerDestroyed(fn TOnServerDestroyed) {
	if m.serverDestroyedPtr != 0 {
		RemoveEventElement(m.serverDestroyedPtr)
	}
	m.serverDestroyedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(220, m.Instance(), m.serverDestroyedPtr)
}

func (m *TCEFServerComponent) SetOnClientConnected(fn TOnClientConnected) {
	if m.clientConnectedPtr != 0 {
		RemoveEventElement(m.clientConnectedPtr)
	}
	m.clientConnectedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(216, m.Instance(), m.clientConnectedPtr)
}

func (m *TCEFServerComponent) SetOnClientDisconnected(fn TOnClientDisconnected) {
	if m.clientDisconnectedPtr != 0 {
		RemoveEventElement(m.clientDisconnectedPtr)
	}
	m.clientDisconnectedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(217, m.Instance(), m.clientDisconnectedPtr)
}

func (m *TCEFServerComponent) SetOnHttpRequest(fn TOnHttpRequest) {
	if m.httpRequestPtr != 0 {
		RemoveEventElement(m.httpRequestPtr)
	}
	m.httpRequestPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(218, m.Instance(), m.httpRequestPtr)
}

func (m *TCEFServerComponent) SetOnWebSocketRequest(fn TOnWebSocketRequest) {
	if m.webSocketRequestPtr != 0 {
		RemoveEventElement(m.webSocketRequestPtr)
	}
	m.webSocketRequestPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(223, m.Instance(), m.webSocketRequestPtr)
}

func (m *TCEFServerComponent) SetOnWebSocketConnected(fn TOnWebSocketConnected) {
	if m.webSocketConnectedPtr != 0 {
		RemoveEventElement(m.webSocketConnectedPtr)
	}
	m.webSocketConnectedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(221, m.Instance(), m.webSocketConnectedPtr)
}

func (m *TCEFServerComponent) SetOnWebSocketMessage(fn TOnWebSocketMessage) {
	if m.webSocketMessagePtr != 0 {
		RemoveEventElement(m.webSocketMessagePtr)
	}
	m.webSocketMessagePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(222, m.Instance(), m.webSocketMessagePtr)
}
