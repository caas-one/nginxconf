package core

import (
	"fmt"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

// UpstreamServerParameters struct
type UpstreamServerParameters struct {
	Weight      string `json:"weight,omitempty" key:"weight" eg:"weight=99"`                   // weight=number
	MaxConns    string `json:"maxConns,omitempty" key:"max_conns" eg:"max_conns=3"`            // max_conns=number
	MaxFails    string `json:"maxFails,omitempty" key:"max_fails" eg:"max_fails=3"`            // max_fails=number
	FailTimeout string `json:"failTimeout,omitempty" key:"fail_timeout" eg:"fail_timeout=30s"` // fail_timeout=time (eg. 30s)
	Backup      string `json:"backup,omitempty" key:"backup" eg:"backup"`                      // backup
	Down        string `json:"down,omitempty" key:"down" eg:"down"`                            // down
}

// UpstreamKeepalive struct
type UpstreamKeepalive struct {
	Time string `json:"time,omitempty" key:"time" eg:"16"`
}

// UpstreamServer struct
type UpstreamServer struct {
	Address string                   `json:"address" key:"server" eg:"10.10.12.45:80"` // eg: server 10.10.12.45:80/app.example.com:80/unix:/tmp/backend3
	Params  UpstreamServerParameters `json:"params,omitempty"`
}

// Upstream struct
type Upstream struct {
	Name              string            `json:"name,omitempty" key:"upstream" eg:"upstream example_app_airsupport_http"` // example_app_airsupport_http
	KeepaliveTimeout  string            `json:"keepaliveTimeout,omitempty" key:"keepalive_timeout" eg:"keepalive_timeout 20s"`
	Method            string            `json:"method,omitempty" key:"ip_hash" eg:"ip_hash"`               // ip_hash
	UpstreamKeepalive UpstreamKeepalive `json:"keepalive,omitempty" key:"keepalive" eg:"keepalive 16"`     // keepalive 16
	Servers           []UpstreamServer  `json:"servers,omitempty" key:"server" eg:"server 10.10.12.45:80"` // server 192.168.240.229:80 weight=99 max_fails=2 fail_timeout=30s
}

// NewUpstream func
func NewUpstream() *Upstream {
	return &Upstream{Servers: make([]UpstreamServer, 0)}
}

// LocationIfBlock struct
type LocationIfBlock struct {
	Condition string   `json:"condition,omitempty" key:"if" eg:"( $request_uri ~* /app-merchant-proxy/ )"`                   // ( $request_uri ~* /app-merchant-proxy/ )
	ProxyPass string   `json:"proxyPass,omitempty" key:"proxy_pass" eg:"proxy_pass http://example_app_airsupport_http"`      // proxy_pass http://example_app_airsupport_http;
	SendFile  string   `json:"sendFile,omitempty" key:"send_file" eg:"send_file on"`                                         // sendfile on;
	Root      string   `json:"root,omitempty" key:"root" eg:"root /data/w3"`                                                 // root /data/w3;
	Rewrite   []string `json:"rewrite,omitempty" key:"rewrite" eg:"rewrite ^/docs/(.*)$ /yop_doc/doc/$1/default.html break"` // rewrite ^/docs/(.*)$ /yop_doc/doc/$1/default.html break;
	Set       []string `json:"set,omitempty" key:"set" eg:"set $limit_rate 4k"`                                              // set $limit_rate 4k;
	AddHeader []string `json:"addHeader,omitempty" key:"add_header" eg:"add_header X-Frame-Options SAMEORIGIN"`
}

// NewLocationIfBlock func
func NewLocationIfBlock() *LocationIfBlock {
	ifBlock := &LocationIfBlock{}
	ifBlock.Rewrite = make([]string, 0)
	ifBlock.Set = make([]string, 0)
	ifBlock.AddHeader = make([]string, 0)
	return ifBlock
}

// Location struct
type Location struct {
	Path                       string            `json:"path" key:"location" eg:"location /video/"`                                                                   // location /video/ { ... }
	Aio                        string            `json:"aio,omitempty" key:"aio" eg:"aio off"`                                                                        // aio off; on | off | threads[=pool];
	AioWrite                   string            `json:"aioWrite,omitempty" key:"aio_write" eg:"aio_write off"`                                                       // aio_write off;
	SendFile                   string            `json:"sendFile,omitempty" key:"sendfile" eg:"sendfile on"`                                                          // sendfile on;
	Root                       string            `json:"root,omitempty" key:"root" eg:"root /data/w3"`                                                                // root /data/w3;
	Alias                      string            `json:"alias,omitempty" key:"alias" eg:"alias /data/w3/images/"`                                                     // alias /data/w3/images/;
	ChunkedTransferEncoding    string            `json:"chunkedTransferEncoding,omitempty" key:"chunked_transfer_encoding" eg:"chunked_transfer_encoding on"`         // chunked_transfer_encoding on;
	ClientBodyBufferSize       string            `json:"clientBodyBufferSize,omitempty" key:"client_body_buffer_size" eg:"client_body_buffer_size 8k"`                // client_body_buffer_size 8k|16k;
	ClientBodyInFileOnly       string            `json:"clientBodyInFileOnly,omitempty" key:"client_body_in_file_only" eg:"client_body_in_file_only off"`             // client_body_in_file_only off;
	ClientBodyInSingleBuffer   string            `json:"clientBodyInSingleBuffer,omitempty" key:"client_body_in_single_buffer" eg:"client_body_in_single_buffer off"` // client_body_in_single_buffer off;
	ClientBodyTempPath         string            `json:"clientBodyTempPath,omitempty" key:"client_body_temp_path" eg:"client_body_temp_path client_body_temp"`        // client_body_temp_path client_body_temp;
	ClientBodyTimeout          string            `json:"clientBodyTimeout,omitempty" key:"client_body_timeout" eg:"client_body_timeout 60s"`                          // client_body_timeout 60s;
	DefaultType                string            `json:"defaultType,omitempty" key:"default_type" eg:"default_type text/plain"`                                       // default_type text/plain;
	DirectIO                   string            `json:"directio,omitempty" key:"directio" eg:"directio 512"`                                                         // directio 512;
	LimitConn                  string            `json:"limitConn,omitempty" key:"limit_conn" eg:"limit_conn conn_zone 1"`
	LimitReq                   string            `json:"limitReq,omitempty" key:"limit_req" eg:"limit_req  zone=qps1 burst=5"`
	DirectioAlignment          string            `json:"directioAlignment,omitempty" key:"directio_alignment" eg:"directio_alignment 512"`                                      // directio_alignment 512;
	DisableSymlinks            string            `json:"disableSymlinks,omitempty" key:"disable_symlinks" eg:"disable_symlinks off"`                                            // disable_symlinks off;
	ErrorPage                  string            `json:"errorPage,omitempty" key:"error_page" eg:"error_page 500 502 503 504 /50x.html"`                                        // error_page 500 502 503 504 /50x.html;
	Internal                   string            `json:"internal,omitempty" key:"internal" eg:"internal"`                                                                       // internal;
	KeepaliveDisable           string            `json:"keepaliveDisable,omitempty" key:"keepalive_disable" eg:"keepalive_disable msie6"`                                       // keepalive_disable msie6;
	KeepaliveRequests          string            `json:"keepaliveRequests,omitempty" key:"keepalive_requests" eg:"keepalive_requests 100"`                                      // keepalive_requests 100;
	KeepaliveTimeout           string            `json:"keepaliveTimeout,omitempty" key:"keepalive_timeout" eg:"keepalive_timeout 75s"`                                         // keepalive_timeout 75s;
	LimitRate                  string            `json:"limitRate,omitempty" key:"limit_rate" eg:"limit_rate 0"`                                                                // limit_rate 0;
	LimitRateAfter             string            `json:"limitRateAfter,omitempty" key:"limit_rate_after" eg:"limit_rate_after 0"`                                               // limit_rate_after 0;
	LingeringClose             string            `json:"lingeringClose,omitempty" key:"lingering_close" eg:"lingering_close on"`                                                // lingering_close on;
	LingeringTime              string            `json:"lingeringTime,omitempty" key:"lingering_time" eg:"lingering_time 30s"`                                                  // lingering_time 30s;
	LingeringTimeout           string            `json:"lingeringTimeout,omitempty" key:"lingering_timeout" eg:"lingering_timeout 5s"`                                          // lingering_timeout 5s;
	LogNotFound                string            `json:"logNotFound,omitempty" key:"log_not_found" eg:"log_not_found on"`                                                       // log_not_found on;
	LogSubrequest              string            `json:"logSubrequest,omitempty" key:"log_subrequest" eg:"log_subrequest off"`                                                  // log_subrequest off;
	MaxRanges                  string            `json:"maxRanges,omitempty" key:"max_ranges" eg:"max_ranges 50"`                                                               // max_ranges number;
	MsiePadding                string            `json:"msiePadding,omitempty" key:"msie_padding" eg:"msie_padding on"`                                                         // msie_padding on;
	MsieRefresh                string            `json:"msieRefresh,omitempty" key:"msie_refresh" eg:"msie_refresh off"`                                                        // msie_refresh off;
	OpenFileCache              string            `json:"openFileCache,omitempty" key:"open_file_cache" eg:"open_file_cache off"`                                                // open_file_cache off;
	OpenFileCacheErrors        string            `json:"openFileCacheErrors,omitempty" key:"open_file_cache_errors" eg:"open_file_cache_errors off"`                            // open_file_cache_errors off;
	OpenFileCacheMinUses       string            `json:"openFileCacheMinUses,omitempty" key:"open_file_cache_min_uses" eg:"open_file_cache_min_uses 1"`                         // open_file_cache_min_uses 1;
	OpenFileCacheValid         string            `json:"openFileCacheValid,omitempty" key:"open_file_cache_valid" eg:"open_file_cache_valid 60s"`                               // open_file_cache_valid 60s;
	OutputBuffers              string            `json:"outputBuffers,omitempty" key:"output_buffers" eg:"output_buffers 2 32k"`                                                // output_buffers 2 32k;
	PortInRedirect             string            `json:"portInRedirect,omitempty" key:"port_in_redirect" eg:"port_in_redirect on"`                                              // port_in_redirect on;
	PostponeOutput             string            `json:"postponeOutput,omitempty" key:"postpone_output" eg:"postpone_output 1460"`                                              // postpone_output 1460;
	ReadAhead                  string            `json:"readAhead,omitempty" key:"read_ahead" eg:"read_ahead 0"`                                                                // read_ahead 0;
	RecursiveErrorPages        string            `json:"recursiveErrorPages,omitempty" key:"recursive_error_pages" eg:"recursive_error_pages off"`                              // recursive_error_pages off;
	Resolver                   string            `json:"resolver,omitempty" key:"resolver" eg:"resolver 127.0.0.1 [::1]:5353"`                                                  // resolver 127.0.0.1 [::1]:5353;
	ResolverTimeout            string            `json:"resolverTimeout,omitempty" key:"resolver_timeout" eg:"resolver_timeout 30s"`                                            // resolver_timeout 30s;
	Satisfy                    string            `json:"satisfy,omitempty" key:"satisfy" eg:"satisfy all"`                                                                      // satisfy all;
	SendLowat                  string            `json:"sendLowat,omitempty" key:"send_lowat" eg:"send_lowat 0"`                                                                // send_lowat 0;
	SendTimeout                string            `json:"sendTimeout,omitempty" key:"send_timeout" eg:"send_timeout 60s"`                                                        // send_timeout 60s;
	SendFileMaxChunk           string            `json:"sendFileMaxChunk,omitempty" key:"sendfile_max_chunk" eg:"sendfile_max_chunk 0"`                                         // sendfile_max_chunk 0;
	ServerNameInRedirect       string            `json:"serverNameInRedirect,omitempty" key:"server_name_in_redirect" eg:"server_name_in_redirect off"`                         // server_name_in_redirect off;
	ServerTokens               string            `json:"serverTokens,omitempty" key:"server_tokens" eg:"server_tokens on"`                                                      // server_tokens on;
	SubrequestOutputBufferSize string            `json:"subrequestOutputBufferSize,omitempty" key:"subrequest_output_buffer_size" eg:"subrequest_output_buffer_size 4k"`        // subrequest_output_buffer_size 4k|8k;
	TcpNodelay                 string            `json:"tcpNodelay,omitempty" key:"tcp_nodelay" eg:"tcp_nodelay on"`                                                            // tcp_nodelay on;
	TcpNopush                  string            `json:"tcpNopush,omitempty" key:"tcp_nopush" eg:"tcp_nopush off"`                                                              // tcp_nopush off;
	TryFiles                   string            `json:"tryFiles,omitempty" key:"try_files" eg:"try_files file ... =code"`                                                      // try_files file ... =code;
	TypesHashBucketSize        string            `json:"typesHashBucketSize,omitempty" key:"types_hash_bucket_size" eg:"types_hash_bucket_size 64"`                             // types_hash_bucket_size 64;
	TypesHashMaxSize           string            `json:"typesHashMaxSize,omitempty" key:"types_hash_max_size" eg:"types_hash_max_size 1024"`                                    // types_hash_max_size 1024;
	AbsoluteRedirect           string            `json:"absoluteRedirect,omitempty" key:"absolute_redirect" eg:"absolute_redirect on"`                                          // absolute_redirect on;
	Etag                       string            `json:"etag,omitempty" key:"etag" eg:"etag on"`                                                                                // etag on;
	AccessLog                  string            `json:"accessLog,omitempty" key:"access_log" eg:"access_log  logs/access.log  main"`                                           // access_log
	LogFormat                  []string          `json:"logFormat,omitempty" key:"log_format" eg:"log_format download '$remote_addr - $remote_user [$time_local] '"`            // log_format download '$remote_addr - $remote_user [$time_local] '
	ProxyBind                  string            `json:"proxyBind,omitempty" key:"proxy_bind" eg:"proxy_bind $remote_addr transparent"`                                         // proxy_bind $remote_addr transparent;
	ProxyBufferSize            string            `json:"proxyBufferSize,omitempty" key:"proxy_buffer_size" eg:"proxy_buffer_size 4k"`                                           // proxy_buffer_size 4k|8k;
	ProxyBuffering             string            `json:"proxyBuffering,omitempty" key:"proxy_buffering" eg:"proxy_buffering on"`                                                // proxy_buffering on;
	ProxyBuffers               string            `json:"proxyBuffers,omitempty" key:"proxy_buffers" eg:"proxy_buffers 8 4k"`                                                    // proxy_buffers 8 4k|8k;
	ProxyBusyBuffersSize       string            `json:"ProxyBusyBuffersSize,omitempty" key:"proxy_busy_buffers_size" eg:"proxy_busy_buffers_size 8k"`                          // proxy_busy_buffers_size 8k|16k;
	ProxyCache                 string            `json:"proxyCache,omitempty" key:"proxy_cache" eg:"proxy_cache off"`                                                           // proxy_cache off;
	ProxyCacheBackgroundUpdate string            `json:"proxyCacheBackgroundUpdate,omitempty" key:"proxy_cache_background_update" eg:"proxy_cache_background_update off"`       // proxy_cache_background_update off;
	ProxyCacheBypass           []string          `json:"proxyCacheBypass,omitempty" key:"proxy_cache_bypass" eg:"proxy_cache_bypass  $cookie_nocache $arg_nocache$arg_comment"` // proxy_cache_bypass  $cookie_nocache $arg_nocache$arg_comment;
	ProxyCacheConvertHead      string            `json:"proxyCacheConvertHead,omitempty" key:"proxy_cache_convert_head" eg:"proxy_cache_convert_head on"`                       // proxy_cache_convert_head on;
	ProxyCacheKey              string            `json:"proxyCacheKey,omitempty" key:"proxy_cache_key" eg:"proxy_cache_key $scheme$proxy_host$request_uri"`                     // proxy_cache_key $scheme$proxy_host$request_uri;
	ProxyCacheLock             string            `json:"proxyCacheLock,omitempty" key:"proxy_cache_lock" eg:"proxy_cache_lock off"`                                             // proxy_cache_lock off;
	ProxyCacheLockAge          string            `json:"proxyCacheLockAge,omitempty" key:"proxy_cache_lock_age" eg:"proxy_cache_lock_age 5s"`                                   // proxy_cache_lock_age 5s;
	ProxyCacheLockTimeout      string            `json:"proxyCacheLockTimeout,omitempty" key:"proxy_cache_lock_timeout" eg:"proxy_cache_lock_timeout 5s"`                       // proxy_cache_lock_timeout 5s;
	ProxyCacheMaxRangeOffset   string            `json:"proxyCacheMaxRangeOffset,omitempty" key:"proxy_cache_max_range_offset" eg:"proxy_cache_max_range_offset 1024"`          // proxy_cache_max_range_offset number;
	ProxyCacheMethods          string            `json:"proxyCacheMethods,omitempty" key:"proxy_cache_methods" eg:"proxy_cache_methods GET HEAD"`                               // proxy_cache_methods GET HEAD;
	ProxyCacheMinUses          string            `json:"proxyCacheMinUses,omitempty" key:"proxy_cache_min_uses" eg:"proxy_cache_min_uses 1"`                                    // proxy_cache_min_uses 1;
	ProxyCachePath             string            `json:"proxyCachePath,omitempty" key:"proxy_cache_path" eg:"proxy_cache_path /data/nginx/cache levels=1:2 keys_zone=one:10m"`  // proxy_cache_path /data/nginx/cache levels=1:2 keys_zone=one:10m;
	ProxyCachePurge            string            `json:"proxyCachePurge,omitempty" key:"proxy_cache_purge" eg:"proxy_cache_purge $purge_method"`                                // proxy_cache_purge $purge_method;
	ProxyCacheUseStale         string            `json:"proxyCacheUseStale,omitempty" key:"proxy_cache_use_stale" eg:"proxy_cache_use_stale off"`                               // proxy_cache_use_stale off;
	ProxyCacheValid            []string          `json:"proxyCacheValid,omitempty" key:"proxy_cache_valid" eg:"proxy_cache_valid 200 302 10m"`                                  // proxy_cache_valid 200 302 10m; proxy_cache_valid 404      1m;
	ProxyConnectTimeout        string            `json:"proxyConnectTimeout,omitempty" key:"proxy_connect_timeout" eg:"proxy_connect_timeout 60s"`                              // proxy_connect_timeout 60s;
	ProxyCookieDomain          string            `json:"proxyCookieDomain,omitempty" key:"proxy_cookie_domain" eg:"proxy_cookie_domain off"`                                    // proxy_cookie_domain off;
	ProxyCookiePath            string            `json:"proxyCookiePath,omitempty" key:"proxy_cookie_path" eg:"proxy_cookie_path off"`                                          // proxy_cookie_path off;
	ProxyForceRanges           string            `json:"proxyForceRanges,omitempty" key:"proxy_force_ranges" eg:"proxy_force_ranges off"`                                       // proxy_force_ranges off;
	ProxyHeadersHashBucketSize string            `json:"proxyHeadersHashBucketSize,omitempty" key:"proxy_headers_hash_bucket_size" eg:"proxy_headers_hash_bucket_size 64"`      // proxy_headers_hash_bucket_size 64;
	ProxyHeadersHashMaxSize    string            `json:"proxyHeadersHashMaxSize,omitempty" key:"proxy_headers_hash_max_size" eg:"proxy_headers_hash_max_size 512"`              // proxy_headers_hash_max_size 512;
	ProxyHttpVersion           string            `json:"proxyHttpVersion,omitempty" key:"proxy_http_version" eg:"proxy_http_version 1.0"`                                       // proxy_http_version 1.0;
	ProxyIgnoreClientAbort     string            `json:"proxyIgnoreClientAbort,omitempty" key:"proxy_ignore_client_abort" eg:"proxy_ignore_client_abort off"`                   // proxy_ignore_client_abort off;
	ProxyIgnoreHeaders         string            `json:"proxyIgnoreHeaders,omitempty" key:"proxy_ignore_headers" eg:"proxy_ignore_headers field"`                               // proxy_ignore_headers field ...;
	ProxyInterceptErrors       string            `json:"proxyInterceptErrors,omitempty" key:"proxy_intercept_errors" eg:"proxy_intercept_errors off"`                           // proxy_intercept_errors off;
	ProxyLimitRate             string            `json:"proxyLimitRate,omitempty" key:"proxy_limit_rate" eg:"proxy_limit_rate 0"`                                               // proxy_limit_rate 0;
	ProxyMaxTempFileSize       string            `json:"proxyMaxTempFileSize,omitempty" key:"proxy_max_temp_file_size" eg:"proxy_max_temp_file_size 1024m"`                     // proxy_max_temp_file_size 1024m;
	ProxyMethod                string            `json:"proxyMethod,omitempty" key:"proxy_method" eg:"proxy_method method"`                                                     // proxy_method method;
	ProxyNextUpstream          string            `json:"proxyNextUpstream,omitempty" key:"proxy_next_upstream" eg:"proxy_next_upstream error"`                                  // proxy_next_upstream error | timeout
	ProxyNextUpstreamTimeout   string            `json:"proxyNextUpstreamTimeout,omitempty" key:"proxy_next_upstream_timeout" eg:"proxy_next_upstream_timeout 0"`               // proxy_next_upstream_timeout 0;
	ProxyNextUpstreamTries     string            `json:"proxyNextUpstreamTries,omitempty" key:"proxy_next_upstream_tries" eg:"proxy_next_upstream_tries 0"`                     // proxy_next_upstream_tries 0;
	ProxyNoCache               string            `json:"proxyNoCache,omitempty" key:"proxy_no_cache" eg:"proxy_no_cache string ..."`                                            // proxy_no_cache string ...;
	ProxyPass                  string            `json:"proxyPass,omitempty" key:"proxy_pass" eg:"proxy_pass http://example_app_airsupport_http"`                               // proxy_pass http://example_app_airsupport_http;
	ProxyPassHeader            string            `json:"proxyPassHeader,omitempty" key:"proxy_pass_header" eg:"proxy_pass_header field"`                                        // proxy_pass_header field
	ProxyPassRequestBody       string            `json:"proxyPassRequestBody,omitempty" key:"proxy_pass_request_body" eg:"proxy_pass_request_body on"`                          // proxy_pass_request_body on;
	ProxyPassRequestHeaders    string            `json:"proxyPassRequestHeaders,omitempty" key:"proxy_pass_request_headers" eg:"proxy_pass_request_headers on"`                 // proxy_pass_request_headers on;
	ProxyReadTimeout           string            `json:"proxyReadTimeout,omitempty" key:"proxy_read_timeout" eg:"proxy_read_timeout 60s"`                                       // proxy_read_timeout 60s;
	ProxyRedirect              string            `json:"proxyRedirect,omitempty" key:"proxy_redirect" eg:"proxy_redirect redirect replacement"`                                 // proxy_redirect redirect replacement;
	ProxyRequestBuffering      string            `json:"proxyRequestBuffering,omitempty" key:"proxy_request_buffering" eg:"proxy_request_buffering on"`                         // proxy_request_buffering on;
	ProxySendLowat             string            `json:"proxySendLowat,omitempty" key:"proxy_send_lowat" eg:"proxy_send_lowat 0"`                                               // proxy_send_lowat 0;
	ProxySendTimeout           string            `json:"proxySendTimeout,omitempty" key:"proxy_send_timeout" eg:"proxy_send_timeout 60s"`                                       // proxy_send_timeout 60s;
	ProxySetBody               string            `json:"proxySetBody,omitempty" key:"proxy_set_body" eg:"proxy_set_body value"`                                                 // proxy_set_body value;
	ProxySetHeader             []string          `json:"proxySetHeader,omitempty" key:"proxy_set_header" eg:"proxy_set_header Host $proxy_host"`                                // proxy_set_header Host $proxy_host; proxy_set_header Connection close;
	ProxySocketKeepalive       string            `json:"proxySocketKeepalive,omitempty" key:"proxy_socket_keepalive" eg:"proxy_socket_keepalive off"`                           // proxy_socket_keepalive off;
	ProxySslCertificate        string            `json:"proxySslCertificate,omitempty" key:"proxy_ssl_certificate" eg:"proxy_ssl_certificate file"`                             // proxy_ssl_certificate file;
	ProxySslCertificateKey     string            `json:"proxySslCertificateKey,omitempty" key:"proxy_ssl_certificate_key" eg:"proxy_ssl_certificate_key file"`                  // proxy_ssl_certificate_key file;
	ProxySslCiphers            string            `json:"proxySslCiphers,omitempty" key:"proxy_ssl_ciphers" eg:"proxy_ssl_ciphers DEFAULT"`                                      // proxy_ssl_ciphers DEFAULT;
	ProxySslCrl                string            `json:"proxySslCrl,omitempty" key:"proxy_ssl_crl" eg:"proxy_ssl_crl file"`                                                     // proxy_ssl_crl file;
	ProxySslName               string            `json:"proxySslName,omitempty" key:"proxy_ssl_name" eg:"proxy_ssl_name name"`                                                  // proxy_ssl_name name;
	ProxySslPasswordFile       string            `json:"proxySslPasswordFile,omitempty" key:"proxy_ssl_password_file" eg:"proxy_ssl_password_file file"`                        // proxy_ssl_password_file file;
	ProxySslProtocols          string            `json:"proxySslProtocols,omitempty" key:"proxy_ssl_protocols" eg:"proxy_ssl_protocols TLSv1 TLSv1.1 TLSv1.2"`                  // proxy_ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
	ProxySslServerName         string            `json:"proxySslServerName,omitempty" key:"proxy_ssl_server_name" eg:"proxy_ssl_server_name off"`                               // proxy_ssl_server_name off;
	ProxySslSessionReuse       string            `json:"proxySslSessionReuse,omitempty" key:"proxy_ssl_session_reuse" eg:"proxy_ssl_session_reuse on"`                          // proxy_ssl_session_reuse on;
	ProxySslTrustedCertificate string            `json:"proxySslTrustedCertificate,omitempty" key:"proxy_ssl_trusted_certificate" eg:"proxy_ssl_trusted_certificate file"`      // proxy_ssl_trusted_certificate file;
	ProxySslVerify             string            `json:"proxySslVerify,omitempty" key:"proxy_ssl_verify" eg:"proxy_ssl_verify off"`                                             // proxy_ssl_verify off;
	ProxySslVerifyDepth        string            `json:"proxySslVerifyDepth,omitempty" key:"proxy_ssl_verify_depth" eg:"proxy_ssl_verify_depth 1"`                              // proxy_ssl_verify_depth 1;
	ProxyStore                 string            `json:"proxyStore,omitempty" key:"proxy_store" eg:"proxy_store off"`                                                           // proxy_store off;
	ProxyStoreAccess           string            `json:"proxyStoreAccess,omitempty" key:"proxy_store_access" eg:"proxy_store_access user:rw"`                                   // proxy_store_access user:rw;
	ProxyTempFileWriteSize     string            `json:"proxyTempFileWriteSize,omitempty" key:"proxy_temp_file_write_size" eg:"proxy_temp_file_write_size 8k"`                  // proxy_temp_file_write_size 8k|16k;
	ProxyTempPath              string            `json:"proxyTempPath,omitempty" key:"proxy_temp_path" eg:"proxy_temp_path proxy_temp"`                                         // proxy_temp_path proxy_temp;
	Allow                      []string          `json:"allow,omitempty" key:"allow" eg:"allow 119.161.147.101"`                                                                // allow 119.161.147.101;
	Deny                       []string          `json:"deny,omitempty" key:"deny" eg:"deny all"`                                                                               // deny all;
	IfBlocks                   []LocationIfBlock `json:"ifInLocation,omitempty" key:"ifInLocation" eg:"if ( $request_uri ~* /app-merchant-proxy/ )"`                            // if ( $request_uri ~* /app-merchant-proxy/ ){ proxy_pass http://example_app_airsupport_http; }
	ClientMaxBodySize          string            `json:"clientMaxBodySize,omitempty" key:"client_max_body_size" eg:"client_max_body_size 1m"`                                   // client_max_body_size 1m;
	Expires                    string            `json:"expires,omitempty" key:"expires" eg:"expires 30s"`                                                                      // expires 30s;
	Rewrite                    []string          `json:"rewrite,omitempty" key:"rewrite" eg:"rewrite ^/docs/(.*)$ /yop_doc/doc/$1/default.html break"`                          // rewrite ^/docs/(.*)$ /yop_doc/doc/$1/default.html break;
	StubStatus                 string            `json:"stubStatus,omitempty" key:"stub_status" eg:"stub_status on"`                                                            // stub_status on;
	Set                        []string          `json:"set,omitempty" key:"set" eg:"set $limit_rate 4k"`                                                                       // set $limit_rate 4k;
	AddHeader                  []string          `json:"addHeader,omitempty" key:"add_header" eg:"add_header X-Frame-Options SAMEORIGIN"`
	Gzip                       string            `json:"gzip,omitempty" key:"gzip" eg:"gzip on"`
	GzipMinLength              string            `json:"gzipMinLength" key:"gzip_min_length" eg:"gzip_min_length 20"`
	GzipBuffers                string            `json:"gzipBuffers" key:"gzip_buffers" eg:"gzip_buffers 32 4k|16 8k"`
	GzipCompLevel              string            `json:"gzipCompLevel" key:"gzip_comp_level" eg:"gzip_comp_level 1"`
	GzipTypes                  string            `json:"gzipTypes" key:"gzip_types" eg:"gzip_types text/html"`
	GzipVary                   string            `json:"gzipVary" key:"gzip_vary" eg:"gzip_vary off"`
	GzipDisable                string            `json:"gzipDisable" key:"gzip_disable" eg:"gzip_disable \"msie6\""`
	GzipHttpVersion            string            `json:"gzipHttpVersion" key:"gzip_disable" eg:"gzip_http_version 1.1"`
	GzipProxied                string            `json:"gzipProxied" key:"gzip_proxied" eg:"gzip_proxied off"`
	ProxyHideHeader            string            `json:"proxyHideHeader,omitempty" key:"proxy_hide_header" eg:"proxy_hide_header X-Powered-By"`
}

// NewLocation  location module
func NewLocation() *Location {
	return &Location{
		LogFormat:        make([]string, 0),
		ProxyCacheBypass: make([]string, 0),
		ProxyCacheValid:  make([]string, 0),
		ProxySetHeader:   make([]string, 0),
		IfBlocks:         make([]LocationIfBlock, 0),
		Rewrite:          make([]string, 0),
		Allow:            make([]string, 0),
		Deny:             make([]string, 0),
		Set:              make([]string, 0),
		AddHeader:        make([]string, 0),
	}
}

// ServerIfBlock serverif block
type ServerIfBlock struct {
	Condition      string   `json:"condition,omitempty"  key:"condition" eg:"( $request_uri ~* /app-merchant-proxy/ )"`           // ( $request_uri ~* /app-merchant-proxy/ )
	ProxyPass      string   `json:"proxyPass,omitempty" key:"proxy_pass" eg:"proxy_pass http://example_app_airsupport_http"`      // proxy_pass http://example_app_airsupport_http;
	SendFile       string   `json:"sendFile,omitempty" key:"sendfile" eg:"sendfile on"`                                           // sendfile on;
	Root           string   `json:"root,omitempty" key:"root" eg:"root /data/w3"`                                                 // root /data/w3;
	Rewrite        []string `json:"rewrite,omitempty" key:"rewrite" eg:"rewrite ^/docs/(.*)$ /yop_doc/doc/$1/default.html break"` // rewrite ^/docs/(.*)$ /yop_doc/doc/$1/default.html break;
	Return         string   `json:"return,omitempty" key:"return" eg:"return 403"`                                                // return 403;
	Set            []string `json:"set,omitempty" key:"set" eg:"set $limit_rate 4k"`                                              // set $limit_rate 4k;
	LimitRateAfter string   `json:"limitRateAfter,omitempty" key:"limit_rate_after" eg:"limit_rate_after 0"`                      // limit_rate_after 0;
}

// NewServerIfBlock new serverif
func NewServerIfBlock() *ServerIfBlock {
	return &ServerIfBlock{
		Rewrite: make([]string, 0),
		Set:     make([]string, 0),
	}
}

// Server server module
type Server struct {
	Name                       string          `json:"name" key:"server_name" eg:"server_name www.example.com"`                                            // server_name www.example.com;
	Listen                     string          `json:"listen" key:"listen" eg:"listen 443 ssl"`                                                            // listen 443 ssl;
	Return                     string          `json:"return,omitempty" key:"return" eg:"return 404"`                                                      // return 404
	Root                       string          `json:"root,omitempty" key:"root" eg:"root  html"`                                                          // root  html;
	Ssl                        string          `json:"ssl,omitempty" key:"ssl" eg:"ssl on"`                                                                // ssl on;
	SslCertificate             string          `json:"sslCertificate,omitempty" key:"ssl_certificate" eg:"ssl_certificate     www.example.com.crt"`        // ssl_certificate     www.example.com.crt;
	SslCertificateKey          string          `json:"sslCertificateKey,omitempty" key:"ssl_certificate_key" eg:"ssl_certificate_key www.example.com.key"` // ssl_certificate_key www.example.com.key;
	SslProtocols               string          `json:"sslProtocols,omitempty" key:"ssl_protocols" eg:"ssl_protocols       TLSv1 TLSv1.1 TLSv1.2"`          // ssl_protocols       TLSv1 TLSv1.1 TLSv1.2;
	SslCiphers                 string          `json:"sslCiphers,omitempty" key:"ssl_ciphers" eg:"ssl_ciphers         HIGH:!aNULL:!MD5"`                   // ssl_ciphers         HIGH:!aNULL:!MD5;
	SslSessionTimeout          string          `json:"sslSessionTimeout,omitempty" key:"ssl_session_timeout" eg:"ssl_session_timeout 5m"`                  //  ssl_session_timeout 5m;
	SslPreferServerCiphers     string          `json:"sslPreferServerCiphers,omitempty" key:"ssl_prefer_server_ciphers" eg:"ssl_prefer_server_ciphers on"` // ssl_prefer_server_ciphers on;
	Rewrite                    []string        `json:"rewrite,omitempty" key:"rewrite" eg:"rewrite \"^/(.*)$\" https://app.example.com/$1 permanent"`      // rewrite "^/(.*)$" https://app.example.com/$1 permanent;
	IfBlocks                   []ServerIfBlock `json:"ifInServer,omitempty" key:"ifInServer"`                                                              //
	Locations                  []Location      `json:"locations,omitempty" key:"locations"`
	LimitConn                  string          `json:"limitConn,omitempty" key:"limit_conn" eg:"limit_conn conn_zone 1"`
	LimitReq                   string          `json:"limitReq,omitempty" key:"limit_req" eg:"limit_req  zone=qps1 burst=5"`
	AccessLog                  string          `json:"accessLog,omitempty" key:"access_log" eg:"access_log"`                                                                                       // access_log
	KeepaliveDisable           string          `json:"keepaliveDisable,omitempty" key:"keepalive_disable" eg:"keepalive_disable msie6"`                                                            // keepalive_disable msie6;
	KeepaliveRequests          string          `json:"keepaliveRequests,omitempty" key:"keepalive_requests" eg:"keepalive_requests 100"`                                                           // keepalive_requests 100;
	KeepaliveTimeout           string          `json:"keepaliveTimeout,omitempty" key:"keepalive_timeout" eg:"keepalive_timeout 75s"`                                                              // keepalive_timeout 75s;
	LargeClientHeaderBuffers   string          `json:"largeClientHeaderBuffers,omitempty" key:"large_client_header_buffers" eg:"large_client_header_buffers 4 8k"`                                 // large_client_header_buffers 4 8k;
	LimitRate                  string          `json:"limitRate,omitempty" key:"limit_rate" eg:"limit_rate 0"`                                                                                     // limit_rate 0;
	LimitRateAfter             string          `json:"limitRateAfter,omitempty" key:"limit_rate_after" eg:"limit_rate_after 0"`                                                                    // limit_rate_after 0;
	LingeringClose             string          `json:"lingeringClose,omitempty" key:"lingering_close" eg:"lingering_close on"`                                                                     // lingering_close on;
	LingeringTime              string          `json:"lingeringTime,omitempty" key:"lingering_time" eg:"lingering_time 30s"`                                                                       // lingering_time 30s;
	LingeringTimeout           string          `json:"lingeringTimeout,omitempty" key:"lingering_timeout" eg:"lingering_timeout 5s"`                                                               // lingering_timeout 5s;
	LogNotFound                string          `json:"logNotFound,omitempty" key:"log_not_found" eg:"log_not_found on"`                                                                            // log_not_found on;
	LogSubrequest              string          `json:"logSubrequest,omitempty" key:"log_subrequest" eg:"log_subrequest off"`                                                                       // log_subrequest off;
	MaxRanges                  string          `json:"maxRanges,omitempty" key:"max_ranges" eg:"max_ranges number"`                                                                                // max_ranges number;
	MergeSlashes               string          `json:"mergeSlashes,omitempty" key:"merge_slashes" eg:"merge_slashes on"`                                                                           // merge_slashes on;
	MsiePadding                string          `json:"msiePadding,omitempty" key:"msie_padding" eg:"msie_padding on"`                                                                              // msie_padding on;
	MsieRefresh                string          `json:"msieRefresh,omitempty" key:"msie_refresh" eg:"msie_refresh off"`                                                                             // msie_refresh off;
	OpenFileCache              string          `json:"openFileCache,omitempty" key:"open_file_cache" eg:"open_file_cache off"`                                                                     // open_file_cache off;
	OpenFileCacheErrors        string          `json:"openFileCacheErrors" key:"open_file_cache_errors" eg:"open_file_cache_errors off"`                                                           // open_file_cache_errors off;
	OpenFileCacheMinUses       string          `json:"openFileCacheMinUses,omitempty" key:"open_file_cache_min_uses" eg:"open_file_cache_min_uses 1"`                                              // open_file_cache_min_uses 1;
	OpenFileCacheValid         string          `json:"openFileCacheValid,omitempty" key:"open_file_cache_valid" eg:"open_file_cache_valid 60s"`                                                    // open_file_cache_valid 60s;
	OutputBuffers              string          `json:"outputBuffers,omitempty" key:"output_buffers" eg:"output_buffers 2 32k"`                                                                     // output_buffers 2 32k;
	PortInRedirect             string          `json:"portInRedirect,omitempty" key:"port_in_redirect" eg:"port_in_redirect on"`                                                                   // port_in_redirect on;
	PostponeOutput             string          `json:"postponeOutput,omitempty" key:"postpone_output" eg:"postpone_output 1460"`                                                                   // postpone_output 1460;
	ReadAhead                  string          `json:"readAhead,omitempty" key:"read_ahead" eg:"read_ahead 0"`                                                                                     // read_ahead 0;
	RecursiveErrorPages        string          `json:"recursiveErrorPages,omitempty" key:"recursive_error_pages" eg:"recursive_error_pages off"`                                                   // recursive_error_pages off;
	RequestPoolSize            string          `json:"requestPoolSize,omitempty" key:"request_pool_size" eg:"request_pool_size 4k"`                                                                // request_pool_size 4k;
	Resolver                   string          `json:"resolver,omitempty" key:"resolver" eg:"resolver 127.0.0.1 [::1]:5353"`                                                                       // resolver 127.0.0.1 [::1]:5353;
	Satisfy                    string          `json:"satisfy,omitempty" key:"satisfy" eg:"satisfy all"`                                                                                           // satisfy all;
	SendLowat                  string          `json:"sendLowat,omitempty" key:"send_lowat" eg:"send_lowat 0"`                                                                                     // send_lowat 0;
	SendTimeout                string          `json:"sendTimeout,omitempty" key:"send_timeout" eg:"send_timeout 60s"`                                                                             // send_timeout 60s;
	SendFile                   string          `json:"sendFile,omitempty" key:"sendfile" eg:"sendfile on"`                                                                                         // sendfile on;
	SendFileMaxChunk           string          `json:"sendFileMaxChunk,omitempty" key:"sendfile_max_chunk" eg:"sendfile_max_chunk 0"`                                                              // sendfile_max_chunk 0;
	ServerNameInRedirect       string          `json:"serverNameInRedirect,omitempty" key:"server_name_in_redirect" eg:"server_name_in_redirect off"`                                              // server_name_in_redirect off;
	SubrequestOutputBufferSize string          `json:"subrequestOutputBufferSize,omitempty" key:"subrequest_output_buffer_size" eg:"subrequest_output_buffer_size 4k|8k"`                          // subrequest_output_buffer_size 4k|8k;
	TcpNodelay                 string          `json:"tcpNodelay,omitempty" key:"tcp_nodelay" eg:"tcp_nodelay on"`                                                                                 //tcp_nodelay on;
	TcpNopush                  string          `json:"tcpNopush,omitempty" key:"tcp_nopush" eg:"tcp_nopush off"`                                                                                   // tcp_nopush off;
	TryFiles                   string          `json:"tryFiles,omitempty" key:"try_files" eg:"try_files file ... =code"`                                                                           // try_files file ... =code;
	UnderscoresInHeaders       string          `json:"underscoresInHeaders,omitempty" key:"underscores_in_headers" eg:"underscores_in_headers off"`                                                // underscores_in_headers off;
	TypesHashBucketSize        string          `json:"typesHashBucketSize,omitempty" key:"types_hash_bucket_size" eg:"types_hash_bucket_size 64"`                                                  // types_hash_bucket_size 64;
	TypesHashMaxSize           string          `json:"typesHashMaxSize,omitempty" key:"types_hash_max_size" eg:"types_hash_max_size 1024"`                                                         // types_hash_max_size 1024;
	Aio                        string          `json:"aio,omitempty" key:"aio" eg:"aio off; on | off | threads[=pool]"`                                                                            // aio off; on | off | threads[=pool];
	AbsoluteRedirect           string          `json:"absoluteRedirect,omitempty" key:"absolute_redirect" eg:"absolute_redirect on"`                                                               // absolute_redirect on;
	AioWrite                   string          `json:"aioWrite,omitempty" key:"aio_write" eg:"aio_write off"`                                                                                      // aio_write off;
	ChunkedTransferEncoding    string          `json:"chunkedTransferEncoding,omitempty" key:"chunked_transfer_encoding" eg:"chunked_transfer_encoding on"`                                        // chunked_transfer_encoding on;
	ClientBodyBufferSize       string          `json:"clientBodyBufferSize,omitempty" key:"client_body_buffer_size" eg:"client_body_buffer_size 8k|16k"`                                           // client_body_buffer_size 8k|16k;
	ClientBodyInFileOnly       string          `json:"clientBodyInFileOnly,omitempty" key:"client_body_in_file_only" eg:"client_body_in_file_only off"`                                            // client_body_in_file_only off;
	ClientBodyInSingleBuffer   string          `json:"clientBodyInSingleBuffer,omitempty" key:"client_body_in_single_buffer" eg:"client_body_in_single_buffer off"`                                // client_body_in_single_buffer off;
	ClientBodyTempPath         string          `json:"clientBodyTempPath,omitempty" key:"client_body_temp_path" eg:"client_body_temp_path client_body_temp"`                                       // client_body_temp_path client_body_temp;
	ClientBodyTimeout          string          `json:"clientBodyTimeout,omitempty" key:"client_body_timeout" eg:"client_body_timeout 60s"`                                                         // client_body_timeout 60s;
	ClientHeaderBufferSize     string          `json:"clientHeaderBufferSize,omitempty" key:"client_header_buffer_size" eg:"client_header_buffer_size 1k"`                                         // client_header_buffer_size 1k;
	ClientHeaderTimeout        string          `json:"clientHeaderTimeout,omitempty" key:"client_header_timeout" eg:"client_header_timeout 60s"`                                                   // client_header_timeout 60s;
	ClientMaxBodySize          string          `json:"clientMaxBodySize,omitempty" key:"client_max_body_size" eg:"client_max_body_size 1m"`                                                        // client_max_body_size 1m;
	ConnectionPoolSize         string          `json:"connectionPoolSize,omitempty" key:"connection_pool_size" eg:"connection_pool_size 256|512"`                                                  // connection_pool_size 256|512;
	DirectIO                   string          `json:"directio,omitempty" key:"directio" eg:"directio 512"`                                                                                        // directio 512;
	DirectioAlignment          string          `json:"directioAlignment,omitempty" key:"directio_alignment" eg:"directio_alignment 512"`                                                           // directio_alignment 512;
	DisableSymlinks            string          `json:"disableSymlinks,omitempty" key:"disable_symlinks" eg:"disable_symlinks off"`                                                                 // disable_symlinks off;
	ErrorPage                  string          `json:"errorPage,omitempty" key:"error_page" eg:"error_page 500 502 503 504 /index.html"`                                                           // error_page 500 502 503 504 /index.html;
	IgnoreInvalidHeaders       string          `json:"ignoreInvalidHeaders,omitempty" key:"ignore_invalid_headers" eg:"ignore_invalid_headers on"`                                                 // ignore_invalid_headers on;
	ProxyBind                  string          `json:"proxyBind,omitempty" key:"proxy_bind" eg:"proxy_bind $remote_addr transparent"`                                                              // proxy_bind $remote_addr transparent;
	ProxyBufferSize            string          `json:"proxyBufferSize,omitempty" key:"proxy_buffer_size" eg:"proxy_buffer_size 4k|8k"`                                                             // proxy_buffer_size 4k|8k;
	ProxyBuffering             string          `json:"proxyBuffering,omitempty" key:"proxy_buffering" eg:"proxy_buffering on"`                                                                     // proxy_buffering on;
	ProxyBuffers               string          `json:"proxyBuffers,omitempty" key:"proxy_buffers" eg:"proxy_buffers 8 4k|8k"`                                                                      // proxy_buffers 8 4k|8k;
	ProxyBusyBuffersSize       string          `json:"ProxyBusyBuffersSize,omitempty" key:"proxy_busy_buffers_size" eg:"proxy_busy_buffers_size 8k|16k"`                                           // proxy_busy_buffers_size 8k|16k;
	ProxyCache                 string          `json:"proxyCache,omitempty" key:"proxy_cache" eg:"proxy_cache off"`                                                                                // proxy_cache off;
	ProxyCacheBackgroundUpdate string          `json:"proxyCacheBackgroundUpdate,omitempty" key:"proxy_cache_background_update" eg:"proxy_cache_background_update off"`                            // proxy_cache_background_update off;
	ProxyCacheBypass           []string        `json:"proxyCacheBypass,omitempty" key:"proxy_cache_bypass" eg:"proxy_cache_bypass  $cookie_nocache $arg_nocache$arg_comment"`                      // proxy_cache_bypass  $cookie_nocache $arg_nocache$arg_comment;
	ProxyCacheConvertHead      string          `json:"proxyCacheConvertHead,omitempty" key:"proxy_cache_convert_head" eg:"proxy_cache_convert_head on"`                                            // proxy_cache_convert_head on;
	ProxyCacheKey              string          `json:"proxyCacheKey,omitempty" key:"proxy_cache_key" eg:"proxy_cache_key $scheme$proxy_host$request_uri"`                                          // proxy_cache_key $scheme$proxy_host$request_uri;
	ProxyCacheLock             string          `json:"proxyCacheLock,omitempty" key:"proxy_cache_lock" eg:"proxy_cache_lock off"`                                                                  // proxy_cache_lock off;
	ProxyCacheLockAge          string          `json:"proxyCacheLockAge,omitempty" key:"proxy_cache_lock_age" eg:"proxy_cache_lock_age 5s"`                                                        // proxy_cache_lock_age 5s;
	ProxyCacheLockTimeout      string          `json:"proxyCacheLockTimeout,omitempty" key:"proxy_cache_lock_timeout" eg:"proxy_cache_lock_timeout 5s"`                                            // proxy_cache_lock_timeout 5s;
	ProxyCacheMaxRangeOffset   string          `json:"proxyCacheMaxRangeOffset,omitempty" key:"proxy_cache_max_range_offset" eg:"proxy_cache_max_range_offset number"`                             // proxy_cache_max_range_offset number;
	ProxyCacheMethods          string          `json:"proxyCacheMethods,omitempty" key:"proxy_cache_methods" eg:"proxy_cache_methods GET HEAD"`                                                    // proxy_cache_methods GET HEAD;
	ProxyCacheMinUses          string          `json:"proxyCacheMinUses,omitempty" key:"proxy_cache_min_uses" eg:"proxy_cache_min_uses 1"`                                                         // proxy_cache_min_uses 1;
	ProxyCachePath             string          `json:"proxyCachePath,omitempty" key:"proxy_cache_path" eg:"proxy_cache_path" eg:"proxy_cache_path /data/nginx/cache levels=1:2 keys_zone=one:10m"` // proxy_cache_path /data/nginx/cache levels=1:2 keys_zone=one:10m;
	ProxyCachePurge            string          `json:"proxyCachePurge,omitempty" key:"proxy_cache_purge" eg:"proxy_cache_purge $purge_method"`                                                     // proxy_cache_purge $purge_method;
	ProxyCacheUseStale         string          `json:"proxyCacheUseStale,omitempty" key:"proxy_cache_use_stale" eg:"proxy_cache_use_stale off"`                                                    // proxy_cache_use_stale off;
	ProxyCacheValid            []string        `json:"proxyCacheValid,omitempty" key:"proxy_cache_valid" eg:"proxy_cache_valid 200 302 10m; proxy_cache_valid 404      1m"`                        // proxy_cache_valid 200 302 10m; proxy_cache_valid 404      1m;
	ProxyConnectTimeout        string          `json:"ProxyConnectTimeout,omitempty" key:"proxy_connect_timeout" eg:"proxy_connect_timeout 60s"`                                                   // proxy_connect_timeout 60s;
	ProxyCookieDomain          string          `json:"proxyCookieDomain,omitempty" key:"proxy_cookie_domain" eg:"proxy_cookie_domain off"`                                                         // proxy_cookie_domain off;
	ProxyCookiePath            string          `json:"proxyCookiePath,omitempty" key:"proxy_cookie_path" eg:"proxy_cookie_path off"`                                                               // proxy_cookie_path off;
	ProxyForceRanges           string          `json:"proxyForceRanges,omitempty" key:"proxy_force_ranges" eg:"proxy_force_ranges off"`                                                            // proxy_force_ranges off;
	ProxyHeadersHashBucketSize string          `json:"proxyHeadersHashBucketSize,omitempty" key:"proxy_headers_hash_bucket_size" eg:"proxy_headers_hash_bucket_size 64"`                           // proxy_headers_hash_bucket_size 64;
	ProxyHeadersHashMaxSize    string          `json:"proxyHeadersHashMaxSize,omitempty" key:"proxy_headers_hash_max_size" eg:"proxy_headers_hash_max_size 512"`                                   // proxy_headers_hash_max_size 512;
	ProxyHttpVersion           string          `json:"proxyHttpVersion,omitempty" key:"proxy_http_version" eg:"proxy_http_version 1.0"`                                                            // proxy_http_version 1.0;
	ProxyIgnoreClientAbort     string          `json:"proxyIgnoreClientAbort,omitempty" key:"proxy_ignore_client_abort" eg:"proxy_ignore_client_abort off"`                                        // proxy_ignore_client_abort off;
	ProxyIgnoreHeaders         string          `json:"proxyIgnoreHeaders,omitempty" key:"proxy_ignore_headers" eg:"proxy_ignore_headers field ..."`                                                // proxy_ignore_headers field ...;
	ProxyInterceptErrors       string          `json:"proxyInterceptErrors,omitempty" key:"proxy_intercept_errors" eg:"proxy_intercept_errors off"`                                                // proxy_intercept_errors off;
	ProxyLimitRate             string          `json:"proxyLimitRate,omitempty" key:"proxy_limit_rate" eg:"proxy_limit_rate 0"`                                                                    // proxy_limit_rate 0;
	ProxyMaxTempFileSize       string          `json:"proxyMaxTempFileSize,omitempty" key:"proxy_max_temp_file_size" eg:"proxy_max_temp_file_size 1024m"`                                          // proxy_max_temp_file_size 1024m;
	ProxyMethod                string          `json:"proxyMethod,omitempty" key:"proxy_method" eg:"proxy_method method"`                                                                          // proxy_method method;
	ProxyNextUpstream          string          `json:"proxyNextUpstream,omitempty" key:"proxy_next_upstream" eg:"proxy_next_upstream error | timeout"`                                             // proxy_next_upstream error | timeout
	ProxyNextUpstreamTimeout   string          `json:"proxyNextUpstreamTimeout,omitempty" key:"proxy_next_upstream_timeout" eg:"proxy_next_upstream_timeout 0"`                                    // proxy_next_upstream_timeout 0;
	ProxyNextUpstreamTries     string          `json:"proxyNextUpstreamTries,omitempty" key:"proxy_next_upstream_tries" eg:"proxy_next_upstream_tries 0"`                                          // proxy_next_upstream_tries 0;
	ProxyNoCache               string          `json:"proxyNoCache,omitempty" key:"proxy_no_cache" eg:"proxy_no_cache string ..."`                                                                 // proxy_no_cache string ...;
	ProxyPass                  string          `json:"proxyPass,omitempty" key:"proxy_pass" eg:"proxy_pass http://example_app_airsupport_http"`                                                    // proxy_pass http://example_app_airsupport_http;
	ProxyPassHeader            string          `json:"proxyPassHeader,omitempty" key:"proxy_pass_header" eg:"proxy_pass_header field"`                                                             // proxy_pass_header field
	ProxyPassRequestBody       string          `json:"proxyPassRequestBody,omitempty" key:"proxy_pass_request_body" eg:"proxy_pass_request_body on"`                                               // proxy_pass_request_body on;
	ProxyPassRequestHeaders    string          `json:"proxyPassRequestHeaders,omitempty" key:"proxy_pass_request_headers" eg:"proxy_pass_request_headers on"`                                      // proxy_pass_request_headers on;
	ProxyReadTimeout           string          `json:"proxyReadTimeout,omitempty" key:"proxy_read_timeout" eg:"proxy_read_timeout 60s"`                                                            // proxy_read_timeout 60s;
	ProxyRedirect              string          `json:"proxyRedirect,omitempty" key:"proxy_redirect" eg:"proxy_redirect redirect replacement"`                                                      // proxy_redirect redirect replacement;
	ProxyRequestBuffering      string          `json:"proxyRequestBuffering,omitempty" key:"proxy_request_buffering" eg:"proxy_request_buffering on"`                                              // proxy_request_buffering on;
	ProxySendLowat             string          `json:"proxySendLowat,omitempty" key:"proxy_send_lowat" eg:"proxy_send_lowat 0"`                                                                    // proxy_send_lowat 0;
	ProxySendTimeout           string          `json:"proxySendTimeout,omitempty" key:"proxy_send_timeout" eg:"proxy_send_timeout 60s"`                                                            // proxy_send_timeout 60s;
	ProxySetBody               string          `json:"proxySetBody,omitempty" key:"proxy_set_body" eg:"proxy_set_body value"`                                                                      // proxy_set_body value;
	ProxySetHeader             []string        `json:"proxySetHeader,omitempty" key:"proxy_set_header" eg:"proxy_set_header Host $proxy_host; proxy_set_header Connection close"`                  // proxy_set_header Host $proxy_host; proxy_set_header Connection close;
	ProxySocketKeepalive       string          `json:"proxySocketKeepalive,omitempty" key:"proxy_socket_keepalive" eg:"proxy_socket_keepalive off"`                                                // proxy_socket_keepalive off;
	ProxySslCertificate        string          `json:"proxySslCertificate,omitempty" key:"proxy_ssl_certificate" eg:"proxy_ssl_certificate file"`                                                  // proxy_ssl_certificate file;
	ProxySslCertificateKey     string          `json:"proxySslCertificateKey,omitempty" key:"proxy_ssl_certificate_key" eg:"proxy_ssl_certificate_key file"`                                       // proxy_ssl_certificate_key file;
	ProxySslCiphers            string          `json:"proxySslCiphers,omitempty" key:"proxy_ssl_ciphers" eg:"proxy_ssl_ciphers DEFAULT"`                                                           // proxy_ssl_ciphers DEFAULT;
	ProxySslCrl                string          `json:"proxySslCrl,omitempty" key:"proxy_ssl_crl" eg:"proxy_ssl_crl file"`                                                                          // proxy_ssl_crl file;
	ProxySslName               string          `json:"proxySslName,omitempty" key:"proxy_ssl_name" eg:"proxy_ssl_name name"`                                                                       // proxy_ssl_name name;
	ProxySslPasswordFile       string          `json:"proxySslPasswordFile,omitempty" key:"proxy_ssl_password_file" eg:"proxy_ssl_password_file file"`                                             // proxy_ssl_password_file file;
	ProxySslProtocols          string          `json:"proxySslProtocols,omitempty" key:"proxy_ssl_protocols" eg:"proxy_ssl_protocols TLSv1 TLSv1.1 TLSv1.2"`                                       // proxy_ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
	ProxySslServerName         string          `json:"proxySslServerName,omitempty" key:"proxy_ssl_server_name" eg:"proxy_ssl_server_name off"`                                                    // proxy_ssl_server_name off;
	ProxySslSessionReuse       string          `json:"proxySslSessionReuse,omitempty" key:"proxy_ssl_session_reuse" eg:"proxy_ssl_session_reuse on"`                                               // proxy_ssl_session_reuse on;
	ProxySslTrustedCertificate string          `json:"proxySslTrustedCertificate,omitempty" key:"proxy_ssl_trusted_certificate" eg:"proxy_ssl_trusted_certificate file"`                           // proxy_ssl_trusted_certificate file;
	ProxySslVerify             string          `json:"proxySslVerify,omitempty" key:"proxy_ssl_verify" eg:"proxy_ssl_verify off"`                                                                  // proxy_ssl_verify off;
	ProxySslVerifyDepth        string          `json:"proxySslVerifyDepth,omitempty" key:"proxy_ssl_verify_depth" eg:"proxy_ssl_verify_depth 1"`                                                   // proxy_ssl_verify_depth 1;
	ProxyStore                 string          `json:"proxyStore,omitempty" key:"proxy_store" eg:"proxy_store off"`                                                                                // proxy_store off;
	ProxyStoreAccess           string          `json:"proxyStoreAccess,omitempty" key:"proxy_store_access" eg:"proxy_store_access user:rw"`                                                        // proxy_store_access user:rw;
	ProxyTempFileWriteSize     string          `json:"proxyTempFileWriteSize,omitempty" key:"proxy_temp_file_write_size" eg:"proxy_temp_file_write_size 8k|16k"`                                   // proxy_temp_file_write_size 8k|16k;
	ProxyTempPath              string          `json:"proxyTempPath,omitempty" key:"proxy_temp_path" eg:"proxy_temp_path proxy_temp"`                                                              // proxy_temp_path proxy_temp;
	Allow                      []string        `json:"allow,omitempty" key:"allow" eg:"allow 119.161.147.101"`                                                                                     // allow 119.161.147.101;
	Deny                       []string        `json:"deny,omitempty" key:"deny" eg:"deny all"`                                                                                                    // deny all;
	Set                        []string        `json:"set,omitempty" key:"set" eg:"set $limit_rate 4k"`                                                                                            // set $limit_rate 4k;
	AddHeader                  []string        `json:"addHeader,omitempty" key:"add_header" eg:"add_header X-Frame-Options SAMEORIGIN"`
	Gzip                       string          `json:"gzip,omitempty" key:"gzip" eg:"gzip on"`
	GzipMinLength              string          `json:"gzip_min_length" key:"gzip_min_length" eg:"gzip_min_length 20"`
	GzipBuffers                string          `json:"gzip_buffers" key:"gzip_buffers" eg:"gzip_buffers 32 4k|16 8k"`
	GzipCompLevel              string          `json:"gzip_comp_level" key:"gzip_comp_level" eg:"gzip_comp_level 1"`
	GzipTypes                  string          `json:"gzip_types" key:"gzip_types" eg:"gzip_types text/html"`
	GzipVary                   string          `json:"gzip_vary" key:"gzip_vary" eg:"gzip_vary off"`
	GzipDisable                string          `json:"gzip_disable" key:"gzip_disable" eg:"gzip_disable \"msie6\""`
	GzipHttpVersion            string          `json:"gzip_http_version" key:"gzip_disable" eg:"gzip_http_version 1.1"`
	GzipProxied                string          `json:"gzip_proxied" key:"gzip_proxied" eg:"gzip_proxied off"`
	ProxyHideHeader            string          `json:"proxyHideHeader,omitempty" key:"proxy_hide_header" eg:"proxy_hide_header X-Powered-By"`
}

// NewServer new server module
func NewServer() *Server {
	return &Server{
		Rewrite:          make([]string, 0),
		IfBlocks:         make([]ServerIfBlock, 0),
		Locations:        make([]Location, 0),
		ProxyCacheBypass: make([]string, 0),
		ProxyCacheValid:  make([]string, 0),
		Allow:            make([]string, 0),
		Deny:             make([]string, 0),
		Set:              make([]string, 0),
		AddHeader:        make([]string, 0),
	}
}

// Http http module
type Http struct {
	Gzip                       string     `json:"gzip,omitempty" key:"gzip" eg:"gzip on"`                                  // gzip on;
	GzipMinLength              string     `json:"gzipMinLength,omitempty" key:"gzip_min_length" eg:"gzip_min_length 1100"` // gzip_min_length 1100;
	GzipBuffers                string     `json:"gzipBuffers,omitempty" key:"gzip_buffers" eg:"gzip_buffers 4 8k"`         // gzip_buffers 4 8k;
	GzipTypes                  string     `json:"gzipTypes,omitempty" key:"gzip_types" eg:"gzip_types text/plain"`         // gzip_types text/plain;
	GzipCompLevel              string     `json:"gzipCompLevel,omitempty" key:"gzip_comp_level" eg:"gzip_comp_level 1"`
	LimitConnZone              string     `json:"limitConnZone,omitempty" key:"limit_conn_zone" eg:"$binary_remote_addr zone=one:10m"`
	LimitReqZone               string     `json:"limitReqZone,omitempty" key:"limit_req_zone" eg:"$binary_remote_addr zone=one:10m rate=1r/s"`
	GzipVary                   string     `json:"gzipVary,omitempty" key:"gzip_vary" eg:"gzip_vary off"`
	GzipDisable                string     `json:"gzipDisable,omitempty" key:"gzip_disable" eg:"gzip_disable \"msie6\""`
	GzipHttpVersion            string     `json:"gzipHttpVersion,omitempty" key:"gzip_disable" eg:"gzip_http_version 1.1"`
	GzipProxied                string     `json:"gzip_proxied,omitempty" key:"gzip_proxied" eg:"gzip_proxied off"`
	Root                       string     `json:"root,omitempty" key:"root" eg:"root /data/w3;"`                                                                             // root /data/w3;
	KeepaliveDisable           string     `json:"keepaliveDisable,omitempty" key:"keepalive_disable" eg:"keepalive_disable msie6"`                                           // keepalive_disable msie6;
	KeepaliveRequests          string     `json:"keepaliveRequests,omitempty" key:"keepalive_requests" eg:"keepalive_requests 100"`                                          // keepalive_requests 100;
	KeepaliveTimeout           string     `json:"keepaliveTimeout,omitempty" key:"keepalive_timeout" eg:"keepalive_timeout 75s"`                                             // keepalive_timeout 75s;
	LargeClientHeaderBuffers   string     `json:"largeClientHeaderBuffers,omitempty" key:"large_client_header_buffers" eg:"large_client_header_buffers 4 8k"`                // large_client_header_buffers 4 8k;
	LimitRate                  string     `json:"limitRate,omitempty" key:"limit_rate" eg:"limit_rate 0"`                                                                    // limit_rate 0;
	LimitRateAfter             string     `json:"limitRateAfter,omitempty" key:"limit_rate_after" eg:"limit_rate_after 0"`                                                   // limit_rate_after 0;
	LingeringClose             string     `json:"lingeringClose,omitempty" key:"lingering_close" eg:"lingering_close on"`                                                    // lingering_close on;
	LingeringTime              string     `json:"lingeringTime,omitempty" key:"lingering_time" eg:"lingering_time 30s"`                                                      // lingering_time 30s;
	LingeringTimeout           string     `json:"lingeringTimeout,omitempty" key:"lingering_timeout" eg:"lingering_timeout 5s"`                                              // lingering_timeout 5s;
	LogNotFound                string     `json:"logNotFound,omitempty" key:"log_not_found" eg:"log_not_found on"`                                                           // log_not_found on;
	LogSubrequest              string     `json:"logSubrequest,omitempty" key:"log_subrequest" eg:"log_subrequest off"`                                                      // log_subrequest off;
	MaxRanges                  string     `json:"maxRanges,omitempty"  key:"max_ranges" eg:"max_ranges number"`                                                              // max_ranges number;
	MsiePadding                string     `json:"msiePadding,omitempty" key:"msie_padding" eg:"msie_padding on"`                                                             // msie_padding on;
	MsieRefresh                string     `json:"msieRefresh,omitempty" key:"msie_refresh" eg:"msie_refresh off"`                                                            // msie_refresh off;
	OpenFileCache              string     `json:"openFileCache,omitempty" key:"open_file_cache" eg:"open_file_cache off"`                                                    // open_file_cache off;
	OpenFileCacheErrors        string     `json:"openFileCacheErrors,omitempty" key:"open_file_cache_errors" eg:"open_file_cache_errors off"`                                // open_file_cache_errors off;
	OpenFileCacheMinUses       string     `json:"openFileCacheMinUses,omitempty" key:"open_file_cache_min_uses" eg:"open_file_cache_min_uses 1"`                             // open_file_cache_min_uses 1;
	OpenFileCacheValid         string     `json:"openFileCacheValid,omitempty" key:"open_file_cache_valid" eg:"open_file_cache_valid 60s"`                                   // open_file_cache_valid 60s;
	OutputBuffers              string     `json:"outputBuffers,omitempty" key:"output_buffers" eg:"output_buffers 2 32k"`                                                    // output_buffers 2 32k;
	PostponeOutput             string     `json:"postponeOutput,omitempty" key:"postpone_output" eg:"postpone_output 1460"`                                                  // postpone_output 1460;
	ReadAhead                  string     `json:"readAhead,omitempty" key:"read_ahead" eg:"read_ahead 0"`                                                                    // read_ahead 0;
	RecursiveErrorPages        string     `json:"recursiveErrorPages,omitempty" key:"recursive_error_pages" eg:"recursive_error_pages off"`                                  // recursive_error_pages off;
	RequestPoolSize            string     `json:"requestPoolSize,omitempty" key:"request_pool_size" eg:"request_pool_size 4k"`                                               // request_pool_size 4k;
	ResetTimedoutConnection    string     `json:"resetTimedoutConnection,omitempty" key:"reset_timedout_connection" eg:"reset_timedout_connection off"`                      // reset_timedout_connection off;
	Resolver                   string     `json:"resolver,omitempty" key:"resolver" eg:"resolver 127.0.0.1 [::1]:5353"`                                                      // resolver 127.0.0.1 [::1]:5353;
	Satisfy                    string     `json:"satisfy,omitempty" key:"satisfy" eg:"satisfy all"`                                                                          // satisfy all;
	SendLowat                  string     `json:"sendLowat,omitempty" key:"send_lowat" eg:"send_lowat 0"`                                                                    // send_lowat 0;
	SendTimeout                string     `json:"sendTimeout,omitempty" key:"send_timeout" eg:"send_timeout 60s"`                                                            // send_timeout 60s;
	SendFile                   string     `json:"sendFile,omitempty" key:"sendfile" eg:"sendfile on"`                                                                        // sendfile on;
	SendFileMaxChunk           string     `json:"sendfileMaxChunk,omitempty" key:"sendfile_max_chunk" eg:"sendfile_max_chunk 0"`                                             // sendfile_max_chunk 0;
	ServerNameInRedirect       string     `json:"serverNameInRedirect,omitempty" key:"server_name_in_redirect" eg:"server_name_in_redirect off"`                             // server_name_in_redirect off;
	ServerNamesHashBucketSize  string     `json:"serverNamesHashBucketSize,omitempty" key:"server_names_hash_bucket_size" eg:"server_names_hash_bucket_size 32|64|128"`      // server_names_hash_bucket_size 32|64|128;
	ServerNamesHashMaxSize     string     `json:"serverNamesHashMaxSize,omitempty" key:"server_names_hash_max_size" eg:"server_names_hash_max_size 512"`                     // server_names_hash_max_size 512;
	SubrequestOutputBufferSize string     `json:"subrequestOutputBufferSize,omitempty" key:"subrequest_output_buffer_size" eg:"subrequest_output_buffer_size 4k|8k"`         // subrequest_output_buffer_size 4k|8k;
	TcpNodelay                 string     `json:"tcpNodelay,omitempty" key:"tcp_nodelay" eg:"tcp_nodelay on"`                                                                //tcp_nodelay on;
	TcpNopush                  string     `json:"tcpNopush,omitempty" key:"tcp_nopush" eg:"tcp_nopush off"`                                                                  // tcp_nopush off;
	VariablesHashMaxSize       string     `json:"variablesHashMaxSize,omitempty" key:"variables_hash_max_size" eg:"variables_hash_max_size 1024"`                            // variables_hash_max_size 1024;
	VariablesHashBucketSize    string     `json:"variablesHashBucketSize,omitempty" key:"variables_hash_bucket_size" eg:"variables_hash_bucket_size 64"`                     // variables_hash_bucket_size 64;
	UnderscoresInHeaders       string     `json:"underscoresInHeaders,omitempty" key:"underscores_in_headers" eg:"underscores_in_headers off"`                               // underscores_in_headers off;
	TypesHashBucketSize        string     `json:"typesHashBucketSize,omitempty" key:"types_hash_bucket_size" eg:"types_hash_bucket_size 64"`                                 // types_hash_bucket_size 64;
	TypesHashMaxSize           string     `json:"typesHashMaxSize,omitempty" key:"types_hash_max_size" eg:"types_hash_max_size 1024"`                                        // types_hash_max_size 1024;
	Aio                        string     `json:"aio,omitempty" key:"aio" eg:"aio off; on | off | threads[=pool]"`                                                           // aio off; on | off | threads[=pool];
	AbsoluteRedirect           string     `json:"absoluteRedirect,omitempty" key:"absolute_redirect" eg:"absolute_redirect on"`                                              // absolute_redirect on;
	AioWrite                   string     `json:"aioWrite,omitempty" key:"aio_write" eg:"aio_write off"`                                                                     // aio_write off;
	ChunkedTransferEncoding    string     `json:"chunkedTransferEncoding,omitempty" key:"chunked_transfer_encoding" eg:"chunked_transfer_encoding on"`                       // chunked_transfer_encoding on;
	ClientBodyBufferSize       string     `json:"clientBodyBufferSize,omitempty" key:"client_body_buffer_size" eg:"client_body_buffer_size 8k|16k"`                          // client_body_buffer_size 8k|16k;
	ClientBodyInFileOnly       string     `json:"clientBodyInFileOnly,omitempty" key:"client_body_in_file_only" eg:"client_body_in_file_only off"`                           // client_body_in_file_only off;
	ClientBodyInSingleBuffer   string     `json:"clientBodyInSingleBuffer,omitempty" key:"client_body_in_single_buffer" eg:"client_body_in_single_buffer off"`               // client_body_in_single_buffer off;
	ClientBodyTempPath         string     `json:"clientBodyTempPath,omitempty" key:"client_body_temp_path" eg:"client_body_temp_path client_body_temp"`                      // client_body_temp_path client_body_temp;
	ClientBodyTimeout          string     `json:"clientBodyTimeout,omitempty" key:"client_body_timeout" eg:"client_body_timeout 60s"`                                        // client_body_timeout 60s;
	ClientHeaderBufferSize     string     `json:"clientHeaderBufferSize,omitempty" key:"client_header_buffer_size" eg:"client_header_buffer_size 1k"`                        // client_header_buffer_size 1k;
	ClientHeaderTimeout        string     `json:"clientHeaderTimeout,omitempty" key:"client_header_timeout" eg:"client_header_timeout 60s"`                                  // client_header_timeout 60s;
	ClientMaxBodySize          string     `json:"clientMaxBodySize,omitempty" key:"client_max_body_size" eg:"client_max_body_size 1m"`                                       // client_max_body_size 1m;
	ConnectionPoolSize         string     `json:"connectionPoolSize,omitempty" key:"connection_pool_size" eg:"connection_pool_size 256|512"`                                 // connection_pool_size 256|512;
	DirectIO                   string     `json:"directio,omitempty" key:"directio" eg:"directio 512"`                                                                       // directio 512;
	DirectioAlignment          string     `json:"directioAlignment,omitempty" key:"directio_alignment" eg:"directio_alignment 512"`                                          // directio_alignment 512;
	DisableSymlinks            string     `json:"disableSymlinks,omitempty" key:"disable_symlinks" eg:"disable_symlinks off"`                                                // disable_symlinks off;
	ErrorPage                  string     `json:"errorPage,omitempty" key:"error_page" eg:"error_page 500 502 503 504 /50x.html"`                                            // error_page 500 502 503 504 /50x.html;
	IgnoreInvalidHeaders       string     `json:"ignoreInvalidHeaders,omitempty" key:"ignore_invalid_headers" eg:"ignore_invalid_headers on"`                                // ignore_invalid_headers on;
	Server                     []Server   `json:"server,omitempty" key:"server" eg:"server {...}"`                                                                           // server {...}
	AccessLog                  string     `json:"accessLog,omitempty" key:"access_log" eg:"access_log"`                                                                      // access_log
	ProxyBind                  string     `json:"proxyBind,omitempty" key:"proxy_bind" eg:"proxy_bind $remote_addr transparent"`                                             // proxy_bind $remote_addr transparent;
	ProxyBufferSize            string     `json:"proxyBufferSize,omitempty" key:"proxy_buffer_size" eg:"proxy_buffer_size 4k|8k"`                                            // proxy_buffer_size 4k|8k;
	ProxyBuffering             string     `json:"proxyBuffering,omitempty" key:"proxy_buffering" eg:"proxy_buffering on"`                                                    // proxy_buffering on;
	ProxyBuffers               string     `json:"proxyBuffers,omitempty" key:"proxy_buffers" eg:"proxy_buffers 8 4k|8k"`                                                     // proxy_buffers 8 4k|8k;
	ProxyBusyBuffersSize       string     `json:"proxyBusyBuffersSize,omitempty" key:"proxy_busy_buffers_size" eg:"proxy_busy_buffers_size 8k|16k"`                          // proxy_busy_buffers_size 8k|16k;
	ProxyCache                 string     `json:"proxyCache,omitempty" key:"proxy_cache" eg:"proxy_cache off"`                                                               // proxy_cache off;
	ProxyCacheBackgroundUpdate string     `json:"proxyCacheBackgroundUpdate,omitempty" key:"proxy_cache_background_update" eg:"proxy_cache_background_update off"`           // proxy_cache_background_update off;
	ProxyCacheBypass           []string   `json:"proxyCacheBypass,omitempty" key:"proxy_cache_bypass" eg:"proxy_cache_bypass  $cookie_nocache $arg_nocache$arg_comment"`     // proxy_cache_bypass  $cookie_nocache $arg_nocache$arg_comment;
	ProxyCacheConvertHead      string     `json:"proxyCacheConvertHead,omitempty" key:"proxy_cache_convert_head" eg:"proxy_cache_convert_head on"`                           // proxy_cache_convert_head on;
	ProxyCacheKey              string     `json:"proxyCacheKey,omitempty" key:"proxy_cache_key" eg:"proxy_cache_key $scheme$proxy_host$request_uri"`                         // proxy_cache_key $scheme$proxy_host$request_uri;
	ProxyCacheLock             string     `json:"proxyCacheLock,omitempty" key:"proxy_cache_lock" eg:"proxy_cache_lock off"`                                                 // proxy_cache_lock off;
	ProxyCacheLockAge          string     `json:"proxyCacheLockAge,omitempty" key:"proxy_cache_lock_age" eg:"proxy_cache_lock_age 5s"`                                       // proxy_cache_lock_age 5s;
	ProxyCacheLockTimeout      string     `json:"proxyCacheLockTimeout,omitempty" key:"proxy_cache_lock_timeout" eg:"proxy_cache_lock_timeout 5s"`                           // proxy_cache_lock_timeout 5s;
	ProxyCacheMaxRangeOffset   string     `json:"proxyCacheMaxRangeOffset,omitempty" key:"proxy_cache_max_range_offset" eg:"proxy_cache_max_range_offset number"`            // proxy_cache_max_range_offset number;
	ProxyCacheMethods          string     `json:"proxyCacheMethods,omitempty" key:"proxy_cache_methods" eg:"proxy_cache_methods GET HEAD"`                                   // proxy_cache_methods GET HEAD;
	ProxyCacheMinUses          string     `json:"proxyCacheMinUses,omitempty" key:"proxy_cache_min_uses" eg:"proxy_cache_min_uses 1"`                                        // proxy_cache_min_uses 1;
	ProxyCachePath             string     `json:"proxyCachePath,omitempty" key:"proxy_cache_path" eg:"proxy_cache_path /data/nginx/cache levels=1:2 keys_zone=one:10m"`      // proxy_cache_path /data/nginx/cache levels=1:2 keys_zone=one:10m;
	ProxyCachePurge            string     `json:"proxyCachePurge,omitempty" key:"proxy_cache_purge" eg:"proxy_cache_purge $purge_method"`                                    // proxy_cache_purge $purge_method;
	ProxyCacheUseStale         string     `json:"proxyCacheUseStale,omitempty" key:"proxy_cache_use_stale" eg:"proxy_cache_use_stale off"`                                   // proxy_cache_use_stale off;
	ProxyCacheValid            []string   `json:"proxyCacheValid,omitempty" key:"proxy_cache_valid" eg:"proxy_cache_valid 200 302 10m; proxy_cache_valid 404 1m"`            // proxy_cache_valid 200 302 10m; proxy_cache_valid 404      1m;
	ProxyConnectTimeout        string     `json:"proxyConnectTimeout,omitempty" key:"proxy_connect_timeout" eg:"proxy_connect_timeout 60s"`                                  // proxy_connect_timeout 60s;
	ProxyCookieDomain          string     `json:"proxyCookieDomain,omitempty" key:"proxy_cookie_domain" eg:"proxy_cookie_domain off"`                                        // proxy_cookie_domain off;
	ProxyCookiePath            string     `json:"proxyCookiePath,omitempty" key:"proxy_cookie_path" eg:"proxy_cookie_path off"`                                              // proxy_cookie_path off;
	ProxyForceRanges           string     `json:"proxyForceRanges,omitempty" key:"proxy_force_ranges" eg:"proxy_force_ranges off"`                                           // proxy_force_ranges off;
	ProxyHeadersHashBucketSize string     `json:"proxyHeadersHashBucketSize,omitempty" key:"proxy_headers_hash_bucket_size" eg:"proxy_headers_hash_bucket_size 64"`          // proxy_headers_hash_bucket_size 64;
	ProxyHeadersHashMaxSize    string     `json:"proxyHeadersHashMaxSize,omitempty" key:"proxy_headers_hash_max_size" eg:"proxy_headers_hash_max_size 512"`                  // proxy_headers_hash_max_size 512;
	ProxyHttpVersion           string     `json:"proxyHttpVersion,omitempty" key:"proxy_http_version" eg:"proxy_http_version 1.0"`                                           // proxy_http_version 1.0;
	ProxyIgnoreClientAbort     string     `json:"proxyIgnoreClientAbort,omitempty" key:"proxy_ignore_client_abort" eg:"proxy_ignore_client_abort off"`                       // proxy_ignore_client_abort off;
	ProxyIgnoreHeaders         string     `json:"proxyIgnoreHeaders,omitempty" key:"proxy_ignore_headers" eg:"proxy_ignore_headers field ..."`                               // proxy_ignore_headers field ...;
	ProxyInterceptErrors       string     `json:"proxyInterceptErrors,omitempty" key:"proxy_intercept_errors" eg:"proxy_intercept_errors off"`                               // proxy_intercept_errors off;
	ProxyLimitRate             string     `json:"proxyLimitRate,omitempty" key:"proxy_limit_rate" eg:"proxy_limit_rate 0"`                                                   // proxy_limit_rate 0;
	ProxyMaxTempFileSize       string     `json:"proxyMaxTempFileSize,omitempty" key:"proxy_max_temp_file_size" eg:"proxy_max_temp_file_size 1024m"`                         // proxy_max_temp_file_size 1024m;
	ProxyMethod                string     `json:"proxyMethod,omitempty" key:"proxy_method" eg:"proxy_method method"`                                                         // proxy_method method;
	ProxyNextUpstream          string     `json:"proxyNextUpstream,omitempty" key:"proxy_next_upstream" eg:"proxy_next_upstream error | timeout"`                            // proxy_next_upstream error | timeout
	ProxyNextUpstreamTimeout   string     `json:"proxyNextUpstreamTimeout,omitempty" key:"proxy_next_upstream_timeout" eg:"proxy_next_upstream_timeout 0"`                   // proxy_next_upstream_timeout 0;
	ProxyNextUpstreamTries     string     `json:"proxyNextUpstreamTries,omitempty" key:"proxy_next_upstream_tries" eg:"proxy_next_upstream_tries 0"`                         // proxy_next_upstream_tries 0;
	ProxyNoCache               string     `json:"proxyNoCache,omitempty" key:"proxy_no_cache" eg:"proxy_no_cache string ..."`                                                // proxy_no_cache string ...;
	ProxyPass                  string     `json:"proxyPass,omitempty" key:"proxy_pass" eg:"proxy_pass http://example_app_airsupport_http"`                                   // proxy_pass http://example_app_airsupport_http;
	ProxyPassHeader            string     `json:"proxyPassHeader,omitempty" key:"proxy_pass_header" eg:"proxy_pass_header field"`                                            // proxy_pass_header field
	ProxyPassRequestBody       string     `json:"proxyPassRequestBody,omitempty" key:"proxy_pass_request_body" eg:"proxy_pass_request_body on"`                              // proxy_pass_request_body on;
	ProxyPassRequestHeaders    string     `json:"proxyPassRequestHeaders,omitempty" key:"proxy_pass_request_headers" eg:"proxy_pass_request_headers on"`                     // proxy_pass_request_headers on;
	ProxyReadTimeout           string     `json:"proxyReadTimeout,omitempty" key:"proxy_read_timeout" eg:"proxy_read_timeout 60s"`                                           // proxy_read_timeout 60s;
	ProxyRedirect              string     `json:"proxyRedirect,omitempty" key:"proxy_redirect" eg:"proxy_redirect redirect replacement"`                                     // proxy_redirect redirect replacement;
	ProxyRequestBuffering      string     `json:"proxyRequestBuffering,omitempty" key:"proxy_request_buffering" eg:"proxy_request_buffering on"`                             // proxy_request_buffering on;
	ProxySendLowat             string     `json:"proxySendLowat,omitempty" key:"proxy_send_lowat" eg:"proxy_send_lowat 0"`                                                   // proxy_send_lowat 0;
	ProxySendTimeout           string     `json:"proxySendTimeout,omitempty" key:"proxy_send_timeout" eg:"proxy_send_timeout 60s"`                                           // proxy_send_timeout 60s;
	ProxySetBody               string     `json:"proxySetBody,omitempty" key:"proxy_set_body" eg:"proxy_set_body value"`                                                     // proxy_set_body value;
	ProxySetHeader             []string   `json:"proxySetHeader,omitempty" key:"proxy_set_header" eg:"proxy_set_header Host $proxy_host; proxy_set_header Connection close"` // proxy_set_header Host $proxy_host; proxy_set_header Connection close;
	ProxySocketKeepalive       string     `json:"proxySocketKeepalive,omitempty" key:"proxy_socket_keepalive" eg:"proxy_socket_keepalive off"`                               // proxy_socket_keepalive off;
	ProxySslCertificate        string     `json:"proxySslCertificate,omitempty" key:"proxy_ssl_certificate" eg:"proxy_ssl_certificate file"`                                 // proxy_ssl_certificate file;
	ProxySslCertificateKey     string     `json:"proxySslCertificateKey,omitempty" key:"proxy_ssl_certificate_key" eg:"proxy_ssl_certificate_key file"`                      // proxy_ssl_certificate_key file;
	ProxySslCiphers            string     `json:"proxySslCiphers,omitempty" key:"proxy_ssl_ciphers" eg:"proxy_ssl_ciphers DEFAULT"`                                          // proxy_ssl_ciphers DEFAULT;
	ProxySslCrl                string     `json:"proxySslCrl,omitempty" key:"proxy_ssl_crl" eg:"proxy_ssl_crl file"`                                                         // proxy_ssl_crl file;
	ProxySslName               string     `json:"proxySslName,omitempty" key:"proxy_ssl_name" eg:"proxy_ssl_name name"`                                                      // proxy_ssl_name name;
	ProxySslPasswordFile       string     `json:"proxySslPasswordFile,omitempty" key:"proxy_ssl_password_file" eg:"proxy_ssl_password_file file"`                            // proxy_ssl_password_file file;
	ProxySslProtocols          string     `json:"proxySslProtocols,omitempty" key:"proxy_ssl_protocols" eg:"proxy_ssl_protocols TLSv1 TLSv1.1 TLSv1.2"`                      // proxy_ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
	ProxySslServerName         string     `json:"proxySslServerName,omitempty" key:"proxy_ssl_server_name" eg:"proxy_ssl_server_name off"`                                   // proxy_ssl_server_name off;
	ProxySslSessionReuse       string     `json:"proxySslSessionReuse,omitempty" key:"proxy_ssl_session_reuse" eg:"proxy_ssl_session_reuse on"`                              // proxy_ssl_session_reuse on;
	ProxySslTrustedCertificate string     `json:"proxySslTrustedCertificate,omitempty" key:"proxy_ssl_trusted_certificate" eg:"proxy_ssl_trusted_certificate file"`          // proxy_ssl_trusted_certificate file;
	ProxySslVerify             string     `json:"proxySslVerify,omitempty" key:"proxy_ssl_verify" eg:"proxy_ssl_verify off"`                                                 // proxy_ssl_verify off;
	ProxySslVerifyDepth        string     `json:"proxySslVerifyDepth,omitempty" key:"proxy_ssl_verify_depth" eg:"proxy_ssl_verify_depth 1"`                                  // proxy_ssl_verify_depth 1;
	ProxyStore                 string     `json:"proxyStore,omitempty" key:"proxy_store" eg:"proxy_store off"`                                                               // proxy_store off;
	ProxyStoreAccess           string     `json:"proxyStoreAccess,omitempty" key:"proxy_store_access" eg:"proxy_store_access user:rw"`                                       // proxy_store_access user:rw;
	ProxyTempFileWriteSize     string     `json:"proxyTempFileWriteSize,omitempty" key:"proxy_temp_file_write_size" eg:"proxy_temp_file_write_size 8k|16k"`                  // proxy_temp_file_write_size 8k|16k;
	ProxyTempPath              string     `json:"proxyTempPath,omitempty" key:"proxy_temp_path" eg:"proxy_temp_path proxy_temp"`                                             // proxy_temp_path proxy_temp;
	Allow                      []string   `json:"allow,omitempty" key:"allow" eg:"allow 119.161.147.101"`                                                                    // allow 119.161.147.101;
	Deny                       []string   `json:"deny,omitempty" key:"deny" eg:"deny all"`                                                                                   // deny all;
	Geo                        []Geo      `json:"geo,omitempty"`
	LogFormat                  []string   `json:"logFormat,omitempty" key:"log_format" eg:"log_format download '$remote_addr - $remote_user [$time_local]'"` // log_format download '$remote_addr - $remote_user [$time_local] '
	Upstream                   []Upstream `json:"upstream,omitempty" key:"upstream" eg:"upstream xxx{}"`                                                     // upstream xxx{}
	DefaultType                string     `json:"defaultType,omitempty" key:"default_type" eg:"default_type text/plain"`                                     // default_type text/plain;
	Includes                   []string   `json:"include,omitempty" key:"include" eg:"include servers/*.conf"`
	AddHeader                  []string   `json:"addHeader,omitempty" key:"add_header" eg:"add_header X-Frame-Options SAMEORIGIN"`
	ProxyHideHeader            string     `json:"proxyHideHeader,omitempty" key:"proxy_hide_header" eg:"proxy_hide_header X-Powered-By"`
}

// NewHttp new http module
func NewHttp() *Http {
	return &Http{
		Server:           make([]Server, 0),
		Upstream:         make([]Upstream, 0),
		ProxyCacheBypass: make([]string, 0),
		ProxyCacheValid:  make([]string, 0),
		ProxySetHeader:   make([]string, 0),
		Geo:              make([]Geo, 0),
		LogFormat:        make([]string, 0),
		Allow:            make([]string, 0),
		Deny:             make([]string, 0),
		Includes:         make([]string, 0),
		AddHeader:        make([]string, 0),
	}
}

// Marshal http json marshal
func (hp *Http) Marshal() (string, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	data, err := json.Marshal(hp)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Unmarshal http json unmarshal
func (hp *Http) Unmarshal(data []byte) error {
	return jsoniter.Unmarshal(data, hp)
}

// UnmarshalFormFile http json unmarshal form file
func (hp *Http) UnmarshalFormFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("ioutil ReadFile error: err=%s", err)
		return err
	}

	return jsoniter.Unmarshal(data, hp)
}

// Geo module
type Geo struct {
	Name           string   `json:"name" key:"geo" eg:"geo [$address] $variable { ... }"`                          // geo [$address] $variable { ... };
	Ranges         string   `json:"ranges,omitempty" key:"ranges"`                                                 // ranges;
	Delete         string   `json:"delete,omitempty" key:"delete" eg:"delete 127.0.0.0/16"`                        // delete 127.0.0.0/16;
	Default        string   `json:"default,omitempty" key:"default" eg:"default 0"`                                // default 0;
	Include        string   `json:"include,omitempty" key:"include" eg:"include conf/geo.conf"`                    // include conf/geo.conf;
	Proxy          string   `json:"proxy,omitempty" key:"proxy" eg:"proxy 192.168.100.0/24"`                       // proxy 192.168.100.0/24;
	ProxyRecursive string   `json:"proxyRecursive,omitempty" key:"proxy_recursive"`                                // proxy_recursive
	List           []string `json:"list,omitempty" key:"list" eg:"127.0.0.0-127.0.0.0 US; 127.0.0.1-127.0.0.1 RU"` // 127.0.0.0-127.0.0.0 US; 127.0.0.1-127.0.0.1 RU;
}

// NewGeo new geo
func NewGeo() *Geo {
	return &Geo{List: make([]string, 0)}
}

// Event event module
type Event struct {
	AcceptMutex       string   `json:"acceptMutex,omitempty" key:"accept_mutex" eg:"accept_mutex off"`                                               // accept_mutex off;
	AcceptMutexDelay  string   `json:"acceptMutexDelay,omitempty" key:"accept_mutex_delay" eg:"accept_mutex_delay 500ms"`                            // accept_mutex_delay 500ms;
	DebugConnection   []string `json:"debugConnection,omitempty" key:"debug_connection" eg:"debug_connection 127.0.0.1; debug_connection localhost"` // debug_connection 127.0.0.1; debug_connection localhost;
	Use               string   `json:"use,omitempty" key:"use" eg:"use method"`                                                                      // use method;
	WorkerAioRequests string   `json:"workerAioRequests,omitempty" key:"worker_aio_requests" eg:"worker_aio_requests 32"`                            // worker_aio_requests 32;
	WorkerConnections string   `json:"workerConnections,omitempty" key:"worker_connections" eg:"worker_connections 512"`                             // worker_connections 512;
}

// NewEvent new event module
func NewEvent() *Event {
	return &Event{DebugConnection: make([]string, 0)}
}

// Global global module
type Global struct {
	Daemon                string   `json:"daemon,omitempty" key:"daemon" eg:"daemon on"`                                                    // daemon on;
	Env                   []string `json:"env,omitempty" key:"env" eg:"env MALLOC_OPTIONS"`                                                 // env MALLOC_OPTIONS;
	ErrorPage             string   `json:"errorPage,omitempty" key:"error_page" eg:"error_page 500 502 503 504 /50x.html"`                  // error_page 500 502 503 504 /50x.html;
	Events                Event    `json:"events,omitempty" key:"events" eg:"events {...}"`                                                 // events {...}
	Include               []string `json:"include,omitempty" key:"include" eg:"include mime.types"`                                         // include mime.types;
	LoadModule            string   `json:"loadModule,omitempty" key:"load_module" eg:"load_module file"`                                    // load_module file;
	LockFile              string   `json:"lockFile,omitempty" key:"lock_file" eg:"lock_file logs/nginx.lock"`                               // lock_file logs/nginx.lock;
	MasterProcess         string   `json:"masterProcess,omitempty" key:"master_process" eg:"master_process on"`                             // master_process on;
	PcreJIT               string   `json:"pcreJIT,omitempty" key:"pcre_jit" eg:"pcre_jit off"`                                              // pcre_jit off;
	Pid                   string   `json:"pid,omitempty" key:"pid" eg:"pid logs/nginx.pid"`                                                 // pid logs/nginx.pid;
	SslEngine             string   `json:"sslEngine,omitempty" key:"ssl_engine" eg:"ssl_engine device"`                                     // ssl_engine device;
	ThreadPool            string   `json:"threadPool,omitempty" key:"thread_pool" eg:"thread_pool default threads=32 max_queue=65536"`      // thread_pool default threads=32 max_queue=65536;
	TimerResolution       string   `json:"timerResolution,omitempty" key:"timer_resolution" eg:"timer_resolution 100ms"`                    // timer_resolution 100ms;
	User                  string   `json:"user,omitempty" key:"user nobody" eg:"user nobody nobody"`                                        // user nobody nobody;
	WorkerCpuAffinity     string   `json:"workerCpuAffinity,omitempty" key:"worker_cpu_affinity" eg:"worker_cpu_affinity auto [cpumask]"`   // worker_cpu_affinity auto [cpumask];
	WorkerPriority        string   `json:"workerPriority,omitempty" key:"worker_priority" eg:"worker_priority number"`                      // worker_priority number;
	WorkerProcesses       string   `json:"workerProcesses,omitempty" key:"worker_processes" eg:"worker_processes 1"`                        // worker_processes 1;
	WorkerRlimitCore      string   `json:"workerRlimitCore,omitempty" key:"worker_rlimit_core" eg:"worker_rlimit_core size"`                // worker_rlimit_core size;
	WorkerRlimitNofile    string   `json:"workerRlimitNofile,omitempty" key:"worker_rlimit_nofile" eg:"worker_rlimit_nofile number"`        // worker_rlimit_nofile number;
	WorkerShutdownTimeout string   `json:"workerShutdownTimeout,omitempty" key:"worker_shutdown_timeout" eg:"worker_shutdown_timeout time"` // worker_shutdown_timeout time;
	WorkingDirectory      string   `json:"workingDirectory,omitempty" key:"working_directory" eg:"working_directory directory"`             // working_directory directory;
	ProxyBind             string   `json:"proxyBind,omitempty" key:"proxy_bind" eg:"proxy_bind $remote_addr transparent"`                   // proxy_bind $remote_addr transparent;
	Http                  []Http   `json:"http,omitempty" key:"http" eg:"http {...}"`                                                       // http {...}
}

// NewGlobal new global module
func NewGlobal() *Global {
	return &Global{
		Env:     make([]string, 0),
		Include: make([]string, 0),
		Http:    make([]Http, 0),
	}
}
