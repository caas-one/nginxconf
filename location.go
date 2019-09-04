package nginx

import (
	"errors"
)

// 判断是否是Location指令
func isLocationDirective(directive string) bool {
	if isEqualString(directive, LocationDirective) {
		return true
	}
	return false
}

// 判断是否是aio指令
func isLocationAioDirective(directive string) bool {
	if isEqualString(directive, LocationAioDirective) {
		return true
	}
	return false
}

// 判断是否是aio_write指令
func isLocationAioWriteDirective(directive string) bool {
	if isEqualString(directive, LocationAioWriteDirective) {
		return true
	}
	return false
}

// 判断是否是limit_conn指令
func isLocationLimitConnDirective(directive string) bool {
	if isEqualString(directive, LocationLimitConnDirective) {
		return true
	}
	return false
}

// 判断是否是limit_req指令
func isLocationLimitReqDirective(directive string) bool {
	if isEqualString(directive, LocationLimitReqDirective) {
		return true
	}
	return false
}

// 判断是否是sendfile指令
func isLocationSendFileDirective(directive string) bool {
	if isEqualString(directive, LocationSendFileDirective) {
		return true
	}
	return false
}

// 判断是否是root指令
func isLocationRootDirective(directive string) bool {
	if isEqualString(directive, LocationRootDirective) {
		return true
	}
	return false
}

// 判断是否rewrite指令
func isLocationAliasDirective(directive string) bool {
	if isEqualString(directive, LocationAliasDirective) {
		return true
	}
	return false
}

// 判断是否是chunked_transfer_encoding指令
func isLocationChunkedTransferEncodingDirective(directive string) bool {
	if isEqualString(directive, LocationChunkedTransferEncodingDirective) {
		return true
	}
	return false
}

// 判断是否是chunked_transfer_encoding指令
func isLocationClientBodyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, LocationClientBodyBufferSizeDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_in_file_only指令
func isLocationClientBodyInFileOnlyDirective(directive string) bool {
	if isEqualString(directive, LocationClientBodyInFileOnlyDirective) {
		return true
	}
	return false
}

// 判断是否是disable_symlinks指令
func isLocationDisableSymlinksDirective(directive string) bool {
	if isEqualString(directive, LocationDisableSymlinksDirective) {
		return true
	}
	return false
}

// 判断是否是server_name_in_redirect指令
func isLocationServerNameInRedirectDirective(directive string) bool {
	if isEqualString(directive, LocationServerNameInRedirectDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_temp_path指令
func isLocationClientBodyTempPathDirective(directive string) bool {
	if isEqualString(directive, LocationClientBodyTempPathDirective) {
		return true
	}
	return false
}

// 判断是否是lingering_close指令
func isLocationLingeringCloseDirective(directive string) bool {
	if isEqualString(directive, LocationLingeringCloseDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_send_lowat指令
func isLocationProxySendLowatDirective(directive string) bool {
	if isEqualString(directive, LocationProxySendLowatDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_timeout指令
func isLocationClientBodyTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationClientBodyTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_lock_age指令
func isLocationProxyCacheLockAgeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheLockAgeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_set_header指令
func isLocationProxySetHeaderDirective(directive string) bool {
	if isEqualString(directive, LocationProxySetHeaderDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_buffers指令
func isLocationProxyBuffersDirective(directive string) bool {
	if isEqualString(directive, LocationProxyBuffersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_lock_timeout指令
func isLocationProxyCacheLockTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheLockTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_use_stale指令
func isLocationProxyCacheUseStaleDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheUseStaleDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cookie_path指令
func isLocationProxyCookiePathDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCookiePathDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_store_access指令
func isLocationProxyStoreAccessDirective(directive string) bool {
	if isEqualString(directive, LocationProxyStoreAccessDirective) {
		return true
	}
	return false
}

// 判断是否是keepalive_timeout指令
func isLocationKeepaliveTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationKeepaliveTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_method指令
func isLocationProxyMethodDirective(directive string) bool {
	if isEqualString(directive, LocationProxyMethodDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_lock指令
func isLocationProxyCacheLockDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheLockDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_next_upstream指令
func isLocationProxyNextUpstreamDirective(directive string) bool {
	if isEqualString(directive, LocationProxyNextUpstreamDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_send_timeout指令
func isLocationProxySendTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationProxySendTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是send_timeout指令
func isLocationSendTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationSendTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_min_uses指令
func isLocationProxyCacheMinUsesDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheMinUsesDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_headers_hash_max_size指令
func isLocationProxyHeadersHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyHeadersHashMaxSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_protocols指令
func isLocationProxySslProtocolsDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslProtocolsDirective) {
		return true
	}
	return false
}

// 判断是否是allow指令
func isLocationAllowDirective(directive string) bool {
	if isEqualString(directive, LocationAllowDirective) {
		return true
	}
	return false
}

// 判断是否是types_hash_max_size指令
func isLocationTypesHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, LocationTypesHashMaxSizeDirective) {
		return true
	}
	return false
}

// 判断是否是etag指令
func isLocationEtagDirective(directive string) bool {
	if isEqualString(directive, LocationEtagDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ignore_headers指令
func isLocationProxyIgnoreHeadersDirective(directive string) bool {
	if isEqualString(directive, LocationProxyIgnoreHeadersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_limit_rate指令
func isLocationProxyLimitRateDirective(directive string) bool {
	if isEqualString(directive, LocationProxyLimitRateDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass_request_headers指令
func isLocationProxyPassRequestHeadersDirective(directive string) bool {
	if isEqualString(directive, LocationProxyPassRequestHeadersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_session_reuse指令
func isLocationProxySslSessionReuseDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslSessionReuseDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_store指令
func isLocationProxyStoreDirective(directive string) bool {
	if isEqualString(directive, LocationProxyStoreDirective) {
		return true
	}
	return false
}

// 判断是否是limit_rate_after指令
func isLocationLimitRateAfterDirective(directive string) bool {
	if isEqualString(directive, LocationLimitRateAfterDirective) {
		return true
	}
	return false
}

// 判断是否是satisfy指令
func isLocationSatisfyDirective(directive string) bool {
	if isEqualString(directive, LocationSatisfyDirective) {
		return true
	}
	return false
}

// 判断是否是subrequest_output_buffer_size指令
func isLocationSubrequestOutputBufferSizeDirective(directive string) bool {
	if isEqualString(directive, LocationSubrequestOutputBufferSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_max_temp_file_size指令
func isLocationProxyMaxTempFileSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyMaxTempFileSizeDirective) {
		return true
	}
	return false
}

// 判断是否是resolver指令
func isLocationResolverDirective(directive string) bool {
	if isEqualString(directive, LocationResolverDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_no_cache指令
func isLocationProxyNoCacheDirective(directive string) bool {
	if isEqualString(directive, LocationProxyNoCacheDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_verify指令
func isLocationProxySslVerifyDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslVerifyDirective) {
		return true
	}
	return false
}

// 判断是否是log_not_found指令
func isLocationLogNotFoundDirective(directive string) bool {
	if isEqualString(directive, LocationLogNotFoundDirective) {
		return true
	}
	return false
}

// 判断是否是open_file_cache_errors指令
func isLocationOpenFileCacheErrorsDirective(directive string) bool {
	if isEqualString(directive, LocationOpenFileCacheErrorsDirective) {
		return true
	}
	return false
}

// 判断是否是server_tokens指令
func isLocationServerTokensDirective(directive string) bool {
	if isEqualString(directive, LocationServerTokensDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_busy_buffers_size指令
func isLocationProxyBusyBuffersSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyBusyBuffersSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_convert_head指令
func isLocationProxyCacheConvertHeadDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheConvertHeadDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_next_upstream_tries指令
func isLocationProxyNextUpstreamTriesDirective(directive string) bool {
	if isEqualString(directive, LocationProxyNextUpstreamTriesDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_read_timeout指令
func isLocationProxyReadTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationProxyReadTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是output_buffers指令
func isLocationOutputBuffersDirective(directive string) bool {
	if isEqualString(directive, LocationOutputBuffersDirective) {
		return true
	}
	return false
}

// 判断是否是internal指令
func isLocationInternalDirective(directive string) bool {
	if isEqualString(directive, LocationInternalDirective) {
		return true
	}
	return false
}

// 判断是否是recursive_error_pages指令
func isLocationRecursiveErrorPagesDirective(directive string) bool {
	if isEqualString(directive, LocationRecursiveErrorPagesDirective) {
		return true
	}
	return false
}

// 判断是否是access_log
func isLocationAccessLogDirective(directive string) bool {
	if isEqualString(directive, LocationAccessLogDirective) {
		return true
	}
	return false
}

// 判断是否是log_format指令
func isLocationLogFormatDirective(directive string) bool {
	if isEqualString(directive, LocationLogFormatDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_buffering指令
func isLocationProxyBufferingDirective(directive string) bool {
	if isEqualString(directive, LocationProxyBufferingDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_ciphers指令
func isLocationProxySslCiphersDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslCiphersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_temp_file_write_size指令
func isLocationProxyTempFileWriteSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyTempFileWriteSizeDirective) {
		return true
	}
	return false
}

// 判断是否是error_page指令
func isLocationErrorPageDirective(directive string) bool {
	if isEqualString(directive, LocationErrorPageDirective) {
		return true
	}
	return false
}

// 判断是否是port_in_redirect指令
func isLocationPortInRedirectDirective(directive string) bool {
	if isEqualString(directive, LocationPortInRedirectDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ignore_client_abort指令
func isLocationProxyIgnoreClientAbortDirective(directive string) bool {
	if isEqualString(directive, LocationProxyIgnoreClientAbortDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_name指令
func isLocationProxySslNameDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslNameDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_purge指令
func isLocationProxyCachePurgeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCachePurgeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_certificate_key指令
func isLocationProxySslCertificateKeyDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslCertificateKeyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_methods指令
func isLocationProxyCacheMethodsDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheMethodsDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_intercept_errors指令
func isLocationProxyInterceptErrorsDirective(directive string) bool {
	if isEqualString(directive, LocationProxyInterceptErrorsDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_next_upstream_timeout指令
func isLocationProxyNextUpstreamTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationProxyNextUpstreamTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是resolver_timeout指令
func isLocationResolverTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationResolverTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是absolute_redirect指令
func isLocationAbsoluteRedirectDirective(directive string) bool {
	if isEqualString(directive, LocationAbsoluteRedirectDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass指令
func isLocationProxyPassDirective(directive string) bool {
	if isEqualString(directive, LocationProxyPassDirective) {
		return true
	}
	return false
}

// 判断是否是keepalive_disable指令
func isLocationKeepaliveDisableDirective(directive string) bool {
	if isEqualString(directive, LocationKeepaliveDisableDirective) {
		return true
	}
	return false
}

// 判断是否是lingering_timeout指令
func isLocationLingeringTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationLingeringTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是max_ranges指令
func isLocationMaxRangesDirective(directive string) bool {
	if isEqualString(directive, LocationMaxRangesDirective) {
		return true
	}
	return false
}

// 判断是否是open_file_cache指令
func isLocationOpenFileCacheDirective(directive string) bool {
	if isEqualString(directive, LocationOpenFileCacheDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_max_range_offset指令
func isLocationProxyCacheMaxRangeOffsetDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheMaxRangeOffsetDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_socket_keepalive指令
func isLocationProxySocketKeepaliveDirective(directive string) bool {
	if isEqualString(directive, LocationProxySocketKeepaliveDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_bind指令
func isLocationProxyBindDirective(directive string) bool {
	if isEqualString(directive, LocationProxyBindDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_headers_hash_bucket_size指令
func isLocationProxyHeadersHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyHeadersHashBucketSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_set_body指令
func isLocationProxySetBodyDirective(directive string) bool {
	if isEqualString(directive, LocationProxySetBodyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_password_file指令
func isLocationProxySslPasswordFileDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslPasswordFileDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_trusted_certificate指令
func isLocationProxySslTrustedCertificateDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslTrustedCertificateDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_verify_depth指令
func isLocationProxySslVerifyDepthDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslVerifyDepthDirective) {
		return true
	}
	return false
}

// 判断是否是limit_rate指令
func isLocationLimitRateDirective(directive string) bool {
	if isEqualString(directive, LocationLimitRateDirective) {
		return true
	}
	return false
}

// 判断是否是sendfile_max_chunk指令
func isLocationSendFileMaxChunkDirective(directive string) bool {
	if isEqualString(directive, LocationSendFileMaxChunkDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache指令
func isLocationProxyCacheDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass_request_body指令
func isLocationProxyPassRequestBodyDirective(directive string) bool {
	if isEqualString(directive, LocationProxyPassRequestBodyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_certificate指令
func isLocationProxySslCertificateDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslCertificateDirective) {
		return true
	}
	return false
}

// 判断是否是deny指令
func isLocationDenyDirective(directive string) bool {
	if isEqualString(directive, LocationDenyDirective) {
		return true
	}
	return false
}

// 判断是否是try_files指令
func isLocationTryFilesDirective(directive string) bool {
	if isEqualString(directive, LocationTryFilesDirective) {
		return true
	}
	return false
}

// 判断是否是open_file_cache_valid指令
func isLocationOpenFileCacheValidDirective(directive string) bool {
	if isEqualString(directive, LocationOpenFileCacheValidDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_key指令
func isLocationProxyCacheKeyDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheKeyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_force_ranges指令
func isLocationProxyForceRangesDirective(directive string) bool {
	if isEqualString(directive, LocationProxyForceRangesDirective) {
		return true
	}
	return false
}

// 判断是否是types_hash_bucket_size指令
func isLocationTypesHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, LocationTypesHashBucketSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_path指令
func isLocationProxyCachePathDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCachePathDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_crl指令
func isLocationProxySslCrlDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslCrlDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_temp_path指令
func isLocationProxyTempPathDirective(directive string) bool {
	if isEqualString(directive, LocationProxyTempPathDirective) {
		return true
	}
	return false
}

// 判断是否是default_type指令
func isLocationDefaultTypeDirective(directive string) bool {
	if isEqualString(directive, LocationDefaultTypeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_connect_timeout指令
func isLocationProxyConnectTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationProxyConnectTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是directio_alignment指令
func isLocationDirectioAlignmentDirective(directive string) bool {
	if isEqualString(directive, LocationDirectioAlignmentDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_bypass指令
func isLocationProxyCacheBypassDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheBypassDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_valid指令
func isLocationProxyCacheValidDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheValidDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_background_update指令
func isLocationProxyCacheBackgroundUpdateDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheBackgroundUpdateDirective) {
		return true
	}
	return false
}

// 判断是否是keepalive_requests指令
func isLocationKeepaliveRequestsDirective(directive string) bool {
	if isEqualString(directive, LocationKeepaliveRequestsDirective) {
		return true
	}
	return false
}

// 判断是否是postpone_output指令
func isLocationPostponeOutputDirective(directive string) bool {
	if isEqualString(directive, LocationPostponeOutputDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass_header指令
func isLocationProxyPassHeaderDirective(directive string) bool {
	if isEqualString(directive, LocationProxyPassHeaderDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_request_buffering指令
func isLocationProxyRequestBufferingDirective(directive string) bool {
	if isEqualString(directive, LocationProxyRequestBufferingDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_in_single_buffer指令
func isLocationClientBodyInSingleBufferDirective(directive string) bool {
	if isEqualString(directive, LocationClientBodyInSingleBufferDirective) {
		return true
	}
	return false
}

// 判断是否是directio指令
func isLocationDirectIODirective(directive string) bool {
	if isEqualString(directive, LocationDirectIODirective) {
		return true
	}
	return false
}

// 判断是否是lingering_time指令
func isLocationLingeringTimeDirective(directive string) bool {
	if isEqualString(directive, LocationLingeringTimeDirective) {
		return true
	}
	return false
}

// 判断是否是msie_padding指令
func isLocationMsiePaddingDirective(directive string) bool {
	if isEqualString(directive, LocationMsiePaddingDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cookie_domain指令
func isLocationProxyCookieDomainDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCookieDomainDirective) {
		return true
	}
	return false
}

// 判断是否是send_lowat指令
func isLocationSendLowatDirective(directive string) bool {
	if isEqualString(directive, LocationSendLowatDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_redirect redirect指令
func isLocationProxyRedirectDirective(directive string) bool {
	if isEqualString(directive, LocationProxyRedirectDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_server_name指令
func isLocationProxySslServerNameDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslServerNameDirective) {
		return true
	}
	return false
}

// 判断是否是read_ahead指令
func isLocationReadAheadDirective(directive string) bool {
	if isEqualString(directive, LocationReadAheadDirective) {
		return true
	}
	return false
}

// 判断是否是tcp_nodelay指令
func isLocationTcpNodelayDirective(directive string) bool {
	if isEqualString(directive, LocationTCPNodelayDirective) {
		return true
	}
	return false
}

// 判断是否是tcp_nopush指令
func isLocationTcpNopushDirective(directive string) bool {
	if isEqualString(directive, LocationTCPNopushDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_buffer_size指令
func isLocationProxyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyBufferSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_http_version指令
func isLocationProxyHttpVersionDirective(directive string) bool {
	if isEqualString(directive, LocationProxyHTTPVersionDirective) {
		return true
	}
	return false
}

// 判断是否是client_max_body_size指令
func isLocationClientMaxBodySizeDirective(directive string) bool {
	if isEqualString(directive, LocationClientMaxBodySizeDirective) {
		return true
	}
	return false
}

// 判断是否是expires指令
func isLocationExpiresDirective(directive string) bool {
	if isEqualString(directive, LocationExpiresDirective) {
		return true
	}
	return false
}

// 判断是否是stub_status指令
func isLocationStubStatusDirective(directive string) bool {
	if isEqualString(directive, LocationStubStatusDirective) {
		return true
	}
	return false
}

// 判断if指令
func isLocationIfBlocksDirective(directive string) bool {
	return isEqualString(directive, LocationIfDirective)
}

// 判断if指令
func isLocationSetDirective(directive string) bool {
	return isEqualString(directive, LocationSetDirective)
}

// 判断add_header指令
func isLocationAddHeaderDirective(directive string) bool {
	return isEqualString(directive, AddHeaderDirective)
}

// 判断是否是gzip指令
func isLocationGzipDirective(directive string) bool {
	return isEqualString(directive, LocationGzipDirective)
}

// 判断是否是gzip_min_length指令
func isLocationGzipMinLengthDirective(directive string) bool {
	return isEqualString(directive, LocationGzipMinLengthDirective)
}

// 判断是否是gzip_buffers指令
func isLocationGzipBuffersDirective(directive string) bool {
	return isEqualString(directive, LocationGzipBuffersDirective)
}

// 判断是否是gzip_comp_level指令
func isLocationGzipCompLevelDirective(directive string) bool {
	return isEqualString(directive, LocationGzipCompLevelDirective)
}

// 判断是否是gzip_types指令
func isLocationGzipTypesDirective(directive string) bool {
	return isEqualString(directive, LocationGzipTypesDirective)
}

// 判断是否是gzip_vary指令
func isLocationGzipVaryDirective(directive string) bool {
	return isEqualString(directive, LocationGzipVaryDirective)
}

// 判断是否是gzip_disable指令
func isLocationGzipDisableDirective(directive string) bool {
	return isEqualString(directive, LocationGzipDisableDirective)
}

// 判断是否是gzip_http_version指令
func isLocationGzipHttpVersionDirective(directive string) bool {
	return isEqualString(directive, LocationGzipHTTPVersionDirective)
}

// 判断是否是gzip_proxied指令
func isLocationGzipProxiedDirective(directive string) bool {
	return isEqualString(directive, LocationGzipProxiedDirective)
}

// 判断是否是proxy_hide_header指令
func isLocationProxyHideHeaderDirective(directive string) bool {
	return isEqualString(directive, LocationProxyHideHeaderDirective)
}

// 处理Location中的指令
// 先处理下面的这些指令，以后需要随时添加
/*
access_log
allow
client_body_buffer_size
client_max_body_size
deny
expires
if
proxy_buffer_size
proxy_buffers
proxy_busy_buffers_size
proxy_cache
proxy_cache_key
proxy_cache_purge
proxy_cache_valid
proxy_connect_timeout
proxy_ignore_client_abort
proxy_next_upstream
proxy_pass
proxy_read_timeout
proxy_redirect
proxy_send_timeout
proxy_set_header
proxy_temp_file_write_size
rewrite
root
stub_status
*/
func ProcessLocation(innerBlock *InnerBlock) (*Location, error) {
	if !isLocationDirective(innerBlock.Directive) {
		return nil, errors.New("Not location directive")
	}

	location := NewLocation()
	// 判断Location是否需要加引号
	args := processQuote(innerBlock.Args)
	// 处理Location的Path部分: location = /index.html {}
	location.Path = processArgsArray(args)

	// 处理location中的指令
	for _, inmostBlock := range innerBlock.InmostBlocks {
		/* 全量的 location */

		// 处理proxy_buffers指令
		if isLocationProxyBuffersDirective(inmostBlock.Directive) {
			location.ProxyBuffers = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_convert_head指令
		if isLocationProxyCacheConvertHeadDirective(inmostBlock.Directive) {
			location.ProxyCacheConvertHead = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_http_version指令
		if isLocationProxyHttpVersionDirective(inmostBlock.Directive) {
			location.ProxyHttpVersion = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_min_uses指令
		if isLocationProxyCacheMinUsesDirective(inmostBlock.Directive) {
			location.ProxyCacheMinUses = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_pass_header指令
		if isLocationProxyPassHeaderDirective(inmostBlock.Directive) {
			location.ProxyPassHeader = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理disable_symlinks指令
		if isLocationDisableSymlinksDirective(inmostBlock.Directive) {
			location.DisableSymlinks = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理keepalive_timeout指令
		if isLocationKeepaliveTimeoutDirective(inmostBlock.Directive) {
			location.KeepaliveTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理lingering_timeout指令
		if isLocationLingeringTimeoutDirective(inmostBlock.Directive) {
			location.LingeringTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理sendfile指令
		if isLocationSendFileDirective(inmostBlock.Directive) {
			location.SendFile = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理error_page指令
		if isLocationErrorPageDirective(inmostBlock.Directive) {
			location.ErrorPage = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_key指令
		if isLocationProxyCacheKeyDirective(inmostBlock.Directive) {
			location.ProxyCacheKey = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_methods指令
		if isLocationProxyCacheMethodsDirective(inmostBlock.Directive) {
			location.ProxyCacheMethods = processArgsArray(inmostBlock.Args)
			continue
		}
		if isLocationLimitConnDirective(inmostBlock.Directive) {
			location.LimitConn = processArgsArray(inmostBlock.Args)
			continue
		}
		if isLocationLimitReqDirective(inmostBlock.Directive) {
			location.LimitReq = processArgsArray(inmostBlock.Args)
			continue
		}
		// 处理alias指令
		if isLocationAliasDirective(inmostBlock.Directive) {
			location.Alias = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理types_hash_max_size指令
		if isLocationTypesHashMaxSizeDirective(inmostBlock.Directive) {
			location.TypesHashMaxSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理absolute_redirect指令
		if isLocationAbsoluteRedirectDirective(inmostBlock.Directive) {
			location.AbsoluteRedirect = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_send_lowat指令
		if isLocationProxySendLowatDirective(inmostBlock.Directive) {
			location.ProxySendLowat = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_store指令
		if isLocationProxyStoreDirective(inmostBlock.Directive) {
			location.ProxyStore = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理try_files指令
		if isLocationTryFilesDirective(inmostBlock.Directive) {
			location.TryFiles = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_busy_buffers_size指令
		if isLocationProxyBusyBuffersSizeDirective(inmostBlock.Directive) {
			location.ProxyBusyBuffersSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_next_upstream_timeout指令
		if isLocationProxyNextUpstreamTimeoutDirective(inmostBlock.Directive) {
			location.ProxyNextUpstreamTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理client_body_temp_path指令
		if isLocationClientBodyTempPathDirective(inmostBlock.Directive) {
			location.ClientBodyTempPath = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理directio指令
		if isLocationDirectIODirective(inmostBlock.Directive) {
			location.DirectIO = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理satisfy指令
		if isLocationSatisfyDirective(inmostBlock.Directive) {
			location.Satisfy = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_store_access指令
		if isLocationProxyStoreAccessDirective(inmostBlock.Directive) {
			location.ProxyStoreAccess = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理expires指令
		if isLocationExpiresDirective(inmostBlock.Directive) {
			location.Expires = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理subrequest_output_buffer_size指令
		if isLocationSubrequestOutputBufferSizeDirective(inmostBlock.Directive) {
			location.SubrequestOutputBufferSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_lock指令
		if isLocationProxyCacheLockDirective(inmostBlock.Directive) {
			location.ProxyCacheLock = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_trusted_certificate指令
		if isLocationProxySslTrustedCertificateDirective(inmostBlock.Directive) {
			location.ProxySslTrustedCertificate = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理default_type指令
		if isLocationDefaultTypeDirective(inmostBlock.Directive) {
			location.DefaultType = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理directio_alignment指令
		if isLocationDirectioAlignmentDirective(inmostBlock.Directive) {
			location.DirectioAlignment = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理keepalive_disable指令
		if isLocationKeepaliveDisableDirective(inmostBlock.Directive) {
			location.KeepaliveDisable = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理read_ahead指令
		if isLocationReadAheadDirective(inmostBlock.Directive) {
			location.ReadAhead = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_verify指令
		if isLocationProxySslVerifyDirective(inmostBlock.Directive) {
			location.ProxySslVerify = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_set_body指令
		if isLocationProxySetBodyDirective(inmostBlock.Directive) {
			location.ProxySetBody = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_socket_keepalive指令
		if isLocationProxySocketKeepaliveDirective(inmostBlock.Directive) {
			location.ProxySocketKeepalive = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理client_max_body_size指令
		if isLocationClientMaxBodySizeDirective(inmostBlock.Directive) {
			location.ClientMaxBodySize = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理log_not_found指令
		if isLocationLogNotFoundDirective(inmostBlock.Directive) {
			location.LogNotFound = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理access_log指令
		if isLocationAccessLogDirective(inmostBlock.Directive) {
			location.AccessLog = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理resolver指令
		if isLocationResolverDirective(inmostBlock.Directive) {
			location.Resolver = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理server_name_in_redirect指令
		if isLocationServerNameInRedirectDirective(inmostBlock.Directive) {
			location.ServerNameInRedirect = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_read_timeout指令
		if isLocationProxyReadTimeoutDirective(inmostBlock.Directive) {
			location.ProxyReadTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理limit_rate指令
		if isLocationLimitRateDirective(inmostBlock.Directive) {
			location.LimitRate = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_request_buffering指令
		if isLocationProxyRequestBufferingDirective(inmostBlock.Directive) {
			location.ProxyRequestBuffering = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理root指令
		if isLocationRootDirective(inmostBlock.Directive) {
			location.Root = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理client_body_buffer_size指令
		if isLocationClientBodyBufferSizeDirective(inmostBlock.Directive) {
			location.ClientBodyBufferSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理max_ranges指令
		if isLocationMaxRangesDirective(inmostBlock.Directive) {
			location.MaxRanges = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理tcp_nopush指令
		if isLocationTcpNopushDirective(inmostBlock.Directive) {
			location.TcpNopush = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_no_cache指令
		if isLocationProxyNoCacheDirective(inmostBlock.Directive) {
			location.ProxyNoCache = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_password_file指令
		if isLocationProxySslPasswordFileDirective(inmostBlock.Directive) {
			location.ProxySslPasswordFile = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_headers_hash_max_size指令
		if isLocationProxyHeadersHashMaxSizeDirective(inmostBlock.Directive) {
			location.ProxyHeadersHashMaxSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_pass_request_headers指令
		if isLocationProxyPassRequestHeadersDirective(inmostBlock.Directive) {
			location.ProxyPassRequestHeaders = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_certificate_key指令
		if isLocationProxySslCertificateKeyDirective(inmostBlock.Directive) {
			location.ProxySslCertificateKey = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_verify_depth指令
		if isLocationProxySslVerifyDepthDirective(inmostBlock.Directive) {
			location.ProxySslVerifyDepth = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理keepalive_requests指令
		if isLocationKeepaliveRequestsDirective(inmostBlock.Directive) {
			location.KeepaliveRequests = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理open_file_cache指令
		if isLocationOpenFileCacheDirective(inmostBlock.Directive) {
			location.OpenFileCache = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理output_buffers指令
		if isLocationOutputBuffersDirective(inmostBlock.Directive) {
			location.OutputBuffers = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_buffering指令
		if isLocationProxyBufferingDirective(inmostBlock.Directive) {
			location.ProxyBuffering = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理recursive_error_pages指令
		if isLocationRecursiveErrorPagesDirective(inmostBlock.Directive) {
			location.RecursiveErrorPages = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_background_update指令
		if isLocationProxyCacheBackgroundUpdateDirective(inmostBlock.Directive) {
			location.ProxyCacheBackgroundUpdate = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_connect_timeout指令
		if isLocationProxyConnectTimeoutDirective(inmostBlock.Directive) {
			location.ProxyConnectTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理client_body_timeout指令
		if isLocationClientBodyTimeoutDirective(inmostBlock.Directive) {
			location.ClientBodyTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_bind指令
		if isLocationProxyBindDirective(inmostBlock.Directive) {
			location.ProxyBind = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_lock_age指令
		if isLocationProxyCacheLockAgeDirective(inmostBlock.Directive) {
			location.ProxyCacheLockAge = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_server_name指令
		if isLocationProxySslServerNameDirective(inmostBlock.Directive) {
			location.ProxySslServerName = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache指令
		if isLocationProxyCacheDirective(inmostBlock.Directive) {
			location.ProxyCache = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_max_range_offset指令
		if isLocationProxyCacheMaxRangeOffsetDirective(inmostBlock.Directive) {
			location.ProxyCacheMaxRangeOffset = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_pass指令
		if isLocationProxyPassDirective(inmostBlock.Directive) {
			location.ProxyPass = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理chunked_transfer_encoding指令
		if isLocationChunkedTransferEncodingDirective(inmostBlock.Directive) {
			location.ChunkedTransferEncoding = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理tcp_nodelay指令
		if isLocationTcpNodelayDirective(inmostBlock.Directive) {
			location.TcpNodelay = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_lock_timeout指令
		if isLocationProxyCacheLockTimeoutDirective(inmostBlock.Directive) {
			location.ProxyCacheLockTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_intercept_errors指令
		if isLocationProxyInterceptErrorsDirective(inmostBlock.Directive) {
			location.ProxyInterceptErrors = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_redirect指令
		if isLocationProxyRedirectDirective(inmostBlock.Directive) {
			location.ProxyRedirect = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_path指令
		if isLocationProxyCachePathDirective(inmostBlock.Directive) {
			location.ProxyCachePath = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_name指令
		if isLocationProxySslNameDirective(inmostBlock.Directive) {
			location.ProxySslName = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_use_stale指令
		if isLocationProxyCacheUseStaleDirective(inmostBlock.Directive) {
			location.ProxyCacheUseStale = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_max_temp_file_size指令
		if isLocationProxyMaxTempFileSizeDirective(inmostBlock.Directive) {
			location.ProxyMaxTempFileSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_next_upstream指令
		if isLocationProxyNextUpstreamDirective(inmostBlock.Directive) {
			location.ProxyNextUpstream = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理limit_rate_after指令
		if isLocationLimitRateAfterDirective(inmostBlock.Directive) {
			location.LimitRateAfter = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理lingering_close指令
		if isLocationLingeringCloseDirective(inmostBlock.Directive) {
			location.LingeringClose = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理server_tokens指令
		if isLocationServerTokensDirective(inmostBlock.Directive) {
			location.ServerTokens = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_protocols指令
		if isLocationProxySslProtocolsDirective(inmostBlock.Directive) {
			location.ProxySslProtocols = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理aio_write指令
		if isLocationAioWriteDirective(inmostBlock.Directive) {
			location.AioWrite = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理client_body_in_file_only指令
		if isLocationClientBodyInFileOnlyDirective(inmostBlock.Directive) {
			location.ClientBodyInFileOnly = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理open_file_cache_errors指令
		if isLocationOpenFileCacheErrorsDirective(inmostBlock.Directive) {
			location.OpenFileCacheErrors = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_next_upstream_tries指令
		if isLocationProxyNextUpstreamTriesDirective(inmostBlock.Directive) {
			location.ProxyNextUpstreamTries = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理lingering_time指令
		if isLocationLingeringTimeDirective(inmostBlock.Directive) {
			location.LingeringTime = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理open_file_cache_valid指令
		if isLocationOpenFileCacheValidDirective(inmostBlock.Directive) {
			location.OpenFileCacheValid = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理postpone_output指令
		if isLocationPostponeOutputDirective(inmostBlock.Directive) {
			location.PostponeOutput = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_limit_rate指令
		if isLocationProxyLimitRateDirective(inmostBlock.Directive) {
			location.ProxyLimitRate = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_temp_file_write_size指令
		if isLocationProxyTempFileWriteSizeDirective(inmostBlock.Directive) {
			location.ProxyTempFileWriteSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理internal指令
		if isLocationInternalDirective(inmostBlock.Directive) {
			location.Internal = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理sendfile_max_chunk指令
		if isLocationSendFileMaxChunkDirective(inmostBlock.Directive) {
			location.SendFileMaxChunk = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理port_in_redirect指令
		if isLocationPortInRedirectDirective(inmostBlock.Directive) {
			location.PortInRedirect = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理resolver_timeout指令
		if isLocationResolverTimeoutDirective(inmostBlock.Directive) {
			location.ResolverTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_purge指令
		if isLocationProxyCachePurgeDirective(inmostBlock.Directive) {
			location.ProxyCachePurge = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_headers_hash_bucket_size指令
		if isLocationProxyHeadersHashBucketSizeDirective(inmostBlock.Directive) {
			location.ProxyHeadersHashBucketSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_send_timeout指令
		if isLocationProxySendTimeoutDirective(inmostBlock.Directive) {
			location.ProxySendTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理client_body_in_single_buffer指令
		if isLocationClientBodyInSingleBufferDirective(inmostBlock.Directive) {
			location.ClientBodyInSingleBuffer = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理send_lowat指令
		if isLocationSendLowatDirective(inmostBlock.Directive) {
			location.SendLowat = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理send_timeout指令
		if isLocationSendTimeoutDirective(inmostBlock.Directive) {
			location.SendTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ignore_client_abort指令
		if isLocationProxyIgnoreClientAbortDirective(inmostBlock.Directive) {
			location.ProxyIgnoreClientAbort = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_buffer_size指令
		if isLocationProxyBufferSizeDirective(inmostBlock.Directive) {
			location.ProxyBufferSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_certificate指令
		if isLocationProxySslCertificateDirective(inmostBlock.Directive) {
			location.ProxySslCertificate = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理msie_padding指令
		if isLocationMsiePaddingDirective(inmostBlock.Directive) {
			location.MsiePadding = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_force_ranges指令
		if isLocationProxyForceRangesDirective(inmostBlock.Directive) {
			location.ProxyForceRanges = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_pass_request_body指令
		if isLocationProxyPassRequestBodyDirective(inmostBlock.Directive) {
			location.ProxyPassRequestBody = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_crl指令
		if isLocationProxySslCrlDirective(inmostBlock.Directive) {
			location.ProxySslCrl = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_temp_path指令
		if isLocationProxyTempPathDirective(inmostBlock.Directive) {
			location.ProxyTempPath = processArgsArray(inmostBlock.Args)
			continue
		}
		// 处理stub_status指令
		if isLocationStubStatusDirective(inmostBlock.Directive) {
			location.StubStatus = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理etag指令
		if isLocationEtagDirective(inmostBlock.Directive) {
			location.Etag = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cookie_domain指令
		if isLocationProxyCookieDomainDirective(inmostBlock.Directive) {
			location.ProxyCookieDomain = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cookie_path指令
		if isLocationProxyCookiePathDirective(inmostBlock.Directive) {
			location.ProxyCookiePath = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ignore_headers指令
		if isLocationProxyIgnoreHeadersDirective(inmostBlock.Directive) {
			location.ProxyIgnoreHeaders = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理aio指令
		if isLocationAioDirective(inmostBlock.Directive) {
			location.Aio = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理types_hash_bucket_size指令
		if isLocationTypesHashBucketSizeDirective(inmostBlock.Directive) {
			location.TypesHashBucketSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_method指令
		if isLocationProxyMethodDirective(inmostBlock.Directive) {
			location.ProxyMethod = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_ciphers指令
		if isLocationProxySslCiphersDirective(inmostBlock.Directive) {
			location.ProxySslCiphers = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_ssl_session_reuse指令
		if isLocationProxySslSessionReuseDirective(inmostBlock.Directive) {
			location.ProxySslSessionReuse = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理gzip指令
		if isLocationGzipDirective(inmostBlock.Directive) {
			location.Gzip = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理gzip_min_length指令
		if isLocationGzipMinLengthDirective(inmostBlock.Directive) {
			location.GzipMinLength = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理gzip_buffers指令
		if isLocationGzipBuffersDirective(inmostBlock.Directive) {
			location.GzipBuffers = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理gzip_comp_level指令
		if isLocationGzipCompLevelDirective(inmostBlock.Directive) {
			location.GzipCompLevel = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理gzip_types指令
		if isLocationGzipTypesDirective(inmostBlock.Directive) {
			location.GzipTypes = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理gzip_vary指令
		if isLocationGzipVaryDirective(inmostBlock.Directive) {
			location.GzipVary = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理gzip_disable指令
		if isLocationGzipDisableDirective(inmostBlock.Directive) {
			location.GzipDisable = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理gzip_http_version指令
		if isLocationGzipHttpVersionDirective(inmostBlock.Directive) {
			location.GzipHttpVersion = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理gzip_proxied指令
		if isLocationGzipProxiedDirective(inmostBlock.Directive) {
			location.GzipProxied = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_hide_header指令
		if isLocationProxyHideHeaderDirective(inmostBlock.Directive) {
			location.ProxyHideHeader = processArgsArray(inmostBlock.Args)
			continue
		}

		// if isLocationIfBlocksDirective(inmostBlock.Directive) {
		// location.IfBlock = append(location.IfBlock, processArgsArray(inmostBlock.Args))

		// }

		// 处理rewrite指令
		if isLocationRewriteRootDirective(inmostBlock.Directive) {
			args := processQuote(inmostBlock.Args)
			location.Rewrite = append(location.Rewrite, processArgsArray(args))
			continue
		}

		// 处理proxy_cache_bypass指令
		if isLocationProxyCacheBypassDirective(inmostBlock.Directive) {
			location.ProxyCacheBypass = append(location.ProxyCacheBypass, processArgsArray(inmostBlock.Args))
			continue
		}

		// 处理proxy_cache指令
		if isLocationProxyCacheDirective(inmostBlock.Directive) {
			location.ProxyCache = processArgsArray(inmostBlock.Args)
			continue
		}

		// 处理proxy_cache_valid指令
		if isLocationProxyCacheValidDirective(inmostBlock.Directive) {
			location.ProxyCacheValid = append(location.ProxyCacheValid, processArgsArray(inmostBlock.Args))
			continue
		}

		// 处理allow指令
		if isLocationAllowDirective(inmostBlock.Directive) {
			location.Allow = append(location.Allow, processArgsArray(inmostBlock.Args))
			continue
		}

		// 处理proxy_set_header指令
		if isLocationProxySetHeaderDirective(inmostBlock.Directive) {
			location.ProxySetHeader = append(location.ProxySetHeader, processArgsArray(inmostBlock.Args))
			continue
		}

		// 处理deny指令
		// if isLocationDenyDirective(inmostBlock.Directive) {
		// 	location.Deny = append(location.Deny, processArgsArray(inmostBlock.Args))
		// 	continue
		// }

		// 处理deny指令
		if isLocationDenyDirective(inmostBlock.Directive) {
			location.Deny = append(location.Deny, processArgsArray(inmostBlock.Args))
			continue
		}
		// 处理set指令
		if isLocationSetDirective(inmostBlock.Directive) {
			location.Set = append(location.Set, processArgsArray(inmostBlock.Args))
		}

		// 处理log_format指令
		if isLocationLogFormatDirective(inmostBlock.Directive) {
			location.LogFormat = append(location.LogFormat, processArgsArray(inmostBlock.Args))
			continue
		}
		// 处理if指令
		if isLocationIfBlocksDirective(inmostBlock.Directive) {
			ifBlock, err := processLocationIfDirective(inmostBlock)
			if err != nil {
				return location, err
			}
			location.IfBlocks = append(location.IfBlocks, *ifBlock)
			continue
		}

		// add_header
		if isLocationAddHeaderDirective(inmostBlock.Directive) {
			// 判断add_header指令里是否需要加引号
			for i := range inmostBlock.Args {
				inmostBlock.Args[i] = nonescapeQuotation(inmostBlock.Args[i])
			}
			args := processQuote(inmostBlock.Args)
			location.AddHeader = append(location.AddHeader, processArgsArray(args))
		}

		//TODO: limit_req
	}
	return location, nil
}
