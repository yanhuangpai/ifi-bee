// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import "github.com/yanhuangpai/ifi-bee/pkg/swarm"

type Server = server

type (
	BytesPostResponse        = bytesPostResponse
	ChunkAddressResponse     = chunkAddressResponse
	SocPostResponse          = socPostResponse
	FeedReferenceResponse    = feedReferenceResponse
	FileUploadResponse       = fileUploadResponse
	TagResponse              = tagResponse
	TagRequest               = tagRequest
	ListTagsResponse         = listTagsResponse
	PinnedChunk              = pinnedChunk
	ListPinnedChunksResponse = listPinnedChunksResponse
	UpdatePinCounter         = updatePinCounter
)

var (
	ContentTypeTar = contentTypeTar
)

var (
	ManifestRootPath                      = manifestRootPath
	ManifestWebsiteIndexDocumentSuffixKey = manifestWebsiteIndexDocumentSuffixKey
	ManifestWebsiteErrorDocumentPathKey   = manifestWebsiteErrorDocumentPathKey
)

var (
	ErrNoResolver           = errNoResolver
	ErrInvalidNameOrAddress = errInvalidNameOrAddress
)

var (
	FeedMetadataEntryOwner = feedMetadataEntryOwner
	FeedMetadataEntryTopic = feedMetadataEntryTopic
	FeedMetadataEntryType  = feedMetadataEntryType
)

func (s *Server) ResolveNameOrAddress(str string) (swarm.Address, error) {
	return s.resolveNameOrAddress(str)
}

func CalculateNumberOfChunks(contentLength int64, isEncrypted bool) int64 {
	return calculateNumberOfChunks(contentLength, isEncrypted)
}
