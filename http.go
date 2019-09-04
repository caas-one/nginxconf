package nginx

import (
	"errors"
)

// determine whether it is a gzip_buffers directive
func isHTTPGzipBuffersDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipBuffersDirective) {
		return true
	}
	return false
}

// determine whether it is a output_buffers directive
func isHTTPOutputBuffersDirective(directive string) bool {
	if isEqualString(directive, HTTPOutputBuffersDirective) {
		return true
	}
	return false
}

// determine whether it is a request_pool_size directive
func isHTTPRequestPoolSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPRequestPoolSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_send_lowat directive
func isHTTPProxySendLowatDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySendLowatDirective) {
		return true
	}
	return false
}

// determine whether it is a absolute_redirect directive
func isHTTPAbsoluteRedirectDirective(directive string) bool {
	if isEqualString(directive, HTTPAbsoluteRedirectDirective) {
		return true
	}
	return false
}

// determine whether it is a directio_alignment directive
func isHTTPDirectioAlignmentDirective(directive string) bool {
	if isEqualString(directive, HTTPDirectioAlignmentDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_request_buffering directive
func isHTTPProxyRequestBufferingDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyRequestBufferingDirective) {
		return true
	}
	return false
}

// determine whether it is a log_subrequest directive
func isHTTPLogSubrequestDirective(directive string) bool {
	if isEqualString(directive, HTTPLogSubrequestDirective) {
		return true
	}
	return false
}

// determine whether it is a postpone_output directive
func isHTTPPostponeOutputDirective(directive string) bool {
	if isEqualString(directive, HTTPPostponeOutputDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_timeout directive
func isHTTPClientBodyTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPClientBodyTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a server directive
func isHTTPServersDirective(directive string) bool {
	if isEqualString(directive, HTTPServersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_force_ranges directive
func isHTTPProxyForceRangesDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyForceRangesDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip directive
func isHTTPGzipDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_conn_zone directive
func isHTTPLimitConnZoneDirective(directive string) bool {
	if isEqualString(directive, HTTPLimitConnZoneDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_req_zone directive
func isHTTPLimitReqZoneDirective(directive string) bool {
	if isEqualString(directive, HTTPLimitReqZoneDirective) {
		return true
	}
	return false
}

// determine whether it is a types_hash_max_size directive
func isHTTPTypesHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPTypesHashMaxSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_buffering directive
func isHTTPProxyBufferingDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyBufferingDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_crl directive
func isHTTPProxySslCrlDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslCrlDirective) {
		return true
	}
	return false
}

// determine whether it is a keepalive_timeout directive
func isHTTPKeepaliveTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPKeepaliveTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_rate directive
func isHTTPLimitRateDirective(directive string) bool {
	if isEqualString(directive, HTTPLimitRateDirective) {
		return true
	}
	return false
}

// determine whether it is a lingering_close directive
func isHTTPLingeringCloseDirective(directive string) bool {
	if isEqualString(directive, HTTPLingeringCloseDirective) {
		return true
	}
	return false
}

// determine whether it is a variables_hash_max_size directive
func isHTTPVariablesHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPVariablesHashMaxSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_in_file_only directive
func isHTTPClientBodyInFileOnlyDirective(directive string) bool {
	if isEqualString(directive, HTTPClientBodyInFileOnlyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_next_upstream directive
func isHTTPProxyNextUpstreamDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyNextUpstreamDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_trusted_certificate directive
func isHTTPProxySslTrustedCertificateDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslTrustedCertificateDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_temp_file_write_size directive
func isHTTPProxyTempFileWriteSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyTempFileWriteSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a max_ranges directive
func isHTTPMaxRangesDirective(directive string) bool {
	if isEqualString(directive, HTTPMaxRangesDirective) {
		return true
	}
	return false
}

// determine whether it is a variables_hash_bucket_size directive
func isHTTPVariablesHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPVariablesHashBucketSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_in_single_buffer directive
func isHTTPClientBodyInSingleBufferDirective(directive string) bool {
	if isEqualString(directive, HTTPClientBodyInSingleBufferDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_use_stale directive
func isHTTPProxyCacheUseStaleDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheUseStaleDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass_header directive
func isHTTPProxyPassHeaderDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyPassHeaderDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass_request_body directive
func isHTTPProxyPassRequestBodyDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyPassRequestBodyDirective) {
		return true
	}
	return false
}

// determine whether it is a sendfile_max_chunk directive
func isHTTPSendFileMaxChunkDirective(directive string) bool {
	if isEqualString(directive, HTTPSendFileMaxChunkDirective) {
		return true
	}
	return false
}

// determine whether it is a aio directive
func isHTTPAioDirective(directive string) bool {
	if isEqualString(directive, HTTPAioDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_min_uses directive
func isHTTPProxyCacheMinUsesDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheMinUsesDirective) {
		return true
	}
	return false
}

// determine whether it is a keepalive_disable directive
func isHTTPKeepaliveDisableDirective(directive string) bool {
	if isEqualString(directive, HTTPKeepaliveDisableDirective) {
		return true
	}
	return false
}

// determine whether it is a error_page directive
func isHTTPErrorPageDirective(directive string) bool {
	if isEqualString(directive, HTTPErrorPageDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_http_version directive
func isHTTPProxyHTTPVersionDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyHTTPVersionDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_verify_depth directive
func isHTTPProxySslVerifyDepthDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslVerifyDepthDirective) {
		return true
	}
	return false
}

// determine whether it is a server_name_in_redirect directive
func isHTTPServerNameInRedirectDirective(directive string) bool {
	if isEqualString(directive, HTTPServerNameInRedirectDirective) {
		return true
	}
	return false
}

// determine whether it is a server_names_hash_bucket_size directive
func isHTTPServerNamesHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPServerNamesHashBucketSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_buffer_size directive
func isHTTPClientBodyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPClientBodyBufferSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_valid directive
func isHTTPProxyCacheValidDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheValidDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_headers_hash_bucket_size directive
func isHTTPProxyHeadersHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyHeadersHashBucketSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_temp_path directive
func isHTTPProxyTempPathDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyTempPathDirective) {
		return true
	}
	return false
}

// determine whether it is a deny directive
func isHTTPDenyDirective(directive string) bool {
	if isEqualString(directive, HTTPDenyDirective) {
		return true
	}
	return false
}

// determine whether it is a send_timeout directive
func isHTTPSendTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPSendTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_max_range_offset directive
func isHTTPProxyCacheMaxRangeOffsetDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheMaxRangeOffsetDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_busy_buffers_size directive
func isHTTPProxyBusyBuffersSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyBusyBuffersSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_lock directive
func isHTTPProxyCacheLockDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheLockDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cookie_domain directive
func isHTTPProxyCookieDomainDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCookieDomainDirective) {
		return true
	}
	return false
}

// determine whether it is a open_file_cache_errors directive
func isHTTPOpenFileCacheErrorsDirective(directive string) bool {
	if isEqualString(directive, HTTPOpenFileCacheErrorsDirective) {
		return true
	}
	return false
}

// determine whether it is a open_file_cache_valid directive
func isHTTPOpenFileCacheValidDirective(directive string) bool {
	if isEqualString(directive, HTTPOpenFileCacheValidDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_background_update directive
func isHTTPProxyCacheBackgroundUpdateDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheBackgroundUpdateDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_intercept_errors directive
func isHTTPProxyInterceptErrorsDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyInterceptErrorsDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_store_access directive
func isHTTPProxyStoreAccessDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyStoreAccessDirective) {
		return true
	}
	return false
}

// determine whether it is a http directive
func isHTTPDirective(directive string) bool {
	if isEqualString(directive, HTTPDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_lock_timeout directive
func isHTTPProxyCacheLockTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheLockTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_send_timeout directive
func isHTTPProxySendTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySendTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a underscores_in_headers directive
func isHTTPUnderscoresInHeadersDirective(directive string) bool {
	if isEqualString(directive, HTTPUnderscoresInHeadersDirective) {
		return true
	}
	return false
}

// determine whether it is a connection_pool_size directive
func isHTTPConnectionPoolSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPConnectionPoolSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cookie_path directive
func isHTTPProxyCookiePathDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCookiePathDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_method directive
func isHTTPProxyMethodDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyMethodDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_read_timeout directive
func isHTTPProxyReadTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyReadTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_server_name directive
func isHTTPProxySslServerNameDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslServerNameDirective) {
		return true
	}
	return false
}

// determine whether it is a large_client_header_buffers directive
func isHTTPLargeClientHeaderBuffersDirective(directive string) bool {
	if isEqualString(directive, HTTPLargeClientHeaderBuffersDirective) {
		return true
	}
	return false
}

// determine whether it is a log_not_found directive
func isHTTPLogNotFoundDirective(directive string) bool {
	if isEqualString(directive, HTTPLogNotFoundDirective) {
		return true
	}
	return false
}

// determine whether it is a open_file_cache directive
func isHTTPOpenFileCacheDirective(directive string) bool {
	if isEqualString(directive, HTTPOpenFileCacheDirective) {
		return true
	}
	return false
}

// determine whether it is a recursive_error_pages directive
func isHTTPRecursiveErrorPagesDirective(directive string) bool {
	if isEqualString(directive, HTTPRecursiveErrorPagesDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_min_length directive
func isHTTPGzipMinLengthDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipMinLengthDirective) {
		return true
	}
	return false
}

// determine whether it is a msie_padding directive
func isHTTPMsiePaddingDirective(directive string) bool {
	if isEqualString(directive, HTTPMsiePaddingDirective) {
		return true
	}
	return false
}

// determine whether it is a msie_refresh directive
func isHTTPMsieRefreshDirective(directive string) bool {
	if isEqualString(directive, HTTPMsieRefreshDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_no_cache directive
func isHTTPProxyNoCacheDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyNoCacheDirective) {
		return true
	}
	return false
}

// determine whether it is a access_log directive
func isHTTPAccessLogDirective(directive string) bool {
	if isEqualString(directive, HTTPAccessLogDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_redirect directive
func isHTTPProxyRedirectDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyRedirectDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_set_header directive
func isHTTPProxySetHeaderDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySetHeaderDirective) {
		return true
	}
	return false
}

// determine whether it is a client_header_timeout directive
func isHTTPClientHeaderTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPClientHeaderTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a ignore_invalid_headers directive
func isHTTPIgnoreInvalidHeadersDirective(directive string) bool {
	if isEqualString(directive, HTTPIgnoreInvalidHeadersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_purge directive
func isHTTPProxyCachePurgeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCachePurgeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_max_temp_file_size directive
func isHTTPProxyMaxTempFileSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyMaxTempFileSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_store directive
func isHTTPProxyStoreDirectiveDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyStoreDirectiveDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_types directive
func isHTTPGzipTypesDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipTypesDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_buffer_size directive
func isHTTPProxyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyBufferSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a send_lowat directive
func isHTTPSendLowatDirective(directive string) bool {
	if isEqualString(directive, HTTPSendLowatDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_connect_timeout directive
func isHTTPProxyConnectTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyConnectTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_next_upstream_tries directive
func isHTTPProxyNextUpstreamTriesDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyNextUpstreamTriesDirective) {
		return true
	}
	return false
}

// determine whether it is a open_file_cache_min_uses directive
func isHTTPOpenFileCacheMinUsesDirective(directive string) bool {
	if isEqualString(directive, HTTPOpenFileCacheMinUsesDirective) {
		return true
	}
	return false
}

// determine whether it is a read_ahead directive
func isHTTPReadAheadDirective(directive string) bool {
	if isEqualString(directive, HTTPReadAheadDirective) {
		return true
	}
	return false
}

// determine whether it is a directio directive
func isHTTPDirectIODirective(directive string) bool {
	if isEqualString(directive, HTTPDirectIODirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ignore_headers directive
func isHTTPProxyIgnoreHeadersDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyIgnoreHeadersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_limit_rate directive
func isHTTPProxyLimitRateDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyLimitRateDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_socket_keepalive directive
func isHTTPProxySocketKeepaliveDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySocketKeepaliveDirective) {
		return true
	}
	return false
}

// determine whether it is a root directive
func isHTTPRootDirective(directive string) bool {
	if isEqualString(directive, HTTPRootDirective) {
		return true
	}
	return false
}

// determine whether it is a subrequest_output_buffer_size directive
func isHTTPSubrequestOutputBufferSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPSubrequestOutputBufferSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a types_hash_bucket_size directive
func isHTTPTypesHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPTypesHashBucketSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a disable_symlinks directive
func isHTTPDisableSymlinksDirective(directive string) bool {
	if isEqualString(directive, HTTPDisableSymlinksDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_convert_head directive
func isHTTPProxyCacheConvertHeadDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheConvertHeadDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_certificate directive
func isHTTPProxySslCertificateDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslCertificateDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_password_file directive
func isHTTPProxySslPasswordFile(directive string) bool {
	if isEqualString(directive, HTTPProxySslPasswordFile) {
		return true
	}
	return false
}

// determine whether it is a proxy_buffers directive
func isHTTPProxyBuffersDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyBuffersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_key directive
func isHTTPProxyCacheKeyDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheKeyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_certificate_key directive
func isHTTPProxySslCertificateKeyDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslCertificateKeyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_verify directive
func isHTTPProxySslVerifyDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslVerifyDirective) {
		return true
	}
	return false
}

// determine whether it is a cp_nodelay directive
func isHTTPTcpNodelayDirective(directive string) bool {
	if isEqualString(directive, HTTPTcpNodelayDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ignore_client_abort directive
func isHTTPProxyIgnoreClientAbortDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyIgnoreClientAbortDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass directive
func isHTTPProxyPassDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyPassDirective) {
		return true
	}
	return false
}

// determine whether it is a allow directive
func isHTTPAllowDirective(directive string) bool {
	if isEqualString(directive, HTTPAllowDirective) {
		return true
	}
	return false
}

// determine whether it is a lingering_time directive
func isHTTPLingeringTimeDirective(directive string) bool {
	if isEqualString(directive, HTTPLingeringTimeDirective) {
		return true
	}
	return false
}

// determine whether it is a tcp_nopush directive
func isHTTPTcpNopushDirective(directive string) bool {
	if isEqualString(directive, HTTPTcpNopushDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass_request_headers directive
func isHTTPProxyPassRequestHeadersDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyPassRequestHeadersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_name directive
func isHTTPProxySslNameDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslNameDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_next_upstream_timeout directive
func isHTTPProxyNextUpstreamTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyNextUpstreamTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a reset_timedout_connection directive
func isHTTPResetTimedoutConnectionDirective(directive string) bool {
	if isEqualString(directive, HTTPResetTimedoutConnectionDirective) {
		return true
	}
	return false
}

// determine whether it is a resolver directive
func isHTTPResolverDirective(directive string) bool {
	if isEqualString(directive, HTTPResolverDirective) {
		return true
	}
	return false
}

// determine whether it is a server_names_hash_max_size directive
func isHTTPServerNamesHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPServerNamesHashMaxSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a client_max_body_size directive
func isHTTPClientMaxBodySizeDirective(directive string) bool {
	if isEqualString(directive, HTTPClientMaxBodySizeDirective) {
		return true
	}
	return false
}

// determine whether it is a aio_write directive
func isHTTPAioWriteDirective(directive string) bool {
	if isEqualString(directive, HTTPAioWriteDirective) {
		return true
	}
	return false
}

// determine whether it is a client_header_buffer_size directive
func isHTTPClientHeaderBufferSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPClientHeaderBufferSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache directive
func isHTTPProxyCacheDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheDirective) {
		return true
	}
	return false
}

// determine whether it is a chunked_transfer_encoding directive
func isHTTPChunkedTransferEncodingDirective(directive string) bool {
	if isEqualString(directive, HTTPChunkedTransferEncodingDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_temp_path directive
func isHTTPClientBodyTempPathDirective(directive string) bool {
	if isEqualString(directive, HTTPClientBodyTempPathDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_bind directive
func isHTTPProxyBindDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyBindDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_lock_age directive
func isHTTPProxyCacheLockAgeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheLockAgeDirective) {
		return true
	}
	return false
}

// determine whether it is a lingering_timeout directive
func isHTTPLingeringTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPLingeringTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a satisfy directive
func isHTTPSatisfyDirective(directive string) bool {
	if isEqualString(directive, HTTPSatisfyDirective) {
		return true
	}
	return false
}

// determine whether it is a sendfile directive
func isHTTPSendFileDirective(directive string) bool {
	if isEqualString(directive, HTTPSendFileDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_methods directive
func isHTTPProxyCacheMethodsDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheMethodsDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_headers_hash_max_size directive
func isHTTPProxyHeadersHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyHeadersHashMaxSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_set_body directive
func isHTTPProxySetBodyDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySetBodyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_ciphers directive
func isHTTPProxySslCiphersDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslCiphersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_protocols directive
func isHTTPProxySslProtocolsDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslProtocolsDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_session_reuse directive
func isHTTPProxySslSessionReuseDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslSessionReuseDirective) {
		return true
	}
	return false
}

// determine whether it is a keepalive_requests directive
func isHTTPKeepaliveRequestsDirective(directive string) bool {
	if isEqualString(directive, HTTPKeepaliveRequestsDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_rate_after directive
func isHTTPLimitRateAfterDirective(directive string) bool {
	if isEqualString(directive, HTTPLimitRateAfterDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_bypass directive
func isHTTPProxyCacheBypassDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheBypassDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_path directive
func isHTTPProxyCachePathDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCachePathDirective) {
		return true
	}
	return false
}

// determine whether it is a geo directive
func isHTTPGeoDirective(directive string) bool {
	if isEqualString(directive, HTTPGeoDirective) {
		return true
	}
	return false
}

// determine whether it is a log_format directive
func isHTTPLogFormatDirective(directive string) bool {
	if isEqualString(directive, HTTPLogFormatDirective) {
		return true
	}
	return false
}

// determine whether it is a server directive
func isHTTPServerDirective(directive string) bool {
	if isEqualString(directive, HTTPServerDirective) {
		return true
	}
	return false
}

// determine whether it is a upsteam directive
func isHTTPUpstreamDirective(directive string) bool {
	if isEqualString(directive, HTTPUpstreamDirective) {
		return true
	}
	return false
}

// determine whether it is a default_type directive
func isHTTPDefaultTypeDirective(directive string) bool {
	if isEqualString(directive, HTTPDefaultTypeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_password_file directive
func isHTTPProxySslPasswordFileDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslPasswordFile) {
		return true
	}
	return false
}

// determine whether it is a proxy_store_access directive
func isHTTPProxyStoreDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyStoreAccessDirective) {
		return true
	}
	return false
}

// determine whether it is a include directive
func isHTTPIncludeDirective(directive string) bool {
	if isEqualString(directive, HTTPIncludeDirective) {
		return true
	}
	return false
}

// determine whether it is a add_header directive
func isHTTPAddHeaderDirective(directive string) bool {
	if isEqualString(directive, AddHeaderDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_comp_level directive
func isHTTPGzipCompLevelDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipCompLevelDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_vary directive
func isHTTPGzipVaryDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipVaryDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_disable directive
func isHTTPGzipDisableDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipDisableDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_http_version directive
func isHTTPGzipHTTPVersionDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipHTTPVersionDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_proxied directive
func isHTTPGzipProxiedDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipProxiedDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_hide_header directive
func isHTTPProxyHideHeaderDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyHideHeaderDirective) {
		return true
	}
	return false
}

// ProcessHTTP process Http module
func ProcessHTTP(parsed *Parsed) (*Http, error) {
	if !isHTTPDirective(parsed.Directive) {
		return nil, errors.New("Not http directive")
	}
	http := NewHttp()
	for _, block := range parsed.Blocks {

		// process gzip_types directive
		if isHTTPGzipTypesDirective(block.Directive) {
			http.GzipTypes = processArgsArray(block.Args)
			continue
		}

		// process log_not_found directive
		if isHTTPLogNotFoundDirective(block.Directive) {
			http.LogNotFound = processArgsArray(block.Args)
			continue
		}

		// process types_hash_max_size directive
		if isHTTPTypesHashMaxSizeDirective(block.Directive) {
			http.TypesHashMaxSize = processArgsArray(block.Args)
			continue
		}

		// process client_header_buffer_size directive
		if isHTTPClientHeaderBufferSizeDirective(block.Directive) {
			http.ClientHeaderBufferSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_headers_hash_max_size directive
		if isHTTPProxyHeadersHashMaxSizeDirective(block.Directive) {
			http.ProxyHeadersHashMaxSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_send_timeout directive
		if isHTTPProxySendTimeoutDirective(block.Directive) {
			http.ProxySendTimeout = processArgsArray(block.Args)
			continue
		}

		// process root directive
		if isHTTPRootDirective(block.Directive) {
			http.Root = processArgsArray(block.Args)
			continue
		}

		// process proxy_cookie_domain directive
		if isHTTPProxyCookieDomainDirective(block.Directive) {
			http.ProxyCookieDomain = processArgsArray(block.Args)
			continue
		}

		// process proxy_pass_request_headers directive
		if isHTTPProxyPassRequestHeadersDirective(block.Directive) {
			http.ProxyPassRequestHeaders = processArgsArray(block.Args)
			continue
		}

		// process proxy_send_lowat directive
		if isHTTPProxySendLowatDirective(block.Directive) {
			http.ProxySendLowat = processArgsArray(block.Args)
			continue
		}

		// process server_name_in_redirect directive
		if isHTTPServerNameInRedirectDirective(block.Directive) {
			http.ServerNameInRedirect = processArgsArray(block.Args)
			continue
		}

		// process types_hash_bucket_size directive
		if isHTTPTypesHashBucketSizeDirective(block.Directive) {
			http.TypesHashBucketSize = processArgsArray(block.Args)
			continue
		}

		// process client_body_temp_path directive
		if isHTTPClientBodyTempPathDirective(block.Directive) {
			http.ClientBodyTempPath = processArgsArray(block.Args)
			continue
		}

		// process proxy_buffers directive
		if isHTTPProxyBuffersDirective(block.Directive) {
			http.ProxyBuffers = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_convert_head on directive
		if isHTTPProxyCacheConvertHeadDirective(block.Directive) {
			http.ProxyCacheConvertHead = processArgsArray(block.Args)
			continue
		}

		// process proxy_socket_keepalive directive
		if isHTTPProxySocketKeepaliveDirective(block.Directive) {
			http.ProxySocketKeepalive = processArgsArray(block.Args)
			continue
		}

		// process proxy_hide_header directive
		if isHTTPProxyHideHeaderDirective(block.Directive) {
			http.ProxyHideHeader = processArgsArray(block.Args)
		}

		// process tcp_nopush directive
		if isHTTPTcpNopushDirective(block.Directive) {
			http.TcpNopush = processArgsArray(block.Args)
			continue
		}

		// process variables_hash_max_size directive
		if isHTTPVariablesHashMaxSizeDirective(block.Directive) {
			http.VariablesHashMaxSize = processArgsArray(block.Args)
			continue
		}

		// process client_body_timeout directive
		if isHTTPClientBodyTimeoutDirective(block.Directive) {
			http.ClientBodyTimeout = processArgsArray(block.Args)
			continue
		}

		// process directio_alignment directive
		if isHTTPDirectioAlignmentDirective(block.Directive) {
			http.DirectioAlignment = processArgsArray(block.Args)
			continue
		}

		// process proxy_busy_buffers_size directive
		if isHTTPProxyBusyBuffersSizeDirective(block.Directive) {
			http.ProxyBusyBuffersSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_password_file directive
		if isHTTPProxySslPasswordFileDirective(block.Directive) {
			http.ProxySslPasswordFile = processArgsArray(block.Args)
			continue
		}

		// process aio_write directive
		if isHTTPAioWriteDirective(block.Directive) {
			http.AioWrite = processArgsArray(block.Args)
			continue
		}

		// process chunked_transfer_encoding directive
		if isHTTPChunkedTransferEncodingDirective(block.Directive) {
			http.ChunkedTransferEncoding = processArgsArray(block.Args)
			continue
		}

		// process open_file_cache directive
		if isHTTPOpenFileCacheDirective(block.Directive) {
			http.OpenFileCache = processArgsArray(block.Args)
			continue
		}

		// process send_timeout directive
		if isHTTPSendTimeoutDirective(block.Directive) {
			http.SendTimeout = processArgsArray(block.Args)
			continue
		}

		// process client_body_buffer_size directive
		if isHTTPClientBodyBufferSizeDirective(block.Directive) {
			http.ClientBodyBufferSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache directive
		if isHTTPProxyCacheDirective(block.Directive) {
			http.ProxyCache = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_ciphers directive
		if isHTTPProxySslCiphersDirective(block.Directive) {
			http.ProxySslCiphers = processArgsArray(block.Args)
			continue
		}

		// process lingering_close directive
		if isHTTPLingeringCloseDirective(block.Directive) {
			http.LingeringClose = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_name directive
		if isHTTPProxySslNameDirective(block.Directive) {
			http.ProxySslName = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_protocols directive
		if isHTTPProxySslProtocolsDirective(block.Directive) {
			http.ProxySslProtocols = processArgsArray(block.Args)
			continue
		}

		// process max_ranges directive
		if isHTTPMaxRangesDirective(block.Directive) {
			http.MaxRanges = processArgsArray(block.Args)
			continue
		}

		// process request_pool_size directive
		if isHTTPRequestPoolSizeDirective(block.Directive) {
			http.RequestPoolSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_force_ranges directive
		if isHTTPProxyForceRangesDirective(block.Directive) {
			http.ProxyForceRanges = processArgsArray(block.Args)
			continue
		}

		// process proxy_max_temp_file_size directive
		if isHTTPProxyMaxTempFileSizeDirective(block.Directive) {
			http.ProxyMaxTempFileSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_pass directive
		if isHTTPProxyPassDirective(block.Directive) {
			http.ProxyPass = processArgsArray(block.Args)
			continue
		}

		// process proxy_store_access directive
		if isHTTPProxyStoreAccessDirective(block.Directive) {
			http.ProxyStoreAccess = processArgsArray(block.Args)
			continue
		}

		// process limit_rate directive
		if isHTTPLimitRateDirective(block.Directive) {
			http.LimitRate = processArgsArray(block.Args)
			continue
		}

		// process satisfy directive
		if isHTTPSatisfyDirective(block.Directive) {
			http.Satisfy = processArgsArray(block.Args)
			continue
		}

		// process sendfile directive
		if isHTTPSendFileDirective(block.Directive) {
			http.SendFile = processArgsArray(block.Args)
			continue
		}

		// process proxy_bind directive
		if isHTTPProxyBindDirective(block.Directive) {
			http.ProxyBind = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_lock directive
		if isHTTPProxyCacheLockDirective(block.Directive) {
			http.ProxyCacheLock = processArgsArray(block.Args)
			continue
		}

		// process disable_symlinks directive
		if isHTTPDisableSymlinksDirective(block.Directive) {
			http.DisableSymlinks = processArgsArray(block.Args)
			continue
		}

		// process proxy_temp_path directive
		if isHTTPProxyTempPathDirective(block.Directive) {
			http.ProxyTempPath = processArgsArray(block.Args)
			continue
		}

		// process gzip directive
		if isHTTPGzipDirective(block.Directive) {
			http.Gzip = processArgsArray(block.Args)
			continue
		}

		// process limit_conn_zone directive
		if isHTTPLimitConnZoneDirective(block.Directive) {
			http.LimitConnZone = processArgsArray(block.Args)
			continue
		}

		// process  limit_req_zone  directive
		if isHTTPLimitReqZoneDirective(block.Directive) {
			http.LimitReqZone = processArgsArray(block.Args)
			continue
		}

		// process open_file_cache_errors directive
		if isHTTPOpenFileCacheErrorsDirective(block.Directive) {
			http.OpenFileCacheErrors = processArgsArray(block.Args)
			continue
		}

		// process open_file_cache_min_uses directive
		if isHTTPOpenFileCacheMinUsesDirective(block.Directive) {
			http.OpenFileCacheMinUses = processArgsArray(block.Args)
			continue
		}

		// process read_ahead directive
		if isHTTPReadAheadDirective(block.Directive) {
			http.ReadAhead = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_verify_depth directive
		if isHTTPProxySslVerifyDepthDirective(block.Directive) {
			http.ProxySslVerifyDepth = processArgsArray(block.Args)
			continue
		}

		// process subrequest_output_buffer_size directive
		if isHTTPSubrequestOutputBufferSizeDirective(block.Directive) {
			http.SubrequestOutputBufferSize = processArgsArray(block.Args)
			continue
		}

		// process variables_hash_bucket_size directive
		if isHTTPVariablesHashBucketSizeDirective(block.Directive) {
			http.VariablesHashBucketSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_key directive
		if isHTTPProxyCacheKeyDirective(block.Directive) {
			http.ProxyCacheKey = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_purge directive
		if isHTTPProxyCachePurgeDirective(block.Directive) {
			http.ProxyCachePurge = processArgsArray(block.Args)
			continue
		}

		// process proxy_http_version directive
		if isHTTPProxyHTTPVersionDirective(block.Directive) {
			http.ProxyHttpVersion = processArgsArray(block.Args)
			continue
		}

		// process proxy_pass_request_body directive
		if isHTTPProxyPassRequestBodyDirective(block.Directive) {
			http.ProxyPassRequestBody = processArgsArray(block.Args)
			continue
		}

		// process gzip_buffers directive
		if isHTTPGzipBuffersDirective(block.Directive) {
			http.GzipBuffers = processArgsArray(block.Args)
			continue
		}

		// process client_body_in_file_only directive
		if isHTTPClientBodyInFileOnlyDirective(block.Directive) {
			http.ClientBodyInFileOnly = processArgsArray(block.Args)
			continue
		}

		// process keepalive_requests directive
		if isHTTPKeepaliveRequestsDirective(block.Directive) {
			http.KeepaliveRequests = processArgsArray(block.Args)
			continue
		}

		// process msie_refresh directive
		if isHTTPMsieRefreshDirective(block.Directive) {
			http.MsieRefresh = processArgsArray(block.Args)
			continue
		}

		// process client_max_body_size directive
		if isHTTPClientMaxBodySizeDirective(block.Directive) {
			http.ClientMaxBodySize = processArgsArray(block.Args)
			continue
		}

		// process proxy_method directive
		if isHTTPProxyMethodDirective(block.Directive) {
			http.ProxyMethod = processArgsArray(block.Args)
			continue
		}

		// process sendfile_max_chunk directive
		if isHTTPSendFileMaxChunkDirective(block.Directive) {
			http.SendFileMaxChunk = processArgsArray(block.Args)
			continue
		}

		// process client_body_in_single_buffer directive
		if isHTTPClientBodyInSingleBufferDirective(block.Directive) {
			http.ClientBodyInSingleBuffer = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_use_stale directive
		if isHTTPProxyCacheUseStaleDirective(block.Directive) {
			http.ProxyCacheUseStale = processArgsArray(block.Args)
			continue
		}

		// process proxy_read_timeout directive
		if isHTTPProxyReadTimeoutDirective(block.Directive) {
			http.ProxyReadTimeout = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_server_name directive
		if isHTTPProxySslServerNameDirective(block.Directive) {
			http.ProxySslServerName = processArgsArray(block.Args)
			continue
		}

		// process proxy_temp_file_write_size directive
		if isHTTPProxyTempFileWriteSizeDirective(block.Directive) {
			http.ProxyTempFileWriteSize = processArgsArray(block.Args)
			continue
		}

		// process aio directive
		if isHTTPAioDirective(block.Directive) {
			http.Aio = processArgsArray(block.Args)
			continue
		}

		// process proxy_cookie_path directive
		if isHTTPProxyCookiePathDirective(block.Directive) {
			http.ProxyCookiePath = processArgsArray(block.Args)
			continue
		}

		// process output_buffers directive
		if isHTTPOutputBuffersDirective(block.Directive) {
			http.OutputBuffers = processArgsArray(block.Args)
			continue
		}

		// process reset_timedout_connection directive
		if isHTTPResetTimedoutConnectionDirective(block.Directive) {
			http.ResetTimedoutConnection = processArgsArray(block.Args)
			continue
		}

		// process access_log directive
		if isHTTPAccessLogDirective(block.Directive) {
			http.AccessLog = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_path directive
		if isHTTPProxyCachePathDirective(block.Directive) {
			http.ProxyCachePath = processArgsArray(block.Args)
			continue
		}

		// process connection_pool_size directive
		if isHTTPConnectionPoolSizeDirective(block.Directive) {
			http.ConnectionPoolSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_next_upstream_tries directive
		if isHTTPProxyNextUpstreamTriesDirective(block.Directive) {
			http.ProxyNextUpstreamTries = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_verify directive
		if isHTTPProxySslVerifyDirective(block.Directive) {
			http.ProxySslVerify = processArgsArray(block.Args)
			continue
		}

		// process resolver directive
		if isHTTPResolverDirective(block.Directive) {
			http.Resolver = processArgsArray(block.Args)
			continue
		}

		// process absolute_redirect directive
		if isHTTPAbsoluteRedirectDirective(block.Directive) {
			http.AbsoluteRedirect = processArgsArray(block.Args)
			continue
		}

		// process proxy_buffer_size directive
		if isHTTPProxyBufferSizeDirective(block.Directive) {
			http.ProxyBufferSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_buffering directive
		if isHTTPProxyBufferingDirective(block.Directive) {
			http.ProxyBuffering = processArgsArray(block.Args)
			continue
		}

		// process lingering_time directive
		if isHTTPLingeringTimeDirective(block.Directive) {
			http.LingeringTime = processArgsArray(block.Args)
			continue
		}

		// process server_names_hash_max_size directive
		if isHTTPServerNamesHashMaxSizeDirective(block.Directive) {
			http.ServerNamesHashMaxSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_max_range_offset directive
		if isHTTPProxyCacheMaxRangeOffsetDirective(block.Directive) {
			http.ProxyCacheMaxRangeOffset = processArgsArray(block.Args)
			continue
		}

		// process keepalive_timeout directive
		if isHTTPKeepaliveTimeoutDirective(block.Directive) {
			http.KeepaliveTimeout = processArgsArray(block.Args)
			continue
		}

		// process msie_padding directive
		if isHTTPMsiePaddingDirective(block.Directive) {
			http.MsiePadding = processArgsArray(block.Args)
			continue
		}

		// process postpone_output directive
		if isHTTPPostponeOutputDirective(block.Directive) {
			http.PostponeOutput = processArgsArray(block.Args)
			continue
		}

		// process send_lowat directive
		if isHTTPSendLowatDirective(block.Directive) {
			http.SendLowat = processArgsArray(block.Args)
			continue
		}

		// process ignore_invalid_headers directive
		if isHTTPIgnoreInvalidHeadersDirective(block.Directive) {
			http.IgnoreInvalidHeaders = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_background_update directive
		if isHTTPProxyCacheBackgroundUpdateDirective(block.Directive) {
			http.ProxyCacheBackgroundUpdate = processArgsArray(block.Args)
			continue
		}

		// process proxy_ignore_headers directive
		if isHTTPProxyIgnoreHeadersDirective(block.Directive) {
			http.ProxyIgnoreHeaders = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_certificate directive
		if isHTTPProxySslCertificateDirective(block.Directive) {
			http.ProxySslCertificate = processArgsArray(block.Args)
			continue
		}

		// process underscores_in_headers directive
		if isHTTPUnderscoresInHeadersDirective(block.Directive) {
			http.UnderscoresInHeaders = processArgsArray(block.Args)
			continue
		}

		// process proxy_connect_timeout directive
		if isHTTPProxyConnectTimeoutDirective(block.Directive) {
			http.ProxyConnectTimeout = processArgsArray(block.Args)
			continue
		}

		// process proxy_intercept_errors directive
		if isHTTPProxyInterceptErrorsDirective(block.Directive) {
			http.ProxyInterceptErrors = processArgsArray(block.Args)
			continue
		}

		// process proxy_request_buffering directive
		if isHTTPProxyRequestBufferingDirective(block.Directive) {
			http.ProxyRequestBuffering = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_crl directive
		if isHTTPProxySslCrlDirective(block.Directive) {
			http.ProxySslCrl = processArgsArray(block.Args)
			continue
		}

		// process default_type directive
		if isHTTPDefaultTypeDirective(block.Directive) {
			http.DefaultType = processArgsArray(block.Args)
			continue
		}

		// process large_client_header_buffers directive
		if isHTTPLargeClientHeaderBuffersDirective(block.Directive) {
			http.LargeClientHeaderBuffers = processArgsArray(block.Args)
			continue
		}

		// process log_subrequest directive
		if isHTTPLogSubrequestDirective(block.Directive) {
			http.LogSubrequest = processArgsArray(block.Args)
			continue
		}

		// process client_header_timeout directive
		if isHTTPClientHeaderTimeoutDirective(block.Directive) {
			http.ClientHeaderTimeout = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_lock_timeout directive
		if isHTTPProxyCacheLockTimeoutDirective(block.Directive) {
			http.ProxyCacheLockTimeout = processArgsArray(block.Args)
			continue
		}

		// process proxy_ignore_client_abort directive
		if isHTTPProxyIgnoreClientAbortDirective(block.Directive) {
			http.ProxyIgnoreClientAbort = processArgsArray(block.Args)
			continue
		}

		// process proxy_set_body directive
		if isHTTPProxySetBodyDirective(block.Directive) {
			http.ProxySetBody = processArgsArray(block.Args)
			continue
		}

		// process cp_nodelay directive
		if isHTTPTcpNodelayDirective(block.Directive) {
			http.TcpNodelay = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_min_uses directive
		if isHTTPProxyCacheMinUsesDirective(block.Directive) {
			http.ProxyCacheMinUses = processArgsArray(block.Args)
			continue
		}

		// process keepalive_disable directive
		if isHTTPKeepaliveDisableDirective(block.Directive) {
			http.KeepaliveDisable = processArgsArray(block.Args)
			continue
		}

		// process lingering_timeout directive
		if isHTTPLingeringTimeoutDirective(block.Directive) {
			http.LingeringTimeout = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_methods directive
		if isHTTPProxyCacheMethodsDirective(block.Directive) {
			http.ProxyCacheMethods = processArgsArray(block.Args)
			continue
		}

		// process proxy_limit_rate directive
		if isHTTPProxyLimitRateDirective(block.Directive) {
			http.ProxyLimitRate = processArgsArray(block.Args)
			continue
		}

		// process directio directive
		if isHTTPDirectIODirective(block.Directive) {
			http.DirectIO = processArgsArray(block.Args)
			continue
		}

		// process error_page directive
		if isHTTPErrorPageDirective(block.Directive) {
			http.ErrorPage = processArgsArray(block.Args)
			continue
		}

		// process proxy_next_upstream directive
		if isHTTPProxyNextUpstreamDirective(block.Directive) {
			http.ProxyNextUpstream = processArgsArray(block.Args)
			continue
		}

		// process proxy_no_cache directive
		if isHTTPProxyNoCacheDirective(block.Directive) {
			http.ProxyNoCache = processArgsArray(block.Args)
			continue
		}

		// process proxy_redirect directive
		if isHTTPProxyRedirectDirective(block.Directive) {
			http.ProxyRedirect = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_certificate_key directive
		if isHTTPProxySslCertificateKeyDirective(block.Directive) {
			http.ProxySslCertificateKey = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_trusted_certificate directive
		if isHTTPProxySslTrustedCertificateDirective(block.Directive) {
			http.ProxySslTrustedCertificate = processArgsArray(block.Args)
			continue
		}

		// process recursive_error_pages directive
		if isHTTPRecursiveErrorPagesDirective(block.Directive) {
			http.RecursiveErrorPages = processArgsArray(block.Args)
			continue
		}

		// process limit_rate_after directive
		if isHTTPLimitRateAfterDirective(block.Directive) {
			http.LimitRateAfter = processArgsArray(block.Args)
			continue
		}

		// process server_names_hash_bucket_size directive
		if isHTTPServerNamesHashBucketSizeDirective(block.Directive) {
			http.ServerNamesHashBucketSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_next_upstream_timeout directive
		if isHTTPProxyNextUpstreamTimeoutDirective(block.Directive) {
			http.ProxyNextUpstreamTimeout = processArgsArray(block.Args)
			continue
		}

		// process proxy_pass_header directive
		if isHTTPProxyPassHeaderDirective(block.Directive) {
			http.ProxyPassHeader = processArgsArray(block.Args)
			continue
		}

		// process proxy_ssl_session_reuse directive
		if isHTTPProxySslSessionReuseDirective(block.Directive) {
			http.ProxySslSessionReuse = processArgsArray(block.Args)
			continue
		}

		// process gzip_min_length directive
		if isHTTPGzipMinLengthDirective(block.Directive) {
			http.GzipMinLength = processArgsArray(block.Args)
			continue
		}

		// process gzip_compl_level directive
		if isHTTPGzipCompLevelDirective(block.Directive) {
			http.GzipCompLevel = processArgsArray(block.Args)
			continue
		}

		// process gzip_types directive
		if isHTTPGzipTypesDirective(block.Directive) {
			http.GzipTypes = processArgsArray(block.Args)
			continue
		}

		// process gzip_vary directive
		if isHTTPGzipVaryDirective(block.Directive) {
			http.GzipVary = processArgsArray(block.Args)
			continue
		}

		// process gzip_disable directive
		if isHTTPGzipDisableDirective(block.Directive) {
			http.GzipDisable = processArgsArray(block.Args)
			continue
		}

		// process gzip_http_version directive
		if isHTTPGzipHTTPVersionDirective(block.Directive) {
			http.GzipHttpVersion = processArgsArray(block.Args)
			continue
		}

		// process gzip_proxied directive
		if isHTTPGzipProxiedDirective(block.Directive) {
			http.GzipProxied = processArgsArray(block.Args)
			continue
		}

		// process open_file_cache_valid directive
		if isHTTPOpenFileCacheValidDirective(block.Directive) {
			http.OpenFileCacheValid = processArgsArray(block.Args)
			continue
		}

		// process proxy_cache_lock_age directive
		if isHTTPProxyCacheLockAgeDirective(block.Directive) {
			http.ProxyCacheLockAge = processArgsArray(block.Args)
			continue
		}

		// process proxy_headers_hash_bucket_size directive
		if isHTTPProxyHeadersHashBucketSizeDirective(block.Directive) {
			http.ProxyHeadersHashBucketSize = processArgsArray(block.Args)
			continue
		}

		// process proxy_store directive
		if isHTTPProxyStoreDirective(block.Directive) {
			http.ProxyStore = processArgsArray(block.Args)
			continue
		}

		// process deny directive
		if isHTTPDenyDirective(block.Directive) {
			http.Deny = append(http.Deny, processArgsArray(block.Args))
		}

		// process proxy_cache_bypass directive
		if isHTTPProxyCacheBypassDirective(block.Directive) {
			http.ProxyCacheBypass = append(http.ProxyCacheBypass, processArgsArray(block.Args))
		}

		// process proxy_cache_valid directive
		if isHTTPProxyCacheValidDirective(block.Directive) {
			http.ProxyCacheValid = append(http.ProxyCacheValid, processArgsArray(block.Args))
		}

		// process proxy_set_header directive
		if isHTTPProxySetHeaderDirective(block.Directive) {
			http.ProxySetHeader = append(http.ProxySetHeader, processArgsArray(block.Args))
		}

		// process allow directive
		if isHTTPAllowDirective(block.Directive) {
			http.Allow = append(http.Allow, processArgsArray(block.Args))
		}

		// process include directive
		if isHTTPIncludeDirective(block.Directive) {
			http.Includes = append(http.Includes, processArgsArray(block.Args))
		}

		// process upstream directive
		if isHTTPUpstreamDirective(block.Directive) {
			upstream, err := ProcessUpstream(&block)
			if err != nil {
				return http, err
			}
			http.Upstream = append(http.Upstream, *upstream)
		}

		// process log_format directive
		if isHTTPLogFormatDirective(block.Directive) {
			http.LogFormat = append(http.LogFormat, processArgsArray(block.Args))
		}

		// process server directive
		if isHTTPServerDirective(block.Directive) {
			server, err := ProcessServer(&block)
			if err != nil {
				return http, err
			}
			http.Server = append(http.Server, *server)
		}

		// add_header
		if isHTTPAddHeaderDirective(block.Directive) {
			// add_header directive need to add quotation marks
			for i := range block.Args {
				block.Args[i] = nonescapeQuotation(block.Args[i])
			}
			args := processQuote(block.Args)
			http.AddHeader = append(http.AddHeader, processArgsArray(args))
		}

		// process geo directive
		if isHTTPGeoDirective(block.Directive) {
			geo, err := ProcessGeo(&block)
			if err != nil {
				return http, err
			}
			http.Geo = append(http.Geo, *geo)
		}

	}
	return http, nil
}
