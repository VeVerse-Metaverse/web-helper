package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type PlayerResponse struct {
	ResponseContext   ResponseContext   `json:"responseContext"`
	PlayabilityStatus PlayabilityStatus `json:"playabilityStatus"`
	StreamingData     StreamingData     `json:"streamingData"`
	PlaybackTracking  PlaybackTracking  `json:"playbackTracking"`
	VideoDetails      VideoDetails      `json:"videoDetails"`
	PlayerConfig      PlayerConfig      `json:"playerConfig"`
}

type ResponseContext struct {
	VisitorData string `json:"visitorData"`
}

type PlayabilityStatus struct {
	Status          string `json:"status"`
	PlayableInEmbed bool   `json:"playableInEmbed"`
}

type ColorInfo struct {
	Primaries               string `json:"primaries,omitempty"`
	TransferCharacteristics string `json:"transferCharacteristics,omitempty"`
	MatrixCoefficients      string `json:"matrixCoefficients,omitempty"`
}

type Range struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Format struct {
	Itag             int       `json:"itag"`
	URL              string    `json:"url"`
	MimeType         string    `json:"mimeType"`
	Bitrate          int       `json:"bitrate"`
	Width            int       `json:"width"`
	Height           int       `json:"height"`
	LastModified     string    `json:"lastModified"`
	ContentLength    string    `json:"contentLength"`
	Quality          string    `json:"quality"`
	FPS              int       `json:"fps"`
	QualityLabel     string    `json:"qualityLabel"`
	ProjectionType   string    `json:"projectionType"`
	AverageBitrate   int       `json:"averageBitrate"`
	AudioQuality     string    `json:"audioQuality,omitempty"`
	ApproxDurationMs string    `json:"approxDurationMs"`
	AudioSampleRate  string    `json:"audioSampleRate,omitempty"`
	AudioChannels    int       `json:"audioChannels,omitempty"`
	HighReplication  bool      `json:"highReplication,omitempty"`
	ColorInfo        ColorInfo `json:"colorInfo,omitempty"`
	InitRange        Range     `json:"initRange,omitempty"`
	IndexRange       Range     `json:"indexRange,omitempty"`
}

type AdaptiveFormat struct {
	Itag             int       `json:"itag"`
	URL              string    `json:"url"`
	MimeType         string    `json:"mimeType"`
	Bitrate          int       `json:"bitrate"`
	Width            int       `json:"width"`
	Height           int       `json:"height"`
	InitRange        Range     `json:"initRange,omitempty"`
	IndexRange       Range     `json:"indexRange,omitempty"`
	LastModified     string    `json:"lastModified"`
	ContentLength    string    `json:"contentLength"`
	Quality          string    `json:"quality"`
	FPS              int       `json:"fps"`
	QualityLabel     string    `json:"qualityLabel"`
	ProjectionType   string    `json:"projectionType"`
	AverageBitrate   int       `json:"averageBitrate"`
	ColorInfo        ColorInfo `json:"colorInfo,omitempty"`
	ApproxDurationMs string    `json:"approxDurationMs"`
}

type StreamingData struct {
	ExpiresInSeconds string           `json:"expiresInSeconds"`
	Formats          []Format         `json:"formats"`
	AdaptiveFormats  []AdaptiveFormat `json:"adaptiveFormats"`
}

type TrackingUrlHeader struct {
	HeaderType string `json:"headerType"`
}

type TrackingUrl struct {
	BaseUrl string              `json:"baseUrl"`
	Headers []TrackingUrlHeader `json:"headers"`
}

type PlaybackTracking struct {
	VideostatsPlaybackUrl  TrackingUrl `json:"videostatsPlaybackUrl"`
	VideostatsDelayplayUrl TrackingUrl `json:"videostatsDelayplayUrl"`
	VideostatsWatchtimeUrl TrackingUrl `json:"videostatsWatchtimeUrl"`
	PtrackingUrl           TrackingUrl `json:"ptrackingUrl"`
	QoeUrl                 TrackingUrl `json:"qoeUrl"`
}

type ThumbnailList struct {
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type Thumbnail struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type VideoDetails struct {
	VideoId           string        `json:"videoId"`
	Title             string        `json:"title"`
	LengthSeconds     string        `json:"lengthSeconds"`
	Keywords          []string      `json:"keywords"`
	ChannelId         string        `json:"channelId"`
	IsOwnerViewing    bool          `json:"isOwnerViewing"`
	ShortDescription  string        `json:"shortDescription"`
	IsCrawlable       bool          `json:"isCrawlable"`
	Thumbnail         ThumbnailList `json:"thumbnail"`
	AllowRatings      bool          `json:"allowRatings"`
	ViewCount         string        `json:"viewCount"`
	Author            string        `json:"author"`
	IsPrivate         bool          `json:"isPrivate"`
	IsUnpluggedCorpus bool          `json:"isUnpluggedCorpus"`
	IsLiveContent     bool          `json:"isLiveContent"`
}

type AudioConfig struct {
	LoudnessDb              float64 `json:"loudnessDb"`
	PerceptualLoudnessDb    float64 `json:"perceptualLoudnessDb"`
	EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
}

type ExoplayerConfig struct {
	UseExoPlayer                                     bool     `json:"useExoPlayer"`
	UseAdaptiveBitrate                               bool     `json:"useAdaptiveBitrate"`
	MaxInitialByteRate                               int      `json:"maxInitialByteRate"`
	MinDurationForQualityIncreaseMs                  int      `json:"minDurationForQualityIncreaseMs"`
	MaxDurationForQualityDecreaseMs                  int      `json:"maxDurationForQualityDecreaseMs"`
	MinDurationToRetainAfterDiscardMs                int      `json:"minDurationToRetainAfterDiscardMs"`
	LowWatermarkMs                                   int      `json:"lowWatermarkMs"`
	HighWatermarkMs                                  int      `json:"highWatermarkMs"`
	LowPoolLoad                                      float64  `json:"lowPoolLoad"`
	HighPoolLoad                                     float64  `json:"highPoolLoad"`
	SufficientBandwidthOverhead                      float64  `json:"sufficientBandwidthOverhead"`
	BufferChunkSizeKb                                int      `json:"bufferChunkSizeKb"`
	HttpConnectTimeoutMs                             int      `json:"httpConnectTimeoutMs"`
	HttpReadTimeoutMs                                int      `json:"httpReadTimeoutMs"`
	NumAudioSegmentsPerFetch                         int      `json:"numAudioSegmentsPerFetch"`
	NumVideoSegmentsPerFetch                         int      `json:"numVideoSegmentsPerFetch"`
	MinDurationForPlaybackStartMs                    int      `json:"minDurationForPlaybackStartMs"`
	EnableExoplayerReuse                             bool     `json:"enableExoplayerReuse"`
	UseRadioTypeForInitialQualitySelection           bool     `json:"useRadioTypeForInitialQualitySelection"`
	BlacklistFormatOnError                           bool     `json:"blacklistFormatOnError"`
	EnableBandaidHttpDataSource                      bool     `json:"enableBandaidHttpDataSource"`
	HttpLoadTimeoutMs                                int      `json:"httpLoadTimeoutMs"`
	CanPlayHdDrm                                     bool     `json:"canPlayHdDrm"`
	VideoBufferSegmentCount                          int      `json:"videoBufferSegmentCount"`
	AudioBufferSegmentCount                          int      `json:"audioBufferSegmentCount"`
	UseAbruptSplicing                                bool     `json:"useAbruptSplicing"`
	MinRetryCount                                    int      `json:"minRetryCount"`
	MinChunksNeededToPreferOffline                   int      `json:"minChunksNeededToPreferOffline"`
	SecondsToMaxAggressiveness                       int      `json:"secondsToMaxAggressiveness"`
	EnableSurfaceviewResizeWorkaround                bool     `json:"enableSurfaceviewResizeWorkaround"`
	EnableVp9IfThresholdsPass                        bool     `json:"enableVp9IfThresholdsPass"`
	MatchQualityToViewportOnUnfullscreen             bool     `json:"matchQualityToViewportOnUnfullscreen"`
	LowAudioQualityConnTypes                         []string `json:"lowAudioQualityConnTypes"`
	UseDashForLiveStreams                            bool     `json:"useDashForLiveStreams"`
	EnableLibvpxVideoTrackRenderer                   bool     `json:"enableLibvpxVideoTrackRenderer"`
	LowAudioQualityBandwidthThresholdBps             int      `json:"lowAudioQualityBandwidthThresholdBps"`
	EnableVariableSpeedPlayback                      bool     `json:"enableVariableSpeedPlayback"`
	PreferOnesieBufferedFormat                       bool     `json:"preferOnesieBufferedFormat"`
	MinimumBandwidthSampleBytes                      int      `json:"minimumBandwidthSampleBytes"`
	UseDashForOtfAndCompletedLiveStreams             bool     `json:"useDashForOtfAndCompletedLiveStreams"`
	DisableCacheAwareVideoFormatEvaluation           bool     `json:"disableCacheAwareVideoFormatEvaluation"`
	UseLiveDvrForDashLiveStreams                     bool     `json:"useLiveDvrForDashLiveStreams"`
	CronetResetTimeoutOnRedirects                    bool     `json:"cronetResetTimeoutOnRedirects"`
	EmitVideoDecoderChangeEvents                     bool     `json:"emitVideoDecoderChangeEvents"`
	OnesieVideoBufferLoadTimeoutMs                   string   `json:"onesieVideoBufferLoadTimeoutMs"`
	OnesieVideoBufferReadTimeoutMs                   string   `json:"onesieVideoBufferReadTimeoutMs"`
	LibvpxEnableGl                                   bool     `json:"libvpxEnableGl"`
	EnableVp9EncryptedIfThresholdsPass               bool     `json:"enableVp9EncryptedIfThresholdsPass"`
	EnableOpus                                       bool     `json:"enableOpus"`
	UsePredictedBuffer                               bool     `json:"usePredictedBuffer"`
	MaxReadAheadMediaTimeMs                          int      `json:"maxReadAheadMediaTimeMs"`
	UseMediaTimeCappedLoadControl                    bool     `json:"useMediaTimeCappedLoadControl"`
	AllowCacheOverrideToLowerQualitiesWithinRange    int      `json:"allowCacheOverrideToLowerQualitiesWithinRange"`
	AllowDroppingUndecodedFrames                     bool     `json:"allowDroppingUndecodedFrames"`
	MinDurationForPlaybackRestartMs                  int      `json:"minDurationForPlaybackRestartMs"`
	ServerProvidedBandwidthHeader                    string   `json:"serverProvidedBandwidthHeader"`
	LiveOnlyPegStrategy                              string   `json:"liveOnlyPegStrategy"`
	EnableRedirectorHostFallback                     bool     `json:"enableRedirectorHostFallback"`
	EnableHighlyAvailableFormatFallbackOnPcr         bool     `json:"enableHighlyAvailableFormatFallbackOnPcr"`
	RecordTrackRendererTimingEvents                  bool     `json:"recordTrackRendererTimingEvents"`
	MinErrorsForRedirectorHostFallback               int      `json:"minErrorsForRedirectorHostFallback"`
	NonHardwareMediaCodecNames                       []string `json:"nonHardwareMediaCodecNames"`
	EnableVp9IfInHardware                            bool     `json:"enableVp9IfInHardware"`
	EnableVp9EncryptedIfInHardware                   bool     `json:"enableVp9EncryptedIfInHardware"`
	UseOpusMedAsLowQualityAudio                      bool     `json:"useOpusMedAsLowQualityAudio"`
	MinErrorsForPcrFallback                          int      `json:"minErrorsForPcrFallback"`
	UseStickyRedirectHttpDataSource                  bool     `json:"useStickyRedirectHttpDataSource"`
	OnlyVideoBandwidth                               bool     `json:"onlyVideoBandwidth"`
	UseRedirectorOnNetworkChange                     bool     `json:"useRedirectorOnNetworkChange"`
	EnableMaxReadaheadAbrThreshold                   bool     `json:"enableMaxReadaheadAbrThreshold"`
	CacheCheckDirectoryWritabilityOnce               bool     `json:"cacheCheckDirectoryWritabilityOnce"`
	PredictorType                                    string   `json:"predictorType"`
	SlidingPercentile                                float64  `json:"slidingPercentile"`
	SlidingWindowSize                                int      `json:"slidingWindowSize"`
	MaxFrameDropIntervalMs                           int      `json:"maxFrameDropIntervalMs"`
	IgnoreLoadTimeoutForFallback                     bool     `json:"ignoreLoadTimeoutForFallback"`
	ServerBweMultiplier                              int      `json:"serverBweMultiplier"`
	DrmMaxKeyfetchDelayMs                            int      `json:"drmMaxKeyfetchDelayMs"`
	MaxResolutionForWhiteNoise                       int      `json:"maxResolutionForWhiteNoise"`
	WhiteNoiseRenderEffectMode                       string   `json:"whiteNoiseRenderEffectMode"`
	EnableLibvpxHdr                                  bool     `json:"enableLibvpxHdr"`
	EnableCacheAwareStreamSelection                  bool     `json:"enableCacheAwareStreamSelection"`
	UseExoCronetDataSource                           bool     `json:"useExoCronetDataSource"`
	WhiteNoiseScale                                  int      `json:"whiteNoiseScale"`
	WhiteNoiseOffset                                 int      `json:"whiteNoiseOffset"`
	PreventVideoFrameLaggingWithLibvpx               bool     `json:"preventVideoFrameLaggingWithLibvpx"`
	EnableMediaCodecHdr                              bool     `json:"enableMediaCodecHdr"`
	EnableMediaCodecSwHdr                            bool     `json:"enableMediaCodecSwHdr"`
	LiveOnlyWindowChunks                             int      `json:"liveOnlyWindowChunks"`
	BearerMinDurationToRetainAfterDiscardMs          []int    `json:"bearerMinDurationToRetainAfterDiscardMs"`
	ForceWidevineL3                                  bool     `json:"forceWidevineL3"`
	UseAverageBitrate                                bool     `json:"useAverageBitrate"`
	UseMedialibAudioTrackRendererForLive             bool     `json:"useMedialibAudioTrackRendererForLive"`
	UseExoPlayerV2                                   bool     `json:"useExoPlayerV2"`
	LogMediaRequestEventsToCsi                       bool     `json:"logMediaRequestEventsToCsi"`
	OnesieFixNonZeroStartTimeFormatSelection         bool     `json:"onesieFixNonZeroStartTimeFormatSelection"`
	LiveOnlyReadaheadStepSizeChunks                  int      `json:"liveOnlyReadaheadStepSizeChunks"`
	LiveOnlyBufferHealthHalfLifeSeconds              int      `json:"liveOnlyBufferHealthHalfLifeSeconds"`
	LiveOnlyMinBufferHealthRatio                     float64  `json:"liveOnlyMinBufferHealthRatio"`
	LiveOnlyMinLatencyToSeekRatio                    int      `json:"liveOnlyMinLatencyToSeekRatio"`
	ManifestlessPartialChunkStrategy                 string   `json:"manifestlessPartialChunkStrategy"`
	IgnoreViewportSizeWhenSticky                     bool     `json:"ignoreViewportSizeWhenSticky"`
	EnableLibvpxFallback                             bool     `json:"enableLibvpxFallback"`
	DisableLibvpxLoopFilter                          bool     `json:"disableLibvpxLoopFilter"`
	EnableVpxMediaView                               bool     `json:"enableVpxMediaView"`
	HdrMinScreenBrightness                           int      `json:"hdrMinScreenBrightness"`
	HdrMaxScreenBrightnessThreshold                  int      `json:"hdrMaxScreenBrightnessThreshold"`
	OnesieDataSourceAboveCacheDataSource             bool     `json:"onesieDataSourceAboveCacheDataSource"`
	HttpNonplayerLoadTimeoutMs                       int      `json:"httpNonplayerLoadTimeoutMs"`
	NumVideoSegmentsPerFetchStrategy                 string   `json:"numVideoSegmentsPerFetchStrategy"`
	MaxVideoDurationPerFetchMs                       int      `json:"maxVideoDurationPerFetchMs"`
	MaxVideoEstimatedLoadDurationMs                  int      `json:"maxVideoEstimatedLoadDurationMs"`
	EstimatedServerClockHalfLife                     int      `json:"estimatedServerClockHalfLife"`
	EstimatedServerClockStrictOffset                 bool     `json:"estimatedServerClockStrictOffset"`
	MinReadAheadMediaTimeMs                          int      `json:"minReadAheadMediaTimeMs"`
	ReadAheadGrowthRate                              int      `json:"readAheadGrowthRate"`
	UseDynamicReadAhead                              bool     `json:"useDynamicReadAhead"`
	UseYtVodMediaSourceForV2                         bool     `json:"useYtVodMediaSourceForV2"`
	EnableV2Gapless                                  bool     `json:"enableV2Gapless"`
	UseLiveHeadTimeMillis                            bool     `json:"useLiveHeadTimeMillis"`
	AllowTrackSelectionWithUpdatedVideoItagsForExoV2 bool     `json:"allowTrackSelectionWithUpdatedVideoItagsForExoV2"`
	MaxAllowableTimeBeforeMediaTimeUpdateSec         int      `json:"maxAllowableTimeBeforeMediaTimeUpdateSec"`
	EnableDynamicHdr                                 bool     `json:"enableDynamicHdr"`
	V2PerformEarlyStreamSelection                    bool     `json:"v2PerformEarlyStreamSelection"`
	V2UsePlaybackStreamSelectionResult               bool     `json:"v2UsePlaybackStreamSelectionResult"`
	V2MinTimeBetweenAbrReevaluationMs                int      `json:"v2MinTimeBetweenAbrReevaluationMs"`
	AvoidReusePlaybackAcrossLoadvideos               bool     `json:"avoidReusePlaybackAcrossLoadvideos"`
	EnableInfiniteNetworkLoadingRetries              bool     `json:"enableInfiniteNetworkLoadingRetries"`
	ReportExoPlayerStateOnTransition                 bool     `json:"reportExoPlayerStateOnTransition"`
	ManifestlessSequenceMethod                       string   `json:"manifestlessSequenceMethod"`
	UseLiveHeadWindow                                bool     `json:"useLiveHeadWindow"`
	EnableDynamicHdrInHardware                       bool     `json:"enableDynamicHdrInHardware"`
	UltralowAudioQualityBandwidthThresholdBps        int      `json:"ultralowAudioQualityBandwidthThresholdBps"`
	IgnoreUnneededSeeksToLiveHead                    bool     `json:"ignoreUnneededSeeksToLiveHead"`
	DrmMetricsQoeLoggingFraction                     float64  `json:"drmMetricsQoeLoggingFraction"`
	UseTimeSeriesBufferPrediction                    bool     `json:"useTimeSeriesBufferPrediction"`
	SlidingPercentileScalar                          int      `json:"slidingPercentileScalar"`
}

type PlayerConfig struct {
	AudioConfig     AudioConfig     `json:"audioConfig"`
	ExoPlayerConfig ExoplayerConfig `json:"exoPlayerConfig"`
}

func GetPlayerResponse(videoID string) (*PlayerResponse, error) {
	url := "https://www.youtube.com/youtubei/v1/player"
	requestBody := fmt.Sprintf(`{
		"videoId": "%s",
		"context": {
			"client": {
				"clientName": "ANDROID_TESTSUITE",
				"clientVersion": "1.9",
				"androidSdkVersion": 30,
				"hl": "en",
				"gl": "US",
				"utcOffsetMinutes": 0
			}
		}
	}`, videoID)

	request, err := http.NewRequest("POST", url, bytes.NewBufferString(requestBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "com.google.android.youtube/17.36.4 (Linux; U; Android 12; GB) gzip")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", response.StatusCode)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	playerResponse, err := parsePlayerResponse(responseBody)
	if err != nil {
		return nil, err
	}

	return playerResponse, nil
}

func parsePlayerResponse(responseBody []byte) (*PlayerResponse, error) {
	playerResponse := &PlayerResponse{}
	err := json.Unmarshal(responseBody, &playerResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse player response: %v", err)
	}

	return playerResponse, nil
}
