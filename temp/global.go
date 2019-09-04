package temp

import (
	"fmt"
	"text/template"
)

// HTTPTpl http template
var (
	globalTpl *template.Template
)

func init() {
	var err error
	globalTpl, err = template.New("global_tpl").Parse(GlobalTemp)
	if err != nil {
		panic(fmt.Errorf("init main_tpl failed err:%v", err))
	}
}

// GetGlobalTpl get global template
func GetGlobalTpl() *template.Template {
	return globalTpl
}

// // MainTempToConf use template reander global nginx config
// func MainTempToConf(conf *core.Global) ([]byte, error) {
// 	content := new(bytes.Buffer)
// 	err := GlobalTpl.Execute(content, conf)
// 	return content.Bytes(), err
// }

// global template
const (
	GlobalTemp = `{{- if .Daemon}}
daemon {{.Daemon}};
{{- end }}
{{- if .ErrorPage}}
error_page {{.ErrorPage}};
{{- end }}
{{- if .LoadModule}}
load_module {{.LoadModule}};
{{- end }}
{{- if .LockFile}}
lock_file {{.LockFile}};
{{- end }}
{{- if .MasterProcess}}
master_process {{.MasterProcess}};
{{- end }}
{{- if .PcreJIT}}
pcre_jit {{.PcreJIT}}:
{{- end }}
{{- if .Pid}}
pid {{.Pid}};
{{- end }}
{{- if .SslEngine}}
ssl_engine {{.SslEngine}};
{{- end }}
{{- if .ThreadPool}}
thread_pool {{.ThreadPool}};
{{- end }}
{{- if .TimerResolution}}
timer_resolution {{.TimerResolution}};
{{- end }}
{{- if .User}}
user {{.User}};
{{- end }}
{{- if .WorkerCpuAffinity}}
worker_cpu_affinity {{.WorkerCpuAffinity}};
{{- end }}
{{- if .WorkerPriority}}
worker_priority {{.WorkerPriority}};
{{- end }}
{{- if .WorkerProcesses}}
worker_processes {{.WorkerProcesses}};
{{- end }}
{{- if .WorkerRlimitCore}}
worker_rlimit_core {{.WorkerRlimitCore}};
{{- end }}
{{- if .WorkerRlimitNofile}}
worker_rlimit_nofile {{.WorkerRlimitNofile}};
{{- end }}
{{- if .WorkerShutdownTimeout}}
worker_shutdown_timeout {{.WorkerShutdownTimeout}};
{{- end }}
{{- if .WorkingDirectory}}
working_directory {{.WorkingDirectory}};
{{- end }}
{{- if .ProxyBind}}
proxy_bind {{.ProxyBind}}
{{- end }}
{{- range $env := .Env}}
env {{.Env}};
{{- end }}
{{- range $include := .Include}}
include {{$include}};
{{- end }}
events {
	{{- if .Events.AcceptMutex}}
	accept_mutex {{.Events.AcceptMutex}};
	{{- end }}
	{{- if .Events.AcceptMutexDelay}}
	accept_mutex_delay {{.Events.AcceptMutexDelay}};
	{{- end }}
	{{- if .Events.Use}}
	use {{.Events.Use}};
	{{- end }}
	{{- if .Events.WorkerAioRequests}}
	worker_aio_requests {{.Events.WorkerAioRequests}};
	{{- end }}
	{{- if .Events.WorkerConnections}}
	worker_connections {{.Events.WorkerConnections}};
	{{- end }}
	{{- range $debug := .Events.DebugConnection}}
	debug_connection {{$debug}}
	{{- end }}
}
{{- range $http := .Http}}
http {
	{{- if .AddHeader}}
	{{range $header := .AddHeader}}
	add_header {{$header}};
	{{- end }}
	{{- end }}
	{{- if .DefaultType}}
	default_type {{.DefaultType}};
	{{- end }}
	log_format main '$remote_addr - $remote_user [$time_local] ' ' "$request" $status $bytes_sent $request_time ' ' "$http_referer" "$http_user_agent" ' ' "$http_x_forwarded_for" ' ' JSESSIONID=$cookie_JSESSIONID ' ' "$http_host" "$scheme" "$ssl_protocol" "$upstream_addr" "$upstream_response_time" ' ' "$http_x_yop_request_id" ';
	log_format download '$remote_addr - $remote_user [$time_local] ' '"$request" $status $bytes_sent $request_time $request_length ' '"$http_referer" "$http_user_agent" ' '"$http_range" "$sent_http_content_range"';
	{{- if .ClientHeaderBufferSize}}
	client_header_buffer_size {{.ClientHeaderBufferSize}};
	{{- end }}
	{{- if .LargeClientHeaderBuffers}}
	large_client_header_buffers {{.LargeClientHeaderBuffers}};
	{{- end }}
	{{- if .ServerNamesHashBucketSize}}
	server_names_hash_bucket_size {{.ServerNamesHashBucketSize}};
	{{- end }}
	{{- if .LimitConnZone}}
	limit_conn_zone {{.LimitConnZone}};
	{{- end }}
	{{- if .LimitReqZone}}
	limit_req_zone {{.LimitReqZone}};
	{{- end }}
	{{- if .Gzip}}
	gzip {{.Gzip}};
	{{- end }}
	{{- if .GzipMinLength}}
	gzip_min_length {{.GzipMinLength}};
	{{- end }}
	{{- if .GzipBuffers}}
	gzip_buffers {{.GzipBuffers}};
	{{- end }}
	{{- if .GzipTypes}}
	gzip_types {{.GzipTypes}};
	{{- end }}
	{{- if .GzipCompLevel}}
    gzip_comp_level {{.GzipCompLevel}};
    {{- end }}
	{{- if .GzipVary}}
    gzip_vary {{.GzipVary}};
    {{- end }}
	{{- if .GzipDisable}}
    gzip_disable {{.GzipDisable}};
    {{- end }}
	{{- if .GzipHttpVersion}}
    gzip_http_version {{.GzipHttpVersion}};
    {{- end }}
	{{- if .GzipProxied}}
    gzip_proxied {{.GzipProxied}};
    {{- end }}
	{{- if .ProxyHideHeader}}
    proxy_hide_header {{.ProxyHideHeader}};
    {{- end }}
	{{- if .OutputBuffers}}
	output_buffers {{.OutputBuffers}};
	{{- end }}
	{{- if .PostponeOutput}}
	postpone_output {{.PostponeOutput}};
	{{- end }}

	{{- if .AccessLog}}
	access_log {{.AccessLog}};
	{{- end }}
	{{- if .ClientHeaderTimeout}}
	client_header_timeout {{.ClientHeaderTimeout}};
	{{- end }}
	{{- if .ClientBodyTimeout}}
	client_body_timeout {{.ClientBodyTimeout}};
	{{- end }}
	{{- if .SendTimeout}}
	send_timeout {{.SendTimeout}};
	{{- end }}
	{{- if .SendFile}}
	sendfile {{.SendFile}};
	{{- end }}
	{{- if .TcpNodelay}}
	tcp_nodelay {{.TcpNodelay}};
	{{- end }}
	{{- if .TcpNopush}}
	tcp_nopush {{.TcpNopush}};
	{{- end }}
	{{- if .KeepaliveTimeout}}
	keepalive_timeout {{.KeepaliveTimeout}};
	{{- end }}

	{{- if .ProxyTempPath}}
	proxy_temp_path {{.ProxyTempPath}};
	{{- end }}
	{{- if .ProxyCachePath}}
	proxy_cache_path {{.ProxyCachePath}};
	{{- end }}
	{{- range $include := .Includes}}
	include {{$include}};
	{{- end }}
	{{- if .Root}}
	root {{.Root}};
	{{- end }}
	{{- if .KeepaliveDisable}}
	keepalive_disable {{.KeepaliveDisable}};
	{{- end }}
	{{- if .KeepaliveRequests}}
	keepalive_requests {{.KeepaliveRequests}};
	{{- end }}
	{{- if .LimitRate}}
	limit_rate {{.LimitRate}};
	{{- end }}
	{{- if .LimitRateAfter}}
	limit_rate_after {{.LimitRateAfter}};
	{{- end }}
	{{- if .LingeringClose}}
	lingering_close {{.LingeringClose}};
	{{- end }}
	{{- if .LingeringTime}}
	lingering_time {{.LingeringTime}};
	{{- end }}
	{{- if .LingeringTimeout}}
	lingering_timeout {{.LingeringTimeout}};
	{{- end }}
	{{- if .LogNotFound}}
	log_not_found {{.LogNotFound}};
	{{- end }}
	{{- if .LogSubrequest}}
	log_subrequest {{.LogSubrequest}};
	{{- end }}
	{{- if .MaxRanges}}
	max_ranges {{.MaxRanges}};
	{{- end }}
	{{- if .MsiePadding}}
	msie_padding {{.MsiePadding}};
	{{- end }}
	{{- if .MsieRefresh}}
	msie_refresh {{.MsieRefresh}};
	{{- end }}
	{{- if .OpenFileCache}}
	open_file_cache {{.OpenFileCache}};
	{{- end }}
	{{- if .OpenFileCacheErrors}}
	open_file_cache_errors {{.OpenFileCacheErrors}};
	{{- end }}
	{{- if .OpenFileCacheMinUses}}
	open_file_cache_min_uses {{.OpenFileCacheMinUses}};
	{{- end }}
	{{- if .OpenFileCacheValid}}
	open_file_cache_valid {{.OpenFileCacheValid}};
	{{- end }}
	{{- if .ReadAhead}}
	read_ahead {{.ReadAhead}};
	{{- end }}
	{{- if .RecursiveErrorPages}}
	recursive_error_pages {{.RecursiveErrorPages}};
	{{- end }}
	{{- if .RequestPoolSize}}
	request_pool_size {{.RequestPoolSize}};
	{{- end }}
	{{- if .ResetTimedoutConnection}}
	reset_timedout_connection {{.ResetTimedoutConnection}};
	{{- end }}
	{{- if .Resolver}}
	resolver {{.Resolver}};
	{{- end }}
	{{- if .Satisfy}}
	satisfy {{.Satisfy}};
	{{- end }}
	{{- if .SendLowat}}
	send_lowat {{.SendLowat}};
	{{- end }}
	{{- if .SendFileMaxChunk}}
	sendfile_max_chunk {{.SendFileMaxChunk}};
	{{- end }}
	{{- if .ServerNameInRedirect}}
	server_name_in_redirect {{.ServerNameInRedirect}};
	{{- end }}
	{{- if .ServerNamesHashMaxSize}}
	server_names_hash_max_size {{.ServerNamesHashMaxSize}};
	{{- end }}
	{{- if .SubrequestOutputBufferSize}}
	subrequest_output_buffer_size {{.SubrequestOutputBufferSize}};
	{{- end }}
	{{- if .VariablesHashMaxSize}}
	variables_hash_max_size {{.VariablesHashMaxSize}};
	{{- end }}
	{{- if .VariablesHashBucketSize}}
	variables_hash_bucket_size {{.VariablesHashBucketSize}};
	{{- end }}
	{{- if .UnderscoresInHeaders}}
	underscores_in_headers {{.UnderscoresInHeaders}};
	{{- end }}
	{{- if .TypesHashBucketSize}}
	types_hash_bucket_size {{.TypesHashBucketSize}};
	{{- end }}
	{{- if .TypesHashMaxSize}}
	types_hash_max_size {{.TypesHashMaxSize}};
	{{- end }}
	{{- if .Aio}}
	aio {{.Aio}};
	{{- end }}
	{{- if .AbsoluteRedirect}}
	absolute_redirect {{.AbsoluteRedirect}};
	{{- end }}
	{{- if .AioWrite}}
	aio_write {{.AioWrite}};
	{{- end }}
	{{- if .ChunkedTransferEncoding}}
	chunked_transfer_encoding {{.ChunkedTransferEncoding}};
	{{- end }}
	{{- if .ClientBodyBufferSize}}
	client_body_buffer_size {{.ClientBodyBufferSize}};
	{{- end }}
	{{- if .ClientBodyInFileOnly}}
	client_body_in_file_only {{.ClientBodyInFileOnly}};
	{{- end }}
	{{- if .ClientBodyInSingleBuffer}}
	client_body_in_single_buffer {{.ClientBodyInSingleBuffer}};
	{{- end }}
	{{- if .ClientBodyTempPath}}
	client_body_temp_path {{.ClientBodyTempPath}};
	{{- end }}
	{{- if .ClientMaxBodySize}}
	client_max_body_size {{.ClientMaxBodySize}};
	{{- end }}
	{{- if .ConnectionPoolSize}}
	connection_pool_size {{.ConnectionPoolSize}};
	{{- end }}
	{{- if .DirectIO}}
	directio {{.DirectIO}};
	{{- end }}
	{{- if .DirectioAlignment}}
	directio_alignment {{.DirectioAlignment}};
	{{- end }}
	{{- if .DisableSymlinks}}
	disable_symlinks {{.DisableSymlinks}};
	{{- end }}
	{{- if .ErrorPage}}
	error_page {{.ErrorPage}};
	{{- end }}
	{{- if .IgnoreInvalidHeaders}}
	ignore_invalid_headers {{.IgnoreInvalidHeaders}};
	{{- end }}
	{{- if .ProxyBind}}
	proxy_bind {{.ProxyBind}};
	{{- end }}
	{{- if .ProxyBufferSize}}
	proxy_buffer_size {{.ProxyBufferSize}};
	{{- end }}
	{{- if .ProxyBuffering}}
	proxy_buffering {{.ProxyBuffering}};
	{{- end }}
	{{- if .ProxyBuffers}}
	proxy_buffers {{.ProxyBuffers}};
	{{- end }}
	{{- if .ProxyBusyBuffersSize}}
	proxy_busy_buffers_size {{.ProxyBusyBuffersSize}};
	{{- end }}
	{{- if .ProxyCache}}
	proxy_cache {{.ProxyCache}};
	{{- end }}
	{{- if .ProxyCacheBackgroundUpdate}}
	proxy_cache_background_update {{.ProxyCacheBackgroundUpdate}};
	{{- end }}
	{{- if .ProxyCacheBypass}}
	{{- range $bypass := .ProxyCacheBypass}}
	proxy_cache_bypass {{$bypass}};
	{{- end }}
	{{- end }}
	{{- if .ProxyCacheConvertHead}}
	proxy_cache_convert_head {{.ProxyCacheConvertHead}};
	{{- end }}
	{{- if .ProxyCacheKey}}
	proxy_cache_key{{.ProxyCacheKey}};
	{{- end }}
	{{- if .ProxyCacheLock}}
	proxy_cache_lock {{.ProxyCacheLock}};
	{{- end }}
	{{- if .ProxyCacheLockAge}}
	proxy_cache_lock_age {{.ProxyCacheLockAge}};
	{{- end }}
	{{- if .ProxyCacheLockTimeout}}
	proxy_cache_lock_timeout {{.ProxyCacheLockTimeout}};
	{{- end }}
	{{- if .ProxyCacheMaxRangeOffset}}
	proxy_cache_max_range_offset {{.ProxyCacheMaxRangeOffset}};
	{{- end }}
	{{- if .ProxyCacheMethods}}
	proxy_cache_methods {{.ProxyCacheMethods}};
	{{- end }}
	{{- if .ProxyCacheMinUses}}
	proxy_cache_min_uses {{.ProxyCacheMinUses}};
	{{- end }}
	{{- if .ProxyCachePurge}}
	proxy_cache_purge {{.ProxyCachePurge}};
	{{- end }}
	{{- if .ProxyCacheUseStale}}
	proxy_cache_use_stale {{.ProxyCacheUseStale}};
	{{- end }}
	{{- if $valid := .ProxyCacheValid}}
	proxy_cache_valid {{$valid}};
	{{- end }}
	{{- if .ProxyConnectTimeout}}
	proxy_connect_timeout {{.ProxyConnectTimeout}};
	{{- end }}
	{{- if .ProxyCookieDomain}}
	proxy_cookie_domain {{.ProxyCookieDomain}};
	{{- end }}
	{{- if .ProxyCookiePath}}
	proxy_cookie_path {{.ProxyCookiePath}};
	{{- end }}
	{{- if .ProxyForceRanges}}
	proxy_force_ranges {{.ProxyForceRanges}};
	{{- end }}
	{{- if .ProxyHeadersHashBucketSize}}
	proxy_headers_hash_bucket_size {{.ProxyHeadersHashBucketSize}};
	{{- end }}
	{{- if .ProxyHeadersHashMaxSize}}
	proxy_headers_hash_max_size {{.ProxyHeadersHashMaxSize}};
	{{- end }}
	{{- if .ProxyHttpVersion}}
	proxy_http_version {{.ProxyHttpVersion}};
	{{- end }}
	{{- if .ProxyIgnoreClientAbort}}
	proxy_ignore_client_abort {{.ProxyIgnoreClientAbort}};
	{{- end }}
	{{- if .ProxyIgnoreHeaders}}
	proxy_ignore_headers {{.ProxyIgnoreHeaders}};
	{{- end }}
	{{- if .ProxyInterceptErrors}}
	proxy_intercept_errors {{.ProxyInterceptErrors}};
	{{- end }}
	{{- if .ProxyLimitRate}}
	proxy_limit_rate {{.ProxyLimitRate}};
	{{- end }}
	{{- if .ProxyMaxTempFileSize}}
	proxy_max_temp_file_size {{.ProxyMaxTempFileSize}};
	{{- end }}
	{{- if .ProxyMethod}}
	proxy_method {{.ProxyMethod}};
	{{- end }}
	{{- if .ProxyNextUpstream}}
	proxy_next_upstream {{.ProxyNextUpstream}};
	{{- end }}
	{{- if .ProxyNextUpstreamTimeout}}
	proxy_next_upstream_timeout {{.ProxyNextUpstreamTimeout}};
	{{- end }}
	{{- if .ProxyNextUpstreamTries}}
	proxy_next_upstream_tries {{.ProxyNextUpstreamTries}};
	{{- end }}
	{{- if .ProxyNoCache}}
	proxy_no_cache {{.ProxyNoCache}};
	{{- end }}
	{{- if .ProxyPass}}
	proxy_pass {{.ProxyPass}};
	{{- end }}
	{{- if .ProxyPassHeader}}
	proxy_pass_header {{.ProxyPassHeader}};
	{{- end }}
	{{- if .ProxyPassRequestBody}}
	proxy_pass_request_body {{.ProxyPassRequestBody}};
	{{- end }}
	{{- if .ProxyPassRequestHeaders}}
	proxy_pass_request_headers {{.ProxyPassRequestHeaders}};
	{{- end }}
	{{- if .ProxyReadTimeout}}
	proxy_read_timeout {{.ProxyReadTimeout}};
	{{- end }}
	{{- if .ProxyRedirect}}
	proxy_redirect {{.ProxyRedirect}};
	{{- end }}
	{{- if .ProxyRequestBuffering}}
	proxy_request_buffering {{.ProxyRequestBuffering}};
	{{- end }}
	{{- if .ProxySendLowat}}
	proxy_send_lowat {{.ProxySendLowat}};
	{{- end }}
	{{- if .ProxySendTimeout}}
	proxy_send_timeout {{.ProxySendTimeout}};
	{{- end }}
	{{- if .ProxySetBody}}
	proxy_set_body {{.ProxySetBody}};
	{{- end }}
	{{- if .ProxySetHeader}}
	{{- range $header := .ProxySetHeader}}
	proxy_set_header {{$header}};
	{{- end }}
	{{- end }}
	{{- if .ProxySocketKeepalive}}
	proxy_socket_keepalive {{.ProxySocketKeepalive}};
	{{- end }}
	{{- if .ProxySslCertificate}}
	proxy_ssl_certificate {{.ProxySslCertificate}};
	{{- end }}
	{{- if .ProxySslCertificateKey}}
	proxy_ssl_certificate_key {{.ProxySslCertificateKey}};
	{{- end }}
	{{- if .ProxySslCiphers}}
	proxy_ssl_ciphers {{.ProxySslCiphers}};
	{{- end }}
	{{- if .ProxySslCrl}}
	proxy_ssl_crl {{.ProxySslCrl}};
	{{- end }}
	{{- if .ProxySslName}}
	proxy_ssl_name {{.ProxySslName}};
	{{- end }}
	{{- if .ProxySslPasswordFile}}
	proxy_ssl_password_file {{.ProxySslPasswordFile}};
	{{- end }}
	{{- if .ProxySslProtocols}}
	proxy_ssl_protocols {{.ProxySslProtocols}};
	{{- end }}
	{{- if .ProxySslServerName}}
	proxy_ssl_server_name {{.ProxySslServerName}};
	{{- end }}
	{{- if .ProxySslSessionReuse}}
	proxy_ssl_session_reuse {{.ProxySslSessionReuse}};
	{{- end }}
	{{- if .ProxySslTrustedCertificate}}
	proxy_ssl_trusted_certificate {{.ProxySslTrustedCertificate}};
	{{- end }}
	{{- if .ProxySslVerify}}
	proxy_ssl_verify {{.ProxySslVerify}};
	{{- end }}
	{{- if .ProxySslVerifyDepth}}
	proxy_ssl_verify_depth {{.ProxySslVerifyDepth}};
	{{- end }}
	{{- if .ProxyStore}}
	proxy_store {{.ProxyStore}};
	{{- end }}
	{{- if .ProxyStoreAccess}}
	proxy_store_access {{.ProxyStoreAccess}};
	{{- end }}
	{{- if .ProxyTempFileWriteSize}}
	proxy_temp_file_write_size {{.ProxyTempFileWriteSize}};
	{{- end }}
	{{- if .Allow}}
	{{- range $allow := .Allow}}
	allow {{$allow}};
	{{- end }}
	{{- end }}
	{{- if .Deny}}
	{{- range $deny := .Deny}}
	deny {{$deny}};
	{{- end }}
	{{- end }}

	{{- range $geo := .Geo}}
	geo {{$geo.Name}} {
		{{- if $geo.Ranges}}
		ranges;
		{{- end }}
		{{- if $geo.Delete}}
		delete {{$geo.Delete}};
		{{- end }}
		{{- if $geo.Default}}
		default {{$geo.Default}};
		{{- end }}
		{{- if $geo.Include}}
		include {{$geo.Include}};
		{{- end }}
		{{- if $geo.Proxy}}
		proxy {{$geo.Proxy}};
		{{- end }}
		{{- if $geo.ProxyRecursive}}
		proxy_recursive {{$geo.ProxyRecursive}}
		{{- end }}
		{{range $list := $geo.List}}
		{{$list}};
		{{- end }}
	}
	{{- end }}

	{{- range $upstream := .Upstream }}
	upstream {{$upstream.Name}} {
		{{- if $upstream.Method}}
		{{$upstream.Method}};
		{{- end -}}
		{{range $server := $upstream.Servers}}
		server {{$server.Address}} {{- if $server.Params.Weight}} weight={{$server.Params.Weight}} {{- end -}} {{- if $server.Params.MaxConns}} max_conns={{$server.Params.MaxConns}} {{- end -}} {{- if $server.Params.MaxFails}} max_fails={{$server.Params.MaxFails}} {{- end -}} {{- if $server.Params.FailTimeout}} fail_timeout={{$server.Params.FailTimeout}} {{- end -}} {{- if $server.Params.Backup}} backup {{- end}} {{- if $server.Params.Down}} down {{- end}};
		{{- end -}}
		{{if $upstream.UpstreamKeepalive.Time}}
		keepalive {{$upstream.UpstreamKeepalive.Time}};
		{{- end}}
		{{- if .KeepaliveTimeout}}
		keepalive_timeout {{.KeepaliveTimeout}};
		{{- end }}
	}
	{{- end }}

	{{- range $server := .Server}}
	server {
		{{- if $server.AddHeader}}
		{{range $header := $server.AddHeader}}
		add_header {{$header}};
		{{- end }}
		{{- end }}
		{{- if $server.Name}}
		server_name {{$server.Name}};
		{{- end }}
		{{- if $server.Listen}}
		listen {{$server.Listen}};
		{{- end }}
		{{- if $server.ErrorPage}}
		error_page {{$server.ErrorPage}};
		{{- end }}
		{{- if $server.Return}}
		return {{$server.Return}};
		{{- end }}
		{{- if $server.Root}}
		root {{$server.Root}};
		{{- end }}
		{{- if $server.Ssl}}
		ssl {{$server.Ssl}};
		{{- end }}
		{{- if $server.SslCertificate}}
		ssl_certificate {{$server.SslCertificate}};
		{{- end }}
		{{- if $server.SslCertificateKey}}
		ssl_certificate_key {{$server.SslCertificateKey}};
		{{- end }}
		{{- if $server.SslProtocols}}
		ssl_protocols {{$server.SslProtocols}};
		{{- end }}
		{{- if $server.SslCiphers}}
		ssl_ciphers {{$server.SslCiphers}};
		{{- end }}
		{{- if $server.SslSessionTimeout}}
		ssl_session_timeout {{$server.SslSessionTimeout}};
		{{- end }}
		{{- if .LimitConn}}
		limit_conn {{.LimitConn}};
		{{- end }}
		{{- if .LimitReq}}
		limit_req {{.LimitReq}};
		{{- end }}
		{{- if $server.SslPreferServerCiphers}}
		ssl_prefer_server_ciphers {{$server.SslPreferServerCiphers}};
		{{- end }}
		{{- if $server.Rewrite}}
		{{range $rewrite := $server.Rewrite}}
		rewrite {{$rewrite}};
		{{- end }}
		{{- end }}
		{{- if .Set}}
		{{range $set := .Set}}
		set {{$set}};
		{{- end }}
		{{- end }}
		{{- if $server.IfBlocks}}
		{{range $ifBlock := $server.IfBlocks}}
		if ({{.Condition}}) {
			{{- if $ifBlock.ProxyPass}}
			proxy_pass {{$ifBlock.ProxyPass}};
			{{- end }}
			{{- if $ifBlock.SendFile}}
			send_file {{$ifBlock.SendFile}};
			{{- end }}
			{{- if $ifBlock.Root}}
			root {{$ifBlock.Root}};
			{{- end }}
			{{- if $ifBlock.Rewrite}}
			{{range $rewrite := $ifBlock.Rewrite}}
			rewrite {{$rewrite}};
			{{- end }}
			{{- end }}
			{{- if $ifBlock.Return}}
			return {{$ifBlock.Return}};
			{{- end }}
			{{- if $ifBlock.Set}}
			{{range $set := $ifBlock.Set}}
			set {{$set}};
			{{- end }}
			{{- end }}
			{{- if $ifBlock.LimitRateAfter}}
			limit_rate_after {{$ifBlock.LimitRateAfter}};
			{{- end }}
		}
		{{- end }}
		{{- end }}

		{{- if $server.AccessLog}}
		access_log {{$server.AccessLog}};
		{{- end }}
		{{- if $server.KeepaliveDisable}}
		keepalive_disable {{$server.KeepAliveDisable}};
		{{- end }}
		{{- if $server.KeepaliveRequests}}
		keepalive_requests {{$server.KeepAliveRequests}};
		{{- end }}
		{{- if $server.KeepaliveTimeout}}
		keepalive_timeout {{$server.KeepAliveTimeout}};
		{{- end }}
		{{- if $server.LargeClientHeaderBuffers}}
		largeClientHeaderBuffers {{$server.LargeClientHeaderBuffers}};
		{{- end }}
		{{- if $server.LimitRate}}
		limit_rate {{$server.LimitRate}};
		{{- end }}
		{{- if $server.LimitRateAfter}}
		limit_rate_after {{$server.LimitRateAfter}};
		{{- end }}
		{{- if $server.LingeringClose}}
		lingering_close {{$server.LingeringClose}};
		{{- end }}
		{{- if $server.LingeringTime}}
		lingering_time {{$server.LingeringTime}};
		{{- end }}
		{{- if $server.LingeringTimeout}}
		lingering_timeout {{$server.LingeringTimeout}};
		{{- end }}
		{{- if $server.LogNotFound}}
		log_not_found {{$server.LogNotFound}};
		{{- end }}
		{{- if $server.LogSubrequest}}
		log_subrequest {{$server.LogSubrequest}};
		{{- end }}
		{{- if $server.MaxRanges}}
		max_ranges {{$server.MaxRanges}};
		{{- end }}
		{{- if $server.MergeSlashes}}
		merge_slashes {{$server.MergeSlashes}};
		{{- end }}
		{{- if $server.MsiePadding}}
		msie_padding {{$server.MsiePadding}};
		{{- end }}
		{{- if $server.MsieRefresh}}
		msie_refresh {{$server.MsieRefresh}};
		{{- end }}
		{{- if $server.OpenFileCache}}
		open_file_cache {{$server.OpenFileCache}};
		{{- end }}
		{{- if $server.OpenFileCacheErrors}}
		open_file_cache_errors {{$server.OpenFileCacheErrors}};
		{{- end }}
		{{- if $server.OpenFileCacheMinUses}}
		open_file_cache_min_uses {{$server.OpenFileCacheMinUses}};
		{{- end }}
		{{- if $server.OpenFileCacheValid}}
		open_file_cache_valid {{$server.OpenFileCacheValid}};
		{{- end }}
		{{- if $server.OutputBuffers}}
		output_buffers {{$server.OutputBuffers}};
		{{- end }}
		{{- if $server.PostponeOutput}}
		postpone_output {{$server.PostponeOutput}};
		{{- end }}
		{{- if $server.ReadAhead}}
		read_ahead {{$server.ReadAhead}};
		{{- end }}
		{{- if $server.RequestPoolSize}}
		request_pool_size {{$server.RequestPoolSize}};
		{{- end }}
		{{- if $server.Resolver}}
		resolver {{$server.Resolver}};
		{{- end }}
		{{- if $server.Satisfy}}
		satisfy {{$server.Satisfy}};
		{{- end }}
		{{- if $server.SendLowat}}
		send_lowat {{$server.SendLowat}};
		{{- end }}
		{{- if $server.SendTimeout}}
		SendTimeout {{$server.SendTimeout}};
		{{- end }}
		{{- if $server.SendFile}}
		sendfile {{$server.SendFile}};
		{{- end }}
		{{- if $server.SendFileMaxChunk}}
		sendfile_max_chunk {{$server.SendFileMaxChunk}};
		{{- end }}
		{{- if $server.ServerNameInRedirect}}
		server_name_in_redirect {{$server.ServerNameInRedirect}};
		{{- end }}
		{{- if $server.SubrequestOutputBufferSize}}
		subrequest_output_buffer_size {{$server.SubrequestOutputBufferSize}};
		{{- end }}
		{{- if $server.TcpNodelay}}
		tcp_nodelay {{$server.TcpNodelay}};
		{{- end }}
		{{- if $server.TcpNopush}}
		tcp_nopush {{$server.TcpNopush}};
		{{- end }}
		{{- if $server.UnderscoresInHeaders}}
		underscores_in_headers {{$server.UnderscoresInHeaders}};
		{{- end }}
		{{- if $server.TypesHashBucketSize}}
		types_hash_bucket_size {{$server.TypesHashBucketSize}};
		{{- end }}
		{{- if $server.TypesHashMaxSize}}
		types_hash_max_size {{$server.TypesHashMaxSize}};
		{{- end }}
		{{- if $server.Aio}}
		aio {{$server.Aio}};
		{{- end }}
		{{- if $server.AbsoluteRedirect}}
		absolute_redirect {{$server.AbsoluteRedirect}};
		{{- end }}
		{{- if $server.AioWrite}}
		aio_write {{$server.AioWrite}};
		{{- end }}
		{{- if $server.ChunkedTransferEncoding}}
		chunked_transfer_encoding {{$server.ChunkedTransferEncoding}};
		{{- end }}
		{{- if $server.ClientBodyBufferSize}}
		client_body_bufer_size {{$server.ClientBodyBufferSize}};
		{{- end }}
		{{- if $server.ClientBodyInFileOnly}}
		client_body_in_file_only {{$server.ClientBodyInFileOnly}};
		{{- end }}
		{{- if $server.ClientBodyInSingleBuffer}}
		client_body_in_single_buffer {{$server.ClientBodyInSingleBuffer}};
		{{- end }}
		{{- if $server.ClientBodyTempPath}}
		client_body_temp_path {{$server.ClientBodyTempPath}};
		{{- end }}
		{{- if $server.ClientBodyTimeout}}
		client_body_timeout {{$server.ClientBodyTimeout}};
		{{- end }}
		{{- if $server.ClientHeaderBufferSize}}
		client_header_buffer_size {{$server.ClientHeaderBufferSize}};
		{{- end }}
		{{- if $server.ClientHeaderTimeout}}
		client_header_timeout {{$server.ClientHeaderTimeout}};
		{{- end }}
		{{- if $server.ClientMaxBodySize}}
		client_max_body_size {{$server.ClientMaxBodySize}};
		{{- end }}
		{{- if $server.ConnectionPoolSize}}
		conenction_pool_size {{$server.ConnectionPoolSize}};
		{{- end }}
		{{- if $server.DirectIO}}
		direct_io {{$server.DirectIO}};
		{{- end }}
		{{- if $server.DirectioAlignment}}
		directionio_alignment {{$server.DirectioAlignment}};
		{{- end }}
		{{- if $server.DisableSymlinks}}
		disable_symlinks {{$server.DisableSymlinks}};
		{{- end }}

		{{- if $server.IgnoreInvalidHeaders}}
		ignore_invalid_headers {{$server.IgnoreInvalidHeaders}};
		{{- end }}
		{{- if $server.AccessLog}}
		access_log {{$server.AccessLog}};
		{{- end }}
		{{- if $server.ProxyBind}}
		proxy_bind {{$server.ProxyBind}};
		{{- end }}
		{{- if $server.ProxyBufferSize}}
		proxy_buffer_size {{$server.ProxyBufferSize}};
		{{- end }}
		{{- if $server.ProxyBuffering}}
		proxy_buffering {{$server.ProxyBuffering}};
		{{- end }}
		{{- if $server.ProxyBuffers}}
		proxy_buffers {{$server.ProxyBuffers}};
		{{- end }}
		{{- if $server.ProxyBusyBuffersSize}}
		proxy_busy_buffers_size {{$server.ProxyBusyBuffersSize}};
		{{- end }}
		{{- if $server.ProxyCache}}
		proxy_cache {{$server.ProxyCache}};
		{{- end }}
		{{- if $server.ProxyCacheBackgroundUpdate}}
		proxy_cache_background_update {{$server.ProxyCacheBackgroundUpdate}};
		{{- end }}
		{{- if $server.ProxyCacheBypass}}
		{{range $proxyCacheBypass := $server.ProxyCacheBypass}}
		proxy_cache_byPass {{$proxyCacheBypass}};
		{{- end }}
		{{- end }}
		{{- if $server.ProxyCacheConvertHead}}
		proxy_cache_convert_head {{$server.ProxyCacheConvertHead}};
		{{- end }}
		{{- if $server.ProxyCacheKey}}
		proxy_cache_key {{$server.ProxyCacheKey}};
		{{- end }}
		{{- if $server.ProxyCacheLock}}
		proxy_cache_lock {{$server.ProxyCacheLock}};
		{{- end }}
		{{- if .ProxyCacheLockTimeout}}
		proxy_cache_lock_timeout {{$server.ProxyCacheLockTimeout}};
		{{- end }}
		{{- if $server.ProxyCacheMinUses}}
		proxy_cache_min_uses {{$server.ProxyCacheMinUses}};
		{{- end }}
		{{- if $server.ProxyCachePath}}
		proxy_cache_path {{$server.ProxyCachePath}};
		{{- end }}
		{{- if $server.ProxyCachePurge}}
		proxy_cache_purge {{$server.ProxyCachePurge}};
		{{- end }}
		{{- if $server.ProxyCacheUseStale}}
		proxy_cache_use_stale {{$server.ProxyCacheUseStale}};
		{{- end }}
		{{- if $server.ProxyCacheValid}}
		{{range $proxyCacheValid := $server.ProxyCacheValid}}
		proxy_cache_valid {{$proxyCacheValid}}
		{{- end }}
		{{- end }}
		{{- if $server.ProxyConnectTimeout}}
		proxy_connect_timeout {{$server.ProxyConnectionTimeout}};
		{{- end }}
		{{- if $server.ProxyCookieDomain}}
		proxy_coookie_domain {{$server.ProxyCookieDomain}};
		{{- end }}
		{{- if $server.ProxyCookiePath}}
		proxy_coookie_path {{$server.ProxyCookiePath}};
		{{- end }}
		{{- if $server.ProxyForceRanges}}
		proxy_force_ranges {{$server.ProxyForceRanges}};
		{{- end }}
		{{- if $server.ProxyHeadersHashBucketSize}}
		proxy_headers_hash_bucket_size {{$server.ProxyHeadersHashBucketSize}};
		{{- end }}
		{{- if $server.ProxyHeadersHashMaxSize}}
		proxy_headers_hash_max_size {{$server.ProxyHeadersHashMaxSize}};
		{{- end }}
		{{- if $server.ProxyHttpVersion}}
		proxy_http_version {{$server.ProxyHttpVersion}};
		{{- end }}
		{{- if $server.ProxyIgnoreClientAbort}}
		proxy_ignore_client_abort {{$server.ProxyIgnoreClientAbort}};
		{{- end }}
		{{- if $server.ProxyIgnoreHeaders}}
		proxy_ignore_headers {{$server.ProxyIgnoreHeaders}};
		{{- end }}
		{{- if $server.ProxyInterceptErrors}}
		proxy_intercept_errors {{$server.ProxyInterceptErrors}};
		{{- end }}
		{{- if $server.ProxyLimitRate}}
		proxy_limit_rate {{$server.ProxyLimitRate}};
		{{- end }}
		{{- if $server.ProxyMaxTempFileSize}}
		proxy_max_temp_file_size {{$server.ProxyMaxTempFileSize}};
		{{- end }}
		{{- if $server.ProxyMethod}}
		proxy_method {{$server.ProxyMethod}};
		{{- end }}
		{{- if $server.ProxyNextUpstream}}
		proxy_next_upstream {{$server.ProxyNextUpstream}};
		{{- end }}
		{{- if $server.ProxyNextUpstreamTimeout}}
		proxy_next_upstream_timeout {{$server.ProxyNextUpstreamTimeout}};
		{{- end }}
		{{- if $server.ProxyNextUpstreamTries}}
		proxy_next_upstream_tries {{$server.ProxyNextUpstreamTries}};
		{{- end }}
		{{- if $server.ProxyNoCache}}
		proxy_no_cache {{$server.ProxyNoCache}};
		{{- end }}
		{{- if $server.ProxyPass}}
		proxy_pass {{$server.ProxyPass}};
		{{- end }}
		{{- if $server.ProxyPassHeader}}
		proxy_pass_header {{$server.ProxyPassHeader}};
		{{- end }}
		{{- if $server.ProxyPassRequestBody}}
		proxy_pass_request_body {{$server.ProxyPassRequestBody}};
		{{- end }}
		{{- if $server.ProxyReadTimeout}}
		proxy_read_timeout {{$server.ProxyReadTimeout}};
		{{- end }}
		{{- if $server.ProxyRedirect}}
		proxy_redirect {{$server.ProxyRedirect}};
		{{- end }}
		{{- if $server.ProxyRequestBuffering}}
		proxy_request_buffering {{$server.ProxyRequestBuffering}};
		{{- end }}
		{{- if $server.ProxySendLowat}}
		proxy_send_lowat {{$server.ProxySendLowat}};
		{{- end }}
		{{- if $server.ProxySendTimeout}}
		proxy_send_timeout {{$server.ProxySendTimeout}};
		{{- end }}
		{{- if $server.ProxySetBody}}
		proxy_set_body {{$server.ProxySetBody}};
		{{- end }}
		{{- if $server.ProxySetHeader}}
		{{range $proxySetHeader := $server.ProxySetHeader}}
		proxy_set_header {{$proxySetHeader}};
		{{- end }}
		{{- end }}
		{{- if $server.ProxySocketKeepalive}}
		proxy_socket_keepalive {{$server.ProxySocketKeepalive}};
		{{- end }}
		{{- if $server.ProxySslCertificate}}
		proxy_ssl_certificate {{$server.ProxySslCertificate}};
		{{- end }}
		{{- if $server.ProxySslCertificate}}
		proxy_ssl_certificate {{$server.ProxySslCertificate}};
		{{- end }}
		{{- if $server.ProxySslCertificateKey}}
		proxy_ssl_certificate_key {{$server.ProxySslCertificateKey}};
		{{- end }}
		{{- if $server.ProxySslCiphers}}
		proxy_ssl_ciphers {{$server.ProxySslCiphers}};
		{{- end }}
		{{- if $server.ProxySslCrl}}
		proxy_ssl_crl {{$server.ProxySslCrl}};
		{{- end }}
		{{- if $server.ProxySslName}}
		proxy_ssl_name {{$server.ProxySslName}};
		{{- end }}
		{{- if $server.ProxySslPasswordFile}}
		proxy_ssl_password_file {{$server.ProxySslPasswordFile}};
		{{- end }}
		{{- if $server.ProxySslProtocols}}
		proxy_ssl_protocols {{$server.ProxySslProtocols}};
		{{- end }}
		{{- if $server.ProxySslServerName}}
		proxy_ssl_server_name {{$server.ProxySslServerName}};
		{{- end }}
		{{- if $server.ProxySslSessionReuse}}
		proxy_ssl_session_reuse {{$server.ProxySslSessionReuse}};
		{{- end }}
		{{- if $server.ProxyStore}}
		proxy_store {{$server.ProxyStore}};
		{{- end }}
		{{- if $server.ProxyStoreAccess}}
		proxy_store_access {{$server.ProxyStoreAccess}};
		{{- end }}
		{{- if $server.ProxyTempFileWriteSize}}
		proxy_temp_file_write_size {{$server.ProxyTempFileWriteSize}};
		{{- end }}
		{{- if $server.ProxyTempPath}}
		proxy_temp_path {{$server.ProxyTempPath}};
		{{- end }}
		{{- if .Gzip}}
        gzip {{.Gzip}};
        {{- end }}
    	{{- if .GzipMinLength}}
        gzip_min_length {{.GzipMinLength}};
        {{- end }}
    	{{- if .GzipBuffers}}
        gzip_buffers {{.GzipBuffers}};
        {{- end }}
    	{{- if .GzipCompLevel}}
        gzip_comp_level {{.GzipCompLevel}};
        {{- end }}
    	{{- if .GzipTypes}}
        gzip_types {{.GzipTypes}};
        {{- end }}
    	{{- if .GzipVary}}
        gzip_vary {{.GzipVary}};
        {{- end }}
    	{{- if .GzipDisable}}
        gzip_disable {{.GzipDisable}};
        {{- end }}
    	{{- if .GzipHttpVersion}}
        gzip_http_version {{.GzipHttpVersion}};
        {{- end }}
    	{{- if .GzipProxied}}
        gzip_proxied {{.GzipProxied}};
        {{- end }}
    	{{- if .ProxyHideHeader}}
        proxy_hide_header {{.ProxyHideHeader}};
        {{- end }}
		{{- if $server.Allow}}
		{{range $allow:= $server.Allow}}
		allow {{$allow}};
		{{- end }}
		{{- end }}
		{{- if $server.Deny}}
		{{range $deny:= $server.Deny}}
		deny {{$deny}};
		{{- end }}
		{{- end }}
		{{- if $server.Locations}}
		{{range $location := .Locations}}
			location  {{$location.Path}} {

				{{- if $location.AddHeader}}
				{{range $header := $location.AddHeader}}
				add_header {{$header}};
				{{- end }}
				{{- end }}
				{{- if $location.Aio}}
				aio {{$location.Aio}};
				{{- end }}
				{{- if $location.AioWrite}}
				aio_write {{$location.AioWrite}};
				{{- end }}
				{{- if $location.SendFile}}
				sendfile {{$location.SendFile}};
				{{- end }}
				{{- if $location.AccessLog}}
				access_log {{$location.AccessLog}};
				{{- end }}
				{{- if $location.Allow}}
				{{- range $allow := $location.Allow}}
				allow {{$allow}};
				{{- end }}
				{{- end }}
				{{- if $location.ClientBodyBufferSize}}
				client_body_buffer_size {{$location.ClientBodyBufferSize}};
				{{- end }}
				{{- if $location.ClientMaxBodySize}}
				client_max_body_size {{$location.ClientMaxBodySize}};
				{{- end }}
				{{- if $location.LimitConn}}
				limit_conn {{$location.LimitConn}};
				{{- end }}
				{{- if $location.LimitReq}}
				limit_req {{$location.LimitReq}};
				{{- end }}
				{{- if $location.LimitRateAfter}}
				limit_rate_after {{$location.LimitRateAfter}};
				{{- end }}
				{{- if $location.LimitRate}}
				limit_rate {{$location.LimitRate}};
				{{- end }}
				{{- if $location.Deny}}
				{{- range $deny := .Deny}}
				deny {{$deny}};
				{{- end }}
				{{- end }}
				{{- if $location.Expires}}
				expires {{$location.Expires}};
				{{- end }}
				{{- if $location.ProxyBufferSize}}
				proxy_buffer_size {{$location.ProxyBufferSize}};
				{{- end }}
				{{- if $location.ProxyBuffers}}
				proxy_buffers {{$location.ProxyBuffers}};
				{{- end }}
				{{- if $location.ProxyBusyBuffersSize}}
				proxy_busy_buffers_size {{$location.ProxyBusyBuffersSize}};
				{{- end }}
				{{- if $location.ProxyCache}}
				proxy_cache {{$location.ProxyCache}};
				{{- end }}
				{{- if $location.ProxyCacheKey}}
				proxy_cache_key {{$location.ProxyCacheKey}};
				{{- end }}
				{{- if .ProxyHttpVersion}}
				proxy_http_version {{.ProxyHttpVersion}};
				{{- end }}
				{{- if $location.ProxyCachePurge}}
				proxy_cache_purge {{$location.ProxyCachePurge}};
				{{- end }}
				{{- if $location.ProxyConnectTimeout}}
				proxy_connect_timeout {{$location.ProxyConnectTimeout}};
				{{- end }}
				{{- if $location.ProxyIgnoreClientAbort}}
				proxy_ignore_client_abort {{$location.ProxyIgnoreClientAbort}};
				{{- end }}
				{{- if $location.ProxyNextUpstream}}
				proxy_next_upstream {{$location.ProxyNextUpstream}};
				{{- end }}
				{{- if $location.ProxyPass}}
				proxy_pass {{$location.ProxyPass}};
				{{- end }}
				{{- if $location.ProxyReadTimeout}}
				proxy_read_timeout {{$location.ProxyReadTimeout}};
				{{- end }}
				{{-  if $location.ProxyRedirect}}
				proxy_redirect {{$location.ProxyRedirect}};
				{{- end }}
				{{- if $location.ProxySendTimeout}}
				proxy_send_timeout {{$location.ProxySendTimeout}};
				{{- end }}
				{{- if $location.ProxyTempFileWriteSize}}
				proxy_temp_file_write_size {{$location.ProxyTempFileWriteSize}};
				{{- end }}
				{{- if $location.Root}}
				root  {{$location.Root}};
				{{- end }}
				{{- if $location.StubStatus}}
				stub_status {{$location.StubStatus}};
				{{- end }}
				{{- if .Gzip}}
				gzip {{.Gzip}};
				{{- end }}
				{{- if .GzipMinLength}}
				gzip_min_length {{.GzipMinLength}};
				{{- end }}
				{{- if .GzipBuffers}}
				gzip_buffers {{.GzipBuffers}};
				{{- end }}
				{{- if .GzipCompLevel}}
				gzip_comp_level {{.GzipCompLevel}};
				{{- end }}
				{{- if .GzipTypes}}
				gzip_types {{.GzipTypes}};
				{{- end }}
				{{- if .GzipVary}}
				gzip_vary {{.GzipVary}};
				{{- end }}
				{{- if .GzipDisable}}
				gzip_disable {{.GzipDisable}};
				{{- end }}
				{{- if .GzipHttpVersion}}
				gzip_http_version {{.GzipHttpVersion}};
				{{- end }}
				{{- if .GzipProxied}}
				gzip_proxied {{.GzipProxied}};
				{{- end }}
				{{- if .ProxyHideHeader}}
				proxy_hide_header {{.ProxyHideHeader}};
				{{- end }}
				{{- if .Set}}
				{{range $set := .Set}}
				set {{$set}};
				{{- end }}
				{{- end }}
				{{- if $location.Rewrite}}
				{{- range $rewrite := .Rewrite}}
				rewrite {{$rewrite}};
				{{- end }}
				{{- end }}
				{{- if $location.ProxyCacheValid}}
				{{- range $proxyCacheValid := $location.ProxyCacheValid}}
				proxy_cache_valid {{$proxyCacheValid}};
				{{- end }}
				{{- end}}
				{{- if $location.ProxySetHeader}}
				{{- range $proxySetHeader := $location.ProxySetHeader}}
				proxy_set_header {{$proxySetHeader}};
				{{- end }}
				{{- end }}
				{{- if $location.IfBlocks}}
				{{- range $ifBlock := .IfBlocks}}
				if ({{.Condition}}) {
					{{- if $ifBlock.AddHeader}}
					{{range $header := $ifBlock.AddHeader}}
					add_header {{$header}};
					{{- end }}
					{{- end }}
					{{- if $ifBlock.ProxyPass}}
					proxy_pass {{$ifBlock.ProxyPass}};
					{{- end }}
					{{- if $ifBlock.SendFile}}
					send_file {{$ifBlock.SendFile}};
					{{- end }}
					{{- if $ifBlock.Root}}
					root {{$ifBlock.Root}};
					{{- end }}
					{{- if $ifBlock.Set}}
					{{range $set := $ifBlock.Set}}
					set {{$set}};
					{{- end }}
					{{- end }}
					{{range $rewrite := $ifBlock.Rewrite}}
					rewrite {{$rewrite}};
					{{- end }}
				}
				{{- end }}
				{{- end }}
			}
		{{- end }}
		{{- end }}
	}
	{{- end }}
}
{{- end }}
`
)
