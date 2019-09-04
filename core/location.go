package core

import (
	"errors"
)

// determine whether it is a Location directive
func isLocationDirective(directive string) bool {
	if isEqualString(directive, LocationDirective) {
		return true
	}
	return false
}

// determine whether it is a aio directive
func isLocationAioDirective(directive string) bool {
	if isEqualString(directive, LocationAioDirective) {
		return true
	}
	return false
}

// determine whether it is a aio_write directive
func isLocationAioWriteDirective(directive string) bool {
	if isEqualString(directive, LocationAioWriteDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_conn directive
func isLocationLimitConnDirective(directive string) bool {
	if isEqualString(directive, LocationLimitConnDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_req directive
func isLocationLimitReqDirective(directive string) bool {
	if isEqualString(directive, LocationLimitReqDirective) {
		return true
	}
	return false
}

// determine whether it is a sendfile directive
func isLocationSendFileDirective(directive string) bool {
	if isEqualString(directive, LocationSendFileDirective) {
		return true
	}
	return false
}

// determine whether it is a root directive
func isLocationRootDirective(directive string) bool {
	if isEqualString(directive, LocationRootDirective) {
		return true
	}
	return false
}

// rewrite directive
func isLocationAliasDirective(directive string) bool {
	if isEqualString(directive, LocationAliasDirective) {
		return true
	}
	return false
}

// determine whether it is a chunked_transfer_encoding directive
func isLocationChunkedTransferEncodingDirective(directive string) bool {
	if isEqualString(directive, LocationChunkedTransferEncodingDirective) {
		return true
	}
	return false
}

// determine whether it is a chunked_transfer_encoding directive
func isLocationClientBodyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, LocationClientBodyBufferSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_in_file_only directive
func isLocationClientBodyInFileOnlyDirective(directive string) bool {
	if isEqualString(directive, LocationClientBodyInFileOnlyDirective) {
		return true
	}
	return false
}

// determine whether it is a disable_symlinks directive
func isLocationDisableSymlinksDirective(directive string) bool {
	if isEqualString(directive, LocationDisableSymlinksDirective) {
		return true
	}
	return false
}

// determine whether it is a server_name_in_redirect directive
func isLocationServerNameInRedirectDirective(directive string) bool {
	if isEqualString(directive, LocationServerNameInRedirectDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_temp_path directive
func isLocationClientBodyTempPathDirective(directive string) bool {
	if isEqualString(directive, LocationClientBodyTempPathDirective) {
		return true
	}
	return false
}

// determine whether it is a lingering_close directive
func isLocationLingeringCloseDirective(directive string) bool {
	if isEqualString(directive, LocationLingeringCloseDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_send_lowat directive
func isLocationProxySendLowatDirective(directive string) bool {
	if isEqualString(directive, LocationProxySendLowatDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_timeout directive
func isLocationClientBodyTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationClientBodyTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_lock_age directive
func isLocationProxyCacheLockAgeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheLockAgeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_set_header directive
func isLocationProxySetHeaderDirective(directive string) bool {
	if isEqualString(directive, LocationProxySetHeaderDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_buffers directive
func isLocationProxyBuffersDirective(directive string) bool {
	if isEqualString(directive, LocationProxyBuffersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_lock_timeout directive
func isLocationProxyCacheLockTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheLockTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_use_stale directive
func isLocationProxyCacheUseStaleDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheUseStaleDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cookie_path directive
func isLocationProxyCookiePathDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCookiePathDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_store_access directive
func isLocationProxyStoreAccessDirective(directive string) bool {
	if isEqualString(directive, LocationProxyStoreAccessDirective) {
		return true
	}
	return false
}

// determine whether it is a keepalive_timeout directive
func isLocationKeepaliveTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationKeepaliveTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_method directive
func isLocationProxyMethodDirective(directive string) bool {
	if isEqualString(directive, LocationProxyMethodDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_lock directive
func isLocationProxyCacheLockDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheLockDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_next_upstream directive
func isLocationProxyNextUpstreamDirective(directive string) bool {
	if isEqualString(directive, LocationProxyNextUpstreamDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_send_timeout directive
func isLocationProxySendTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationProxySendTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a send_timeout directive
func isLocationSendTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationSendTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_min_uses directive
func isLocationProxyCacheMinUsesDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheMinUsesDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_headers_hash_max_size directive
func isLocationProxyHeadersHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyHeadersHashMaxSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_protocols directive
func isLocationProxySslProtocolsDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslProtocolsDirective) {
		return true
	}
	return false
}

// determine whether it is a allow directive
func isLocationAllowDirective(directive string) bool {
	if isEqualString(directive, LocationAllowDirective) {
		return true
	}
	return false
}

// determine whether it is a types_hash_max_size directive
func isLocationTypesHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, LocationTypesHashMaxSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a etag directive
func isLocationEtagDirective(directive string) bool {
	if isEqualString(directive, LocationEtagDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ignore_headers directive
func isLocationProxyIgnoreHeadersDirective(directive string) bool {
	if isEqualString(directive, LocationProxyIgnoreHeadersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_limit_rate directive
func isLocationProxyLimitRateDirective(directive string) bool {
	if isEqualString(directive, LocationProxyLimitRateDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass_request_headers directive
func isLocationProxyPassRequestHeadersDirective(directive string) bool {
	if isEqualString(directive, LocationProxyPassRequestHeadersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_session_reuse directive
func isLocationProxySslSessionReuseDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslSessionReuseDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_store directive
func isLocationProxyStoreDirective(directive string) bool {
	if isEqualString(directive, LocationProxyStoreDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_rate_after directive
func isLocationLimitRateAfterDirective(directive string) bool {
	if isEqualString(directive, LocationLimitRateAfterDirective) {
		return true
	}
	return false
}

// determine whether it is a satisfy directive
func isLocationSatisfyDirective(directive string) bool {
	if isEqualString(directive, LocationSatisfyDirective) {
		return true
	}
	return false
}

// determine whether it is a subrequest_output_buffer_size directive
func isLocationSubrequestOutputBufferSizeDirective(directive string) bool {
	if isEqualString(directive, LocationSubrequestOutputBufferSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_max_temp_file_size directive
func isLocationProxyMaxTempFileSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyMaxTempFileSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a resolver directive
func isLocationResolverDirective(directive string) bool {
	if isEqualString(directive, LocationResolverDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_no_cache directive
func isLocationProxyNoCacheDirective(directive string) bool {
	if isEqualString(directive, LocationProxyNoCacheDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_verify directive
func isLocationProxySslVerifyDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslVerifyDirective) {
		return true
	}
	return false
}

// determine whether it is a log_not_found directive
func isLocationLogNotFoundDirective(directive string) bool {
	if isEqualString(directive, LocationLogNotFoundDirective) {
		return true
	}
	return false
}

// determine whether it is a open_file_cache_errors directive
func isLocationOpenFileCacheErrorsDirective(directive string) bool {
	if isEqualString(directive, LocationOpenFileCacheErrorsDirective) {
		return true
	}
	return false
}

// determine whether it is a server_tokens directive
func isLocationServerTokensDirective(directive string) bool {
	if isEqualString(directive, LocationServerTokensDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_busy_buffers_size directive
func isLocationProxyBusyBuffersSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyBusyBuffersSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_convert_head directive
func isLocationProxyCacheConvertHeadDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheConvertHeadDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_next_upstream_tries directive
func isLocationProxyNextUpstreamTriesDirective(directive string) bool {
	if isEqualString(directive, LocationProxyNextUpstreamTriesDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_read_timeout directive
func isLocationProxyReadTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationProxyReadTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a output_buffers directive
func isLocationOutputBuffersDirective(directive string) bool {
	if isEqualString(directive, LocationOutputBuffersDirective) {
		return true
	}
	return false
}

// determine whether it is a internal directive
func isLocationInternalDirective(directive string) bool {
	if isEqualString(directive, LocationInternalDirective) {
		return true
	}
	return false
}

// determine whether it is a recursive_error_pages directive
func isLocationRecursiveErrorPagesDirective(directive string) bool {
	if isEqualString(directive, LocationRecursiveErrorPagesDirective) {
		return true
	}
	return false
}

// determine whether it is a access_log
func isLocationAccessLogDirective(directive string) bool {
	if isEqualString(directive, LocationAccessLogDirective) {
		return true
	}
	return false
}

// determine whether it is a log_format directive
func isLocationLogFormatDirective(directive string) bool {
	if isEqualString(directive, LocationLogFormatDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_buffering directive
func isLocationProxyBufferingDirective(directive string) bool {
	if isEqualString(directive, LocationProxyBufferingDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_ciphers directive
func isLocationProxySslCiphersDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslCiphersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_temp_file_write_size directive
func isLocationProxyTempFileWriteSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyTempFileWriteSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a error_page directive
func isLocationErrorPageDirective(directive string) bool {
	if isEqualString(directive, LocationErrorPageDirective) {
		return true
	}
	return false
}

// determine whether it is a port_in_redirect directive
func isLocationPortInRedirectDirective(directive string) bool {
	if isEqualString(directive, LocationPortInRedirectDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ignore_client_abort directive
func isLocationProxyIgnoreClientAbortDirective(directive string) bool {
	if isEqualString(directive, LocationProxyIgnoreClientAbortDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_name directive
func isLocationProxySslNameDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslNameDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_purge directive
func isLocationProxyCachePurgeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCachePurgeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_certificate_key directive
func isLocationProxySslCertificateKeyDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslCertificateKeyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_methods directive
func isLocationProxyCacheMethodsDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheMethodsDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_intercept_errors directive
func isLocationProxyInterceptErrorsDirective(directive string) bool {
	if isEqualString(directive, LocationProxyInterceptErrorsDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_next_upstream_timeout directive
func isLocationProxyNextUpstreamTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationProxyNextUpstreamTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a resolver_timeout directive
func isLocationResolverTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationResolverTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a absolute_redirect directive
func isLocationAbsoluteRedirectDirective(directive string) bool {
	if isEqualString(directive, LocationAbsoluteRedirectDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass directive
func isLocationProxyPassDirective(directive string) bool {
	if isEqualString(directive, LocationProxyPassDirective) {
		return true
	}
	return false
}

// determine whether it is a keepalive_disable directive
func isLocationKeepaliveDisableDirective(directive string) bool {
	if isEqualString(directive, LocationKeepaliveDisableDirective) {
		return true
	}
	return false
}

// determine whether it is a lingering_timeout directive
func isLocationLingeringTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationLingeringTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a max_ranges directive
func isLocationMaxRangesDirective(directive string) bool {
	if isEqualString(directive, LocationMaxRangesDirective) {
		return true
	}
	return false
}

// determine whether it is a open_file_cache directive
func isLocationOpenFileCacheDirective(directive string) bool {
	if isEqualString(directive, LocationOpenFileCacheDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_max_range_offset directive
func isLocationProxyCacheMaxRangeOffsetDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheMaxRangeOffsetDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_socket_keepalive directive
func isLocationProxySocketKeepaliveDirective(directive string) bool {
	if isEqualString(directive, LocationProxySocketKeepaliveDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_bind directive
func isLocationProxyBindDirective(directive string) bool {
	if isEqualString(directive, LocationProxyBindDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_headers_hash_bucket_size directive
func isLocationProxyHeadersHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyHeadersHashBucketSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_set_body directive
func isLocationProxySetBodyDirective(directive string) bool {
	if isEqualString(directive, LocationProxySetBodyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_password_file directive
func isLocationProxySslPasswordFileDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslPasswordFileDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_trusted_certificate directive
func isLocationProxySslTrustedCertificateDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslTrustedCertificateDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_verify_depth directive
func isLocationProxySslVerifyDepthDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslVerifyDepthDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_rate directive
func isLocationLimitRateDirective(directive string) bool {
	if isEqualString(directive, LocationLimitRateDirective) {
		return true
	}
	return false
}

// determine whether it is a sendfile_max_chunk directive
func isLocationSendFileMaxChunkDirective(directive string) bool {
	if isEqualString(directive, LocationSendFileMaxChunkDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache directive
func isLocationProxyCacheDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass_request_body directive
func isLocationProxyPassRequestBodyDirective(directive string) bool {
	if isEqualString(directive, LocationProxyPassRequestBodyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_certificate directive
func isLocationProxySslCertificateDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslCertificateDirective) {
		return true
	}
	return false
}

// determine whether it is a deny directive
func isLocationDenyDirective(directive string) bool {
	if isEqualString(directive, LocationDenyDirective) {
		return true
	}
	return false
}

// determine whether it is a try_files directive
func isLocationTryFilesDirective(directive string) bool {
	if isEqualString(directive, LocationTryFilesDirective) {
		return true
	}
	return false
}

// determine whether it is a open_file_cache_valid directive
func isLocationOpenFileCacheValidDirective(directive string) bool {
	if isEqualString(directive, LocationOpenFileCacheValidDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_key directive
func isLocationProxyCacheKeyDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheKeyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_force_ranges directive
func isLocationProxyForceRangesDirective(directive string) bool {
	if isEqualString(directive, LocationProxyForceRangesDirective) {
		return true
	}
	return false
}

// determine whether it is a types_hash_bucket_size directive
func isLocationTypesHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, LocationTypesHashBucketSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_path directive
func isLocationProxyCachePathDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCachePathDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_crl directive
func isLocationProxySslCrlDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslCrlDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_temp_path directive
func isLocationProxyTempPathDirective(directive string) bool {
	if isEqualString(directive, LocationProxyTempPathDirective) {
		return true
	}
	return false
}

// determine whether it is a default_type directive
func isLocationDefaultTypeDirective(directive string) bool {
	if isEqualString(directive, LocationDefaultTypeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_connect_timeout directive
func isLocationProxyConnectTimeoutDirective(directive string) bool {
	if isEqualString(directive, LocationProxyConnectTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a directio_alignment directive
func isLocationDirectioAlignmentDirective(directive string) bool {
	if isEqualString(directive, LocationDirectioAlignmentDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_bypass directive
func isLocationProxyCacheBypassDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheBypassDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_valid directive
func isLocationProxyCacheValidDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheValidDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_background_update directive
func isLocationProxyCacheBackgroundUpdateDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCacheBackgroundUpdateDirective) {
		return true
	}
	return false
}

// determine whether it is a keepalive_requests directive
func isLocationKeepaliveRequestsDirective(directive string) bool {
	if isEqualString(directive, LocationKeepaliveRequestsDirective) {
		return true
	}
	return false
}

// determine whether it is a postpone_output directive
func isLocationPostponeOutputDirective(directive string) bool {
	if isEqualString(directive, LocationPostponeOutputDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass_header directive
func isLocationProxyPassHeaderDirective(directive string) bool {
	if isEqualString(directive, LocationProxyPassHeaderDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_request_buffering directive
func isLocationProxyRequestBufferingDirective(directive string) bool {
	if isEqualString(directive, LocationProxyRequestBufferingDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_in_single_buffer directive
func isLocationClientBodyInSingleBufferDirective(directive string) bool {
	if isEqualString(directive, LocationClientBodyInSingleBufferDirective) {
		return true
	}
	return false
}

// determine whether it is a directio directive
func isLocationDirectIODirective(directive string) bool {
	if isEqualString(directive, LocationDirectIODirective) {
		return true
	}
	return false
}

// determine whether it is a lingering_time directive
func isLocationLingeringTimeDirective(directive string) bool {
	if isEqualString(directive, LocationLingeringTimeDirective) {
		return true
	}
	return false
}

// determine whether it is a msie_padding directive
func isLocationMsiePaddingDirective(directive string) bool {
	if isEqualString(directive, LocationMsiePaddingDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cookie_domain directive
func isLocationProxyCookieDomainDirective(directive string) bool {
	if isEqualString(directive, LocationProxyCookieDomainDirective) {
		return true
	}
	return false
}

// determine whether it is a send_lowat directive
func isLocationSendLowatDirective(directive string) bool {
	if isEqualString(directive, LocationSendLowatDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_redirect redirect directive
func isLocationProxyRedirectDirective(directive string) bool {
	if isEqualString(directive, LocationProxyRedirectDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_server_name directive
func isLocationProxySslServerNameDirective(directive string) bool {
	if isEqualString(directive, LocationProxySslServerNameDirective) {
		return true
	}
	return false
}

// determine whether it is a read_ahead directive
func isLocationReadAheadDirective(directive string) bool {
	if isEqualString(directive, LocationReadAheadDirective) {
		return true
	}
	return false
}

// determine whether it is a tcp_nodelay directive
func isLocationTCPNodelayDirective(directive string) bool {
	if isEqualString(directive, LocationTCPNodelayDirective) {
		return true
	}
	return false
}

// determine whether it is a tcp_nopush directive
func isLocationTCPNopushDirective(directive string) bool {
	if isEqualString(directive, LocationTCPNopushDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_buffer_size directive
func isLocationProxyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, LocationProxyBufferSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_http_version directive
func isLocationProxyHTTPVersionDirective(directive string) bool {
	if isEqualString(directive, LocationProxyHTTPVersionDirective) {
		return true
	}
	return false
}

// determine whether it is a client_max_body_size directive
func isLocationClientMaxBodySizeDirective(directive string) bool {
	if isEqualString(directive, LocationClientMaxBodySizeDirective) {
		return true
	}
	return false
}

// determine whether it is a expires directive
func isLocationExpiresDirective(directive string) bool {
	if isEqualString(directive, LocationExpiresDirective) {
		return true
	}
	return false
}

// determine whether it is a stub_status directive
func isLocationStubStatusDirective(directive string) bool {
	if isEqualString(directive, LocationStubStatusDirective) {
		return true
	}
	return false
}

// if directive
func isLocationIfBlocksDirective(directive string) bool {
	return isEqualString(directive, LocationIfDirective)
}

// if directive
func isLocationSetDirective(directive string) bool {
	return isEqualString(directive, LocationSetDirective)
}

// add_header directive
func isLocationAddHeaderDirective(directive string) bool {
	return isEqualString(directive, AddHeaderDirective)
}

// determine whether it is a gzip directive
func isLocationGzipDirective(directive string) bool {
	return isEqualString(directive, LocationGzipDirective)
}

// determine whether it is a gzip_min_length directive
func isLocationGzipMinLengthDirective(directive string) bool {
	return isEqualString(directive, LocationGzipMinLengthDirective)
}

// determine whether it is a gzip_buffers directive
func isLocationGzipBuffersDirective(directive string) bool {
	return isEqualString(directive, LocationGzipBuffersDirective)
}

// determine whether it is a gzip_comp_level directive
func isLocationGzipCompLevelDirective(directive string) bool {
	return isEqualString(directive, LocationGzipCompLevelDirective)
}

// determine whether it is a gzip_types directive
func isLocationGzipTypesDirective(directive string) bool {
	return isEqualString(directive, LocationGzipTypesDirective)
}

// determine whether it is a gzip_vary directive
func isLocationGzipVaryDirective(directive string) bool {
	return isEqualString(directive, LocationGzipVaryDirective)
}

// determine whether it is a gzip_disable directive
func isLocationGzipDisableDirective(directive string) bool {
	return isEqualString(directive, LocationGzipDisableDirective)
}

// determine whether it is a gzip_http_version directive
func isLocationGzipHTTPVersionDirective(directive string) bool {
	return isEqualString(directive, LocationGzipHTTPVersionDirective)
}

// determine whether it is a gzip_proxied directive
func isLocationGzipProxiedDirective(directive string) bool {
	return isEqualString(directive, LocationGzipProxiedDirective)
}

// determine whether it is a proxy_hide_header directive
func isLocationProxyHideHeaderDirective(directive string) bool {
	return isEqualString(directive, LocationProxyHideHeaderDirective)
}

// ProcessLocation processLocation location module
func ProcessLocation(innerBlock *InnerBlock) (*Location, error) {
	if !isLocationDirective(innerBlock.Directive) {
		return nil, errors.New("not location directive")
	}

	location := NewLocation()
	// location condition need to add quotation marks
	args := processQuote(innerBlock.Args)
	// process Location Path field: location = /index.html {}
	location.Path = processArgsArray(args)

	// process location directive
	for _, inmostBlock := range innerBlock.InmostBlocks {
		// process proxy_buffers directive
		if isLocationProxyBuffersDirective(inmostBlock.Directive) {
			location.ProxyBuffers = processArgsArray(inmostBlock.Args)
			continue
		}
		// process proxy_cache_convert_head directive
		if isLocationProxyCacheConvertHeadDirective(inmostBlock.Directive) {
			location.ProxyCacheConvertHead = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_http_version directive
		if isLocationProxyHTTPVersionDirective(inmostBlock.Directive) {
			location.ProxyHttpVersion = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_min_uses directive
		if isLocationProxyCacheMinUsesDirective(inmostBlock.Directive) {
			location.ProxyCacheMinUses = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_pass_header directive
		if isLocationProxyPassHeaderDirective(inmostBlock.Directive) {
			location.ProxyPassHeader = processArgsArray(inmostBlock.Args)
			continue
		}

		// process disable_symlinks directive
		if isLocationDisableSymlinksDirective(inmostBlock.Directive) {
			location.DisableSymlinks = processArgsArray(inmostBlock.Args)
			continue
		}

		// process keepalive_timeout directive
		if isLocationKeepaliveTimeoutDirective(inmostBlock.Directive) {
			location.KeepaliveTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// process lingering_timeout directive
		if isLocationLingeringTimeoutDirective(inmostBlock.Directive) {
			location.LingeringTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// process sendfile directive
		if isLocationSendFileDirective(inmostBlock.Directive) {
			location.SendFile = processArgsArray(inmostBlock.Args)
			continue
		}

		// process error_page directive
		if isLocationErrorPageDirective(inmostBlock.Directive) {
			location.ErrorPage = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_key directive
		if isLocationProxyCacheKeyDirective(inmostBlock.Directive) {
			location.ProxyCacheKey = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_methods directive
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
		// process alias directive
		if isLocationAliasDirective(inmostBlock.Directive) {
			location.Alias = processArgsArray(inmostBlock.Args)
			continue
		}

		// process types_hash_max_size directive
		if isLocationTypesHashMaxSizeDirective(inmostBlock.Directive) {
			location.TypesHashMaxSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// process absolute_redirect directive
		if isLocationAbsoluteRedirectDirective(inmostBlock.Directive) {
			location.AbsoluteRedirect = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_send_lowat directive
		if isLocationProxySendLowatDirective(inmostBlock.Directive) {
			location.ProxySendLowat = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_store directive
		if isLocationProxyStoreDirective(inmostBlock.Directive) {
			location.ProxyStore = processArgsArray(inmostBlock.Args)
			continue
		}

		// process try_files directive
		if isLocationTryFilesDirective(inmostBlock.Directive) {
			location.TryFiles = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_busy_buffers_size directive
		if isLocationProxyBusyBuffersSizeDirective(inmostBlock.Directive) {
			location.ProxyBusyBuffersSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_next_upstream_timeout directive
		if isLocationProxyNextUpstreamTimeoutDirective(inmostBlock.Directive) {
			location.ProxyNextUpstreamTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// process client_body_temp_path directive
		if isLocationClientBodyTempPathDirective(inmostBlock.Directive) {
			location.ClientBodyTempPath = processArgsArray(inmostBlock.Args)
			continue
		}

		// process directio directive
		if isLocationDirectIODirective(inmostBlock.Directive) {
			location.DirectIO = processArgsArray(inmostBlock.Args)
			continue
		}

		// process satisfy directive
		if isLocationSatisfyDirective(inmostBlock.Directive) {
			location.Satisfy = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_store_access directive
		if isLocationProxyStoreAccessDirective(inmostBlock.Directive) {
			location.ProxyStoreAccess = processArgsArray(inmostBlock.Args)
			continue
		}

		// process expires directive
		if isLocationExpiresDirective(inmostBlock.Directive) {
			location.Expires = processArgsArray(inmostBlock.Args)
			continue
		}

		// process subrequest_output_buffer_size directive
		if isLocationSubrequestOutputBufferSizeDirective(inmostBlock.Directive) {
			location.SubrequestOutputBufferSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_lock directive
		if isLocationProxyCacheLockDirective(inmostBlock.Directive) {
			location.ProxyCacheLock = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_trusted_certificate directive
		if isLocationProxySslTrustedCertificateDirective(inmostBlock.Directive) {
			location.ProxySslTrustedCertificate = processArgsArray(inmostBlock.Args)
			continue
		}

		// process default_type directive
		if isLocationDefaultTypeDirective(inmostBlock.Directive) {
			location.DefaultType = processArgsArray(inmostBlock.Args)
			continue
		}

		// process directio_alignment directive
		if isLocationDirectioAlignmentDirective(inmostBlock.Directive) {
			location.DirectioAlignment = processArgsArray(inmostBlock.Args)
			continue
		}

		// process keepalive_disable directive
		if isLocationKeepaliveDisableDirective(inmostBlock.Directive) {
			location.KeepaliveDisable = processArgsArray(inmostBlock.Args)
			continue
		}

		// process read_ahead directive
		if isLocationReadAheadDirective(inmostBlock.Directive) {
			location.ReadAhead = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_verify directive
		if isLocationProxySslVerifyDirective(inmostBlock.Directive) {
			location.ProxySslVerify = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_set_body directive
		if isLocationProxySetBodyDirective(inmostBlock.Directive) {
			location.ProxySetBody = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_socket_keepalive directive
		if isLocationProxySocketKeepaliveDirective(inmostBlock.Directive) {
			location.ProxySocketKeepalive = processArgsArray(inmostBlock.Args)
			continue
		}

		// process client_max_body_size directive
		if isLocationClientMaxBodySizeDirective(inmostBlock.Directive) {
			location.ClientMaxBodySize = processArgsArray(inmostBlock.Args)
			continue
		}

		// process log_not_found directive
		if isLocationLogNotFoundDirective(inmostBlock.Directive) {
			location.LogNotFound = processArgsArray(inmostBlock.Args)
			continue
		}

		// process access_log directive
		if isLocationAccessLogDirective(inmostBlock.Directive) {
			location.AccessLog = processArgsArray(inmostBlock.Args)
			continue
		}

		// process resolver directive
		if isLocationResolverDirective(inmostBlock.Directive) {
			location.Resolver = processArgsArray(inmostBlock.Args)
			continue
		}

		// process server_name_in_redirect directive
		if isLocationServerNameInRedirectDirective(inmostBlock.Directive) {
			location.ServerNameInRedirect = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_read_timeout directive
		if isLocationProxyReadTimeoutDirective(inmostBlock.Directive) {
			location.ProxyReadTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// process limit_rate directive
		if isLocationLimitRateDirective(inmostBlock.Directive) {
			location.LimitRate = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_request_buffering directive
		if isLocationProxyRequestBufferingDirective(inmostBlock.Directive) {
			location.ProxyRequestBuffering = processArgsArray(inmostBlock.Args)
			continue
		}

		// process root directive
		if isLocationRootDirective(inmostBlock.Directive) {
			location.Root = processArgsArray(inmostBlock.Args)
			continue
		}

		// process client_body_buffer_size directive
		if isLocationClientBodyBufferSizeDirective(inmostBlock.Directive) {
			location.ClientBodyBufferSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// process max_ranges directive
		if isLocationMaxRangesDirective(inmostBlock.Directive) {
			location.MaxRanges = processArgsArray(inmostBlock.Args)
			continue
		}

		// process tcp_nopush directive
		if isLocationTCPNopushDirective(inmostBlock.Directive) {
			location.TcpNopush = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_no_cache directive
		if isLocationProxyNoCacheDirective(inmostBlock.Directive) {
			location.ProxyNoCache = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_password_file directive
		if isLocationProxySslPasswordFileDirective(inmostBlock.Directive) {
			location.ProxySslPasswordFile = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_headers_hash_max_size directive
		if isLocationProxyHeadersHashMaxSizeDirective(inmostBlock.Directive) {
			location.ProxyHeadersHashMaxSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_pass_request_headers directive
		if isLocationProxyPassRequestHeadersDirective(inmostBlock.Directive) {
			location.ProxyPassRequestHeaders = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_certificate_key directive
		if isLocationProxySslCertificateKeyDirective(inmostBlock.Directive) {
			location.ProxySslCertificateKey = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_verify_depth directive
		if isLocationProxySslVerifyDepthDirective(inmostBlock.Directive) {
			location.ProxySslVerifyDepth = processArgsArray(inmostBlock.Args)
			continue
		}

		// process keepalive_requests directive
		if isLocationKeepaliveRequestsDirective(inmostBlock.Directive) {
			location.KeepaliveRequests = processArgsArray(inmostBlock.Args)
			continue
		}

		// process open_file_cache directive
		if isLocationOpenFileCacheDirective(inmostBlock.Directive) {
			location.OpenFileCache = processArgsArray(inmostBlock.Args)
			continue
		}

		// process output_buffers directive
		if isLocationOutputBuffersDirective(inmostBlock.Directive) {
			location.OutputBuffers = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_buffering directive
		if isLocationProxyBufferingDirective(inmostBlock.Directive) {
			location.ProxyBuffering = processArgsArray(inmostBlock.Args)
			continue
		}

		// process recursive_error_pages directive
		if isLocationRecursiveErrorPagesDirective(inmostBlock.Directive) {
			location.RecursiveErrorPages = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_background_update directive
		if isLocationProxyCacheBackgroundUpdateDirective(inmostBlock.Directive) {
			location.ProxyCacheBackgroundUpdate = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_connect_timeout directive
		if isLocationProxyConnectTimeoutDirective(inmostBlock.Directive) {
			location.ProxyConnectTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// process client_body_timeout directive
		if isLocationClientBodyTimeoutDirective(inmostBlock.Directive) {
			location.ClientBodyTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_bind directive
		if isLocationProxyBindDirective(inmostBlock.Directive) {
			location.ProxyBind = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_lock_age directive
		if isLocationProxyCacheLockAgeDirective(inmostBlock.Directive) {
			location.ProxyCacheLockAge = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_server_name directive
		if isLocationProxySslServerNameDirective(inmostBlock.Directive) {
			location.ProxySslServerName = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache directive
		if isLocationProxyCacheDirective(inmostBlock.Directive) {
			location.ProxyCache = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_max_range_offset directive
		if isLocationProxyCacheMaxRangeOffsetDirective(inmostBlock.Directive) {
			location.ProxyCacheMaxRangeOffset = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_pass directive
		if isLocationProxyPassDirective(inmostBlock.Directive) {
			location.ProxyPass = processArgsArray(inmostBlock.Args)
			continue
		}

		// process chunked_transfer_encoding directive
		if isLocationChunkedTransferEncodingDirective(inmostBlock.Directive) {
			location.ChunkedTransferEncoding = processArgsArray(inmostBlock.Args)
			continue
		}

		// process tcp_nodelay directive
		if isLocationTCPNodelayDirective(inmostBlock.Directive) {
			location.TcpNodelay = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_lock_timeout directive
		if isLocationProxyCacheLockTimeoutDirective(inmostBlock.Directive) {
			location.ProxyCacheLockTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_intercept_errors directive
		if isLocationProxyInterceptErrorsDirective(inmostBlock.Directive) {
			location.ProxyInterceptErrors = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_redirect directive
		if isLocationProxyRedirectDirective(inmostBlock.Directive) {
			location.ProxyRedirect = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_path directive
		if isLocationProxyCachePathDirective(inmostBlock.Directive) {
			location.ProxyCachePath = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_name directive
		if isLocationProxySslNameDirective(inmostBlock.Directive) {
			location.ProxySslName = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_use_stale directive
		if isLocationProxyCacheUseStaleDirective(inmostBlock.Directive) {
			location.ProxyCacheUseStale = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_max_temp_file_size directive
		if isLocationProxyMaxTempFileSizeDirective(inmostBlock.Directive) {
			location.ProxyMaxTempFileSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_next_upstream directive
		if isLocationProxyNextUpstreamDirective(inmostBlock.Directive) {
			location.ProxyNextUpstream = processArgsArray(inmostBlock.Args)
			continue
		}

		// process limit_rate_after directive
		if isLocationLimitRateAfterDirective(inmostBlock.Directive) {
			location.LimitRateAfter = processArgsArray(inmostBlock.Args)
			continue
		}

		// process lingering_close directive
		if isLocationLingeringCloseDirective(inmostBlock.Directive) {
			location.LingeringClose = processArgsArray(inmostBlock.Args)
			continue
		}

		// process server_tokens directive
		if isLocationServerTokensDirective(inmostBlock.Directive) {
			location.ServerTokens = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_protocols directive
		if isLocationProxySslProtocolsDirective(inmostBlock.Directive) {
			location.ProxySslProtocols = processArgsArray(inmostBlock.Args)
			continue
		}

		// process aio_write directive
		if isLocationAioWriteDirective(inmostBlock.Directive) {
			location.AioWrite = processArgsArray(inmostBlock.Args)
			continue
		}

		// process client_body_in_file_only directive
		if isLocationClientBodyInFileOnlyDirective(inmostBlock.Directive) {
			location.ClientBodyInFileOnly = processArgsArray(inmostBlock.Args)
			continue
		}

		// process open_file_cache_errors directive
		if isLocationOpenFileCacheErrorsDirective(inmostBlock.Directive) {
			location.OpenFileCacheErrors = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_next_upstream_tries directive
		if isLocationProxyNextUpstreamTriesDirective(inmostBlock.Directive) {
			location.ProxyNextUpstreamTries = processArgsArray(inmostBlock.Args)
			continue
		}

		// process lingering_time directive
		if isLocationLingeringTimeDirective(inmostBlock.Directive) {
			location.LingeringTime = processArgsArray(inmostBlock.Args)
			continue
		}

		// process open_file_cache_valid directive
		if isLocationOpenFileCacheValidDirective(inmostBlock.Directive) {
			location.OpenFileCacheValid = processArgsArray(inmostBlock.Args)
			continue
		}

		// process postpone_output directive
		if isLocationPostponeOutputDirective(inmostBlock.Directive) {
			location.PostponeOutput = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_limit_rate directive
		if isLocationProxyLimitRateDirective(inmostBlock.Directive) {
			location.ProxyLimitRate = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_temp_file_write_size directive
		if isLocationProxyTempFileWriteSizeDirective(inmostBlock.Directive) {
			location.ProxyTempFileWriteSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// process internal directive
		if isLocationInternalDirective(inmostBlock.Directive) {
			location.Internal = processArgsArray(inmostBlock.Args)
			continue
		}

		// process sendfile_max_chunk directive
		if isLocationSendFileMaxChunkDirective(inmostBlock.Directive) {
			location.SendFileMaxChunk = processArgsArray(inmostBlock.Args)
			continue
		}

		// process port_in_redirect directive
		if isLocationPortInRedirectDirective(inmostBlock.Directive) {
			location.PortInRedirect = processArgsArray(inmostBlock.Args)
			continue
		}

		// process resolver_timeout directive
		if isLocationResolverTimeoutDirective(inmostBlock.Directive) {
			location.ResolverTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_purge directive
		if isLocationProxyCachePurgeDirective(inmostBlock.Directive) {
			location.ProxyCachePurge = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_headers_hash_bucket_size directive
		if isLocationProxyHeadersHashBucketSizeDirective(inmostBlock.Directive) {
			location.ProxyHeadersHashBucketSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_send_timeout directive
		if isLocationProxySendTimeoutDirective(inmostBlock.Directive) {
			location.ProxySendTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// process client_body_in_single_buffer directive
		if isLocationClientBodyInSingleBufferDirective(inmostBlock.Directive) {
			location.ClientBodyInSingleBuffer = processArgsArray(inmostBlock.Args)
			continue
		}

		// process send_lowat directive
		if isLocationSendLowatDirective(inmostBlock.Directive) {
			location.SendLowat = processArgsArray(inmostBlock.Args)
			continue
		}

		// process send_timeout directive
		if isLocationSendTimeoutDirective(inmostBlock.Directive) {
			location.SendTimeout = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ignore_client_abort directive
		if isLocationProxyIgnoreClientAbortDirective(inmostBlock.Directive) {
			location.ProxyIgnoreClientAbort = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_buffer_size directive
		if isLocationProxyBufferSizeDirective(inmostBlock.Directive) {
			location.ProxyBufferSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_certificate directive
		if isLocationProxySslCertificateDirective(inmostBlock.Directive) {
			location.ProxySslCertificate = processArgsArray(inmostBlock.Args)
			continue
		}

		// process msie_padding directive
		if isLocationMsiePaddingDirective(inmostBlock.Directive) {
			location.MsiePadding = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_force_ranges directive
		if isLocationProxyForceRangesDirective(inmostBlock.Directive) {
			location.ProxyForceRanges = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_pass_request_body directive
		if isLocationProxyPassRequestBodyDirective(inmostBlock.Directive) {
			location.ProxyPassRequestBody = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_crl directive
		if isLocationProxySslCrlDirective(inmostBlock.Directive) {
			location.ProxySslCrl = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_temp_path directive
		if isLocationProxyTempPathDirective(inmostBlock.Directive) {
			location.ProxyTempPath = processArgsArray(inmostBlock.Args)
			continue
		}
		// process stub_status directive
		if isLocationStubStatusDirective(inmostBlock.Directive) {
			location.StubStatus = processArgsArray(inmostBlock.Args)
			continue
		}

		// process etag directive
		if isLocationEtagDirective(inmostBlock.Directive) {
			location.Etag = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cookie_domain directive
		if isLocationProxyCookieDomainDirective(inmostBlock.Directive) {
			location.ProxyCookieDomain = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cookie_path directive
		if isLocationProxyCookiePathDirective(inmostBlock.Directive) {
			location.ProxyCookiePath = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ignore_headers directive
		if isLocationProxyIgnoreHeadersDirective(inmostBlock.Directive) {
			location.ProxyIgnoreHeaders = processArgsArray(inmostBlock.Args)
			continue
		}

		// process aio directive
		if isLocationAioDirective(inmostBlock.Directive) {
			location.Aio = processArgsArray(inmostBlock.Args)
			continue
		}

		// process types_hash_bucket_size directive
		if isLocationTypesHashBucketSizeDirective(inmostBlock.Directive) {
			location.TypesHashBucketSize = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_method directive
		if isLocationProxyMethodDirective(inmostBlock.Directive) {
			location.ProxyMethod = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_ciphers directive
		if isLocationProxySslCiphersDirective(inmostBlock.Directive) {
			location.ProxySslCiphers = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_ssl_session_reuse directive
		if isLocationProxySslSessionReuseDirective(inmostBlock.Directive) {
			location.ProxySslSessionReuse = processArgsArray(inmostBlock.Args)
			continue
		}

		// process gzip directive
		if isLocationGzipDirective(inmostBlock.Directive) {
			location.Gzip = processArgsArray(inmostBlock.Args)
			continue
		}

		// process gzip_min_length directive
		if isLocationGzipMinLengthDirective(inmostBlock.Directive) {
			location.GzipMinLength = processArgsArray(inmostBlock.Args)
			continue
		}

		// process gzip_buffers directive
		if isLocationGzipBuffersDirective(inmostBlock.Directive) {
			location.GzipBuffers = processArgsArray(inmostBlock.Args)
			continue
		}

		// process gzip_comp_level directive
		if isLocationGzipCompLevelDirective(inmostBlock.Directive) {
			location.GzipCompLevel = processArgsArray(inmostBlock.Args)
			continue
		}

		// process gzip_types directive
		if isLocationGzipTypesDirective(inmostBlock.Directive) {
			location.GzipTypes = processArgsArray(inmostBlock.Args)
			continue
		}

		// process gzip_vary directive
		if isLocationGzipVaryDirective(inmostBlock.Directive) {
			location.GzipVary = processArgsArray(inmostBlock.Args)
			continue
		}

		// process gzip_disable directive
		if isLocationGzipDisableDirective(inmostBlock.Directive) {
			location.GzipDisable = processArgsArray(inmostBlock.Args)
			continue
		}

		// process gzip_http_version directive
		if isLocationGzipHTTPVersionDirective(inmostBlock.Directive) {
			location.GzipHttpVersion = processArgsArray(inmostBlock.Args)
			continue
		}

		// process gzip_proxied directive
		if isLocationGzipProxiedDirective(inmostBlock.Directive) {
			location.GzipProxied = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_hide_header directive
		if isLocationProxyHideHeaderDirective(inmostBlock.Directive) {
			location.ProxyHideHeader = processArgsArray(inmostBlock.Args)
			continue
		}

		// if isLocationIfBlocksDirective(inmostBlock.Directive) {
		// location.IfBlock = append(location.IfBlock, processArgsArray(inmostBlock.Args))

		// }

		// process rewrite directive
		if isLocationRewriteRootDirective(inmostBlock.Directive) {
			args := processQuote(inmostBlock.Args)
			location.Rewrite = append(location.Rewrite, processArgsArray(args))
			continue
		}

		// process proxy_cache_bypass directive
		if isLocationProxyCacheBypassDirective(inmostBlock.Directive) {
			location.ProxyCacheBypass = append(location.ProxyCacheBypass, processArgsArray(inmostBlock.Args))
			continue
		}

		// process proxy_cache directive
		if isLocationProxyCacheDirective(inmostBlock.Directive) {
			location.ProxyCache = processArgsArray(inmostBlock.Args)
			continue
		}

		// process proxy_cache_valid directive
		if isLocationProxyCacheValidDirective(inmostBlock.Directive) {
			location.ProxyCacheValid = append(location.ProxyCacheValid, processArgsArray(inmostBlock.Args))
			continue
		}

		// process allow directive
		if isLocationAllowDirective(inmostBlock.Directive) {
			location.Allow = append(location.Allow, processArgsArray(inmostBlock.Args))
			continue
		}

		// process proxy_set_header directive
		if isLocationProxySetHeaderDirective(inmostBlock.Directive) {
			location.ProxySetHeader = append(location.ProxySetHeader, processArgsArray(inmostBlock.Args))
			continue
		}

		// process deny directive
		// if isLocationDenyDirective(inmostBlock.Directive) {
		// 	location.Deny = append(location.Deny, processArgsArray(inmostBlock.Args))
		// 	continue
		// }

		// process deny directive
		if isLocationDenyDirective(inmostBlock.Directive) {
			location.Deny = append(location.Deny, processArgsArray(inmostBlock.Args))
			continue
		}
		// process set directive
		if isLocationSetDirective(inmostBlock.Directive) {
			location.Set = append(location.Set, processArgsArray(inmostBlock.Args))
		}

		// process log_format directive
		if isLocationLogFormatDirective(inmostBlock.Directive) {
			location.LogFormat = append(location.LogFormat, processArgsArray(inmostBlock.Args))
			continue
		}
		// process if directive
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
			// add_header directive need to add quotation marks
			for i := range inmostBlock.Args {
				inmostBlock.Args[i] = nonescapeQuotation(inmostBlock.Args[i])
			}
			args := processQuote(inmostBlock.Args)
			location.AddHeader = append(location.AddHeader, processArgsArray(args))
		}

	}
	return location, nil
}
