ifeq ($(wildcard $(TOP)/vendor/amlogic/common/miracast-hdcp2),)
LOCAL_PATH:= $(call my-dir)

ifneq ($(filter arm arm64, $(TARGET_ARCH)),)

include $(CLEAR_VARS)
LOCAL_MODULE := libstagefright_hdcp
LOCAL_MODULE_CLASS := SHARED_LIBRARIES
LOCAL_MODULE_SUFFIX := .so
LOCAL_SRC_FILES_32 := arm/$(LOCAL_MODULE)$(LOCAL_MODULE_SUFFIX)
LOCAL_SRC_FILES_64 := arm64/$(LOCAL_MODULE)$(LOCAL_MODULE_SUFFIX)
LOCAL_PROPRIETARY_MODULE := true
LOCAL_STRIP_MODULE := false
LOCAL_MODULE_TAGS := optional
LOCAL_SHARED_LIBRARIES := libc++ libc libcutils libdl liblog libm libteec libutils
LOCAL_MULTILIB := both
LOCAL_MODULE_PATH_32 := $(TARGET_OUT_VENDOR)/lib/
LOCAL_MODULE_PATH_64 := $(TARGET_OUT_VENDOR)/lib64/
LOCAL_LICENSE_KINDS := SPDX-license-identifier-Apache-2.0
LOCAL_LICENSE_CONDITIONS := by_exception_only notice restricted proprietary by_exception_only
LOCAL_NOTICE_FILE := $(LOCAL_PATH)/../LICENSE
include $(BUILD_PREBUILT)

include $(CLEAR_VARS)

TA_UUID := 807798e0-f011-11e5-a5fe-0002a5d5c51b
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
LOCAL_32_BIT_ONLY := true
LOCAL_MODULE_SUFFIX := $(TA_SUFFIX)
LOCAL_STRIP_MODULE := false
LOCAL_MODULE_CLASS := SHARED_LIBRARIES
LOCAL_MODULE_PATH := $(TARGET_OUT_VENDOR)/lib/teetz
LOCAL_LICENSE_KINDS := SPDX-license-identifier-Apache-2.0
LOCAL_LICENSE_CONDITIONS := by_exception_only notice restricted proprietary by_exception_only
LOCAL_NOTICE_FILE := $(LOCAL_PATH)/../LICENSE
include $(BUILD_PREBUILT)

endif # TARGET_ARCH == arm
endif



