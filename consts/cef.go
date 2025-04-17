package consts

import "github.com/cyber-xxm/energy/v2/types"

type CEF_NET_ERROR = types.Int32

const (
	/// Supported error code values.
	/// Ranges:
	///     0- 99 System related errors
	///   100-199 Connection related errors
	///   200-299 Certificate errors
	///   300-399 HTTP errors
	///   400-499 Cache errors
	///   500-599 ?
	///   600-699 FTP errors
	///   700-799 Certificate manager errors
	///   800-899 DNS resolver errors
	/// <para>TCefErrorCode values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see></para>
	/// <para><see href="https://chromium.googlesource.com/chromium/src/+/master/net/base/net_error_list.h">For the complete list of error values see include/base/internal/cef_net_error_list.h which includes this Chromium source file /net/base/net_error_list.h</see></para>
	ERR_NONE                               CEF_NET_ERROR = 0
	ERR_IO_PENDING                                       = -1
	ERR_FAILED                                           = -2
	ERR_ABORTED                                          = -3
	ERR_INVALID_ARGUMENT                                 = -4
	ERR_INVALID_HANDLE                                   = -5
	ERR_FILE_NOT_FOUND                                   = -6
	ERR_TIMED_OUT                                        = -7
	ERR_FILE_TOO_BIG                                     = -8
	ERR_UNEXPECTED                                       = -9
	ERR_ACCESS_DENIED                                    = -10
	ERR_NOT_IMPLEMENTED                                  = -11
	ERR_INSUFFICIENT_RESOURCES                           = -12
	ERR_OUT_OF_MEMORY                                    = -13
	ERR_UPLOAD_FILE_CHANGED                              = -14
	ERR_SOCKET_NOT_CONNECTED                             = -15
	ERR_FILE_EXISTS                                      = -16
	ERR_FILE_PATH_TOO_LONG                               = -17
	ERR_FILE_NO_SPACE                                    = -18
	ERR_FILE_VIRUS_INFECTED                              = -19
	ERR_BLOCKED_BY_CLIENT                                = -20
	ERR_NETWORK_CHANGED                                  = -21
	ERR_BLOCKED_BY_ADMINISTRATOR                         = -22
	ERR_SOCKET_IS_CONNECTED                              = -23
	ERR_BLOCKED_ENROLLMENT_CHECK_PENDING                 = -24
	ERR_UPLOAD_STREAM_REWIND_NOT_SUPPORTED               = -25
	ERR_CONTEXT_SHUT_DOWN                                = -26
	ERR_BLOCKED_BY_RESPONSE                              = -27
	ERR_BLOCKED_BY_XSS_AUDITOR                           = -28
	ERR_CLEARTEXT_NOT_PERMITTED                          = -29
	ERR_BLOCKED_BY_CSP                                   = -30
	ERR_H2_OR_QUIC_REQUIRED                              = -31
	//
	ERR_CONNECTION_CLOSED                 = -100
	ERR_CONNECTION_RESET                  = -101
	ERR_CONNECTION_REFUSED                = -102
	ERR_CONNECTION_ABORTED                = -103
	ERR_CONNECTION_FAILED                 = -104
	ERR_NAME_NOT_RESOLVED                 = -105
	ERR_INTERNET_DISCONNECTED             = -106
	ERR_SSL_PROTOCOL_ERROR                = -107
	ERR_ADDRESS_INVALID                   = -108
	ERR_ADDRESS_UNREACHABLE               = -109
	ERR_SSL_CLIENT_AUTH_CERT_NEEDED       = -110
	ERR_TUNNEL_CONNECTION_FAILED          = -111
	ERR_NO_SSL_VERSIONS_ENABLED           = -112
	ERR_SSL_VERSION_OR_CIPHER_MISMATCH    = -113
	ERR_SSL_RENEGOTIATION_REQUESTED       = -114
	ERR_PROXY_AUTH_UNSUPPORTED            = -115
	ERR_CERT_ERROR_IN_SSL_RENEGOTIATION   = -116
	ERR_BAD_SSL_CLIENT_AUTH_CERT          = -117
	ERR_CONNECTION_TIMED_OUT              = -118
	ERR_HOST_RESOLVER_QUEUE_TOO_LARGE     = -119
	ERR_SOCKS_CONNECTION_FAILED           = -120
	ERR_SOCKS_CONNECTION_HOST_UNREACHABLE = -121
	ERR_ALPN_NEGOTIATION_FAILED           = -122
	ERR_SSL_NO_RENEGOTIATION              = -123
	ERR_WINSOCK_UNEXPECTED_WRITTEN_BYTES  = -124
	ERR_SSL_DECOMPRESSION_FAILURE_ALERT   = -125
	ERR_SSL_BAD_RECORD_MAC_ALERT          = -126
	ERR_PROXY_AUTH_REQUESTED              = -127
	//
	ERR_PROXY_CONNECTION_FAILED              = -130
	ERR_MANDATORY_PROXY_CONFIGURATION_FAILED = -131
	//
	ERR_PRECONNECT_MAX_SOCKET_LIMIT               = -133
	ERR_SSL_CLIENT_AUTH_PRIVATE_KEY_ACCESS_DENIED = -134
	ERR_SSL_CLIENT_AUTH_CERT_NO_PRIVATE_KEY       = -135
	ERR_PROXY_CERTIFICATE_INVALID                 = -136
	ERR_NAME_RESOLUTION_FAILED                    = -137
	ERR_NETWORK_ACCESS_DENIED                     = -138
	ERR_TEMPORARILY_THROTTLED                     = -139
	ERR_HTTPS_PROXY_TUNNEL_RESPONSE_REDIRECT      = -140
	ERR_SSL_CLIENT_AUTH_SIGNATURE_FAILED          = -141
	ERR_MSG_TOO_BIG                               = -142
	//
	ERR_WS_PROTOCOL_ERROR = -145
	//
	ERR_ADDRESS_IN_USE                    = -147
	ERR_SSL_HANDSHAKE_NOT_COMPLETED       = -148
	ERR_SSL_BAD_PEER_PUBLIC_KEY           = -149
	ERR_SSL_PINNED_KEY_NOT_IN_CERT_CHAIN  = -150
	ERR_CLIENT_AUTH_CERT_TYPE_UNSUPPORTED = -151
	//
	ERR_SSL_DECRYPT_ERROR_ALERT     = -153
	ERR_WS_THROTTLE_QUEUE_TOO_LARGE = -154
	//
	ERR_SSL_SERVER_CERT_CHANGED = -156
	//
	ERR_SSL_UNRECOGNIZED_NAME_ALERT             = -159
	ERR_SOCKET_SET_RECEIVE_BUFFER_SIZE_ERROR    = -160
	ERR_SOCKET_SET_SEND_BUFFER_SIZE_ERROR       = -161
	ERR_SOCKET_RECEIVE_BUFFER_SIZE_UNCHANGEABLE = -162
	ERR_SOCKET_SEND_BUFFER_SIZE_UNCHANGEABLE    = -163
	ERR_SSL_CLIENT_AUTH_CERT_BAD_FORMAT         = -164
	//
	ERR_ICANN_NAME_COLLISION                      = -166
	ERR_SSL_SERVER_CERT_BAD_FORMAT                = -167
	ERR_CT_STH_PARSING_FAILED                     = -168
	ERR_CT_STH_INCOMPLETE                         = -169
	ERR_UNABLE_TO_REUSE_CONNECTION_FOR_PROXY_AUTH = -170
	ERR_CT_CONSISTENCY_PROOF_PARSING_FAILED       = -171
	ERR_SSL_OBSOLETE_CIPHER                       = -172
	ERR_WS_UPGRADE                                = -173
	ERR_READ_IF_READY_NOT_IMPLEMENTED             = -174
	//
	ERR_NO_BUFFER_SPACE                      = -176
	ERR_SSL_CLIENT_AUTH_NO_COMMON_ALGORITHMS = -177
	ERR_EARLY_DATA_REJECTED                  = -178
	ERR_WRONG_VERSION_ON_EARLY_DATA          = -179
	ERR_TLS13_DOWNGRADE_DETECTED             = -180
	ERR_SSL_KEY_USAGE_INCOMPATIBLE           = -181
	ERR_INVALID_ECH_CONFIG_LIST              = -182
	ERR_ECH_NOT_NEGOTIATED                   = -183
	ERR_ECH_FALLBACK_CERTIFICATE_INVALID     = -184
	//
	ERR_CERT_COMMON_NAME_INVALID        = -200
	ERR_CERT_DATE_INVALID               = -201
	ERR_CERT_AUTHORITY_INVALID          = -202
	ERR_CERT_CONTAINS_ERRORS            = -203
	ERR_CERT_NO_REVOCATION_MECHANISM    = -204
	ERR_CERT_UNABLE_TO_CHECK_REVOCATION = -205
	ERR_CERT_REVOKED                    = -206
	ERR_CERT_INVALID                    = -207
	ERR_CERT_WEAK_SIGNATURE_ALGORITHM   = -208
	//
	ERR_CERT_NON_UNIQUE_NAME              = -210
	ERR_CERT_WEAK_KEY                     = -211
	ERR_CERT_NAME_CONSTRAINT_VIOLATION    = -212
	ERR_CERT_VALIDITY_TOO_LONG            = -213
	ERR_CERTIFICATE_TRANSPARENCY_REQUIRED = -214
	ERR_CERT_SYMANTEC_LEGACY              = -215
	//
	ERR_CERT_KNOWN_INTERCEPTION_BLOCKED = -217
	//
	ERR_CERT_END = -219
	//
	ERR_INVALID_URL           = -300
	ERR_DISALLOWED_URL_SCHEME = -301
	ERR_UNKNOWN_URL_SCHEME    = -302
	ERR_INVALID_REDIRECT      = -303
	//
	ERR_TOO_MANY_REDIRECTS = -310
	ERR_UNSAFE_REDIRECT    = -311
	ERR_UNSAFE_PORT        = -312
	//
	ERR_INVALID_RESPONSE         = -320
	ERR_INVALID_CHUNKED_ENCODING = -321
	ERR_METHOD_NOT_SUPPORTED     = -322
	ERR_UNEXPECTED_PROXY_AUTH    = -323
	ERR_EMPTY_RESPONSE           = -324
	ERR_RESPONSE_HEADERS_TOO_BIG = -325
	//
	ERR_PAC_SCRIPT_FAILED                         = -327
	ERR_REQUEST_RANGE_NOT_SATISFIABLE             = -328
	ERR_MALFORMED_IDENTITY                        = -329
	ERR_CONTENT_DECODING_FAILED                   = -330
	ERR_NETWORK_IO_SUSPENDED                      = -331
	ERR_SYN_REPLY_NOT_RECEIVED                    = -332
	ERR_ENCODING_CONVERSION_FAILED                = -333
	ERR_UNRECOGNIZED_FTP_DIRECTORY_LISTING_FORMAT = -334
	//
	ERR_NO_SUPPORTED_PROXIES                          = -336
	ERR_SPDY_PROTOCOL_ERROR                           = -337
	ERR_INVALID_AUTH_CREDENTIALS                      = -338
	ERR_UNSUPPORTED_AUTH_SCHEME                       = -339
	ERR_ENCODING_DETECTION_FAILED                     = -340
	ERR_MISSING_AUTH_CREDENTIALS                      = -341
	ERR_UNEXPECTED_SECURITY_LIBRARY_STATUS            = -342
	ERR_MISCONFIGURED_AUTH_ENVIRONMENT                = -343
	ERR_UNDOCUMENTED_SECURITY_LIBRARY_STATUS          = -344
	ERR_RESPONSE_BODY_TOO_BIG_TO_DRAIN                = -345
	ERR_RESPONSE_HEADERS_MULTIPLE_CONTENT_LENGTH      = -346
	ERR_INCOMPLETE_SPDY_HEADERS                       = -347
	ERR_PAC_NOT_IN_DHCP                               = -348
	ERR_RESPONSE_HEADERS_MULTIPLE_CONTENT_DISPOSITION = -349
	ERR_RESPONSE_HEADERS_MULTIPLE_LOCATION            = -350
	ERR_SPDY_SERVER_REFUSED_STREAM                    = -351
	ERR_SPDY_PING_FAILED                              = -352
	//
	ERR_CONTENT_LENGTH_MISMATCH     = -354
	ERR_INCOMPLETE_CHUNKED_ENCODING = -355
	ERR_QUIC_PROTOCOL_ERROR         = -356
	ERR_RESPONSE_HEADERS_TRUNCATED  = -357
	ERR_QUIC_HANDSHAKE_FAILED       = -358
	//
	ERR_SPDY_INADEQUATE_TRANSPORT_SECURITY      = -360
	ERR_SPDY_FLOW_CONTROL_ERROR                 = -361
	ERR_SPDY_FRAME_SIZE_ERROR                   = -362
	ERR_SPDY_COMPRESSION_ERROR                  = -363
	ERR_PROXY_AUTH_REQUESTED_WITH_NO_CONNECTION = -364
	ERR_HTTP_1_1_REQUIRED                       = -365
	ERR_PROXY_HTTP_1_1_REQUIRED                 = -366
	ERR_PAC_SCRIPT_TERMINATED                   = -367
	//
	ERR_INVALID_HTTP_RESPONSE                      = -370
	ERR_CONTENT_DECODING_INIT_FAILED               = -371
	ERR_SPDY_RST_STREAM_NO_ERROR_RECEIVED          = -372
	ERR_SPDY_PUSHED_STREAM_NOT_AVAILABLE           = -373
	ERR_SPDY_CLAIMED_PUSHED_STREAM_RESET_BY_SERVER = -374
	ERR_TOO_MANY_RETRIES                           = -375
	ERR_SPDY_STREAM_CLOSED                         = -376
	ERR_SPDY_CLIENT_REFUSED_STREAM                 = -377
	ERR_SPDY_PUSHED_RESPONSE_DOES_NOT_MATCH        = -378
	ERR_HTTP_RESPONSE_CODE_FAILURE                 = -379
	ERR_QUIC_CERT_ROOT_NOT_KNOWN                   = -380
	ERR_QUIC_GOAWAY_REQUEST_CAN_BE_RETRIED         = -381
	//
	ERR_CACHE_MISS                    = -400
	ERR_CACHE_READ_FAILURE            = -401
	ERR_CACHE_WRITE_FAILURE           = -402
	ERR_CACHE_OPERATION_NOT_SUPPORTED = -403
	ERR_CACHE_OPEN_FAILURE            = -404
	ERR_CACHE_CREATE_FAILURE          = -405
	ERR_CACHE_RACE                    = -406
	ERR_CACHE_CHECKSUM_READ_FAILURE   = -407
	ERR_CACHE_CHECKSUM_MISMATCH       = -408
	ERR_CACHE_LOCK_TIMEOUT            = -409
	ERR_CACHE_AUTH_FAILURE_AFTER_READ = -410
	ERR_CACHE_ENTRY_NOT_SUITABLE      = -411
	ERR_CACHE_DOOM_FAILURE            = -412
	ERR_CACHE_OPEN_OR_CREATE_FAILURE  = -413
	//
	ERR_INSECURE_RESPONSE                                     = -501
	ERR_NO_PRIVATE_KEY_FOR_CERT                               = -502
	ERR_ADD_USER_CERT_FAILED                                  = -503
	ERR_INVALID_SIGNED_EXCHANGE                               = -504
	ERR_INVALID_WEB_BUNDLE                                    = -505
	ERR_TRUST_TOKEN_OPERATION_FAILED                          = -506
	ERR_TRUST_TOKEN_OPERATION_SUCCESS_WITHOUT_SENDING_REQUEST = -507
	//
	ERR_FTP_FAILED                = -601
	ERR_FTP_SERVICE_UNAVAILABLE   = -602
	ERR_FTP_TRANSFER_ABORTED      = -603
	ERR_FTP_FILE_BUSY             = -604
	ERR_FTP_SYNTAX_ERROR          = -605
	ERR_FTP_COMMAND_NOT_SUPPORTED = -606
	ERR_FTP_BAD_COMMAND_SEQUENCE  = -607
	//
	ERR_PKCS12_IMPORT_BAD_PASSWORD = -701
	ERR_PKCS12_IMPORT_FAILED       = -702
	ERR_IMPORT_CA_CERT_NOT_CA      = -703
	ERR_IMPORT_CERT_ALREADY_EXISTS = -704
	ERR_IMPORT_CA_CERT_FAILED      = -705
	ERR_IMPORT_SERVER_CERT_FAILED  = -706
	ERR_PKCS12_IMPORT_INVALID_MAC  = -707
	ERR_PKCS12_IMPORT_INVALID_FILE = -708
	ERR_PKCS12_IMPORT_UNSUPPORTED  = -709
	ERR_KEY_GENERATION_FAILED      = -710
	//
	ERR_PRIVATE_KEY_EXPORT_FAILED          = -712
	ERR_SELF_SIGNED_CERT_GENERATION_FAILED = -713
	ERR_CERT_DATABASE_CHANGED              = -714
	//
	ERR_DNS_MALFORMED_RESPONSE  = -800
	ERR_DNS_SERVER_REQUIRES_TCP = -801
	ERR_DNS_SERVER_FAILED       = -802
	ERR_DNS_TIMED_OUT           = -803
	ERR_NS_CACHE_MISS           = -804
	ERR_DNS_SEARCH_EMPTY        = -805
	ERR_DNS_SORT_ERROR          = -806
	//
	ERR_DNS_SECURE_RESOLVER_HOSTNAME_RESOLUTION_FAILED = -808
	ERR_DNS_NAME_HTTPS_ONLY                            = -809
)
const (
	/// <summary>
	/// command_id constants declared in cef_command_ids.h and used by some callbacks in ICefCommandHandler
	/// cef_command_ids.h is generated in /include/cef_command_ids.h
	/// </summary>
	/// <remarks>
	/// <para><see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/app/chrome_command_ids.h">The command_id values are also available in chrome/app/chrome_command_ids.h</see></para>
	/// </remarks>
	IDC_MinimumLabelValue                                         = 4000
	IDC_BACK                                                      = 33000
	IDC_FORWARD                                                   = 33001
	IDC_RELOAD                                                    = 33002
	IDC_HOME                                                      = 33003
	IDC_OPEN_CURRENT_URL                                          = 33004
	IDC_STOP                                                      = 33006
	IDC_RELOAD_BYPASSING_CACHE                                    = 33007
	IDC_RELOAD_CLEARING_CACHE                                     = 33009
	IDC_NEW_WINDOW                                                = 34000
	IDC_NEW_INCOGNITO_WINDOW                                      = 34001
	IDC_PIN_TO_START_SCREEN                                       = 34005
	IDC_CLOSE_WINDOW                                              = 34012
	IDC_ALWAYS_ON_TOP                                             = 34013
	IDC_NEW_TAB                                                   = 34014
	IDC_CLOSE_TAB                                                 = 34015
	IDC_SELECT_NEXT_TAB                                           = 34016
	IDC_SELECT_PREVIOUS_TAB                                       = 34017
	IDC_SELECT_TAB_0                                              = 34018
	IDC_SELECT_TAB_1                                              = 34019
	IDC_SELECT_TAB_2                                              = 34020
	IDC_SELECT_TAB_3                                              = 34021
	IDC_SELECT_TAB_4                                              = 34022
	IDC_SELECT_TAB_5                                              = 34023
	IDC_SELECT_TAB_6                                              = 34024
	IDC_SELECT_TAB_7                                              = 34025
	IDC_SELECT_LAST_TAB                                           = 34026
	IDC_DUPLICATE_TAB                                             = 34027
	IDC_RESTORE_TAB                                               = 34028
	IDC_SHOW_AS_TAB                                               = 34029
	IDC_FULLSCREEN                                                = 34030
	IDC_EXIT                                                      = 34031
	IDC_MOVE_TAB_NEXT                                             = 34032
	IDC_MOVE_TAB_PREVIOUS                                         = 34033
	IDC_SEARCH                                                    = 34035
	IDC_DEBUG_FRAME_TOGGLE                                        = 34038
	IDC_WINDOW_MENU                                               = 34045
	IDC_MINIMIZE_WINDOW                                           = 34046
	IDC_MAXIMIZE_WINDOW                                           = 34047
	IDC_ALL_WINDOWS_FRONT                                         = 34048
	IDC_NAME_WINDOW                                               = 34049
	IDC_TOGGLE_MULTITASK_MENU                                     = 34050
	IDC_USE_SYSTEM_TITLE_BAR                                      = 34051
	IDC_RESTORE_WINDOW                                            = 34052
	IDC_OPEN_IN_PWA_WINDOW                                        = 34053
	IDC_MOVE_TAB_TO_NEW_WINDOW                                    = 34054
	IDC_COPY_URL                                                  = 34060
	IDC_OPEN_IN_CHROME                                            = 34061
	IDC_SITE_SETTINGS                                             = 34062
	IDC_WEB_APP_MENU_APP_INFO                                     = 34063
	IDC_VISIT_DESKTOP_OF_LRU_USER_2                               = 34080
	IDC_VISIT_DESKTOP_OF_LRU_USER_3                               = 34081
	IDC_VISIT_DESKTOP_OF_LRU_USER_4                               = 34082
	IDC_VISIT_DESKTOP_OF_LRU_USER_5                               = 34083
	IDC_BOOKMARK_THIS_TAB                                         = 35000
	IDC_BOOKMARK_ALL_TABS                                         = 35001
	IDC_VIEW_SOURCE                                               = 35002
	IDC_PRINT                                                     = 35003
	IDC_SAVE_PAGE                                                 = 35004
	IDC_EMAIL_PAGE_LOCATION                                       = 35006
	IDC_BASIC_PRINT                                               = 35007
	IDC_SAVE_CREDIT_CARD_FOR_PAGE                                 = 35008
	IDC_TRANSLATE_PAGE                                            = 35009
	IDC_MANAGE_PASSWORDS_FOR_PAGE                                 = 35010
	IDC_ROUTE_MEDIA                                               = 35011
	IDC_WINDOW_MUTE_SITE                                          = 35012
	IDC_WINDOW_PIN_TAB                                            = 35013
	IDC_WINDOW_GROUP_TAB                                          = 35014
	IDC_MIGRATE_LOCAL_CREDIT_CARD_FOR_PAGE                        = 35015
	IDC_SEND_TAB_TO_SELF                                          = 35016
	IDC_FOCUS_THIS_TAB                                            = 35017
	IDC_QRCODE_GENERATOR                                          = 35021
	IDC_WINDOW_CLOSE_TABS_TO_RIGHT                                = 35022
	IDC_WINDOW_CLOSE_OTHER_TABS                                   = 35023
	IDC_NEW_TAB_TO_RIGHT                                          = 35024
	IDC_SAVE_AUTOFILL_ADDRESS                                     = 35025
	IDC_OFFERS_AND_REWARDS_FOR_PAGE                               = 35026
	IDC_WEBAUTHN                                                  = 35027
	IDC_SHARING_HUB                                               = 35028
	IDC_SHARING_HUB_MENU                                          = 35029
	IDC_VIRTUAL_CARD_MANUAL_FALLBACK                              = 35030
	IDC_SHARING_HUB_SCREENSHOT                                    = 35031
	IDC_VIRTUAL_CARD_ENROLL                                       = 35032
	IDC_FOLLOW                                                    = 35033
	IDC_UNFOLLOW                                                  = 35034
	IDC_SAVE_IBAN_FOR_PAGE                                        = 35035
	IDC_MUTE_TARGET_SITE                                          = 35050
	IDC_PIN_TARGET_TAB                                            = 35051
	IDC_GROUP_TARGET_TAB                                          = 35052
	IDC_DUPLICATE_TARGET_TAB                                      = 35053
	IDC_CUT                                                       = 36000
	IDC_COPY                                                      = 36001
	IDC_PASTE                                                     = 36003
	IDC_EDIT_MENU                                                 = 36004
	IDC_FIND                                                      = 37000
	IDC_FIND_NEXT                                                 = 37001
	IDC_FIND_PREVIOUS                                             = 37002
	IDC_CLOSE_FIND_OR_STOP                                        = 37003
	IDC_FIND_MENU                                                 = 37100
	IDC_ZOOM_MENU                                                 = 38000
	IDC_ZOOM_PLUS                                                 = 38001
	IDC_ZOOM_NORMAL                                               = 38002
	IDC_ZOOM_MINUS                                                = 38003
	IDC_ZOOM_PERCENT_DISPLAY                                      = 38004
	IDC_FOCUS_TOOLBAR                                             = 39000
	IDC_FOCUS_LOCATION                                            = 39001
	IDC_FOCUS_SEARCH                                              = 39002
	IDC_FOCUS_MENU_BAR                                            = 39003
	IDC_FOCUS_NEXT_PANE                                           = 39004
	IDC_FOCUS_PREVIOUS_PANE                                       = 39005
	IDC_FOCUS_BOOKMARKS                                           = 39006
	IDC_FOCUS_INACTIVE_POPUP_FOR_ACCESSIBILITY                    = 39007
	IDC_FOCUS_WEB_CONTENTS_PANE                                   = 39009
	IDC_OPEN_FILE                                                 = 40000
	IDC_CREATE_SHORTCUT                                           = 40002
	IDC_DEVELOPER_MENU                                            = 40003
	IDC_DEV_TOOLS                                                 = 40004
	IDC_DEV_TOOLS_CONSOLE                                         = 40005
	IDC_TASK_MANAGER                                              = 40006
	IDC_DEV_TOOLS_DEVICES                                         = 40007
	IDC_FEEDBACK                                                  = 40008
	IDC_SHOW_BOOKMARK_BAR                                         = 40009
	IDC_SHOW_HISTORY                                              = 40010
	IDC_SHOW_BOOKMARK_MANAGER                                     = 40011
	IDC_SHOW_DOWNLOADS                                            = 40012
	IDC_CLEAR_BROWSING_DATA                                       = 40013
	IDC_IMPORT_SETTINGS                                           = 40014
	IDC_OPTIONS                                                   = 40015
	IDC_EDIT_SEARCH_ENGINES                                       = 40016
	IDC_VIEW_PASSWORDS                                            = 40017
	IDC_ABOUT                                                     = 40018
	IDC_HELP_PAGE_VIA_KEYBOARD                                    = 40019
	IDC_HELP_PAGE_VIA_MENU                                        = 40020
	IDC_SHOW_APP_MENU                                             = 40021
	IDC_MANAGE_EXTENSIONS                                         = 40022
	IDC_DEV_TOOLS_INSPECT                                         = 40023
	IDC_UPGRADE_DIALOG                                            = 40024
	IDC_PROFILING_ENABLED                                         = 40028
	IDC_BOOKMARKS_MENU                                            = 40029
	IDC_SHOW_SIGNIN                                               = 40030
	IDC_EXTENSION_ERRORS                                          = 40031
	IDC_SHOW_SETTINGS_CHANGE_FIRST                                = 40033
	IDC_SHOW_SETTINGS_CHANGE_LAST                                 = 40133
	IDC_SHOW_AVATAR_MENU                                          = 40134
	IDC_EXTENSION_INSTALL_ERROR_FIRST                             = 40135
	IDC_EXTENSION_INSTALL_ERROR_LAST                              = 40235
	IDC_TOGGLE_REQUEST_TABLET_SITE                                = 40236
	IDC_DEV_TOOLS_TOGGLE                                          = 40237
	IDC_RECENT_TABS_MENU                                          = 40239
	IDC_RECENT_TABS_NO_DEVICE_TABS                                = 40240
	IDC_SHOW_SETTINGS_RESET_BUBBLE                                = 40241
	IDC_SHOW_SYNC_ERROR                                           = 40242
	IDC_DISTILL_PAGE                                              = 40243
	IDC_HELP_MENU                                                 = 40244
	IDC_SHOW_SRT_BUBBLE                                           = 40246
	IDC_ELEVATED_RECOVERY_DIALOG                                  = 40247
	IDC_TAKE_SCREENSHOT                                           = 40248
	IDC_MORE_TOOLS_MENU                                           = 40249
	IDC_TOGGLE_FULLSCREEN_TOOLBAR                                 = 40250
	IDC_CUSTOMIZE_TOUCH_BAR                                       = 40251
	IDC_SHOW_BETA_FORUM                                           = 40252
	IDC_TOGGLE_JAVASCRIPT_APPLE_EVENTS                            = 40253
	IDC_INSTALL_PWA                                               = 40254
	IDC_SHOW_MANAGEMENT_PAGE                                      = 40255
	IDC_PASTE_AND_GO                                              = 40256
	IDC_SHOW_SAVE_LOCAL_CARD_SIGN_IN_PROMO_IF_APPLICABLE          = 40257
	IDC_CLOSE_SIGN_IN_PROMO                                       = 40258
	IDC_SHOW_FULL_URLS                                            = 40259
	IDC_CARET_BROWSING_TOGGLE                                     = 40260
	IDC_TOGGLE_QUICK_COMMANDS                                     = 40261
	IDC_CHROME_TIPS                                               = 40263
	IDC_CHROME_WHATS_NEW                                          = 40264
	IDC_LACROS_DATA_MIGRATION                                     = 40265
	IDC_PERFORMANCE                                               = 40266
	IDC_SPELLCHECK_SUGGESTION_0                                   = 41000
	IDC_SPELLCHECK_SUGGESTION_1                                   = 41001
	IDC_SPELLCHECK_SUGGESTION_2                                   = 41002
	IDC_SPELLCHECK_SUGGESTION_3                                   = 41003
	IDC_SPELLCHECK_SUGGESTION_4                                   = 41004
	IDC_SPELLCHECK_MENU                                           = 41005
	IDC_SPELLCHECK_LANGUAGES_FIRST                                = 41006
	IDC_SPELLCHECK_LANGUAGES_LAST                                 = 41106
	IDC_CHECK_SPELLING_WHILE_TYPING                               = 41107
	IDC_SPELLPANEL_TOGGLE                                         = 41109
	IDC_SPELLCHECK_ADD_TO_DICTIONARY                              = 41110
	IDC_SPELLCHECK_MULTI_LINGUAL                                  = 41111
	IDC_WRITING_DIRECTION_MENU                                    = 41120
	IDC_WRITING_DIRECTION_DEFAULT                                 = 41121
	IDC_WRITING_DIRECTION_LTR                                     = 41122
	IDC_WRITING_DIRECTION_RTL                                     = 41123
	IDC_TRANSLATE_ORIGINAL_LANGUAGE_BASE                          = 42100
	IDC_TRANSLATE_TARGET_LANGUAGE_BASE                            = 42400
	IDC_VIEW_MENU                                                 = 44000
	IDC_FILE_MENU                                                 = 44001
	IDC_CHROME_MENU                                               = 44002
	IDC_HIDE_APP                                                  = 44003
	IDC_HISTORY_MENU                                              = 46000
	IDC_TAB_MENU                                                  = 46001
	IDC_PROFILE_MAIN_MENU                                         = 46100
	IDC_INPUT_METHODS_MENU                                        = 46300
	IDC_CONTENT_CONTEXT_CUSTOM_FIRST                              = 47000
	IDC_CONTENT_CONTEXT_CUSTOM_LAST                               = 48000
	IDC_EXTENSIONS_CONTEXT_CUSTOM_FIRST                           = 49000
	IDC_EXTENSIONS_CONTEXT_CUSTOM_LAST                            = 50000
	IDC_CONTENT_CONTEXT_OPENLINKNEWTAB                            = 50100
	IDC_CONTENT_CONTEXT_OPENLINKNEWWINDOW                         = 50101
	IDC_CONTENT_CONTEXT_OPENLINKOFFTHERECORD                      = 50102
	IDC_CONTENT_CONTEXT_SAVELINKAS                                = 50103
	IDC_CONTENT_CONTEXT_COPYLINKLOCATION                          = 50104
	IDC_CONTENT_CONTEXT_COPYEMAILADDRESS                          = 50105
	IDC_CONTENT_CONTEXT_OPENLINKWITH                              = 50106
	IDC_CONTENT_CONTEXT_COPYLINKTEXT                              = 50107
	IDC_CONTENT_CONTEXT_OPENLINKINPROFILE                         = 50108
	IDC_CONTENT_CONTEXT_OPENLINKBOOKMARKAPP                       = 50109
	IDC_CONTENT_CONTEXT_SAVEIMAGEAS                               = 50110
	IDC_CONTENT_CONTEXT_COPYIMAGELOCATION                         = 50111
	IDC_CONTENT_CONTEXT_COPYIMAGE                                 = 50112
	IDC_CONTENT_CONTEXT_OPENIMAGENEWTAB                           = 50113
	IDC_CONTENT_CONTEXT_SEARCHWEBFORIMAGE                         = 50114
	IDC_CONTENT_CONTEXT_OPEN_ORIGINAL_IMAGE_NEW_TAB               = 50115
	IDC_CONTENT_CONTEXT_LOAD_IMAGE                                = 50116
	IDC_CONTENT_CONTEXT_SEARCHLENSFORIMAGE                        = 50117
	IDC_CONTENT_CONTEXT_TRANSLATEIMAGEWITHWEB                     = 50118
	IDC_CONTENT_CONTEXT_TRANSLATEIMAGEWITHLENS                    = 50119
	IDC_CONTENT_CONTEXT_SAVEAVAS                                  = 50120
	IDC_CONTENT_CONTEXT_COPYAVLOCATION                            = 50121
	IDC_CONTENT_CONTEXT_OPENAVNEWTAB                              = 50122
	IDC_CONTENT_CONTEXT_PICTUREINPICTURE                          = 50123
	IDC_CONTENT_CONTEXT_PLAYPAUSE                                 = 50130
	IDC_CONTENT_CONTEXT_MUTE                                      = 50131
	IDC_CONTENT_CONTEXT_LOOP                                      = 50132
	IDC_CONTENT_CONTEXT_CONTROLS                                  = 50133
	IDC_CONTENT_CONTEXT_ROTATECW                                  = 50134
	IDC_CONTENT_CONTEXT_ROTATECCW                                 = 50135
	IDC_CONTENT_CONTEXT_COPY                                      = 50140
	IDC_CONTENT_CONTEXT_CUT                                       = 50141
	IDC_CONTENT_CONTEXT_PASTE                                     = 50142
	IDC_CONTENT_CONTEXT_DELETE                                    = 50143
	IDC_CONTENT_CONTEXT_UNDO                                      = 50144
	IDC_CONTENT_CONTEXT_REDO                                      = 50145
	IDC_CONTENT_CONTEXT_SELECTALL                                 = 50146
	IDC_CONTENT_CONTEXT_PASTE_AND_MATCH_STYLE                     = 50147
	IDC_CONTENT_CONTEXT_COPYLINKTOTEXT                            = 50148
	IDC_CONTENT_CONTEXT_RESHARELINKTOTEXT                         = 50149
	IDC_CONTENT_CONTEXT_REMOVELINKTOTEXT                          = 50150
	IDC_CONTENT_CONTEXT_TRANSLATE                                 = 50151
	IDC_CONTENT_CONTEXT_INSPECTELEMENT                            = 50152
	IDC_CONTENT_CONTEXT_VIEWPAGEINFO                              = 50153
	IDC_CONTENT_CONTEXT_LANGUAGE_SETTINGS                         = 50154
	IDC_CONTENT_CONTEXT_LOOK_UP                                   = 50155
	IDC_CONTENT_CONTEXT_NO_SPELLING_SUGGESTIONS                   = 50156
	IDC_CONTENT_CONTEXT_SPELLING_SUGGESTION                       = 50157
	IDC_CONTENT_CONTEXT_SPELLING_TOGGLE                           = 50158
	IDC_CONTENT_CONTEXT_OPEN_IN_READING_MODE                      = 50159
	IDC_CONTENT_CONTEXT_INSPECTBACKGROUNDPAGE                     = 50161
	IDC_CONTENT_CONTEXT_RELOAD_PACKAGED_APP                       = 50162
	IDC_CONTENT_CONTEXT_RESTART_PACKAGED_APP                      = 50163
	IDC_CONTENT_CONTEXT_LENS_REGION_SEARCH                        = 50164
	IDC_CONTENT_CONTEXT_WEB_REGION_SEARCH                         = 50165
	IDC_CONTENT_CONTEXT_GENERATEPASSWORD                          = 50166
	IDC_CONTENT_CONTEXT_EXIT_FULLSCREEN                           = 50167
	IDC_CONTENT_CONTEXT_SHOWALLSAVEDPASSWORDS                     = 50168
	IDC_CONTENT_CONTEXT_PARTIAL_TRANSLATE                         = 50169
	IDC_CONTENT_CONTEXT_RELOADFRAME                               = 50170
	IDC_CONTENT_CONTEXT_VIEWFRAMESOURCE                           = 50171
	IDC_CONTENT_CONTEXT_VIEWFRAMEINFO                             = 50172
	IDC_CONTENT_CONTEXT_ADD_A_NOTE                                = 50175
	IDC_CONTENT_CONTEXT_GOTOURL                                   = 50180
	IDC_CONTENT_CONTEXT_SEARCHWEBFOR                              = 50181
	IDC_CONTENT_CONTEXT_OPEN_WITH1                                = 50190
	IDC_CONTENT_CONTEXT_OPEN_WITH2                                = 50191
	IDC_CONTENT_CONTEXT_OPEN_WITH3                                = 50192
	IDC_CONTENT_CONTEXT_OPEN_WITH4                                = 50193
	IDC_CONTENT_CONTEXT_OPEN_WITH5                                = 50194
	IDC_CONTENT_CONTEXT_OPEN_WITH6                                = 50195
	IDC_CONTENT_CONTEXT_OPEN_WITH7                                = 50196
	IDC_CONTENT_CONTEXT_OPEN_WITH8                                = 50197
	IDC_CONTENT_CONTEXT_OPEN_WITH9                                = 50198
	IDC_CONTENT_CONTEXT_OPEN_WITH10                               = 50199
	IDC_CONTENT_CONTEXT_OPEN_WITH11                               = 50200
	IDC_CONTENT_CONTEXT_OPEN_WITH12                               = 50201
	IDC_CONTENT_CONTEXT_OPEN_WITH13                               = 50202
	IDC_CONTENT_CONTEXT_OPEN_WITH14                               = 50203
	IDC_CONTENT_CONTEXT_EMOJI                                     = 50210
	IDC_BOOKMARK_BAR_OPEN_ALL                                     = 51000
	IDC_BOOKMARK_BAR_OPEN_ALL_NEW_WINDOW                          = 51001
	IDC_BOOKMARK_BAR_OPEN_ALL_INCOGNITO                           = 51002
	IDC_BOOKMARK_BAR_OPEN_INCOGNITO                               = 51003
	IDC_BOOKMARK_BAR_OPEN_ALL_NEW_TAB_GROUP                       = 51004
	IDC_BOOKMARK_BAR_RENAME_FOLDER                                = 51005
	IDC_BOOKMARK_BAR_EDIT                                         = 51006
	IDC_BOOKMARK_BAR_REMOVE                                       = 51007
	IDC_BOOKMARK_BAR_UNDO                                         = 51008
	IDC_BOOKMARK_BAR_REDO                                         = 51009
	IDC_BOOKMARK_BAR_ADD_NEW_BOOKMARK                             = 51010
	IDC_BOOKMARK_BAR_NEW_FOLDER                                   = 51011
	IDC_BOOKMARK_MANAGER                                          = 51012
	IDC_BOOKMARK_BAR_ALWAYS_SHOW                                  = 51013
	IDC_BOOKMARK_BAR_SHOW_APPS_SHORTCUT                           = 51014
	IDC_BOOKMARK_BAR_SHOW_READING_LIST                            = 51015
	IDC_BOOKMARK_BAR_SHOW_MANAGED_BOOKMARKS                       = 51016
	IDC_BOOKMARK_BAR_TRACK_PRICE_FOR_SHOPPING_BOOKMARK            = 51017
	IDC_BOOKMARK_BAR_UNTRACK_PRICE_FOR_SHOPPING_BOOKMARK          = 51018
	IDC_BOOKMARK_BAR_ADD_TO_BOOKMARKS_BAR                         = 51019
	IDC_BOOKMARK_BAR_REMOVE_FROM_BOOKMARKS_BAR                    = 51020
	IDC_CONTENT_CONTEXT_SHARING_CLICK_TO_CALL_SINGLE_DEVICE       = 51030
	IDC_CONTENT_CONTEXT_SHARING_CLICK_TO_CALL_MULTIPLE_DEVICES    = 51031
	IDC_CONTENT_CONTEXT_SHARING_SHARED_CLIPBOARD_SINGLE_DEVICE    = 51032
	IDC_CONTENT_CONTEXT_SHARING_SHARED_CLIPBOARD_MULTIPLE_DEVICES = 51033
	IDC_CONTENT_CONTEXT_GENERATE_QR_CODE                          = 51034
	IDC_CONTENT_CONTEXT_SHARING_SUBMENU                           = 51035
	IDC_CONTENT_CLIPBOARD_HISTORY_MENU                            = 51036
	IDC_STATUS_TRAY_KEEP_CHROME_RUNNING_IN_BACKGROUND             = 51100
	IDC_MEDIA_ROUTER_ABOUT                                        = 51200
	IDC_MEDIA_ROUTER_HELP                                         = 51201
	IDC_MEDIA_ROUTER_LEARN_MORE                                   = 51202
	IDC_MEDIA_ROUTER_ALWAYS_SHOW_TOOLBAR_ACTION                   = 51204
	IDC_MEDIA_ROUTER_SHOWN_BY_POLICY                              = 51206
	IDC_MEDIA_ROUTER_SHOW_IN_TOOLBAR                              = 51207
	IDC_MEDIA_ROUTER_TOGGLE_MEDIA_REMOTING                        = 51208
	IDC_MEDIA_TOOLBAR_CONTEXT_REPORT_CAST_ISSUE                   = 51209
	IDC_MEDIA_TOOLBAR_CONTEXT_SHOW_OTHER_SESSIONS                 = 51210
	IDC_MEDIA_STREAM_DEVICE_STATUS_TRAY                           = 51300
	IDC_MEDIA_CONTEXT_MEDIA_STREAM_CAPTURE_LIST_FIRST             = 51301
	IDC_MEDIA_CONTEXT_MEDIA_STREAM_CAPTURE_LIST_LAST              = 51399
	IDC_MEDIA_STREAM_DEVICE_ALWAYS_ALLOW                          = 51400
	IDC_CONTENT_CONTEXT_PROTOCOL_HANDLER_FIRST                    = 52000
	IDC_CONTENT_CONTEXT_PROTOCOL_HANDLER_LAST                     = 52199
	IDC_CONTENT_CONTEXT_PROTOCOL_HANDLER_SETTINGS                 = 52200
	IDC_OPEN_LINK_IN_PROFILE_FIRST                                = 52300
	IDC_OPEN_LINK_IN_PROFILE_LAST                                 = 52399
	IDC_CONTENT_CONTEXT_START_SMART_SELECTION_ACTION1             = 52400
	IDC_CONTENT_CONTEXT_START_SMART_SELECTION_ACTION2             = 52401
	IDC_CONTENT_CONTEXT_START_SMART_SELECTION_ACTION3             = 52402
	IDC_CONTENT_CONTEXT_START_SMART_SELECTION_ACTION4             = 52403
	IDC_CONTENT_CONTEXT_START_SMART_SELECTION_ACTION5             = 52404
	IDC_CONTENT_CONTEXT_ACCESSIBILITY_LABELS_TOGGLE               = 52410
	IDC_CONTENT_CONTEXT_ACCESSIBILITY_LABELS                      = 52411
	IDC_CONTENT_CONTEXT_ACCESSIBILITY_LABELS_TOGGLE_ONCE          = 52412
	IDC_CONTENT_CONTEXT_QUICK_ANSWERS_INLINE_ANSWER               = 52413
	IDC_CONTENT_CONTEXT_QUICK_ANSWERS_INLINE_QUERY                = 52414
	IDC_RUN_SCREEN_AI_VISUAL_ANNOTATIONS                          = 52420
	IDC_CONTENT_CONTEXT_PDF_OCR                                   = 52421
	IDC_CONTENT_CONTEXT_PDF_OCR_ALWAYS                            = 52422
	IDC_CONTENT_CONTEXT_PDF_OCR_ONCE                              = 52423
	IDC_TAB_SEARCH                                                = 52500
	IDC_TAB_SEARCH_CLOSE                                          = 52501
	IDC_DEBUG_TOGGLE_TABLET_MODE                                  = 52510
	IDC_DEBUG_PRINT_VIEW_TREE                                     = 52511
	IDC_DEBUG_PRINT_VIEW_TREE_DETAILS                             = 52512
	IDC_CONTENT_CONTEXT_AUTOFILL_FEEDBACK                         = 52990
	IDC_CONTENT_CONTEXT_AUTOFILL_CUSTOM_FIRST                     = 53000
	IDC_CONTENT_CONTEXT_AUTOFILL_CUSTOM_LAST                      = 53250
	IDC_LIVE_CAPTION                                              = 53251
	IDC_MANAGE_HID_DEVICES_FIRST                                  = 53260
	IDC_MANAGE_HID_DEVICES_LAST                                   = 53269
	IDC_FIRST_UNBOUNDED_MENU                                      = 0
)
const (
	/// <summary>
	/// Supported certificate status code values. See net\cert\cert_status_flags.h
	/// for more information. CERT_STATUS_NONE is new in CEF because we use an
	/// enum while cert_status_flags.h uses a typedef and static const variables.
	/// </summary>
	/// <remarks>
	/// <para>TCefCertStatus values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cert_status_t)</see></para>
	/// </remarks>
	CERT_STATUS_NONE                       = 0
	CERT_STATUS_COMMON_NAME_INVALID        = 1 << 0
	CERT_STATUS_DATE_INVALID               = 1 << 1
	CERT_STATUS_AUTHORITY_INVALID          = 1 << 2
	CERT_STATUS_NO_REVOCATION_MECHANISM    = 1 << 4
	CERT_STATUS_UNABLE_TO_CHECK_REVOCATION = 1 << 5
	CERT_STATUS_REVOKED                    = 1 << 6
	CERT_STATUS_INVALID                    = 1 << 7
	CERT_STATUS_WEAK_SIGNATURE_ALGORITHM   = 1 << 8
	CERT_STATUS_NON_UNIQUE_NAME            = 1 << 10
	CERT_STATUS_WEAK_KEY                   = 1 << 11
	CERT_STATUS_PINNED_KEY_MISSING         = 1 << 13
	CERT_STATUS_NAME_CONSTRAINT_VIOLATION  = 1 << 14
	CERT_STATUS_VALIDITY_TOO_LONG          = 1 << 15
	CERT_STATUS_IS_EV                      = 1 << 16
	CERT_STATUS_REV_CHECKING_ENABLED       = 1 << 17
	CERT_STATUS_SHA1_SIGNATURE_PRESENT     = 1 << 19
	CERT_STATUS_CT_COMPLIANCE_FAILED       = 1 << 20
	CERT_STATUS_FIRST_ERROR                = CERT_STATUS_COMMON_NAME_INVALID
	CERT_STATUS_LAST_ERROR                 = CERT_STATUS_VALIDITY_TOO_LONG
)
const (
	/// <summary>
	/// V8 access control values.
	/// </summary>
	/// <remarks>
	/// <para>TCefV8AccessControls values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_v8_accesscontrol_t)</see></para>
	/// </remarks>
	V8_ACCESS_CONTROL_DEFAULT               = 0
	V8_ACCESS_CONTROL_ALL_CAN_READ          = 1 << 0
	V8_ACCESS_CONTROL_ALL_CAN_WRITE         = 1 << 1
	V8_ACCESS_CONTROL_PROHIBITS_OVERWRITING = 1 << 2
)

const (
	/// <summary>
	/// Source is a link click or the JavaScript window.open function. This is
	/// also the default value for requests like sub-resource loads that are not
	/// navigations.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_LINK = 0
	/// <summary>
	/// Source is some other "explicit" navigation. This is the default value for
	/// navigations where the actual type is unknown. See also
	/// TT_DIRECT_LOAD_FLAG.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_EXPLICIT = 1
	/// <summary>
	/// User got to this page through a suggestion in the UI (for example, via the
	/// destinations page). Chrome runtime only.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_AUTO_BOOKMARK = 2
	/// <summary>
	/// Source is a subframe navigation. This is any content that is automatically
	/// loaded in a non-toplevel frame. For example, if a page consists of several
	/// frames containing ads, those ad URLs will have this transition type.
	/// The user may not even realize the content in these pages is a separate
	/// frame, so may not care about the URL.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_AUTO_SUBFRAME = 3
	/// <summary>
	/// Source is a subframe navigation explicitly requested by the user that will
	/// generate new navigation entries in the back/forward list. These are
	/// probably more important than frames that were automatically loaded in
	/// the background because the user probably cares about the fact that this
	/// link was loaded.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_MANUAL_SUBFRAME = 4
	/// <summary>
	/// User got to this page by typing in the URL bar and selecting an entry
	/// that did not look like a URL.  For example, a match might have the URL
	/// of a Google search result page, but appear like "Search Google for ...".
	/// These are not quite the same as EXPLICIT navigations because the user
	/// didn't type or see the destination URL. Chrome runtime only.
	/// See also TT_KEYWORD.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_GENERATED = 5
	/// <summary>
	/// This is a toplevel navigation. This is any content that is automatically
	/// loaded in a toplevel frame.  For example, opening a tab to show the ASH
	/// screen saver, opening the devtools window, opening the NTP after the safe
	/// browsing warning, opening web-based dialog boxes are examples of
	/// AUTO_TOPLEVEL navigations. Chrome runtime only.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_AUTO_TOPLEVEL = 6
	/// <summary>
	/// Source is a form submission by the user. NOTE: In some situations
	/// submitting a form does not result in this transition type. This can happen
	/// if the form uses a script to submit the contents.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_FORM_SUBMIT = 7
	/// <summary>
	/// Source is a "reload" of the page via the Reload function or by re-visiting
	/// the same URL. NOTE: This is distinct from the concept of whether a
	/// particular load uses "reload semantics" (i.e. bypasses cached data).
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_RELOAD = 8
	/// <summary>
	/// The url was generated from a replaceable keyword other than the default
	/// search provider. If the user types a keyword (which also applies to
	/// tab-to-search) in the omnibox this qualifier is applied to the transition
	/// type of the generated url. TemplateURLModel then may generate an
	/// additional visit with a transition type of TT_KEYWORD_GENERATED against
	/// the url 'http://' + keyword. For example, if you do a tab-to-search
	/// against wikipedia the generated url has a transition qualifer of
	/// TT_KEYWORD, and TemplateURLModel generates a visit for 'wikipedia.org'
	/// with a transition type of TT_KEYWORD_GENERATED. Chrome runtime only.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_KEYWORD = 9
	/// <summary>
	/// Corresponds to a visit generated for a keyword. See description of
	/// TT_KEYWORD for more details. Chrome runtime only.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_KEYWORD_GENERATED = 10
	/// <summary>
	/// General mask defining the bits used for the source values.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_SOURCE_MASK = 0x000000FF

	/// <summary>
	/// Attempted to visit a URL but was blocked.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_BLOCKED_FLAG = 0x00800000
	/// <summary>
	/// Used the Forward or Back function to navigate among browsing history.
	/// Will be ORed to the transition type for the original load.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_FORWARD_BACK_FLAG = 0x01000000
	/// <summary>
	/// Loaded a URL directly via CreateBrowser, LoadURL or LoadRequest.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_DIRECT_LOAD_FLAG = 0x02000000
	/// <summary>
	/// User is navigating to the home page. Chrome runtime only.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_HOME_PAGE_FLAG = 0x04000000
	/// <summary>
	/// The transition originated from an external application; the exact
	/// definition of this is embedder dependent. Chrome runtime and
	/// extension system only.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_FROM_API_FLAG = 0x08000000
	/// <summary>
	/// The beginning of a navigation chain.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_CHAIN_START_FLAG = 0x10000000
	/// <summary>
	/// The last transition in a redirect chain.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_CHAIN_END_FLAG = 0x20000000
	/// <summary>
	/// Redirects caused by JavaScript or a meta refresh tag on the page.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_CLIENT_REDIRECT_FLAG = 0x40000000
	/// <summary>
	/// Redirects sent from the server by HTTP headers.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_SERVER_REDIRECT_FLAG = 0x80000000
	/// <summary>
	/// Used to test whether a transition involves a redirect.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_IS_REDIRECT_MASK = 0xC0000000
	/// <summary>
	/// General mask defining the bits used for the qualifiers.
	/// </summary>
	/// <remarks>
	/// <para>TCefTransitionType value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see></para>
	/// </remarks>
	TT_QUALIFIER_MASK = 0xFFFFFF00

	/// <summary>
	/// Default behavior.
	/// </summary>
	/// <remarks>
	/// <para>TCefUrlRequestFlags value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see></para>
	/// </remarks>
	UR_FLAG_NONE = 0
	/// <summary>
	/// If set the cache will be skipped when handling the request. Setting this
	/// value is equivalent to specifying the "Cache-Control: no-cache" request
	/// header. Setting this value in combination with UR_FLAG_ONLY_FROM_CACHE
	/// will cause the request to fail.
	/// </summary>
	/// <remarks>
	/// <para>TCefUrlRequestFlags value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see></para>
	/// </remarks>
	UR_FLAG_SKIP_CACHE = 1 << 0
	/// <summary>
	/// If set the request will fail if it cannot be served from the cache (or
	/// some equivalent local store). Setting this value is equivalent to
	/// specifying the "Cache-Control: only-if-cached" request header. Setting
	/// this value in combination with UR_FLAG_SKIP_CACHE or UR_FLAG_DISABLE_CACHE
	/// will cause the request to fail.
	/// </summary>
	/// <remarks>
	/// <para>TCefUrlRequestFlags value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see></para>
	/// </remarks>
	UR_FLAG_ONLY_FROM_CACHE = 1 << 1
	/// <summary>
	/// If set the cache will not be used at all. Setting this value is equivalent
	/// to specifying the "Cache-Control: no-store" request header. Setting this
	/// value in combination with UR_FLAG_ONLY_FROM_CACHE will cause the request
	/// to fail.
	/// </summary>
	/// <remarks>
	/// <para>TCefUrlRequestFlags value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see></para>
	/// </remarks>
	UR_FLAG_DISABLE_CACHE = 1 << 2
	/// <summary>
	/// If set user name, password, and cookies may be sent with the request, and
	/// cookies may be saved from the response.
	/// </summary>
	/// <remarks>
	/// <para>TCefUrlRequestFlags value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see></para>
	/// </remarks>
	UR_FLAG_ALLOW_STORED_CREDENTIALS = 1 << 3
	/// <summary>
	/// If set upload progress events will be generated when a request has a body.
	/// </summary>
	/// <remarks>
	/// <para>TCefUrlRequestFlags value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see></para>
	/// </remarks>
	UR_FLAG_REPORT_UPLOAD_PROGRESS = 1 << 4
	/// <summary>
	/// If set the ICefURLRequestClient.OnDownloadData method will not be called.
	/// </summary>
	/// <remarks>
	/// <para>TCefUrlRequestFlags value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see></para>
	/// </remarks>
	UR_FLAG_NO_DOWNLOAD_DATA = 1 << 5
	/// <summary>
	/// If set 5XX redirect errors will be propagated to the observer instead of
	/// automatically re-tried. This currently only applies for requests
	/// originated in the browser process.
	/// </summary>
	/// <remarks>
	/// <para>TCefUrlRequestFlags value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see></para>
	/// </remarks>
	UR_FLAG_NO_RETRY_ON_5XX = 1 << 6
	/// <summary>
	/// If set 3XX responses will cause the fetch to halt immediately rather than
	/// continue through the redirect.
	/// </summary>
	/// <remarks>
	/// <para>TCefUrlRequestFlags value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see></para>
	/// </remarks>
	UR_FLAG_STOP_ON_REDIRECT = 1 << 7

	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_UNKNOWN = 0
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_UI = 1 << 0
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_MOUSE = 1 << 1
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_MUTATION = 1 << 2
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_KEYBOARD = 1 << 3
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_TEXT = 1 << 4
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_COMPOSITION = 1 << 5
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_DRAG = 1 << 6
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_CLIPBOARD = 1 << 7
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_MESSAGE = 1 << 8
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_WHEEL = 1 << 9
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_BEFORE_TEXT_INSERTED = 1 << 10
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_OVERFLOW = 1 << 11
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_PAGE_TRANSITION = 1 << 12
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_POPSTATE = 1 << 13
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_PROGRESS = 1 << 14
	/// <summary>
	/// DOM event category flag.
	/// </summary>
	/// <remarks>
	/// <para>TCefDomEventCategory value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see></para>
	/// </remarks>
	DOM_EVENT_CATEGORY_XMLHTTPREQUEST_PROGRESS = 1 << 15

	/// <summary>
	/// Don't unescape anything at all.
	/// </summary>
	/// <remarks>
	/// <para>TCefUriUnescapeRule values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see></para>
	/// </remarks>
	UU_NONE = 0
	/// <summary>
	/// Don't unescape anything special, but all normal unescaping will happen.
	/// This is a placeholder and can't be combined with other flags (since it's
	/// just the absence of them). All other unescape rules imply "normal" in
	/// addition to their special meaning. Things like escaped letters, digits,
	/// and most symbols will get unescaped with this mode.
	/// </summary>
	/// <remarks>
	/// <para>TCefUriUnescapeRule values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see></para>
	/// </remarks>
	UU_NORMAL = 1 << 0
	/// <summary>
	/// Convert %20 to spaces. In some places where we're showing URLs, we may
	/// want this. In places where the URL may be copied and pasted out, then
	/// you wouldn't want this since it might not be interpreted in one piece
	/// by other applications.
	/// </summary>
	/// <remarks>
	/// <para>TCefUriUnescapeRule values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see></para>
	/// </remarks>
	UU_SPACES = 1 << 1
	/// <summary>
	/// Unescapes '/' and '\\'. If these characters were unescaped, the resulting
	/// URL won't be the same as the source one. Moreover, they are dangerous to
	/// unescape in strings that will be used as file paths or names. This value
	/// should only be used when slashes don't have special meaning, like data
	/// URLs.
	/// </summary>
	/// <remarks>
	/// <para>TCefUriUnescapeRule values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see></para>
	/// </remarks>
	UU_PATH_SEPARATORS = 1 << 2
	/// <summary>
	/// Unescapes various characters that will change the meaning of URLs,
	/// including '%', '+', '&', '#'. Does not unescape path separators.
	/// If these characters were unescaped, the resulting URL won't be the same
	/// as the source one. This flag is used when generating final output like
	/// filenames for URLs where we won't be interpreting as a URL and want to do
	/// as much unescaping as possible.
	/// </summary>
	/// <remarks>
	/// <para>TCefUriUnescapeRule values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see></para>
	/// </remarks>
	UU_URL_SPECIAL_CHARS_EXCEPT_PATH_SEPARATORS = 1 << 3
	/// <summary>
	/// URL queries use "+" for space. This flag controls that replacement.
	/// </summary>
	/// <remarks>
	/// <para>TCefUriUnescapeRule values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see></para>
	/// </remarks>
	UU_REPLACE_PLUS_WITH_SPACE = 1 << 4

	/// <summary>
	/// Supported SSL version values.
	/// </summary>
	/// <remarks>
	/// <para>TCefSSLVersion values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_ssl_version_t)</see></para>
	/// <para><see href="https://source.chromium.org/chromium/chromium/src/+/main:net/ssl/ssl_connection_status_flags.h">See net/ssl/ssl_connection_status_flags.h for more information.</see></para>
	/// </remarks>
	SSL_CONNECTION_VERSION_UNKNOWN = 0
	SSL_CONNECTION_VERSION_SSL2    = 1
	SSL_CONNECTION_VERSION_SSL3    = 2
	SSL_CONNECTION_VERSION_TLS1    = 3
	SSL_CONNECTION_VERSION_TLS1_1  = 4
	SSL_CONNECTION_VERSION_TLS1_2  = 5
	SSL_CONNECTION_VERSION_TLS1_3  = 6
	SSL_CONNECTION_VERSION_QUIC    = 7

	/// <summary>
	/// Supported SSL content status flags. See content/public/common/ssl_status.h
	/// for more information.
	/// </summary>
	/// <remarks>
	/// <para>TCefSSLContentStatus values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_ssl_content_status_t)</see></para>
	/// </remarks>
	SSL_CONTENT_NORMAL_CONTENT             = 0
	SSL_CONTENT_DISPLAYED_INSECURE_CONTENT = 1 << 0
	SSL_CONTENT_RAN_INSECURE_CONTENT       = 1 << 1

	/// <summary>
	/// Default behavior.
	/// </summary>
	/// <remarks>
	/// <para>TCefJsonWriterOptions value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_writer_options_t)</see></para>
	/// </remarks>
	JSON_WRITER_DEFAULT = 0
	/// <summary>
	/// This option instructs the writer that if a Binary value is encountered,
	/// the value (and key if within a dictionary) will be omitted from the
	/// output, and success will be returned. Otherwise, if a binary value is
	/// encountered, failure will be returned.
	/// </summary>
	/// <remarks>
	/// <para>TCefJsonWriterOptions value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_writer_options_t)</see></para>
	/// </remarks>
	JSON_WRITER_OMIT_BINARY_VALUES = 1 << 0
	/// <summary>
	/// This option instructs the writer to write doubles that have no fractional
	/// part as a normal integer (i.e., without using exponential notation
	/// or appending a '.0') as long as the value is within the range of a
	/// 64-bit int.
	/// </summary>
	/// <remarks>
	/// <para>TCefJsonWriterOptions value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_writer_options_t)</see></para>
	/// </remarks>
	JSON_WRITER_OMIT_DOUBLE_TYPE_PRESERVATION = 1 << 1
	/// <summary>
	/// Return a slightly nicer formatted json string (pads with whitespace to
	/// help with readability).
	/// </summary>
	/// <remarks>
	/// <para>TCefJsonWriterOptions value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_writer_options_t)</see></para>
	/// </remarks>
	JSON_WRITER_PRETTY_PRINT = 1 << 2

	/// <summary>
	/// Print job duplex mode values.
	/// </summary>
	/// <remarks>
	/// <para>TCefDuplexMode values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_duplex_mode_t)</see></para>
	/// </remarks>
	DUPLEX_MODE_UNKNOWN    = -1
	DUPLEX_MODE_SIMPLEX    = 0
	DUPLEX_MODE_LONG_EDGE  = 1
	DUPLEX_MODE_SHORT_EDGE = 2

	/// <summary>
	/// Result codes for ICefMediaRouter.CreateRoute. Should be kept in sync with
	/// Chromium's media_router::mojom::RouteRequestResultCode type.
	/// </summary>
	/// <remarks>
	/// <para>TCefMediaRouterCreateResult values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_route_create_result_t)</see></para>
	/// </remarks>
	CEF_MRCR_UNKNOWN_ERROR            = 0
	CEF_MRCR_OK                       = 1
	CEF_MRCR_TIMED_OUT                = 2
	CEF_MRCR_ROUTE_NOT_FOUND          = 3
	CEF_MRCR_SINK_NOT_FOUND           = 4
	CEF_MRCR_INVALID_ORIGIN           = 5
	CEF_MRCR_NO_SUPPORTED_PROVIDER    = 7
	CEF_MRCR_CANCELLED                = 8
	CEF_MRCR_ROUTE_ALREADY_EXISTS     = 9
	CEF_MRCR_ROUTE_ALREADY_TERMINATED = 11

	/// <summary>
	/// Supported quick menu state bit flags.
	/// </summary>
	/// <remarks>
	/// <para>TCefQuickMenuEditStateFlags values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_quick_menu_edit_state_flags_t)</see></para>
	/// </remarks>
	QM_EDITFLAG_NONE         = 0
	QM_EDITFLAG_CAN_ELLIPSIS = 1 << 0
	QM_EDITFLAG_CAN_CUT      = 1 << 1
	QM_EDITFLAG_CAN_COPY     = 1 << 2
	QM_EDITFLAG_CAN_PASTE    = 1 << 3

	/// <summary>
	/// Values indicating what state of the touch handle is set.
	/// </summary>
	/// <remarks>
	/// <para>TCefTouchHandleStateFlags values.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_touch_handle_state_flags_t)</see></para>
	/// </remarks>
	CEF_THS_FLAG_NONE        = 0
	CEF_THS_FLAG_ENABLED     = 1 << 0
	CEF_THS_FLAG_ORIENTATION = 1 << 1
	CEF_THS_FLAG_ORIGIN      = 1 << 2
	CEF_THS_FLAG_ALPHA       = 1 << 3

	/// <summary>
	/// No permission.
	/// </summary>
	/// <remarks>
	/// <para>TCefMediaAccessPermissionTypes value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t)</see></para>
	/// </remarks>
	CEF_MEDIA_PERMISSION_NONE = 0
	/// <summary>
	/// Device audio capture permission.
	/// </summary>
	/// <remarks>
	/// <para>TCefMediaAccessPermissionTypes value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t)</see></para>
	/// </remarks>
	CEF_MEDIA_PERMISSION_DEVICE_AUDIO_CAPTURE = 1 << 0
	/// <summary>
	/// Device video capture permission.
	/// </summary>
	/// <remarks>
	/// <para>TCefMediaAccessPermissionTypes value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t)</see></para>
	/// </remarks>
	CEF_MEDIA_PERMISSION_DEVICE_VIDEO_CAPTURE = 1 << 1
	/// <summary>
	/// Desktop audio capture permission.
	/// </summary>
	/// <remarks>
	/// <para>TCefMediaAccessPermissionTypes value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t)</see></para>
	/// </remarks>
	CEF_MEDIA_PERMISSION_DESKTOP_AUDIO_CAPTURE = 1 << 2
	/// <summary>
	/// Desktop video capture permission.
	/// </summary>
	/// <remarks>
	/// <para>TCefMediaAccessPermissionTypes value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t)</see></para>
	/// </remarks>
	CEF_MEDIA_PERMISSION_DESKTOP_VIDEO_CAPTURE = 1 << 3

	/// <summary>
	/// Platform API hash.
	/// </summary>
	/// <remarks>
	/// <para>ucef_api_hash parameter.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/cef_api_hash.h">CEF source file: /include/cef_api_hash.h</see></para>
	/// </remarks>
	CEF_API_HASH_PLATFORM = 0
	/// <summary>
	/// Universal API hash.
	/// </summary>
	/// <remarks>
	/// <para>ucef_api_hash parameter.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/cef_api_hash.h">CEF source file: /include/cef_api_hash.h</see></para>
	/// </remarks>
	CEF_API_HASH_UNIVERSAL = 1
	/// <summary>
	/// Commit hash.
	/// </summary>
	/// <remarks>
	/// <para>ucef_api_hash parameter.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/cef_api_hash.h">CEF source file: /include/cef_api_hash.h</see></para>
	/// </remarks>
	CEF_COMMIT_HASH = 2

	/// <summary>
	/// No interrupt reason.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_NONE = 0
	/// <summary>
	/// Generic file operation failure.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_FAILED = 1
	/// <summary>
	/// The file cannot be accessed due to security restrictions.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_ACCESS_DENIED = 2
	/// <summary>
	/// There is not enough room on the drive.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_NO_SPACE = 3
	/// <summary>
	/// The directory or file name is too long.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_NAME_TOO_LONG = 5
	/// <summary>
	/// The file is too large for the file system to handle.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_LARGE = 6
	/// <summary>
	/// The file contains a virus.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_VIRUS_INFECTED = 7
	/// <summary>
	/// The file was in use. Too many files are opened at once. We have run out of
	/// memory.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_TRANSIENT_ERROR = 10
	/// <summary>
	/// The file was blocked due to local policy.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_BLOCKED = 11
	/// <summary>
	/// An attempt to check the safety of the download failed due to unexpected
	/// reasons. See http://crbug.com/153212.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_SECURITY_CHECK_FAILED = 12
	/// <summary>
	/// An attempt was made to seek past the end of a file in opening
	/// a file (as part of resuming a previously interrupted download).
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_SHORT = 13
	/// <summary>
	/// The partial file didn't match the expected hash.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_HASH_MISMATCH = 14
	/// <summary>
	/// The source and the target of the download were the same.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_SAME_AS_SOURCE = 15
	/// <summary>
	/// Generic network failure.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_NETWORK_FAILED = 20
	/// <summary>
	/// The network operation timed out.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_NETWORK_TIMEOUT = 21
	/// <summary>
	/// The network connection has been lost.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_NETWORK_DISCONNECTED = 22
	/// <summary>
	/// The server has gone down.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_NETWORK_SERVER_DOWN = 23
	/// <summary>
	/// The network request was invalid. This may be due to the original URL or a
	/// redirected URL:
	/// - Having an unsupported scheme.
	/// - Being an invalid URL.
	/// - Being disallowed by policy.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_NETWORK_INVALID_REQUEST = 24
	/// <summary>
	/// The server indicates that the operation has failed (generic).
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_FAILED = 30
	/// <summary>
	/// The server does not support range requests.
	/// Internal use only:  must restart from the beginning.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_NO_RANGE = 31
	/// <summary>
	/// The server does not have the requested data.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_BAD_CONTENT = 33
	/// <summary>
	/// Server didn't authorize access to resource.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_UNAUTHORIZED = 34
	/// <summary>
	/// Server certificate problem.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_CERT_PROBLEM = 35
	/// <summary>
	/// Server access forbidden.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_FORBIDDEN = 36
	/// <summary>
	/// Unexpected server response. This might indicate that the responding server
	/// may not be the intended server.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_UNREACHABLE = 37
	/// <summary>
	/// The server sent fewer bytes than the content-length header. It may
	/// indicate that the connection was closed prematurely, or the Content-Length
	/// header was invalid. The download is only interrupted if strong validators
	/// are present. Otherwise, it is treated as finished.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_CONTENT_LENGTH_MISMATCH = 38
	/// <summary>
	/// An unexpected cross-origin redirect happened.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_CROSS_ORIGIN_REDIRECT = 39
	/// <summary>
	/// The user canceled the download.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_USER_CANCELED = 40
	/// <summary>
	/// The user shut down the browser.
	/// Internal use only:  resume pending downloads if possible.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_USER_SHUTDOWN = 41
	/// <summary>
	/// The browser crashed.
	/// Internal use only:  resume pending downloads if possible.
	/// </summary>
	/// <remarks>
	/// <para>TCefDownloadInterruptReason value.</para>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
	/// </remarks>
	CEF_DOWNLOAD_INTERRUPT_REASON_CRASH = 50

	ABOUTBLANK_URI      = "about:blank"
	DEVTOOLS_WINDOWNAME = "DevTools"

	/// <summary>
	/// Direct proxy type: Never use a proxy.
	/// </summary>
	CEF_PROXYTYPE_DIRECT = 0
	/// <summary>
	/// Auto_detect proxy type: Auto detect proxy settings.
	/// </summary>
	CEF_PROXYTYPE_AUTODETECT = 1
	/// <summary>
	/// System proxy type: Use system proxy settings.
	/// </summary>
	CEF_PROXYTYPE_SYSTEM = 2
	/// <summary>
	/// Fixed_servers proxy type: Use fixed proxy servers.
	/// </summary>
	CEF_PROXYTYPE_FIXED_SERVERS = 3
	/// <summary>
	/// Pac_script proxy type: Use a .pac proxy script.
	/// </summary>
	CEF_PROXYTYPE_PAC_SCRIPT = 4

	/// <summary>
	/// Used in the severity parameter in the 'cef_log' function, also known as 'CefLog' in CEF4Delphi.
	/// The log severities are used to index into the array of names, see log_severity_names.
	/// /base/allocator/partition_allocator/partition_alloc_base/logging.h
	/// /base/logging.cc
	/// </summary>
	/// <remarks>
	/// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see></para>
	/// </remarks>
	CEF_LOG_SEVERITY_VERBOSE = -1
	CEF_LOG_SEVERITY_INFO    = 0
	CEF_LOG_SEVERITY_WARNING = 1
	CEF_LOG_SEVERITY_ERROR   = 2
	/// <summary>This severity log level causes a crash.</summary>
	CEF_LOG_SEVERITY_FATAL                      = 3
	CEF_MAX_CONNECTIONS_PER_PROXY_DEFAULT_VALUE = 32
	CEF_MAX_CONNECTIONS_PER_PROXY_MIN_VALUE     = 7
	CEF_MAX_CONNECTIONS_PER_PROXY_MAX_VALUE     = 99
	CEF_COOKIE_PREF_DEFAULT                     = 0
	CEF_COOKIE_PREF_ALLOW                       = 1
	CEF_COOKIE_PREF_BLOCK                       = 2
	CEF_DEFAULT_ENABLEFOCUSDELAY                = 500 // delay in ms to enable the browser focus
)

// / Event type used by TChromiumCore.SimulateKeyEvent
type TSimulatedCefKeyEventType = int32

const (
	KetKeyDown TSimulatedCefKeyEventType = iota
	KetKeyUp
	KetRawKeyDown
	KetChar
)

// Modifier values used in the Input.dispatchTouchEvent and Input.dispatchMouseEvent DevTools methods.
// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchTouchEvent
type TCefMousetouchEventModifiers = int32

const (
	CEF_MOUSETOUCH_EVENT_MODIFIERS_NONE    TCefMousetouchEventModifiers = 0
	CEF_MOUSETOUCH_EVENT_MODIFIERS_ALT     TCefMousetouchEventModifiers = 1 << 0
	CEF_MOUSETOUCH_EVENT_MODIFIERS_CTRL    TCefMousetouchEventModifiers = 1 << 1
	CEF_MOUSETOUCH_EVENT_MODIFIERS_METACMD TCefMousetouchEventModifiers = 1 << 2
	CEF_MOUSETOUCH_EVENT_MODIFIERS_SHIFT   TCefMousetouchEventModifiers = 1 << 3
)

// / Key location value used in the TChromiumCore.dispatchKeyEvent DevTools method.
type TCefKeyLocation = int32

const (
	CEF_KEYLOCATION_NONE TCefKeyLocation = iota
	CEF_KEYLOCATION_LEFT
	CEF_KEYLOCATION_RIGHT
)

// / Mouse button in the TChromiumCore.SimulateMouseEvent function.
// / <para><see href="https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent">See the Input.dispatchMouseEvent DevTools method</see></para>
type TCefSimulatedMouseButton = int32

const (
	CEF_SIMULATEDMOUSEBUTTON_NONE TCefSimulatedMouseButton = iota
	CEF_SIMULATEDMOUSEBUTTON_LEFT
	CEF_SIMULATEDMOUSEBUTTON_MIDDLE
	CEF_SIMULATEDMOUSEBUTTON_RIGHT
	CEF_SIMULATEDMOUSEBUTTON_BACK
	CEF_SIMULATEDMOUSEBUTTON_FORWARD
)

// Modifier values used in the Input.dispatchMouseEvent DevTools method.
// A number indicating which buttons are pressed on the mouse when a mouse event is triggered. Left=1, Right=2, Middle=4, Back=8, Forward=16, None=0.
// https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent
type TCefPressedMouseButtons = int32

const (
	CEF_PRESSED_MOUSE_BUTTONS_NONE    TCefPressedMouseButtons = 0
	CEF_PRESSED_MOUSE_BUTTONS_LEFT    TCefPressedMouseButtons = 1 << 0
	CEF_PRESSED_MOUSE_BUTTONS_RIGHT   TCefPressedMouseButtons = 1 << 1
	CEF_PRESSED_MOUSE_BUTTONS_MIDDLE  TCefPressedMouseButtons = 1 << 2
	CEF_PRESSED_MOUSE_BUTTONS_BACK    TCefPressedMouseButtons = 1 << 3
	CEF_PRESSED_MOUSE_BUTTONS_FORWARD TCefPressedMouseButtons = 1 << 4
)

// / Pointer type in the TChromiumCore.SimulateMouseEvent function.
// / <para><see href="https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent">See the Input.dispatchMouseEvent DevTools method</see></para>
type TCefSimulatedPointerType = int32

const (
	CEF_SIMULATEDPOINTERTYPE_MOUSE TCefSimulatedPointerType = iota
	CEF_SIMULATEDPOINTERTYPE_PEN
)

// / Type of touch event in the TChromiumCore.SimulateTouchEvent function.
// / <para><see href="https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchTouchEvent">See the Input.dispatchTouchEvent DevTools method</see></para>
type TCefSimulatedTouchEventType = int32

const (
	TouchStart TCefSimulatedTouchEventType = iota
	TouchEnd
	TouchMove
	TouchCancel
)

// / <summary>
// / Storage types used by the Storage.clearDataForOrigin DevTools method in TChromiumCore.ClearDataForOrigin.
// / </summary>
type TCefClearDataStorageTypes = int32

const (
	CdstAppCache TCefClearDataStorageTypes = iota
	CdstCookies
	CdstFileSystems
	CdstIndexeddb
	CdstLocalStorage
	CdstShaderCache
	CdstWebsql
	CdstServiceWorkers
	CdstCacheStorage
	CdstAll // default
)

// / Blink editing commands used by the "Input.dispatchKeyEvent" DevTools method.
// / <para><see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</see></para>
// / <para><see href="https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/editing/commands/editor_command_names.h">See the Chromium sources.</see></para>
type TCefEditingCommand = int32

const (
	EcNone TCefEditingCommand = iota
	EcAlignCenter
	EcAlignJustified
	EcAlignLeft
	EcAlignRight
	EcBackColor
	EcBackwardDelete
	EcBold
	EcCopy
	EcCreateLink
	EcCut
	EcDefaultParagraphSeparator
	EcDelete
	EcDeleteBackward
	EcDeleteBackwardByDecomposingPreviousCharacter
	EcDeleteForward
	EcDeleteToBeginningOfLine
	EcDeleteToBeginningOfParagraph
	EcDeleteToEndOfLine
	EcDeleteToEndOfParagraph
	EcDeleteToMark
	EcDeleteWordBackward
	EcDeleteWordForward
	EcFindString
	EcFontName
	EcFontSize
	EcFontSizeDelta
	EcForeColor
	EcFormatBlock
	EcForwardDelete
	EcHiliteColor
	EcIgnoreSpelling
	EcIndent
	EcInsertBacktab
	EcInsertHorizontalRule
	EcInsertHTML
	EcInsertImage
	EcInsertLineBreak
	EcInsertNewline
	EcInsertNewlineInQuotedContent
	EcInsertOrderedList
	EcInsertParagraph
	EcInsertTab
	EcInsertText
	EcInsertUnorderedList
	EcItalic
	EcJustifyCenter
	EcJustifyFull
	EcJustifyLeft
	EcJustifyNone
	EcJustifyRight
	EcMakeTextWritingDirectionLeftToRight
	EcMakeTextWritingDirectionNatural
	EcMakeTextWritingDirectionRightToLeft
	EcMoveBackward
	EcMoveBackwardAndModifySelection
	EcMoveDown
	EcMoveDownAndModifySelection
	EcMoveForward
	EcMoveForwardAndModifySelection
	EcMoveLeft
	EcMoveLeftAndModifySelection
	EcMovePageDown
	EcMovePageDownAndModifySelection
	EcMovePageUp
	EcMovePageUpAndModifySelection
	EcMoveParagraphBackward
	EcMoveParagraphBackwardAndModifySelection
	EcMoveParagraphForward
	EcMoveParagraphForwardAndModifySelection
	EcMoveRight
	EcMoveRightAndModifySelection
	EcMoveToBeginningOfDocument
	EcMoveToBeginningOfDocumentAndModifySelection
	EcMoveToBeginningOfLine
	EcMoveToBeginningOfLineAndModifySelection
	EcMoveToBeginningOfParagraph
	EcMoveToBeginningOfParagraphAndModifySelection
	EcMoveToBeginningOfSentence
	EcMoveToBeginningOfSentenceAndModifySelection
	EcMoveToEndOfDocument
	EcMoveToEndOfDocumentAndModifySelection
	EcMoveToEndOfLine
	EcMoveToEndOfLineAndModifySelection
	EcMoveToEndOfParagraph
	EcMoveToEndOfParagraphAndModifySelection
	EcMoveToEndOfSentence
	EcMoveToEndOfSentenceAndModifySelection
	EcMoveToLeftEndOfLine
	EcMoveToLeftEndOfLineAndModifySelection
	EcMoveToRightEndOfLine
	EcMoveToRightEndOfLineAndModifySelection
	EcMoveUp
	EcMoveUpAndModifySelection
	EcMoveWordBackward
	EcMoveWordBackwardAndModifySelection
	EcMoveWordForward
	EcMoveWordForwardAndModifySelection
	EcMoveWordLeft
	EcMoveWordLeftAndModifySelection
	EcMoveWordRight
	EcMoveWordRightAndModifySelection
	EcOutdent
	EcOverWrite
	EcPaste
	EcPasteAndMatchStyle
	EcPasteGlobalSelection
	EcPrint
	EcRedo
	EcRemoveFormat
	EcScrollLineDown
	EcScrollLineUp
	EcScrollPageBackward
	EcScrollPageForward
	EcScrollToBeginningOfDocument
	EcScrollToEndOfDocument
	EcSelectAll
	EcSelectLine
	EcSelectParagraph
	EcSelectSentence
	EcSelectToMark
	EcSelectWord
	EcSetMark
	EcStrikethrough
	EcStyleWithCSS
	EcSubscript
	EcSuperscript
	EcSwapWithMark
	EcToggleBold
	EcToggleItalic
	EcToggleUnderline
	EcTranspose
	EcUnderline
	EcUndo
	EcUnlink
	EcUnscript
	EcUnselect
	EcUseCSS
	EcYank
	EcYankAndSelect
)

// / Type of mouse event in the TChromiumCore.SimulateMouseEvent function.
// / <para><see href="https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent">See the Input.dispatchMouseEvent DevTools method</see></para>
type TCefSimulatedMouseEventType = int32

const (
	MousePressed TCefSimulatedMouseEventType = iota
	MouseReleased
	MouseMoved
	MouseWheel
)
