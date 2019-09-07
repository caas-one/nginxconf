package temp

import (
	"fmt"
	"text/template"
)

// ServerTpl server template
var (
	serverTpl *template.Template
)

func init() {
	var err error
	serverTpl, err = template.New("server_tpl").Parse(ServerTemp)
	if err != nil {
		panic(fmt.Errorf("init server_tpl failed err:%v", err))
	}
}

// GetServerTpl get server template
func GetServerTpl() *template.Template {
	return serverTpl
}

// server template
const (
	ServerTemp = `server {
	{{- if .Name}}
	server_name {{.Name}};
	{{- end }}
	{{- if .Listen}}
	listen {{.Listen}};
	{{- end }}
	{{- if .ErrorPage}}
	error_page {{.ErrorPage}};
	{{- end }}
	{{- if .Return}}
	return {{.Return}};
	{{- end }}
	{{- if .Root}}
	root {{.Root}};
	{{- end }}
	{{- if .Ssl}}
	ssl {{.Ssl}};
	{{- end }}
	{{- if .SslCertificate}}
	ssl_certificate {{.SslCertificate}};
	{{- end }}
	{{- if .SslCertificateKey}}
	ssl_certificate_key {{.SslCertificateKey}};
	{{- end }}
	{{- if .SslProtocols}}
	ssl_protocols {{.SslProtocols}};
	{{- end }}
	{{- if .SslCiphers}}
	ssl_ciphers {{.SslCiphers}};
	{{- end }}
	{{- if .SslSessionTimeout}}
	ssl_session_timeout {{.SslSessionTimeout}};
	{{- end }}
	{{- if .LimitConn}}
	limit_conn {{.LimitConn}};
	{{- end }}
	{{- if .LimitReq}}
	limit_req {{.LimitReq}};
	{{- end }}
	{{- if .SslPreferServerCiphers}}
	ssl_prefer_server_ciphers {{.SslPreferServerCiphers}};
	{{- end }}
	{{- if .Set}}
    {{range $set := .Set}}
    set {{$set}};
    {{- end }}
    {{- end }}
	{{- if .Rewrite}}
	{{range $rewrite := .Rewrite}}
	rewrite {{$rewrite}};
	{{- end }}
	{{- end }}
	{{- if .IfBlocks}}
	{{range $ifBlocks := .IfBlocks}}
	if ({{.Condition}}) {
		{{- if .ProxyPass}}
		proxy_pass {{.ProxyPass}};
		{{- end }}
		{{- if .SendFile}}
		send_file {{.SendFile}};
		{{- end }}
		{{- if .Root}}
		root {{.Root}};
		{{- end }}
		{{- if .Set}}
		{{range $set := .Set}}
		set {{$set}};
		{{- end }}
		{{- end }}
		{{- if .Rewrite}}
		{{range $rewrite := .Rewrite}}
		rewrite {{$rewrite}};
		{{- end }}
		{{- end }}
		{{- if .Return}}
		return {{.Return}};
		{{- end }}
		{{- if .LimitRateAfter}}
		limit_rate_after {{.LimitRateAfter}};
		{{- end }}
	}
	{{- end }}
	{{- end }}
	
	{{- if .AccessLog}}
	access_log {{.AccessLog}};
	{{- end }}
	{{- if .KeepaliveDisable}}
	keepalive_disable {{.KeepAliveDisable}};
	{{- end }}
	{{- if .KeepaliveRequests}}
	keepalive_requests {{.KeepAliveRequests}};
	{{- end }}
	{{- if .KeepaliveTimeout}}
	keepalive_timeout {{.KeepAliveTimeout}};
	{{- end }}
	{{- if .LargeClientHeaderBuffers}}
	largeClientHeaderBuffers {{.LargeClientHeaderBuffers}};
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
	{{- if .MergeSlashes}}
	merge_slashes {{.MergeSlashes}};
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
	{{- if .OutputBuffers}}
	output_buffers {{.OutputBuffers}};
	{{- end }}
	{{- if .PostponeOutput}}
	postpone_output {{.PostponeOutput}};
	{{- end }}
	{{- if .ReadAhead}}
	read_ahead {{.ReadAhead}};
	{{- end }}
	{{- if .RequestPoolSize}}
	request_pool_size {{.RequestPoolSize}};
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
	{{- if .SendTimeout}}
	SendTimeout {{.SendTimeout}};
	{{- end }}
	{{- if .SendFile}}
	sendfile {{.SendFile}};
	{{- end }}
	{{- if .SendFileMaxChunk}}
	sendfile_max_chunk {{.SendFileMaxChunk}};
	{{- end }}
	{{- if .ServerNameInRedirect}}
	server_name_in_redirect {{.ServerNameInRedirect}};
	{{- end }}
	{{- if .SubrequestOutputBufferSize}}
	subrequest_output_buffer_size {{.SubrequestOutputBufferSize}};
	{{- end }}
	{{- if .TcpNodelay}}
	tcp_nodelay {{.TcpNodelay}};
	{{- end }}
	{{- if .TcpNopush}}
	tcp_nopush {{.TcpNopush}};
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
	client_body_bufer_size {{.ClientBodyBufferSize}};
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
	{{- if .ClientBodyTimeout}}
	client_body_timeout {{.ClientBodyTimeout}};
	{{- end }}
	{{- if .ClientHeaderBufferSize}}
	client_header_buffer_size {{.ClientHeaderBufferSize}};
	{{- end }}
	{{- if .ClientHeaderTimeout}}
	client_header_timeout {{.ClientHeaderTimeout}};
	{{- end }}
	{{- if .ClientMaxBodySize}}
	client_max_body_size {{.ClientMaxBodySize}};
	{{- end }}
	{{- if .ConnectionPoolSize}}
	connection_pool_size {{.ConnectionPoolSize}};
	{{- end }}
	{{- if .DirectIO}}
	direct_io {{.DirectIO}};
	{{- end }}
	{{- if .DirectioAlignment}}
	directionio_alignment {{.DirectioAlignment}};
	{{- end }}
	{{- if .DisableSymlinks}}
	disable_symlinks {{.DisableSymlinks}};
	{{- end }}
	
	{{- if .IgnoreInvalidHeaders}}
	ignore_invalid_headers {{.IgnoreInvalidHeaders}};
	{{- end }}
	{{- if .AccessLog}}
	access_log {{.AccessLog}};
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
	{{range $proxyCacheBypass := .ProxyCacheBypass}}
	proxy_cache_byPass {{$proxyCacheBypass}};
	{{- end }}
	{{- end }}
	{{- if .ProxyCacheConvertHead}}
	proxy_cache_convert_head {{.ProxyCacheConvertHead}};
	{{- end }}
	{{- if .ProxyCacheKey}}
	proxy_cache_key {{.ProxyCacheKey}};
	{{- end }}
	{{- if .ProxyCacheLock}}
	proxy_cache_lock {{.ProxyCacheLock}};
	{{- end }}
	{{- if .ProxyCacheLockTimeout}}
	proxy_cache_lock_timeout {{.ProxyCacheLockTimeout}};
	{{- end }}
	{{- if .ProxyCacheMinUses}}
	proxy_cache_min_uses {{.ProxyCacheMinUses}};
	{{- end }}
	{{- if .ProxyCachePath}}
	proxy_cache_path {{.ProxyCachePath}};
	{{- end }}
	{{- if .ProxyCachePurge}}
	proxy_cache_purge {{.ProxyCachePurge}};
	{{- end }}
	{{- if .ProxyCacheUseStale}}
	proxy_cache_use_stale {{.ProxyCacheUseStale}};
	{{- end }}
	{{- if .ProxyCacheValid}}
	{{range $proxyCacheValid := .ProxyCacheValid}}
	proxy_cache_valid {{$proxyCacheValid}}
	{{- end }}
	{{- end }}
	{{- if .ProxyConnectTimeout}}
	proxy_connect_timeout {{.ProxyConnectionTimeout}};
	{{- end }}
	{{- if .ProxyCookieDomain}}
	proxy_coookie_domain {{.ProxyCookieDomain}};
	{{- end }}
	{{- if .ProxyCookiePath}}
	proxy_coookie_path {{.ProxyCookiePath}};
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
	{{range $proxySetHeader := .ProxySetHeader}}
	proxy_set_header {{$proxySetHeader}};
	{{- end }}
	{{- end }}
	{{- if .ProxySocketKeepalive}}
	proxy_socket_keepalive {{.ProxySocketKeepalive}};
	{{- end }}
	{{- if .ProxySslCertificate}}
	proxy_ssl_certificate {{.ProxySslCertificate}};
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
	{{- if .ProxyStore}}
	proxy_store {{.ProxyStore}};
	{{- end }}
	{{- if .ProxyStoreAccess}}
	proxy_store_access {{.ProxyStoreAccess}};
	{{- end }}
	{{- if .ProxyTempFileWriteSize}}
	proxy_temp_file_write_size {{.ProxyTempFileWriteSize}};
	{{- end }}
	{{- if .ProxyTempPath}}
	proxy_temp_path {{.ProxyTempPath}};
	{{- end }}
	{{- if .Allow}}
	{{range $allow:= .Allow}}
	allow {{$allow}};
	{{- end }}
	{{- end }}
	{{- if .Deny}}
	{{range $deny:= .Deny}}
	deny {{$deny}};
	{{- end }}
	{{- end }}
	{{- if .AddHeader}}
	{{range $header := .AddHeader}}
	add_header {{$header}};
	{{- end }}
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
	{{- if .Locations}}
	{{range $locations:= .Locations}}
	location  {{.Path}} {
	
		{{- if .AccessLog}}
		access_log {{.AccessLog}};
		{{- end }}
		{{- if .Aio}}
		aio {{.Aio}};
		{{- end }}
		{{- if .AioWrite}}
		aio_write {{.AioWrite}};
		{{- end }}
		{{- if .SendFile}}
		sendfile {{.SendFile}};
		{{- end }}
		{{- if .Allow}}
		{{- range $allow := .Allow}}
		allow {{$allow}};
		{{- end }}
		{{- end }}
		{{- if .ClientBodyBufferSize}}
		client_body_buffer_size {{.ClientBodyBufferSize}};
		{{- end }}
		{{- if .ClientMaxBodySize}}
		client_max_body_size {{.ClientMaxBodySize}};
		{{- end }}
		{{- if .LimitConn}}
		limit_conn {{.LimitConn}};
		{{- end }}
		{{- if .LimitReq}}
		limit_req {{.LimitReq}};
		{{- end }}
		{{- if .LimitRate}}
		limit_rate {{.LimitRate}};
		{{- end }}
		{{- if .LimitRateAfter}}
		limit_rate_after {{.LimitRateAfter}};
		{{- end }}
		{{- if .Deny}}
		{{- range $deny := .Deny}}
		deny {{$deny}};
		{{- end }}
		{{- end }}
		{{- if .Expires}}
		expires {{.Expires}};
		{{- end }}
		{{- if .ProxyBufferSize}}
		proxy_buffer_size {{.ProxyBufferSize}};
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
		{{- if .ProxyCacheKey}}
		proxy_cache_key {{.ProxyCacheKey}};
		{{- end }}
		{{- if .ProxyCachePurge}}
		proxy_cache_purge {{.ProxyCachePurge}};
		{{- end }}
		{{- if .ProxyConnectTimeout}}
		proxy_connect_timeout {{.ProxyConnectTimeout}};
		{{- end }}
		{{- if .ProxyIgnoreClientAbort}}
		proxy_ignore_client_abort {{.ProxyIgnoreClientAbort}};
		{{- end }}
		{{- if .ProxyNextUpstream}}
		proxy_next_upstream {{.ProxyNextUpstream}};
		{{- end }}
		{{- if .ProxyPass}}
		proxy_pass {{.ProxyPass}};
		{{- end }}
		{{- if .ProxyReadTimeout}}
		proxy_read_timeout {{.ProxyReadTimeout}};
		{{- end }}
		{{-  if .ProxyRedirect}}
		proxy_redirect {{.ProxyRedirect}};
		{{- end }}
		{{- if .ProxyHttpVersion}}
		proxy_http_version {{.ProxyHttpVersion}};
		{{- end }}
		{{- if .ProxySendTimeout}}
		proxy_send_timeout {{.ProxySendTimeout}};
		{{- end }}
		{{- if .ProxyTempFileWriteSize}}
		proxy_temp_file_write_size {{.ProxyTempFileWriteSize}};
		{{- end }}
		{{- if .Root}}
		root  {{.Root}};
		{{- end }}
		{{- if .StubStatus}}
		stub_status {{.StubStatus}};
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
		{{- if .Rewrite}}
		{{- range $rewrite := .Rewrite}}
		rewrite {{$rewrite}};
		{{- end }}
		{{- end }}
		{{- if .ProxyCacheValid}}
		{{- range $proxyCacheValid := .ProxyCacheValid}}
		proxy_cache_valid {{$proxyCacheValid}};
		{{- end }}
		{{- end}}
		{{- if .ProxySetHeader}}
		{{- range $proxySetHeader := .ProxySetHeader}}
		proxy_set_header {{$proxySetHeader}};
		{{- end }}
		{{- end }}
		{{- if .AddHeader}}
		{{range $header := .AddHeader}}
		add_header {{$header}};
		{{- end }}
		{{- end }}
		{{- if .IfBlocks}}
		{{- range $ifBlocks := .IfBlocks}}
		if ({{.Condition}}) {
			{{- if .ProxyPass}}
			proxy_pass {{.ProxyPass}};
			{{- end }}
			{{- if .SendFile}}
			send_file {{.SendFile}};
			{{- end }}
			{{- if .Root}}
			root {{.Root}};
			{{- end }}
			{{- if $ifBlocks.Set}}
			{{range $set := $ifBlocks.Set}}
			set {{$set}};
			{{- end }}
			{{- end }}
			{{range $rewrite := .Rewrite}}
			rewrite {{$rewrite}};
			{{- end }}
			{{- if $ifBlocks.AddHeader}}
			{{range $header := $ifBlocks.AddHeader}}
			add_header {{$header}};
			{{- end }}
			{{- end }}
		}
		{{- end }}
		{{- end }}
	}
	{{- end }}
	{{- end }}
}`
)
