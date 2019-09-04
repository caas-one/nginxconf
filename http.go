package nginx

import (
	"errors"
)

// 判断是否是gzip_buffers指令
func isHTTPGzipBuffersDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipBuffersDirective) {
		return true
	}
	return false
}

// 判断是否是output_buffers指令
func isHTTPOutputBuffersDirective(directive string) bool {
	if isEqualString(directive, HTTPOutputBuffersDirective) {
		return true
	}
	return false
}

// 判断是否是request_pool_size指令
func isHTTPRequestPoolSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPRequestPoolSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_send_lowat指令
func isHTTPProxySendLowatDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySendLowatDirective) {
		return true
	}
	return false
}

// 判断是否是absolute_redirect指令
func isHTTPAbsoluteRedirectDirective(directive string) bool {
	if isEqualString(directive, HTTPAbsoluteRedirectDirective) {
		return true
	}
	return false
}

// 判断是否是directio_alignment指令
func isHTTPDirectioAlignmentDirective(directive string) bool {
	if isEqualString(directive, HTTPDirectioAlignmentDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_request_buffering指令
func isHTTPProxyRequestBufferingDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyRequestBufferingDirective) {
		return true
	}
	return false
}

// 判断是否是log_subrequest指令
func isHTTPLogSubrequestDirective(directive string) bool {
	if isEqualString(directive, HTTPLogSubrequestDirective) {
		return true
	}
	return false
}

// 判断是否是postpone_output指令
func isHTTPPostponeOutputDirective(directive string) bool {
	if isEqualString(directive, HTTPPostponeOutputDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_timeout指令
func isHTTPClientBodyTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPClientBodyTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是server指令
func isHTTPServersDirective(directive string) bool {
	if isEqualString(directive, HTTPServersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_force_ranges指令
func isHTTPProxyForceRangesDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyForceRangesDirective) {
		return true
	}
	return false
}

// 判断是否是gzip指令
func isHTTPGzipDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipDirective) {
		return true
	}
	return false
}

// 判断是否是limit_conn_zone指令
func isHTTPLimitConnZoneDirective(directive string) bool {
	if isEqualString(directive, HTTPLimitConnZoneDirective) {
		return true
	}
	return false
}

// 判断是否是limit_req_zone指令
func isHTTPLimitReqZoneDirective(directive string) bool {
	if isEqualString(directive, HTTPLimitReqZoneDirective) {
		return true
	}
	return false
}

// 判断是否是types_hash_max_size指令
func isHTTPTypesHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPTypesHashMaxSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_buffering指令
func isHTTPProxyBufferingDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyBufferingDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_crl指令
func isHTTPProxySslCrlDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslCrlDirective) {
		return true
	}
	return false
}

// 判断是否是keepalive_timeout指令
func isHTTPKeepaliveTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPKeepaliveTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是limit_rate指令
func isHTTPLimitRateDirective(directive string) bool {
	if isEqualString(directive, HTTPLimitRateDirective) {
		return true
	}
	return false
}

// 判断是否是lingering_close指令
func isHTTPLingeringCloseDirective(directive string) bool {
	if isEqualString(directive, HTTPLingeringCloseDirective) {
		return true
	}
	return false
}

// 判断是否是variables_hash_max_size指令
func isHTTPVariablesHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPVariablesHashMaxSizeDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_in_file_only指令
func isHTTPClientBodyInFileOnlyDirective(directive string) bool {
	if isEqualString(directive, HTTPClientBodyInFileOnlyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_next_upstream指令
func isHTTPProxyNextUpstreamDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyNextUpstreamDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_trusted_certificate指令
func isHTTPProxySslTrustedCertificateDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslTrustedCertificateDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_temp_file_write_size指令
func isHTTPProxyTempFileWriteSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyTempFileWriteSizeDirective) {
		return true
	}
	return false
}

// 判断是否是max_ranges指令
func isHTTPMaxRangesDirective(directive string) bool {
	if isEqualString(directive, HTTPMaxRangesDirective) {
		return true
	}
	return false
}

// 判断是否是variables_hash_bucket_size指令
func isHTTPVariablesHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPVariablesHashBucketSizeDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_in_single_buffer指令
func isHTTPClientBodyInSingleBufferDirective(directive string) bool {
	if isEqualString(directive, HTTPClientBodyInSingleBufferDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_use_stale指令
func isHTTPProxyCacheUseStaleDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheUseStaleDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass_header指令
func isHTTPProxyPassHeaderDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyPassHeaderDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass_request_body指令
func isHTTPProxyPassRequestBodyDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyPassRequestBodyDirective) {
		return true
	}
	return false
}

// 判断是否是sendfile_max_chunk指令
func isHTTPSendFileMaxChunkDirective(directive string) bool {
	if isEqualString(directive, HTTPSendFileMaxChunkDirective) {
		return true
	}
	return false
}

// 判断是否是aio指令
func isHTTPAioDirective(directive string) bool {
	if isEqualString(directive, HTTPAioDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_min_uses指令
func isHTTPProxyCacheMinUsesDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheMinUsesDirective) {
		return true
	}
	return false
}

// 判断是否是keepalive_disable指令
func isHTTPKeepaliveDisableDirective(directive string) bool {
	if isEqualString(directive, HTTPKeepaliveDisableDirective) {
		return true
	}
	return false
}

// 判断是否是error_page指令
func isHTTPErrorPageDirective(directive string) bool {
	if isEqualString(directive, HTTPErrorPageDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_http_version指令
func isHTTPProxyHTTPVersionDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyHTTPVersionDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_verify_depth指令
func isHTTPProxySslVerifyDepthDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslVerifyDepthDirective) {
		return true
	}
	return false
}

// 判断是否是server_name_in_redirect指令
func isHTTPServerNameInRedirectDirective(directive string) bool {
	if isEqualString(directive, HTTPServerNameInRedirectDirective) {
		return true
	}
	return false
}

// 判断是否是server_names_hash_bucket_size指令
func isHTTPServerNamesHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPServerNamesHashBucketSizeDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_buffer_size指令
func isHTTPClientBodyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPClientBodyBufferSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_valid指令
func isHTTPProxyCacheValidDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheValidDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_headers_hash_bucket_size指令
func isHTTPProxyHeadersHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyHeadersHashBucketSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_temp_path指令
func isHTTPProxyTempPathDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyTempPathDirective) {
		return true
	}
	return false
}

// 判断是否是deny指令
func isHTTPDenyDirective(directive string) bool {
	if isEqualString(directive, HTTPDenyDirective) {
		return true
	}
	return false
}

// 判断是否是send_timeout指令
func isHTTPSendTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPSendTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_max_range_offset指令
func isHTTPProxyCacheMaxRangeOffsetDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheMaxRangeOffsetDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_busy_buffers_size指令
func isHTTPProxyBusyBuffersSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyBusyBuffersSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_lock指令
func isHTTPProxyCacheLockDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheLockDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cookie_domain指令
func isHTTPProxyCookieDomainDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCookieDomainDirective) {
		return true
	}
	return false
}

// 判断是否是open_file_cache_errors指令
func isHTTPOpenFileCacheErrorsDirective(directive string) bool {
	if isEqualString(directive, HTTPOpenFileCacheErrorsDirective) {
		return true
	}
	return false
}

// 判断是否是open_file_cache_valid指令
func isHTTPOpenFileCacheValidDirective(directive string) bool {
	if isEqualString(directive, HTTPOpenFileCacheValidDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_background_update指令
func isHTTPProxyCacheBackgroundUpdateDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheBackgroundUpdateDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_intercept_errors指令
func isHTTPProxyInterceptErrorsDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyInterceptErrorsDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_store_access指令
func isHTTPProxyStoreAccessDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyStoreAccessDirective) {
		return true
	}
	return false
}

// 判断是否是http指令
func isHTTPDirective(directive string) bool {
	if isEqualString(directive, HTTPDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_lock_timeout指令
func isHTTPProxyCacheLockTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheLockTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_send_timeout指令
func isHTTPProxySendTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySendTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是underscores_in_headers指令
func isHTTPUnderscoresInHeadersDirective(directive string) bool {
	if isEqualString(directive, HTTPUnderscoresInHeadersDirective) {
		return true
	}
	return false
}

// 判断是否是connection_pool_size指令
func isHTTPConnectionPoolSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPConnectionPoolSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cookie_path指令
func isHTTPProxyCookiePathDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCookiePathDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_method指令
func isHTTPProxyMethodDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyMethodDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_read_timeout指令
func isHTTPProxyReadTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyReadTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_server_name指令
func isHTTPProxySslServerNameDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslServerNameDirective) {
		return true
	}
	return false
}

// 判断是否是large_client_header_buffers指令
func isHTTPLargeClientHeaderBuffersDirective(directive string) bool {
	if isEqualString(directive, HTTPLargeClientHeaderBuffersDirective) {
		return true
	}
	return false
}

// 判断是否是log_not_found指令
func isHTTPLogNotFoundDirective(directive string) bool {
	if isEqualString(directive, HTTPLogNotFoundDirective) {
		return true
	}
	return false
}

// 判断是否是open_file_cache指令
func isHTTPOpenFileCacheDirective(directive string) bool {
	if isEqualString(directive, HTTPOpenFileCacheDirective) {
		return true
	}
	return false
}

// 判断是否是recursive_error_pages指令
func isHTTPRecursiveErrorPagesDirective(directive string) bool {
	if isEqualString(directive, HTTPRecursiveErrorPagesDirective) {
		return true
	}
	return false
}

// 判断是否是gzip_min_length指令
func isHTTPGzipMinLengthDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipMinLengthDirective) {
		return true
	}
	return false
}

// 判断是否是msie_padding指令
func isHTTPMsiePaddingDirective(directive string) bool {
	if isEqualString(directive, HTTPMsiePaddingDirective) {
		return true
	}
	return false
}

// 判断是否是msie_refresh指令
func isHTTPMsieRefreshDirective(directive string) bool {
	if isEqualString(directive, HTTPMsieRefreshDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_no_cache指令
func isHTTPProxyNoCacheDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyNoCacheDirective) {
		return true
	}
	return false
}

// 判断是否是access_log指令
func isHTTPAccessLogDirective(directive string) bool {
	if isEqualString(directive, HTTPAccessLogDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_redirect指令
func isHTTPProxyRedirectDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyRedirectDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_set_header指令
func isHTTPProxySetHeaderDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySetHeaderDirective) {
		return true
	}
	return false
}

// 判断是否是client_header_timeout指令
func isHTTPClientHeaderTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPClientHeaderTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是ignore_invalid_headers指令
func isHTTPIgnoreInvalidHeadersDirective(directive string) bool {
	if isEqualString(directive, HTTPIgnoreInvalidHeadersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_purge指令
func isHTTPProxyCachePurgeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCachePurgeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_max_temp_file_size指令
func isHTTPProxyMaxTempFileSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyMaxTempFileSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_store指令
func isHTTPProxyStoreDirectiveDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyStoreDirectiveDirective) {
		return true
	}
	return false
}

// 判断是否是gzip_types指令
func isHTTPGzipTypesDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipTypesDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_buffer_size指令
func isHTTPProxyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyBufferSizeDirective) {
		return true
	}
	return false
}

// 判断是否是send_lowat指令
func isHTTPSendLowatDirective(directive string) bool {
	if isEqualString(directive, HTTPSendLowatDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_connect_timeout指令
func isHTTPProxyConnectTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyConnectTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_next_upstream_tries指令
func isHTTPProxyNextUpstreamTriesDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyNextUpstreamTriesDirective) {
		return true
	}
	return false
}

// 判断是否是open_file_cache_min_uses指令
func isHTTPOpenFileCacheMinUsesDirective(directive string) bool {
	if isEqualString(directive, HTTPOpenFileCacheMinUsesDirective) {
		return true
	}
	return false
}

// 判断是否是read_ahead指令
func isHTTPReadAheadDirective(directive string) bool {
	if isEqualString(directive, HTTPReadAheadDirective) {
		return true
	}
	return false
}

// 判断是否是directio指令
func isHTTPDirectIODirective(directive string) bool {
	if isEqualString(directive, HTTPDirectIODirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ignore_headers指令
func isHTTPProxyIgnoreHeadersDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyIgnoreHeadersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_limit_rate指令
func isHTTPProxyLimitRateDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyLimitRateDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_socket_keepalive指令
func isHTTPProxySocketKeepaliveDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySocketKeepaliveDirective) {
		return true
	}
	return false
}

// 判断是否是root指令
func isHTTPRootDirective(directive string) bool {
	if isEqualString(directive, HTTPRootDirective) {
		return true
	}
	return false
}

// 判断是否是subrequest_output_buffer_size指令
func isHTTPSubrequestOutputBufferSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPSubrequestOutputBufferSizeDirective) {
		return true
	}
	return false
}

// 判断是否是types_hash_bucket_size指令
func isHTTPTypesHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPTypesHashBucketSizeDirective) {
		return true
	}
	return false
}

// 判断是否是disable_symlinks指令
func isHTTPDisableSymlinksDirective(directive string) bool {
	if isEqualString(directive, HTTPDisableSymlinksDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_convert_head指令
func isHTTPProxyCacheConvertHeadDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheConvertHeadDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_certificate指令
func isHTTPProxySslCertificateDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslCertificateDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_password_file指令
func isHTTPProxySslPasswordFile(directive string) bool {
	if isEqualString(directive, HTTPProxySslPasswordFile) {
		return true
	}
	return false
}

// 判断是否是proxy_buffers指令
func isHTTPProxyBuffersDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyBuffersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_key指令
func isHTTPProxyCacheKeyDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheKeyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_certificate_key指令
func isHTTPProxySslCertificateKeyDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslCertificateKeyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_verify指令
func isHTTPProxySslVerifyDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslVerifyDirective) {
		return true
	}
	return false
}

// 判断是否是cp_nodelay指令
func isHTTPTcpNodelayDirective(directive string) bool {
	if isEqualString(directive, HTTPTcpNodelayDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ignore_client_abort指令
func isHTTPProxyIgnoreClientAbortDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyIgnoreClientAbortDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass指令
func isHTTPProxyPassDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyPassDirective) {
		return true
	}
	return false
}

// 判断是否是allow指令
func isHTTPAllowDirective(directive string) bool {
	if isEqualString(directive, HTTPAllowDirective) {
		return true
	}
	return false
}

// 判断是否是lingering_time指令
func isHTTPLingeringTimeDirective(directive string) bool {
	if isEqualString(directive, HTTPLingeringTimeDirective) {
		return true
	}
	return false
}

// 判断是否是tcp_nopush指令
func isHTTPTcpNopushDirective(directive string) bool {
	if isEqualString(directive, HTTPTcpNopushDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass_request_headers指令
func isHTTPProxyPassRequestHeadersDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyPassRequestHeadersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_name指令
func isHTTPProxySslNameDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslNameDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_next_upstream_timeout指令
func isHTTPProxyNextUpstreamTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyNextUpstreamTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是reset_timedout_connection指令
func isHTTPResetTimedoutConnectionDirective(directive string) bool {
	if isEqualString(directive, HTTPResetTimedoutConnectionDirective) {
		return true
	}
	return false
}

// 判断是否是resolver指令
func isHTTPResolverDirective(directive string) bool {
	if isEqualString(directive, HTTPResolverDirective) {
		return true
	}
	return false
}

// 判断是否是server_names_hash_max_size指令
func isHTTPServerNamesHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPServerNamesHashMaxSizeDirective) {
		return true
	}
	return false
}

// 判断是否是client_max_body_size指令
func isHTTPClientMaxBodySizeDirective(directive string) bool {
	if isEqualString(directive, HTTPClientMaxBodySizeDirective) {
		return true
	}
	return false
}

// 判断是否是aio_write指令
func isHTTPAioWriteDirective(directive string) bool {
	if isEqualString(directive, HTTPAioWriteDirective) {
		return true
	}
	return false
}

// 判断是否是client_header_buffer_size指令
func isHTTPClientHeaderBufferSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPClientHeaderBufferSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache指令
func isHTTPProxyCacheDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheDirective) {
		return true
	}
	return false
}

// 判断是否是chunked_transfer_encoding指令
func isHTTPChunkedTransferEncodingDirective(directive string) bool {
	if isEqualString(directive, HTTPChunkedTransferEncodingDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_temp_path指令
func isHTTPClientBodyTempPathDirective(directive string) bool {
	if isEqualString(directive, HTTPClientBodyTempPathDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_bind指令
func isHTTPProxyBindDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyBindDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_lock_age指令
func isHTTPProxyCacheLockAgeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheLockAgeDirective) {
		return true
	}
	return false
}

// 判断是否是lingering_timeout指令
func isHTTPLingeringTimeoutDirective(directive string) bool {
	if isEqualString(directive, HTTPLingeringTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是satisfy指令
func isHTTPSatisfyDirective(directive string) bool {
	if isEqualString(directive, HTTPSatisfyDirective) {
		return true
	}
	return false
}

// 判断是否是sendfile指令
func isHTTPSendFileDirective(directive string) bool {
	if isEqualString(directive, HTTPSendFileDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_methods指令
func isHTTPProxyCacheMethodsDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheMethodsDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_headers_hash_max_size指令
func isHTTPProxyHeadersHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyHeadersHashMaxSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_set_body指令
func isHTTPProxySetBodyDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySetBodyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_ciphers指令
func isHTTPProxySslCiphersDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslCiphersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_protocols指令
func isHTTPProxySslProtocolsDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslProtocolsDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_session_reuse指令
func isHTTPProxySslSessionReuseDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslSessionReuseDirective) {
		return true
	}
	return false
}

// 判断是否是keepalive_requests指令
func isHTTPKeepaliveRequestsDirective(directive string) bool {
	if isEqualString(directive, HTTPKeepaliveRequestsDirective) {
		return true
	}
	return false
}

// 判断是否是limit_rate_after指令
func isHTTPLimitRateAfterDirective(directive string) bool {
	if isEqualString(directive, HTTPLimitRateAfterDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_bypass指令
func isHTTPProxyCacheBypassDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCacheBypassDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_path指令
func isHTTPProxyCachePathDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyCachePathDirective) {
		return true
	}
	return false
}

// 判断是否是geo指令
func isHTTPGeoDirective(directive string) bool {
	if isEqualString(directive, HTTPGeoDirective) {
		return true
	}
	return false
}

// 判断是否是log_format指令
func isHTTPLogFormatDirective(directive string) bool {
	if isEqualString(directive, HTTPLogFormatDirective) {
		return true
	}
	return false
}

// 判断是否是server指令
func isHTTPServerDirective(directive string) bool {
	if isEqualString(directive, HTTPServerDirective) {
		return true
	}
	return false
}

// 判断是否是upsteam指令
func isHTTPUpstreamDirective(directive string) bool {
	if isEqualString(directive, HTTPUpstreamDirective) {
		return true
	}
	return false
}

// 判断是否是default_type指令
func isHTTPDefaultTypeDirective(directive string) bool {
	if isEqualString(directive, HTTPDefaultTypeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_password_file指令
func isHTTPProxySslPasswordFileDirective(directive string) bool {
	if isEqualString(directive, HTTPProxySslPasswordFile) {
		return true
	}
	return false
}

// 判断是否是proxy_store_access指令
func isHTTPProxyStoreDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyStoreAccessDirective) {
		return true
	}
	return false
}

// 判断是否是include指令
func isHTTPIncludeDirective(directive string) bool {
	if isEqualString(directive, HTTPIncludeDirective) {
		return true
	}
	return false
}

// 判断是否为add_header指令
func isHTTPAddHeaderDirective(directive string) bool {
	if isEqualString(directive, AddHeaderDirective) {
		return true
	}
	return false
}

// 判断是否为gzip_comp_level指令
func isHTTPGzipCompLevelDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipCompLevelDirective) {
		return true
	}
	return false
}

// 判断是否为gzip_vary指令
func isHTTPGzipVaryDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipVaryDirective) {
		return true
	}
	return false
}

// 判断是否为gzip_disable指令
func isHTTPGzipDisableDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipDisableDirective) {
		return true
	}
	return false
}

// 判断是否为gzip_http_version指令
func isHTTPGzipHttpVersionDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipHTTPVersionDirective) {
		return true
	}
	return false
}

// 判断是否为gzip_proxied指令
func isHTTPGzipProxiedDirective(directive string) bool {
	if isEqualString(directive, HTTPGzipProxiedDirective) {
		return true
	}
	return false
}

// 判断是否为proxy_hide_header指令
func isHTTPProxyHideHeaderDirective(directive string) bool {
	if isEqualString(directive, HTTPProxyHideHeaderDirective) {
		return true
	}
	return false
}

// 处理Http部分指令
func ProcessHttp(parsed *Parsed) (*Http, error) {
	if !isHTTPDirective(parsed.Directive) {
		return nil, errors.New("Not http directive")
	}
	http := NewHttp()
	for _, block := range parsed.Blocks {

		// 处理gzip_types指令
		if isHTTPGzipTypesDirective(block.Directive) {
			http.GzipTypes = processArgsArray(block.Args)
			continue
		}

		// 处理log_not_found指令
		if isHTTPLogNotFoundDirective(block.Directive) {
			http.LogNotFound = processArgsArray(block.Args)
			continue
		}

		// 处理types_hash_max_size指令
		if isHTTPTypesHashMaxSizeDirective(block.Directive) {
			http.TypesHashMaxSize = processArgsArray(block.Args)
			continue
		}

		// 处理client_header_buffer_size指令
		if isHTTPClientHeaderBufferSizeDirective(block.Directive) {
			http.ClientHeaderBufferSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_headers_hash_max_size指令
		if isHTTPProxyHeadersHashMaxSizeDirective(block.Directive) {
			http.ProxyHeadersHashMaxSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_send_timeout指令
		if isHTTPProxySendTimeoutDirective(block.Directive) {
			http.ProxySendTimeout = processArgsArray(block.Args)
			continue
		}

		// 处理root指令
		if isHTTPRootDirective(block.Directive) {
			http.Root = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cookie_domain指令
		if isHTTPProxyCookieDomainDirective(block.Directive) {
			http.ProxyCookieDomain = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_pass_request_headers指令
		if isHTTPProxyPassRequestHeadersDirective(block.Directive) {
			http.ProxyPassRequestHeaders = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_send_lowat指令
		if isHTTPProxySendLowatDirective(block.Directive) {
			http.ProxySendLowat = processArgsArray(block.Args)
			continue
		}

		// 处理server_name_in_redirect指令
		if isHTTPServerNameInRedirectDirective(block.Directive) {
			http.ServerNameInRedirect = processArgsArray(block.Args)
			continue
		}

		// 处理types_hash_bucket_size指令
		if isHTTPTypesHashBucketSizeDirective(block.Directive) {
			http.TypesHashBucketSize = processArgsArray(block.Args)
			continue
		}

		// 处理client_body_temp_path指令
		if isHTTPClientBodyTempPathDirective(block.Directive) {
			http.ClientBodyTempPath = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_buffers指令
		if isHTTPProxyBuffersDirective(block.Directive) {
			http.ProxyBuffers = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_convert_head on指令
		if isHTTPProxyCacheConvertHeadDirective(block.Directive) {
			http.ProxyCacheConvertHead = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_socket_keepalive指令
		if isHTTPProxySocketKeepaliveDirective(block.Directive) {
			http.ProxySocketKeepalive = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_hide_header指令
		if isHTTPProxyHideHeaderDirective(block.Directive) {
			http.ProxyHideHeader = processArgsArray(block.Args)
		}

		// 处理tcp_nopush指令
		if isHTTPTcpNopushDirective(block.Directive) {
			http.TcpNopush = processArgsArray(block.Args)
			continue
		}

		// 处理variables_hash_max_size指令
		if isHTTPVariablesHashMaxSizeDirective(block.Directive) {
			http.VariablesHashMaxSize = processArgsArray(block.Args)
			continue
		}

		// 处理client_body_timeout指令
		if isHTTPClientBodyTimeoutDirective(block.Directive) {
			http.ClientBodyTimeout = processArgsArray(block.Args)
			continue
		}

		// 处理directio_alignment指令
		if isHTTPDirectioAlignmentDirective(block.Directive) {
			http.DirectioAlignment = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_busy_buffers_size指令
		if isHTTPProxyBusyBuffersSizeDirective(block.Directive) {
			http.ProxyBusyBuffersSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_password_file指令
		if isHTTPProxySslPasswordFileDirective(block.Directive) {
			http.ProxySslPasswordFile = processArgsArray(block.Args)
			continue
		}

		// 处理aio_write指令
		if isHTTPAioWriteDirective(block.Directive) {
			http.AioWrite = processArgsArray(block.Args)
			continue
		}

		// 处理chunked_transfer_encoding指令
		if isHTTPChunkedTransferEncodingDirective(block.Directive) {
			http.ChunkedTransferEncoding = processArgsArray(block.Args)
			continue
		}

		// 处理open_file_cache指令
		if isHTTPOpenFileCacheDirective(block.Directive) {
			http.OpenFileCache = processArgsArray(block.Args)
			continue
		}

		// 处理send_timeout指令
		if isHTTPSendTimeoutDirective(block.Directive) {
			http.SendTimeout = processArgsArray(block.Args)
			continue
		}

		// 处理client_body_buffer_size指令
		if isHTTPClientBodyBufferSizeDirective(block.Directive) {
			http.ClientBodyBufferSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache指令
		if isHTTPProxyCacheDirective(block.Directive) {
			http.ProxyCache = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_ciphers指令
		if isHTTPProxySslCiphersDirective(block.Directive) {
			http.ProxySslCiphers = processArgsArray(block.Args)
			continue
		}

		// 处理lingering_close指令
		if isHTTPLingeringCloseDirective(block.Directive) {
			http.LingeringClose = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_name指令
		if isHTTPProxySslNameDirective(block.Directive) {
			http.ProxySslName = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_protocols指令
		if isHTTPProxySslProtocolsDirective(block.Directive) {
			http.ProxySslProtocols = processArgsArray(block.Args)
			continue
		}

		// 处理max_ranges指令
		if isHTTPMaxRangesDirective(block.Directive) {
			http.MaxRanges = processArgsArray(block.Args)
			continue
		}

		// 处理request_pool_size指令
		if isHTTPRequestPoolSizeDirective(block.Directive) {
			http.RequestPoolSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_force_ranges指令
		if isHTTPProxyForceRangesDirective(block.Directive) {
			http.ProxyForceRanges = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_max_temp_file_size指令
		if isHTTPProxyMaxTempFileSizeDirective(block.Directive) {
			http.ProxyMaxTempFileSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_pass指令
		if isHTTPProxyPassDirective(block.Directive) {
			http.ProxyPass = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_store_access指令
		if isHTTPProxyStoreAccessDirective(block.Directive) {
			http.ProxyStoreAccess = processArgsArray(block.Args)
			continue
		}

		// 处理limit_rate指令
		if isHTTPLimitRateDirective(block.Directive) {
			http.LimitRate = processArgsArray(block.Args)
			continue
		}

		// 处理satisfy指令
		if isHTTPSatisfyDirective(block.Directive) {
			http.Satisfy = processArgsArray(block.Args)
			continue
		}

		// 处理sendfile指令
		if isHTTPSendFileDirective(block.Directive) {
			http.SendFile = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_bind指令
		if isHTTPProxyBindDirective(block.Directive) {
			http.ProxyBind = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_lock指令
		if isHTTPProxyCacheLockDirective(block.Directive) {
			http.ProxyCacheLock = processArgsArray(block.Args)
			continue
		}

		// 处理disable_symlinks指令
		if isHTTPDisableSymlinksDirective(block.Directive) {
			http.DisableSymlinks = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_temp_path指令
		if isHTTPProxyTempPathDirective(block.Directive) {
			http.ProxyTempPath = processArgsArray(block.Args)
			continue
		}

		// 处理gzip指令
		if isHTTPGzipDirective(block.Directive) {
			http.Gzip = processArgsArray(block.Args)
			continue
		}

		// 处理limit_conn_zone指令
		if isHTTPLimitConnZoneDirective(block.Directive) {
			http.LimitConnZone = processArgsArray(block.Args)
			continue
		}

		// 处理 limit_req_zone 指令
		if isHTTPLimitReqZoneDirective(block.Directive) {
			http.LimitReqZone = processArgsArray(block.Args)
			continue
		}

		// 处理open_file_cache_errors指令
		if isHTTPOpenFileCacheErrorsDirective(block.Directive) {
			http.OpenFileCacheErrors = processArgsArray(block.Args)
			continue
		}

		// 处理open_file_cache_min_uses指令
		if isHTTPOpenFileCacheMinUsesDirective(block.Directive) {
			http.OpenFileCacheMinUses = processArgsArray(block.Args)
			continue
		}

		// 处理read_ahead指令
		if isHTTPReadAheadDirective(block.Directive) {
			http.ReadAhead = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_verify_depth指令
		if isHTTPProxySslVerifyDepthDirective(block.Directive) {
			http.ProxySslVerifyDepth = processArgsArray(block.Args)
			continue
		}

		// 处理subrequest_output_buffer_size指令
		if isHTTPSubrequestOutputBufferSizeDirective(block.Directive) {
			http.SubrequestOutputBufferSize = processArgsArray(block.Args)
			continue
		}

		// 处理variables_hash_bucket_size指令
		if isHTTPVariablesHashBucketSizeDirective(block.Directive) {
			http.VariablesHashBucketSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_key指令
		if isHTTPProxyCacheKeyDirective(block.Directive) {
			http.ProxyCacheKey = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_purge指令
		if isHTTPProxyCachePurgeDirective(block.Directive) {
			http.ProxyCachePurge = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_http_version指令
		if isHTTPProxyHTTPVersionDirective(block.Directive) {
			http.ProxyHttpVersion = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_pass_request_body指令
		if isHTTPProxyPassRequestBodyDirective(block.Directive) {
			http.ProxyPassRequestBody = processArgsArray(block.Args)
			continue
		}

		// 处理gzip_buffers指令
		if isHTTPGzipBuffersDirective(block.Directive) {
			http.GzipBuffers = processArgsArray(block.Args)
			continue
		}

		// 处理client_body_in_file_only指令
		if isHTTPClientBodyInFileOnlyDirective(block.Directive) {
			http.ClientBodyInFileOnly = processArgsArray(block.Args)
			continue
		}

		// 处理keepalive_requests指令
		if isHTTPKeepaliveRequestsDirective(block.Directive) {
			http.KeepaliveRequests = processArgsArray(block.Args)
			continue
		}

		// 处理msie_refresh指令
		if isHTTPMsieRefreshDirective(block.Directive) {
			http.MsieRefresh = processArgsArray(block.Args)
			continue
		}

		// 处理client_max_body_size指令
		if isHTTPClientMaxBodySizeDirective(block.Directive) {
			http.ClientMaxBodySize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_method指令
		if isHTTPProxyMethodDirective(block.Directive) {
			http.ProxyMethod = processArgsArray(block.Args)
			continue
		}

		// 处理sendfile_max_chunk指令
		if isHTTPSendFileMaxChunkDirective(block.Directive) {
			http.SendFileMaxChunk = processArgsArray(block.Args)
			continue
		}

		// 处理client_body_in_single_buffer指令
		if isHTTPClientBodyInSingleBufferDirective(block.Directive) {
			http.ClientBodyInSingleBuffer = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_use_stale指令
		if isHTTPProxyCacheUseStaleDirective(block.Directive) {
			http.ProxyCacheUseStale = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_read_timeout指令
		if isHTTPProxyReadTimeoutDirective(block.Directive) {
			http.ProxyReadTimeout = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_server_name指令
		if isHTTPProxySslServerNameDirective(block.Directive) {
			http.ProxySslServerName = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_temp_file_write_size指令
		if isHTTPProxyTempFileWriteSizeDirective(block.Directive) {
			http.ProxyTempFileWriteSize = processArgsArray(block.Args)
			continue
		}

		// 处理aio指令
		if isHTTPAioDirective(block.Directive) {
			http.Aio = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cookie_path指令
		if isHTTPProxyCookiePathDirective(block.Directive) {
			http.ProxyCookiePath = processArgsArray(block.Args)
			continue
		}

		// 处理output_buffers指令
		if isHTTPOutputBuffersDirective(block.Directive) {
			http.OutputBuffers = processArgsArray(block.Args)
			continue
		}

		// 处理reset_timedout_connection指令
		if isHTTPResetTimedoutConnectionDirective(block.Directive) {
			http.ResetTimedoutConnection = processArgsArray(block.Args)
			continue
		}

		// 处理access_log指令
		if isHTTPAccessLogDirective(block.Directive) {
			http.AccessLog = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_path指令
		if isHTTPProxyCachePathDirective(block.Directive) {
			http.ProxyCachePath = processArgsArray(block.Args)
			continue
		}

		// 处理connection_pool_size指令
		if isHTTPConnectionPoolSizeDirective(block.Directive) {
			http.ConnectionPoolSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_next_upstream_tries指令
		if isHTTPProxyNextUpstreamTriesDirective(block.Directive) {
			http.ProxyNextUpstreamTries = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_verify指令
		if isHTTPProxySslVerifyDirective(block.Directive) {
			http.ProxySslVerify = processArgsArray(block.Args)
			continue
		}

		// 处理resolver指令
		if isHTTPResolverDirective(block.Directive) {
			http.Resolver = processArgsArray(block.Args)
			continue
		}

		// 处理absolute_redirect指令
		if isHTTPAbsoluteRedirectDirective(block.Directive) {
			http.AbsoluteRedirect = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_buffer_size指令
		if isHTTPProxyBufferSizeDirective(block.Directive) {
			http.ProxyBufferSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_buffering指令
		if isHTTPProxyBufferingDirective(block.Directive) {
			http.ProxyBuffering = processArgsArray(block.Args)
			continue
		}

		// 处理lingering_time指令
		if isHTTPLingeringTimeDirective(block.Directive) {
			http.LingeringTime = processArgsArray(block.Args)
			continue
		}

		// 处理server_names_hash_max_size指令
		if isHTTPServerNamesHashMaxSizeDirective(block.Directive) {
			http.ServerNamesHashMaxSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_max_range_offset指令
		if isHTTPProxyCacheMaxRangeOffsetDirective(block.Directive) {
			http.ProxyCacheMaxRangeOffset = processArgsArray(block.Args)
			continue
		}

		// 处理keepalive_timeout指令
		if isHTTPKeepaliveTimeoutDirective(block.Directive) {
			http.KeepaliveTimeout = processArgsArray(block.Args)
			continue
		}

		// 处理msie_padding指令
		if isHTTPMsiePaddingDirective(block.Directive) {
			http.MsiePadding = processArgsArray(block.Args)
			continue
		}

		// 处理postpone_output指令
		if isHTTPPostponeOutputDirective(block.Directive) {
			http.PostponeOutput = processArgsArray(block.Args)
			continue
		}

		// 处理send_lowat指令
		if isHTTPSendLowatDirective(block.Directive) {
			http.SendLowat = processArgsArray(block.Args)
			continue
		}

		// 处理ignore_invalid_headers指令
		if isHTTPIgnoreInvalidHeadersDirective(block.Directive) {
			http.IgnoreInvalidHeaders = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_background_update指令
		if isHTTPProxyCacheBackgroundUpdateDirective(block.Directive) {
			http.ProxyCacheBackgroundUpdate = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ignore_headers指令
		if isHTTPProxyIgnoreHeadersDirective(block.Directive) {
			http.ProxyIgnoreHeaders = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_certificate指令
		if isHTTPProxySslCertificateDirective(block.Directive) {
			http.ProxySslCertificate = processArgsArray(block.Args)
			continue
		}

		// 处理underscores_in_headers指令
		if isHTTPUnderscoresInHeadersDirective(block.Directive) {
			http.UnderscoresInHeaders = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_connect_timeout指令
		if isHTTPProxyConnectTimeoutDirective(block.Directive) {
			http.ProxyConnectTimeout = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_intercept_errors指令
		if isHTTPProxyInterceptErrorsDirective(block.Directive) {
			http.ProxyInterceptErrors = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_request_buffering指令
		if isHTTPProxyRequestBufferingDirective(block.Directive) {
			http.ProxyRequestBuffering = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_crl指令
		if isHTTPProxySslCrlDirective(block.Directive) {
			http.ProxySslCrl = processArgsArray(block.Args)
			continue
		}

		// 处理default_type指令
		if isHTTPDefaultTypeDirective(block.Directive) {
			http.DefaultType = processArgsArray(block.Args)
			continue
		}

		// 处理large_client_header_buffers指令
		if isHTTPLargeClientHeaderBuffersDirective(block.Directive) {
			http.LargeClientHeaderBuffers = processArgsArray(block.Args)
			continue
		}

		// 处理log_subrequest指令
		if isHTTPLogSubrequestDirective(block.Directive) {
			http.LogSubrequest = processArgsArray(block.Args)
			continue
		}

		// 处理client_header_timeout指令
		if isHTTPClientHeaderTimeoutDirective(block.Directive) {
			http.ClientHeaderTimeout = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_lock_timeout指令
		if isHTTPProxyCacheLockTimeoutDirective(block.Directive) {
			http.ProxyCacheLockTimeout = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ignore_client_abort指令
		if isHTTPProxyIgnoreClientAbortDirective(block.Directive) {
			http.ProxyIgnoreClientAbort = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_set_body指令
		if isHTTPProxySetBodyDirective(block.Directive) {
			http.ProxySetBody = processArgsArray(block.Args)
			continue
		}

		// 处理cp_nodelay指令
		if isHTTPTcpNodelayDirective(block.Directive) {
			http.TcpNodelay = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_min_uses指令
		if isHTTPProxyCacheMinUsesDirective(block.Directive) {
			http.ProxyCacheMinUses = processArgsArray(block.Args)
			continue
		}

		// 处理keepalive_disable指令
		if isHTTPKeepaliveDisableDirective(block.Directive) {
			http.KeepaliveDisable = processArgsArray(block.Args)
			continue
		}

		// 处理lingering_timeout指令
		if isHTTPLingeringTimeoutDirective(block.Directive) {
			http.LingeringTimeout = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_methods指令
		if isHTTPProxyCacheMethodsDirective(block.Directive) {
			http.ProxyCacheMethods = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_limit_rate指令
		if isHTTPProxyLimitRateDirective(block.Directive) {
			http.ProxyLimitRate = processArgsArray(block.Args)
			continue
		}

		// 处理directio指令
		if isHTTPDirectIODirective(block.Directive) {
			http.DirectIO = processArgsArray(block.Args)
			continue
		}

		// 处理error_page指令
		if isHTTPErrorPageDirective(block.Directive) {
			http.ErrorPage = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_next_upstream指令
		if isHTTPProxyNextUpstreamDirective(block.Directive) {
			http.ProxyNextUpstream = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_no_cache指令
		if isHTTPProxyNoCacheDirective(block.Directive) {
			http.ProxyNoCache = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_redirect指令
		if isHTTPProxyRedirectDirective(block.Directive) {
			http.ProxyRedirect = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_certificate_key指令
		if isHTTPProxySslCertificateKeyDirective(block.Directive) {
			http.ProxySslCertificateKey = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_trusted_certificate指令
		if isHTTPProxySslTrustedCertificateDirective(block.Directive) {
			http.ProxySslTrustedCertificate = processArgsArray(block.Args)
			continue
		}

		// 处理recursive_error_pages指令
		if isHTTPRecursiveErrorPagesDirective(block.Directive) {
			http.RecursiveErrorPages = processArgsArray(block.Args)
			continue
		}

		// 处理limit_rate_after指令
		if isHTTPLimitRateAfterDirective(block.Directive) {
			http.LimitRateAfter = processArgsArray(block.Args)
			continue
		}

		// 处理server_names_hash_bucket_size指令
		if isHTTPServerNamesHashBucketSizeDirective(block.Directive) {
			http.ServerNamesHashBucketSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_next_upstream_timeout指令
		if isHTTPProxyNextUpstreamTimeoutDirective(block.Directive) {
			http.ProxyNextUpstreamTimeout = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_pass_header指令
		if isHTTPProxyPassHeaderDirective(block.Directive) {
			http.ProxyPassHeader = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_ssl_session_reuse指令
		if isHTTPProxySslSessionReuseDirective(block.Directive) {
			http.ProxySslSessionReuse = processArgsArray(block.Args)
			continue
		}

		// 处理gzip_min_length指令
		if isHTTPGzipMinLengthDirective(block.Directive) {
			http.GzipMinLength = processArgsArray(block.Args)
			continue
		}

		// 处理gzip_compl_level指令
		if isHTTPGzipCompLevelDirective(block.Directive) {
			http.GzipCompLevel = processArgsArray(block.Args)
			continue
		}

		// 处理gzip_types指令
		if isHTTPGzipTypesDirective(block.Directive) {
			http.GzipTypes = processArgsArray(block.Args)
			continue
		}

		// 处理gzip_vary指令
		if isHTTPGzipVaryDirective(block.Directive) {
			http.GzipVary = processArgsArray(block.Args)
			continue
		}

		// 处理gzip_disable指令
		if isHTTPGzipDisableDirective(block.Directive) {
			http.GzipDisable = processArgsArray(block.Args)
			continue
		}

		// 处理gzip_http_version指令
		if isHTTPGzipHttpVersionDirective(block.Directive) {
			http.GzipHttpVersion = processArgsArray(block.Args)
			continue
		}

		// 处理gzip_proxied指令
		if isHTTPGzipProxiedDirective(block.Directive) {
			http.GzipProxied = processArgsArray(block.Args)
			continue
		}

		// 处理open_file_cache_valid指令
		if isHTTPOpenFileCacheValidDirective(block.Directive) {
			http.OpenFileCacheValid = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_cache_lock_age指令
		if isHTTPProxyCacheLockAgeDirective(block.Directive) {
			http.ProxyCacheLockAge = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_headers_hash_bucket_size指令
		if isHTTPProxyHeadersHashBucketSizeDirective(block.Directive) {
			http.ProxyHeadersHashBucketSize = processArgsArray(block.Args)
			continue
		}

		// 处理proxy_store指令
		if isHTTPProxyStoreDirective(block.Directive) {
			http.ProxyStore = processArgsArray(block.Args)
			continue
		}

		// 处理deny指令
		if isHTTPDenyDirective(block.Directive) {
			http.Deny = append(http.Deny, processArgsArray(block.Args))
		}

		// 处理proxy_cache_bypass指令
		if isHTTPProxyCacheBypassDirective(block.Directive) {
			http.ProxyCacheBypass = append(http.ProxyCacheBypass, processArgsArray(block.Args))
		}

		// 处理proxy_cache_valid指令
		if isHTTPProxyCacheValidDirective(block.Directive) {
			http.ProxyCacheValid = append(http.ProxyCacheValid, processArgsArray(block.Args))
		}

		// 处理proxy_set_header指令
		if isHTTPProxySetHeaderDirective(block.Directive) {
			http.ProxySetHeader = append(http.ProxySetHeader, processArgsArray(block.Args))
		}

		// 处理allow指令
		if isHTTPAllowDirective(block.Directive) {
			http.Allow = append(http.Allow, processArgsArray(block.Args))
		}

		// 处理include指令
		if isHTTPIncludeDirective(block.Directive) {
			http.Includes = append(http.Includes, processArgsArray(block.Args))
		}

		// 处理upstream指令
		if isHTTPUpstreamDirective(block.Directive) {
			upstream, err := ProcessUpstream(&block)
			if err != nil {
				return http, err
			}
			http.Upstream = append(http.Upstream, *upstream)
		}

		// 处理log_format指令
		if isHTTPLogFormatDirective(block.Directive) {
			http.LogFormat = append(http.LogFormat, processArgsArray(block.Args))
		}

		// 处理server指令
		if isHTTPServerDirective(block.Directive) {
			server, err := ProcessServer(&block)
			if err != nil {
				return http, err
			}
			http.Server = append(http.Server, *server)
		}

		// add_header
		if isHTTPAddHeaderDirective(block.Directive) {
			// 判断add_header指令里是否需要加引号
			for i := range block.Args {
				block.Args[i] = nonescapeQuotation(block.Args[i])
			}
			args := processQuote(block.Args)
			http.AddHeader = append(http.AddHeader, processArgsArray(args))
		}

		// 处理geo指令
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
