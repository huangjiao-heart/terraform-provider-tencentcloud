---
subcategory: "Media Processing Service(MPS)"
layout: "tencentcloud"
page_title: "TencentCloud: tencentcloud_mps_schedule"
sidebar_current: "docs-tencentcloud-resource-mps_schedule"
description: |-
  Provides a resource to create a mps schedule
---

# tencentcloud_mps_schedule

Provides a resource to create a mps schedule

## Example Usage

### Create a schedule through COS bucket

```hcl
data "tencentcloud_cos_bucket_object" "object" {
  bucket = "keep-bucket-${local.app_id}"
  key    = "/mps-test/test.mov"
}

resource "tencentcloud_cos_bucket" "output" {
  bucket      = "tf-bucket-mps-schedule-output-${local.app_id}"
  force_clean = true
  acl         = "public-read"
}

resource "tencentcloud_mps_schedule" "schedule" {
  schedule_name = "tf_test_mps_schedule_%d"

  trigger {
    type = "CosFileUpload"
    cos_file_upload_trigger {
      bucket  = data.tencentcloud_cos_bucket_object.object.bucket
      region  = "%s"
      dir     = "/upload/"
      formats = ["flv", "mov"]
    }
  }

  activities {
    activity_type   = "input"
    reardrive_index = [1, 2]
  }

  activities {
    activity_type   = "action-trans"
    reardrive_index = [3]
    activity_para {
      transcode_task {
        definition = 10
      }
    }
  }

  activities {
    activity_type   = "action-trans"
    reardrive_index = [6, 7]
    activity_para {
      transcode_task {
        definition = 10
      }
    }
  }

  activities {
    activity_type   = "action-trans"
    reardrive_index = [4, 5]
    activity_para {
      transcode_task {
        definition = 10
      }
    }
  }

  activities {
    activity_type   = "action-trans"
    reardrive_index = [10]
    activity_para {
      transcode_task {
        definition = 10
      }
    }
  }

  activities {
    activity_type   = "action-trans"
    reardrive_index = [10]
    activity_para {
      transcode_task {
        definition = 10
      }
    }
  }

  activities {
    activity_type   = "action-trans"
    reardrive_index = [10]
    activity_para {
      transcode_task {
        definition = 10
      }
    }
  }

  activities {
    activity_type   = "action-trans"
    reardrive_index = [8]
    activity_para {
      transcode_task {
        definition = 10
      }
    }
  }

  activities {
    activity_type   = "action-trans"
    reardrive_index = [9]
    activity_para {
      transcode_task {
        definition = 10
      }
    }
  }

  activities {
    activity_type   = "action-trans"
    reardrive_index = [10]
    activity_para {
      transcode_task {
        definition = 10
      }
    }
  }

  activities {
    activity_type = "output"
  }

  output_storage {
    type = "COS"
    cos_output_storage {
      bucket = tencentcloud_cos_bucket.output.bucket
      region = "%s"
    }
  }

  output_dir = "output/"
}
```

## Argument Reference

The following arguments are supported:

* `activities` - (Required, List) The subtasks of the scheme.
* `schedule_name` - (Required, String) The scheme name (max 128 characters). This name should be unique across your account.
* `trigger` - (Required, List) The trigger of the scheme. If a file is uploaded to the specified bucket, the scheme will be triggered.
* `output_dir` - (Optional, String) The directory to save the media processing output file, which must start and end with `/`, such as `/movie/201907/`.If you do not specify this, the file will be saved to the trigger directory.
* `output_storage` - (Optional, List) The bucket to save the output file. If you do not specify this parameter, the bucket in `Trigger` will be used.
* `task_notify_config` - (Optional, List) The notification configuration. If you do not specify this parameter, notifications will not be sent.

The `activities` object supports the following:

* `activity_type` - (Required, String) The subtask type. `input`: The start. `output`: The end. `action-trans`: Transcoding. `action-samplesnapshot`: Sampled screencapturing. `action-AIAnalysis`: Content analysis. `action-AIRecognition`: Content recognition. `action-aiReview`: Content moderation. `action-animated-graphics`: Animated screenshot generation. `action-image-sprite`: Image sprite generation. `action-snapshotByTimeOffset`: Time point screencapturing. `action-adaptive-substream`: Adaptive bitrate streaming.Note: This field may return null, indicating that no valid values can be obtained.
* `activity_para` - (Optional, List) The parameters of a subtask.Note: This field may return null, indicating that no valid values can be obtained.
* `reardrive_index` - (Optional, Set) The indexes of the subsequent actions. Note: This field may return null, indicating that no valid values can be obtained.

The `activity_para` object of `activities` supports the following:

* `adaptive_dynamic_streaming_task` - (Optional, List) An adaptive bitrate streaming task.
* `ai_analysis_task` - (Optional, List) A content analysis task.
* `ai_content_review_task` - (Optional, List) A content moderation task.
* `ai_recognition_task` - (Optional, List) A content recognition task.
* `animated_graphic_task` - (Optional, List) An animated screenshot generation task.
* `image_sprite_task` - (Optional, List) An image sprite generation task.
* `sample_snapshot_task` - (Optional, List) A sampled screencapturing task.
* `snapshot_by_time_offset_task` - (Optional, List) A time point screencapturing task.
* `transcode_task` - (Optional, List) A transcoding task.

The `adaptive_dynamic_streaming_task` object of `activity_para` supports the following:

* `definition` - (Required, Int) Adaptive bitrate streaming template ID.
* `add_on_subtitles` - (Optional, List) Subtitle files to insert.Note: This field may return null, indicating that no valid value can be obtained.
* `output_object_path` - (Optional, String) The relative or absolute output path of the manifest file after being transcoded to adaptive bitrate streaming. If this parameter is left empty, a relative path in the following format will be used by default: `{inputName}_adaptiveDynamicStreaming_{definition}.{format}`.
* `output_storage` - (Optional, List) Target bucket of an output file after being transcoded to adaptive bitrate streaming. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: this field may return null, indicating that no valid values can be obtained.
* `segment_object_name` - (Optional, String) The relative output path of the segment file after being transcoded to adaptive bitrate streaming (in HLS format only). If this parameter is left empty, a relative path in the following format will be used by default: `{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}_{segmentNumber}.{format}`.
* `sub_stream_object_name` - (Optional, String) The relative output path of the substream file after being transcoded to adaptive bitrate streaming. If this parameter is left empty, a relative path in the following format will be used by default: `{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}.{format}`.
* `watermark_set` - (Optional, List) List of up to 10 image or text watermarks.

The `add_on_subtitles` object of `adaptive_dynamic_streaming_task` supports the following:

* `subtitle` - (Optional, List) Subtitle file.Note: This field may return null, indicating that no valid value can be obtained.
* `type` - (Optional, String) The inserting type. Valid values: subtitle-stream:Insert title track close-caption-708:CEA-708 subtitle encode to SEI frame close-caption-608:CEA-608 subtitle encode to SEI frameNote: This field may return null, indicating that no valid value can be obtained.

The `add_on_subtitles` object of `override_parameter` supports the following:

* `subtitle` - (Optional, List) Subtitle file.Note: This field may return null, indicating that no valid value can be obtained.
* `type` - (Optional, String) The inserting type. Valid values: `subtitle-stream`:Insert title track. `close-caption-708`:CEA-708 subtitle encode to SEI frame. `close-caption-608`:CEA-608 subtitle encode to SEI frame. Note: This field may return null, indicating that no valid value can be obtained.

The `addon_audio_stream` object of `override_parameter` supports the following:

* `type` - (Required, String) The input type. Valid values: `COS`: A COS bucket address.  `URL`: A URL.  `AWS-S3`: An AWS S3 bucket address. Currently, this type is only supported for transcoding tasks.
* `cos_input_info` - (Optional, List) The information of the COS object to process. This parameter is valid and required when `Type` is `COS`.
* `s3_input_info` - (Optional, List) The information of the AWS S3 object processed. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.
* `url_input_info` - (Optional, List) The URL of the object to process. This parameter is valid and required when `Type` is `URL`.Note: This field may return null, indicating that no valid value can be obtained.

The `ai_analysis_task` object of `activity_para` supports the following:

* `definition` - (Required, Int) Video content analysis template ID.
* `extended_parameter` - (Optional, String) An extended parameter, whose value is a stringfied JSON.Note: This parameter is for customers with special requirements. It needs to be customized offline.Note: This field may return null, indicating that no valid values can be obtained.

The `ai_content_review_task` object of `activity_para` supports the following:

* `definition` - (Required, Int) Video content audit template ID.

The `ai_recognition_task` object of `activity_para` supports the following:

* `definition` - (Required, Int) Intelligent video recognition template ID.

The `animated_graphic_task` object of `activity_para` supports the following:

* `definition` - (Required, Int) Animated image generating template ID.
* `end_time_offset` - (Required, Float64) End time of an animated image in a video in seconds.
* `start_time_offset` - (Required, Float64) Start time of an animated image in a video in seconds.
* `output_object_path` - (Optional, String) Output path to a generated animated image file, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_animatedGraphic_{definition}.{format}`.
* `output_storage` - (Optional, List) Target bucket of a generated animated image file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: This field may return null, indicating that no valid values can be obtained.

The `audio_template` object of `override_parameter` supports the following:

* `audio_channel` - (Optional, Int) Audio channel system. Valid values: 1: Mono 2: Dual 6: StereoWhen the media is packaged in audio format (FLAC, OGG, MP3, M4A), the sound channel cannot be set to stereo.
* `bitrate` - (Optional, Int) Audio stream bitrate in Kbps. Value range: 0 and [26, 256]. If the value is 0, the bitrate of the audio stream will be the same as that of the original audio.
* `codec` - (Optional, String) Audio stream codec.When the outer `Container` parameter is `mp3`, the valid value is: libmp3lame.When the outer `Container` parameter is `ogg` or `flac`, the valid value is: flac.When the outer `Container` parameter is `m4a`, the valid values include: libfdk_aac; libmp3lame; ac3.When the outer `Container` parameter is `mp4` or `flv`, the valid values include: libfdk_aac: More suitable for mp4; libmp3lame: More suitable for flv; mp2.When the outer `Container` parameter is `hls`, the valid values include: libfdk_aac; libmp3lame.
* `sample_rate` - (Optional, Int) Audio stream sample rate. Valid values: 32,000 44,100 48,000In Hz.
* `stream_selects` - (Optional, Set) The audio tracks to retain. All audio tracks are retained by default.

The `audio_template` object of `raw_parameter` supports the following:

* `bitrate` - (Required, Int) Audio stream bitrate in Kbps. Value range: 0 and [26, 256].If the value is 0, the bitrate of the audio stream will be the same as that of the original audio.
* `codec` - (Required, String) Audio stream codec.When the outer `Container` parameter is `mp3`, the valid value is: libmp3lame.When the outer `Container` parameter is `ogg` or `flac`, the valid value is: flac.When the outer `Container` parameter is `m4a`, the valid values include: libfdk_aac; libmp3lame; ac3.When the outer `Container` parameter is `mp4` or `flv`, the valid values include: libfdk_aac: more suitable for mp4; libmp3lame: more suitable for flv.When the outer `Container` parameter is `hls`, the valid values include: libfdk_aac; libmp3lame.
* `sample_rate` - (Required, Int) Audio stream sample rate. Valid values: 32,000 44,100 48,000In Hz.
* `audio_channel` - (Optional, Int) Audio channel system. Valid values: 1: Mono 2: Dual 6: StereoWhen the media is packaged in audio format (FLAC, OGG, MP3, M4A), the sound channel cannot be set to stereo.Default value: 2.

The `aws_s3_file_upload_trigger` object of `trigger` supports the following:

* `s3_bucket` - (Required, String) The AWS S3 bucket bound to the scheme.
* `s3_region` - (Required, String) The region of the AWS S3 bucket.
* `aws_sqs` - (Optional, List) The SQS queue of the AWS S3 bucket.Note: The queue must be in the same region as the bucket.Note: This field may return null, indicating that no valid values can be obtained.
* `dir` - (Optional, String) The bucket directory bound. It must be an absolute path that starts and ends with `/`, such as `/movie/201907/`. If you do not specify this, the root directory will be bound.	.
* `formats` - (Optional, Set) The file formats that will trigger the scheme, such as [mp4, flv, mov]. If you do not specify this, the upload of files in any format will trigger the scheme.	.
* `s3_secret_id` - (Optional, String) The key ID of the AWS S3 bucket.Note: This field may return null, indicating that no valid values can be obtained.
* `s3_secret_key` - (Optional, String) The key of the AWS S3 bucket.Note: This field may return null, indicating that no valid values can be obtained.

The `aws_sqs` object of `aws_s3_file_upload_trigger` supports the following:

* `sqs_queue_name` - (Required, String) The name of the SQS queue.
* `sqs_region` - (Required, String) The region of the SQS queue.
* `s3_secret_id` - (Optional, String) The key ID required to read from/write to the SQS queue.
* `s3_secret_key` - (Optional, String) The key required to read from/write to the SQS queue.

The `aws_sqs` object of `task_notify_config` supports the following:

* `sqs_queue_name` - (Required, String) The name of the SQS queue.
* `sqs_region` - (Required, String) The region of the SQS queue.
* `s3_secret_id` - (Optional, String) The key ID required to read from/write to the SQS queue.
* `s3_secret_key` - (Optional, String) The key required to read from/write to the SQS queue.

The `cos_file_upload_trigger` object of `trigger` supports the following:

* `bucket` - (Required, String) Name of the COS bucket bound to a workflow, such as `TopRankVideo-125xxx88`.
* `region` - (Required, String) Region of the COS bucket bound to a workflow, such as `ap-chongiqng`.
* `dir` - (Optional, String) Input path directory bound to a workflow, such as `/movie/201907/`. If this parameter is left empty, the `/` root directory will be used.
* `formats` - (Optional, Set) Format list of files that can trigger a workflow, such as [mp4, flv, mov]. If this parameter is left empty, files in all formats can trigger the workflow.

The `cos_input_info` object of `addon_audio_stream` supports the following:

* `bucket` - (Required, String) The COS bucket of the object to process, such as `TopRankVideo-125xxx88`.
* `object` - (Required, String) The path of the object to process, such as `/movie/201907/WildAnimal.mov`.
* `region` - (Required, String) The region of the COS bucket, such as `ap-chongqing`.

The `cos_input_info` object of `head_set` supports the following:

* `bucket` - (Required, String) The COS bucket of the object to process, such as `TopRankVideo-125xxx88`.
* `object` - (Required, String) The path of the object to process, such as `/movie/201907/WildAnimal.mov`.
* `region` - (Required, String) The region of the COS bucket, such as `ap-chongqing`.

The `cos_input_info` object of `image_content` supports the following:

* `bucket` - (Required, String) The COS bucket of the object to process, such as `TopRankVideo-125xxx88`.
* `object` - (Required, String) The path of the object to process, such as `/movie/201907/WildAnimal.mov`.
* `region` - (Required, String) The region of the COS bucket, such as `ap-chongqing`.

The `cos_input_info` object of `subtitle` supports the following:

* `bucket` - (Required, String) The COS bucket of the object to process, such as TopRankVideo-125xxx88.
* `object` - (Required, String) The path of the object to process, such as /movie/201907/WildAnimal.mov.
* `region` - (Required, String) The region of the COS bucket, such as ap-chongqing.

The `cos_input_info` object of `tail_set` supports the following:

* `bucket` - (Required, String) The COS bucket of the object to process, such as `TopRankVideo-125xxx88`.
* `object` - (Required, String) The path of the object to process, such as `/movie/201907/WildAnimal.mov`.
* `region` - (Required, String) The region of the COS bucket, such as `ap-chongqing`.

The `cos_output_storage` object of `output_storage` supports the following:

* `bucket` - (Optional, String) The bucket to which the output file of media processing is saved, such as `TopRankVideo-125xxx88`. If this parameter is left empty, the value of the upper layer will be inherited.
* `region` - (Optional, String) The region of the output bucket, such as `ap-chongqing`. If this parameter is left empty, the value of the upper layer will be inherited.

The `head_set` object of `head_tail_parameter` supports the following:

* `type` - (Required, String) The input type. Valid values: `COS`: A COS bucket address.  `URL`: A URL.  `AWS-S3`: An AWS S3 bucket address. Currently, this type is only supported for transcoding tasks.
* `cos_input_info` - (Optional, List) The information of the COS object to process. This parameter is valid and required when `Type` is `COS`.
* `s3_input_info` - (Optional, List) The information of the AWS S3 object processed. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.
* `url_input_info` - (Optional, List) The URL of the object to process. This parameter is valid and required when `Type` is `URL`.Note: This field may return null, indicating that no valid value can be obtained.

The `head_tail_parameter` object of `transcode_task` supports the following:

* `head_set` - (Optional, List) Opening credits list.
* `tail_set` - (Optional, List) Closing credits list.

The `image_content` object of `image_template` supports the following:

* `type` - (Required, String) The input type. Valid values: `COS`: A COS bucket address.  `URL`: A URL.  `AWS-S3`: An AWS S3 bucket address. Currently, this type is only supported for transcoding tasks.
* `cos_input_info` - (Optional, List) The information of the COS object to process. This parameter is valid and required when `Type` is `COS`.
* `s3_input_info` - (Optional, List) The information of the AWS S3 object processed. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.
* `url_input_info` - (Optional, List) The URL of the object to process. This parameter is valid and required when `Type` is `URL`.Note: This field may return null, indicating that no valid value can be obtained.

The `image_sprite_task` object of `activity_para` supports the following:

* `definition` - (Required, Int) ID of an image sprite generating template.
* `object_number_format` - (Optional, List) Rule of the `{number}` variable in the image sprite output path.Note: This field may return null, indicating that no valid values can be obtained.
* `output_object_path` - (Optional, String) Output path to a generated image sprite file, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_imageSprite_{definition}_{number}.{format}`.
* `output_storage` - (Optional, List) Target bucket of a generated image sprite. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: This field may return null, indicating that no valid values can be obtained.
* `web_vtt_object_name` - (Optional, String) Output path to the WebVTT file after an image sprite is generated, which can only be a relative path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_imageSprite_{definition}.{format}`.

The `image_template` object of `raw_parameter` supports the following:

* `image_content` - (Required, List) Input content of watermark image. JPEG and PNG images are supported.
* `height` - (Optional, String) Watermark height. % and px formats are supported: If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height; If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px.Default value: 0 px, which means that `Height` will be proportionally scaled according to the aspect ratio of the original watermark image.
* `repeat_type` - (Optional, String) Repeat type of an animated watermark. Valid values: `once`: no longer appears after watermark playback ends. `repeat_last_frame`: stays on the last frame after watermark playback ends. `repeat` (default): repeats the playback until the video ends.
* `width` - (Optional, String) Watermark width. % and px formats are supported: If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width; If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px.Default value: 10%.

The `mosaic_set` object of `transcode_task` supports the following:

* `coordinate_origin` - (Optional, String) Origin position, which currently can only be: TopLeft: the origin of coordinates is in the top-left corner of the video, and the origin of the blur is in the top-left corner of the image or text.Default value: TopLeft.
* `end_time_offset` - (Optional, Float64) End time offset of blur in seconds. If this parameter is left empty or 0 is entered, the blur will exist till the last video frame; If this value is greater than 0 (e.g., n), the blur will exist till second n; If this value is smaller than 0 (e.g., -n), the blur will exist till second n before the last video frame.
* `height` - (Optional, String) Blur height. % and px formats are supported: If the string ends in %, the `Height` of the blur will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height; If the string ends in px, the `Height` of the blur will be in px; for example, `100px` means that `Height` is 100 px.Default value: 10%.
* `start_time_offset` - (Optional, Float64) Start time offset of blur in seconds. If this parameter is left empty or 0 is entered, the blur will appear upon the first video frame. If this parameter is left empty or 0 is entered, the blur will appear upon the first video frame; If this value is greater than 0 (e.g., n), the blur will appear at second n after the first video frame; If this value is smaller than 0 (e.g., -n), the blur will appear at second n before the last video frame.
* `width` - (Optional, String) Blur width. % and px formats are supported: If the string ends in %, the `Width` of the blur will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width; If the string ends in px, the `Width` of the blur will be in px; for example, `100px` means that `Width` is 100 px.Default value: 10%.
* `x_pos` - (Optional, String) The horizontal position of the origin of the blur relative to the origin of coordinates of the video. % and px formats are supported: If the string ends in %, the `XPos` of the blur will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width; If the string ends in px, the `XPos` of the blur will be the specified px; for example, `100px` means that `XPos` is 100 px.Default value: 0 px.
* `y_pos` - (Optional, String) Vertical position of the origin of blur relative to the origin of coordinates of video. % and px formats are supported: If the string ends in %, the `YPos` of the blur will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height; If the string ends in px, the `YPos` of the blur will be the specified px; for example, `100px` means that `YPos` is 100 px.Default value: 0 px.

The `object_number_format` object of `image_sprite_task` supports the following:

* `increment` - (Optional, Int) Increment of the `{number}` variable. Default value: 1.
* `initial_value` - (Optional, Int) Start value of the `{number}` variable. Default value: 0.
* `min_length` - (Optional, Int) Minimum length of the `{number}` variable. A placeholder will be used if the variable length is below the minimum requirement. Default value: 1.
* `place_holder` - (Optional, String) Placeholder used when the `{number}` variable length is below the minimum requirement. Default value: 0.

The `object_number_format` object of `sample_snapshot_task` supports the following:

* `increment` - (Optional, Int) Increment of the `{number}` variable. Default value: 1.
* `initial_value` - (Optional, Int) Start value of the `{number}` variable. Default value: 0.
* `min_length` - (Optional, Int) Minimum length of the `{number}` variable. A placeholder will be used if the variable length is below the minimum requirement. Default value: 1.
* `place_holder` - (Optional, String) Placeholder used when the `{number}` variable length is below the minimum requirement. Default value: 0.

The `object_number_format` object of `snapshot_by_time_offset_task` supports the following:

* `increment` - (Optional, Int) Increment of the `{number}` variable. Default value: 1.
* `initial_value` - (Optional, Int) Start value of the `{number}` variable. Default value: 0.
* `min_length` - (Optional, Int) Minimum length of the `{number}` variable. A placeholder will be used if the variable length is below the minimum requirement. Default value: 1.
* `place_holder` - (Optional, String) Placeholder used when the `{number}` variable length is below the minimum requirement. Default value: 0.

The `object_number_format` object of `transcode_task` supports the following:

* `increment` - (Optional, Int) Increment of the `{number}` variable. Default value: 1.
* `initial_value` - (Optional, Int) Start value of the `{number}` variable. Default value: 0.
* `min_length` - (Optional, Int) Minimum length of the `{number}` variable. A placeholder will be used if the variable length is below the minimum requirement. Default value: 1.
* `place_holder` - (Optional, String) Placeholder used when the `{number}` variable length is below the minimum requirement. Default value: 0.

The `output_storage` object of `adaptive_dynamic_streaming_task` supports the following:

* `type` - (Required, String) The storage type for a media processing output file. Valid values: `COS`: Tencent Cloud COS `AWS-S3`: AWS S3. This type is only supported for AWS tasks, and the output bucket must be in the same region as the bucket of the source file.
* `cos_output_storage` - (Optional, List) The location to save the output object in COS. This parameter is valid and required when `Type` is COS.Note: This field may return null, indicating that no valid value can be obtained.
* `s3_output_storage` - (Optional, List) The AWS S3 bucket to save the output file. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.

The `output_storage` object of `animated_graphic_task` supports the following:

* `type` - (Required, String) The storage type for a media processing output file. Valid values: `COS`: Tencent Cloud COS `AWS-S3`: AWS S3. This type is only supported for AWS tasks, and the output bucket must be in the same region as the bucket of the source file.
* `cos_output_storage` - (Optional, List) The location to save the output object in COS. This parameter is valid and required when `Type` is COS.Note: This field may return null, indicating that no valid value can be obtained.
* `s3_output_storage` - (Optional, List) The AWS S3 bucket to save the output file. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.

The `output_storage` object of `image_sprite_task` supports the following:

* `type` - (Required, String) The storage type for a media processing output file. Valid values: `COS`: Tencent Cloud COS `AWS-S3`: AWS S3. This type is only supported for AWS tasks, and the output bucket must be in the same region as the bucket of the source file.
* `cos_output_storage` - (Optional, List) The location to save the output object in COS. This parameter is valid and required when `Type` is COS.Note: This field may return null, indicating that no valid value can be obtained.
* `s3_output_storage` - (Optional, List) The AWS S3 bucket to save the output file. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.

The `output_storage` object of `sample_snapshot_task` supports the following:

* `type` - (Required, String) The storage type for a media processing output file. Valid values: `COS`: Tencent Cloud COS `AWS-S3`: AWS S3. This type is only supported for AWS tasks, and the output bucket must be in the same region as the bucket of the source file.
* `cos_output_storage` - (Optional, List) The location to save the output object in COS. This parameter is valid and required when `Type` is COS.Note: This field may return null, indicating that no valid value can be obtained.
* `s3_output_storage` - (Optional, List) The AWS S3 bucket to save the output file. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.

The `output_storage` object of `snapshot_by_time_offset_task` supports the following:

* `type` - (Required, String) The storage type for a media processing output file. Valid values: `COS`: Tencent Cloud COS `AWS-S3`: AWS S3. This type is only supported for AWS tasks, and the output bucket must be in the same region as the bucket of the source file.
* `cos_output_storage` - (Optional, List) The location to save the output object in COS. This parameter is valid and required when `Type` is COS.Note: This field may return null, indicating that no valid value can be obtained.
* `s3_output_storage` - (Optional, List) The AWS S3 bucket to save the output file. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.

The `output_storage` object of `transcode_task` supports the following:

* `type` - (Required, String) The storage type for a media processing output file. Valid values: `COS`: Tencent Cloud COS `AWS-S3`: AWS S3. This type is only supported for AWS tasks, and the output bucket must be in the same region as the bucket of the source file.
* `cos_output_storage` - (Optional, List) The location to save the output object in COS. This parameter is valid and required when `Type` is COS.Note: This field may return null, indicating that no valid value can be obtained.
* `s3_output_storage` - (Optional, List) The AWS S3 bucket to save the output file. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.

The `output_storage` object supports the following:

* `type` - (Required, String) The storage type for a media processing output file. Valid values: `COS`: Tencent Cloud COS `AWS-S3`: AWS S3. This type is only supported for AWS tasks, and the output bucket must be in the same region as the bucket of the source file.
* `cos_output_storage` - (Optional, List) The location to save the output object in COS. This parameter is valid and required when `Type` is COS.Note: This field may return null, indicating that no valid value can be obtained.
* `s3_output_storage` - (Optional, List) The AWS S3 bucket to save the output file. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.

The `override_parameter` object of `transcode_task` supports the following:

* `add_on_subtitles` - (Optional, List) Subtitle files to insert.Note: This field may return null, indicating that no valid value can be obtained.
* `addon_audio_stream` - (Optional, List) The information of the external audio track to add.Note: This field may return null, indicating that no valid values can be obtained.
* `audio_template` - (Optional, List) Audio stream configuration parameter.
* `container` - (Optional, String) Container format. Valid values: mp4, flv, hls, mp3, flac, ogg, and m4a; mp3, flac, ogg, and m4a are formats of audio files.
* `remove_audio` - (Optional, Int) Whether to remove audio data. Valid values: 0: retain 1: remove.
* `remove_video` - (Optional, Int) Whether to remove video data. Valid values: 0: retain 1: remove.
* `std_ext_info` - (Optional, String) Transcoding extension field.Note: This field may return null, indicating that no valid value can be obtained.
* `subtitle_template` - (Optional, List) The subtitle settings.
* `tehd_config` - (Optional, List) TESHD transcoding parameter.
* `video_template` - (Optional, List) Video stream configuration parameter.

The `raw_parameter` object of `transcode_task` supports the following:

* `container` - (Required, String) Container. Valid values: mp4; flv; hls; mp3; flac; ogg; m4a. Among them, mp3, flac, ogg, and m4a are for audio files.
* `audio_template` - (Optional, List) Audio stream configuration parameter. This field is required when `RemoveAudio` is 0.
* `remove_audio` - (Optional, Int) Whether to remove audio data. Valid values: 0: retain; 1: remove.Default value: 0.
* `remove_video` - (Optional, Int) Whether to remove video data. Valid values: 0: retain; 1: remove.Default value: 0.
* `tehd_config` - (Optional, List) TESHD transcoding parameter.
* `video_template` - (Optional, List) Video stream configuration parameter. This field is required when `RemoveVideo` is 0.

The `raw_parameter` object of `watermark_set` supports the following:

* `type` - (Required, String) Watermark type. Valid values: image: image watermark.
* `coordinate_origin` - (Optional, String) Origin position, which currently can only be: TopLeft: the origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text.Default value: TopLeft.
* `image_template` - (Optional, List) Image watermark template. This field is required when `Type` is `image` and is invalid when `Type` is `text`.
* `x_pos` - (Optional, String) The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported: If the string ends in %, the `XPos` of the watermark will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width; If the string ends in px, the `XPos` of the watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.Default value: 0 px.
* `y_pos` - (Optional, String) The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported: If the string ends in %, the `YPos` of the watermark will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height; If the string ends in px, the `YPos` of the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.Default value: 0 px.

The `s3_input_info` object of `addon_audio_stream` supports the following:

* `s3_bucket` - (Required, String) The AWS S3 bucket.
* `s3_object` - (Required, String) The path of the AWS S3 object.
* `s3_region` - (Required, String) The region of the AWS S3 bucket.
* `s3_secret_id` - (Optional, String) The key ID required to access the AWS S3 object.
* `s3_secret_key` - (Optional, String) The key required to access the AWS S3 object.

The `s3_input_info` object of `head_set` supports the following:

* `s3_bucket` - (Required, String) The AWS S3 bucket.
* `s3_object` - (Required, String) The path of the AWS S3 object.
* `s3_region` - (Required, String) The region of the AWS S3 bucket.
* `s3_secret_id` - (Optional, String) The key ID required to access the AWS S3 object.
* `s3_secret_key` - (Optional, String) The key required to access the AWS S3 object.

The `s3_input_info` object of `image_content` supports the following:

* `s3_bucket` - (Required, String) The AWS S3 bucket.
* `s3_object` - (Required, String) The path of the AWS S3 object.
* `s3_region` - (Required, String) The region of the AWS S3 bucket.
* `s3_secret_id` - (Optional, String) The key ID required to access the AWS S3 object.
* `s3_secret_key` - (Optional, String) The key required to access the AWS S3 object.

The `s3_input_info` object of `subtitle` supports the following:

* `s3_bucket` - (Required, String) S3 bucket.Note: This field may return null, indicating that no valid value can be obtained.
* `s3_object` - (Required, String) The path of the AWS S3 object.Note: This field may return null, indicating that no valid value can be obtained.
* `s3_region` - (Required, String) The region of the AWS S3 bucket, support:  us-east-1  eu-west-3Note: This field may return null, indicating that no valid value can be obtained.
* `s3_secret_id` - (Optional, String) The key ID required to access the AWS S3 object.Note: This field may return null, indicating that no valid value can be obtained.
* `s3_secret_key` - (Optional, String) The key required to access the AWS S3 object.Note: This field may return null, indicating that no valid value can be obtained.

The `s3_input_info` object of `tail_set` supports the following:

* `s3_bucket` - (Required, String) The AWS S3 bucket.
* `s3_object` - (Required, String) The path of the AWS S3 object.
* `s3_region` - (Required, String) The region of the AWS S3 bucket.
* `s3_secret_id` - (Optional, String) The key ID required to access the AWS S3 object.
* `s3_secret_key` - (Optional, String) The key required to access the AWS S3 object.

The `s3_output_storage` object of `output_storage` supports the following:

* `s3_bucket` - (Required, String) The AWS S3 bucket.
* `s3_region` - (Required, String) The region of the AWS S3 bucket.
* `s3_secret_id` - (Optional, String) The key ID required to upload files to the AWS S3 object.
* `s3_secret_key` - (Optional, String) The key required to upload files to the AWS S3 object.

The `sample_snapshot_task` object of `activity_para` supports the following:

* `definition` - (Required, Int) Sampled screencapturing template ID.
* `object_number_format` - (Optional, List) Rule of the `{number}` variable in the sampled screenshot output path.Note: This field may return null, indicating that no valid values can be obtained.
* `output_object_path` - (Optional, String) Output path to a generated sampled screenshot, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_sampleSnapshot_{definition}_{number}.{format}`.
* `output_storage` - (Optional, List) Target bucket of a sampled screenshot. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: This field may return null, indicating that no valid values can be obtained.
* `watermark_set` - (Optional, List) List of up to 10 image or text watermarks.Note: This field may return null, indicating that no valid values can be obtained.

The `snapshot_by_time_offset_task` object of `activity_para` supports the following:

* `definition` - (Required, Int) ID of a time point screencapturing template.
* `ext_time_offset_set` - (Optional, Set) List of screenshot time points in the format of `s` or `%`: If the string ends in `s`, it means that the time point is in seconds; for example, `3.5s` means that the time point is the 3.5th second; If the string ends in `%`, it means that the time point is the specified percentage of the video duration; for example, `10%` means that the time point is 10% of the video duration.
* `object_number_format` - (Optional, List) Rule of the `{number}` variable in the time point screenshot output path.Note: This field may return null, indicating that no valid values can be obtained.
* `output_object_path` - (Optional, String) Output path to a generated time point screenshot, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_snapshotByTimeOffset_{definition}_{number}.{format}`.
* `output_storage` - (Optional, List) Target bucket of a generated time point screenshot file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: This field may return null, indicating that no valid values can be obtained.
* `watermark_set` - (Optional, List) List of up to 10 image or text watermarks.Note: This field may return null, indicating that no valid values can be obtained.

The `subtitle_template` object of `override_parameter` supports the following:

* `font_alpha` - (Optional, Float64) The text transparency. Value range: 0-1. 0: Completely transparent 1: Completely opaqueDefault value: 1.
* `font_color` - (Optional, String) The font color in 0xRRGGBB format. Default value: 0xFFFFFF (white).
* `font_size` - (Optional, String) The font size (pixels). If this is not specified, the font size in the subtitle file will be used.
* `font_type` - (Optional, String) The font type. Valid values: `hei.ttf` `song.ttf` `simkai.ttf` `arial.ttf` (for English only). The default is `hei.ttf`.
* `path` - (Optional, String) The URL of the subtitles to add to the video.
* `stream_index` - (Optional, Int) The subtitle track to add to the video. If both `Path` and `StreamIndex` are specified, `Path` will be used. You need to specify at least one of the two parameters.

The `subtitle` object of `add_on_subtitles` supports the following:

* `type` - (Required, String) The input type. Valid values:  COS:A COS bucket address  URL:A URL  AWS-S3:An AWS S3 bucket address. Currently, this type is only supported for transcoding tasks.
* `cos_input_info` - (Optional, List) The information of the COS object to process. This parameter is valid and required when Type is COS.
* `s3_input_info` - (Optional, List) The information of the AWS S3 object processed. This parameter is required if Type is AWS-S3. Note: This field may return null, indicating that no valid value can be obtained.
* `url_input_info` - (Optional, List) The URL of the object to process. This parameter is valid and required when Type is URL.Note: This field may return null, indicating that no valid value can be obtained.

The `subtitle` object of `add_on_subtitles` supports the following:

* `type` - (Required, String) The input type. Valid values:  `COS`:A COS bucket address. `URL`:A URL. `AWS-S3`:An AWS S3 bucket address. Currently, this type is only supported for transcoding tasks.
* `cos_input_info` - (Optional, List) The information of the COS object to process. This parameter is valid and required when Type is COS.
* `s3_input_info` - (Optional, List) The information of the AWS S3 object processed. This parameter is required if Type is AWS-S3.Note: This field may return null, indicating that no valid value can be obtained.
* `url_input_info` - (Optional, List) The URL of the object to process. This parameter is valid and required when Type is URL.Note: This field may return null, indicating that no valid value can be obtained.

The `tail_set` object of `head_tail_parameter` supports the following:

* `type` - (Required, String) The input type. Valid values: `COS`: A COS bucket address.  `URL`: A URL.  `AWS-S3`: An AWS S3 bucket address. Currently, this type is only supported for transcoding tasks.
* `cos_input_info` - (Optional, List) The information of the COS object to process. This parameter is valid and required when `Type` is `COS`.
* `s3_input_info` - (Optional, List) The information of the AWS S3 object processed. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.
* `url_input_info` - (Optional, List) The URL of the object to process. This parameter is valid and required when `Type` is `URL`.Note: This field may return null, indicating that no valid value can be obtained.

The `task_notify_config` object supports the following:

* `aws_sqs` - (Optional, List) The AWS SQS queue. This parameter is required if `NotifyType` is `AWS-SQS`.Note: This field may return null, indicating that no valid values can be obtained.
* `cmq_model` - (Optional, String) The CMQ or TDMQ-CMQ model. Valid values: Queue, Topic.
* `cmq_region` - (Optional, String) The CMQ or TDMQ-CMQ region, such as `sh` (Shanghai) or `bj` (Beijing).
* `notify_mode` - (Optional, String) Workflow notification method. Valid values: Finish, Change. If this parameter is left empty, `Finish` will be used.
* `notify_type` - (Optional, String) The notification type. Valid values: `CMQ`: This value is no longer used. Please use `TDMQ-CMQ` instead. `TDMQ-CMQ`: Message queue `URL`: If `NotifyType` is set to `URL`, HTTP callbacks are sent to the URL specified by `NotifyUrl`. HTTP and JSON are used for the callbacks. The packet contains the response parameters of the `ParseNotification` API. `SCF`: This notification type is not recommended. You need to configure it in the SCF console. `AWS-SQS`: AWS queue. This type is only supported for AWS tasks, and the queue must be in the same region as the AWS bucket.Note: If you do not pass this parameter or pass in an empty string, `CMQ` will be used. To use a different notification type, specify this parameter accordingly.
* `notify_url` - (Optional, String) HTTP callback URL, required if `NotifyType` is set to `URL`.
* `queue_name` - (Optional, String) The CMQ or TDMQ-CMQ queue to receive notifications. This parameter is valid when `CmqModel` is `Queue`.
* `topic_name` - (Optional, String) The CMQ or TDMQ-CMQ topic to receive notifications. This parameter is valid when `CmqModel` is `Topic`.

The `tehd_config` object of `override_parameter` supports the following:

* `max_video_bitrate` - (Optional, Int) Maximum bitrate. If this parameter is left empty, no modification will be made.
* `type` - (Optional, String) TESHD type. Valid values: TEHD-100: TESHD-100.If this parameter is left blank, no modification will be made.

The `tehd_config` object of `raw_parameter` supports the following:

* `type` - (Required, String) TESHD type. Valid values: TEHD-100: TESHD-100.If this parameter is left empty, TESHD will not be enabled.
* `max_video_bitrate` - (Optional, Int) Maximum bitrate, which is valid when `Type` is `TESHD`.If this parameter is left empty or 0 is entered, there will be no upper limit for bitrate.

The `transcode_task` object of `activity_para` supports the following:

* `definition` - (Required, Int) ID of a video transcoding template.
* `end_time_offset` - (Optional, Float64) End time offset of a transcoded video, in seconds. If this parameter is left empty or set to 0, the transcoded video will end at the same time as the original video. If this parameter is set to a positive number (n for example), the transcoded video will end at the nth second of the original video. If this parameter is set to a negative number (-n for example), the transcoded video will end at the nth second before the end of the original video.
* `head_tail_parameter` - (Optional, List) Opening and closing credits parametersNote: this field may return `null`, indicating that no valid value was found.
* `mosaic_set` - (Optional, List) List of blurs. Up to 10 ones can be supported.
* `object_number_format` - (Optional, List) Rule of the `{number}` variable in the output path after transcoding.Note: This field may return null, indicating that no valid values can be obtained.
* `output_object_path` - (Optional, String) Path to a primary output file, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_transcode_{definition}.{format}`.
* `output_storage` - (Optional, List) Target bucket of an output file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: This field may return null, indicating that no valid values can be obtained.
* `override_parameter` - (Optional, List) Video transcoding custom parameter, which is valid when `Definition` is not 0.When any parameters in this structure are entered, they will be used to override corresponding parameters in templates.This parameter is used in highly customized scenarios. We recommend you only use `Definition` to specify the transcoding parameter.Note: this field may return `null`, indicating that no valid value was found.
* `raw_parameter` - (Optional, List) Custom video transcoding parameter, which is valid if `Definition` is 0.This parameter is used in highly customized scenarios. We recommend you use `Definition` to specify the transcoding parameter preferably.
* `segment_object_name` - (Optional, String) Path to an output file part (the path to ts during transcoding to HLS), which can only be a relative path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_transcode_{definition}_{number}.{format}`.
* `start_time_offset` - (Optional, Float64) Start time offset of a transcoded video, in seconds. If this parameter is left empty or set to 0, the transcoded video will start at the same time as the original video. If this parameter is set to a positive number (n for example), the transcoded video will start at the nth second of the original video. If this parameter is set to a negative number (-n for example), the transcoded video will start at the nth second before the end of the original video.
* `watermark_set` - (Optional, List) List of up to 10 image or text watermarks.Note: This field may return null, indicating that no valid values can be obtained.

The `trigger` object supports the following:

* `type` - (Required, String) The trigger type. Valid values: `CosFileUpload`: Tencent Cloud COS trigger. `AwsS3FileUpload`: AWS S3 trigger. Currently, this type is only supported for transcoding tasks and schemes (not supported for workflows).
* `aws_s3_file_upload_trigger` - (Optional, List) The AWS S3 trigger. This parameter is valid and required if `Type` is `AwsS3FileUpload`.Note: Currently, the key for the AWS S3 bucket, the trigger SQS queue, and the callback SQS queue must be the same.Note: This field may return null, indicating that no valid values can be obtained.
* `cos_file_upload_trigger` - (Optional, List) This parameter is required and valid when `Type` is `CosFileUpload`, indicating the COS trigger rule.Note: This field may return null, indicating that no valid values can be obtained.

The `url_input_info` object of `addon_audio_stream` supports the following:

* `url` - (Required, String) URL of a video.

The `url_input_info` object of `head_set` supports the following:

* `url` - (Required, String) URL of a video.

The `url_input_info` object of `image_content` supports the following:

* `url` - (Required, String) URL of a video.

The `url_input_info` object of `subtitle` supports the following:

* `url` - (Required, String) URL of a video.

The `url_input_info` object of `tail_set` supports the following:

* `url` - (Required, String) URL of a video.

The `video_template` object of `override_parameter` supports the following:

* `bitrate` - (Optional, Int) Bitrate of a video stream in Kbps. Value range: 0 and [128, 35,000].If the value is 0, the bitrate of the video will be the same as that of the source video.
* `codec` - (Optional, String) The video codec. Valid values: libx264: H.264 libx265: H.265 av1: AOMedia Video 1Note: You must specify a resolution (not higher than 640 x 480) if the H.265 codec is used.Note: You can only use the AOMedia Video 1 codec for MP4 files.
* `content_adapt_stream` - (Optional, Int) Whether to enable adaptive encoding. Valid values: 0: Disable 1: EnableDefault value: 0. If this parameter is set to `1`, multiple streams with different resolutions and bitrates will be generated automatically. The highest resolution, bitrate, and quality of the streams are determined by the values of `width` and `height`, `Bitrate`, and `Vcrf` in `VideoTemplate` respectively. If these parameters are not set in `VideoTemplate`, the highest resolution generated will be the same as that of the source video, and the highest video quality will be close to VMAF 95. To use this parameter or learn about the billing details of adaptive encoding, please contact your sales rep.
* `fill_type` - (Optional, String) Fill type. Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported:  stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.
* `fps` - (Optional, Int) Video frame rate in Hz. Value range: [0, 100].If the value is 0, the frame rate will be the same as that of the source video.
* `gop` - (Optional, Int) Frame interval between I keyframes. Value range: 0 and [1,100000]. If this parameter is 0, the system will automatically set the GOP length.
* `height` - (Optional, Int) Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096].
* `resolution_adaptive` - (Optional, String) Resolution adaption. Valid values: open: Enabled. When resolution adaption is enabled, `Width` indicates the long side of a video, while `Height` indicates the short side. close: Disabled. When resolution adaption is disabled, `Width` indicates the width of a video, while `Height` indicates the height.Note: When resolution adaption is enabled, `Width` cannot be smaller than `Height`.
* `vcrf` - (Optional, Int) The control factor of video constant bitrate. Value range: [0, 51]. This parameter will be disabled if you enter `0`.It is not recommended to specify this parameter if there are no special requirements.
* `width` - (Optional, Int) Maximum value of the width (or long side) of a video stream in px. Value range: 0 and [128, 4,096]. If both `Width` and `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled; If both `Width` and `Height` are not 0, the custom resolution will be used.

The `video_template` object of `raw_parameter` supports the following:

* `bitrate` - (Required, Int) The video bitrate (Kbps). Value range: 0 and [128, 35000].If the value is 0, the bitrate of the video will be the same as that of the source video.
* `codec` - (Required, String) The video codec. Valid values: `libx264`: H.264 `libx265`: H.265 `av1`: AOMedia Video 1Note: You must specify a resolution (not higher than 640 x 480) if the H.265 codec is used.Note: You can only use the AOMedia Video 1 codec for MP4 files.
* `fps` - (Required, Int) The video frame rate (Hz). Value range: [0, 100].If the value is 0, the frame rate will be the same as that of the source video.Note: For adaptive bitrate streaming, the value range of this parameter is [0, 60].
* `fill_type` - (Optional, String) The fill mode, which indicates how a video is resized when the video's original aspect ratio is different from the target aspect ratio. Valid values: stretch: Stretch the image frame by frame to fill the entire screen. The video image may become squashed or stretched after transcoding. black: Keep the image&#39;s original aspect ratio and fill the blank space with black bars. white: Keep the image's original aspect ratio and fill the blank space with white bars. gauss: Keep the image's original aspect ratio and apply Gaussian blur to the blank space.Default value: black.Note: Only `stretch` and `black` are supported for adaptive bitrate streaming.
* `gop` - (Optional, Int) Frame interval between I keyframes. Value range: 0 and [1,100000].If this parameter is 0 or left empty, the system will automatically set the GOP length.
* `height` - (Optional, Int) Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096]. If both `Width` and `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled; If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
* `resolution_adaptive` - (Optional, String) Resolution adaption. Valid values: open: Enabled. When resolution adaption is enabled, `Width` indicates the long side of a video, while `Height` indicates the short side. close: Disabled. When resolution adaption is disabled, `Width` indicates the width of a video, while `Height` indicates the height.Default value: open.Note: When resolution adaption is enabled, `Width` cannot be smaller than `Height`.
* `vcrf` - (Optional, Int) The control factor of video constant bitrate. Value range: [1, 51]If this parameter is specified, CRF (a bitrate control method) will be used for transcoding. (Video bitrate will no longer take effect.)It is not recommended to specify this parameter if there are no special requirements.
* `width` - (Optional, Int) Maximum value of the width (or long side) of a video stream in px. Value range: 0 and [128, 4,096]. If both `Width` and `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled; If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.

The `watermark_set` object of `adaptive_dynamic_streaming_task` supports the following:

* `definition` - (Required, Int) ID of a watermarking template.
* `end_time_offset` - (Optional, Float64) End time offset of a watermark in seconds. If this parameter is left empty or 0 is entered, the watermark will exist till the last video frame; If this value is greater than 0 (e.g., n), the watermark will exist till second n; If this value is smaller than 0 (e.g., -n), the watermark will exist till second n before the last video frame.
* `raw_parameter` - (Optional, List) Custom watermark parameter, which is valid if `Definition` is 0.This parameter is used in highly customized scenarios. We recommend you use `Definition` to specify the watermark parameter preferably.Custom watermark parameter is not available for screenshot.
* `start_time_offset` - (Optional, Float64) Start time offset of a watermark in seconds. If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame. If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame; If this value is greater than 0 (e.g., n), the watermark will appear at second n after the first video frame; If this value is smaller than 0 (e.g., -n), the watermark will appear at second n before the last video frame.
* `svg_content` - (Optional, String) SVG content of up to 2,000,000 characters. This field is required only when the watermark type is `SVG`.SVG watermark is not available for screenshot.
* `text_content` - (Optional, String) Text content of up to 100 characters. This field is required only when the watermark type is text.Text watermark is not available for screenshot.

The `watermark_set` object of `sample_snapshot_task` supports the following:

* `definition` - (Required, Int) ID of a watermarking template.
* `end_time_offset` - (Optional, Float64) End time offset of a watermark in seconds. If this parameter is left empty or 0 is entered, the watermark will exist till the last video frame; If this value is greater than 0 (e.g., n), the watermark will exist till second n; If this value is smaller than 0 (e.g., -n), the watermark will exist till second n before the last video frame.
* `raw_parameter` - (Optional, List) Custom watermark parameter, which is valid if `Definition` is 0.This parameter is used in highly customized scenarios. We recommend you use `Definition` to specify the watermark parameter preferably.Custom watermark parameter is not available for screenshot.
* `start_time_offset` - (Optional, Float64) Start time offset of a watermark in seconds. If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame. If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame; If this value is greater than 0 (e.g., n), the watermark will appear at second n after the first video frame; If this value is smaller than 0 (e.g., -n), the watermark will appear at second n before the last video frame.
* `svg_content` - (Optional, String) SVG content of up to 2,000,000 characters. This field is required only when the watermark type is `SVG`.SVG watermark is not available for screenshot.
* `text_content` - (Optional, String) Text content of up to 100 characters. This field is required only when the watermark type is text.Text watermark is not available for screenshot.

The `watermark_set` object of `snapshot_by_time_offset_task` supports the following:

* `definition` - (Required, Int) ID of a watermarking template.
* `end_time_offset` - (Optional, Float64) End time offset of a watermark in seconds. If this parameter is left empty or 0 is entered, the watermark will exist till the last video frame; If this value is greater than 0 (e.g., n), the watermark will exist till second n; If this value is smaller than 0 (e.g., -n), the watermark will exist till second n before the last video frame.
* `raw_parameter` - (Optional, List) Custom watermark parameter, which is valid if `Definition` is 0.This parameter is used in highly customized scenarios. We recommend you use `Definition` to specify the watermark parameter preferably.Custom watermark parameter is not available for screenshot.
* `start_time_offset` - (Optional, Float64) Start time offset of a watermark in seconds. If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame. If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame; If this value is greater than 0 (e.g., n), the watermark will appear at second n after the first video frame; If this value is smaller than 0 (e.g., -n), the watermark will appear at second n before the last video frame.
* `svg_content` - (Optional, String) SVG content of up to 2,000,000 characters. This field is required only when the watermark type is `SVG`.SVG watermark is not available for screenshot.
* `text_content` - (Optional, String) Text content of up to 100 characters. This field is required only when the watermark type is text.Text watermark is not available for screenshot.

The `watermark_set` object of `transcode_task` supports the following:

* `definition` - (Required, Int) ID of a watermarking template.
* `end_time_offset` - (Optional, Float64) End time offset of a watermark in seconds. If this parameter is left empty or 0 is entered, the watermark will exist till the last video frame; If this value is greater than 0 (e.g., n), the watermark will exist till second n; If this value is smaller than 0 (e.g., -n), the watermark will exist till second n before the last video frame.
* `raw_parameter` - (Optional, List) Custom watermark parameter, which is valid if `Definition` is 0.This parameter is used in highly customized scenarios. We recommend you use `Definition` to specify the watermark parameter preferably.Custom watermark parameter is not available for screenshot.
* `start_time_offset` - (Optional, Float64) Start time offset of a watermark in seconds. If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame. If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame; If this value is greater than 0 (e.g., n), the watermark will appear at second n after the first video frame; If this value is smaller than 0 (e.g., -n), the watermark will appear at second n before the last video frame.
* `svg_content` - (Optional, String) SVG content of up to 2,000,000 characters. This field is required only when the watermark type is `SVG`.SVG watermark is not available for screenshot.
* `text_content` - (Optional, String) Text content of up to 100 characters. This field is required only when the watermark type is text.Text watermark is not available for screenshot.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

mps schedule can be imported using the id, e.g.

```
terraform import tencentcloud_mps_schedule.schedule schedule_id
```

