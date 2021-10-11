package main

import (
	"github.com/bitmovin/bitmovin-go/bitmovin"
	"github.com/bitmovin/bitmovin-go/bitmovintypes"
	"github.com/bitmovin/bitmovin-go/models"
	"github.com/bitmovin/bitmovin-go/services"
)

func main (){
	// _ = os.Setenv("GCS_BUCKET", "bitmovin-demo")
	// _ = os.Setenv("GCS_BUCKET_INPUT", "bitmovin-demo")

	// file, err := os.Open("./INPUT_sea.mp4")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// filePath := common.GetFileNameWithoutExt(file.Name()) + "/"

	// if err := repository.GCSUpload(file, filePath, file.Name()); err != nil {
	// 	panic(err)
	// }

	//1.BitmovinAPIクライアントの初期化（アカウントとの紐付け）
	bitmovin := bitmovin.NewBitmovin(
		"76dabab9-e4ad-43e3-80ab-9019bb877ece",
		"https://api.bitmovin.com/v1/", 5,
	)

	//BitmovinにGCSの動画ファイル（bitmovin-demo/INPUT_sea/sea.mp4）の情報を設定
	name := "bitmovin-demo"
	accessKey := "GOOG1E66L3YSOFOA4T3MMBO2JYAPQ2L4MO44CI2RIZETPI27TDTO46M6IL6NY"
	secretKey := "BV3Z0aoeLvnfczjudxZhg58TrXWfKpVw6+h0r4rL"
	bucketName := "bitmovin-demo"

	inputService := services.NewGCSInputService(bitmovin)
	input := &models.GCSInput{
		Name:       &name,
		AccessKey:  &accessKey,
		SecretKey:  &secretKey,
		BucketName: &bucketName,
	}
	inputResp, err := inputService.Create(input)
	if err != nil {
		panic(err)
	}

	//GCSOutputでGCSに出力の情報を作成
	outpuService := services.NewGCSOutputService(bitmovin)
	output := &models.GCSOutput{
		Name:        &name,
		AccessKey:   &accessKey,
		SecretKey:   &secretKey,
		BucketName:  &bucketName,
		CloudRegion: bitmovintypes.GoogleCloudRegionUSEast1,
	}
	outputResp, err := outpuService.Create(output)
	if err != nil {
		panic(err)
	}
	outputId := outputResp.Data.Result.ID
	// outputId := "f22f2b27"

	//H264を使用しコーデック構成情報を作成（引数：Name, Bitrate, Width, Profile）
	codecConfigVideoBitrate1 := int64(1500000)
	codecConfigVideoWidth1 := int64(1024)
	h264S1 := services.NewH264CodecConfigurationService(bitmovin)
	videoCodecConfiguration1 := &models.H264CodecConfiguration{
		Bitrate: &codecConfigVideoBitrate1,
		Width:   &codecConfigVideoWidth1,
		Profile: bitmovintypes.H264ProfileHigh,
	}
	videoCodecConfiguration1Resp, err := h264S1.Create(videoCodecConfiguration1)
	if err != nil {
		panic(err)
	}
	codecConfigVideoName2 := "Getting Started H264 Codec Config 2"
	codecConfigVideoBitrate2 := int64(1000000)
	codecConfigVideoWidth2 := int64(768)
	h264S2 := services.NewH264CodecConfigurationService(bitmovin)
	videoCodecConfiguration2 := &models.H264CodecConfiguration{
		Name:    &codecConfigVideoName2,
		Bitrate: &codecConfigVideoBitrate2,
		Width:   &codecConfigVideoWidth2,
		Profile: bitmovintypes.H264ProfileHigh,
	}
	videoCodecConfiguration2Resp, err := h264S2.Create(videoCodecConfiguration2)
	if err != nil {
		panic(err)
	}
	codecConfigVideoName3 := "Getting Started H264 Codec Config 3"
	codecConfigVideoBitrate3 := int64(750000)
	codecConfigVideoWidth3 := int64(640)
	h264S3 := services.NewH264CodecConfigurationService(bitmovin)
	videoCodecConfiguration3 := &models.H264CodecConfiguration{
		Name:    &codecConfigVideoName3,
		Bitrate: &codecConfigVideoBitrate3,
		Width:   &codecConfigVideoWidth3,
		Profile: bitmovintypes.H264ProfileHigh,
	}

	videoCodecConfiguration3Resp, err := h264S3.Create(videoCodecConfiguration3)
	if err != nil {
		panic(err)
	}

	name = "Getting Started Audio Codec Config"
	bitrate := int64(128000)
	aacS := services.NewAACCodecConfigurationService(bitmovin)
	aacConfig := &models.AACCodecConfiguration{
		Name:    &name,
		Bitrate: &bitrate,
	}
	aacResp, err := aacS.Create(aacConfig)
	if err != nil {
		panic(err)
	}

	//エンコーディングの作成（引数：Name, CloudRegion）
	name = "Getting Started Encoding(nomoto)"
	encodingS := services.NewEncodingService(bitmovin)
	encoding := &models.Encoding{
		Name:        &name,
		CloudRegion: bitmovintypes.CloudRegionGoogleUSEast1,
	}
	encodingResp, err := encodingS.Create(encoding)
	if err != nil {
		panic(err)
	}
	inputPath := "INPUT_sea/sea.mp4"

	videoInputStream1 := models.InputStream{
		InputID:       inputResp.Data.Result.ID,
		InputPath:     &inputPath,
		SelectionMode: bitmovintypes.SelectionModeAuto,
	}

	vis1 := []models.InputStream{videoInputStream1}
	videoStream1 := &models.Stream{
		CodecConfigurationID: videoCodecConfiguration1Resp.Data.Result.ID,
		InputStreams:         vis1,
	}
	videoStream1Resp, err := encodingS.AddStream(*encodingResp.Data.Result.ID, videoStream1)
	if err != nil {
		panic(err)
	}

	videoInputStream2 := models.InputStream{
		InputID:       inputResp.Data.Result.ID,
		InputPath:     &inputPath,
		SelectionMode: bitmovintypes.SelectionModeAuto,
	}

	vis2 := []models.InputStream{videoInputStream2}
	videoStream2 := &models.Stream{
		CodecConfigurationID: videoCodecConfiguration2Resp.Data.Result.ID,
		InputStreams:         vis2,
	}
	videoStream2Resp, err := encodingS.AddStream(*encodingResp.Data.Result.ID, videoStream2)
	if err != nil {
		panic(err)
	}

	videoInputStream3 := models.InputStream{
		InputID:       inputResp.Data.Result.ID,
		InputPath:     &inputPath,
		SelectionMode: bitmovintypes.SelectionModeAuto,
	}

	vis3 := []models.InputStream{videoInputStream3}
	videoStream3 := &models.Stream{
		CodecConfigurationID: videoCodecConfiguration3Resp.Data.Result.ID,
		InputStreams:         vis3,
	}
	videoStream3Resp, err := encodingS.AddStream(*encodingResp.Data.Result.ID, videoStream3)
	if err != nil {
		panic(err)
	}

	audioInputStream := models.InputStream{
		InputID:       inputResp.Data.Result.ID,
		InputPath:     &inputPath,
		SelectionMode: bitmovintypes.SelectionModeAuto,
	}

	ais := []models.InputStream{audioInputStream}
	audioStream := &models.Stream{
		CodecConfigurationID: aacResp.Data.Result.ID,
		InputStreams:         ais,
	}
	aacStreamResp, err := encodingS.AddStream(*encodingResp.Data.Result.ID, audioStream)
	if err != nil {
		panic(err)
	}

	//FMP4にてMuxingの作成
	aclEntry := models.ACLItem{
		Permission: bitmovintypes.ACLPermissionPrivate,
	}
	acl := []models.ACLItem{aclEntry}

	segmentLength := 4.0
	outputPath := "OUTPUT"
	segmentNaming := "seg_%number%.m4s"
	initSegmentName := "init.mp4"

	videoMuxingStream1 := models.StreamItem{
		StreamID: videoStream1Resp.Data.Result.ID,
	}

	outputPath1 := outputPath + "/video/1024_1500000/fmp4"
	videoMuxingOutput1 := models.Output{
		OutputID:   outputId,
		OutputPath: &outputPath1,
		ACL:        acl,
	}

	videoMuxing1 := &models.FMP4Muxing{
		SegmentLength:   &segmentLength,
		SegmentNaming:   &segmentNaming,
		InitSegmentName: &initSegmentName,
		Streams:         []models.StreamItem{videoMuxingStream1},
		Outputs:         []models.Output{videoMuxingOutput1},
	}

	videoMuxing1Resp, err := encodingS.AddFMP4Muxing(*encodingResp.Data.Result.ID, videoMuxing1)
	if err != nil {
		panic(err)
	}

	videoMuxingStream2 := models.StreamItem{
		StreamID: videoStream2Resp.Data.Result.ID,
	}

	outputPath2 := outputPath + "/video/768_1000000/fmp4"
	videoMuxingOutput2 := models.Output{
		OutputID:   outputId,
		OutputPath: &outputPath2,
		ACL:        acl,
	}

	videoMuxing2 := &models.FMP4Muxing{
		SegmentLength:   &segmentLength,
		SegmentNaming:   &segmentNaming,
		InitSegmentName: &initSegmentName,
		Streams:         []models.StreamItem{videoMuxingStream2},
		Outputs:         []models.Output{videoMuxingOutput2},
	}

	videoMuxing2Resp, err := encodingS.AddFMP4Muxing(*encodingResp.Data.Result.ID, videoMuxing2)
	if err != nil {
		panic(err)
	}

	videoMuxingStream3 := models.StreamItem{
		StreamID: videoStream3Resp.Data.Result.ID,
	}

	outputPath3 := outputPath + "/video/640_750000/fmp4"
	videoMuxingOutput3 := models.Output{
		OutputID:   outputId,
		OutputPath: &outputPath3,
		ACL:        acl,
	}

	videoMuxing3 := &models.FMP4Muxing{
		SegmentLength:   &segmentLength,
		SegmentNaming:   &segmentNaming,
		InitSegmentName: &initSegmentName,
		Streams:         []models.StreamItem{videoMuxingStream3},
		Outputs:         []models.Output{videoMuxingOutput3},
	}

	videoMuxing3Resp, err := encodingS.AddFMP4Muxing(*encodingResp.Data.Result.ID, videoMuxing3)
	if err != nil {
		panic(err)
	}

	audioMuxingStream := models.StreamItem{
		StreamID: aacStreamResp.Data.Result.ID,
	}

	outputPathFmp4 := outputPath + "/audio/128000/fmp4"
	audioMuxingFmp4Output := models.Output{
		OutputID:   outputId,
		OutputPath: &outputPathFmp4,
		ACL:        acl,
	}

	audioFMP4Muxing := &models.FMP4Muxing{
		SegmentLength:   &segmentLength,
		SegmentNaming:   &segmentNaming,
		InitSegmentName: &initSegmentName,
		Streams:         []models.StreamItem{audioMuxingStream},
		Outputs:         []models.Output{audioMuxingFmp4Output},
	}
	audioFMP4MuxingResp, err := encodingS.AddFMP4Muxing(*encodingResp.Data.Result.ID, audioFMP4Muxing)
	if err != nil {
		panic(err)
	}

	//Muxing(TS)
	// aclEntry := models.ACLItem{
	// 	Permission: bitmovintypes.ACLPermissionPublicRead,
	// }
	// acl := []models.ACLItem{aclEntry}

	// segmentLength := 4.0
	// outputPath := "/OUTPUT"
	// segmentNaming := "seg_%number%.ts"

	// videoMuxingStream1 := models.StreamItem{
	// 	StreamID: videoStream1Resp.Data.Result.ID,
	// }

	// outputPathTs1 := outputPath + "/video/1024_1500000/ts"
	// videoMuxingOutput1 := models.Output{
	// 	OutputID:   &outputId,
	// 	OutputPath: &outputPathTs1,
	// 	ACL:        acl,
	// }

	// videoMuxing1 := &models.TSMuxing{
	// 	SegmentLength: &segmentLength,
	// 	SegmentNaming: &segmentNaming,
	// 	Streams:       []models.StreamItem{videoMuxingStream1},
	// 	Outputs:       []models.Output{videoMuxingOutput1},
	// }

	// videoMuxing1Resp, err := encodingS.AddTSMuxing(*encodingResp.Data.Result.ID, videoMuxing1)
	// if err != nil {
	// 	panic(nil)
	// }

	// videoMuxingStream2 := models.StreamItem{
	// 	StreamID: videoStream2Resp.Data.Result.ID,
	// }

	// outputPathTs2 := outputPath + "/video/768_1000000/ts"
	// videoMuxingOutput2 := models.Output{
	// 	OutputID:   &outputId,
	// 	OutputPath: &outputPathTs2,
	// 	ACL:        acl,
	// }

	// videoMuxing2 := &models.TSMuxing{
	// 	SegmentLength: &segmentLength,
	// 	SegmentNaming: &segmentNaming,
	// 	Streams:       []models.StreamItem{videoMuxingStream2},
	// 	Outputs:       []models.Output{videoMuxingOutput2},
	// }

	// videoMuxing2Resp, err := encodingS.AddTSMuxing(*encodingResp.Data.Result.ID, videoMuxing2)
	// if err != nil {
	// 	panic(nil)
	// }

	// videoMuxingStream3 := models.StreamItem{
	// 	StreamID: videoStream3Resp.Data.Result.ID,
	// }

	// outputPathTs3 := outputPath + "/video/640_750000/ts"
	// videoMuxingOutput3 := models.Output{
	// 	OutputID:   &outputId,
	// 	OutputPath: &outputPathTs3,
	// 	ACL:        acl,
	// }

	// videoMuxing3 := &models.TSMuxing{
	// 	SegmentLength: &segmentLength,
	// 	SegmentNaming: &segmentNaming,
	// 	Streams:       []models.StreamItem{videoMuxingStream3},
	// 	Outputs:       []models.Output{videoMuxingOutput3},
	// }

	// videoMuxing3Resp, err := encodingS.AddTSMuxing(*encodingResp.Data.Result.ID, videoMuxing3)
	// if err != nil {
	// 	panic(nil)
	// }

	// audioMuxingStream := models.StreamItem{
	// 	StreamID: aacStreamResp.Data.Result.ID,
	// }

	// outputPathTs := outputPath + "/audio/128000/ts"
	// audioMuxingTsOutput := models.Output{
	// 	OutputID:   &outputId,
	// 	OutputPath: &outputPathTs,
	// 	ACL:        acl,
	// }

	// audioTSMuxing := &models.TSMuxing{
	// 	SegmentLength: &segmentLength,
	// 	SegmentNaming: &segmentNaming,
	// 	Streams:       []models.StreamItem{audioMuxingStream},
	// 	Outputs:       []models.Output{audioMuxingTsOutput},
	// }
	// audioTSMuxingResp, err := encodingS.AddTSMuxing(*encodingResp.Data.Result.ID, audioTSMuxing)
	// if err != nil {
	// 	panic(nil)
	// }

	//エンコード（引数：encodingResp.Data.Result.ID）
	_, err = encodingS.Start(*encodingResp.Data.Result.ID)
	if err != nil {
		panic(err)
	}

	//Step 9: Create a Manifest(DASH)
	manifestOutput := models.Output{
		OutputID:   outputId,
		OutputPath: &outputPath,
		ACL:        acl,
	}

	name = "Getting Started Manifest"
	manifestName := "manifest.mpd"
	manifest := &models.DashManifest{
		Name:         &name,
		ManifestName: &manifestName,
		Outputs:      []models.Output{manifestOutput},
	}

	service := services.NewDashManifestService(bitmovin)
	manifestResp, err := service.Create(manifest)
	if err != nil {
		panic(err)
	}

	period := &models.Period{}
	periodResp, err := service.AddPeriod(*manifestResp.Data.Result.ID, period)
	if err != nil {
		panic(err)
	}

	vas := &models.VideoAdaptationSet{}
	vasResp, err := service.AddVideoAdaptationSet(*manifestResp.Data.Result.ID, *periodResp.Data.Result.ID, vas)
	if err != nil {
		panic(err)
	}

	language := "en"
	aas := &models.AudioAdaptationSet{
		Language: &language,
	}
	aasResp, err := service.AddAudioAdaptationSet(*manifestResp.Data.Result.ID, *periodResp.Data.Result.ID, aas)
	if err != nil {
		panic(err)
	}

	segmentPathAudio := "audio/128000/fmp4"
	fmp4RepAudio := &models.FMP4Representation{
		Type:        bitmovintypes.FMP4RepresentationTypeTemplate,
		MuxingID:    audioFMP4MuxingResp.Data.Result.ID,
		EncodingID:  encodingResp.Data.Result.ID,
		SegmentPath: &segmentPathAudio,
	}
	_, err = service.AddFMP4Representation(*manifestResp.Data.Result.ID, *periodResp.Data.Result.ID, *aasResp.Data.Result.ID, fmp4RepAudio)
	if err != nil {
		panic(err)
	}

	segmentPathVideo1 := "video/1024_1500000/fmp4"
	fmp4Rep1 := &models.FMP4Representation{
		Type:        bitmovintypes.FMP4RepresentationTypeTemplate,
		MuxingID:    videoMuxing1Resp.Data.Result.ID,
		EncodingID:  encodingResp.Data.Result.ID,
		SegmentPath: &segmentPathVideo1,
	}
	_, err = service.AddFMP4Representation(*manifestResp.Data.Result.ID, *periodResp.Data.Result.ID, *vasResp.Data.Result.ID, fmp4Rep1)
	if err != nil {
		panic(err)
	}

	segmentPathVideo2 := "video/768_1000000/fmp4"
	fmp4Rep2 := &models.FMP4Representation{
		Type:        bitmovintypes.FMP4RepresentationTypeTemplate,
		MuxingID:    videoMuxing2Resp.Data.Result.ID,
		EncodingID:  encodingResp.Data.Result.ID,
		SegmentPath: &segmentPathVideo2,
	}
	_, err = service.AddFMP4Representation(*manifestResp.Data.Result.ID, *periodResp.Data.Result.ID, *vasResp.Data.Result.ID, fmp4Rep2)
	if err != nil {
		panic(err)
	}

	segmentPathVideo3 := "video/640_750000/fmp4"
	fmp4Rep3 := &models.FMP4Representation{
		Type:        bitmovintypes.FMP4RepresentationTypeTemplate,
		MuxingID:    videoMuxing3Resp.Data.Result.ID,
		EncodingID:  encodingResp.Data.Result.ID,
		SegmentPath: &segmentPathVideo3,
	}
	_, err = service.AddFMP4Representation(*manifestResp.Data.Result.ID, *periodResp.Data.Result.ID, *vasResp.Data.Result.ID, fmp4Rep3)
	if err != nil {
		panic(err)
	}

	_, err = service.Start(*manifestResp.Data.Result.ID)
	if err != nil {
		panic(err)
	}

	//Step 9: Create a Manifest(DASH)
	// manifestOutput := models.Output{
	// 	OutputID:   &outputId,
	// 	OutputPath: &outputPath,
	// 	ACL:        acl,
	// }

	// name = "Getting Started Manifest"
	// manifestName := "manifest.m3u8"
	// manifest := &models.HLSManifest{
	// 	Name:         &name,
	// 	ManifestName: &manifestName,
	// 	Outputs:      []models.Output{manifestOutput},
	// }

	// service := services.NewHLSManifestService(bitmovin)

	// manifestResp, err := service.Create(manifest)
	// if err != nil {
	// 	panic(nil)
	// }

	// uriAudio := "audiomedia.m3u8"
	// groupIDAudio := "audio_group"
	// languageAudio := "en"
	// nameAudio := "HLS Audio Media"
	// segmentPathAudio := "audio/128000/ts"
	// audioMediaInfo := &models.MediaInfo{
	// 	Type:        bitmovintypes.MediaTypeAudio,
	// 	URI:         &uriAudio,
	// 	GroupID:     &groupIDAudio,
	// 	Language:    &languageAudio,
	// 	Name:        &nameAudio,
	// 	SegmentPath: &segmentPathAudio,
	// 	EncodingID:  encodingResp.Data.Result.ID,
	// 	StreamID:    aacStreamResp.Data.Result.ID,
	// 	MuxingID:    audioTSMuxingResp.Data.Result.ID,
	// }

	// _, err = service.AddMediaInfo(*manifestResp.Data.Result.ID, audioMediaInfo)
	// if err != nil {
	// 	panic(nil)
	// }

	// audio := "audio_group"

	// segmentPathVideo1 := "video/1024_1500000/ts"
	// uriVideo1 := "video1.m3u8"
	// videoStreamInfo1 := &models.StreamInfo{
	// 	Audio:       &audio,
	// 	SegmentPath: &segmentPathVideo1,
	// 	URI:         &uriVideo1,
	// 	EncodingID:  encodingResp.Data.Result.ID,
	// 	StreamID:    videoStream1Resp.Data.Result.ID,
	// 	MuxingID:    videoMuxing1Resp.Data.Result.ID,
	// }
	// _, err = service.AddStreamInfo(*manifestResp.Data.Result.ID, videoStreamInfo1)
	// if err != nil {
	// 	panic(nil)
	// }

	// segmentPathVideo2 := "video/768_1000000/ts"
	// uriVideo2 := "video2.m3u8"
	// videoStreamInfo2 := &models.StreamInfo{
	// 	Audio:       &audio,
	// 	SegmentPath: &segmentPathVideo2,
	// 	URI:         &uriVideo2,
	// 	EncodingID:  encodingResp.Data.Result.ID,
	// 	StreamID:    videoStream2Resp.Data.Result.ID,
	// 	MuxingID:    videoMuxing2Resp.Data.Result.ID,
	// }
	// _, err = service.AddStreamInfo(*manifestResp.Data.Result.ID, videoStreamInfo2)
	// if err != nil {
	// 	panic(nil)
	// }

	// segmentPathVideo3 := "video/640_750000/ts"
	// uriVideo3 := "video3.m3u8"
	// videoStreamInfo3 := &models.StreamInfo{
	// 	Audio:       &audio,
	// 	SegmentPath: &segmentPathVideo3,
	// 	URI:         &uriVideo3,
	// 	EncodingID:  encodingResp.Data.Result.ID,
	// 	StreamID:    videoStream3Resp.Data.Result.ID,
	// 	MuxingID:    videoMuxing3Resp.Data.Result.ID,
	// }
	// _, err = service.AddStreamInfo(*manifestResp.Data.Result.ID, videoStreamInfo3)
	// if err != nil {
	// 	panic(nil)
	// }

	// _, err = service.Start(*manifestResp.Data.Result.ID)
	// if err != nil {
	// 	panic(nil)
	// }
}
