#ifndef HIDL_TYPES_H_
#define HIDL_TYPES_H_

#include <android/hardware/drm/1.1/types.h>
#include <android/hardware/drm/1.2/ICryptoPlugin.h>
#include <android/hardware/drm/1.2/IDrmPlugin.h>
#include <android/hardware/drm/1.2/IDrmPluginListener.h>
#include <android/hardware/drm/1.2/types.h>
#include <android/hardware/drm/1.3/ICryptoFactory.h>
#include <android/hardware/drm/1.3/IDrmFactory.h>
#include <android/hardware/drm/1.4/ICryptoFactory.h>
#include <android/hardware/drm/1.4/ICryptoPlugin.h>
#include <android/hardware/drm/1.4/IDrmFactory.h>
#include <android/hardware/drm/1.4/IDrmPlugin.h>
#include <hidl/HidlTransportSupport.h>

using ::android::hardware::configureRpcThreadpool;
using ::android::hardware::hidl_array;
using ::android::hardware::hidl_handle;
using ::android::hardware::hidl_memory;
using ::android::hardware::hidl_string;
using ::android::hardware::hidl_vec;
using ::android::hardware::joinRpcThreadpool;
using ::android::hardware::Return;
using ::android::hardware::Void;
using ::android::sp;

namespace drm = ::android::hardware::drm;
using drm::V1_0::BufferType;
using drm::V1_0::DestinationBuffer;
using drm::V1_0::EventType;
using drm::V1_0::IDrmPluginListener;
using drm::V1_0::KeyRequestType;
using drm::V1_0::KeyStatus;
using drm::V1_0::KeyStatusType;
using drm::V1_0::KeyType;
using drm::V1_0::KeyValue;
using drm::V1_0::Mode;
using drm::V1_0::Pattern;
using drm::V1_0::SecureStop;
using drm::V1_0::SecureStopId;
using drm::V1_0::SharedBuffer;
using drm::V1_0::Status;
using drm::V1_0::SubSample;
using drm::V1_1::DrmMetricGroup;
using drm::V1_1::HdcpLevel;
using drm::V1_1::SecureStopRelease;
using drm::V1_1::SecurityLevel;
using drm::V1_2::ICryptoPlugin;
using drm::V1_2::IDrmPlugin;
using drm::V1_2::KeySetId;
using drm::V1_2::OfflineLicenseState;
using drm::V1_3::ICryptoFactory;
using drm::V1_3::IDrmFactory;

typedef drm::V1_1::KeyRequestType KeyRequestType_V1_1;
typedef drm::V1_2::KeyStatus KeyStatus_V1_2;
typedef drm::V1_2::KeyStatusType KeyStatusType_V1_2;
typedef drm::V1_2::IDrmPluginListener IDrmPluginListener_V1_2;
typedef drm::V1_0::Status Status_V1_0;
typedef drm::V1_2::Status Status_V1_2;
typedef drm::V1_2::HdcpLevel HdcpLevel_V1_2;

#endif
