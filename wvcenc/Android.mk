LOCAL_PATH:= $(call my-dir)
include $(LOCAL_PATH)/oemcryptolevel.mk
ifneq ($(filter arm arm64, $(TARGET_ARCH)),)

include $(CLEAR_VARS)
ifeq ($(LOCAL_OEMCRYPTO_LEVEL),1)
LOCAL_MODULE := liboemcrypto
LOCAL_MULTILIB := both
LOCAL_LICENSE_KINDS := SPDX-license-identifier-Apache-2.0 SPDX-license-identifier-BSD SPDX-license-identifier-LGPL legacy_by_exception_only
LOCAL_LICENSE_CONDITIONS := by_exception_only notice restricted
LOCAL_MODULE_CLASS := SHARED_LIBRARIES
LOCAL_MODULE_SUFFIX := .so
LOCAL_SRC_FILES_32 := arm/$(LOCAL_MODULE)$(LOCAL_MODULE_SUFFIX)
LOCAL_SRC_FILES_64 := arm64/$(LOCAL_MODULE)$(LOCAL_MODULE_SUFFIX)
LOCAL_PROPRIETARY_MODULE := true
LOCAL_STRIP_MODULE := false
LOCAL_MODULE_TAGS := optional
LOCAL_SHARED_LIBRARIES := libcrypto libcutils liblog libteec libutils libz
include $(BUILD_PREBUILT)
endif

ifneq ($(USE_PRESIGNED_TA),true)
include $(CLEAR_VARS)
ifeq ($(LOCAL_OEMCRYPTO_LEVEL),1)

TA_UUID := e043cde0-61d0-11e5-9c26-0002a5d5c51b
TA_SUFFIX := .ta

ifeq ($(PLATFORM_TDK_VERSION), 38)
	PLATFORM_TDK_PATH := $(BOARD_AML_VENDOR_PATH)/tdk_v3
	ifeq ($(BOARD_AML_SOC_TYPE),)
		LOCAL_TA := ta/v3/signed/$(TA_UUID)$(TA_SUFFIX)
	else
		LOCAL_TA := ta/v3/dev/$(BOARD_AML_SOC_TYPE)/$(TA_UUID)$(TA_SUFFIX)
	endif
else
	PLATFORM_TDK_PATH := $(BOARD_AML_VENDOR_PATH)/tdk
	LOCAL_TA := ta/v2/signed/$(TA_UUID)$(TA_SUFFIX)
endif

LOCAL_SRC_FILES := $(LOCAL_TA)
LOCAL_MODULE := $(TA_UUID)
LOCAL_LICENSE_KINDS := SPDX-license-identifier-Apache-2.0 SPDX-license-identifier-BSD SPDX-license-identifier-LGPL legacy_by_exception_only
LOCAL_LICENSE_CONDITIONS := by_exception_only notice restricted
LOCAL_MODULE_SUFFIX := $(TA_SUFFIX)
LOCAL_STRIP_MODULE := false
LOCAL_MODULE_CLASS := SHARED_LIBRARIES
LOCAL_MODULE_PATH := $(TARGET_OUT_VENDOR)/lib/teetz
include $(BUILD_PREBUILT)

endif # LOCAL_OEMCRYPTO_LEVEL == 1
endif # USE_PRESIGNED_TA != true

endif # TARGET_ARCH == arm
