# Migration Guide - upgrading to ibm-cos-sdk-go-config@2.0.0

## Breaking Changes

- [Changes in Go Mod](#changes-in-go-mod)
- [Changes in Package Import](#changes-in-package-import)
- [Changes in API Patch Models](#changes-in-api-patch-models)
- [Changes in Agent Name](#changes-in-agent-name)

------------------------------------------------------------------------------------------------------------------------------------------------

### Changes in Go Mod

```js
|=================================================================================================================|
|                  current(v1.2.1)             |                                new(v2.0.0)                       |
|=================================================================================================================|
| github.com/IBM/ibm-cos-sdk-go-config v1.2.1  |    github.com/IBM/ibm-cos-sdk-go-config/v2 v2.0.0                |
|                                              |                                                                  |
|=================================================================================================================|
```

### Changes in Package Import

There is a change in the import statement and a new package needs to be imported as below.

```js
|==================================================================================================================================|
|                  current(v1.2.1)                             |                                new(v2.0.0)                        |
|==================================================================================================================================|
|github.com/IBM/ibm-cos-sdk-go-config/resourceconfigurationv1  | github.com/IBM/ibm-cos-sdk-go-config/v2/resourceconfigurationv1   |
|                                                              |                                                                   |
|==================================================================================================================================|
```

### Changes in API Patch Models

All patches must be supplied as `BucketPatch` with the map `map[string]interface`.

### Activity Tracking

```js

  |============================================================================================================================================|
  |                    current(v1.2.1)                               |                                new(v2.0.0)                              |
  |============================================================================================================================================|
  |                                                                  |   patchNameMap := make(map [string]interface{})                         |
  |  updateBucketConfigOptions := &rc.UpdateBucketConfigOptions{     |   patchNameMap["activity_tracking"] = &rc.ActivityTracking{             |
  |      Bucket: core.StringPtr("BucketName"),                       |       ReadDataEvents:     core.BoolPtr(bool),                           |
  |      ActivityTracking: &rc.ActivityTracking{                     |       WriteDataEvents:    core.BoolPtr(bool),                           |
  |        ReadDataEvents:     core.BoolPtr(bool),                   |       ActivityTrackerCrn: core.StringPtr(activityTrackerCrn),           |
  |        WriteDataEvents:    core.BoolPtr(bool),                   |   }                                                                     |
  |        ActivityTrackerCrn: core.StringPtr(activityTrackerCrn),   |                                                                         |
  |    },                                                            |   updateBucketConfigOptions := &rc.UpdateBucketConfigOptions{           |
  |  }                                                               |       Bucket: core.StringPtr("BucketName"),                             |
  |                                                                  |       BucketPatch: patchNameMap,                                        |
  |  res, e := client.UpdateBucketConfig(updateBucketConfigOptions)  |   }                                                                     |
  |                                                                  |   res, e := client.UpdateBucketConfig(updateBucketConfigOptions)        |
  |                                                                  |                                                                         |
  |============================================================================================================================================|

```

### Hard Quota

```js

  |============================================================================================================================================|
  |                    current(v1.2.1)                               |                                new(v2.0.0)                              |
  |============================================================================================================================================|
  |                                                                  |   bucketPatchModel := new(rc.BucketPatch)                               |
  |  setOptions.SetBucket(suite.bucketName)                          |   bucketPatchModel.HardQuota = core.Int64Ptr(int64(1024))               |
  |  setOptions.SetBucket(suite.bucketName)                          |   bucketPatchModelAsPatch, asPatchErr := bucketPatchModel.AsPatch()     |
  |  setOptions.SetHardQuota(1024)                                   |   setOptions := new(rc.UpdateBucketConfigOptions)                       |
  |  res, e := suite.client.UpdateBucketConfig(setOptions)           |   setOptions.SetBucket(suite.bucketName)                                |
  |                                                                  |   setOptions.BucketPatch = bucketPatchModelAsPatch                      |
  |                                                                  |   res, e := suite.client.UpdateBucketConfig(setOptions)                 |
  |                                                                  |                                                                         |
  |============================================================================================================================================|

```

### Metrics Monitoring

```js

  |============================================================================================================================================|
  |                    current(v1.2.1)                               |                                new(v2.0.0)                              |
  |============================================================================================================================================|
  |                                                                  |   patchNameMap := make(map [string]interface{})                         |
  |  updateBucketConfigOptions := &rc.UpdateBucketConfigOptions{     |   patchNameMap["metrics_monitoring"] = &rc.MetricsMonitoring{           |
  |      Bucket: core.StringPtr("BucketName"),                       |       MetricsMonitoringCrn: core.StringPtr(sysdigCrn),                  |
  |      MetricsMonitoring: &rc.MetricsMonitoring{                   |       UsageMetricsEnabled:  core.BoolPtr(bool),                         |
  |        MetricsMonitoringCrn: core.StringPtr(sysdigCrn),          |   }                                                                     |
  |        UsageMetricsEnabled:  core.BoolPtr(bool),                 |                                                                         |
  |    },                                                            |   updateBucketConfigOptions := &rc.UpdateBucketConfigOptions{           |
  |  }                                                               |       Bucket: core.StringPtr("BucketName"),                             |
  |                                                                  |       BucketPatch: patchNameMap,                                        |
  |  res, e := client.UpdateBucketConfig(updateBucketConfigOptions)  |   }                                                                     |
  |                                                                  |   res, e := client.UpdateBucketConfig(updateBucketConfigOptions)        |
  |                                                                  |                                                                         |
  |============================================================================================================================================|

```

### Firewall

```js
  |============================================================================================================================================|
  |                    current(v1.2.1)                               |                                new(v2.0.0)                              |
  |============================================================================================================================================|
  |                                                                  |   patchNameMap := make(map [string]interface{})                         |
  |  updateBucketConfigOptions := &rc.UpdateBucketConfigOptions{     |   patchNameMap["firewall"] = &rc.Firewall{                              |
  |      Bucket: core.StringPtr("BucketName"),                       |       AllowedIp: []string{"IPAddres"},                                  |
  |      Firewall: &rc.Firewall{                                     |   }                                                                     |
  |        AllowedIp: []string{"IPAddress"},                         |                                                                         |
  |    },                                                            |   updateBucketConfigOptions := &rc.UpdateBucketConfigOptions{           |
  |  }                                                               |       Bucket: core.StringPtr("BucketName"),                             |
  |                                                                  |       BucketPatch: patchNameMap,                                        |
  |  res, e := client.UpdateBucketConfig(updateBucketConfigOptions)  |   }                                                                     |
  |                                                                  |   res, e := client.UpdateBucketConfig(updateBucketConfigOptions)        |
  |                                                                  |                                                                         |
  |============================================================================================================================================|

```

### Changes in Agent Name

```js

  |======================================================================|
  |      current(v1.2.1)             |           new(v2.0.0)             |
  |======================================================================|
  | ibm-cos-resource-config-sdk-go   |     ibm-cos-sdk-go-config         |
  |======================================================================|

```

------------------------------------------------------------------------------------------------------------------------------------------------
